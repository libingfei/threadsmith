package reanchor

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestMissingCheckpointRequiresFirstRunRefresh(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateUnknown {
		t.Fatalf("anchor_state = %s, want %s", result.AnchorState, AnchorStateUnknown)
	}
	if result.NextAction != NextActionReadRequiredThenContinue {
		t.Fatalf("next_action = %s", result.NextAction)
	}
	if len(result.RequiredReads) == 0 {
		t.Fatal("expected required reads for first-run refresh")
	}
	if result.CheckpointUpdate.WriteAllowed {
		t.Fatal("checkpoint update must not be write_allowed before required reads")
	}
}

func TestAllFilesUnchanged(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Checkpoint.State = checkpointForRequest(t, req)
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateUnchanged {
		t.Fatalf("anchor_state = %s, want %s", result.AnchorState, AnchorStateUnchanged)
	}
	if result.NextAction != NextActionContinue {
		t.Fatalf("next_action = %s", result.NextAction)
	}
	if len(result.RequiredReads) != 0 {
		t.Fatalf("required_reads = %d, want 0", len(result.RequiredReads))
	}
	if !result.CheckpointUpdate.WriteAllowed {
		t.Fatal("checkpoint update should be allowed when nothing requires reads")
	}
}

func TestLayer0Changed(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Checkpoint.State = checkpointForRequest(t, req)
	writeFile(t, root, "docs/anchor_pm/00_framework_baseline/baseline.md", "# Baseline\n\nBaseline id: `package-first-v2.0`\n")
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateChanged {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
	if result.NextAction != NextActionRunFrameworkUpgradeFlow {
		t.Fatalf("next_action = %s", result.NextAction)
	}
	assertLayerChanged(t, result, Layer0FrameworkBaseline)
}

func TestLayer1Changed(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Checkpoint.State = checkpointForRequest(t, req)
	writeFile(t, root, "docs/anchor_pm/01_thread_definitions/product_manager.md", "# Product Manager\n\nchanged\n")
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.NextAction != NextActionRefreshThreadDefinition {
		t.Fatalf("next_action = %s", result.NextAction)
	}
	assertLayerChanged(t, result, Layer1ThreadDefinition)
}

func TestLayer2InboundChanged(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "reanchor_detector_core")
	req.Checkpoint.State = checkpointForRequest(t, req)
	writeFile(t, root, "docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md", "# Shared\n\nchanged\n")
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.NextAction != NextActionReadRequiredThenContinue {
		t.Fatalf("next_action = %s", result.NextAction)
	}
	assertLayerChanged(t, result, Layer2CrossThreadShared)
	if len(result.RequiredReads) == 0 {
		t.Fatal("expected inbound changed file to be required read")
	}
}

func TestLayer3LocalChanged(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Checkpoint.State = checkpointForRequest(t, req)
	writeFile(t, root, "docs/module_state/product_manager.md", "# Product Manager State\n\nchanged\n")
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.NextAction != NextActionReadRequiredThenContinue {
		t.Fatalf("next_action = %s", result.NextAction)
	}
	assertLayerChanged(t, result, Layer3ThreadLocalMemory)
}

func TestMissingRequiredBlocks(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Checkpoint.State = checkpointForRequest(t, req)
	if err := os.Remove(filepath.Join(root, "docs/module_state/product_manager.md")); err != nil {
		t.Fatal(err)
	}
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateBlocked {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
	if result.NextAction != NextActionFailSafeFullReanchor {
		t.Fatalf("next_action = %s", result.NextAction)
	}
}

func TestMissingOptionalContinues(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Checkpoint.State = checkpointForRequest(t, req)
	if err := os.Remove(filepath.Join(root, "docs/anchor_pm/review_log.md")); err != nil {
		t.Fatal(err)
	}
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateUnchanged {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
}

func TestInvalidPathOutsideProjectRootBlocks(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Registry.Files = append(req.Registry.Files, RegistryFileEntry{
		ID:         "bad",
		Layer:      Layer1ThreadDefinition,
		Path:       "../outside.md",
		Required:   true,
		CheckMode:  CheckModeHash,
		ReadPolicy: ReadPolicyContentIfChanged,
	})
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateBlocked {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
	if result.FileStatuses[0].Status != FileStatusInvalidPath {
		t.Fatalf("status = %s", result.FileStatuses[0].Status)
	}
}

func TestRequiredReadsOverLimitBlocks(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Options.MaxRequiredReads = 1
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateBlocked {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
	if result.NextAction != NextActionFailSafeFullReanchor {
		t.Fatalf("next_action = %s", result.NextAction)
	}
}

func TestUnknownFingerprintRequiresRead(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Registry.Files = []RegistryFileEntry{{
		ID:         "unknown",
		Layer:      Layer3ThreadLocalMemory,
		Path:       "docs/module_state/product_manager.md",
		Required:   true,
		CheckMode:  "unsupported",
		ReadPolicy: ReadPolicyContentIfChanged,
	}}
	req.Checkpoint.State = &Checkpoint{
		SchemaVersion: CheckpointSchemaV1,
		ThreadID:      "product_manager",
		Anchors:       []CheckpointAnchor{},
	}
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateUnknown {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
	if len(result.RequiredReads) != 1 {
		t.Fatalf("required_reads = %d", len(result.RequiredReads))
	}
}

func TestMTimeSizeFallbackDetectsChange(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Registry.Files = []RegistryFileEntry{{
		ID:         "mtime",
		Layer:      Layer3ThreadLocalMemory,
		Path:       "docs/module_state/product_manager.md",
		Required:   true,
		CheckMode:  CheckModeMTimeSize,
		ReadPolicy: ReadPolicyContentIfChanged,
	}}
	req.Checkpoint.State = checkpointForRequest(t, req)
	writeFile(t, root, "docs/module_state/product_manager.md", "# Product Manager State\n\nchanged size\n")
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateChanged {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
}

func TestCRLFContentHashStableWhenUnchanged(t *testing.T) {
	root := fixtureProject(t)
	writeFile(t, root, "docs/module_state/product_manager.md", "# Product Manager State\r\n\r\nsame\r\n")
	req := baseRequest(root, "product_manager")
	req.Registry.Files = []RegistryFileEntry{{
		ID:         "crlf",
		Layer:      Layer3ThreadLocalMemory,
		Path:       "docs/module_state/product_manager.md",
		Required:   true,
		CheckMode:  CheckModeHash,
		ReadPolicy: ReadPolicyContentIfChanged,
	}}
	req.Checkpoint.State = checkpointForRequest(t, req)
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateUnchanged {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
}

func TestPeriodicReanchorDue(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Conversation.RoundCount = 10
	req.Checkpoint.State = checkpointForRequest(t, req)
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if !result.PeriodicDue {
		t.Fatal("expected periodic_due")
	}
	if result.AnchorState != AnchorStateUnchanged {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
}

func TestCloseoutLocalMemoryUpdate(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Operation = OperationOrdinaryThreadCloseout
	req.Checkpoint.State = checkpointForRequest(t, req)
	req.CloseoutEvents = []CloseoutEvent{{
		Type:       "local_memory_changed",
		Layer:      Layer3ThreadLocalMemory,
		TargetPath: "docs/module_state/product_manager/key_memory.md",
		Reason:     "durable product principle changed",
	}}
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.RequiredUpdates) != 1 {
		t.Fatalf("required_updates = %d", len(result.RequiredUpdates))
	}
	if result.RequiredUpdates[0].UpdateType != "local_memory" {
		t.Fatalf("update_type = %s", result.RequiredUpdates[0].UpdateType)
	}
}

func TestCloseoutCrossThreadDependencyUpdate(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Operation = OperationOrdinaryThreadCloseout
	req.Checkpoint.State = checkpointForRequest(t, req)
	req.CloseoutEvents = []CloseoutEvent{{
		Type:         "cross_thread_dependency_changed",
		Layer:        Layer2CrossThreadShared,
		TargetPath:   "docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md",
		TargetThread: "reanchor_detector_core",
		Reason:       "detector requirements changed",
	}}
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.RequiredUpdates) != 1 {
		t.Fatalf("required_updates = %d", len(result.RequiredUpdates))
	}
	if result.NextAction != NextActionPerformCrossThreadHandoff {
		t.Fatalf("next_action = %s", result.NextAction)
	}
}

func TestBackslashPathInput(t *testing.T) {
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Registry.Files = []RegistryFileEntry{{
		ID:         "backslash",
		Layer:      Layer3ThreadLocalMemory,
		Path:       `docs\module_state\product_manager.md`,
		Required:   true,
		CheckMode:  CheckModeHash,
		ReadPolicy: ReadPolicyContentIfChanged,
	}}
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.FileStatuses[0].Path != "docs/module_state/product_manager.md" {
		t.Fatalf("path = %s", result.FileStatuses[0].Path)
	}
}

func TestUnreadableRequiredFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows ACLs do not reliably model chmod unreadable fixtures")
	}
	if os.Geteuid() == 0 {
		t.Skip("root can read files regardless of mode")
	}
	root := fixtureProject(t)
	req := baseRequest(root, "product_manager")
	req.Checkpoint.State = checkpointForRequest(t, req)
	path := filepath.Join(root, "docs/module_state/product_manager.md")
	if err := os.Chmod(path, 0); err != nil {
		t.Fatal(err)
	}
	defer os.Chmod(path, 0o600)
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateBlocked {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
	if result.NextAction != NextActionStopForReadError {
		t.Fatalf("next_action = %s", result.NextAction)
	}
}

func TestSymlinkTraversalBlocks(t *testing.T) {
	root := fixtureProject(t)
	outside := t.TempDir()
	writeFileAbs(t, filepath.Join(outside, "outside.md"), "outside")
	linkPath := filepath.Join(root, "docs/module_state/product_manager/link.md")
	if err := os.Symlink(filepath.Join(outside, "outside.md"), linkPath); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	req := baseRequest(root, "product_manager")
	req.Registry.Files = []RegistryFileEntry{{
		ID:         "symlink",
		Layer:      Layer3ThreadLocalMemory,
		Path:       "docs/module_state/product_manager/link.md",
		Required:   true,
		CheckMode:  CheckModeHash,
		ReadPolicy: ReadPolicyContentIfChanged,
	}}
	result, err := Evaluate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.AnchorState != AnchorStateBlocked {
		t.Fatalf("anchor_state = %s", result.AnchorState)
	}
	if result.FileStatuses[0].Status != FileStatusInvalidPath {
		t.Fatalf("status = %s", result.FileStatuses[0].Status)
	}
}

func fixtureProject(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	files := map[string]string{
		"AGENTS.md":                                                                    "project rules\n",
		"docs/anchor_pm/current_version.md":                                            "# Current\n\nCurrent coordination version: `package-first-v1.0`\n",
		"docs/anchor_pm/contracts.md":                                                  "# Contracts\n",
		"docs/anchor_pm/thread_initialization.md":                                      "# Thread init\n",
		"docs/anchor_pm/interaction_guide.md":                                          "# Interaction\n",
		"docs/anchor_pm/review_log.md":                                                 "# Review\n",
		"docs/anchor_pm/simplification.md":                                             "# Simplification\n",
		"docs/anchor_pm/00_framework_baseline/baseline.md":                             "# Baseline\n\nBaseline id: `package-first-v1.0`\n",
		"docs/anchor_pm/01_thread_definitions/index.md":                                "# Index\n",
		"docs/anchor_pm/01_thread_definitions/product_manager.md":                      "# Product Manager\n",
		"docs/anchor_pm/01_thread_definitions/reanchor_detector_core.md":               "# Reanchor Detector Core\n",
		"docs/anchor_pm/02_shared_state/index.md":                                      "# Shared Index\n",
		"docs/anchor_pm/02_shared_state/project_version_notice.md":                     "# Version Notice\n",
		"docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md": "# Shared\n",
		"docs/anchor_pm/03_thread_local_memory.md":                                     "# Local Memory\n",
		"docs/module_state/product_manager.md":                                         "# Product Manager State\n",
		"docs/module_state/product_manager/current_state.md":                           "# Current State\n",
		"docs/module_state/reanchor_detector_core.md":                                  "# Detector State\n",
		"docs/module_state/reanchor_detector_core/current_state.md":                    "# Current State\n",
	}
	for path, content := range files {
		writeFile(t, root, path, content)
	}
	return root
}

func baseRequest(root, thread string) ReanchorRequest {
	return ReanchorRequest{
		SchemaVersion: RequestSchemaV1,
		Operation:     OperationOrdinaryThreadStart,
		ProjectRoot:   root,
		Thread: ThreadRequest{
			ID:   thread,
			Name: thread,
		},
		Conversation: Conversation{
			RoundCount: 7,
		},
		Checkpoint: CheckpointInput{
			Path: "docs/anchor_pm/.state/" + thread + ".anchor_state.json",
		},
		Registry: RegistryInput{
			Mode: defaultRegistryMode,
		},
		Options: Options{
			HashAlgorithm:    "sha256",
			MaxRequiredReads: 20,
		},
	}
}

func checkpointForRequest(t *testing.T, req ReanchorRequest) *Checkpoint {
	t.Helper()
	initial := req
	initial.Checkpoint.State = nil
	result, err := Evaluate(context.Background(), initial)
	if err != nil {
		t.Fatal(err)
	}
	cp := result.CheckpointUpdate.State
	return &cp
}

func assertLayerChanged(t *testing.T, result ReanchorResult, layer string) {
	t.Helper()
	for _, changed := range result.ChangedLayers {
		if changed == layer {
			return
		}
	}
	t.Fatalf("changed layer %s not found in %#v", layer, result.ChangedLayers)
}

func writeFile(t *testing.T, root, rel, content string) {
	t.Helper()
	writeFileAbs(t, filepath.Join(root, filepath.FromSlash(rel)), content)
}

func writeFileAbs(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}
