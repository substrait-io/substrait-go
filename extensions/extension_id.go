// SPDX-License-Identifier: Apache-2.0

package extensions

// ExtensionID identifies an object declared by a Substrait extension.
type ExtensionID struct {
	// URN identifies the extension file that declares the object.
	URN string
	// Name identifies the object within the extension file.
	Name string
}

// TypeID is the unique identifier for a Substrait extension type.
type TypeID ExtensionID

// FunctionID is the unique identifier for a Substrait extension function.
// For lookups, Name may be the simple function name. As a unique identifier,
// Name is the compound function name.
type FunctionID ExtensionID

// TypeVariationID is the unique identifier for a Substrait extension type variation.
type TypeVariationID ExtensionID
