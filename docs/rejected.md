# Rejected resource types

This document records **commonly submitted resource types** that do not fit [README.md](../README.md) scope for **Awesome AI Agent Factories**.

README lists systems, theories, frameworks, and practices for managing **groups** of agents and their **flows**: coordination, orchestration, delegation, routing, handoffs, shared state, and group-level evaluation. Submissions that lack that agent-factory angle are declined. This file gives maintainers a stable, citeable reference when the same out-of-scope patterns recur.

Manual review is the process today. Phase 4 automation may later warn on scope keywords or match against documented rejection patterns; Phase 7 content seeding does not populate this file with README entries. Neither phase is implemented yet.

## Related documents

| Document | Use when declining |
|----------|-------------------|
| [docs/review-policy.md](review-policy.md) | PR review checklist, removal triggers, `rejected` label, and triage labels |
| [docs/taxonomy.md](taxonomy.md) | Category-level include/exclude rules and scope terminology |
| [CONTRIBUTING.md](../CONTRIBUTING.md) | Entry format, scope exclusions, and submission expectations |
| [README.md](../README.md) | Canonical scope statement and section headings |
| [docs/historical.md](historical.md) | Historically important but inactive resources (not out-of-scope types) |
| [MAINTAINERS.md](../MAINTAINERS.md) | Merge expectations, stale-PR handling, and maintainer responsibilities |

## When to document here versus one-off declines

Not every declined pull request needs a new note in this file. Use the decision table below.

| Situation | Action | Rationale |
|-----------|--------|-----------|
| **First-time or unusual out-of-scope submission** | Decline with the [PR review checklist](review-policy.md#pr-review-checklist), apply the `rejected` label, and cite [README scope](../README.md) or [docs/taxonomy.md](taxonomy.md). | A single bad fit does not require updating this reference. |
| **Repeat submission of the same out-of-scope type** | Decline and **cite this file** (section link below) in the closing comment. Apply the `rejected` label. | Contributors and maintainers get a durable explanation without re-debating scope each time. |
| **New recurring pattern not yet listed** | Decline the pull request, then add or extend a category in [Commonly rejected categories](#commonly-rejected-categories) via maintainer pull request when the pattern appears more than once or is obviously common (for example, generic SDK spam). | Keeps the document factual and avoids speculative categories for one-off mistakes. |
| **Inactive but historically important resource** | Do not add here. Route to [docs/historical.md](historical.md) per [review-policy removal rules](review-policy.md#where-relocated-entries-go). | Rejection documents scope mismatch, not archival status. |
| **Duplicate, broken link, or misleading description** | Handle with checklist and outcome labels (`duplicate`, `broken-link`) per [docs/review-policy.md](review-policy.md). Cite this file only when scope is the primary reason. | Operational quality issues are not "commonly rejected types." |
| **Malicious, compromised, or spammy submission** | Reject immediately. Follow [SECURITY.md](../SECURITY.md) for urgent response. Apply `rejected` for spam; do not treat security incidents as taxonomy updates. | Security response takes precedence over scope documentation. |

**One-off declines** use triage and outcome labels from [docs/review-policy.md](review-policy.md) (`needs-scope-review`, `rejected`, and related labels) without editing this file.

**Documented rejection notes** belong here when maintainers expect the same category of submission to appear repeatedly and want a neutral, reusable explanation.

## Commonly rejected categories

The categories below align with [README.md](../README.md) scope exclusions and [CONTRIBUTING.md](../CONTRIBUTING.md). Each states why the type is out of scope for an agent-factory list.

### Single-agent chatbots

**Why rejected:** The resource centers on one conversational agent without managing a **group** of agents, task routing between roles, or multi-step **flows**. A chatbot wrapper, persona bot, or single-assistant product does not help readers design agent factories unless the submission documents explicit multi-agent orchestration (supervision, delegation, or handoffs).

**Agent-factory angle required:** Multiple agents with defined roles, or a documented flow where tasks move between agents or human checkpoints.

**See also:** [docs/taxonomy.md](taxonomy.md) exclude rules under Theories, Case Studies, and Frameworks.

### Generic LLM SDKs

**Why rejected:** The resource is a general-purpose library for calling language models (completion, chat, embeddings) without orchestrating agent groups, workflows, or handoffs. Model access alone is infrastructure, not agent-factory design.

**Agent-factory angle required:** APIs or patterns for coordinating multiple agents, shared state across agents, or flow control—not merely invoking a model.

**See also:** [docs/taxonomy.md](taxonomy.md) Frameworks exclude rules (generic model SDKs).

### Model provider SDKs

**Why rejected:** Official or third-party SDKs for a specific model provider (OpenAI, Anthropic, Google, and similar) focus on authentication, requests, and streaming for single-model use. They are not agent-factory frameworks unless the linked documentation or project explicitly addresses multi-agent orchestration as its primary purpose.

**Agent-factory angle required:** The submission must point to multi-agent coordination features, not the provider's base client library.

### Prompt collections

**Why rejected:** Curated prompts, prompt marketplaces, or prompt-engineering guides without orchestration structure do not manage agent groups or flows. Prompt text alone does not implement delegation, routing, shared workspaces, or group evaluation.

**Agent-factory angle required:** Runnable multi-agent workflows, templates with role topology, or documented handoff patterns—not isolated prompt strings.

**See also:** [docs/taxonomy.md](taxonomy.md) Blog Posts and Examples and Templates exclude rules.

### Standalone RAG tools

**Why rejected:** Retrieval-augmented generation pipelines that index documents and answer queries typically serve a single agent or a single query path. Unless the resource documents multi-agent retrieval, specialist routing, or group evaluation over retrieved context, it falls outside agent-factory scope.

**Agent-factory angle required:** Multiple agents consuming or producing shared retrieval state, router-specialist splits, or benchmarked group RAG workflows.

### Generic observability tools

**Why rejected:** Logging, tracing, metrics, or APM products aimed at any software stack do not belong here unless they document agent-specific group coordination, multi-agent trace analysis, or flow-level evaluation as a first-class feature described in the linked resource.

**Agent-factory angle required:** Observability of **agent teams or flows** (handoffs, role failures, group-level success metrics), not generic application monitoring.

### Non-agent workflow engines

**Why rejected:** General workflow or BPM engines (visual flow builders, cron orchestrators, ETL pipelines) without agent-specific usage manage tasks and dependencies, not groups of autonomous agents with delegation and handoffs. A generic automation tool is excluded unless the submission shows agent-oriented routing, tool use between agents, or multi-agent supervision.

**Agent-factory angle required:** Documented use as an agent factory (agent roles, LLM-driven steps with handoffs, or multi-agent coordination), not arbitrary DAG execution.

**See also:** [docs/taxonomy.md](taxonomy.md) Frameworks include rules (workflow-oriented **agent** frameworks).

### Primarily marketing or launch content

**Why rejected:** Vendor launch posts, product announcements, and promotional landing pages violate [CONTRIBUTING.md](../CONTRIBUTING.md) tone rules. They rarely provide durable technical detail about managing agent groups.

**Agent-factory angle required:** Technical depth comparable to [docs/taxonomy.md](taxonomy.md) Blog Posts include rules (architecture, failure analysis, evaluation writeups).

**Note:** Marketing submissions are usually one-off declines. Add a maintainer note here only when a **category** of promotional submissions recurs (for example, repeated affiliate listicles).

## How maintainers cite this file

When closing a pull request or issue for an out-of-scope submission:

1. **State the outcome briefly** — for example, "Out of scope for Awesome AI Agent Factories."
2. **Name the matching category** from [Commonly rejected categories](#commonly-rejected-categories) with a link to the anchor (for example, `docs/rejected.md#generic-llm-sdks`).
3. **Point to scope sources** — [README.md](../README.md) scope section and the relevant [docs/taxonomy.md](taxonomy.md) exclude rule when helpful.
4. **Apply the `rejected` outcome label** per [docs/review-policy.md](review-policy.md#outcome-labels).
5. **Invite a resubmission only when appropriate** — if the contributor can show a genuine multi-agent or flow-management angle, ask them to update the pull request description and re-run the [PR review checklist](review-policy.md#pr-review-checklist) rather than arguing scope in comment threads.

Example closing comment:

```markdown
Thanks for the submission. This resource is a generic LLM client library without multi-agent orchestration, which is out of scope for this list.

See [Generic LLM SDKs](rejected.md#generic-llm-sdks) and [README scope](../README.md). Closing as out of scope; applied label `rejected`.
```

For issues (suggest-a-resource templates), use the same citation pattern so contributors receive a consistent explanation before opening a pull request.

## Maintainer workflow for updating this document

1. **Confirm the pattern is recurring** — at least two similar submissions or obvious community spam patterns.
2. **Draft neutral language** — describe the type and the required agent-factory angle; avoid naming specific vendors unless necessary for clarity.
3. **Cross-check** [docs/taxonomy.md](taxonomy.md) and [CONTRIBUTING.md](../CONTRIBUTING.md) so exclusion language stays aligned.
4. **Open a maintainer pull request** — one category or clarification per change when possible; no bulk rewrites unrelated to observed submissions.
5. **Reference [MAINTAINERS.md](../MAINTAINERS.md)** for review and merge when the edit is maintainer-initiated.

Do not add individual project names as a "wall of shame." This file documents **types**, not a blacklist of contributors.

## Future automation (Phase 4) and maintenance (Phase 9)

Rejection categories and agent-factory rationale in this document are intended inputs for Phase 4 scope warnings and for quarterly maintainer reviews described in the project roadmap. Implementation of automated scope matching is not part of this repository yet.
