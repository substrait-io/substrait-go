# substrait-go

Experimental Go bindings for [substrait](https://substrait.io)

[![release status](https://github.com/substrait-io/substrait-go/actions/workflows/release.yml/badge.svg)](https://github.com/substrait-io/substrait-go/actions/workflows/release.yml)
[![codecov](https://codecov.io/gh/substrait-io/substrait-go/branch/main/graph/badge.svg?token=7YXPNM3AMJ)](https://codecov.io/gh/substrait-io/substrait-go)
## Note:

This is work in progress still, things still to do:

- [ ] Expression parsing
- [x] Reading in extension yamls
- [x] CI building and testing the implementation
- [ ] Serialization/Deserialization of some expression types:
  - [x] IfThen
  - [x] SwitchExpression
  - [x] SingularOrList
  - [x] MultiOrList
  - [x] Cast
  - [x] Nested
  - [x] Subquery
- [ ] Serialization/Deserialization of Plan and Relations
  - [x] Plan
  - [x] PlanRel
  - [x] Rel
    - [x] ReadRel
    - [x] FilterRel
    - [x] FetchRel
    - [x] AggregateRel
    - [x] SortRel
    - [x] JoinRel
    - [x] ProjectRel
    - [x] SetRel
    - [x] ExtensionSingleRel
    - [x] ExtensionMultiRel
    - [x] ExtensionLeafRel
    - [x] CrossRel
    - [x] HashJoinRel
    - [x] MergeJoinRel
  - [ ] DdlRel
  - [ ] WriteRel
  - [ ] ExchangeRel
- [x] Plan Building helpers
  - [ ] ReadRel
    - [x] NamedScanReadRel
    - [x] VirtualTableReadRel
    - [ ] ExtensionTableReadRel
    - [ ] LocalFileReadRel
  - [x] FilterRel
  - [x] FetchRel
  - [x] AggregateRel
  - [x] SortRel
  - [x] JoinRel
  - [x] ProjectRel
  - [x] SetRel
  - [x] CrossRel
  - [ ] HashJoinRel
  - [ ] MergeJoinRel
  - [ ] DdlRel
  - [ ] WriteRel
  - [ ] ExchangeRel

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