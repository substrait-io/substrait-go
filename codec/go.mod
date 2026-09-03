// SPDX-License-Identifier: Apache-2.0

module github.com/substrait-io/substrait-go/codec

go 1.23.0

toolchain go1.24.4

require (
	github.com/google/go-cmp v0.7.0
	github.com/stretchr/testify v1.10.0
	github.com/substrait-io/substrait-go/v9 v9.0.0
	github.com/substrait-io/substrait-protobuf/go v0.85.0
	google.golang.org/protobuf v1.36.6
)

require (
	cloud.google.com/go v0.121.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/creasty/defaults v1.8.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/goccy/go-yaml v1.17.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/substrait-io/substrait v0.87.0 // indirect
	golang.org/x/exp v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// core is developed in the same repo and unreleased; resolve it locally.
replace github.com/substrait-io/substrait-go/v9 => ../
