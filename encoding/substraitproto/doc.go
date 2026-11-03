// SPDX-License-Identifier: Apache-2.0

// Package substraitproto converts between substrait-go's domain types and the generated
// Substrait protobuf messages. It is a separate module so that serialization is
// optional.
//
// substraitproto imports the core. The core must never import substraitproto.
package substraitproto
