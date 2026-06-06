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
| [docs/historical.md](historical.md) | Relocating historically important but inactive resources |
| [docs/rejected.md](rejected.md) | Commonly submitted out-of-scope resource types |
| [SECURITY.md](../SECURITY.md) | Malicious links, typosquatting, supply-chain concerns, and compromised projects |

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

## Removal and relocation

Maintainers remove or relocate README entries when keeping them would mislead readers or violate list scope. This section operationalizes the checklist above for **decline**, **remove**, and **move** outcomes. It aligns with [MAINTAINERS.md](../MAINTAINERS.md) removal and relocation policy and does not override [SECURITY.md](../SECURITY.md) for urgent security response.

### When to remove or move

Apply one or more of the following triggers before removing an entry from [README.md](../README.md) or declining a pull request:

| Trigger | Maintainer action |
|---------|-------------------|
| **Dead link for more than 30 days** | Remove the entry or decline the pull request. Apply the `broken-link` label. Prefer a replacement canonical URL when the resource still exists elsewhere; otherwise remove or relocate per historical importance below. |
| **Archived and not historically important** | Remove the entry. Do not keep abandoned projects in README when they no longer help readers understand agent-factory work. |
| **Misleading description** | Request an edit or remove the entry. Descriptions must match the linked resource and stay encyclopedic per [CONTRIBUTING.md](../CONTRIBUTING.md). |
| **No longer about agent groups or flows** | Remove the entry or decline the pull request. Scope drift away from coordination, orchestration, delegation, routing, handoffs, shared state, or group-level evaluation is out of scope. |
| **Primarily marketing** | Decline or remove. Vendor launch posts, promotional copy, and hype language violate [CONTRIBUTING.md](../CONTRIBUTING.md) tone rules. |
| **Malicious, compromised, or spammy** | Remove immediately or reject before merge. Follow [SECURITY.md](../SECURITY.md) for reporting channels and urgent response. Apply `rejected` for out-of-scope spam; escalate credible compromise reports per SECURITY.md rather than debating scope in the pull request thread alone. |

When removal is non-urgent, comment on the issue or pull request with the reason and applicable label. Urgent security concerns may result in immediate removal with a brief follow-up note, consistent with [MAINTAINERS.md](../MAINTAINERS.md).

### Where relocated entries go

Not every removal is a hard delete. Route outcomes as follows:

| Outcome | Destination | When to use |
|---------|-------------|-------------|
| **Relocate to historical archive** | [docs/historical.md](historical.md) | The resource is **historically important** but **inactive** (archived, unmaintained, or superseded) and still worth citing for readers and future reviewers. Remove from README, add to `historical.md`, and note the reason in the pull request. Apply the `historical` label. |
| **Document as commonly rejected** | [docs/rejected.md](rejected.md) | The submission type is **commonly submitted** but **out of scope** (for example, single-agent chatbots, generic LLM SDKs, prompt collections). Decline the pull request and cite `rejected.md` when the pattern repeats. Apply the `rejected` label. See [docs/taxonomy.md](taxonomy.md) for category-level include/exclude rules. |
| **Remove outright** | *(no list entry)* | The resource is duplicate, spam, unsafe, not historically important, or otherwise should not remain traceable on the list. |

Until [docs/rejected.md](rejected.md) exists, maintainers may still apply the triggers and labels above for commonly rejected types and document the intended relocation in the pull request comment. [docs/historical.md](historical.md) holds durable citations for historically important inactive resources.

### Removal workflow summary

1. Confirm the trigger using the checklist and, for security items, [SECURITY.md](../SECURITY.md).
2. Choose **merge**, **request changes**, **decline**, **remove from README**, or **relocate** to `historical.md` or `rejected.md`.
3. Apply the matching **outcome label** (`duplicate`, `broken-link`, `historical`, or `rejected`) or a **triage label** while the contributor addresses feedback.
4. Cross-reference [MAINTAINERS.md](../MAINTAINERS.md) for stale-pull-request handling when authors do not respond.

## Recommended labels

Apply GitHub labels during triage and when closing pull requests so outcomes stay consistent before Phase 4 automation. Labels fall into two groups: **resource category** (where an accepted entry belongs) and **triage or outcome** (review state or final disposition).

### Resource category labels

Each `resource:*` label maps to a README section. Apply one category label when the section is clear; use `needs-scope-review` when it is not.

| Label | README section | Use when |
|-------|----------------|----------|
| `resource:theory` | Theories | Foundational ideas for organizing agent groups (societies, roles, planning, blackboard systems, swarm intelligence, and similar). |
| `resource:pattern` | Coordination Patterns | Reusable multi-agent system shapes (supervisor-worker, planner-executor, router-specialist, debate, pipeline, hub-and-spoke, and similar). |
| `resource:framework` | Frameworks | Software that manages agent groups, flows, handoffs, or orchestration. |
| `resource:protocol` | Protocols and Interfaces | Standards or conventions for agent interaction, tool interfaces, handoff schemas, and message formats. |
| `resource:benchmark` | Benchmarks | Benchmarks that evaluate group or workflow behavior. |
| `resource:paper` | Research Papers | Academic or technical papers about agent groups and flows. |
| `resource:blog` | Blog Posts | Technical writing on architectures, production lessons, or multi-agent design patterns. |
| `resource:case-study` | Case Studies | Real-world deployments or operations involving multiple agents or agent flows. |
| `resource:example` | Examples and Templates | Runnable or forkable examples (supervisor-worker demos, crew templates, handoff workflows, and similar). |

Category labels help filter open pull requests by intended destination. They do not by themselves approve a submission; maintainers still complete the full checklist.

### Triage labels

Triage labels signal that a pull request needs maintainer or contributor action before merge. Remove or replace them once the issue is resolved.

| Label | Use when |
|-------|----------|
| `needs-scope-review` | Scope or README section fit is unclear. Consult [docs/taxonomy.md](taxonomy.md) and ask the contributor for an agent-factory justification. |
| `needs-link-review` | The URL is broken, redirects suspiciously, or may not be canonical. Verify link health manually until automated link checking exists. |
| `needs-maintainer-review` | A second maintainer opinion is required (borderline scope, category dispute, or maintainer-authored change per [MAINTAINERS.md](../MAINTAINERS.md)). |

### Outcome labels

Outcome labels describe a final or near-final disposition. Prefer citing [docs/review-policy.md](review-policy.md), [docs/historical.md](historical.md), or [docs/rejected.md](rejected.md) in the closing comment when applying these labels.

| Label | Use when |
|-------|----------|
| `duplicate` | The URL or an equivalent resource already appears in [README.md](../README.md). |
| `broken-link` | The destination is persistently unreachable or unsuitable (including dead for more than 30 days). |
| `historical` | The entry is removed from README and relocated to [docs/historical.md](historical.md) as historically important but inactive. |
| `rejected` | The submission is out of scope, spam, or a commonly rejected type documented in [docs/rejected.md](rejected.md). |

Do not treat `resource:*` labels as substitutes for outcome labels. A pull request may carry both during review (for example, `resource:framework` and `needs-scope-review`) until scope is confirmed.

## How maintainers use this checklist

1. Read the pull request description and diff against [CONTRIBUTING.md](../CONTRIBUTING.md).
2. Walk through all nine checklist sections above.
3. Apply **triage labels** (`needs-scope-review`, `needs-link-review`, `needs-maintainer-review`) or a **resource category label** when helpful.
4. For merge decisions, follow [MAINTAINERS.md](../MAINTAINERS.md) (approval required, one resource per pull request, feedback addressed before merge).
5. For section-fit disputes, cite [docs/taxonomy.md](taxonomy.md) when asking the contributor to move or recategorize an entry.
6. When declining, removing, or relocating, follow [Removal and relocation](#removal-and-relocation) and apply the matching **outcome label**.

## Future automation (Phase 4)

This checklist is written so each item can later map to an automated check or warning (format validation, duplicate detection, alphabetization, banned phrases, scope keywords) without changing maintainer intent. Phase 4 implementation is **not** part of this document; no automation is claimed to be running until workflows are merged and documented in the repository.
