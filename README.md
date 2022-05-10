# substrait-go

Experimental Go bindings for [substrait](https://substrait.io)

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