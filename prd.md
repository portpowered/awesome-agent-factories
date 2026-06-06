# PRD: Phase 4 Go README Checker Core

## Context

### Customer ask

Phase 4 automation: add the root Go module and custom README checker requested in `docs/internal/customer-ask.md`. Implement `go.mod`, `go.sum` if needed, and `internal/checks` with focused code split across `main.go`, `readme.go`, `links.go`, `sections.go`, `alphabetical.go`, and `duplicates.go` or an equivalent small Go package structure that preserves those responsibilities. The checker must validate that README exists; required sections exist; Contents links match required headings; README entries use `- [Name](URL) - Description.`; descriptions end with a period; duplicate URLs and duplicate resource names are rejected; entries are alphabetized within sections; bare URLs are rejected; tracking-heavy or suspicious URLs are reported; marketing phrases such as `revolutionary`, `game-changing`, `best`, `ultimate`, and `cutting-edge` are rejected unless part of a resource title; and scope keywords are warnings rather than hard failures. Keep empty starter sections allowed until the content phase fills them. Include deterministic Go tests or table tests for the checker behavior using in-memory or fixture markdown. Do not add GitHub Actions, issue templates, or bulk README resource entries in this item. Acceptance: `go test ./...` passes, `go run ./internal/checks` passes against the current README, failures explain the specific resource-entry rule, and existing Phase 1-3 docs remain aligned.

### Problem

Phases 1–3 established README structure, governance, and manual review policy, but contributors and maintainers still rely on hand review to catch formatting mistakes, duplicate links, alphabetization errors, promotional language, and weak scope signals. There is no executable enforcement of the rules documented in `CONTRIBUTING.md` and `docs/review-policy.md`. Without a deterministic local checker, Phase 4 automation cannot converge and Phase 5 CI cannot be wired safely.

### Solution

Add a root Go module (`github.com/portpowered/awesome-agent-factories`) and an `internal/checks` command package that reads the repository `README.md`, runs focused validations aligned with Phase 1–3 policy, emits actionable hard failures and non-blocking warnings, and exits non-zero only on hard failures. Cover behavior with table-driven Go tests using in-memory markdown fixtures. Keep resource sections empty until Phase 7 content seeding.

## Introduction

This work item delivers the **core custom README checker** for **Awesome AI Agent Factories**. It is the first Phase 4 deliverable and a prerequisite for Makefile/CI integration (`phase-4-makefile-and-lint-config`). The checker encodes CONTRIBUTING and review-policy rules as observable CLI behavior: pass, fail with a specific rule citation, or warn without failing.

## Project-level acceptance criteria

- [ ] Root `go.mod` (and `go.sum` if dependencies are used) exist; module path is `github.com/portpowered/awesome-agent-factories`.
- [ ] `internal/checks` is split across `main.go`, `readme.go`, `links.go`, `sections.go`, `alphabetical.go`, and `duplicates.go` (or equivalent files preserving the same responsibilities).
- [ ] `go test ./...` passes with deterministic table tests covering hard-failure and warning behaviors.
- [ ] `go run ./internal/checks` exits 0 against the current repository `README.md` (empty resource sections are allowed).
- [ ] Hard failures name the violated resource-entry rule (section, entry line, and rule such as duplicate URL, bad format, or marketing phrase).
- [ ] Scope-keyword and suspicious/tracking URL findings are warnings only and do not cause a non-zero exit when no hard failures exist.
- [ ] No GitHub Actions workflows, issue/PR templates, Makefile, bulk README entries, or unrelated repository refactors are added.
- [ ] Quality gate: `go test ./...` passes; `go build ./internal/checks` succeeds; checker behavior remains aligned with Phase 1–3 docs (`CONTRIBUTING.md`, `docs/review-policy.md`, `README.md` scope and headings).

## Goals

- Encode README structure and entry-format rules from Phase 1–3 as executable checks.
- Give contributors fast, specific feedback before maintainer review.
- Allow empty starter resource sections until Phase 7 content seeding.
- Separate hard failures (must fix) from warnings (scope relevance, tracking URLs).
- Provide test-backed confidence for future Makefile and CI wiring.

## User stories

### phase-4-go-readme-checker-core-001: Bootstrap Go module and checks package layout

**Description:** As a maintainer introducing Phase 4 automation, I want a root Go module and an `internal/checks` package skeleton so later validation stories have a stable build and ownership boundary.

**Acceptance Criteria:**

- [ ] `go.mod` exists at repository root with module path `github.com/portpowered/awesome-agent-factories` and a supported Go toolchain version.
- [ ] Package `internal/checks` exists with files `main.go`, `readme.go`, `links.go`, `sections.go`, `alphabetical.go`, and `duplicates.go` (or equivalent split preserving those responsibilities).
- [ ] `go build ./internal/checks` succeeds from repository root.
- [ ] Typecheck passes

### phase-4-go-readme-checker-core-002: Load README and parse document structure

**Description:** As a contributor running the checker locally, I want a clear failure when `README.md` is missing and reliable parsing of headings and sections so downstream rules operate on structured content.

**Acceptance Criteria:**

- [ ] Running the checker when `README.md` is absent exits non-zero with a message stating the README is missing.
- [ ] The checker reads `README.md` from the repository root (or an explicit path flag if provided) and identifies `##` section headings and their body content.
- [ ] Table tests use in-memory markdown fixtures to verify missing-README failure and basic section extraction.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-4-go-readme-checker-core-003: Validate required sections and Contents-heading alignment

**Description:** As a maintainer guarding list structure, I want required README sections and Contents anchor links to match headings so the Phase 1 skeleton cannot drift silently.

**Acceptance Criteria:**

- [ ] Hard failure when any required resource section heading is missing: Theories, Coordination Patterns, Frameworks, Protocols and Interfaces, Benchmarks, Research Papers, Blog Posts, Case Studies, Examples and Templates, Related Lists.
- [ ] Hard failure when `## Contents` is missing or when a Contents list item anchor does not match an existing `##` heading (and vice versa for required resource sections listed in Contents).
- [ ] Empty resource sections (heading plus optional intro prose, zero `- [` entries) pass without error.
- [ ] Table tests cover missing section, missing Contents entry, and valid empty-section fixture.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-4-go-readme-checker-core-004: Enforce resource entry list-item format

**Description:** As a contributor adding a resource, I want invalid list-item shapes rejected with a specific format error so I can fix entries to match `- [Name](URL) - Description.`

**Acceptance Criteria:**

- [ ] Each resource entry in a resource section must match the pattern `- [Name](URL) - Description` (single markdown link, hyphen separator, description text).
- [ ] Hard failure for malformed lines in resource sections (bare bullets, missing link, missing description, extra links on the same line, or wrong separator) with a message identifying the section, approximate line, and expected format rule.
- [ ] Non-entry lines (section intro paragraphs, blank lines) in resource sections do not false-positive as entries.
- [ ] Table tests cover valid entry, missing period (handled in next story), and multiple invalid shapes.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-4-go-readme-checker-core-005: Enforce description punctuation and banned marketing phrases

**Description:** As a maintainer enforcing encyclopedic tone, I want descriptions to end with a period and promotional phrases rejected unless they appear in the official resource title (link text).

**Acceptance Criteria:**

- [ ] Hard failure when a resource description does not end with `.` (period).
- [ ] Hard failure when description text (after the hyphen separator, excluding link text) contains banned marketing phrases case-insensitively: `revolutionary`, `game-changing`, `best`, `ultimate`, `cutting-edge`.
- [ ] Banned phrase in link text `[Name]` is allowed when it is part of the official resource title; the same phrase only in the description still fails.
- [ ] Failure messages cite the entry and the specific rule (missing period vs banned phrase).
- [ ] Table tests cover period violation, phrase in description, and phrase allowed in title only.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-4-go-readme-checker-core-006: Reject duplicate URLs and duplicate resource names

**Description:** As a maintainer preventing duplicate submissions, I want duplicate URLs and duplicate link-text names rejected anywhere in README resource entries.

**Acceptance Criteria:**

- [ ] Hard failure when the same normalized URL appears on more than one resource entry (across any sections).
- [ ] Hard failure when the same resource name (link text) appears on more than one entry, compared case-insensitively unless names differ only by case per CONTRIBUTING guidance.
- [ ] Failure messages identify both colliding entries (section and name/URL).
- [ ] Table tests cover duplicate URL, duplicate name, and clean distinct entries.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-4-go-readme-checker-core-007: Enforce alphabetical order within each resource section

**Description:** As a contributor inserting a resource, I want alphabetization violations reported per section so I can place entries in the correct order by link text.

**Acceptance Criteria:**

- [ ] Hard failure when entries in a resource section are not sorted alphabetically by link text using case-insensitive comparison.
- [ ] Failure message names the section and the out-of-order adjacent entries.
- [ ] Sections with zero or one entry pass alphabetization check.
- [ ] Table tests cover in-order, out-of-order, and single-entry sections.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-4-go-readme-checker-core-008: Reject bare URLs and warn on tracking-heavy or suspicious links

**Description:** As a maintainer protecting link quality, I want bare URLs rejected and tracking-heavy or suspicious URLs surfaced as warnings without failing an otherwise valid README.

**Acceptance Criteria:**

- [ ] Hard failure for bare URLs in resource sections (http/https URLs not inside `[text](url)` markdown link syntax).
- [ ] Warning (not hard failure) when a resource URL contains common tracking or suspicious query parameters or patterns (for example `utm_`, `fbclid`, `gclid`, `mc_eid`, or excessive redirect/shortener hosts documented in implementation).
- [ ] Warnings are printed to stderr (or equivalent) with entry context; exit code remains 0 when only warnings are present.
- [ ] Table tests cover bare URL failure, tracking URL warning, and clean canonical URL.
- [ ] Typecheck passes
- [ ] Tests pass

### phase-4-go-readme-checker-core-009: Emit scope-keyword warnings without failing valid structure

**Description:** As a maintainer nudging agent-factory relevance, I want scope-related keyword gaps reported as warnings so contributors strengthen descriptions without blocking empty starter sections.

**Acceptance Criteria:**

- [ ] For each resource entry with a description, emit a warning when the description does not contain any positive scope keyword from the CONTRIBUTING agent-factory angle set (coordination, orchestration, delegation, routing, handoffs, shared state, group evaluation / group-level evaluation, multi-agent, agent group, or close stems).
- [ ] Scope keyword findings are warnings only; they do not cause non-zero exit by themselves.
- [ ] Entries in empty sections produce no scope warnings (no entries to evaluate).
- [ ] Table tests cover description with scope keyword (no warning), without scope keyword (warning), and empty section (no warning).
- [ ] Typecheck passes
- [ ] Tests pass

### phase-4-go-readme-checker-core-010: CLI integration passes current README with aligned policy docs

**Description:** As a factory operator completing Phase 4 core work, I want `go run ./internal/checks` to run all checks against the real README with clear aggregated output so the repository is ready for Makefile/CI wiring.

**Acceptance Criteria:**

- [ ] `go run ./internal/checks` from repository root runs structure, format, duplicate, alphabetization, URL, marketing, and scope checks in one command.
- [ ] Exit code 0 when the current `README.md` has no hard failures (empty resource sections allowed).
- [ ] Exit code non-zero when any hard failure exists; stderr lists each failure with a specific resource-entry rule message.
- [ ] `go test ./...` passes for the full `internal/checks` package.
- [ ] Phase 1–3 docs are not contradicted: checker rules match `CONTRIBUTING.md` entry format, marketing prohibitions, alphabetization, and duplicate policy; no doc claims GitHub Actions or CI already enforce checks (Makefile/CI remain out of scope for this item).
- [ ] Typecheck passes
- [ ] Tests pass

## Functional requirements

- **FR-1:** Provide root `go.mod` with module `github.com/portpowered/awesome-agent-factories`.
- **FR-2:** Implement `internal/checks` as a runnable command package with responsibilities split across `main.go`, `readme.go`, `links.go`, `sections.go`, `alphabetical.go`, and `duplicates.go`.
- **FR-3:** Fail when `README.md` is missing or unreadable from the expected repository root path.
- **FR-4:** Require `## Contents` and the ten Phase 1 resource section headings; Contents anchor links must match those headings.
- **FR-5:** Allow resource sections with no `- [` entries (empty starter sections).
- **FR-6:** Validate each resource entry matches `- [Name](URL) - Description.` structure.
- **FR-7:** Require descriptions to end with `.`; reject banned marketing phrases in description text unless only present in link text title.
- **FR-8:** Reject duplicate URLs and duplicate resource names across the README.
- **FR-9:** Require case-insensitive alphabetical order by link text within each resource section.
- **FR-10:** Reject bare URLs in resource sections.
- **FR-11:** Warn on tracking-heavy or suspicious URLs without failing exit status alone.
- **FR-12:** Warn when descriptions lack positive agent-factory scope keywords; never treat scope warnings as hard failures.
- **FR-13:** Emit human-readable messages identifying section, entry, and violated rule for every hard failure.
- **FR-14:** Cover behaviors with deterministic Go table tests using in-memory or fixture markdown (no network calls).

## Non-goals

- GitHub Actions workflows (Phase 5).
- Root `Makefile`, `.golangci.yml`, `.markdownlint.yaml`, `.lychee.toml`, or shell wrapper scripts (`phase-4-makefile-and-lint-config`).
- Issue templates or PR templates (Phase 6).
- Seeding README resource entries or enforcing minimum entries per section (Phase 7).
- Live HTTP link checking (deferred to lychee/Makefile `links` target).
- Rewriting Phase 1 README scope prose or Phase 2/3 governance documents beyond alignment verification.
- Enforcing new-category minimum-three-entries rule until sections contain entries.

## High-level technical design

### Package ownership

| File | Responsibility |
|------|----------------|
| `main.go` | CLI entry, flag parsing (optional README path), orchestration, exit codes, stderr formatting for failures vs warnings |
| `readme.go` | Load README bytes, top-level document model, marketing-phrase and scope-keyword analysis helpers |
| `sections.go` | Required headings, Contents parsing, section body boundaries, entry line extraction |
| `links.go` | Entry format regex/parse, URL normalization, bare URL detection, tracking/suspicious URL heuristics |
| `alphabetical.go` | Per-section sort validation by link text |
| `duplicates.go` | Global URL and name collision detection |

### State and side effects

- **Input:** Read-only file read of `README.md` (default: repository root relative to working directory or module root).
- **Output:** Structured findings slice `{Severity: error|warning, Rule, Section, Line, Message}` printed and aggregated; `os.Exit(1)` only if any error-severity finding exists.
- **No network I/O** in core checker tests or default run.

### Contracts

- Resource sections are the ten Phase 1 category headings under `##`.
- Entry lines begin with `- [` in resource section bodies.
- URL normalization for duplicate detection: lowercase scheme/host, trim trailing slash, strip fragment (document choice in implementation; test-backed).

### Test layer

- Table-driven `_test.go` files colocated in `internal/checks`.
- Fixtures as multiline string constants in tests (preferred) or small `testdata/*.md` files.
- Each story’s behavioral criteria maps to at least one test case before story completion.

### Dependency order

1. Module bootstrap → 2. Parse README → 3. Sections/Contents → 4. Entry format → 5. Punctuation/marketing → 6. Duplicates → 7. Alphabetize → 8. URL hygiene → 9. Scope warnings → 10. CLI integration on real README.

## Supporting technical and UX considerations

- Prefer standard library only unless a tiny dependency materially simplifies markdown parsing; if added, pin in `go.sum`.
- Error messages should mirror CONTRIBUTING vocabulary ("description must end with a period", "duplicate URL", "not alphabetized by link text") so contributors can self-serve fixes.
- Warnings should be prefixed distinctly (for example `warning:`) from hard failures (`error:`).
- Do not require entries in empty sections during Phase 4; scope-keyword warnings apply only when entries exist.
- `docs/review-policy.md` already states automation is future work; this item adds the checker binary but does not require editing review docs unless an implementer finds a direct contradiction—default is no doc churn.
- Running from repository root is the primary UX; optional `-readme path` flag is acceptable for tests and local debugging.

## Success metrics

- `go test ./...` and `go run ./internal/checks` pass on a clean checkout with current Phase 1 README skeleton.
- A deliberately invalid fixture entry produces a failure message naming the exact rule within one run.
- Warnings alone never block exit code 0.
- Phase 4 Makefile story can wire `check:` to `go run ./internal/checks` without code changes to validation rules.

## Open questions

None. Empty-section allowance, warning vs failure split, file layout, and acceptance commands are fully specified in the customer ask and Phase 1–3 governance docs.
