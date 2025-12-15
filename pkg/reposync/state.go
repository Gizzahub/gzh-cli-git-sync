package reposync

import "context"

// StateStore persists run state for resume/audit.
type StateStore interface {
	Save(ctx context.Context, state RunState) error
	Load(ctx context.Context) (RunState, error)
}

// RunState captures progress for resuming operations.
type RunState struct {
	Items []RunStateItem
}

// RunStateItem tracks per-repo status.
type RunStateItem struct {
	Repo    RepoSpec
	Status  RunStatus
	Message string
}

// RunStatus represents the last known state of a repository.
type RunStatus string

const (
	RunStatusPending RunStatus = "pending"
	RunStatusRunning RunStatus = "running"
	RunStatusDone    RunStatus = "done"
	RunStatusFailed  RunStatus = "failed"
)
