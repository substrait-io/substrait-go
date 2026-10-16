# Releases

Core and codec are versioned independently. Release Please prepares one combined
release PR for modules with releasable changes.

## Prepare and publish

1. Merge changes using Conventional Commits. Release Please uses the commit
   messages and changed file paths to propose a version for each affected module.
   Changes only under `codec/` do not bump core.
2. Review the proposed versions and changelogs. For a major release, update the
   affected module path and imports before merging. Release Please does not make
   those Go source changes or validate module-path majors automatically.
3. Merge the release PR. Release Please creates tags and GitHub releases for the
   affected modules. Core tags use `vX.Y.Z`; codec tags use `codec/vX.Y.Z`.

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

## Repository setup

Configure a personal access token as the `RELEASE_PLEASE_TOKEN` repository secret,
with permission to create release PRs, labels, and releases. This allows generated
release PRs to trigger CI. The built-in `GITHUB_TOKEN` does not normally trigger
workflows for resources it creates. See
[the action's authentication guidance](https://github.com/googleapis/release-please-action#github-credentials).

The manifest and bootstrap commit start at core's latest stable release,
`v8.1.1`. Codec has no stable release yet, so its configuration sets the initial
version to `9.0.0` to match its existing module path. The first combined release
PR proposes `v9.0.0` and `codec/v9.0.0`. Codec already requires core `v9.0.0`, so
that initial publication depends on core publishing successfully too.

Update the bootstrap values if stable releases are published before this
migration merges. Release Please maintains the manifest after initialization.
Existing alpha tags are not changed.
