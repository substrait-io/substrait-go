// SPDX-License-Identifier: Apache-2.0

// Package wire converts between substrait-go's domain types and the generated
// Substrait protobuf messages.
//
// This folder is temporary. Issue #280 moves substrait-protobuf out of core so
// a consumer that only builds and inspects plans no longer drags the generated
// types into its dependency graph. The conversions collect here first, in-tree,
// so each type can move over one at a time. Once every proto conversion lives
// here, the whole package moves into its own codec module and leaves core.
package wire
