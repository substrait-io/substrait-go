# substrait-go

Experimental Go bindings for [substrait](https://substrait.io)

## Note:

This is work in progress still, things still to do:

- [ ] Expression parsing
- [ ] Reading in extension yamls
- [ ] CI building and testing the implementation
- [ ] Serialization/Deserialization of some expression types:
  - [ ] IfThen
  - [ ] SwitchExpression
  - [ ] SingularOrList
  - [ ] MultiOrList
  - [ ] Cast
  - [ ] Nested
  - [ ] Subquery
- [ ] Plan Building helpers

As this is built out, you can expect refactors and other changes to the
structure of the package for the time being. **The API should not yet be
considered stable.**

## Generate from proto files

### Install buf

First ensure you have `buf` installed by following https://docs.buf.build/installation.

### Install go plugin

Run the following to install the Go plugin for protobuf:

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Ensure that your GOPATH is on your path:

```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

### Run go generate

As long as buf and the Go protobuf plugin are installed, you can 
simply run `go generate` to generate the updated `.pb.go` files. It
will generate them by referencing the primary substrait-io repository.

You can then commit the updated files.