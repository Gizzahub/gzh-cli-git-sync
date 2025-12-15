package main

import (
	"fmt"
	"os"

	"github.com/Gizzahub/gzh-cli-reposync/pkg/reposync"
	"github.com/Gizzahub/gzh-cli-reposync/pkg/reposynccli"
)

func main() {
	planner := reposync.StaticPlanner{}
	executor := reposync.NoopExecutor{}
	state := reposync.NewInMemoryStateStore()
	orchestrator := reposync.NewOrchestrator(planner, executor, state)

	factory := reposynccli.CommandFactory{
		Use:          "gz-reposync",
		Short:        "Repository synchronization",
		Orchestrator: orchestrator,
		SpecLoader:   reposynccli.FileSpecLoader{},
	}

	root := factory.NewRootCmd()
	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
