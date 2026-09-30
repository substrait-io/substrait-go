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

// FunctionID is the unique identifier for a Substrait extension function variant.
type FunctionID struct {
	// URN identifies the extension file that declares the function.
	URN string
	// Signature is the registered compound function variant name.
	Signature string
}

// TypeVariationID is the unique identifier for a Substrait extension type variation.
type TypeVariationID ExtensionID
