# Historical resources

This document holds **historically important but inactive** resources that no longer belong in [README.md](../README.md).

README lists resources that help readers understand agent-factory work today: maintained projects, durable references, and academically relevant sources. When a resource was influential but is now **archived**, **unmaintained**, or **superseded**, maintainers remove it from README and relocate it here so readers and future reviewers retain traceability without implying the project is still active.

Manual relocation is the process today. Phase 4 automation may later validate entry format and alphabetization in this file; Phase 7 content seeding may populate entries from maintainer decisions. Neither phase is implemented yet.

## Related documents

| Document | Use when relocating |
|----------|---------------------|
| [docs/review-policy.md](review-policy.md) | Removal triggers, `historical` label, and PR review checklist |
| [MAINTAINERS.md](../MAINTAINERS.md) | Merge expectations, stale-PR handling, and maintainer responsibilities |
| [README.md](../README.md) | Canonical active list and section headings |
| [CONTRIBUTING.md](../CONTRIBUTING.md) | Entry format, tone, and alphabetization rules |
| [docs/taxonomy.md](taxonomy.md) | Original README section when recording where the entry lived |
| [docs/rejected.md](rejected.md) | Commonly submitted out-of-scope types (not historical archives) |

## When to move here versus other outcomes

Use the decision table below after applying the removal triggers in [docs/review-policy.md](review-policy.md#when-to-remove-or-move).

| Situation | Destination | Rationale |
|-----------|-------------|-----------|
| **Historically important but inactive** | This file (`docs/historical.md`) | The resource shaped agent-factory thinking, appears in citations or case studies, or remains a useful reference even though it is archived or unmaintained. |
| **Archived and not historically important** | Remove outright | No durable value for readers; keeping a trace would clutter the list. |
| **Out of scope or commonly rejected type** | [docs/rejected.md](rejected.md) | The submission never fit README scope (for example, single-agent chatbots or generic LLM SDKs), not merely inactive. |
| **Duplicate, spam, unsafe, or malicious** | Remove outright | Follow [SECURITY.md](../SECURITY.md) for urgent cases; do not archive spam or compromised projects here. |
| **Dead link, not historically important** | Remove outright | Apply the `broken-link` label; relocate only when historical importance justifies keeping a citation with an updated or archived URL. |
| **Still maintained and in scope** | Keep in [README.md](../README.md) | Active or academically relevant resources stay on the main list. |

**Historical importance** means the resource materially advanced multi-agent coordination, orchestration, delegation, routing, handoffs, shared state, or group-level evaluation—or it is widely cited as a foundation readers still encounter in papers, blog posts, or frameworks. Mere age or nostalgia is not sufficient.

## Entry format

Each relocated entry uses the same awesome-list line format as [README.md](../README.md) and [CONTRIBUTING.md](../CONTRIBUTING.md):

```markdown
- [Resource Name](URL) - Description.
```

Requirements:

- **Link text** is the official resource name (project, paper, or article title).
- **URL** is the best available canonical or archived destination (official homepage, paper page, or permanent archive link when the primary site is gone).
- **Description** is concise, factual, and non-promotional. State what the resource was and why it mattered for agent-factory work. End with a period.
- **Former section** (optional but recommended): note in the pull request or commit message which README section the entry came from; when helpful, add a short HTML comment above the entry in this file: `<!-- Former README section: Frameworks -->`.

Do not invent placeholder entries. Add lines only when maintainers relocate real resources from README or accept a pull request that documents a historical move.

## Alphabetization

- Order entries **alphabetically by link text** (the text inside `[...]`), using case-insensitive comparison.
- Maintain a single list under [Historical entries](#historical-entries) unless the file later grows large enough to warrant subsection splits; if split by former README category, alphabetize **within each subsection**.
- When inserting an entry, change only the lines needed to preserve order—no unrelated reordering.

## Maintainer workflow

When moving an entry from README to this file:

1. **Confirm the trigger** using [docs/review-policy.md](review-policy.md) (inactive or archived, historically important, still worth citing).
2. **Remove** the line from the correct README section.
3. **Add** the entry here in alphabetical order, adjusting the URL or description if needed to reflect inactive status factually (for example, "archived framework for multi-agent orchestration").
4. **Open or update a pull request** that includes both the README removal and the `historical.md` addition. One resource per pull request when the change originates from a contributor submission; maintainer batch relocations may group related housekeeping when [MAINTAINERS.md](../MAINTAINERS.md) allows.
5. **Note the reason** in the pull request description: archived date, superseded by what, link-health context, and former README section.
6. **Apply the `historical` outcome label** and reference this file in the closing comment.
7. **Follow [MAINTAINERS.md](../MAINTAINERS.md)** for approval and merge when the change is maintainer-initiated.

For link-dead resources that remain historically important, prefer an archived or alternate canonical URL in the entry. If no durable URL exists, discuss in the pull request before keeping a broken link here; [docs/review-policy.md](review-policy.md) treats persistent broken links as removal triggers.

## Historical entries

*No entries yet. Maintainers add resources here when relocating them from README per the workflow above.*

## Future automation (Phase 4) and content seeding (Phase 7)

Entry format, alphabetization, and relocation criteria in this document are intended inputs for automated checks and for Phase 7 seeding of historical citations. Implementation of those phases is not part of this repository yet.
