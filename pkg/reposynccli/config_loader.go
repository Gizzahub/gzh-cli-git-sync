package reposynccli

import (
	"context"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Gizzahub/gzh-cli-git-sync/pkg/reposync"
)

// ConfigData contains plan + run inputs loaded from a config file.
type ConfigData struct {
	Plan reposync.PlanRequest
	Run  reposync.RunOptions
}

// SpecLoader loads repository specs from a source (e.g., YAML file).
type SpecLoader interface {
	Load(ctx context.Context, path string) (ConfigData, error)
}

// FileSpecLoader loads configuration from a YAML file on disk.
type FileSpecLoader struct {
	// Optional defaults if the file omits values.
	DefaultStrategy reposync.Strategy
	DefaultParallel int
	DefaultRetries  int
}

type fileConfig struct {
	Strategy       string      `yaml:"strategy"`
	Parallel       int         `yaml:"parallel"`
	MaxRetries     int         `yaml:"maxRetries"`
	Resume         bool        `yaml:"resume"`
	DryRun         bool        `yaml:"dryRun"`
	CleanupOrphans bool        `yaml:"cleanupOrphans"`
	Repositories   []repoEntry `yaml:"repositories"`
}

type repoEntry struct {
	Name          string `yaml:"name"`
	Provider      string `yaml:"provider"`
	URL           string `yaml:"url"`
	TargetPath    string `yaml:"targetPath"`
	Strategy      string `yaml:"strategy"`
	AssumePresent bool   `yaml:"assumePresent"`
}

// Load implements SpecLoader.
func (l FileSpecLoader) Load(_ context.Context, path string) (ConfigData, error) {
	if path == "" {
		return ConfigData{}, errors.New("config path is required")
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return ConfigData{}, fmt.Errorf("read config: %w", err)
	}

	var cfg fileConfig
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return ConfigData{}, fmt.Errorf("parse config: %w", err)
	}

	if len(cfg.Repositories) == 0 {
		return ConfigData{}, errors.New("config has no repositories")
	}

	defaultStrategy := l.DefaultStrategy
	if defaultStrategy == "" {
		defaultStrategy = reposync.StrategyReset
	}

	defaultParallel := l.DefaultParallel
	if defaultParallel <= 0 {
		defaultParallel = 4
	}

	defaultRetries := l.DefaultRetries
	if defaultRetries < 0 {
		defaultRetries = 1
	}

	parsedStrategy, err := reposync.ParseStrategy(cfg.Strategy)
	if err != nil {
		return ConfigData{}, err
	}
	if cfg.Strategy == "" {
		parsedStrategy = defaultStrategy
	}

	plan := reposync.PlanRequest{
		Input: reposync.PlanInput{
			Repos: make([]reposync.RepoSpec, 0, len(cfg.Repositories)),
		},
		Options: reposync.PlanOptions{
			DefaultStrategy: parsedStrategy,
			CleanupOrphans:  cfg.CleanupOrphans,
		},
	}

	for _, repo := range cfg.Repositories {
		if repo.Name == "" || repo.URL == "" || repo.TargetPath == "" {
			return ConfigData{}, fmt.Errorf("repository entry is missing required fields (name/url/targetPath)")
		}

		repoStrategy := parsedStrategy
		if repo.Strategy != "" {
			repoStrategy, err = reposync.ParseStrategy(repo.Strategy)
			if err != nil {
				return ConfigData{}, fmt.Errorf("repository %s: %w", repo.Name, err)
			}
		}

		plan.Input.Repos = append(plan.Input.Repos, reposync.RepoSpec{
			Name:          repo.Name,
			Provider:      repo.Provider,
			CloneURL:      repo.URL,
			TargetPath:    repo.TargetPath,
			Strategy:      repoStrategy,
			AssumePresent: repo.AssumePresent,
		})
	}

	run := reposync.RunOptions{
		Parallel:   cfg.Parallel,
		MaxRetries: cfg.MaxRetries,
		Resume:     cfg.Resume,
		DryRun:     cfg.DryRun,
	}

	if run.Parallel <= 0 {
		run.Parallel = defaultParallel
	}
	if run.MaxRetries < 0 {
		run.MaxRetries = defaultRetries
	}

	return ConfigData{
		Plan: plan,
		Run:  run,
	}, nil
}
