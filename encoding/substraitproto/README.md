# substrait-go/encoding/substraitproto

The substraitproto module converts substrait-go's domain types to and from generated
Substrait protobuf messages. It is separate from the core module so core does not
depend on `substrait-protobuf`. Import substraitproto only when you need the protobuf wire
format (issue #280).

    go get github.com/substrait-io/substrait-go/encoding/substraitproto

For local development, use a Go workspace to build core and this module from
the checkout. See [cross-module development](../../CONTRIBUTING.md#cross-module-development)
for setup and test commands.

Without a workspace, this module uses the published core version declared in
`go.mod`, which is also how CI builds and tests it.
