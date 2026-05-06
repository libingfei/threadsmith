package reanchor

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

func registryForRequest(req normalizedRequest) ([]RegistryFileEntry, error) {
	if len(req.Registry.Files) > 0 {
		out := make([]RegistryFileEntry, 0, len(req.Registry.Files))
		for _, entry := range req.Registry.Files {
			out = append(out, normalizeEntry(entry))
		}
		return out, nil
	}
	if req.Registry.Mode != defaultRegistryMode {
		return nil, fmt.Errorf("unsupported registry mode %q", req.Registry.Mode)
	}
	return discoverDefaultRegistry(req)
}

func discoverDefaultRegistry(req normalizedRequest) ([]RegistryFileEntry, error) {
	var entries []RegistryFileEntry
	seen := map[string]bool{}
	add := func(entry RegistryFileEntry) {
		entry = normalizeEntry(entry)
		key := entry.Layer + "\x00" + entry.Path
		if seen[key] {
			return
		}
		seen[key] = true
		entries = append(entries, entry)
	}

	add(RegistryFileEntry{
		ID:         "layer0.baseline",
		Layer:      Layer0FrameworkBaseline,
		Path:       "docs/anchor_pm/00_framework_baseline/baseline.md",
		Required:   true,
		Category:   "baseline",
		CheckMode:  CheckModeVersion,
		ReadPolicy: ReadPolicyContentIfChanged,
	})

	add(RegistryFileEntry{
		ID:         "layer1.thread_definition.index",
		Layer:      Layer1ThreadDefinition,
		Path:       "docs/anchor_pm/01_thread_definitions/index.md",
		Required:   true,
		Category:   "definition_index",
		CheckMode:  CheckModeHash,
		ReadPolicy: ReadPolicyContentIfChanged,
	})
	add(RegistryFileEntry{
		ID:          "layer1.thread_definition." + req.Thread.ID,
		Layer:       Layer1ThreadDefinition,
		Path:        "docs/anchor_pm/01_thread_definitions/" + req.Thread.ID + ".md",
		Required:    true,
		OwnerThread: req.Thread.ID,
		Category:    "definition",
		CheckMode:   CheckModeHash,
		ReadPolicy:  ReadPolicyContentIfChanged,
	})

	if req.includeCompat {
		add(RegistryFileEntry{
			ID:         "compat.agents",
			Layer:      Layer1ThreadDefinition,
			Path:       "AGENTS.md",
			Required:   true,
			Category:   "project_rules",
			CheckMode:  CheckModeHash,
			ReadPolicy: ReadPolicyContentIfChanged,
		})
		add(RegistryFileEntry{
			ID:         "compat.contracts",
			Layer:      Layer1ThreadDefinition,
			Path:       "docs/anchor_pm/contracts.md",
			Required:   true,
			Category:   "contracts",
			CheckMode:  CheckModeHash,
			ReadPolicy: ReadPolicyContentIfChanged,
		})
		add(RegistryFileEntry{
			ID:         "compat.thread_initialization",
			Layer:      Layer1ThreadDefinition,
			Path:       "docs/anchor_pm/thread_initialization.md",
			Required:   false,
			Category:   "thread_initialization",
			CheckMode:  CheckModeHash,
			ReadPolicy: ReadPolicyContentIfChanged,
		})
		add(RegistryFileEntry{
			ID:                  "compat.current_version",
			Layer:               Layer0FrameworkBaseline,
			Path:                "docs/anchor_pm/current_version.md",
			Required:            true,
			Category:            "current_version",
			CheckMode:           CheckModeVersion,
			ReadPolicy:          ReadPolicyContentIfChanged,
			CompatibilitySource: "docs/anchor_pm/00_framework_baseline/baseline.md",
		})
	}

	add(RegistryFileEntry{
		ID:         "layer2.shared_state.index",
		Layer:      Layer2CrossThreadShared,
		Path:       "docs/anchor_pm/02_shared_state/index.md",
		Required:   false,
		Category:   "shared_index",
		CheckMode:  CheckModeHash,
		ReadPolicy: ReadPolicyContentIfRelevant,
	})
	addSharedStateFiles(req, add)
	addHandoffNamedFiles(req, add)
	if req.includeCompat {
		for _, path := range []string{
			"docs/anchor_pm/interaction_guide.md",
			"docs/anchor_pm/review_log.md",
			"docs/anchor_pm/simplification.md",
		} {
			add(RegistryFileEntry{
				ID:         stableID("compat.shared", path),
				Layer:      Layer2CrossThreadShared,
				Path:       path,
				Required:   false,
				Category:   "shared_compatibility",
				CheckMode:  CheckModeHash,
				ReadPolicy: ReadPolicyContentIfRelevant,
			})
		}
	}

	add(RegistryFileEntry{
		ID:         "layer3.local_memory.index",
		Layer:      Layer3ThreadLocalMemory,
		Path:       "docs/anchor_pm/03_thread_local_memory.md",
		Required:   false,
		Category:   "local_memory_index",
		CheckMode:  CheckModeHash,
		ReadPolicy: ReadPolicyContentIfChanged,
	})
	add(RegistryFileEntry{
		ID:          "layer3.module_state." + req.Thread.ID,
		Layer:       Layer3ThreadLocalMemory,
		Path:        "docs/module_state/" + req.Thread.ID + ".md",
		Required:    true,
		OwnerThread: req.Thread.ID,
		Category:    "current_state",
		CheckMode:   CheckModeHash,
		ReadPolicy:  ReadPolicyContentIfChanged,
	})
	addLocalCategoryFiles(req, add)

	switch req.Operation {
	case OperationThreadManagement:
		addThreadDefinitionFiles(req, add)
	case OperationDetectorDevelopment:
		addAnchorPMDevelopmentFiles(req, add)
	case OperationValidation:
		addValidationFiles(req, add)
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Layer == entries[j].Layer {
			return entries[i].Path < entries[j].Path
		}
		return entries[i].Layer < entries[j].Layer
	})
	return entries, nil
}

func addSharedStateFiles(req normalizedRequest, add func(RegistryFileEntry)) {
	root := filepath.Join(req.projectRoot, "docs", "anchor_pm", "02_shared_state")
	files, _ := filepath.Glob(filepath.Join(root, "*.md"))
	sort.Strings(files)
	for _, abs := range files {
		rel, err := filepath.Rel(req.projectRoot, abs)
		if err != nil {
			continue
		}
		rel = filepath.ToSlash(rel)
		name := filepath.Base(rel)
		if name == "index.md" {
			continue
		}
		source, target, directed := parseDirectedSharedName(name)
		isInbound := directed && target == req.Thread.ID
		isMirror := !directed
		if !isInbound && !isMirror {
			continue
		}
		add(RegistryFileEntry{
			ID:           stableID("layer2.shared", rel),
			Layer:        Layer2CrossThreadShared,
			Path:         rel,
			Required:     false,
			SourceThread: source,
			TargetThread: target,
			Category:     sharedCategory(name, directed),
			CheckMode:    CheckModeHash,
			ReadPolicy:   ReadPolicyContentIfRelevant,
		})
	}
}

func addHandoffNamedFiles(req normalizedRequest, add func(RegistryFileEntry)) {
	for _, path := range req.Handoff.NamedFiles {
		add(RegistryFileEntry{
			ID:           stableID("handoff.named", path),
			Layer:        layerForHandoffPath(path),
			Path:         path,
			Required:     true,
			SourceThread: req.Handoff.SourceThread,
			TargetThread: req.Handoff.TargetThread,
			Category:     "handoff_named",
			CheckMode:    CheckModeHash,
			ReadPolicy:   ReadPolicyContentIfRelevant,
		})
	}
}

func addLocalCategoryFiles(req normalizedRequest, add func(RegistryFileEntry)) {
	root := filepath.Join(req.projectRoot, "docs", "module_state", req.Thread.ID)
	files, _ := filepath.Glob(filepath.Join(root, "*.md"))
	sort.Strings(files)
	for _, abs := range files {
		rel, err := filepath.Rel(req.projectRoot, abs)
		if err != nil {
			continue
		}
		rel = filepath.ToSlash(rel)
		add(RegistryFileEntry{
			ID:          stableID("layer3.category", rel),
			Layer:       Layer3ThreadLocalMemory,
			Path:        rel,
			Required:    false,
			OwnerThread: req.Thread.ID,
			Category:    strings.TrimSuffix(filepath.Base(rel), ".md"),
			CheckMode:   CheckModeHash,
			ReadPolicy:  ReadPolicyContentIfChanged,
		})
	}
}

func addThreadDefinitionFiles(req normalizedRequest, add func(RegistryFileEntry)) {
	root := filepath.Join(req.projectRoot, "docs", "anchor_pm", "01_thread_definitions")
	files, _ := filepath.Glob(filepath.Join(root, "*.md"))
	sort.Strings(files)
	for _, abs := range files {
		rel, err := filepath.Rel(req.projectRoot, abs)
		if err != nil {
			continue
		}
		rel = filepath.ToSlash(rel)
		add(RegistryFileEntry{
			ID:         stableID("thread_management.definition", rel),
			Layer:      Layer1ThreadDefinition,
			Path:       rel,
			Required:   true,
			Category:   "definition",
			CheckMode:  CheckModeHash,
			ReadPolicy: ReadPolicyContentIfChanged,
		})
	}
}

func addAnchorPMDevelopmentFiles(req normalizedRequest, add func(RegistryFileEntry)) {
	for _, path := range []string{
		"README.md",
		"PRODUCT_PRINCIPLES.md",
		"MVP_SPEC.md",
		"ANCHOR_PM_FRAMEWORK_DESIGN.md",
		"docs/anchor_pm/internal_function_spec.md",
		"docs/anchor_pm/contract_state_detector.md",
		"docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md",
	} {
		add(RegistryFileEntry{
			ID:         stableID("anchorpm.dev", path),
			Layer:      LayerAnchorPMDevelopment,
			Path:       path,
			Required:   false,
			Category:   "development_reference",
			CheckMode:  CheckModeHash,
			ReadPolicy: ReadPolicyContentIfRelevant,
		})
	}
	files, _ := filepath.Glob(filepath.Join(req.projectRoot, "ANCHOR_PM_INSTALL_PROMPT*.md"))
	sort.Strings(files)
	for _, abs := range files {
		rel, err := filepath.Rel(req.projectRoot, abs)
		if err == nil {
			add(RegistryFileEntry{
				ID:         stableID("anchorpm.dev", rel),
				Layer:      LayerAnchorPMDevelopment,
				Path:       filepath.ToSlash(rel),
				Required:   false,
				Category:   "install_prompt",
				CheckMode:  CheckModeHash,
				ReadPolicy: ReadPolicyContentIfRelevant,
			})
		}
	}
	_ = filepath.WalkDir(filepath.Join(req.projectRoot, "packages"), func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(req.projectRoot, path)
		if err != nil {
			return nil
		}
		add(RegistryFileEntry{
			ID:         stableID("anchorpm.dev.package", rel),
			Layer:      LayerAnchorPMDevelopment,
			Path:       filepath.ToSlash(rel),
			Required:   false,
			Category:   "package_file",
			CheckMode:  CheckModeHash,
			ReadPolicy: ReadPolicyContentIfRelevant,
		})
		return nil
	})
}

func addValidationFiles(req normalizedRequest, add func(RegistryFileEntry)) {
	root := filepath.Join(req.projectRoot, "docs", "anchor_pm", "reports")
	files, _ := filepath.Glob(filepath.Join(root, "*.md"))
	sort.Strings(files)
	for _, abs := range files {
		rel, err := filepath.Rel(req.projectRoot, abs)
		if err != nil {
			continue
		}
		add(RegistryFileEntry{
			ID:         stableID("validation.report", rel),
			Layer:      Layer2CrossThreadShared,
			Path:       filepath.ToSlash(rel),
			Required:   false,
			Category:   "validation_report",
			CheckMode:  CheckModeHash,
			ReadPolicy: ReadPolicyContentIfRelevant,
		})
	}
}

func parseDirectedSharedName(name string) (source, target string, ok bool) {
	if !strings.HasSuffix(name, ".md") || !strings.Contains(name, "__to_") {
		return "", "", false
	}
	stem := strings.TrimSuffix(name, ".md")
	parts := strings.Split(stem, "__to_")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func sharedCategory(name string, directed bool) string {
	if directed {
		return "directed_dependency"
	}
	return strings.TrimSuffix(name, ".md")
}

func layerForHandoffPath(path string) string {
	path = filepath.ToSlash(strings.ReplaceAll(path, "\\", "/"))
	switch {
	case strings.HasPrefix(path, "docs/module_state/"):
		return Layer3ThreadLocalMemory
	case strings.HasPrefix(path, "docs/anchor_pm/02_shared_state/"):
		return Layer2CrossThreadShared
	case strings.HasPrefix(path, "docs/anchor_pm/01_thread_definitions/"):
		return Layer1ThreadDefinition
	case strings.HasPrefix(path, "docs/anchor_pm/00_framework_baseline/"):
		return Layer0FrameworkBaseline
	default:
		return Layer2CrossThreadShared
	}
}
