# Releases

Core and codec are versioned independently. Release Please prepares a separate
release PR for each module with releasable changes.

## Prepare and publish

1. Merge changes using Conventional Commits. Release Please uses the commit
   messages and changed file paths to propose a version for each affected module.
   Changes only under `codec/` do not bump core.
2. Review the proposed versions and changelogs. For a major release, update the
   affected module path and imports before merging. Release Please does not make
   those Go source changes or validate module-path majors automatically.
3. Merge a module's release PR when it is ready. Release Please creates its tag
   and GitHub release. Core tags use `vX.Y.Z`; codec tags use `codec/vX.Y.Z`.

Release PRs update the affected module changelogs and
`.release-please-manifest.json`. Publication does not modify files or create
release-only commits. Publishing multiple modules is not atomic; inspect the
existing tags and releases before retrying a failed workflow. Never delete or
move published tags.

## Updating codec's core dependency

Update the requirement in `codec/go.mod` in a normal PR when codec needs a newer
core version. Publish that core version first. For example, core can release
`v9.1.0`, then a codec dependency update can ship as `codec/v9.0.1` if it does not
break codec's API. The two version numbers do not need to match.

The local `replace` in `codec/go.mod` lets development and CI use core from the
checkout. Go ignores that replacement when codec is used as a dependency, so
check that the declared core version contains the APIs codec needs.
