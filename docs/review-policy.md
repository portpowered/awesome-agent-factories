# Review policy

This document governs **manual human review** of resource pull requests for **Awesome AI Agent Factories** until Phase 4 automation (Go checks, GitHub Actions, and related tooling) is added to the repository.

Automation does **not** run today. Maintainers apply the checklist below by hand so weak or out-of-scope entries can be rejected consistently before automated enforcement exists. Future automation should adopt the same questions and outcomes described here.

## Related documents

| Document | Use when reviewing |
|----------|-------------------|
| [CONTRIBUTING.md](../CONTRIBUTING.md) | Entry format, scope rules, marketing prohibitions, alphabetization, and one-resource-per-PR discipline |
| [MAINTAINERS.md](../MAINTAINERS.md) | Merge expectations, stale-PR handling, and maintainer responsibilities |
| [docs/taxonomy.md](taxonomy.md) | Section-fit disputes — which README category a resource belongs in |
| [README.md](../README.md) | Canonical scope statement and section headings |

Removal, relocation, and label guidance will be added to this document in a follow-up change. Until then, see [MAINTAINERS.md](../MAINTAINERS.md) for removal and relocation policy.

## PR review checklist

Every resource pull request must be answerable to the questions below. Maintainers may request changes, apply triage labels, or decline the submission when any answer is unsatisfactory.

Use this checklist during review. Contributors should self-check the same items before opening a pull request.

### 1. What is the resource?

- [ ] The pull request identifies the resource by its **official name** (project, paper, or article title).
- [ ] The linked URL is the **canonical destination** for that resource (homepage, paper page, or primary article URL).
- [ ] The entry in [README.md](../README.md) uses the required format: `- [Resource Name](URL) - Description.` as defined in [CONTRIBUTING.md](../CONTRIBUTING.md).

### 2. Which README section does it belong in?

- [ ] The entry is placed in the correct section among: Theories, Coordination Patterns, Frameworks, Protocols and Interfaces, Benchmarks, Research Papers, Blog Posts, Case Studies, Examples and Templates, or Related Lists.
- [ ] If section fit is unclear, consult [docs/taxonomy.md](taxonomy.md) for category definitions, examples, and include/exclude rules before merging or requesting a move.
- [ ] New categories meet [CONTRIBUTING.md](../CONTRIBUTING.md) minimum-entry requirements when applicable.

### 3. Does it manage groups of agents or flows?

- [ ] The resource addresses **agent-factory** work: systems, theories, frameworks, or practices for managing **groups** of agents and their **flows**.
- [ ] Single-agent chatbots, generic LLM SDKs, and other out-of-scope types described in [README.md](../README.md) are excluded unless the pull request explains a direct multi-agent or flow-management angle.

### 4. Does it involve coordination, orchestration, delegation, routing, handoffs, shared state, or group-level evaluation?

- [ ] At least one scope dimension applies: **coordination**, **orchestration**, **delegation**, **routing**, **handoffs**, **shared state**, or **group-level evaluation**.
- [ ] The pull request description explains how the resource helps with one or more of these dimensions (not only that it "uses AI" or "has agents").

### 5. Is it maintained, historically important, or academically relevant?

- [ ] The resource is **actively maintained**, **historically important** to the field, or **academically relevant** (for example, a cited paper or durable reference).
- [ ] Abandoned, archived, or inactive projects are not listed in [README.md](../README.md) unless they remain historically important; relocation options are described in [MAINTAINERS.md](../MAINTAINERS.md).

### 6. Is the description factual and non-promotional?

- [ ] The description is **concise** and **encyclopedic** — it states what the resource is and why it matters for agent-factory work.
- [ ] The description avoids marketing language prohibited in [CONTRIBUTING.md](../CONTRIBUTING.md) (for example, "revolutionary", "game-changing", "best", or "ultimate" unless part of an official title).
- [ ] The description **ends with a period**.

### 7. Is it a duplicate?

- [ ] The URL does not already appear anywhere in [README.md](../README.md).
- [ ] The resource is not a duplicate under a different URL or name already covered by another entry.

### 8. Is the link stable?

- [ ] The URL resolves to the intended resource at review time.
- [ ] The link appears to be an **official** or **durable** destination (not a redirect chain to unrelated content, a temporary campaign page, or a fork that does not represent the resource).
- [ ] Maintainers note link-health concerns even though automated link checking is not yet enabled.

### 9. Is the entry alphabetized?

- [ ] The entry is in **alphabetical order by link text** within its target section, per [CONTRIBUTING.md](../CONTRIBUTING.md).
- [ ] Only the affected section's ordering was changed when adding the new entry (no unrelated reordering).

## How maintainers use this checklist

1. Read the pull request description and diff against [CONTRIBUTING.md](../CONTRIBUTING.md).
2. Walk through all nine checklist sections above.
3. For merge decisions, follow [MAINTAINERS.md](../MAINTAINERS.md) (approval required, one resource per pull request, feedback addressed before merge).
4. For section-fit disputes, cite [docs/taxonomy.md](taxonomy.md) when asking the contributor to move or recategorize an entry.

## Future automation (Phase 4)

This checklist is written so each item can later map to an automated check or warning (format validation, duplicate detection, alphabetization, banned phrases, scope keywords) without changing maintainer intent. Phase 4 implementation is **not** part of this document; no automation is claimed to be running until workflows are merged and documented in the repository.
