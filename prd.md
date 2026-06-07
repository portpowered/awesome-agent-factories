# PRD: Phase 5 Workflow Integration Repair

## Introduction

Converge the completed `phase-5-link-awesome-maintenance-workflows` GitHub Actions work onto `main` before Phase 6 starts. Local `main` already includes the merged core `.github/workflows/ci.yml` from `origin/main` (via `phase-5-core-ci-workflow`), but the remaining Phase 5 maintenance workflows and contributor documentation updates exist only on `origin/phase-5-link-awesome-maintenance-workflows`. This repair selectively integrates those artifacts onto `main`—link checking, awesome-list linting, scheduled maintenance, and the `CONTRIBUTING.md` GitHub Actions section—while preserving the existing core CI workflow and leaving Phase 6 templates and README content out of scope.

## Context

### Customer ask

Phase 5 integration repair: converge the completed GitHub Actions work onto `main` before Phase 6 starts. Local `main` has the merged core `.github/workflows/ci.yml` from `origin/main`, but the completed `phase-5-link-awesome-maintenance-workflows` output is still only on its branch. Integrate the remaining Phase 5 outputs into `main`: `.github/workflows/link-check.yml`, `.github/workflows/awesome-lint.yml`, `.github/workflows/scheduled-maintenance.yml`, and the `CONTRIBUTING.md` GitHub Actions documentation if still applicable. Preserve the already-integrated core CI workflow. Acceptance: all four requested workflow files exist on `main`; link check runs on pull requests to `main` and weekly; awesome lint runs on pull requests to `main` and pushes to `main`; scheduled maintenance runs monthly, performs read-only custom README checks and link checks, and does not mutate repository content; workflow commands remain aligned with local Makefile/config behavior; local `make check`, `make test`, `go test ./...`, and `git diff --check` pass.

### Problem

Phase 5 is split across two merged/in-flight branches. Core CI (`ci.yml`) is on `main`, but link checking, awesome-list linting, and scheduled maintenance workflows remain only on `phase-5-link-awesome-maintenance-workflows`. Contributors reading `CONTRIBUTING.md` on `main` are told GitHub Actions are not configured, while the completed branch already documents workflow-to-local-command mappings. A naive merge risks clobbering the merged `ci.yml` or reintroducing planner tracking-file conflicts; a partial merge leaves Phase 5 incomplete and blocks Phase 6.

### Solution

Create an integration repair branch (`phase-5-workflow-integration-repair`) from current `main`, selectively bring in the three maintenance workflow files and the `CONTRIBUTING.md` GitHub Actions section from `origin/phase-5-link-awesome-maintenance-workflows`, verify `ci.yml` is unchanged, and confirm all workflows are read-only with triggers and commands aligned to the root `Makefile` and `.lychee.toml`. Verify convergence with `make check`, `make test`, `go test ./...`, and `git diff --check` before merging to `main`.

## Goals

- Land all four Phase 5 workflow files on `main`: `ci.yml` (preserved), `link-check.yml`, `awesome-lint.yml`, and `scheduled-maintenance.yml`.
- Ensure workflow triggers match the customer ask: link check on PRs to `main` and weekly; awesome lint on PRs and pushes to `main`; scheduled maintenance monthly with read-only README and link checks.
- Keep workflows non-mutating and aligned with local commands (`make check`, `make links`, `npx awesome-lint`) without duplicating checker or lychee rules in YAML.
- Update `CONTRIBUTING.md` so contributors can map CI failures to local reproduction commands.
- Preserve the merged core CI workflow and pass all local quality gates before Phase 6.

## Project-Level Acceptance Criteria

- [ ] All four workflow files exist on the integrated branch: `.github/workflows/ci.yml`, `.github/workflows/link-check.yml`, `.github/workflows/awesome-lint.yml`, and `.github/workflows/scheduled-maintenance.yml`.
- [ ] `link-check.yml` triggers on `pull_request` to `main` and a weekly `schedule` cron; it scans `README.md` and `docs/*.md` with lychee honoring root `.lychee.toml` and performs no content mutation.
- [ ] `awesome-lint.yml` triggers on `pull_request` and `push` to `main`; it runs `npx awesome-lint` (or equivalent) and performs no content mutation.
- [ ] `scheduled-maintenance.yml` triggers on a monthly `schedule` cron; it runs custom README checks (`make check` or `go run ./internal/checks`) and link checks; it performs no commits, auto-fixes, or other repository mutations.
- [ ] `.github/workflows/ci.yml` is preserved from current `main` (Go format check, `go test ./...`, README checks via `make check` / `go run ./internal/checks` on PR and push to `main`).
- [ ] `CONTRIBUTING.md` documents the three maintenance workflows and maps each to its local equivalent command.
- [ ] Workflow step names or comments identify local equivalents (`make links`, `make check`, `npx awesome-lint`); checker and lychee settings are not duplicated in workflow YAML.
- [ ] Quality gate: from repository root on the integrated branch, `make check`, `make test`, `go test ./...`, and `git diff --check` all exit 0.

## User Stories

### phase-5-workflow-integration-repair-001: Integrate link-check workflow onto main

**Description:** As a maintainer, I want link checking enforced on pull requests to `main` and on a weekly schedule so broken links in `README.md` and documentation are caught before or soon after merge.

**Acceptance Criteria:**

- [ ] `.github/workflows/link-check.yml` exists on the integration branch with workflow name `Link Check`.
- [ ] Workflow `on` includes `pull_request` with `branches: [main]` and `schedule` with weekly cron (`0 8 * * 1` or equivalent Monday schedule).
- [ ] A job on `ubuntu-latest` checks out the repository with `actions/checkout@v4` and runs lychee on `README.md` and `docs/*.md` via `lycheeverse/lychee-action@v2` (or equivalent) so root `.lychee.toml` settings apply.
- [ ] Step names or comments identify the local equivalent (`make links`).
- [ ] The workflow performs no file writes, commits, or content mutation.
- [ ] Typecheck passes

### phase-5-workflow-integration-repair-002: Integrate awesome-lint workflow onto main

**Description:** As a maintainer, I want awesome-list lint rules enforced on pull requests and pushes to `main` so list structure violations are blocked before they land on the default branch.

**Acceptance Criteria:**

- [ ] `.github/workflows/awesome-lint.yml` exists on the integration branch with workflow name `Awesome Lint`.
- [ ] Workflow `on` includes `pull_request` with `branches: [main]` and `push` with `branches: [main]`.
- [ ] A job on `ubuntu-latest` checks out the repository with `actions/checkout@v4` and runs `npx awesome-lint`.
- [ ] Step names or comments identify the local equivalent (`npx awesome-lint`).
- [ ] The workflow performs no file writes, commits, or content mutation.
- [ ] Typecheck passes

### phase-5-workflow-integration-repair-003: Integrate scheduled maintenance workflow onto main

**Description:** As a maintainer, I want a monthly scheduled job that re-validates README rules and link health on `main` so drift is detected even when no pull requests are open.

**Acceptance Criteria:**

- [ ] `.github/workflows/scheduled-maintenance.yml` exists on the integration branch with workflow name `Scheduled Maintenance`.
- [ ] Workflow `on` includes `schedule` with monthly cron (`0 9 1 * *` or equivalent first-of-month schedule).
- [ ] A job on `ubuntu-latest` checks out the repository, sets up Go with `actions/setup-go@v5` (`go-version: stable`), and runs custom README checks via `make check` (preferred) or `go run ./internal/checks`.
- [ ] The same job runs link checks on `README.md` and `docs/*.md` using the same lychee approach as `link-check.yml`, honoring `.lychee.toml`.
- [ ] Step names or comments identify local equivalents (`make check`, `make links`).
- [ ] The workflow performs no file writes, commits, issue bots, or auto-fix steps.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-5-workflow-integration-repair-004: Update CONTRIBUTING.md with GitHub Actions documentation

**Description:** As a contributor, I want `CONTRIBUTING.md` to document which GitHub workflows run and how they map to local commands so I can reproduce CI failures before pushing.

**Acceptance Criteria:**

- [ ] `CONTRIBUTING.md` replaces the statement that GitHub Actions are not configured with a `### GitHub Actions` subsection under Local checks.
- [ ] The subsection includes a table mapping Link Check, Awesome Lint, and Scheduled Maintenance workflows to their triggers and local equivalents (`make links`, `npx awesome-lint`, `make check` and `make links`).
- [ ] Documentation states workflows are read-only and that link checks honor `.lychee.toml` while README rules live in `internal/checks` (not duplicated in workflow YAML).
- [ ] Links in the table point to `.github/workflows/link-check.yml`, `.github/workflows/awesome-lint.yml`, and `.github/workflows/scheduled-maintenance.yml`.
- [ ] No README list entries or unrelated governance prose changes are introduced.
- [ ] Typecheck passes

### phase-5-workflow-integration-repair-005: Preserve core CI workflow during integration

**Description:** As a maintainer, I want the already-merged core CI workflow preserved unchanged so Phase 5 integration does not regress Go formatting, test, or README check gates on pull requests and pushes to `main`.

**Acceptance Criteria:**

- [ ] `.github/workflows/ci.yml` on the integration branch matches current `main` (workflow name `CI`; triggers on `pull_request` and `push` to `main`; Go format check on `internal/`; `go test ./...`; README checks via `make check` with `go run ./internal/checks` fallback).
- [ ] No integration change modifies `ci.yml` step commands, action versions, or trigger configuration unless fixing a merge conflict in favor of the `main` version.
- [ ] Core CI workflow remains read-only and non-mutating.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-5-workflow-integration-repair-006: Verify Phase 5 workflow convergence on integrated main

**Description:** As a planner preparing the Phase 5 loopback review, I want end-to-end verification that all four workflows exist, triggers and read-only behavior match the customer ask, and local quality gates pass so Phase 6 can start from a converged `main`.

**Acceptance Criteria:**

- [ ] All four workflow files exist on the integration branch: `ci.yml`, `link-check.yml`, `awesome-lint.yml`, and `scheduled-maintenance.yml`.
- [ ] Observable trigger configuration matches the customer ask (link check: PR + weekly; awesome lint: PR + push to `main`; scheduled maintenance: monthly; core CI: PR + push to `main`).
- [ ] From repository root: `make check`, `make test`, `go test ./...`, and `git diff --check` all exit 0.
- [ ] No issue templates, PR templates, or bulk README resource entries are added as part of this integration.
- [ ] Only `.github/workflows/link-check.yml`, `.github/workflows/awesome-lint.yml`, `.github/workflows/scheduled-maintenance.yml`, and narrowly scoped `CONTRIBUTING.md` updates are introduced—no unrelated cleanup.
- [ ] Typecheck passes
- [ ] Tests pass

## High-Level Technical Design

Integration follows a selective file-level merge from `origin/phase-5-link-awesome-maintenance-workflows` onto current `main`:

```
main (has ci.yml)
  │
  ├── cherry-pick / copy: link-check.yml, awesome-lint.yml, scheduled-maintenance.yml
  ├── patch: CONTRIBUTING.md Local checks → GitHub Actions subsection
  └── preserve: ci.yml unchanged
```

**Workflow ownership:**

| Workflow | Triggers | Local equivalent | Config source |
| --- | --- | --- | --- |
| `ci.yml` | PR + push to `main` | `make check`, `go test ./...`, gofmt | `Makefile`, `internal/checks` |
| `link-check.yml` | PR to `main` + weekly | `make links` | `.lychee.toml` |
| `awesome-lint.yml` | PR + push to `main` | `npx awesome-lint` | awesome-lint defaults |
| `scheduled-maintenance.yml` | Monthly | `make check`, `make links` | `Makefile`, `.lychee.toml`, `internal/checks` |

**Merge conflict policy:** If `ci.yml` conflicts, keep `main`'s version. If `CONTRIBUTING.md` conflicts, merge the GitHub Actions subsection from the candidate branch while preserving any newer `main` prose outside Local checks.

**Read-only guarantees:** None of the workflows use write permissions, `gofmt -w`, commit bots, or auto-fix actions. Scheduled maintenance reports failures without mutating tracked content.

## Functional Requirements

- FR-1: Integrate `.github/workflows/link-check.yml` with PR-to-`main` and weekly schedule triggers.
- FR-2: Integrate `.github/workflows/awesome-lint.yml` with PR-to-`main` and push-to-`main` triggers.
- FR-3: Integrate `.github/workflows/scheduled-maintenance.yml` with monthly schedule, README checks, and link checks.
- FR-4: Update `CONTRIBUTING.md` to document workflow-to-local-command mapping.
- FR-5: Preserve `.github/workflows/ci.yml` from current `main` without regression.
- FR-6: Verify local gates (`make check`, `make test`, `go test ./...`, `git diff --check`) pass on the integrated branch.

## Non-Goals

- No changes to `ci.yml` beyond conflict resolution in favor of `main`.
- No issue templates, PR templates, or Phase 6 deliverables.
- No README content entries or bulk documentation rewrites.
- No duplication of checker logic or lychee accept/exclude rules inside workflow YAML.
- No workflow steps that mutate repository content (commits, auto-fix, format-with-write).

## Supporting Technical and UX Considerations

- Use maintained action versions already validated on the candidate branch (`actions/checkout@v4`, `actions/setup-go@v5`, `lycheeverse/lychee-action@v2`).
- Quote workflow step names containing colons to keep YAML valid (lesson from `phase-5-core-ci-workflow`).
- `CONTRIBUTING.md` is the contributor-facing map from CI failure to local reproduction; keep the table concise and link to workflow files.
- Integration should not replace `prd.json`, `prd.md`, or `progress.txt` on `main` with older branch snapshots unless an explicit planner update is required.

## Success Metrics

- All four Phase 5 workflow files exist on `main` after merge.
- Contributors can reproduce any maintenance workflow failure locally using documented `make` or `npx` commands.
- Local quality gates pass with zero regressions to checker behavior or core CI.
- Phase 5 checklist in `docs/internal/checklist.md` can be marked converged for workflow integration.

## Open Questions

- None. The candidate branch `origin/phase-5-link-awesome-maintenance-workflows` is validated and the integration surface is a straightforward selective merge with `ci.yml` preservation.
