package reposync

import "context"

// Planner produces a Plan from desired repositories and current state.
// Concrete implementation will live in future steps; this placeholder defines
// the interface surface for consumers and CLI wiring.
type Planner interface {
	Plan(ctx context.Context, input PlanInput) (Plan, error)
}

// PlanInput captures desired repositories and optional context (e.g., host
// aliases, path rules). It is intentionally minimal for now; richer fields
// will be added in follow-up steps.
type PlanInput struct {
	Repos []RepoSpec
}

// Plan is the result of planning (e.g., clone/pull/fetch/delete actions).
// Details will be expanded as the orchestration logic lands.
type Plan struct {
	Actions []Action
}

// RepoSpec describes a repository to manage.
type RepoSpec struct {
	Name       string
	Provider   string
	CloneURL   string
	TargetPath string
}

// Action describes a single operation in a plan.
type Action struct {
	Repo   RepoSpec
	Type   ActionType
	Reason string
}

// ActionType enumerates planned operations.
type ActionType string

const (
	ActionClone  ActionType = "clone"
	ActionUpdate ActionType = "update"
	ActionSkip   ActionType = "skip"
	ActionDelete ActionType = "delete"
)
