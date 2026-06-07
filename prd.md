# PRD: Phase 7 Protocols, Benchmarks, and Research Papers Seed

## Introduction

Extend Phase 7 README content seeding from the four foundational sections already on `main` (Theories, Coordination Patterns, Frameworks, and Related Lists) into Protocols and Interfaces, Benchmarks, and Research Papers. This batch adds the next high-confidence curated slice of canonical resources that directly enable or evaluate multi-agent coordination, agent-to-agent contracts, handoffs, shared state, workflow orchestration, or group-level agent performance. Blog Posts, Case Studies, and Examples and Templates remain empty for a later batch.

## Context

**Customer ask:** Populate Protocols and Interfaces, Benchmarks, and Research Papers with roughly 3–5 stable entries per protocol/benchmark section and 5–8 stable entries for Research Papers. Use exact resource names as link text, concise factual descriptions ending in periods, and alphabetical order within each section. Update `docs/taxonomy.md` so its Phase 7 status note reflects that foundational seeding is already present on `main` for Theories, Coordination Patterns, Frameworks, and Related Lists, and that this batch extends seeding into the three target sections. Keep `make check`, `make test`, and `git diff --check` passing.

**Problem:** The README has populated foundational sections but leaves Protocols and Interfaces, Benchmarks, and Research Papers empty despite taxonomy definitions and automated checks already enforcing entry format, alphabetization, scope keywords, and section structure. Contributors and planners lack canonical starting points for contracts, evaluation harnesses, and research literature in the agent-factory domain.

**Solution:** Add a curated, non-promotional seed set to each target README section following existing entry conventions, then update the taxonomy Phase 7 status paragraph to document completed and remaining seeding scope accurately.

## Goals

- Seed Protocols and Interfaces with 5 canonical agent-interaction and contract specifications.
- Seed Benchmarks with 5 canonical group- or workflow-oriented evaluation suites.
- Seed Research Papers with 7 widely cited papers on multi-agent coordination, orchestration, and societies.
- Keep entries alphabetized by link text, scope-aligned, and free of banned marketing language.
- Update `docs/taxonomy.md` Phase 7 status to reflect sections seeded on `main` versus sections deferred.
- Pass all repository quality gates with no unrelated changes.

## User Stories

### US-001: Seed Protocols and Interfaces section

**Description:** As a contributor designing multi-agent handoffs, I want the Protocols and Interfaces section to list canonical interaction contracts so I can find interoperable standards without guessing category fit.

**Acceptance Criteria:**

- [ ] `## Protocols and Interfaces` contains exactly these five entries in alphabetical order by link text:
  - [Agent2Agent Protocol](https://google.github.io/A2A/) — open agent-to-agent interoperability specification for discovery, task delegation, and message exchange across heterogeneous agent runtimes.
  - [AGNTCY Agent Connect Protocol](https://docs.agntcy.org/) — specification for secure agent-to-agent connectivity, identity, and message routing across distributed multi-agent networks.
  - [FIPA Agent Communication Language](http://www.fipa.org/repository/aclspecs.html) — standard message-representation and interaction-protocol vocabulary for structured coordination among software agents.
  - [Model Context Protocol](https://modelcontextprotocol.io/) — open protocol for sharing tools, resources, and context between agents and runtimes in multi-step orchestration flows.
  - [Open Agent Schema Framework](https://github.com/agntcy/oasf) — schema specification for describing agent capabilities and interfaces to support routing and delegation in multi-agent systems.
- [ ] Each entry uses format `- [Resource Name](URL) - Description.` with description ending in a period and containing at least one agent-factory scope keyword enforced by `make check`.
- [ ] No new entries appear in Blog Posts, Case Studies, or Examples and Templates.
- [ ] `make check` exits 0 after the README update.
- [ ] Typecheck passes.

### US-002: Seed Benchmarks section

**Description:** As a maintainer evaluating agent factories, I want the Benchmarks section to list canonical group- and workflow-oriented evaluation suites so readers can measure coordination and multi-step flow behavior.

**Acceptance Criteria:**

- [ ] `## Benchmarks` contains exactly these five entries in alphabetical order by link text:
  - [AgentBench](https://github.com/THUDM/AgentBench) — benchmark suite measuring LLM agents on multi-turn tool use, delegation, and environment interaction across diverse coordination tasks.
  - [AgentBoard](https://github.com/hkust-nlp/AgentBoard) — evaluation framework tracking longitudinal multi-step agent progress across planning, memory, and handoff-heavy scenarios.
  - [MultiAgentBench](https://github.com/MultiAgentBench/MultiAgentBench) — benchmark measuring collaboration quality, coordination, and role specialization across multi-agent teams.
  - [SWE-bench](https://www.swebench.com/) — software-engineering benchmark evaluating multi-step coding workflows with tool use, delegation, and verification between agent roles.
  - [WebArena](https://webarena.dev/) — realistic web-environment benchmark for autonomous agents performing multi-step navigation, planning, and orchestrated tool actions.
- [ ] Each entry follows README entry format and passes scope-keyword and marketing-phrase checks.
- [ ] No duplicate URLs or resource names (case-insensitive) with entries in other README sections.
- [ ] `make check` exits 0 after the README update.
- [ ] Typecheck passes.

### US-003: Seed Research Papers section

**Description:** As a researcher or practitioner, I want the Research Papers section to list foundational multi-agent coordination literature so the list anchors academic context for agent-factory design.

**Acceptance Criteria:**

- [ ] `## Research Papers` contains exactly these seven entries in alphabetical order by link text:
  - [A Survey on Large Language Model based Autonomous Agents](https://arxiv.org/abs/2308.11432) — survey of autonomous-agent architectures covering planning, memory, tool use, and multi-agent orchestration patterns.
  - [AutoGen: Enabling Next-Gen LLM Applications via Multi-Agent Conversation](https://arxiv.org/abs/2308.08155) — framework paper on conversational multi-agent orchestration, role specialization, and delegated task flows.
  - [CAMEL: Communicative Agents for Mind Exploration of Large Language Model Society](https://arxiv.org/abs/2303.17760) — role-playing agent society design for scalable multi-agent communication and task coordination.
  - [Communicative Agents for Software Development](https://arxiv.org/abs/2307.07924) — multi-agent software-development workflow study on role-based delegation, handoffs, and collaborative orchestration.
  - [Generative Agents: Interactive Simulacra of Human Behavior](https://arxiv.org/abs/2304.03442) — architecture for agent societies with shared memory, planning, and emergent group-level coordination in interactive environments.
  - [Large Language Model based Multi-Agents: A Survey of Progress and Challenges](https://arxiv.org/abs/2402.01680) — survey of LLM multi-agent systems covering coordination, communication, planning, and group-level evaluation methods.
  - [MetaGPT: Meta Programming for A Multi-Agent Collaborative Framework](https://arxiv.org/abs/2308.00352) — multi-agent framework paper on standardized operating procedures, role workflows, and orchestrated software-team coordination.
- [ ] Paper link text uses exact publication titles; descriptions are factual, end with periods, and include agent-factory scope keywords.
- [ ] Framework README entries (for example AutoGen, MetaGPT) remain in Frameworks; paper entries link to arXiv or equivalent research URLs, not duplicate framework homepages.
- [ ] `make check` exits 0 after the README update.
- [ ] Typecheck passes.

### US-004: Update taxonomy Phase 7 seeding status

**Description:** As a planner or maintainer reading taxonomy guidance, I want the Phase 7 status note to accurately describe which README sections are seeded on `main` and which remain for later batches.

**Acceptance Criteria:**

- [ ] `docs/taxonomy.md` Phase 7 status paragraph states that foundational seeding is present on `main` for Theories, Coordination Patterns, Frameworks, and Related Lists.
- [ ] The same paragraph states this batch extends seeding into Protocols and Interfaces, Benchmarks, and Research Papers.
- [ ] The paragraph explicitly defers Blog Posts, Case Studies, and Examples and Templates to a later batch.
- [ ] No contradictory claims remain that Protocols and Interfaces, Benchmarks, or Research Papers are empty after this batch lands.
- [ ] Typecheck passes.

### US-005: Verify repository quality gate

**Description:** As a maintainer merging Phase 7 content, I want end-to-end verification that the seeded README and taxonomy update pass all automated checks without whitespace regressions.

**Acceptance Criteria:**

- [ ] From repository root, `make check` exits 0.
- [ ] From repository root, `make test` exits 0.
- [ ] From repository root, `git diff --check` exits 0.
- [ ] README Contents anchor list still matches section headings; no new top-level sections are introduced.
- [ ] Blog Posts, Case Studies, and Examples and Templates sections remain without resource entries.
- [ ] Typecheck passes.
- [ ] Tests pass.

## Functional Requirements

- FR-1: Add five alphabetized resource entries under `## Protocols and Interfaces` per US-001.
- FR-2: Add five alphabetized resource entries under `## Benchmarks` per US-002.
- FR-3: Add seven alphabetized resource entries under `## Research Papers` per US-003.
- FR-4: Update `docs/taxonomy.md` Phase 7 seeding status per US-004.
- FR-5: All new descriptions must satisfy automated scope-keyword checks and avoid banned marketing phrases defined in `internal/checks`.
- FR-6: No duplicate resource names or normalized URLs across README sections.

## Non-Goals

- Seeding Blog Posts, Case Studies, or Examples and Templates.
- Adding new README sections or changing scope exclusions.
- Modifying Go check rules, GitHub workflows, or issue templates.
- Link-check remediation for third-party sites beyond what existing automation handles.
- Expanding Related Lists, Theories, Coordination Patterns, or Frameworks beyond their current seeded entries.

## Technical Design

This batch is README and taxonomy documentation only. No application code changes are required unless a check failure exposes a pre-existing defect.

**Content contract:** Each entry is a single markdown list item:

```markdown
- [Exact Resource Name](canonical-url) - Factual description ending with a period.
```

**Validation path:** `make check` runs Go-based README validators in `internal/checks` for section presence, Contents alignment, entry format, alphabetical order, duplicate detection, scope keywords, and banned marketing phrases. `make test` runs the same package's unit tests.

**Dependency order:** Protocols, Benchmarks, and Research Papers seeding stories are independent and may land in one commit, but acceptance is verified incrementally so partial failures are diagnosable. Taxonomy update depends on README content being finalized. Quality gate runs last.

## Supporting Considerations

- Prefer durable canonical URLs (project homepages, arXiv abs links, specification docs) over ephemeral blog or social links.
- Descriptions must imply agent-factory relevance through scope keywords (`coordination`, `orchestration`, `delegation`, `routing`, `handoffs`, `shared state`, `group-level evaluation`) without marketing tone.
- Framework papers (AutoGen, MetaGPT) belong in Research Papers; shipped frameworks stay in Frameworks with distinct URLs.
- MCP is included because taxonomy explicitly allows MCP resources relevant to multi-agent flow management.

## Success Metrics

- Three previously empty README sections contain stable, reviewer-verifiable canonical entries.
- `make check` and `make test` pass on the first merge attempt without follow-up content fixes.
- Taxonomy Phase 7 status requires no further correction before the next content batch.
- A contributor can self-categorize a protocol, benchmark, or paper submission by comparing against seeded examples.

## Open Questions

None. Resource selection, counts, and deferred sections are specified in the customer ask.
