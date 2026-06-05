# Awesome AI Agent Factories — Build Checklist

> Goal: build a curated Awesome List focused on managing groups of AI agents and their flows: coordination, orchestration, delegation, routing, handoffs, shared state, evaluation, and governance.

## 0. Repository Structure

Target repo:

```txt
awesome-ai-agent-factories/
├── README.md
├── LICENSE
├── CONTRIBUTING.md
├── CODE_OF_CONDUCT.md
├── SECURITY.md
├── MAINTAINERS.md
├── Makefile
├── go.mod
├── go.sum
├── .gitignore
├── .editorconfig
├── .markdownlint.yaml
├── .lychee.toml
├── .golangci.yml
├── .github/
│   ├── PULL_REQUEST_TEMPLATE.md
│   ├── ISSUE_TEMPLATE/
│   │   ├── add-resource.yml
│   │   ├── remove-resource.yml
│   │   ├── broken-link.yml
│   │   └── config.yml
│   └── workflows/
│       ├── ci.yml
│       ├── link-check.yml
│       ├── awesome-lint.yml
│       └── scheduled-maintenance.yml
├── internal/
│   └── checks/
│       ├── main.go
│       ├── readme.go
│       ├── links.go
│       ├── sections.go
│       ├── alphabetical.go
│       └── duplicates.go
├── scripts/
│   ├── check-readme.sh
│   └── normalize-readme.sh
└── docs/
    ├── taxonomy.md
    ├── review-policy.md
    ├── historical.md
    └── rejected.md
```

## 1. Phase One — Set Up the Single Main Document

The first deliverable is one strong `README.md`.

### README.md checklist

* [ ] Add title: `Awesome AI Agent Factories`.
* [ ] Add header image or logo.
* [ ] Add Awesome badge.
* [ ] Add short description.
* [ ] Add scope statement.
* [ ] Add `Contents`.
* [ ] Add section skeletons.
* [ ] Add contribution note at bottom.
* [ ] Add license note at bottom.
* [ ] Add neutral, non-promotional language.

Suggested opening:

```md
# Awesome AI Agent Factories [![Awesome](https://awesome.re/badge.svg)](https://awesome.re)

> A curated list of theories, frameworks, benchmarks, research, and writing about managing groups of AI agents and their flows.

AI agent factories are systems for coordinating groups of agents: assigning roles, routing tasks, managing handoffs, supervising execution, sharing state, and evaluating group-level behavior.
```

### Initial README shape

```md
## Contents

- [Theories](#theories)
- [Coordination Patterns](#coordination-patterns)
- [Frameworks](#frameworks)
- [Protocols and Interfaces](#protocols-and-interfaces)
- [Benchmarks](#benchmarks)
- [Research Papers](#research-papers)
- [Blog Posts](#blog-posts)
- [Case Studies](#case-studies)
- [Examples and Templates](#examples-and-templates)
- [Related Lists](#related-lists)
```

### Scope section

```md
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
```

## 2. Phase Two — Add Repo Governance Files

Before adding lots of content, add the files that keep contributions sane.

### Required files

* [ ] `LICENSE`
* [ ] `CONTRIBUTING.md`
* [ ] `CODE_OF_CONDUCT.md`
* [ ] `SECURITY.md`
* [ ] `MAINTAINERS.md`
* [ ] `.gitignore`
* [ ] `.editorconfig`

### Recommended license

* [ ] Use `CC0-1.0` for the list content.
* [ ] Mention that linked resources retain their own licenses.

### CONTRIBUTING.md should define

* [ ] One resource per PR.
* [ ] Exact resource name as link text.
* [ ] Concise factual description.
* [ ] Description must end with a period.
* [ ] No marketing language.
* [ ] No duplicate links.
* [ ] Alphabetical order inside sections.
* [ ] New category requires at least 3 entries.
* [ ] Blog posts must be technical.
* [ ] Frameworks must manage groups, flows, handoffs, or orchestration.
* [ ] Generic single-agent SDKs are excluded.
* [ ] Every submission must explain the agent-factory angle.

### MAINTAINERS.md should define

* [ ] Current maintainers.
* [ ] Review responsibilities.
* [ ] Merge policy.
* [ ] Stale PR policy.
* [ ] Removal policy.
* [ ] Security/contact policy.

### SECURITY.md should cover

* [ ] Malicious links.
* [ ] Typosquatting.
* [ ] Supply-chain concerns.
* [ ] Abandoned or compromised projects.
* [ ] Where to report concerns.

## 3. Phase Three — Add Review Process

The list should reject weak entries before automation even runs.

### PR review checklist

Every PR should answer:

* [ ] What is the resource?
* [ ] Which section does it belong in?
* [ ] Does it manage groups of agents or flows?
* [ ] Does it involve coordination, orchestration, delegation, routing, handoffs, shared state, or group evaluation?
* [ ] Is it maintained, historically important, or academically relevant?
* [ ] Is the description factual and non-promotional?
* [ ] Is it a duplicate?
* [ ] Is the link stable?
* [ ] Is the entry alphabetized?

### Recommended labels

```txt
resource:theory
resource:pattern
resource:framework
resource:protocol
resource:benchmark
resource:paper
resource:blog
resource:case-study
resource:example
needs-scope-review
needs-link-review
needs-maintainer-review
duplicate
broken-link
historical
rejected
```

### Removal policy

Remove or move entries when:

* [ ] The link is dead for more than 30 days.
* [ ] The project is archived and not historically important.
* [ ] The description is misleading.
* [ ] The resource is no longer about agent groups or flows.
* [ ] The resource is primarily marketing.
* [ ] The project appears malicious, compromised, or spammy.

Use `docs/historical.md` for historically important but inactive resources.

Use `docs/rejected.md` for commonly submitted resources that do not fit the scope.

## 4. Phase Four — Add Automation

Use automation to enforce structure, link health, and contribution discipline.

### Makefile

```makefile
.PHONY: all lint check links fmt test

all: check

lint:
	golangci-lint run ./...

check:
	go run ./internal/checks

links:
	lychee README.md docs/*.md

fmt:
	gofmt -w internal

test:
	go test ./...
```

### Go module

```txt
go.mod
go.sum
```

Use Go for custom README checks:

* [ ] Duplicate links.
* [ ] Duplicate names.
* [ ] Alphabetical order.
* [ ] Required sections.
* [ ] Empty sections.
* [ ] Minimum entries per new section.
* [ ] Entry format.
* [ ] Description punctuation.
* [ ] Banned phrases.
* [ ] Suspicious URLs.
* [ ] Scope keywords, as a warning rather than a hard failure.

### internal/checks

Suggested checks:

```txt
internal/checks/main.go
internal/checks/readme.go
internal/checks/links.go
internal/checks/sections.go
internal/checks/alphabetical.go
internal/checks/duplicates.go
```

The custom checker should validate:

* [ ] README exists.
* [ ] Required sections exist.
* [ ] `Contents` matches headings.
* [ ] Entries use `- [Name](URL) - Description.` format.
* [ ] Descriptions end with a period.
* [ ] No duplicate URLs.
* [ ] No duplicate resource names.
* [ ] Entries are alphabetized within sections.
* [ ] No bare URLs.
* [ ] No tracking-heavy links where avoidable.
* [ ] No obvious marketing phrases such as `revolutionary`, `game-changing`, `best`, `ultimate`, `cutting-edge`, unless part of an official title.

### .golangci.yml

```yaml
run:
  timeout: 5m

linters:
  enable:
    - govet
    - staticcheck
    - errcheck
    - ineffassign
    - unused
    - gofmt
    - gofumpt
    - revive
    - misspell

issues:
  exclude-use-default: false
```

### .markdownlint.yaml

```yaml
MD013: false
MD033: false
MD041: false
MD024:
  siblings_only: true
```

### .lychee.toml

```toml
max_retries = 2
timeout = 20
accept = [200, 204, 206, 301, 302, 403, 429]

exclude = [
  "localhost",
  "127.0.0.1"
]
```

### .gitignore

```gitignore
.DS_Store
.idea/
.vscode/
*.log
tmp/
dist/
coverage.out
```

### .editorconfig

```ini
root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true

[*.md]
trim_trailing_whitespace = false

[*.go]
indent_style = tab

[*.{yml,yaml,toml,json,md}]
indent_style = space
indent_size = 2
```

## 5. Phase Five — Add GitHub Actions

### .github/workflows/ci.yml

```yaml
name: CI

on:
  pull_request:
    branches: [main]
  push:
    branches: [main]

jobs:
  go-checks:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: stable

      - name: Download dependencies
        run: go mod download

      - name: Format check
        run: test -z "$(gofmt -l internal)"

      - name: Go tests
        run: go test ./...

      - name: Custom README checks
        run: go run ./internal/checks
```

### .github/workflows/link-check.yml

```yaml
name: Link Check

on:
  pull_request:
    branches: [main]
  schedule:
    - cron: "0 8 * * 1"

jobs:
  link-check:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Check links
        uses: lycheeverse/lychee-action@v2
        with:
          args: --verbose --no-progress README.md docs/*.md
```

### .github/workflows/awesome-lint.yml

```yaml
name: Awesome Lint

on:
  pull_request:
    branches: [main]
  push:
    branches: [main]

jobs:
  awesome-lint:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4
      - run: npx awesome-lint
```

### .github/workflows/scheduled-maintenance.yml

```yaml
name: Scheduled Maintenance

on:
  schedule:
    - cron: "0 9 1 * *"

jobs:
  maintenance:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: stable

      - name: Run custom checks
        run: go run ./internal/checks

      - name: Check links
        uses: lycheeverse/lychee-action@v2
        with:
          args: --verbose --no-progress README.md docs/*.md
```

## 6. Phase Six — Add Issue and PR Templates

### .github/PULL_REQUEST_TEMPLATE.md

```md
## Resource

Name:

URL:

Section:

## Why it belongs

Explain how this resource helps people manage groups of agents or their flows.

## Checklist

- [ ] I added one resource only.
- [ ] This resource is about agent groups, flows, orchestration, routing, handoffs, shared state, or group-level evaluation.
- [ ] The entry uses the required format.
- [ ] The description is concise, factual, and non-promotional.
- [ ] The description ends with a period.
- [ ] The entry is alphabetized.
- [ ] This is not a duplicate.
- [ ] The link is live.
```

### .github/ISSUE_TEMPLATE/add-resource.yml

```yaml
name: Suggest a resource
description: Suggest a resource for the list.
title: "Add: "
labels: ["needs-scope-review"]
body:
  - type: input
    id: name
    attributes:
      label: Resource name
    validations:
      required: true

  - type: input
    id: url
    attributes:
      label: URL
    validations:
      required: true

  - type: dropdown
    id: section
    attributes:
      label: Section
      options:
        - Theories
        - Coordination Patterns
        - Frameworks
        - Protocols and Interfaces
        - Benchmarks
        - Research Papers
        - Blog Posts
        - Case Studies
        - Examples and Templates
        - Related Lists
    validations:
      required: true

  - type: textarea
    id: relevance
    attributes:
      label: Why does this belong?
      description: Explain how it helps manage groups of agents or their flows.
    validations:
      required: true
```

## 7. Phase Seven — Add Initial Content

Only after the document, governance, review process, and automation exist.

### Seed targets

* [ ] 8–12 Theories.
* [ ] 8–12 Coordination Patterns.
* [ ] 8–12 Frameworks.
* [ ] 5–8 Protocols and Interfaces.
* [ ] 5–8 Benchmarks.
* [ ] 15–25 Research Papers.
* [ ] 10–15 Blog Posts.
* [ ] 3–5 Case Studies.
* [ ] 5–10 Examples and Templates.
* [ ] 3–5 Related Lists.

### Content rules

Every item must answer one of these questions:

* [ ] How do groups of agents coordinate?
* [ ] How are tasks routed, delegated, or handed off?
* [ ] How is shared state managed?
* [ ] How are agent teams supervised?
* [ ] How are multi-agent flows evaluated?
* [ ] How are agent groups deployed or governed?
* [ ] How do humans interact with agent teams?

## 8. Phase Eight — Category Definitions

### Theories

Foundational ideas for organizing agent groups.

Examples:

* Agent societies
* Role-based organizations
* Hierarchical planning
* Blackboard systems
* Actor model
* Swarm intelligence
* Market-based coordination
* Human-agent teaming

### Coordination Patterns

Reusable system shapes.

Examples:

* Supervisor-worker
* Planner-executor
* Router-specialist
* Critic-reviewer
* Debate
* Swarm
* Pipeline
* Hub-and-spoke
* Shared workspace
* Blackboard
* Committee
* Auction or market-based allocation

### Frameworks

Software that manages agent groups, flows, handoffs, or orchestration.

Include:

* Multi-agent frameworks
* Agent graph frameworks
* Crew/team frameworks
* Workflow-oriented agent frameworks
* Handoff-oriented frameworks

Exclude:

* Generic model SDKs
* Single-agent chatbot frameworks
* Prompt-only libraries

### Protocols and Interfaces

Standards or conventions for interaction.

Include:

* Agent-to-agent communication
* Tool interfaces
* Shared memory interfaces
* Handoff schemas
* Workflow schemas
* Task queue interfaces
* Message formats
* MCP-related resources when relevant to multi-agent flow management

### Benchmarks

Benchmarks that evaluate group or workflow behavior.

Include:

* Multi-agent collaboration
* Long-horizon task completion
* Tool-use trajectories
* Web automation
* Software engineering workflows
* Debate or negotiation
* Human-agent collaboration

### Research Papers

Academic or technical papers about agent groups and flows.

Subsections may include:

* Surveys
* Multi-agent coordination
* Communication
* Planning and delegation
* Agent societies
* Evaluation
* Safety and governance

### Blog Posts

Technical writing that explains real patterns, architectures, or failures.

Include:

* Architecture breakdowns
* Production lessons
* Failure analyses
* Evaluation writeups
* Multi-agent design patterns

Exclude:

* Launch posts
* Shallow trend pieces
* Vendor marketing
* Prompt collections

### Case Studies

Real-world uses of multiple agents or agent flows.

Include:

* Production deployments
* Enterprise agent teams
* Software engineering agent teams
* Research agents
* Operations agents
* Support agents with escalation

### Examples and Templates

Runnable or forkable examples.

Include:

* Supervisor-worker examples
* Agent graph examples
* Crew/team templates
* Debate/review workflows
* Multi-agent software engineering examples
* Agent handoff examples

## 9. Phase Nine — Maintenance

### Weekly

* [ ] Review broken-link reports.
* [ ] Review new issues.
* [ ] Triage PRs.
* [ ] Remove obvious spam.

### Monthly

* [ ] Review stale PRs.
* [ ] Check archived projects.
* [ ] Re-run full link check.
* [ ] Review category balance.
* [ ] Move inactive but important resources to `docs/historical.md`.

### Quarterly

* [ ] Review scope.
* [ ] Prune weak entries.
* [ ] Refresh benchmarks.
* [ ] Refresh papers.
* [ ] Review whether new categories are needed.
* [ ] Update `docs/taxonomy.md`.
* [ ] Update `docs/rejected.md`.

## 10. Phase Ten — Public Launch

Launch only when:

* [ ] README has strong scope.
* [ ] Governance files are present.
* [ ] CI is passing.
* [ ] Link checker is passing.
* [ ] Custom Go checks are passing.
* [ ] At least 50 high-quality entries exist.
* [ ] At least 2 maintainers have reviewed the taxonomy.
* [ ] Contribution process is clear.
* [ ] Removal policy is clear.
* [ ] The list is not just a generic AI agents list.
