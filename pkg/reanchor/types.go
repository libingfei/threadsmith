package reanchor

const (
	RequestSchemaV1    = "anchorpm.reanchor.request.v1"
	ResultSchemaV1     = "anchorpm.reanchor.result.v1"
	CheckpointSchemaV1 = "anchorpm.reanchor.checkpoint.v1"

	Layer0FrameworkBaseline  = "layer_0_framework_baseline"
	Layer1ThreadDefinition   = "layer_1_thread_definition"
	Layer2CrossThreadShared  = "layer_2_cross_thread_shared_state"
	Layer3ThreadLocalMemory  = "layer_3_thread_local_memory"
	LayerAnchorPMDevelopment = "anchor_pm_development"

	OperationOrdinaryThreadStart    = "ordinary_thread_start"
	OperationPeriodicReanchor       = "periodic_reanchor"
	OperationOrdinaryThreadCloseout = "ordinary_thread_closeout"
	OperationThreadManagement       = "thread_management"
	OperationFrameworkUpgrade       = "framework_upgrade"
	OperationCrossThreadHandoff     = "cross_thread_handoff"
	OperationDetectorDevelopment    = "detector_development"
	OperationValidation             = "validation"

	FileStatusUnchanged       = "unchanged"
	FileStatusChanged         = "changed"
	FileStatusUnknown         = "unknown"
	FileStatusMissingRequired = "missing_required"
	FileStatusMissingOptional = "missing_optional"
	FileStatusUnreadable      = "unreadable"
	FileStatusInvalidPath     = "invalid_path"

	AnchorStateUnchanged = "unchanged"
	AnchorStateChanged   = "changed"
	AnchorStateUnknown   = "unknown"
	AnchorStateBlocked   = "blocked"

	NextActionContinue                  = "continue"
	NextActionReadRequiredThenContinue  = "read_required_files_then_continue"
	NextActionRunFrameworkUpgradeFlow   = "run_framework_upgrade_flow"
	NextActionRefreshThreadDefinition   = "refresh_thread_definition"
	NextActionHandoffToThreadManagement = "handoff_to_thread_management"
	NextActionPerformCrossThreadHandoff = "perform_cross_thread_handoff"
	NextActionFailSafeFullReanchor      = "fail_safe_full_reanchor"
	NextActionStopForReadError          = "stop_for_read_error"

	ReadModeMetadataOnly     = "metadata_only"
	ReadModeFull             = "full"
	ReadModeSection          = "section"
	ReadModeGeneratedSummary = "generated_summary"

	CheckModeVersion   = "version"
	CheckModeHash      = "hash"
	CheckModeMTimeSize = "mtime_size"
	CheckModeExists    = "exists"

	ReadPolicyNever             = "never"
	ReadPolicyContentIfChanged  = "content_if_changed"
	ReadPolicyContentIfRelevant = "content_if_relevant"
	ReadPolicyAlwaysForOwner    = "always_for_owner"
)

type ReanchorRequest struct {
	SchemaVersion  string          `json:"schema_version"`
	Operation      string          `json:"operation"`
	ProjectRoot    string          `json:"project_root"`
	Thread         ThreadRequest   `json:"thread"`
	Conversation   Conversation    `json:"conversation"`
	Checkpoint     CheckpointInput `json:"checkpoint"`
	Registry       RegistryInput   `json:"registry"`
	Handoff        HandoffInput    `json:"handoff"`
	Options        Options         `json:"options"`
	CloseoutEvents []CloseoutEvent `json:"closeout_events,omitempty"`
}

type ThreadRequest struct {
	ID                 string `json:"id"`
	Name               string `json:"name,omitempty"`
	IsThreadManagement bool   `json:"is_thread_management"`
}

type Conversation struct {
	RoundCount    int    `json:"round_count"`
	ForcePeriodic bool   `json:"force_periodic"`
	UserTaskHint  string `json:"user_task_hint,omitempty"`
}

type CheckpointInput struct {
	Path  string      `json:"path,omitempty"`
	State *Checkpoint `json:"state,omitempty"`
}

type RegistryInput struct {
	Mode  string              `json:"mode"`
	Files []RegistryFileEntry `json:"files"`
}

type HandoffInput struct {
	NamedFiles   []string `json:"named_files"`
	SourceThread string   `json:"source_thread,omitempty"`
	TargetThread string   `json:"target_thread,omitempty"`
}

type Options struct {
	HashAlgorithm             string `json:"hash_algorithm,omitempty"`
	ReadFileContents          bool   `json:"read_file_contents"`
	IncludeCompatibilityFiles *bool  `json:"include_compatibility_files,omitempty"`
	MaxRequiredReads          int    `json:"max_required_reads,omitempty"`
}

type RegistryFileEntry struct {
	ID                  string `json:"id"`
	Layer               string `json:"layer"`
	Path                string `json:"path"`
	Required            bool   `json:"required"`
	OwnerThread         string `json:"owner_thread,omitempty"`
	SourceThread        string `json:"source_thread,omitempty"`
	TargetThread        string `json:"target_thread,omitempty"`
	Category            string `json:"category,omitempty"`
	CheckMode           string `json:"check_mode"`
	ReadPolicy          string `json:"read_policy"`
	CompatibilitySource string `json:"compatibility_source,omitempty"`
}

type Checkpoint struct {
	SchemaVersion string             `json:"schema_version"`
	ProjectID     string             `json:"project_id,omitempty"`
	ThreadID      string             `json:"thread_id"`
	UpdatedAt     string             `json:"updated_at,omitempty"`
	RoundCount    int                `json:"round_count"`
	LastOperation string             `json:"last_operation,omitempty"`
	Anchors       []CheckpointAnchor `json:"anchors"`
}

type CheckpointAnchor struct {
	ID              string `json:"id"`
	Layer           string `json:"layer"`
	Path            string `json:"path"`
	Fingerprint     string `json:"fingerprint"`
	FingerprintType string `json:"fingerprint_type"`
}

type ReanchorResult struct {
	SchemaVersion     string            `json:"schema_version"`
	Operation         string            `json:"operation"`
	ThreadID          string            `json:"thread_id"`
	CheckedAt         string            `json:"checked_at"`
	AnchorState       string            `json:"anchor_state"`
	PeriodicDue       bool              `json:"periodic_due"`
	ChangedLayers     []string          `json:"changed_layers"`
	FileStatuses      []FileStatus      `json:"file_statuses"`
	RequiredReads     []RequiredRead    `json:"required_reads"`
	BlockedBy         []BlockedBy       `json:"blocked_by"`
	NextAction        string            `json:"next_action"`
	CheckpointUpdate  CheckpointUpdate  `json:"checkpoint_update"`
	MinimalChatOutput MinimalChatOutput `json:"minimal_chat_output"`
	RequiredUpdates   []RequiredUpdate  `json:"required_updates,omitempty"`
}

type FileStatus struct {
	ID                  string  `json:"id"`
	Layer               string  `json:"layer"`
	Path                string  `json:"path"`
	PreviousFingerprint string  `json:"previous_fingerprint,omitempty"`
	CurrentFingerprint  string  `json:"current_fingerprint,omitempty"`
	Status              string  `json:"status"`
	Error               *string `json:"error"`
}

type RequiredRead struct {
	Path     string `json:"path"`
	Layer    string `json:"layer"`
	Reason   string `json:"reason"`
	ReadMode string `json:"read_mode"`
	Priority string `json:"priority"`
}

type BlockedBy struct {
	Path   string `json:"path,omitempty"`
	Layer  string `json:"layer,omitempty"`
	Status string `json:"status"`
	Reason string `json:"reason"`
	Error  string `json:"error,omitempty"`
}

type CheckpointUpdate struct {
	WriteAllowed bool       `json:"write_allowed"`
	Path         string     `json:"path"`
	State        Checkpoint `json:"state"`
}

type MinimalChatOutput struct {
	AnchorState string `json:"anchor_state"`
	Refresh     string `json:"refresh"`
	Reason      string `json:"reason,omitempty"`
	Next        string `json:"next"`
}

type CloseoutEvent struct {
	Type         string `json:"type"`
	Layer        string `json:"layer"`
	TargetPath   string `json:"target_path"`
	Reason       string `json:"reason,omitempty"`
	SourceThread string `json:"source_thread,omitempty"`
	TargetThread string `json:"target_thread,omitempty"`
}

type RequiredUpdate struct {
	Path         string `json:"path"`
	Layer        string `json:"layer"`
	UpdateType   string `json:"update_type"`
	OwnerThread  string `json:"owner_thread,omitempty"`
	TargetThread string `json:"target_thread,omitempty"`
	Reason       string `json:"reason,omitempty"`
}
