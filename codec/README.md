# substrait-go/codec

The codec module converts substrait-go's domain types to and from generated
Substrait protobuf messages. It is separate from the core module so core does not
depend on `substrait-protobuf`. Import codec only when you need the protobuf wire
format (issue #280).

    go get github.com/substrait-io/substrait-go/codec/v9

Core and codec are versioned independently. Update codec's core requirement in
`go.mod` when codec needs a newer core version, after that version is published.

For local development, the `replace` directive in `codec/go.mod` uses the core
module in this repository. The replacement stays in released source; Go ignores
it when codec is used as a dependency.

Release Please prepares a combined release PR for modules with releasable
changes. After that PR merges, it publishes their tags and GitHub releases.
Codec tags use the `codec/vX.Y.Z` prefix. See
[the release process](../RELEASING.md).
