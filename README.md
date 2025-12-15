# gzh-cli-reposync

Repository synchronization engine for gzh ecosystem.

## Goals

- Library-first: plan/execute/state/progress APIs to manage many repositories.
- Shared CLI: usable both as standalone `gz-reposync` and embedded as `gz reposync`.
- Provider-agnostic: takes repo specs from callers (e.g., gitforge) instead of assuming one hosting service.

## High-level pieces

- `plan`: turn desired repos + local state into actionable steps (clone/pull/fetch/delete).
- `executor`: run steps with strategies, parallelism, retries, and hooks.
- `state`: persist/restore runs for resume and auditing.
- `progress`: surface events for CLI/metrics/UX.

This module purposely does **not** bundle the entire `synclone` implementation from `gzh-cli`; instead it rethinks the orchestration layer to be reusable outside the main CLI.

## CLI (shared)

The `pkg/reposynccli` package provides a Cobra command tree that can be used as:

- Standalone binary: `gz-reposync run --config reposync.yaml`
- Embedded: attach `CommandFactory.NewRootCmd()` as a subcommand under `gz`.

### Minimal config example

```yaml
strategy: reset          # default strategy (reset|pull|fetch)
parallel: 4
maxRetries: 1
dryRun: true
repositories:
  - name: example
    provider: github
    url: https://github.com/org/example.git
    targetPath: ./repos/example
```
