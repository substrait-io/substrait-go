# substrait-go/codec

The codec module converts substrait-go's domain types to and from generated
Substrait protobuf messages. It is separate from the core module so core does not
depend on `substrait-protobuf`. Import codec only when you need the protobuf wire
format (issue #280).

    go get github.com/substrait-io/substrait-go/codec/v9

Core and codec are released with matching versions. A `codec/vX.Y.Z` release
declares a dependency on core `vX.Y.Z`.

For local development, the `replace` directive in `codec/go.mod` uses the core
module in this repository. The `scripts/release-codec.sh` script prepares a
release by updating the core dependency, removing the replacement, and creating
a separate release commit and codec tag. The release workflow pushes that tag.
