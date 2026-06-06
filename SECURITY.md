# Security Policy

This document explains how **Awesome AI Agent Factories** handles security concerns about resources linked from the curated list. It applies to readers, contributors, and maintainers.

This repository is a manually maintained awesome list. We do not host third-party code, packages, or binaries. Linked resources retain their own licenses and security posture. See [LICENSE](LICENSE) for list-content licensing.

## Scope

This policy covers:

- Destinations linked from [README.md](README.md) entries
- URLs proposed in open pull requests before they are merged
- Suspicious install instructions, binaries, or dependency guidance associated with listed resources

It does not cover vulnerabilities in this repository's own markdown or GitHub configuration. Use the reporting channels below for those cases as well; maintainers will route reports appropriately.

## Threats we watch for

### Malicious links

A listed URL may point to phishing pages, malware downloads, credential harvesting, or other harmful destinations.

Maintainers treat credible reports seriously. When a link appears malicious:

- The entry may be removed promptly or the pull request rejected before merge.
- Maintainers may add a note in the issue or pull request without repeating a dangerous URL when a safer reference is available.
- Repeated or deliberate submission of harmful links may be treated as a conduct issue under [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

We do not run automated malware scanning today. Link review is manual during pull request review and when reports are received.

### Typosquatting and confusing names

Typosquatting includes URLs or project names that closely resemble well-known projects to trick readers (for example, swapped characters, misleading domains, or copycat repositories).

Before submitting or approving an entry:

- Verify the URL matches the official project, paper, or article site.
- Compare the repository owner, domain, and package name to the resource you intend to cite.
- Be cautious with recently created repositories that mirror popular names.

Maintainers may reject or remove entries when the destination is likely to confuse readers or impersonate another project.

### Supply-chain concerns

Many listed resources are installable software. Readers may run install commands or add dependencies based on list entries.

Be alert to:

- Package or repository hijacking after maintainer turnover
- Install instructions that pipe remote scripts to a shell without verification
- Unvetted binaries, container images, or dependency confusion across ecosystems
- Sudden changes in package ownership, release signing, or default install paths

This list does not vet every dependency graph. Contributors should link to canonical sources, and readers should follow normal supply-chain hygiene before installing third-party software.

If you believe a listed resource's distribution channel has been compromised, report it even if the README description is still accurate.

### Abandoned or compromised projects

Not every inactive project is a security problem. We still act when a resource is unsafe to recommend.

Maintainers may remove or relocate an entry when:

- The project is **compromised** — confirmed account takeover, malicious releases, or credible reports of harmful behavior
- The project is **abandoned and risky** — archived with no maintenance, broken security posture, or known unfixed critical issues that make the link misleading
- The destination is **dead or unstable** — persistent broken links or repeated redirects to unrelated sites
- The description no longer matches the resource — scope drift, marketing takeover, or content that no longer addresses agent-factory topics

Historically important but inactive resources may be moved to archival documentation when that process exists. Until then, maintainers may remove unsafe destinations and note the reason in the relevant issue or pull request.

## What this repository does not guarantee

- We do not continuously monitor every linked site for compromise.
- We do not certify that listed frameworks, packages, or papers are safe to run.
- We do not operate a bug bounty or paid security program.

Automated link checking and repository CI enforcement are planned for later phases. Until those workflows exist, rely on maintainer review and community reports.

## Reporting concerns

Report security issues through one of these channels:

1. **GitHub Security Advisories (preferred for sensitive reports)** — [Open a private vulnerability report](https://github.com/portpowered/awesome-agent-factories/security/advisories/new) for this repository when disclosure should not be public yet. Use this for compromised list entries, malicious proposed links, or security issues in repository automation once it exists.

2. **GitHub issue** — Open a [new issue](https://github.com/portpowered/awesome-agent-factories/issues/new) for non-sensitive reports such as suspicious linked resources, typosquatting, or outdated entries. Include the README section and entry text when possible.

3. **Contact a maintainer directly** — Message a repository maintainer on GitHub:
   - [@portpowered](https://github.com/portpowered)
   - [@AndreasAbdi](https://github.com/AndreasAbdi)

Maintainer review responsibilities and contact policy are described in [MAINTAINERS.md](MAINTAINERS.md).

### What to include

- The README section and entry text, or a pull request link
- The URL you consider harmful, confusing, or compromised
- What you observed (malware, typosquatting, hijacked package, dead link, etc.)
- Links to public advisories, issues, or news reports when available
- Whether you need a private response

## Maintainer response

Maintainers review reports in good faith and as availability allows. There is no 24/7 security desk.

Typical responses include:

- Removing or rejecting the affected entry
- Requesting a corrected canonical URL from the contributor
- Asking for more evidence before action
- Locking or closing abusive submissions

We aim to acknowledge credible security reports within a reasonable time. Response time depends on maintainer availability.

## Related policies

- [CONTRIBUTING.md](CONTRIBUTING.md) — submission rules, scope, and duplicate-link discipline
- [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) — behavior expectations in issues and pull requests
- [MAINTAINERS.md](MAINTAINERS.md) — review, merge, removal, and contact policy
