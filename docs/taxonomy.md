# Taxonomy

This document defines the README categories for **Awesome AI Agent Factories**. Category names and headings align with [README.md](../README.md) section headings and [CONTRIBUTING.md](../CONTRIBUTING.md) submission rules.

Use these definitions during manual review when deciding which section a resource belongs in. For the full pull-request checklist, see [docs/review-policy.md](review-policy.md). For commonly rejected types, see [docs/rejected.md](rejected.md).

## Related documents

| Document | Use when categorizing |
|----------|----------------------|
| [README.md](../README.md) | Canonical scope statement and section headings |
| [CONTRIBUTING.md](../CONTRIBUTING.md) | Entry format, scope exclusions, and category minimums |
| [docs/review-policy.md](review-policy.md) | PR review checklist and recommended labels |
| [docs/historical.md](historical.md) | Historically important but inactive resources |
| [docs/rejected.md](rejected.md) | Commonly submitted out-of-scope types |

## How to use this taxonomy

1. Confirm the resource fits [README.md](../README.md) scope: systems, theories, frameworks, or practices for managing **groups** of agents and their **flows**.
2. Match the resource to the category whose **include** rules best describe it.
3. Apply **exclude** rules to reject misfiled or out-of-scope submissions.
4. When two categories fit, prefer the category that describes the resource's **primary** contribution (for example, a paper about a coordination pattern belongs in Research Papers, not Coordination Patterns).
5. Apply the matching `resource:*` label from [docs/review-policy.md](review-policy.md) when triaging pull requests.

## Theories

**README section:** [Theories](../README.md#theories)
**Label:** `resource:theory`

Foundational ideas for organizing agent groups: how societies form roles, how tasks decompose, how information is shared, and how groups coordinate without naming a specific product.

### Representative examples

- Agent societies
- Role-based organizations
- Hierarchical planning
- Blackboard systems
- Actor model
- Swarm intelligence
- Market-based coordination
- Human-agent teaming

### Include

- Conceptual models or formalisms for multi-agent organization, cooperation, competition, or planning.
- Academic or industry frameworks that describe **how groups of agents should be structured**, not a single shipped library.
- Historical or enduring ideas that readers apply when designing agent factories.

### Exclude

- Runnable frameworks or SDKs (use **Frameworks**).
- Named coordination topologies described only as implementation patterns (use **Coordination Patterns**).
- Single-agent cognitive architectures with no group or flow angle.

## Coordination Patterns

**README section:** [Coordination Patterns](../README.md#coordination-patterns)
**Label:** `resource:pattern`

Reusable system shapes for multi-agent work: who decides, who executes, how messages flow, and how subtasks combine.

### Representative examples

- Supervisor-worker
- Planner-executor
- Router-specialist
- Critic-reviewer
- Debate
- Swarm
- Pipeline
- Hub-and-spoke
- Shared workspace
- Blackboard
- Committee
- Auction or market-based allocation

### Include

- Documented patterns for task routing, delegation, handoffs, supervision, or group deliberation.
- Architecture writeups whose main value is the **topology** (roles and message flow), whether or not tied to one framework.
- Pattern catalogs or pattern languages aimed at multi-agent system design.

### Exclude

- General software design patterns with no agent or LLM group context.
- Framework-specific tutorials where the framework itself is the primary subject (use **Frameworks** or **Examples and Templates**).
- Benchmarks or papers whose primary contribution is measurement, not pattern description (use **Benchmarks** or **Research Papers**).

## Frameworks

**README section:** [Frameworks](../README.md#frameworks)
**Label:** `resource:framework`

Software that manages agent groups, flows, handoffs, or orchestration.

### Representative examples

- **Multi-agent frameworks** — libraries for defining and running multiple agents with shared runtime concerns.
- **Agent graph frameworks** — graph-structured agent workflows with explicit nodes and edges.
- **Crew or team frameworks** — role-based agent teams with delegation between members.
- **Workflow-oriented agent frameworks** — DAG or state-machine style orchestration over agent steps.
- **Handoff-oriented frameworks** — first-class support for transferring control, context, or tasks between agents.

### Include

- Open-source or commercial frameworks whose core purpose is **multi-agent** coordination, orchestration, routing, or handoffs.
- Tools that provide runtime primitives for agent teams, swarms, supervisors, or shared execution state.
- Frameworks that expose APIs for defining flows across multiple agents, not only a single chat loop.

### Exclude

- **Generic model SDKs** (provider clients without multi-agent orchestration primitives).
- **Single-agent chatbot frameworks** (one conversational agent, no group or flow management).
- **Prompt-only libraries** (templates or prompt packs without orchestration runtime).
- Generic workflow engines with no agent-specific usage unless the pull request documents a direct agent-factory application.

## Protocols and Interfaces

**README section:** [Protocols and Interfaces](../README.md#protocols-and-interfaces)
**Label:** `resource:protocol`

Standards or conventions for how agents, tools, and runtimes interact in multi-agent flows.

### Representative examples

- Agent-to-agent communication protocols
- Tool interfaces shared across agents
- Shared memory or shared workspace interfaces
- Handoff schemas and task-transfer conventions
- Workflow or message schemas for multi-step agent runs
- Task queue interfaces between agents
- Message formats for agent collaboration
- MCP-related resources when relevant to **multi-agent flow management** (not single-tool demos alone)

### Include

- Specifications, schemas, or interface docs that multiple implementations can adopt.
- Interoperability layers for routing, delegation, or shared state across agents.
- Protocol documentation where the primary artifact is the **contract**, not an application built on it.

### Exclude

- Application frameworks that happen to use a protocol internally (use **Frameworks**).
- Single-vendor APIs with no broader agent-interaction story unless they define a reusable handoff or messaging contract.
- Generic HTTP or RPC libraries unrelated to agent coordination.

## Benchmarks

**README section:** [Benchmarks](../README.md#benchmarks)
**Label:** `resource:benchmark`

Benchmarks that evaluate **group** or **workflow** behavior, not isolated single-turn model quality alone.

### Representative examples

- Multi-agent collaboration benchmarks
- Long-horizon task completion suites involving multiple steps or roles
- Tool-use trajectory benchmarks with delegation or handoffs
- Web automation benchmarks driven by agent teams
- Software engineering workflow benchmarks (multi-agent coding or review)
- Debate or negotiation evaluation sets
- Human-agent collaboration benchmarks

### Include

- Datasets, leaderboards, or evaluation harnesses measuring coordination, handoffs, shared planning, or team outcomes.
- Benchmarks where success criteria depend on **multiple agents** or **multi-step flows**.
- Reproducible evaluation methods cited for group-level agent behavior.

### Exclude

- Single-model chat benchmarks with no multi-agent or flow dimension.
- Generic ML benchmarks unrelated to agents or tool use.
- Marketing leaderboard pages without a durable evaluation methodology.

## Research Papers

**README section:** [Research Papers](../README.md#research-papers)
**Label:** `resource:paper`

Academic or technical papers about agent groups and flows.

### Representative examples

- Surveys of multi-agent systems or LLM agent teams
- Papers on multi-agent coordination, communication, or planning
- Work on delegation, routing, or handoffs between agents
- Studies of agent societies, swarms, or organizational structure
- Group-level evaluation, safety, or governance research
- Human-agent teaming in collaborative task settings

### Include

- Peer-reviewed or widely cited preprints whose contribution is **research**, not a product landing page.
- Technical reports with clear methods and results about multi-agent coordination or orchestration.
- Papers that introduce or analyze patterns, benchmarks, or protocols (file here even if the topic overlaps other categories).

### Exclude

- Blog posts and informal writeups (use **Blog Posts**).
- Vendor whitepapers that are primarily promotional (use **Blog Posts** only if technical and non-promotional per [CONTRIBUTING.md](../CONTRIBUTING.md)).
- Executable codebases or frameworks (use **Frameworks** or **Examples and Templates**).

### Optional subthemes

Maintainers may group entries mentally (not necessarily as README subsections) by themes such as surveys, coordination, communication, planning and delegation, agent societies, evaluation, and safety or governance.

## Blog Posts

**README section:** [Blog Posts](../README.md#blog-posts)
**Label:** `resource:blog`

Technical writing that explains real patterns, architectures, failures, or lessons from multi-agent systems.

### Representative examples

- Architecture breakdowns of production agent teams
- Postmortems or failure analyses of multi-agent deployments
- Evaluation writeups for group workflows
- Multi-agent design pattern articles with concrete system detail
- Engineering notes on routing, handoffs, or shared state in live systems

### Include

- Posts that teach **how** groups of agents were designed, operated, or measured.
- Vendor or independent writing that stays factual and technical per [CONTRIBUTING.md](../CONTRIBUTING.md).
- Articles with enough depth that a maintainer can verify claims against the linked content.

### Exclude

- **Launch posts** and product announcements without technical depth.
- **Shallow trend pieces** that restate common talking points.
- **Vendor marketing** or hype-driven content.
- **Prompt collections** (see [docs/rejected.md](rejected.md)).

## Case Studies

**README section:** [Case Studies](../README.md#case-studies)
**Label:** `resource:case-study`

Real-world uses of multiple agents or agent flows in production, enterprise, or sustained operational settings.

### Representative examples

- Production deployments of agent teams
- Enterprise agent teams with documented workflows
- Software engineering agent teams (code, review, test, deploy)
- Research operations using coordinated research agents
- Operations or SRE agents with escalation between roles
- Support agents with handoff to specialists or humans

### Include

- Documented deployments where **multiple agents or explicit flows** were central to the solution.
- Case writeups with concrete context: problem, architecture, roles, and outcomes.
- Customer or internal stories that remain encyclopedic, not promotional.

### Exclude

- Hypothetical architectures with no real deployment narrative.
- Single-chatbot customer support stories without multi-agent or flow structure.
- Pure product pages without a case narrative.

## Examples and Templates

**README section:** [Examples and Templates](../README.md#examples-and-templates)
**Label:** `resource:example`

Runnable or forkable examples that demonstrate multi-agent coordination, not abstract pattern descriptions alone.

### Representative examples

- Supervisor-worker examples
- Agent graph examples
- Crew or team templates
- Debate or review workflows
- Multi-agent software engineering examples
- Agent handoff examples

### Include

- Repositories, notebooks, or templates readers can run or adapt to build agent factories.
- Minimal demos whose purpose is **illustration** of a topology or framework feature.
- Starter kits focused on multi-agent flows, handoffs, or shared workspaces.

### Exclude

- Production frameworks (use **Frameworks**).
- Pattern documentation without runnable or forkable artifacts (use **Coordination Patterns**).
- Prompt-only snippets without orchestration code or structure.

## Related Lists

**README section:** [Related Lists](../README.md#related-lists)
**Label:** `resource:related-list`

Curated awesome lists or indexes that complement this repository. This taxonomy does not define detailed include/exclude rules for Related Lists; apply [README.md](../README.md) scope and [CONTRIBUTING.md](../CONTRIBUTING.md) the same way as other sections.

### Include

- Awesome lists or resource indexes that substantially cover multi-agent orchestration, agent teams, or agent-factory topics.

### Exclude

- Generic AI lists with no meaningful multi-agent or flow-management focus.

## Category quick reference

| README section | `resource:*` label | Primary question |
|----------------|-------------------|------------------|
| Theories | `resource:theory` | What foundational idea organizes agent groups? |
| Coordination Patterns | `resource:pattern` | What reusable topology describes the work? |
| Frameworks | `resource:framework` | What software orchestrates multiple agents? |
| Protocols and Interfaces | `resource:protocol` | What contract governs interaction? |
| Benchmarks | `resource:benchmark` | How is group or flow behavior measured? |
| Research Papers | `resource:paper` | What research advances group or flow understanding? |
| Blog Posts | `resource:blog` | What technical article teaches real lessons? |
| Case Studies | `resource:case-study` | What real deployment used multiple agents or flows? |
| Examples and Templates | `resource:example` | What can readers run or fork? |
| Related Lists | `resource:related-list` | What complementary awesome list or index covers agent orchestration? |

## Automation, workflow gates, and Phase 7 content seeding

Category names, include/exclude rules, and examples in this document informed Phase 4 automated README checks (section validation, Contents alignment, entry format, alphabetization) implemented in [`internal/checks`](../internal/checks) with local commands `make check` and `make test`. Phase 5 GitHub workflow gates (CI, Link Check, Awesome Lint, and Scheduled Maintenance) run the same or equivalent checks — see [docs/review-policy.md](review-policy.md#automated-enforcement) and [CONTRIBUTING.md](../CONTRIBUTING.md#github-actions).

Phase 4 structural checks and Phase 5 workflow gates are **implemented** in this repository.

Phase 7 foundational content seeding has **started** in this batch for [Theories](../README.md#theories), [Coordination Patterns](../README.md#coordination-patterns), [Frameworks](../README.md#frameworks), and [Related Lists](../README.md#related-lists) in [README.md](../README.md). Blog Posts, Case Studies, Benchmarks, Research Papers, Protocols and Interfaces, and Examples and Templates remain empty for later batches. Maintainers still apply manual review for scope, quality, and category fit per [docs/review-policy.md](review-policy.md).
