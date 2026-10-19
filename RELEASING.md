# Releases

Core and codec are versioned independently. Release Please prepares a separate
release PR for each module with releasable changes. The `main` branch publishes
stable releases; `alpha` publishes prereleases.

## Prepare and publish

1. Merge changes using Conventional Commits. Release Please uses the commit
   messages and changed file paths to propose a version for each affected module.
   Changes only under `codec/` do not bump core.
2. Review the proposed versions and changelogs. For a major release, update the
   affected module path and imports before merging. Release Please does not make
   those Go source changes or validate module-path majors automatically.
3. Merge a module's release PR when it is ready. Release Please creates its tag
   and GitHub release. Core tags use `vX.Y.Z`; codec tags use `codec/vX.Y.Z`.

Release PRs update the affected module changelogs and the channel's version
manifest. Publication does not modify files or create release-only commits.
Publishing multiple modules is not atomic; inspect the existing tags and
releases before retrying a failed workflow. Never delete or move published tags.

## Alpha releases

Develop prereleases on `alpha`. Merging an alpha release PR creates an alpha tag
and marks its GitHub release as a prerelease. Core tags use
`vX.Y.Z-alpha.N`; codec tags use `codec/vX.Y.Z-alpha.N`. Release Please tracks
these versions in `.release-please-manifest-alpha.json`, separately from stable
releases, and increments the alpha counter for subsequent releases in a series.

To start a new series, add a `Release-As` footer to a releasable commit affecting
the module. For example, after stable `v9.1.0`, a core feature commit can select
the next series with:

```text
feat(expr): add a new expression

Release-As: 9.2.0-alpha.0
```

Use separate PRs when core and codec need different target versions. Review
alpha release PRs to ensure their versions are newer than existing tags.

When ready for a stable release, bring the implementation changes into `main`
through a normal squash-merged PR, without alpha-only `Release-As` footers.
Keep the stable manifest unchanged; do not copy alpha versions into it. Merge
the resulting stable release PR to publish. Alpha tags remain prereleases and
are never moved or renamed.

## Updating codec's core dependency

Update the requirement in `codec/go.mod` in a normal PR when codec needs a newer
core version. Publish that core version first. For example, core can release
`v9.1.0`, then a codec dependency update can ship as `codec/v9.0.1` if it does not
break codec's API. The two version numbers do not need to match.

The same ordering applies to alphas. Publish a core alpha before updating codec
to require it. Before publishing stable codec, update its requirement to a
compatible published stable core version.

The local `replace` in `codec/go.mod` lets development and CI use core from the
checkout. Go ignores that replacement when codec is used as a dependency, so
check that the declared core version contains the APIs codec needs.
