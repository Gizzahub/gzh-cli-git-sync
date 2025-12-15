# gzh-cli-git-sync

Repository synchronization engine for the gzh ecosystem.

## Goals

- Library-first: plan/execute/state/progress APIs to manage many repositories.
- Shared CLI: usable both as standalone `gz-git-sync` and embedded as `gz git-sync`.
- Provider-agnostic: takes repo specs from callers (e.g., gitforge) instead of assuming one hosting service.

## High-level pieces

- `plan`: turn desired repos + local state into actionable steps (clone/pull/fetch/delete).
- `executor`: run steps with strategies, parallelism, retries, and hooks.
- `state`: persist/restore runs for resume and auditing.
- `progress`: surface events for CLI/metrics/UX.

This module purposely does **not** bundle the entire `synclone` implementation from `gzh-cli`; instead it rethinks the orchestration layer to be reusable outside the main CLI.

- Default executor uses [`gzh-cli-git`](https://github.com/Gizzahub/gzh-cli-git) to perform real Git operations.

## CLI (shared)

The `pkg/reposynccli` package provides a Cobra command tree that can be used as:

- Standalone binary: `gz-git-sync run --config git-sync.yaml`
- Embedded: attach `CommandFactory.NewRootCmd()` as a subcommand under `gz` (e.g., `gz git-sync`).

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
