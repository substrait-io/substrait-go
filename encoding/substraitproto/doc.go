// SPDX-License-Identifier: Apache-2.0

// Package codec converts between substrait-go's domain types and the generated
// Substrait protobuf messages. It is a separate module so that serialization is
// optional.
//
// codec imports the core. The core must never import codec.
package codec
