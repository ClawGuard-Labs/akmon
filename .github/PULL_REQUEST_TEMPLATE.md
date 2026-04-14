## Description

Brief description of what this PR does.

## Type of change

- [ ] Bug fix
- [ ] New feature
- [ ] Security fix
- [ ] Documentation
- [ ] Refactor / cleanup
- [ ] CI / build / tooling
- [ ] Breaking change
- [ ] Other (describe below)

## How to test

Steps or commands to verify the change (e.g. build, run monitor, trigger a rule).

```
# paste the commands you ran and what you observed
```

## Checklist

- [ ] Code follows project style (`make fmt`)
- [ ] `make build` passes
- [ ] `make lint` passes (or N/A — explain)
- [ ] `make test` passes where applicable
- [ ] Commits are signed off (`git commit -s`) per DCO
- [ ] Commit messages follow [Conventional Commits](https://www.conventionalcommits.org/)
- [ ] Documentation updated if user-facing behavior changed
- [ ] For detection-rule changes: opened against [onyx-templates](https://github.com/ClawGuard-Labs/onyx-templates) instead
- [ ] For eBPF changes: `make verify` passes; new kernel version requirements (if any) are noted below
- [ ] For breaking changes: listed under "Breaking changes" below with migration guidance

## Breaking changes

<!-- Delete this section if not applicable. Otherwise describe what breaks and how users should migrate. -->

## Related issues

Closes #
