# Releases

Core and codec share release versions. A breaking change in either module
requires updating both module paths and the affected imports before release.

## Prepare and publish

1. Merge changes using Conventional Commits. Release Please opens or updates one
   release PR targeting `main`.
2. Review the proposed version, changelog, and codec's updated core requirement.
   The release-version CI check validates both module majors and that requirement.
   If the major changes, update both module paths and affected imports before
   merging the release PR.
3. Merge the release PR. The workflow publishes core as `vX.Y.Z` and adds
   `codec/vX.Y.Z` at the same commit. Publication does not modify files or create
   release-only commits. Codec has a Go module tag, not a separate GitHub release.

The release PR updates `CHANGELOG.md`, `.release-please-manifest.json`, and the
core requirement marked with `// x-release-please-version` in `codec/go.mod`.
The local `replace` stays in that file so tests can use core from this checkout.
Go ignores dependency modules' replacement directives for consumers.

Between merging the release PR and publishing core, codec temporarily requires
an unpublished core version. The local replacement keeps development and CI
working, but consumers of that unreleased commit may fail until core is published.

## Repository setup

Configure a personal access token as the `RELEASE_PLEASE_TOKEN` repository secret,
with permission to create release PRs, labels, and releases. This allows generated release PRs to trigger CI. The built-in
`GITHUB_TOKEN` does not normally trigger workflows for resources it creates. See
[the action's authentication guidance](https://github.com/googleapis/release-please-action#github-credentials).
Require the `Release Version` check and normal build/test checks before merging
release PRs.

The manifest is bootstrapped from the latest stable core release, `8.1.1`. Update
that starting value if another stable release is published before this migration
merges. Release Please maintains it after the first release PR. Existing alpha
tags are not changed.

## Retry codec tagging

Core publication and codec tagging are not atomic. If core publishes but codec
tagging fails, rerunning Release Please may not retry tagging after the core
release is marked complete.

To create a missing codec tag for an existing core release:

```sh
git fetch origin --tags
core_tag=v9.1.0
git tag "codec/$core_tag" "$core_tag^{commit}"
git push origin "refs/tags/codec/$core_tag"
```

If the codec tag already exists, verify it points to the core release commit.
Never delete or move a published tag.
