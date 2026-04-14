# Changelog

All notable changes to this project are documented here.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- `Dockerfile.dev` — reproducible Linux toolchain (Go, clang, libbpf, Node) for macOS/Windows contributors.
- `SECURITY.md`, `SUPPORT.md`, `CHANGELOG.md`, `.github/CODEOWNERS`, `.github/dependabot.yml`.
- GitHub Actions CI: `go-build-test`, `go-lint`, `bpf-build`, `ui-build` jobs.
- GitHub Actions release workflow with `goreleaser` producing `akmon_linux_{amd64,arm64}.tar.gz` on tag.
- `.golangci.yml` enforcing `gofmt`, `govet`, `errcheck`, `staticcheck`, `ineffassign`, `gosec`.
- `--cors-origin` flag (repeatable) on the dashboard API.
- Package-level doc comments on every `internal/*` package.
- `//go:build linux` constraint on every Go file under `cmd/` and `internal/`.
- `SPDX-License-Identifier: MIT` header on BPF C sources.
- FAQ, Community, and Download sections in the README.
- DCO sign-off and Conventional Commits guidance in `CONTRIBUTING.md`.

### Changed
- **BREAKING:** Project renamed from `ClawSec` to `Akmon`. Module path is now `github.com/ClawGuard-Labs/akmon`. Binary is `akmon`. Install paths are `/usr/local/bin/akmon`, `/usr/lib/akmon/`, `/etc/akmon/`, `/var/log/akmon/`, `/etc/systemd/system/akmon.service`, `/etc/logrotate.d/akmon`.
- Companion templates repository is now [`ClawGuard-Labs/akmon-templates`](https://github.com/ClawGuard-Labs/akmon-templates) (previously `clawsec-templates`).
- Dashboard API `Access-Control-Allow-Origin` now defaults to `http://localhost:9090` and `http://127.0.0.1:9090` instead of `*`. Extend via `--cors-origin`.
- `internal/graphapi` shutdown now derives its context from the parent, with a bounded timeout, so host cancellation propagates correctly.
- `config.yaml`: added missing `category: runtime` to the `cargo` and `rust-ai` process entries.
- LICENSE year updated to 2025–2026 and copyright holder to "The Akmon Contributors".

### Fixed
- `README.md` security-advisory link pointed to a non-existent org; now points to `github.com/ClawGuard-Labs/akmon/security/advisories/new`.
- `scripts/gen_vmlinux.sh` bpftool discovery no longer relies on unquoted `ls` output (word-splitting hazard).
- `Makefile` `verify` target now uses a `trap` to clean up `/sys/fs/bpf/akmon_verify` even when bpftool is interrupted.
- `tests/helpers_test.go` uses `t.Fatalf` instead of `panic` on log-file creation failure.
- Several places in `internal/output` and `internal/loader` previously discarded write and `Close` errors; they now log at `warn` level.

### Removed
- Stale `TODO.md` reference in `internal/provenance/tracker.go`.

## [0.1.0] — 2026-04-14

Initial public release of Akmon as an open-source project.

- Kernel-level behavioral monitoring for AI/ML workloads on Linux using eBPF (CO-RE, BTF).
- YAML-based detection engine with process, file, network, session, and chain matchers.
- Nuclei v3 integration for active scanning of local AI services (Qdrant, Ollama, vLLM, ChromaDB, Weaviate, Milvus, …) when a process connects to them.
- Session correlation with process-tree reconstruction.
- NDJSON, grouped JSON, and live SSE output modes.
- Optional React dashboard (Vite) embedded in the binary via `go:embed`.
- `systemd` unit, logrotate config, dependency checker, and `make install`/`make uninstall` targets.
