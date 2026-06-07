# PRD: Phase 6–7 Documentation Convergence Repair

## Introduction

Finish the repo-facing and planner-facing documentation reconciliation that the prior `phase-6-7-docs-convergence-cleanup` batch was meant to land. Phase 6 community intake templates and Phase 4–5 automation are present on `main`, but key documentation still drifts: `README.md` presents **Community** as a curated list category in the table of contents, `docs/review-policy.md` still claims automation does not exist, and `factory/docs/overview.md` still describes the old AI model reference website. This repair batch converges those documents so Phase 7 content seeding starts from one accurate source of truth.

## Context

### Customer ask

Repair documentation convergence before Phase 7 content seeding. Update `README.md`, `docs/review-policy.md`, and `factory/docs/overview.md`; keep them aligned with `CONTRIBUTING.md` and `docs/taxonomy.md` on curated categories and contributor/planner workflow. Leave `make check`, `go test ./...`, and `git diff --check` clean. Do not start Phase 7 content seeding in this batch.

### Concrete problem

A prior cleanup idea completed in the factory queue, but filesystem review shows the intended documentation changes never landed on `main`:

- **README navigation drift** — The `## Contents` list ends with `[Community](#community)` instead of reflecting the ten curated resource categories defined in `CONTRIBUTING.md` and `docs/taxonomy.md`. `Related Lists` is a curated category (with its `## Related Lists` section) but is correctly omitted from Contents per awesome-list conventions enforced by `internal/checks`. `Community` is a governance footer (conduct and security pointers), not a curated list section, and must not appear in Contents as if it were.
- **Review-policy automation drift** — `docs/review-policy.md` opens with “Automation does **not** run today” and references Phase 4 as future work, even though `make check`, `make test`, `make links`, and four GitHub workflows (`ci.yml`, `link-check.yml`, `awesome-lint.yml`, `scheduled-maintenance.yml`) are live on `main`.
- **Factory overview drift** — `factory/docs/overview.md` describes an “AI model reference website,” points to nonexistent `docs/documentation-site-pages-needed.md`, and omits the actual Awesome List customer ask and planner-owned state files for this repository.

### High-level solution

Apply a narrow, documentation-only convergence pass: fix README Contents/footer semantics, rewrite review-policy automation language to match the live toolchain while preserving manual ownership of scope and quality judgments, repair factory planner guidance for this repo, and run a final cross-document consistency check across the five named policy files before the Phase 7 content gate.

## Goals

- Present one coherent story for contributors: ten curated README categories, local `make` checks, GitHub workflow gates, and maintainer manual review responsibilities.
- Present one coherent story for factory planners: this repository builds an Awesome List; phase control lives in `docs/internal/customer-ask.md`; progress is logged in planner-owned state files.
- Remove stale references to pre-migration project scope (AI model reference site, missing documentation-site roadmap file).
- Keep scope narrow: documentation convergence and planner guidance only—no Phase 7 list entries, no unrelated governance rewrites.

## Project-level acceptance criteria

- [ ] `README.md` `## Contents` lists `Scope` plus the nine curated resource sections that belong in the table of contents (Theories through Examples and Templates), omits `Related Lists` and `Contributing` per awesome-list Contents conventions, and does **not** list `Community`.
- [ ] `README.md` retains `## Related Lists` as the tenth curated category section and `## Community` as a non-curated governance footer after `## Contributing`.
- [ ] `docs/review-policy.md` accurately describes live automation (`make check`, `make test`, `make links`, GitHub CI/link-check/awesome-lint/scheduled-maintenance workflows) and states that manual review still owns scope fit, quality, convergence, and borderline judgments automation cannot make.
- [ ] `factory/docs/overview.md` describes the **Awesome AI Agent Factories** repository, references planner-owned `docs/internal/customer-ask.md`, `docs/internal/checklist.md`, and `docs/internal/progress.txt`, and removes stale references to the old AI model reference website and `docs/documentation-site-pages-needed.md`.
- [ ] `README.md`, `CONTRIBUTING.md`, `docs/review-policy.md`, `docs/taxonomy.md`, and `factory/docs/overview.md` agree on the ten curated categories and the contributor/planner workflow.
- [ ] Quality gate: `make check`, `go test ./...`, and `git diff --check` all exit 0 from the repository root.

## User Stories

### US-001: Converge README Contents with curated categories

**Description:** As a contributor reading the list, I want the README table of contents to reflect the ten curated resource categories so I know where to place submissions without treating governance links as list content.

**Acceptance Criteria:**

- [ ] `## Contents` includes anchor links for `Scope` and for Theories, Coordination Patterns, Frameworks, Protocols and Interfaces, Benchmarks, Research Papers, Blog Posts, Case Studies, and Examples and Templates—in that canonical order.
- [ ] `## Contents` does **not** include `Related Lists`, `Contributing`, or `Community` (matching awesome-list Contents conventions enforced by `make check`).
- [ ] `## Related Lists` remains present as the tenth curated category section with its existing intro text.
- [ ] `## Contributing` and `## Community` remain after the curated sections; `## Community` covers conduct and security expectations only, not list entries.
- [ ] `make check` exits 0 after the README update.
- [ ] Typecheck passes

### US-002: Update review policy for live automation with manual judgment ownership

**Description:** As a maintainer, I want `docs/review-policy.md` to describe what automation enforces today and what still requires human judgment so reviewers and contributors share accurate expectations.

**Acceptance Criteria:**

- [ ] The introduction no longer states that automation does not run or that Phase 4 tooling is absent from the repository.
- [ ] A contributor-readable section documents live local automation: `make check` (Go README structure and entry rules), `make test` / `go test ./...`, and optional `make links` / `make lint`, with pointers to `CONTRIBUTING.md` for the GitHub Actions mapping table.
- [ ] A contributor-readable section documents live GitHub workflow gates: `ci.yml`, `link-check.yml`, `awesome-lint.yml`, and `scheduled-maintenance.yml`, described as read-only enforcement aligned with local commands.
- [ ] Checklist items that claim automated link checking is unavailable are updated to reflect `make links`, `link-check.yml`, and scheduled link checks.
- [ ] The document explicitly states manual review still owns scope fit, section-fit disputes, agent-factory relevance, quality bar, removal/relocation decisions, and convergence judgments automation cannot make.
- [ ] The former “Future automation (Phase 4)” section is reframed: Phase 4 structural checks and Phase 5 workflow gates are present; only not-yet-implemented enhancements (for example scope-keyword warnings) remain future work.
- [ ] Typecheck passes

### US-003: Repair factory planner overview for this repository

**Description:** As a meta-planner operating the agent factory, I want `factory/docs/overview.md` to describe this Awesome List repository and its planner state files so batch submission guidance matches reality.

**Acceptance Criteria:**

- [ ] The overview identifies the project as **Awesome AI Agent Factories**—a curated awesome list for agent-factory coordination, orchestration, and flows—not an AI model reference website.
- [ ] The “Read First” list includes `factory/factory.json`, `factory/workstations/ideafy/AGENTS.md`, `docs/internal/customer-ask.md`, `docs/internal/checklist.md`, and `docs/internal/progress.txt`.
- [ ] The “Read First” list does **not** reference `docs/documentation-site-pages-needed.md` or other paths that do not exist in this repository.
- [ ] Phase control text points to `docs/internal/customer-ask.md` as the source for current phase authorization and the Awesome List build goal.
- [ ] Existing factory workflow sections (work types, workstation flow, batch submission, state inspection, repair) remain accurate and refer to `you docs agents` / `you docs batch-inputs` and `factory/docs/batch-input-example.json`.
- [ ] Typecheck passes

### US-004: Align taxonomy and cross-document category workflow language

**Description:** As a maintainer preparing Phase 7 seeding, I want taxonomy and sibling policy docs to agree with README and CONTRIBUTING on categories and workflow so planners and contributors read the same rules.

**Acceptance Criteria:**

- [ ] `docs/taxonomy.md` lists the same ten curated README categories as `CONTRIBUTING.md` and does not describe Phase 4 automation as unimplemented when referring to current repository state.
- [ ] `docs/taxonomy.md` “Future automation” language distinguishes implemented structural checks (Phase 4) and workflow gates (Phase 5) from not-yet-started Phase 7 content seeding.
- [ ] `CONTRIBUTING.md` category list, README section headings, and `docs/review-policy.md` checklist section list name the same ten categories with matching headings.
- [ ] `factory/docs/overview.md` does not contradict contributor category or local-check guidance in `CONTRIBUTING.md`.
- [ ] No new README resource entries, issue-template edits, or unrelated governance file rewrites are introduced in this story.
- [ ] Typecheck passes

### US-005: Verify documentation convergence quality gate

**Description:** As a planner running convergence review before Phase 7, I want end-to-end verification that the repaired documentation passes repository quality gates and contains no whitespace regressions.

**Acceptance Criteria:**

- [ ] From repository root, `make check` exits 0 (README structure, Contents alignment, entry rules).
- [ ] From repository root, `go test ./...` exits 0.
- [ ] From repository root, `git diff --check` exits 0 (no conflict markers or whitespace errors in changed files).
- [ ] A reviewer can read `README.md` Contents, `CONTRIBUTING.md` categories, `docs/taxonomy.md`, `docs/review-policy.md`, and `factory/docs/overview.md` without finding contradictory claims about whether automation runs or what the ten curated categories are.
- [ ] No Phase 7 list entries are added.
- [ ] Typecheck passes
- [ ] Tests pass

## Functional Requirements

- FR-1: README Contents lists Scope plus nine in-TOC curated sections; Related Lists, Contributing, and Community are excluded from Contents.
- FR-2: README retains all ten curated `##` section headings including Related Lists.
- FR-3: Review policy documents both automated enforcement surfaces and manual review ownership.
- FR-4: Factory overview describes Awesome List scope, planner state files, and current factory workflow without stale external-project references.
- FR-5: Taxonomy, CONTRIBUTING, review policy, and README use identical category names and headings.
- FR-6: Repository quality commands (`make check`, `go test ./...`, `git diff --check`) pass after changes.

## Non-Goals

- Seeding Phase 7 README content (theories, frameworks, papers, and other entries).
- Rewriting `MAINTAINERS.md`, `SECURITY.md`, or GitHub issue/PR templates unless a direct contradiction with repaired docs is discovered during the cross-doc pass.
- Adding new automated checks or changing `internal/checks` behavior beyond what README edits require to keep `make check` passing.
- Bulk link-check or awesome-lint remediation across external URLs.
- Creating `docs/documentation-site-pages-needed.md` or other documentation-site scaffolding.

## High-level Technical Design

This batch is documentation-only. No application runtime, API, or UI changes are involved.

**Document ownership map**

| Concern | Primary file | Consumers |
| --- | --- | --- |
| Curated category names and TOC | `README.md` | Contributors, awesome-lint, `internal/checks` |
| Submission rules and local/CI commands | `CONTRIBUTING.md` | Contributors, review policy |
| Category definitions | `docs/taxonomy.md` | Maintainers, review policy |
| Review checklist and automation split | `docs/review-policy.md` | Maintainers, contributors (self-check) |
| Factory planner guidance | `factory/docs/overview.md` | Ideafy meta-planner, batch executors |

**Contents conventions (behavioral contract)**

`internal/checks` enforces awesome-list Contents rules: required resource section headings must exist; Contents must link to each in-TOC curated section; `Related Lists` and `Contributing` headings must exist but are omitted from Contents. `Community` is not a curated category and must not appear in Contents.

**Automation vs manual review split**

| Automated today | Manual maintainer judgment |
| --- | --- |
| README section presence and Contents alignment | Scope and agent-factory relevance |
| Entry format, punctuation, banned phrases | Section-fit disputes and borderline resources |
| Duplicate URL/name detection | Quality bar and historical importance |
| Alphabetization within sections | Removal, relocation, and convergence decisions |
| Go format/tests in CI | Whether a submission should merge |
| Link checks (local `make links`, scheduled workflow) | Canonical URL and durability judgments |
| awesome-lint on PR/push | Whether marketing tone is acceptable |

**Implementation order**

1. README Contents/footer (US-001) — establishes canonical navigation.
2. Review policy automation language (US-002) — independent but references CONTRIBUTING CI table.
3. Factory overview repair (US-003) — independent planner doc.
4. Taxonomy and cross-doc alignment (US-004) — depends on prior doc content settling.
5. Quality gate verification (US-005) — final convergence proof.

## Supporting Technical and UX Considerations

- Prefer minimal diffs: update prose in place rather than restructuring unrelated sections.
- When reframing “future automation” sections, preserve maintainer intent (checklist questions and labels) while correcting factual claims about what runs today.
- `factory/docs/overview.md` is bundled factory documentation; keep workstation flow diagrams and `you` CLI commands accurate.
- Contributors discover categories through README Contents and CONTRIBUTING; maintainers use taxonomy and review policy—wording should be consistent but not duplicated verbatim across every file.
- Do not introduce trailing whitespace; `git diff --check` is part of the quality gate.

## Success Metrics

- Zero contradictory “automation does not run” statements across the five named policy files after repair.
- README Contents matches CONTRIBUTING’s ten-category model without listing Community as curated content.
- Factory overview references only files that exist on `main`.
- `make check` and `go test ./...` pass on first attempt after merge.
- Phase 7 content seeding can begin without a follow-up documentation repair batch.

## Open Questions

None. Scope, target files, and quality gate are defined by the customer ask and current repository state.
