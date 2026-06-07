# PRD: Phase 7 Source Quality Audit

## Introduction

Phase 7 README section expansions are complete: all ten curated resource sections on `main` now carry seeded entries. This batch is **convergence cleanup**—a maintainer-defensible quality audit of `README.md` and contributor-facing docs so the Phase 7 manual gate on **stable URLs** and **non-promotional sourcing** can pass review.

The concrete change: systematically review every README resource entry and aligned contributor prose for source hygiene (canonical URLs, scope fit, encyclopedic tone), remediate redirecting or fragile links, remove or replace weak or marketing-heavy entries that do not meet documented scope, and apply only narrow follow-on edits to contributor docs where audited outcomes must be reflected. **Preserve** the existing section taxonomy, Contents navigation, and README governance structure (Scope, Contributing, Community).

## Context

### Customer ask

Phase 7 convergence cleanup: after the README section expansions are complete, audit `README.md` and the contributor-facing docs for content quality and source hygiene. Resolve redirecting URLs where a more canonical destination is available, remove or replace any weak or marketing-heavy entries that do not meet the documented scope, and make any narrow follow-on edits needed so the Phase 7 manual gate about stable URLs and non-promotional sourcing is defensible. Preserve the existing section taxonomy and README structure.

### Problem

Bulk Phase 7 seeding prioritized coverage density across ten curated sections. Some entries may still use redirecting URLs, campaign or version-fragile deep links, promotional description tone, or borderline scope fit that automated checks (`make check`, `make links`) cannot judge. Without a focused audit pass, maintainers cannot confidently sign off the Phase 7 manual gate in [docs/review-policy.md](../../docs/review-policy.md) (checklist items 6 and 8) and [CONTRIBUTING.md](../../CONTRIBUTING.md) (canonical URLs, marketing prohibitions).

### Solution

Run a section-preserving quality audit: canonicalize URLs, remove or replace weak entries per [docs/taxonomy.md](../../docs/taxonomy.md) include/exclude rules and [docs/review-policy.md](../../docs/review-policy.md) removal guidance, rewrite remaining descriptions to encyclopedic factual tone with explicit agent-factory relevance, and make narrow contributor-doc updates only where the audited state must be documented. Verify with `make check`, `make test`, and `make links`.

## Goals

- Every remaining README resource URL points to a canonical, durable official destination where one exists.
- No README entry survives the audit with prohibited marketing language or weak sourcing that fails documented scope.
- Weak or out-of-scope entries are removed, replaced with stronger alternatives, or relocated to [docs/historical.md](../../docs/historical.md) when historically important.
- Contributor-facing docs stay aligned with the audited README without restructuring taxonomy or section headings.
- Phase 7 manual gate on stable URLs and non-promotional sourcing is reviewer-verifiable from the resulting diff and passing checks.

## Project-Level Acceptance Criteria

- [ ] Every README resource entry in all ten curated sections uses a canonical stable URL (no unnecessary redirect hops, shortened links, or version-fragile deep links when an equally canonical root or official page exists).
- [ ] No README entry retains prohibited marketing language or hype tone per [CONTRIBUTING.md](../../CONTRIBUTING.md); descriptions are concise, factual, encyclopedic, and end with a period.
- [ ] No README entry remains that fails [docs/taxonomy.md](../../docs/taxonomy.md) include rules or [README.md](../../README.md) scope for its section; weak entries are removed, replaced, or relocated per [docs/review-policy.md](../../docs/review-policy.md).
- [ ] README section taxonomy, Contents navigation anchors, and governance sections (Scope, Contributing, Community) are preserved; entries stay alphabetized by link text within each section.
- [ ] Contributor-facing docs ([CONTRIBUTING.md](../../CONTRIBUTING.md), [docs/taxonomy.md](../../docs/taxonomy.md), [docs/review-policy.md](../../docs/review-policy.md)) receive only narrow follow-on edits needed to reflect the audited quality bar—no category renames or structural rewrites.
- [ ] `make links` reports no failures for `README.md` and `docs/*.md` after remediation.
- [ ] Quality gate: `make check`, `make test`, and `git diff --check` all pass from the repository root.

## User Stories

### US-001: Canonicalize README resource URLs

**Description:** As a maintainer signing off Phase 7, I want every README entry to link to canonical stable destinations so readers and link checks do not depend on redirect chains or fragile deep links.

**Acceptance Criteria:**

- [ ] Each resource URL in all ten curated README sections resolves to the official homepage, paper page, repository root, or other durable canonical destination for that resource—not a redirect hop, shortener, tracking URL, or version-fragile path when a more canonical target exists.
- [ ] URL changes preserve exact entry format `- [Resource Name](URL) - Description.` and alphabetization by link text within each section.
- [ ] No duplicate URLs are introduced anywhere in `README.md`.
- [ ] `make links` exits 0 for `README.md` after URL remediation.
- [ ] `make check` exits 0 after URL remediation.
- [ ] Typecheck passes
- [ ] Tests pass

### US-002: Remove or replace weak and out-of-scope README entries

**Description:** As a reader trusting this curated list, I want weak, borderline, or out-of-scope entries removed or replaced so every remaining row meets the documented agent-factory scope bar.

**Acceptance Criteria:**

- [ ] Every entry removed or replaced is documented in the pull request or story notes with the failing criterion (scope, sourcing quality, duplicate, or unmaintained per [docs/review-policy.md](../../docs/review-policy.md)).
- [ ] Removed entries that are historically important but inactive are relocated to [docs/historical.md](../../docs/historical.md) with factual relocation prose—not left as dead README rows.
- [ ] Replaced entries use stronger canonical sources that clearly fit [docs/taxonomy.md](../../docs/taxonomy.md) include rules for their section and include at least one agent-factory scope keyword (coordination, orchestration, delegation, routing, handoffs, shared state, or group-level evaluation).
- [ ] Affected sections remain alphabetized by link text; no section headings or Contents anchors are renamed or reordered.
- [ ] `make check` exits 0 after entry remediation.
- [ ] Typecheck passes
- [ ] Tests pass

### US-003: Harden README descriptions to encyclopedic non-promotional tone

**Description:** As a contributor learning submission standards, I want every on-main README description to model factual encyclopedic tone so the Phase 7 non-promotional sourcing gate is defensible.

**Acceptance Criteria:**

- [ ] No README resource description contains prohibited marketing language from [CONTRIBUTING.md](../../CONTRIBUTING.md) (for example "revolutionary", "game-changing", "best", "ultimate", or calls to action).
- [ ] Every description states what the resource is and why it matters for agent-factory work in one or two factual sentences ending with a period.
- [ ] Descriptions for entries touched in US-001 or US-002 are reviewed; remaining seeded entries are scanned and rewritten only where tone fails the encyclopedic bar.
- [ ] `make check` exits 0 after description hardening.
- [ ] Typecheck passes
- [ ] Tests pass

### US-004: Converge contributor-facing docs with audited quality outcomes

**Description:** As a maintainer running the Phase 7 manual gate, I want contributor docs to reflect that seeded content passed URL stability and non-promotional sourcing review without restructuring taxonomy.

**Acceptance Criteria:**

- [ ] [docs/taxonomy.md](../../docs/taxonomy.md) Phase 7 status prose records that post-seeding source-quality audit is complete (or in progress only if earlier stories are not yet merged) without renaming categories or altering the ten-section taxonomy table.
- [ ] [docs/review-policy.md](../../docs/review-policy.md) and [CONTRIBUTING.md](../../CONTRIBUTING.md) require no contradictory guidance relative to the audited README entries; apply only narrow clarifications if audit findings exposed a doc gap (for example canonical-URL examples or relocation cross-references).
- [ ] No unrelated governance rewrites, new categories, or README structural changes are introduced in contributor docs.
- [ ] `make links` exits 0 for `docs/*.md` after any doc edits.
- [ ] Typecheck passes
- [ ] Tests pass

### US-005: Verify Phase 7 source-quality audit quality gates

**Description:** As a maintainer merging this batch, I want end-to-end verification that the audited README and docs pass automated checks and preserve list structure.

**Acceptance Criteria:**

- [ ] From repository root, `make check` exits 0.
- [ ] From repository root, `make test` exits 0.
- [ ] From repository root, `make links` exits 0 for `README.md` and `docs/*.md`.
- [ ] `git diff --check` reports no whitespace errors on changed files.
- [ ] README retains ten curated resource sections with unchanged Contents navigation anchors and governance sections (Scope, Contributing, Community).
- [ ] Changed content files are limited to `README.md`, contributor-facing docs (`CONTRIBUTING.md`, `docs/taxonomy.md`, `docs/review-policy.md`, and `docs/historical.md` only when relocations occurred), and planning artifacts for this batch.
- [ ] Typecheck passes
- [ ] Tests pass

## Functional Requirements

- FR-1: Audit all resource entries across the ten curated README sections (Theories through Related Lists) for URL stability per [docs/review-policy.md](../../docs/review-policy.md) checklist item 8.
- FR-2: Upgrade redirecting, shortened, tracking, or version-fragile URLs to canonical official destinations when available.
- FR-3: Remove, replace, or relocate entries that fail scope, sourcing quality, or maintenance bar per [docs/taxonomy.md](../../docs/taxonomy.md) and [docs/review-policy.md](../../docs/review-policy.md) removal triggers.
- FR-4: Rewrite descriptions to encyclopedic factual tone per [CONTRIBUTING.md](../../CONTRIBUTING.md) and checklist item 6 in [docs/review-policy.md](../../docs/review-policy.md).
- FR-5: Preserve README section taxonomy, Contents TOC anchors, entry format, and alphabetization rules.
- FR-6: Apply narrow contributor-doc follow-on edits only where needed to document audited Phase 7 quality outcomes.
- FR-7: Pass `make check`, `make test`, and `make links` before merge.

## Non-Goals

- Adding new README entries to increase section density (handled by sibling Phase 7 expansion batches).
- Renaming README sections, changing Contents navigation, or restructuring the ten-category taxonomy.
- Broad unrelated cleanup (refactors, CI changes, Go checker rule changes, or awesome-lint config edits).
- Populating [docs/rejected.md](../../docs/rejected.md) with new rejection patterns unless a repeated submission type emerged during audit.
- Automated enforcement of banned-phrase or canonical-URL heuristics in `internal/checks` (manual gate only in this batch).

## High-Level Technical Design

This batch is documentation curation, not application code. Work proceeds in dependency order:

1. **URL pass** — For each README entry, fetch or inspect the linked destination. Prefer repository roots over `/tree/main/...` deep paths when the root is equally canonical; prefer official docs roots over blog campaign URLs; prefer paper pages (arXiv, proceedings) over aggregator mirrors. Record before/after URLs in commit messages or story notes.
2. **Entry quality pass** — Evaluate scope fit against [README.md](../../README.md) Scope and section include rules. Remove outright when duplicate, spam, or out of scope; replace when a stronger canonical source exists; relocate to `docs/historical.md` when historically important but inactive.
3. **Tone pass** — Rewrite descriptions for entries that remain, ensuring agent-factory relevance keywords and terminal periods.
4. **Doc convergence** — Update Phase 7 status prose in `docs/taxonomy.md` and only necessary cross-references in `CONTRIBUTING.md` / `docs/review-policy.md`.
5. **Verification** — Run `make check`, `make test`, `make links`, and `git diff --check`.

No new packages, APIs, or UI surfaces are introduced. Existing `internal/checks` validates structure post-edit; lychee validates link health.

## Supporting Technical and UX Considerations

- **Check layers:** `make check` enforces section headings, Contents alignment, entry format, and alphabetization. `make links` (lychee with [`.lychee.toml`](../../.lychee.toml)) validates HTTP reachability. Neither replaces maintainer judgment on canonical URL choice—document rationale when upgrading a link lychee already passes.
- **Relocation protocol:** Follow [docs/review-policy.md](../../docs/review-policy.md) removal table—historically important inactive resources go to `docs/historical.md` with factual prose, not README stubs.
- **Alphabetization:** Any removal, replacement, or link-text change may require re-sorting within the affected section only.
- **Contributor UX:** Remaining entries should model the submission format contributors see in [CONTRIBUTING.md](../../CONTRIBUTING.md)—exact resource name as link text, canonical URL, factual description.

## Success Metrics

- Zero README link failures in `make links` after merge.
- Zero entries flagged in maintainer self-check against review-policy items 6 and 8.
- No net increase in README entry count unless a removed weak entry is replaced one-for-one with a stronger source.
- Phase 7 manual gate can be signed off with a reviewer-verifiable diff and passing `make check` / `make test` / `make links`.

## Open Questions

- None blocking implementation. Specific URLs and entries to change are determined during the audit pass against live destinations and documented scope rules.
