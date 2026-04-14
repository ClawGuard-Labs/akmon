# Contributing to Onyx

Thank you for your interest in contributing. This document explains how to get set up, report issues, and submit changes.

## Code of Conduct

This project adheres to a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you agree to uphold it.

## How to contribute

- **Report bugs** — Open an issue with the bug report template. Include environment, steps to reproduce, and expected vs actual behavior.
- **Suggest features** — Open an issue with the feature request template. Describe the use case and, if you can, a proposed design.
- **Contribute code** — Fork the repo, create a branch, make your changes, run tests, and open a pull request.

## Development setup

### Prerequisites

- Linux with kernel **≥ 5.15** and BTF (`/sys/kernel/btf/vmlinux`) — **required for runtime**
- **Go 1.22+**
- Root or `CAP_BPF`, `CAP_PERFMON`, `CAP_NET_ADMIN` to run the monitor
- For eBPF recompilation: `clang`, `llvm-strip`, `bpftool`

macOS / Windows contributors: use the Docker toolchain below. The monitor itself cannot run outside Linux.

### Build and run

```bash
git clone https://github.com/ClawGuard-Labs/onyx.git
git clone https://github.com/ClawGuard-Labs/onyx-templates.git
cd onyx
make build
```

For running the monitor locally, clone **onyx-templates** under `onyx/onyx-templates` or pass `--behavioral-templates` / `--nuclei-templates` (see [README — Quick Start](README.md#quick-start)). For system install, `sudo make install` reads YAML from `../onyx-templates` by default (`TEMPLATES_SRC`).

See [README.md](README.md#quick-start) for full options and [Build Targets](README.md#build-targets).

### Building from macOS / Windows via Docker

Onyx is Linux-only at runtime, but the build toolchain runs inside a provided container image.

```bash
# Build the dev image once
docker build -t onyx-dev -f Dockerfile.dev .

# Persistent Go caches (faster subsequent builds)
docker volume create onyx-gomod
docker volume create onyx-gobuild

# Run any build/test command against the repo
docker run --rm \
    -v "$PWD:/src" \
    -v onyx-gomod:/go/pkg/mod \
    -v onyx-gobuild:/root/.cache/go-build \
    -w /src onyx-dev go build ./...

docker run --rm -v "$PWD:/src" -w /src onyx-dev make bpf
docker run --rm -v "$PWD:/src" -w /src onyx-dev bash -c 'cd ui && npm ci && npm run build'
```

### Useful make targets

- `make build` — Build Go binary and embed eBPF object
- `make bpf` — Recompile eBPF C (requires clang)
- `make gen-vmlinux` — Regenerate `bpf/vmlinux.h` from kernel BTF
- `make run` — Build and run as root
- `make fmt` — Format Go and C source
- `make lint` — Run `golangci-lint` (install separately; see [install guide](https://golangci-lint.run/welcome/install/))
- `make test` — Run template detection tests
- `make clean` — Remove build artifacts

## Pull request process

1. **Branch** — Create a branch from `main` (e.g. `feat/add-xyz`, `fix/issue-123`).
2. **Scope** — Keep one logical change per PR when possible.
3. **Description** — Use the PR template: what changed, why, and how to test.
4. **Tests** — Ensure the project builds and any existing tests pass (`make test`).
5. **Sign off** — This project uses the [Developer Certificate of Origin](https://developercertificate.org/). Every commit must be signed off: `git commit -s`. PRs from unsigned commits will be blocked by CI.
6. **Review** — A maintainer will review and may request changes. See `.github/CODEOWNERS` for who reviews which areas.

## Commit messages

We follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<optional-scope>): <short summary>

<body (optional)>

Signed-off-by: Your Name <your@email>
```

Common types:

- `feat:` — new user-visible capability
- `fix:` — bug fix
- `security:` — security-relevant change (CORS, auth, validation)
- `docs:` — documentation only
- `refactor:` — internal change with no behavior delta
- `test:` — tests only
- `chore:` / `build:` / `ci:` — tooling, dependencies, CI

Keep the summary under 72 characters, imperative mood ("add", not "adds" or "added"). The body explains the *why*; the code already shows the *what*.

## Style and conventions

- **Go** — `gofmt` on every file; run `make fmt`. Follow standard Go style. CI enforces `golangci-lint` — see `.golangci.yml`. Exported identifiers must have a godoc comment beginning with the identifier name.
- **eBPF / C** — match existing style in `bpf/`; run `make fmt` for `clang-format`. All BPF C files carry `SPDX-License-Identifier: MIT` at the top.
- **Build tags** — Go files under `internal/` and `cmd/` target Linux only. If you add a new file, start it with `//go:build linux`.
- **Templates** — YAML rules live in **[onyx-templates](https://github.com/ClawGuard-Labs/onyx-templates)**; follow [AUTHORING.md](https://github.com/ClawGuard-Labs/onyx-templates/blob/main/AUTHORING.md) and the layout there.

## Testing

- `make test` runs the template detection suite in `tests/`.
- Integration tests that need eBPF load must run on a Linux host with BTF. They skip automatically on macOS.
- When adding a new detection tag or matcher type, add a corresponding test case in `tests/`.

## Linting

`golangci-lint` is configured via `.golangci.yml` (gofmt, govet, errcheck, staticcheck, ineffassign, gosec). CI runs it on every PR.

```bash
# Install locally:
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin"

# Run:
make lint
```

## eBPF changes

When you touch `bpf/monitor.bpf.c`, `bpf/common.h`, or the loader in `internal/loader/`:

1. Regenerate `bpf/vmlinux.h` if your kernel differs materially from the checked-in one: `make gen-vmlinux`.
2. Rebuild and dry-run verify: `make bpf && make verify` — bpftool will load the program into `/sys/fs/bpf/onyx_verify` and unload it.
3. Test on at least one additional kernel minor version if possible (5.15 LTS and 6.x).
4. Call out any new required capabilities or kernel config in the PR description.

## Issue triage

- Maintainers triage new issues within ~5 business days.
- Labels: `bug`, `enhancement`, `security`, `good first issue`, `help wanted`, `needs-repro`, `needs-info`.
- Security-sensitive issues must go through [SECURITY.md](./SECURITY.md) — do not open public issues for them.

## Adding detection rules

Open PRs against **[onyx-templates](https://github.com/ClawGuard-Labs/onyx-templates)** (not this repo).

- **Behavioral rules** — add `.yaml` under `behavioral-templates/session|file|process|network/`. See [Detection Templates](README.md#detection-templates) and onyx-templates [AUTHORING.md](https://github.com/ClawGuard-Labs/onyx-templates/blob/main/AUTHORING.md).
- **Nuclei rules** — add Nuclei v3 HTTP templates under `nuclei-templates/ai-services/`.

## Recognition

Contributors are acknowledged in release notes and the GitHub contributor graph. If you prefer to remain anonymous, say so in your PR or issue.
