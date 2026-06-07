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
- [BDI Software Agent](https://en.wikipedia.org/wiki/BDI_software_agent) - Deliberative agent architecture modeling beliefs, desires, and intentions for rational planning and coordinated delegation in multi-agent groups.
- [Blackboard Architecture](https://en.wikipedia.org/wiki/Blackboard_architecture) - Shared-workspace coordination model where specialists contribute to emergent group solutions.
- [Contract Net Protocol](https://en.wikipedia.org/wiki/Contract_Net_Protocol) - Market-based delegation protocol for task allocation among autonomous agents.
- [Distributed Artificial Intelligence](https://en.wikipedia.org/wiki/Distributed_artificial_intelligence) - Foundational field studying cooperation, coordination, and problem decomposition across networks of autonomous agents.
- [Hierarchical Task Network](https://en.wikipedia.org/wiki/Hierarchical_task_network) - Planning formalism for decomposing complex tasks through hierarchical delegation and coordinated subgoal orchestration among specialized agents.
- [Swarm Intelligence](https://en.wikipedia.org/wiki/Swarm_intelligence) - Collective coordination principles for decentralized group-level behavior.

## Coordination Patterns

Reusable system shapes for coordinating agent groups and flows.

- [Agent orchestration](https://openai.github.io/openai-agents-python/multi_agent/) - OpenAI Agents SDK guide comparing manager-style orchestration with specialist handoffs.
- [AI Agent Orchestration Patterns](https://learn.microsoft.com/en-us/azure/architecture/ai-ml/guide/ai-agent-design-patterns) - Architecture guide for sequential, concurrent, group-chat, and handoff orchestration topologies.
- [Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents) - Engineering patterns including orchestrator-workers, routing, and evaluator-optimizer flows.
- [Choose a design pattern for your agentic AI system](https://docs.cloud.google.com/architecture/choose-design-pattern-agentic-ai-system) - Google Cloud Architecture Center guide comparing multi-agent sequential, parallel, coordinator, swarm, and review-and-critique orchestration topologies for delegating specialized sub-tasks.
- [Improving Factuality and Reasoning in Language Models through Multiagent Debate](https://proceedings.mlr.press/v235/du24e.html) - ICML proceedings defining a debate topology where multiple model instances use multi-agent coordination through iterative argument exchange and critique to improve reasoning quality.
- [Multi-agent](https://langchain-ai.github.io/langgraph/concepts/multi_agent/) - LangGraph concepts for supervisor, swarm, and hierarchical handoff topologies.
- [Saga orchestration patterns](https://docs.aws.amazon.com/prescriptive-guidance/latest/agentic-ai-patterns/saga-orchestration-patterns.html) - AWS Prescriptive Guidance on centralized orchestration where coordinator agents decompose tasks, use delegation to specialist worker agents, and coordinate fault-tolerant multi-step agentic workflows.
- [Scheduler Agent Supervisor](https://learn.microsoft.com/en-us/azure/architecture/patterns/scheduler-agent-supervisor) - Azure Architecture Center pattern separating scheduler workflow orchestration, agent execution, and supervisor recovery for reliable multi-step delegation pipelines.

## Frameworks

Software that manages agent groups, flows, handoffs, or orchestration.

- [AgentScope](https://github.com/agentscope-ai/agentscope) - Agent-oriented programming framework for message-based multi-agent orchestration, pipeline coordination, and delegated tool-use workflows.
- [AutoGen](https://github.com/microsoft/autogen) - Framework for multi-agent conversations, group chat orchestration, and delegated agent workflows.
- [CAMEL](https://github.com/camel-ai/camel) - Communicative agents framework for role-playing multi-agent societies with task delegation and coordinated conversation orchestration.
- [CrewAI](https://github.com/crewAIInc/crewAI) - Role-based crew framework for task delegation and hierarchical agent coordination.
- [LangGraph](https://github.com/langchain-ai/langgraph) - Graph-based orchestration library for stateful multi-agent workflows and handoffs.
- [MetaGPT](https://github.com/FoundationAgents/MetaGPT) - Multi-agent framework that assigns software-company roles for coordination of planning and execution flows.
- [Semantic Kernel](https://github.com/microsoft/semantic-kernel) - Microsoft SDK for composing multi-agent workflows with planner orchestration, handoffs, and coordinated routing across specialized agent plugins.
- [Symphony](https://github.com/openai/symphony) - Orchestration spec and service for dispatching isolated coding-agent runs from issue-tracker workflows.

## Protocols and Interfaces

Standards or conventions for agent interaction, handoffs, and shared interfaces.

- [Agent2Agent Protocol](https://a2a-protocol.org/) - Open agent-to-agent interoperability specification for discovery, task delegation, and message exchange across heterogeneous agent runtimes.
- [AGNTCY Agent Connect Protocol](https://docs.agntcy.org/) - Specification for secure agent-to-agent connectivity, identity, and message routing across distributed multi-agent networks.
- [FIPA Agent Communication Language](http://www.fipa.org/repository/aclspecs.html) - Standard message-representation and interaction-protocol vocabulary for structured coordination among software agents.
- [Model Context Protocol](https://modelcontextprotocol.io/docs/getting-started/intro) - Open protocol for sharing tools, resources, and context between agents and runtimes in multi-step orchestration flows.
- [Open Agent Schema Framework](https://github.com/agntcy/oasf) - Schema specification for describing agent capabilities and interfaces to support routing and delegation in multi-agent systems.

## Benchmarks

Benchmarks that evaluate group or workflow behavior.

- [AgentBench](https://github.com/THUDM/AgentBench) - Multi-environment benchmark for evaluating LLM-as-agent reasoning, tool use, and multi-turn orchestration across diverse interactive task flows.
- [AgentBoard](https://hkust-nlp.github.io/agentboard/) - Analytical evaluation board for multi-turn LLM agents spanning embodied, web, game, and tool domains with group-level evaluation metrics.
- [MultiAgentBench](https://github.com/ulab-uiuc/MARBLE) - Benchmark for evaluating collaboration, competition, and coordination among LLM-based multi-agent systems across diverse interactive scenarios.
- [SWE-bench](https://www.swebench.com/) - Benchmark of real GitHub issues for evaluating agent workflows with task delegation across multi-step software engineering flows.
- [WebArena](https://webarena.dev/) - Self-hostable web environment benchmarking autonomous agents on long-horizon multi-step tasks across realistic site orchestration flows.

## Research Papers

Academic or technical papers about agent groups and flows.

- [A Survey on Large Language Model based Autonomous Agents](https://arxiv.org/abs/2308.11432) - Survey of LLM-based autonomous agents covering planning, memory, and action modules for multi-step orchestration and coordination in agent factories.
- [AgentPoison: Red-teaming LLM Agents via Poisoning Memory or Tool Knowledge](https://arxiv.org/abs/2407.12784) - Security research on poisoning shared memory and tool knowledge to test multi-agent safety and governance vulnerabilities in agent coordination workflows.
- [AgentVerse: Facilitating Multi-Agent Collaboration and Exploring Emergent Behaviors](https://arxiv.org/abs/2308.10848) - Framework for assembling multi-agent groups with customizable coordination topologies to study emergent collaboration and delegated task specialization.
- [AutoGen: Enabling Next-Gen LLM Applications via Multi-Agent Conversation](https://arxiv.org/abs/2308.08155) - Research on conversational multi-agent orchestration patterns for task delegation, group chat, and cooperative problem solving across specialized agents.
- [CAMEL: Communicative Agents for Mind Exploration of Large Language Model Society](https://arxiv.org/abs/2303.17760) - Framework for role-playing communicative agents that coordinate through structured dialogue to explore emergent multi-agent behavior and delegation strategies.
- [ChatEval: Towards Better LLM-based Evaluation through Multi-Agent Debate](https://arxiv.org/abs/2308.07201) - Evaluation framework using multi-agent debate among LLM judges to improve response quality assessment through coordinated critique and group-level evaluation.
- [Communicative Agents for Software Development](https://arxiv.org/abs/2307.07924) - ChatDev paper describing multi-agent software teams with role-based delegation, handoffs, and coordinated orchestration flows across development phases.
- [Exploring Collaboration Mechanisms for LLM Agents: A Social Psychology View](https://arxiv.org/abs/2310.02124) - Study applying social psychology principles to LLM agent group dynamics for improving cooperation, coordination, and role-based delegation in multi-agent teams.
- [Generative Agents: Interactive Simulacra of Human Behavior](https://arxiv.org/abs/2304.03442) - Architecture for believable agent societies with shared state, memory, and planning that coordinate daily activities in interactive simulated environments.
- [HuggingGPT: Solving AI Tasks with ChatGPT and its Friends in Hugging Face](https://arxiv.org/abs/2303.17580) - LLM-powered orchestration system that delegates subtasks to specialized expert models through planning, routing, and coordinated handoffs across heterogeneous agents.
- [Large Language Model based Multi-Agents: A Survey of Progress and Challenges](https://arxiv.org/abs/2402.01680) - IJCAI survey of LLM multi-agent systems covering agent communication, profiling, coordination paradigms, and group-level evaluation benchmarks.
- [MetaGPT: Meta Programming for A Multi-Agent Collaborative Framework](https://arxiv.org/abs/2308.00352) - Multi-agent collaborative framework assigning software-company roles with structured message passing for coordination of planning and execution flows.
- [Mindstorms in Natural Language-Based Societies of Mind](https://arxiv.org/abs/2305.17066) - Research on natural-language societies of mind where multiple LLM agents use modular specialization and emergent group-level coordination flows for collective reasoning.
- [More Agents Is All You Need](https://arxiv.org/abs/2312.07974) - Empirical study showing sampling-and-voting across multiple LLM agent instances improves task performance through parallel coordination and majority orchestration.
- [ProAgent: Building Proactive Cooperative Agents with Large Language Models](https://arxiv.org/abs/2308.11339) - Framework for proactive cooperative agents that anticipate teammate needs and coordinate embodied multi-agent collaboration through shared planning and delegation.

## Blog Posts

Technical writing that explains real patterns, architectures, or failures in multi-agent systems.

- [Build multi-agent systems with LangGraph and Amazon Bedrock](https://aws.amazon.com/blogs/machine-learning/build-multi-agent-systems-with-langgraph-and-amazon-bedrock/) - AWS blog post on designing supervisor-worker and peer orchestration workflows with LangGraph state graphs and coordinated agent handoffs on Amazon Bedrock.
- [How to build Agentic Systems: The Missing Architecture for Production AI Agents](https://blog.crewai.com/agentic-systems-with-crewai/) - CrewAI engineering post on combining deterministic Flow backbones with autonomous Crew delegation so multi-agent orchestration stays traceable, observable, and production-safe.
- [How we built our multi-agent research system](https://www.anthropic.com/engineering/multi-agent-research-system) - Anthropic engineering post on Claude Research's orchestrator-worker architecture with parallel subagents, tool design, and production coordination lessons.
- [How We Deployed our Multi-Agent Flow to LangGraph Cloud](https://www.langchain.com/blog/how-we-deployed-our-multi-agent-flow-to-langgraph-cloud-2) - LangChain post on shipping a GPT-Researcher multi-agent graph to LangGraph Cloud with API-triggered orchestration, shared state coordination, and LangSmith group-level evaluation traces.
- [LangGraph: Multi-Agent Workflows](https://www.langchain.com/blog/langgraph-multi-agent-workflows) - LangChain post explaining multi-agent orchestration patterns implemented with LangGraph graphs, including network, supervisor, and hierarchical coordination designs.
- [Magentic-One: A Generalist Multi-Agent System for Solving Complex Tasks](https://www.microsoft.com/en-us/research/articles/magentic-one-a-generalist-multi-agent-system-for-solving-complex-tasks/) - Microsoft Research article on an orchestrator-led AutoGen multi-agent architecture with task-ledger planning, progress tracking, specialized agent delegation, and error-recovery coordination for open-ended web and file workflows.
- [Multi Agent Collaboration with Strands](https://aws.amazon.com/blogs/devops/multi-agent-collaboration-with-strands/) - AWS DevOps blog on evolving supervisor-into-arbiter orchestration with semantic blackboard coordination, dynamic agent routing, and delegated task planning across adaptive multi-agent teams.
- [Multi-Agent AI Coordination: The Distributed Systems Challenge Nobody's Talking About](https://jonathangardner.io/multi-agent-ai-is-a-distributed-systems-problem-plan-accordingly/) - Essay framing multi-agent coordination as a distributed systems problem covering orchestration, timeouts, circuit breakers, and saga-style recovery.
- [Why agent orchestration is harder than kubernetes - Lessons while building Agentflow](https://dev.to/siddhantkcode/why-agent-orchestration-is-harder-than-kubernetes-lessons-while-building-agentflow-4jm3) - DEV Community post contrasting container orchestration with agent workflows on semantic checkpointing, failure detection, and dynamic task decomposition.
- [Why Multi-Agent LLM Systems Fail (and How to Build Ones That Don't)](https://tianpan.co/blog/2025-10-14-why-multi-agent-llm-systems-fail) - Failure analysis of multi-agent LLM execution traces covering specification, inter-agent coordination misalignment, and termination issues with mitigation guidance.

## Case Studies

Real-world uses of multiple agents or agent flows.

- [Allianz Project Nemo Agentic AI Claims](https://www.insurancenews.com.au/insurtech/ai-trial-provides-blueprint-for-future-allianz-says) - Allianz case narrative on Project Nemo, a seven-agent claims workflow with planner orchestration, specialist verification agents, human payout handoffs, and roughly 80% faster low-complexity claim settlement in Australia.
- [Faire: Swarm-Coding with Multiple Background Agents for Large-Scale Code Maintenance](https://www.zenml.io/llmops-database/swarm-coding-with-multiple-background-agents-for-large-scale-code-maintenance) - Faire case writeup on coordinating hierarchical GitHub Copilot background agents with MCP delegation to parallelize feature-flag cleanup and maintenance pull requests across production repositories.
- [How Planview built a scalable AI Assistant for portfolio and project management using Amazon Bedrock](https://aws.amazon.com/blogs/machine-learning/how-planview-built-a-scalable-ai-assistant-for-portfolio-and-project-management-using-amazon-bedrock/) - Planview deployment story on Planview Copilot, a Bedrock-backed multi-agent assistant with a central orchestrator routing portfolio questions to specialized application agents for enterprise project workflows.
- [How Schroders built its multi-agent financial analysis research assistant](https://cloud.google.com/blog/topics/customers/how-schroders-built-its-multi-agent-financial-analysis-research-assistant) - Schroders customer case on a Vertex AI Agent Builder research assistant using specialized analyst agents, LangGraph parent-child orchestration, and evaluation gates to shorten equity research coordination cycles.
- [How Thomson Reuters built an Agentic Platform Engineering Hub with Amazon Bedrock AgentCore](https://aws.amazon.com/blogs/machine-learning/how-thomson-reuters-built-an-agentic-platform-engineering-hub-with-amazon-bedrock-agentcore/) - Thomson Reuters platform engineering case on an AgentCore hub where a central orchestrator delegates AWS provisioning and patching tasks to service-specific agents with human-in-the-loop approval handoffs.

## Examples and Templates

Runnable or forkable examples of multi-agent coordination.

- [AutoGen Python samples](https://github.com/microsoft/autogen/tree/main/python/samples) - Runnable AutoGen sample projects demonstrating conversational multi-agent orchestration, group chat coordination, and delegated tool-use workflows.
- [ChatDev](https://github.com/OpenBMB/ChatDev) - Runnable multi-agent software-development demo where specialized agents coordinate through chat-based handoffs across design, coding, and review phases.
- [CrewAI Examples](https://github.com/crewAIInc/crewAI-examples) - Forkable crew templates demonstrating hierarchical task delegation, role-based agent coordination, and orchestrated multi-step workflows.
- [GPT Researcher](https://github.com/assafelovic/gpt-researcher) - Multi-agent research assistant template with planner-editor coordination, parallel sub-agent delegation, and shared research orchestration flows.
- [LangGraph examples](https://github.com/langchain-ai/langgraph/tree/main/examples) - Runnable graph examples for supervisor, swarm, and hierarchical multi-agent orchestration patterns with stateful handoffs.
- [OpenAI Agents Python examples](https://github.com/openai/openai-agents-python/tree/main/examples) - Runnable samples for manager orchestration, specialist handoffs, and guardrailed multi-agent coordination flows.

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
