---
name: Bug report
about: Report a bug or unexpected behavior in Onyx
title: '[Bug] '
labels: bug
assignees: ''
---

## Description

A clear description of the bug.

## Environment

- **Distro / OS:** (e.g. Ubuntu 22.04, Debian 12, Fedora 40, RHEL 9)
- **`uname -a`:**
- **Kernel BTF present:** (`test -f /sys/kernel/btf/vmlinux && echo yes` → )
- **Go version:** (`go version`)
- **Onyx version / commit:** (from `onyx --version` or `git rev-parse HEAD`)
- **Install mode:** (from release tarball / `make install` / manual `./bin/onyx`)
- **Run mode:** (as root / with `setcap`; flags used, e.g. `--grouped`, `--sse`, `--ui`)
- **AI workload involved (if any):** (e.g. Ollama 0.3.12, Qdrant 1.11, vLLM, a custom Python agent)

## Steps to reproduce

1.
2.
3.

## Expected behavior

What you expected to happen.

## Actual behavior

What actually happened. Include:

- relevant log lines (`--log-level debug` is most useful)
- BPF verifier output if loading failed
- NDJSON output snippet if a rule misfired / didn't fire

## Additional context

Any other details, config snippets, templates involved, or a minimal reproducing script.
