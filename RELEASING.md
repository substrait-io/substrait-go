# Releases

Core and `encoding/substraitproto` are versioned independently. Release Please
prepares a separate release PR for each module with releasable changes. The `main`
branch publishes stable releases; `alpha` publishes prereleases.

## Prepare and publish

1. Merge changes using [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
   Release Please uses the commit messages and changed file paths to propose a
   version for each affected module.
   Changes only under `encoding/substraitproto/` do not bump core.
2. Review the proposed versions and changelogs. For a major release, update the
   affected module path and imports before merging. Release Please does not make
   those Go source changes or validate module-path majors automatically.
3. Merge a module's release PR when it is ready. Release Please creates its tag
   and GitHub release. Core tags use `vX.Y.Z`; substraitproto tags use
   `encoding/substraitproto/vX.Y.Z`. Its first stable release is `v0.1.0`.

There is currently no CI or publication-time check that the proposed release
major matches the module path in `go.mod`. Verify this manually before merging a
release PR. For example, `v10.0.0` requires a module path ending in `/v10`, not
`/v9`. GitHub can publish a mismatched tag, but Go consumers will reject that
version. A fix is tracked in
[#133](https://github.com/substrait-io/substrait-go/issues/133).

Release PRs update the affected module changelogs and the channel's version
manifest. Publication does not modify files or create release-only commits.
Publishing multiple modules is not atomic; inspect the existing tags and
releases before retrying a failed workflow. Never delete or move published tags.

## Alpha releases

Develop prereleases on `alpha`. Merging an alpha release PR creates an alpha tag
and marks its GitHub release as a prerelease. Core tags use
`vX.Y.Z-alpha.N`; substraitproto tags use
`encoding/substraitproto/vX.Y.Z-alpha.N`. Release Please tracks
these versions in `.release-please-manifest-alpha.json`, separately from stable
releases, and increments the alpha counter for subsequent releases in a series.

To start a new series, add a `Release-As` footer to a releasable commit affecting
the module. For example, after stable `v9.1.0`, a core feature commit can select
the next series with:

```text
feat(expr): add a new expression

Release-As: 9.2.0-alpha.0
```

Use separate PRs when core and substraitproto need different target versions.
Review alpha release PRs to ensure their versions are newer than existing tags.

When ready for a stable release, bring the implementation changes into `main`
through a normal squash-merged PR, without alpha-only `Release-As` footers.
Keep the stable manifest unchanged; do not copy alpha versions into it. Merge
the resulting stable release PR to publish. Alpha tags remain prereleases and
are never moved or renamed.

## Updating substraitproto's core dependency

Update the requirement in `encoding/substraitproto/go.mod` in a normal PR when
substraitproto needs a newer core version. Publish that core version first. For
example, core can release `v9.1.0`, then a dependency update can ship as
`encoding/substraitproto/v0.1.1` if it does not break substraitproto's API. The two
version numbers do not need to match.

The same ordering applies to alphas. Publish a core alpha before updating
substraitproto to require it. Before publishing stable substraitproto, update its
requirement to a compatible published stable core version.

The repository's `go.work` uses core from the checkout for local development.
CI sets `GOWORK=off` to test each module against its declared dependencies. Check
that the declared core version contains the APIs substraitproto needs. See
[cross-module development](CONTRIBUTING.md#cross-module-development).
