package reanchor

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const defaultRegistryMode = "default_layered"

type normalizedRequest struct {
	ReanchorRequest
	projectRoot    string
	checkpointPath string
	includeCompat  bool
	maxReads       int
	hashAlgorithm  string
}

type fingerprint struct {
	value string
	typ   string
}

// Evaluate compares the registered Anchor PM files with the previous
// checkpoint and returns a deterministic reanchor decision. It never writes
// checkpoints or anchor documents.
func Evaluate(ctx context.Context, req ReanchorRequest) (ReanchorResult, error) {
	nreq, err := normalizeRequest(req)
	if err != nil {
		return ReanchorResult{}, err
	}

	checkedAt := time.Now().UTC().Format(time.RFC3339)
	result := ReanchorResult{
		SchemaVersion: ResultSchemaV1,
		Operation:     nreq.Operation,
		ThreadID:      nreq.Thread.ID,
		CheckedAt:     checkedAt,
		AnchorState:   AnchorStateUnchanged,
		PeriodicDue:   periodicDue(nreq.Operation, nreq.Conversation),
		ChangedLayers: []string{},
		FileStatuses:  []FileStatus{},
		RequiredReads: []RequiredRead{},
		BlockedBy:     []BlockedBy{},
		NextAction:    NextActionContinue,
		CheckpointUpdate: CheckpointUpdate{
			WriteAllowed: false,
			Path:         nreq.checkpointPath,
			State: Checkpoint{
				SchemaVersion: CheckpointSchemaV1,
				ProjectID:     filepath.Base(nreq.projectRoot),
				ThreadID:      nreq.Thread.ID,
				UpdatedAt:     checkedAt,
				RoundCount:    nreq.Conversation.RoundCount + 1,
				LastOperation: nreq.Operation,
				Anchors:       []CheckpointAnchor{},
			},
		},
	}

	registry, err := registryForRequest(nreq)
	if err != nil {
		return ReanchorResult{}, err
	}

	checkpoint, checkpointFound, err := loadCheckpoint(nreq)
	if err != nil {
		return ReanchorResult{}, err
	}
	prior := checkpointIndex(checkpoint)

	handoffNamed := makePathSet(nreq.Handoff.NamedFiles)
	currentAnchors := make([]CheckpointAnchor, 0, len(registry))
	readSeen := map[string]bool{}
	changedLayerSet := map[string]bool{}

	for _, entry := range registry {
		if err := ctx.Err(); err != nil {
			return ReanchorResult{}, err
		}
		entry = normalizeEntry(entry)
		status, anchor := evaluateFile(nreq, entry, prior, checkpointFound)
		result.FileStatuses = append(result.FileStatuses, status)
		if anchor.Fingerprint != "" {
			currentAnchors = append(currentAnchors, anchor)
		}
		if statusCountsAsChanged(status.Status) {
			changedLayerSet[entry.Layer] = true
		}

		if block, ok := blockingReason(entry, status); ok {
			result.BlockedBy = append(result.BlockedBy, block)
			continue
		}
		if requiredReadNeeded(entry, status, checkpointFound, handoffNamed) {
			addRequiredRead(&result.RequiredReads, readSeen, RequiredRead{
				Path:     status.Path,
				Layer:    status.Layer,
				Reason:   readReason(entry, status, checkpointFound, handoffNamed),
				ReadMode: ReadModeFull,
				Priority: readPriority(entry, status),
			})
		}
	}

	if len(result.RequiredReads) > nreq.maxReads {
		result.BlockedBy = append(result.BlockedBy, BlockedBy{
			Status: "required_reads_exceeded",
			Reason: fmt.Sprintf("required reads %d exceed max_required_reads %d", len(result.RequiredReads), nreq.maxReads),
		})
	}

	sort.Slice(currentAnchors, func(i, j int) bool {
		if currentAnchors[i].Layer == currentAnchors[j].Layer {
			return currentAnchors[i].Path < currentAnchors[j].Path
		}
		return currentAnchors[i].Layer < currentAnchors[j].Layer
	})
	result.CheckpointUpdate.State.Anchors = currentAnchors

	result.RequiredUpdates = closeoutUpdates(nreq)
	for _, update := range result.RequiredUpdates {
		changedLayerSet[update.Layer] = true
	}

	result.ChangedLayers = sortedKeys(changedLayerSet)
	result.AnchorState = decideAnchorState(result, checkpointFound)
	result.NextAction = decideNextAction(result, checkpointFound)
	result.CheckpointUpdate.WriteAllowed = result.AnchorState == AnchorStateUnchanged && len(result.RequiredReads) == 0 && len(result.BlockedBy) == 0
	result.MinimalChatOutput = minimalChatOutput(result, checkpointFound)

	return result, nil
}

func normalizeRequest(req ReanchorRequest) (normalizedRequest, error) {
	if req.SchemaVersion == "" {
		req.SchemaVersion = RequestSchemaV1
	}
	if req.SchemaVersion != RequestSchemaV1 {
		return normalizedRequest{}, fmt.Errorf("unsupported request schema_version %q", req.SchemaVersion)
	}
	if req.Operation == "" {
		req.Operation = OperationOrdinaryThreadStart
	}
	if req.Thread.ID == "" {
		return normalizedRequest{}, errors.New("thread.id is required")
	}
	if req.ProjectRoot == "" {
		return normalizedRequest{}, errors.New("project_root is required")
	}
	root := filepath.Clean(filepath.FromSlash(strings.ReplaceAll(req.ProjectRoot, "\\", "/")))
	if !filepath.IsAbs(root) {
		return normalizedRequest{}, fmt.Errorf("project_root must be absolute: %q", req.ProjectRoot)
	}
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return normalizedRequest{}, err
	}
	if req.Registry.Mode == "" {
		req.Registry.Mode = defaultRegistryMode
	}
	includeCompat := true
	if req.Options.IncludeCompatibilityFiles != nil {
		includeCompat = *req.Options.IncludeCompatibilityFiles
	}
	hashAlgorithm := req.Options.HashAlgorithm
	if hashAlgorithm == "" {
		hashAlgorithm = "sha256"
	}
	if hashAlgorithm != "sha256" {
		return normalizedRequest{}, fmt.Errorf("unsupported hash_algorithm %q", hashAlgorithm)
	}
	maxReads := req.Options.MaxRequiredReads
	if maxReads <= 0 {
		maxReads = 20
	}
	checkpointPath := req.Checkpoint.Path
	if checkpointPath == "" {
		checkpointPath = fmt.Sprintf("docs/anchor_pm/.state/%s.anchor_state.json", req.Thread.ID)
	}
	return normalizedRequest{
		ReanchorRequest: req,
		projectRoot:     absRoot,
		checkpointPath:  filepath.ToSlash(checkpointPath),
		includeCompat:   includeCompat,
		maxReads:        maxReads,
		hashAlgorithm:   hashAlgorithm,
	}, nil
}

func loadCheckpoint(req normalizedRequest) (*Checkpoint, bool, error) {
	if req.Checkpoint.State != nil {
		return req.Checkpoint.State, true, nil
	}
	abs, _, ok := resolveUnderRoot(req.projectRoot, req.checkpointPath)
	if !ok {
		return nil, false, fmt.Errorf("checkpoint path is outside project_root: %s", req.checkpointPath)
	}
	data, err := os.ReadFile(abs)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, false, nil
		}
		return nil, false, err
	}
	var checkpoint Checkpoint
	if err := json.Unmarshal(data, &checkpoint); err != nil {
		return nil, false, err
	}
	return &checkpoint, true, nil
}

func checkpointIndex(checkpoint *Checkpoint) map[string]CheckpointAnchor {
	index := map[string]CheckpointAnchor{}
	if checkpoint == nil {
		return index
	}
	for _, anchor := range checkpoint.Anchors {
		index[anchorKey(anchor.ID, anchor.Path)] = anchor
		index[anchorKey("", anchor.Path)] = anchor
	}
	return index
}

func evaluateFile(req normalizedRequest, entry RegistryFileEntry, prior map[string]CheckpointAnchor, checkpointFound bool) (FileStatus, CheckpointAnchor) {
	status := FileStatus{
		ID:     entry.ID,
		Layer:  entry.Layer,
		Path:   filepath.ToSlash(entry.Path),
		Status: FileStatusUnknown,
		Error:  nil,
	}
	anchor := CheckpointAnchor{
		ID:    entry.ID,
		Layer: entry.Layer,
		Path:  filepath.ToSlash(entry.Path),
	}

	absPath, relPath, ok := resolveUnderRoot(req.projectRoot, entry.Path)
	if !ok {
		msg := "path is outside project_root"
		status.Path = filepath.ToSlash(entry.Path)
		status.Status = FileStatusInvalidPath
		status.Error = &msg
		return status, anchor
	}
	status.Path = relPath
	anchor.Path = relPath

	if existingPathLeavesRoot(req.projectRoot, absPath) {
		msg := "resolved path leaves project_root"
		status.Status = FileStatusInvalidPath
		status.Error = &msg
		return status, anchor
	}

	info, err := os.Stat(absPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			if entry.Required {
				status.Status = FileStatusMissingRequired
			} else {
				status.Status = FileStatusMissingOptional
			}
			return status, anchor
		}
		msg := err.Error()
		status.Status = FileStatusUnreadable
		status.Error = &msg
		return status, anchor
	}
	if info.IsDir() {
		msg := "registered path is a directory, not a file"
		status.Status = FileStatusInvalidPath
		status.Error = &msg
		return status, anchor
	}

	fp, err := computeFingerprint(absPath, info, entry.CheckMode)
	if err != nil {
		msg := err.Error()
		status.Status = FileStatusUnreadable
		status.Error = &msg
		return status, anchor
	}
	if fp.value == "" || fp.typ == "unknown" {
		status.Status = FileStatusUnknown
		return status, anchor
	}
	anchor.Fingerprint = fp.value
	anchor.FingerprintType = fp.typ
	status.CurrentFingerprint = fp.value

	prev, ok := prior[anchorKey(entry.ID, relPath)]
	if !ok {
		prev, ok = prior[anchorKey("", relPath)]
	}
	if ok {
		status.PreviousFingerprint = prev.Fingerprint
	}
	if !checkpointFound || !ok || prev.Fingerprint == "" {
		status.Status = FileStatusUnknown
		return status, anchor
	}
	if prev.Fingerprint != fp.value {
		status.Status = FileStatusChanged
		return status, anchor
	}
	status.Status = FileStatusUnchanged
	return status, anchor
}

func computeFingerprint(path string, info fs.FileInfo, checkMode string) (fingerprint, error) {
	switch checkMode {
	case "", CheckModeHash:
		return hashFingerprint(path)
	case CheckModeVersion:
		if version, err := explicitVersion(path); err == nil && version != "" {
			return fingerprint{value: "version:" + version, typ: "version"}, nil
		}
		return hashFingerprint(path)
	case CheckModeMTimeSize:
		return fingerprint{
			value: fmt.Sprintf("mtime_size:%d:%d", info.ModTime().UTC().UnixNano(), info.Size()),
			typ:   "mtime_size",
		}, nil
	case CheckModeExists:
		return fingerprint{value: "exists:true", typ: "exists"}, nil
	default:
		return fingerprint{typ: "unknown"}, nil
	}
}

func hashFingerprint(path string) (fingerprint, error) {
	f, err := os.Open(path)
	if err != nil {
		return fingerprint{}, err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return fingerprint{}, err
	}
	return fingerprint{value: "sha256:" + hex.EncodeToString(h.Sum(nil)), typ: "sha256"}, nil
}

func explicitVersion(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	for _, raw := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(raw)
		for _, prefix := range []string{"Baseline id:", "Current coordination version:", "Version:"} {
			if strings.HasPrefix(line, prefix) {
				value := strings.TrimSpace(strings.TrimPrefix(line, prefix))
				value = strings.Trim(value, "` ")
				if value != "" {
					return value, nil
				}
			}
		}
	}
	return "", nil
}

func resolveUnderRoot(root, inputPath string) (absPath, relPath string, ok bool) {
	if inputPath == "" {
		return "", "", false
	}
	cleanInput := strings.ReplaceAll(inputPath, "\\", "/")
	var abs string
	if filepath.IsAbs(filepath.FromSlash(cleanInput)) {
		abs = filepath.Clean(filepath.FromSlash(cleanInput))
	} else {
		abs = filepath.Join(root, filepath.FromSlash(cleanInput))
	}
	abs, _ = filepath.Abs(abs)
	rel, err := filepath.Rel(root, abs)
	if err != nil || rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) || filepath.IsAbs(rel) {
		return abs, filepath.ToSlash(rel), false
	}
	return abs, filepath.ToSlash(rel), true
}

func existingPathLeavesRoot(root, absPath string) bool {
	rootEval, err := filepath.EvalSymlinks(root)
	if err != nil {
		rootEval = root
	}
	pathEval, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		return false
	}
	rel, err := filepath.Rel(rootEval, pathEval)
	return err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) || filepath.IsAbs(rel)
}

func normalizeEntry(entry RegistryFileEntry) RegistryFileEntry {
	if entry.ID == "" {
		entry.ID = stableID(entry.Layer, entry.Path)
	}
	if entry.CheckMode == "" {
		entry.CheckMode = CheckModeHash
	}
	if entry.ReadPolicy == "" {
		entry.ReadPolicy = ReadPolicyContentIfChanged
	}
	entry.Path = filepath.ToSlash(strings.ReplaceAll(entry.Path, "\\", "/"))
	return entry
}

func blockingReason(entry RegistryFileEntry, status FileStatus) (BlockedBy, bool) {
	switch status.Status {
	case FileStatusInvalidPath:
		return BlockedBy{Path: status.Path, Layer: status.Layer, Status: status.Status, Reason: "invalid anchor path", Error: deref(status.Error)}, true
	case FileStatusMissingRequired:
		return BlockedBy{Path: status.Path, Layer: status.Layer, Status: status.Status, Reason: "required anchor file is missing"}, true
	case FileStatusUnreadable:
		if entry.Required {
			return BlockedBy{Path: status.Path, Layer: status.Layer, Status: status.Status, Reason: "required anchor file is unreadable", Error: deref(status.Error)}, true
		}
	}
	return BlockedBy{}, false
}

func requiredReadNeeded(entry RegistryFileEntry, status FileStatus, checkpointFound bool, handoffNamed map[string]bool) bool {
	if status.Status == FileStatusUnchanged || status.Status == FileStatusMissingOptional || status.Status == FileStatusMissingRequired || status.Status == FileStatusInvalidPath || status.Status == FileStatusUnreadable {
		return false
	}
	if !checkpointFound && entry.Required {
		return true
	}
	if entry.Required && (status.Status == FileStatusChanged || status.Status == FileStatusUnknown) {
		return true
	}
	if isInbound(entry) && (status.Status == FileStatusChanged || status.Status == FileStatusUnknown) {
		return true
	}
	if handoffNamed[filepath.ToSlash(entry.Path)] && (status.Status == FileStatusChanged || status.Status == FileStatusUnknown) {
		return true
	}
	return entry.ReadPolicy == ReadPolicyAlwaysForOwner && (status.Status == FileStatusChanged || status.Status == FileStatusUnknown)
}

func readReason(entry RegistryFileEntry, status FileStatus, checkpointFound bool, handoffNamed map[string]bool) string {
	if !checkpointFound {
		return "missing checkpoint requires first-run refresh"
	}
	if handoffNamed[filepath.ToSlash(entry.Path)] {
		return "handoff named file requires refresh"
	}
	if isInbound(entry) {
		return "inbound shared dependency changed"
	}
	if status.Status == FileStatusUnknown {
		return "fingerprint state is unknown"
	}
	return "registered anchor file changed"
}

func readPriority(entry RegistryFileEntry, status FileStatus) string {
	if entry.Required || status.Status == FileStatusUnknown || isInbound(entry) {
		return "high"
	}
	return "normal"
}

func addRequiredRead(reads *[]RequiredRead, seen map[string]bool, read RequiredRead) {
	key := read.Path + "\x00" + read.Layer
	if seen[key] {
		return
	}
	seen[key] = true
	*reads = append(*reads, read)
}

func decideAnchorState(result ReanchorResult, checkpointFound bool) string {
	if len(result.BlockedBy) > 0 {
		return AnchorStateBlocked
	}
	if !checkpointFound {
		return AnchorStateUnknown
	}
	for _, status := range result.FileStatuses {
		if status.Status == FileStatusUnknown {
			return AnchorStateUnknown
		}
	}
	if len(result.ChangedLayers) > 0 || len(result.RequiredUpdates) > 0 {
		return AnchorStateChanged
	}
	return AnchorStateUnchanged
}

func decideNextAction(result ReanchorResult, checkpointFound bool) string {
	if len(result.BlockedBy) > 0 {
		for _, block := range result.BlockedBy {
			if block.Status == FileStatusUnreadable {
				return NextActionStopForReadError
			}
		}
		return NextActionFailSafeFullReanchor
	}
	if result.Operation == OperationOrdinaryThreadCloseout {
		for _, update := range result.RequiredUpdates {
			if update.Layer == Layer2CrossThreadShared {
				return NextActionPerformCrossThreadHandoff
			}
		}
	}
	if !checkpointFound {
		return NextActionReadRequiredThenContinue
	}
	if contains(result.ChangedLayers, Layer0FrameworkBaseline) {
		return NextActionRunFrameworkUpgradeFlow
	}
	if contains(result.ChangedLayers, Layer1ThreadDefinition) {
		return NextActionRefreshThreadDefinition
	}
	if len(result.RequiredReads) > 0 {
		return NextActionReadRequiredThenContinue
	}
	return NextActionContinue
}

func minimalChatOutput(result ReanchorResult, checkpointFound bool) MinimalChatOutput {
	refresh := "none"
	if len(result.ChangedLayers) > 0 {
		refresh = strings.Join(shortLayers(result.ChangedLayers), ", ")
	}
	reason := ""
	if !checkpointFound {
		reason = "checkpoint missing or not supplied"
	} else if len(result.BlockedBy) > 0 {
		reason = result.BlockedBy[0].Reason
	} else if len(result.RequiredReads) > 0 {
		reason = result.RequiredReads[0].Reason
	}
	next := "continue within current thread scope"
	switch result.NextAction {
	case NextActionReadRequiredThenContinue:
		next = "read required files before work continues"
	case NextActionRunFrameworkUpgradeFlow:
		next = "run framework upgrade flow before ordinary work"
	case NextActionRefreshThreadDefinition:
		next = "refresh this thread definition before work"
	case NextActionPerformCrossThreadHandoff:
		next = "perform required cross-thread handoff update"
	case NextActionFailSafeFullReanchor:
		next = "stop and perform fail-safe full reanchor"
	case NextActionStopForReadError:
		next = "stop for anchor file read error"
	}
	return MinimalChatOutput{
		AnchorState: result.AnchorState,
		Refresh:     refresh,
		Reason:      reason,
		Next:        next,
	}
}

func closeoutUpdates(req normalizedRequest) []RequiredUpdate {
	updates := []RequiredUpdate{}
	if req.Operation != OperationOrdinaryThreadCloseout {
		return updates
	}
	for _, event := range req.CloseoutEvents {
		if event.TargetPath == "" {
			continue
		}
		layer := event.Layer
		if layer == "" {
			layer = Layer3ThreadLocalMemory
		}
		updateType := "local_memory"
		if event.Type == "cross_thread_dependency_changed" || layer == Layer2CrossThreadShared {
			updateType = "cross_thread_dependency"
			layer = Layer2CrossThreadShared
		}
		updates = append(updates, RequiredUpdate{
			Path:         filepath.ToSlash(strings.ReplaceAll(event.TargetPath, "\\", "/")),
			Layer:        layer,
			UpdateType:   updateType,
			OwnerThread:  req.Thread.ID,
			TargetThread: event.TargetThread,
			Reason:       event.Reason,
		})
	}
	return updates
}

func periodicDue(operation string, conversation Conversation) bool {
	return operation == OperationPeriodicReanchor || conversation.ForcePeriodic || (conversation.RoundCount > 0 && conversation.RoundCount%10 == 0)
}

func statusCountsAsChanged(status string) bool {
	return status != FileStatusUnchanged && status != FileStatusMissingOptional
}

func anchorKey(id, path string) string {
	return id + "\x00" + filepath.ToSlash(path)
}

func stableID(layer, path string) string {
	path = strings.Trim(filepath.ToSlash(path), "/")
	path = strings.NewReplacer("/", ".", "-", "_", " ", "_").Replace(path)
	if layer == "" {
		return path
	}
	return layer + "." + path
}

func isInbound(entry RegistryFileEntry) bool {
	return entry.Layer == Layer2CrossThreadShared && entry.TargetThread != "" && entry.SourceThread != ""
}

func makePathSet(paths []string) map[string]bool {
	set := map[string]bool{}
	for _, path := range paths {
		set[filepath.ToSlash(strings.ReplaceAll(path, "\\", "/"))] = true
	}
	return set
}

func sortedKeys(set map[string]bool) []string {
	out := make([]string, 0, len(set))
	for key := range set {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}

func shortLayers(layers []string) []string {
	out := make([]string, 0, len(layers))
	for _, layer := range layers {
		switch layer {
		case Layer0FrameworkBaseline:
			out = append(out, "layer_0")
		case Layer1ThreadDefinition:
			out = append(out, "layer_1")
		case Layer2CrossThreadShared:
			out = append(out, "layer_2")
		case Layer3ThreadLocalMemory:
			out = append(out, "layer_3")
		default:
			out = append(out, layer)
		}
	}
	return out
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func deref(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
