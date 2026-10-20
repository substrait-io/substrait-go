# substrait-go/encoding/substraitproto

The substraitproto module converts substrait-go's domain types to and from generated
Substrait protobuf messages. It is separate from the core module so core does not
depend on `substrait-protobuf`. Import substraitproto only when you need the protobuf wire
format (issue #280).

    go get github.com/substrait-io/substrait-go/encoding/substraitproto

The repository's `go.work` builds core and this module from the checkout during
local development. See [cross-module development](../../CONTRIBUTING.md#cross-module-development)
for build and test commands.

CI sets `GOWORK=off` to build and test against the published core version
declared in `go.mod`.
