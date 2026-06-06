# Contributing

Thank you for helping improve **Awesome AI Agent Factories**. This document covers submission format and discipline rules so your pull request matches list conventions before maintainers review scope.

See [README.md](README.md) for list scope and categories.

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
- If the resource is already listed under another name or section, do not submit a duplicate; comment on the existing entry's issue or open a discussion instead.

Maintainers will reject pull requests that introduce duplicate links.

## Alphabetical order

Within each README section, entries must be sorted in **alphabetical order by link text** (the text inside `[...]`).

- Compare case-insensitively unless two names differ only by case.
- Insert your entry in the correct position; do not append to the end of a section if that breaks order.
- When renaming link text, re-sort the affected section if needed.

## Before you open a pull request

1. Confirm your change adds **one resource** only.
2. Use the required entry format with the **exact resource name** as link text.
3. Write a **concise, factual** description that **ends with a period** and avoids marketing language.
4. Verify the URL is **not already present** in [README.md](README.md).
5. Place the entry in the correct section in **alphabetical order** by link text.

Maintainers review scope, link stability, and formatting manually. Following these rules reduces review churn and helps your pull request move forward.
