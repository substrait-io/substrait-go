// SPDX-License-Identifier: Apache-2.0

module github.com/substrait-io/substrait-go/v9

go 1.23.0

toolchain go1.24.4

require (
	cloud.google.com/go v0.121.0
	github.com/antlr4-go/antlr/v4 v4.13.1
	github.com/cockroachdb/apd/v3 v3.2.1
	github.com/creasty/defaults v1.8.0
	github.com/goccy/go-yaml v1.17.1
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.10.0
	github.com/substrait-io/substrait-packaging/go/substrait-antlr v0.0.0-00010101000000-000000000000
	github.com/substrait-io/substrait-packaging/go/substrait-extensions v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.36.6
)

// DO NOT MERGE: the go/substrait-* pre-release branches only exist on the
// vbarua/substrait-packaging fork, so point the canonical module paths at it.
replace (
	github.com/substrait-io/substrait-packaging/go/substrait-antlr => github.com/vbarua/substrait-packaging/go/substrait-antlr v0.0.0-20261124184703-e1fabd7846db
	github.com/substrait-io/substrait-packaging/go/substrait-extensions => github.com/vbarua/substrait-packaging/go/substrait-extensions v0.0.0-20261124184638-88b3a4ff5851
	github.com/substrait-io/substrait-packaging/go/substrait-protobuf => github.com/vbarua/substrait-packaging/go/substrait-protobuf v0.0.0-20261124184706-c9daca29e656
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	golang.org/x/exp v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
