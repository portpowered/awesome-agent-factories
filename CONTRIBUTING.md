# Contributing

Thank you for helping improve **Awesome AI Agent Factories**. This document covers submission format, scope, category rules, and discipline so your pull request matches list conventions before maintainers review it.

See [README.md](README.md) for the canonical scope statement and section headings.

For detailed category definitions, examples, and include/exclude rules, see [docs/taxonomy.md](docs/taxonomy.md). For the full maintainer review checklist (which you can self-check before submitting), see [docs/review-policy.md](docs/review-policy.md).

## One resource per pull request

Each pull request must add or change **one resource** only.

- Do not bundle multiple unrelated entries in a single pull request.
- Do not mix list edits with unrelated repository changes.
- If you need to fix formatting for an existing entry, limit the pull request to that single entry.

## Entry format

Add entries to the correct section in [README.md](README.md) using this format:

```markdown
- [Resource Name](URL) - Description.
```

Requirements:

- **Link text** must be the **exact resource name** (official project, paper, or article title). Do not use generic link text such as "here", "link", or a domain name unless that is the official name.
- **URL** must point to the canonical resource (project homepage, paper page, or primary article URL).
- **Description** follows the hyphen separator and must end with a **period**.

### Valid example

```markdown
- [AutoGen](https://github.com/microsoft/autogen) - A framework for building multi-agent applications with conversable agents and group chat orchestration.
```

### Invalid examples

```markdown
- [Click here](https://example.com/framework) - A revolutionary game-changing platform.
```

Problems: generic link text, marketing language, description tone is promotional rather than factual.

```markdown
- [Example Framework](https://example.com/framework) - Multi-agent orchestration library
```

Problems: description does not end with a period.

```markdown
- [example.com](https://example.com/framework) - Framework for agent coordination.
```

Problems: link text is a domain, not the resource name (unless the domain is the official name).

## Descriptions

Descriptions must be:

- **Concise** — one or two sentences; state what the resource is and why it matters for agent-factory work.
- **Factual** — describe capabilities, scope, or content; do not sell or hype the resource.
- **Terminal punctuation** — every description must end with a period.

### Marketing language is prohibited

Do not use promotional or superlative phrasing in descriptions. Maintain a neutral encyclopedic tone.

Disallowed examples (unless they appear in an official title you are quoting):

- "revolutionary", "game-changing", "best", "ultimate", "cutting-edge"
- "must-have", "industry-leading", "unparalleled", "world-class"
- Calls to action such as "check it out" or "you won't want to miss this"

Prefer factual statements such as:

- "A library for supervisor-worker multi-agent workflows."
- "A survey of coordination patterns in LLM-based agent systems."

## Duplicate links

Do not add a URL that already appears anywhere in [README.md](README.md).

Before opening a pull request:

- Search [README.md](README.md) for the URL you plan to add.
- If the resource is already listed under another name or section, do not submit a duplicate; open an issue from the GitHub issue chooser (for example **Suggest a resource** or **Request resource removal**) or start a discussion instead.

Maintainers will reject pull requests that introduce duplicate links.

## Alphabetical order

Within each README section, entries must be sorted in **alphabetical order by link text** (the text inside `[...]`).

- Compare case-insensitively unless two names differ only by case.
- Insert your entry in the correct position; do not append to the end of a section if that breaks order.
- When renaming link text, re-sort the affected section if needed.

## Scope

This list focuses on **AI agent factories**: systems, theories, frameworks, and practices for managing **groups of agents** and their **flows**.

### Included

Submissions should fit one or more of these themes:

- Multi-agent orchestration
- Agent teams, swarms, crews, societies, and supervisors
- Task delegation, routing, handoffs, and flow control
- Agent communication protocols
- Shared state and shared workspaces
- Group-level benchmarks and evaluation methods
- Research on agent coordination, cooperation, competition, and organization
- Production case studies involving groups of agents

### Excluded unless directly relevant

Do not submit resources that are primarily about the following unless they clearly address group coordination, flows, or orchestration:

- Single-agent chatbots
- Generic LLM SDKs
- Prompt collections
- Standalone RAG tools
- Generic observability tools
- Generic workflow engines with no agent-specific usage
- Model provider SDKs

When a borderline resource is directly relevant, explain in your pull request how it helps manage groups of agents or their flows.

## Categories

Add each resource to **one** existing README section. Choose the section that best matches the resource type.

Contributors must choose from these ten categories (README section headings):

1. **Theories** — Foundational ideas for organizing agent groups.
2. **Coordination Patterns** — Reusable system shapes for coordinating agent groups and flows.
3. **Frameworks** — Software that manages agent groups, flows, handoffs, or orchestration.
4. **Protocols and Interfaces** — Standards or conventions for agent interaction, handoffs, and shared interfaces.
5. **Benchmarks** — Benchmarks that evaluate group or workflow behavior.
6. **Research Papers** — Academic or technical papers about agent groups and flows.
7. **Blog Posts** — Technical writing about real patterns, architectures, or failures in multi-agent systems.
8. **Case Studies** — Real-world uses of multiple agents or agent flows.
9. **Examples and Templates** — Runnable or forkable examples of multi-agent coordination.
10. **Related Lists** — Other awesome lists and curated resources related to agent orchestration and multi-agent systems.

Do not invent new top-level README sections in a single-resource pull request.

For detailed definitions, representative examples, and include/exclude rules per category, see [docs/taxonomy.md](docs/taxonomy.md).

### New categories

If you believe the list needs a **new category**, open a pull request that adds **at least three entries** for that category in the same change.

- All three entries must meet scope and formatting rules.
- Each entry must be alphabetized within the new section.
- Explain in the pull request why the existing ten categories are insufficient.

Maintainers may reject category proposals with fewer than three qualifying entries.

## Section-specific rules

### Frameworks

Framework submissions must help users **manage groups of agents, flows, handoffs, or orchestration**.

Include:

- Multi-agent frameworks
- Agent graph frameworks
- Crew, team, or society frameworks
- Workflow-oriented agent frameworks
- Handoff-oriented frameworks

Exclude:

- Generic model or inference SDKs
- Single-agent chatbot frameworks
- Prompt-only libraries

If the project supports both single-agent and multi-agent modes, your description and pull request must focus on the multi-agent or orchestration capabilities.

### Blog posts

Blog posts must be **technical**. Acceptable topics include architecture breakdowns, production lessons, failure analyses, evaluation writeups, and multi-agent design patterns.

Exclude:

- Product launch posts
- Shallow trend pieces
- Vendor marketing or sponsored content
- Prompt collections without substantive systems discussion

### Research papers, case studies, benchmarks, and examples

These categories still require an agent-factory angle. A paper about a single model, a benchmark for isolated tool use, or an example of one chatbot is out of scope unless it materially addresses coordination, orchestration, or group-level behavior.

## Agent-factory relevance

Every submission must explain the **agent-factory angle**: how the resource helps people manage **groups of agents** or their **flows**.

In your pull request description, state how the resource relates to at least one of:

- Coordination
- Orchestration
- Delegation
- Routing
- Handoffs
- Shared state
- Group-level evaluation

Your README description should imply this relevance factually. Do not rely on marketing claims; state what the resource does for multi-agent or flow management.

### Valid relevance examples

- "Documents a supervisor-worker pattern for delegating subtasks across specialized agents."
- "Defines a message protocol for agent-to-agent handoffs in a shared workspace."
- "Benchmarks collaboration quality across teams of tool-using agents."

### Invalid relevance examples

- "Useful for anyone building with LLMs." (no group or flow angle)
- "Great SDK for chat applications." (single-agent focus)
- "Popular in the AI community." (no technical relevance)

## Local checks

Run these commands from the repository root before you open a pull request:

- **`make check`** — validate README structure and entry rules with the Go checker.
- **`make test`** — run Go checker tests (`go test ./...`).
- **`make all`** — same as `make check`; use this as the primary pre-submit shortcut.

Optional targets:

- **`make lint`** — Go static analysis with [golangci-lint](https://golangci-lint.run/welcome/install/). Install with `go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest` or follow upstream installation docs for your platform.
- **`make links`** — markdown link checking with [lychee](https://lychee.cli.rs/guides/getting-started/). Install with `brew install lychee` or follow upstream installation docs for your platform.

Thin shell wrappers `scripts/check-readme.sh` and `scripts/normalize-readme.sh` delegate to `make check` and `make fmt` when you prefer script entry points.

### GitHub Actions

These workflows mirror the local commands above (read-only; no repository mutations):

| Workflow | Triggers | Local equivalent |
| --- | --- | --- |
| [CI](.github/workflows/ci.yml) | Pull requests and pushes to `main` | `make test`, `make check` |
| [Link Check](.github/workflows/link-check.yml) | Pull requests to `main`, weekly (Mondays) | `make links` |
| [Awesome Lint](.github/workflows/awesome-lint.yml) | Pull requests and pushes to `main` | `npx awesome-lint` |
| [Scheduled Maintenance](.github/workflows/scheduled-maintenance.yml) | Monthly (1st of month) | `make check`, `make links` |

The CI workflow runs Go format checks, `go test ./...`, and README validation (`make check` / `internal/checks`). Run the local commands above to reproduce workflow failures before pushing. Link checks honor root [`.lychee.toml`](.lychee.toml); README validation rules live in [`internal/checks`](internal/checks)—workflow YAML does not duplicate them. Maintainer merge expectations and recurring stewardship cadence are in [MAINTAINERS.md](MAINTAINERS.md).

## Before you open a pull request

GitHub pre-fills the repository [**pull request template**](.github/PULL_REQUEST_TEMPLATE.md) when you open a new pull request. Complete its resource metadata and checklist acknowledgements before requesting review.

For issues, GitHub disables blank reports; use the issue chooser (**Suggest a resource**, **Request resource removal**, or **Report a broken link**) or the contact links to this document and [docs/review-policy.md](docs/review-policy.md).

1. Confirm your change adds **one resource** only (unless proposing a new category with at least three entries).
2. Choose one of the **ten README categories** (or meet the new-category rule above); use [docs/taxonomy.md](docs/taxonomy.md) when section fit is unclear.
3. Verify the resource fits **scope** and is not a generic single-agent SDK or other excluded type unless directly relevant.
4. Use the required entry format with the **exact resource name** as link text.
5. Write a **concise, factual** description that **ends with a period**, avoids marketing language, and reflects the **agent-factory angle**.
6. Verify the URL is **not already present** in [README.md](README.md).
7. Place the entry in the correct section in **alphabetical order** by link text.
8. In the pull request description, explain **how the resource helps manage groups of agents or their flows**.
9. Self-check against the questions in [docs/review-policy.md](docs/review-policy.md) to reduce review churn.

Automated checks cover README structure, Go tests, link health, and awesome-list conventions; maintainers still review scope, canonical URLs, and quality manually. Following these rules reduces review churn and helps your pull request move forward.

## Security

If you find a malicious, typosquatted, compromised, or otherwise unsafe linked resource, see [SECURITY.md](SECURITY.md) for what we watch for and how to report concerns.

## Related policies

- [README.md](README.md) — scope statement and section headings
- [docs/taxonomy.md](docs/taxonomy.md) — category definitions, examples, and include/exclude rules
- [docs/review-policy.md](docs/review-policy.md) — maintainer PR review checklist and recommended labels
- [MAINTAINERS.md](MAINTAINERS.md) — review, merge, stale pull request, and removal policy
- [SECURITY.md](SECURITY.md) — malicious links, typosquatting, supply-chain concerns, and reporting
- [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) — community behavior in issues and pull requests
- [LICENSE](LICENSE) — license for curated list content
