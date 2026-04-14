# Security Policy

Akmon is a kernel-level security sensor. Vulnerabilities in it directly affect the security posture of the hosts it runs on, so we take reports seriously and treat them privately until a fix is available.

## Supported versions

| Version | Status |
|---------|--------|
| `main` | Actively supported. All security fixes land here first. |
| Latest minor release | Supported — security fixes backported. |
| Older releases | Not supported. Please upgrade. |

## Reporting a vulnerability

**Do not open a public GitHub issue for security problems.**

Preferred channel:

- Open a private [GitHub Security Advisory](https://github.com/ClawGuard-Labs/akmon/security/advisories/new) on this repository.

What to include:

- A clear description of the issue and its impact.
- Steps to reproduce (PoC code, exact `akmon` flags, kernel version, distro).
- Any relevant logs, crash output, or BPF verifier output.
- Whether the issue is already public anywhere (blog post, tweet, etc.).
- How you'd like to be credited (name/handle, or anonymous).

## What to expect

| Stage | Target |
|-------|--------|
| Acknowledgement of your report | within **3 business days** |
| Initial triage and severity assessment | within **7 business days** |
| Fix in `main` (for confirmed, in-scope issues) | within **30 days** for high/critical; longer for lower severity |
| Coordinated disclosure | we'll agree a disclosure date with you once a fix is ready |

For critical issues (RCE, privilege escalation, sandbox escape) we'll prioritize and communicate more aggressively.

## Scope

**In scope:**

- The Akmon monitor binary and its eBPF programs (`cmd/`, `internal/`, `bpf/`).
- The dashboard HTTP API (`internal/graphapi/`).
- Build and release artifacts published on this repository.
- Documentation that could lead users into an insecure configuration.

**Out of scope:**

- Vulnerabilities in upstream dependencies (please report to the upstream project). If the way we use a dependency is the issue, that *is* in scope.
- Findings that require an already-root attacker on the host (Akmon runs as root; the threat model starts from unprivileged attacker).
- Denial-of-service against the dashboard API achievable only from the host itself (the API is intended for localhost use).
- Issues in the companion [akmon-templates](https://github.com/ClawGuard-Labs/akmon-templates) repository — please open an advisory there instead.

## Credit

Unless you ask otherwise, we credit reporters in the advisory and release notes. Let us know in your report if you'd prefer to stay anonymous.

## PGP / encrypted contact

If a GitHub Security Advisory is not possible for your situation, open an issue titled "Request private security contact" with no details, and a maintainer will reach out to arrange an encrypted channel.
