// SPDX-License-Identifier: Apache-2.0

module github.com/substrait-io/substrait-go/codec

go 1.23.0

toolchain go1.24.4

require (
	github.com/google/go-cmp v0.7.0
	github.com/substrait-io/substrait-go/v8 v8.0.0-00010101000000-000000000000
	github.com/substrait-io/substrait-protobuf/go v0.85.0
	google.golang.org/protobuf v1.36.6
)

require cloud.google.com/go v0.121.0 // indirect

// The codec builds against the working tree rather than a published core, so a change to a core
// type does not need a release first. The require above is deliberately unresolvable: the codec
// cannot be published until a core carrying the domain types ships.
replace github.com/substrait-io/substrait-go/v8 => ../
