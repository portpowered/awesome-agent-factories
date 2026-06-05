# Factory Overview

This factory coordinates autonomous work for the AI model reference website.
The ideafy workstation is the meta-planner. It chooses phase-scoped batches,
submits ideas, and records progress. Executors implement PRD stories in
worktrees. Review gates the resulting PRs.

## Read First

Before submitting work, read:

* `factory/factory.json`
* `factory/workstations/ideafy/AGENTS.md`
* `docs/internal/customer-ask.md`
* `docs/internal/checklist.md`
* `docs/internal/progress.txt`
* `docs/documentation-site-pages-needed.md`
* `you docs agents`
* `you docs batch-inputs`

## Phase Control

Current phase authorization lives in:

```txt
docs/internal/customer-ask.md
```

The meta-planner may dry-run batches during planning. It must not submit a real
batch unless `customer-ask.md` sets `realSubmissionAuthorized: true` or the
customer explicitly authorizes submission in the current conversation.

Phase work is review-gated through Phase 10. After Phase 10, long-tail backfill
may run mostly autonomously in small batches, still with batch summaries and
review.

## Work Types

Configured work types:

```txt
thoughts       meta-planner loopback work
idea           product/implementation idea submitted by ideafy
plan           PRD planning output from an idea
task           executor/review implementation work
cron-triggers  runtime trigger type
```

Use `idea`, singular, for implementation proposals.
Use `thoughts`, plural, for ideafy loopback.

## Workstation Flow

```txt
thoughts:init -> ideafy -> thoughts:complete

idea:init -> plan -> idea:to-complete + plan:init
plan:init -> setup-workspace -> plan:complete + task:init
task:init -> process -> task:in-review
task:in-review -> review -> task:to-complete
idea:to-complete + task:to-complete with the same name -> consume
```

## Batch Submission

Use the canonical `FACTORY_REQUEST_BATCH` shape from `you docs batch-inputs`.

For a running factory, prefer:

```sh
you submit batch <path>
```

Always dry-run first:

```sh
you submit batch --dry-run <path>
```

For watched-folder operator ingress, use:

```txt
factory/inputs/BATCH/default/<request_id>.json
```

The checked-in example is:

```txt
factory/docs/batch-input-example.json
```

## State Inspection

Use:

```sh
you work list
you session list
```

`you work list` shows durable work state. `you session list` shows active or
recent runtime sessions. Check both before deciding that work is stuck or before
submitting a new batch.

## Repair

Use:

```sh
you work move
```

only for deliberate workflow repair. Record every manual move in
`docs/internal/progress.txt` with the work item, old state, new state, reason,
and expected next workstation. Do not use work moves to skip implementation,
review, or validation.

## Local State Files

```txt
docs/internal/customer-ask.md  current phase and submission authorization
docs/internal/checklist.md     high-level phase and customer ask tracking
docs/internal/progress.txt     append-only meta-planner progress log
```
