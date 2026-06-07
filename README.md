# Awesome AI Agent Factories [![Awesome](https://awesome.re/badge.svg)](https://awesome.re)

<p align="center">
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 48" width="200" height="48" role="img" aria-label="Agent factory mark: three connected nodes">
    <rect x="4" y="8" width="32" height="32" rx="4" fill="none" stroke="#333" stroke-width="2"/>
    <rect x="84" y="8" width="32" height="32" rx="4" fill="none" stroke="#333" stroke-width="2"/>
    <rect x="164" y="8" width="32" height="32" rx="4" fill="none" stroke="#333" stroke-width="2"/>
    <line x1="36" y1="24" x2="84" y2="24" stroke="#333" stroke-width="2"/>
    <line x1="116" y1="24" x2="164" y2="24" stroke="#333" stroke-width="2"/>
  </svg>
</p>

> A curated list of theories, frameworks, benchmarks, research, and writing about managing groups of AI agents and their flows.

AI agent factories are systems for coordinating groups of agents: assigning roles, routing tasks, managing handoffs, supervising execution, sharing state, and evaluating group-level behavior.

## Contents

<!--lint disable awesome-toc-->
- [Scope](#scope)
- [Theories](#theories)
- [Coordination Patterns](#coordination-patterns)
- [Frameworks](#frameworks)
- [Protocols and Interfaces](#protocols-and-interfaces)
- [Benchmarks](#benchmarks)
- [Research Papers](#research-papers)
- [Blog Posts](#blog-posts)
- [Case Studies](#case-studies)
- [Examples and Templates](#examples-and-templates)
<!--lint enable awesome-toc-->

## Scope

This list focuses on AI agent factories: systems, theories, frameworks, and practices for managing groups of agents and their flows.

Included:

- Multi-agent orchestration
- Agent teams, swarms, crews, societies, and supervisors
- Task delegation, routing, handoffs, and flow control
- Agent communication protocols
- Shared state and shared workspaces
- Group-level benchmarks and evaluation methods
- Research on agent coordination, cooperation, competition, and organization
- Production case studies involving groups of agents

Excluded unless directly relevant:

- Single-agent chatbots
- Generic LLM SDKs
- Prompt collections
- Standalone RAG tools
- Generic observability tools
- Generic workflow engines with no agent-specific usage
- Model provider SDKs

## Theories

Foundational ideas for organizing agent groups.

- [Actor Model](https://en.wikipedia.org/wiki/Actor_model) - Message-passing model where autonomous agents coordinate through asynchronous delegation.
- [An Introduction to MultiAgent Systems](https://www.cs.ox.ac.uk/people/michael.wooldridge/pubs/imas/IMAS2e.html) - Textbook on multi-agent coordination, communication, and cooperation.
- [Blackboard Architecture](https://en.wikipedia.org/wiki/Blackboard_architecture) - Shared-workspace coordination model where specialists contribute to emergent group solutions.
- [Contract Net Protocol](https://en.wikipedia.org/wiki/Contract_Net_Protocol) - Market-based delegation protocol for task allocation among autonomous agents.
- [Swarm Intelligence](https://en.wikipedia.org/wiki/Swarm_intelligence) - Collective coordination principles for decentralized group-level behavior.

## Coordination Patterns

Reusable system shapes for coordinating agent groups and flows.

- [Agent orchestration](https://openai.github.io/openai-agents-python/multi_agent/) - OpenAI Agents SDK guide comparing manager-style orchestration with specialist handoffs.
- [AI Agent Orchestration Patterns](https://learn.microsoft.com/en-us/azure/architecture/ai-ml/guide/ai-agent-design-patterns) - Architecture guide for sequential, concurrent, group-chat, and handoff orchestration topologies.
- [Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents) - Engineering patterns including orchestrator-workers, routing, and evaluator-optimizer flows.
- [Multi-agent](https://langchain-ai.github.io/langgraph/concepts/multi_agent/) - LangGraph concepts for supervisor, swarm, and hierarchical handoff topologies.

## Frameworks

Software that manages agent groups, flows, handoffs, or orchestration.

- [AutoGen](https://github.com/microsoft/autogen) - Framework for multi-agent conversations, group chat orchestration, and delegated agent workflows.
- [CrewAI](https://github.com/crewAIInc/crewAI) - Role-based crew framework for task delegation and hierarchical agent coordination.
- [LangGraph](https://github.com/langchain-ai/langgraph) - Graph-based orchestration library for stateful multi-agent workflows and handoffs.
- [MetaGPT](https://github.com/FoundationAgents/MetaGPT) - Multi-agent framework that assigns software-company roles for coordination of planning and execution flows.
- [Symphony](https://github.com/openai/symphony) - Orchestration spec and service for dispatching isolated coding-agent runs from issue-tracker workflows.

## Protocols and Interfaces

Standards or conventions for agent interaction, handoffs, and shared interfaces.

## Benchmarks

Benchmarks that evaluate group or workflow behavior.

## Research Papers

Academic or technical papers about agent groups and flows.

## Blog Posts

Technical writing that explains real patterns, architectures, or failures in multi-agent systems.

## Case Studies

Real-world uses of multiple agents or agent flows.

## Examples and Templates

Runnable or forkable examples of multi-agent coordination.

## Related Lists

Other awesome lists and curated resources related to agent orchestration and multi-agent systems.

- [Awesome AI agents](https://github.com/e2b-dev/awesome-ai-agents) - Curated index of autonomous and multi-agent agent projects covering orchestration frameworks and coordination tooling.
- [Awesome Generative Multi-Agent Systems](https://github.com/richardblythman/awesome-multi-agent-systems) - Curated libraries and resources for generative multi-agent orchestration and coordination patterns.
- [Awesome Multi-Agent Learning](https://github.com/WeiChengTseng/awesome-multi-agent) - Curated multi-agent reinforcement-learning papers and tutorials on coordination and delegation strategies.
- [Awesome Multi-Agent Papers](https://github.com/kyegomez/awesome-multi-agent-papers) - Compilation of multi-agent systems research papers on coordination, orchestration, and group-level evaluation.

## Contributing

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for the full submission path.

**Scope** — whether your resource fits agent-factory topics (groups of agents, flows, orchestration).

**Category** — which of the ten curated README sections applies (Theories, Coordination Patterns, Frameworks, Protocols and Interfaces, Benchmarks, Research Papers, Blog Posts, Case Studies, Examples and Templates, or Related Lists).

**Entry format** — `- [Resource Name](URL) - Description.` with the exact resource name as link text.

**Agent-factory relevance** — what to state in your pull request about coordination, orchestration, handoffs, or related flows.

Before opening a pull request, use GitHub's [**pull request template**](.github/PULL_REQUEST_TEMPLATE.md) (pre-filled for new PRs) and read [docs/taxonomy.md](docs/taxonomy.md) for category definitions and [docs/review-policy.md](docs/review-policy.md) for the maintainer review checklist you can self-check against. For issues, pick a structured form from the GitHub issue chooser (blank issues are disabled).

[MAINTAINERS.md](MAINTAINERS.md) describes review and merge policy.

## Community

See [CODE_OF_CONDUCT.md](https://github.com/portpowered/awesome-agent-factories/blob/main/CODE_OF_CONDUCT.md) for behavior expectations in issues and pull requests. See [SECURITY.md](https://github.com/portpowered/awesome-agent-factories/blob/main/SECURITY.md) for reporting malicious, typosquatted, or compromised linked resources.

---

[![CC0](https://mirrors.creativecommons.org/presskit/buttons/88x31/svg/cc-zero.svg)](https://creativecommons.org/publicdomain/zero/1.0/)

List content is released under [CC0-1.0](LICENSE). Linked third-party resources retain their own licenses.
