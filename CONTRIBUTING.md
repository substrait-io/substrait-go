# Contributing to Substrait Go

Welcome! This page provides some orientation and recommendations on how to get
the best results when engaging with the community.

## Contributor License Agreement

Substrait requires all contributors to sign the
[Contributor License Agreement (CLA)](https://cla-assistant.io/substrait-io/substrait)
before their contributions can be merged. A GitHub app checks this on every pull
request and guides new contributors through signing it.

## The specification is the source of truth

Substrait Go is an implementation of the
[Substrait specification](https://substrait.io/); it does not define Substrait
semantics. Review behavioral changes against the spec text and the `.proto`
comments in
[`substrait-io/substrait`](https://github.com/substrait-io/substrait).

Where the spec is genuinely unclear, don't settle it here — raise a clarification
issue in
[`substrait-io/substrait`](https://github.com/substrait-io/substrait/issues) or
bring it to the [community](https://substrait.io/community/) channels rather than
encoding a guess, and record the open question in the pull request so the
assumption stays reviewable.

## Commit Conventions

Substrait Go follows
[conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit
message structure, with an optional scope naming the affected package — for
example `fix(expr): initialize MapLiteral KeyValue entries before use`. Append `!`
after the type or scope for a breaking change (`refactor(extensions)!: …`).

Releases are cut by
[`go-semantic-release`](https://github.com/Nightapes/go-semantic-release), which
derives the next version and the release notes from these commit messages, so the
type and the `!` marker directly determine what gets released. Pull requests are
squash-merged, so please ensure that your PR title and initial comment together
form a valid commit message.

## Building and testing

```sh
go build ./...
go test ./...
```

CI builds and tests on Linux, Windows and macOS, and reports coverage from the
Linux job to Codecov.

## Linting

Linting is done with [`golangci-lint`](https://golangci-lint.run) using the
configuration in [`.golangci.yml`](.golangci.yml). The `Code Linting` CI job is
authoritative for whether a change is clean:

```sh
golangci-lint run
```

`.golangci.yml` is a `version: "2"` configuration, so it requires golangci-lint
2.x; CI currently runs 2.1.6.

## License headers

Go files carry an SPDX license header as their first line:

```go
// SPDX-License-Identifier: Apache-2.0
```

Please include one in new files. This is a convention rather than something CI
currently enforces.
