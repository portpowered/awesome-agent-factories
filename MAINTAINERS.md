# Maintainers

This document describes who maintains **Awesome AI Agent Factories**, how pull requests are reviewed and merged, and when list entries may be removed or stalled work is closed.

Review is **manual**. There is no automated merge bot, scheduled link checker, or dedicated moderation staff unless such workflows are added to the repository and documented here.

## Current maintainers

| Maintainer | GitHub | Role |
|------------|--------|------|
| Repository owner | [@portpowered](https://github.com/portpowered) | List stewardship, scope decisions, security and conduct reports |
| Collaborator | [@AndreasAbdi](https://github.com/AndreasAbdi) | Pull request review, scope and formatting checks |

To reach maintainers, open a GitHub issue, comment on the relevant pull request, or send a GitHub message to one of the handles above. For sensitive security reports, use the channels in [SECURITY.md](SECURITY.md).

## Review responsibilities

Maintainers review pull requests and issues in good faith and as availability allows. There is no guaranteed response-time SLA.

Apply the full [docs/review-policy.md](docs/review-policy.md) PR review checklist for each resource submission. The checklist operationalizes the summary below; use [docs/taxonomy.md](docs/taxonomy.md) for section-fit disputes.

For each resource submission, maintainers check:

- **Scope fit** — The resource belongs on an agent-factory list: it helps manage groups of agents or their flows (coordination, orchestration, delegation, routing, handoffs, shared state, or group-level evaluation). See [CONTRIBUTING.md](CONTRIBUTING.md) for Included, Excluded, and category rules.
- **Category** — The entry is placed in the correct README section, or a new category meets the minimum entry count and justification. See [docs/taxonomy.md](docs/taxonomy.md) when category fit is disputed.
- **Agent-factory relevance** — The pull request description explains how the resource relates to multi-agent or flow management.
- **Formatting** — Entry uses `- [Resource Name](URL) - Description.` with the exact resource name as link text and a concise factual description ending with a period.
- **Tone** — Descriptions avoid marketing language and stay encyclopedic.
- **Duplicates** — The URL does not already appear anywhere in [README.md](README.md).
- **Alphabetization** — The entry is in alphabetical order by link text within its section.
- **Link stability** — The URL appears to be the canonical, official destination for the resource at review time.
- **Conduct** — The submission and discussion follow [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

When declining out-of-scope submissions, cite [docs/rejected.md](docs/rejected.md) for commonly rejected types. When relocating historically important but inactive resources, follow [docs/historical.md](docs/historical.md) and apply the `historical` label per [docs/review-policy.md](docs/review-policy.md).

Maintainers may request edits, reject out-of-scope submissions, or ask for more evidence before merge. They are not required to research every linked project beyond what is needed for a reasonable list-quality decision.

## Merge policy

- **Maintainer approval required** — Resource pull requests are merged only after a maintainer approves the change.
- **One resource per pull request** — Bundled unrelated entries are not merged; contributors should split work per [CONTRIBUTING.md](CONTRIBUTING.md).
- **Feedback first** — Authors should address maintainer review comments or explain why a requested change is unnecessary before merge.
- **No silent scope exceptions** — Borderline resources need an explicit agent-factory justification in the pull request; maintainers should not merge entries that only fit single-agent or generic SDK use cases without a clear orchestration angle.
- **Maintainer-authored changes** — When a maintainer submits a substantive list change, another maintainer should review when possible. Small typo fixes may be merged without a second review.

Merges are performed manually on GitHub. There is no required CI gate, awesome-lint job, or Go check in this repository yet.

## Stale pull request policy

Pull requests may stall when authors do not respond to feedback or when list conventions change before merge.

- After **60 days** with no author or maintainer activity, a maintainer may comment asking whether the contributor still intends to finish the pull request.
- If there is no response within **14 days** of that comment, or the change is clearly outdated or superseded, maintainers may **close** the pull request without merge.
- Closing a stale pull request is not a rejection of the resource. Contributors may reopen an updated branch or open a fresh pull request that meets current [CONTRIBUTING.md](CONTRIBUTING.md) rules.

Draft pull requests and pull requests blocked on an explicit maintainer question follow the same inactivity window from the last meaningful activity.

## Removal and relocation policy

Maintainers may remove or relocate README entries when keeping them would mislead readers or violate list scope. Detailed triggers, labels, and workflow are in [docs/review-policy.md](docs/review-policy.md#removal-and-relocation).

Common reasons for removal:

- **Dead or unstable links** — The destination is persistently broken, redirects to unrelated content, or has been unreachable for an extended period (for example, more than 30 days) without a suitable replacement URL.
- **Archived or abandoned projects** — The resource is no longer maintained in a way that makes the entry misleading, especially when combined with security or accuracy concerns.
- **Compromised or malicious resources** — Credible reports or evidence of account takeover, malware, typosquatting, or harmful install paths. See [SECURITY.md](SECURITY.md).
- **Scope drift** — The linked resource no longer addresses agent-factory topics, or never did despite the description.
- **Misleading descriptions** — Marketing language, inaccurate capabilities, or descriptions that do not match the linked resource.
- **Duplicates** — The same URL or an equivalent resource already appears elsewhere in the list.
- **Spam or bad-faith submissions** — Promotional flooding, irrelevant links, or repeated violations after maintainer feedback.

When removal is non-urgent, maintainers prefer to comment on the relevant issue or pull request with the reason. Urgent security concerns may result in immediate removal with a brief follow-up note.

**Relocation outcomes:**

- **Historically important but inactive** — Remove from [README.md](README.md) and add to [docs/historical.md](docs/historical.md) with the `historical` label. See [docs/review-policy.md](docs/review-policy.md).
- **Commonly submitted out-of-scope types** — Decline with reference to [docs/rejected.md](docs/rejected.md) and the `rejected` label when appropriate.
- **Unsafe or misleading entries** — Remove outright rather than leaving them in place.

## Security and contact policy

Security concerns about linked resources, typosquatting, supply-chain risk, or compromised projects are handled under [SECURITY.md](SECURITY.md).

Use these channels:

1. **GitHub Security Advisories** — [Private vulnerability report](https://github.com/portpowered/awesome-agent-factories/security/advisories/new) for sensitive issues.
2. **GitHub issue** — Public reports for suspicious links, typosquatting, or outdated entries.
3. **Maintainer contact** — [@portpowered](https://github.com/portpowered) or [@AndreasAbdi](https://github.com/AndreasAbdi) on GitHub.

Conduct reports are covered in [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md). Maintainers route security and conduct reports to an available maintainer; there is no separate security operations team.

## What maintainers do not do today

Unless this document is updated after new automation lands:

- We do not run continuous link monitoring across all README entries.
- We do not enforce formatting or scope with repository CI, Go checks, or awesome-lint.
- We do not guarantee 24/7 review or incident response.

Community reports and pull request review are the primary quality controls until Phase 4 automation is added. Review docs under `docs/` describe manual expectations today and are intended inputs for later automated checks.

## Related policies

- [CONTRIBUTING.md](CONTRIBUTING.md) — submission format, scope, categories, and agent-factory relevance
- [docs/review-policy.md](docs/review-policy.md) — PR review checklist, removal rules, and recommended labels
- [docs/taxonomy.md](docs/taxonomy.md) — category definitions for section-fit decisions
- [docs/historical.md](docs/historical.md) — historically important but inactive resources
- [docs/rejected.md](docs/rejected.md) — commonly submitted out-of-scope types
- [SECURITY.md](SECURITY.md) — threats, reporting, and response for linked resources
- [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) — behavior expectations in issues and pull requests
- [LICENSE](LICENSE) — license for curated list content
