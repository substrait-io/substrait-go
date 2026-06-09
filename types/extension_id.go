package types

// ExtensionID identifies an object declared by a Substrait extension.
type ExtensionID struct {
	// URN identifies the extension file that declares the object.
	URN string
	// Name identifies the object within the extension file.
	Name string
}

// ExtensionTypeID identifies a user-defined type declared by a Substrait extension.
type ExtensionTypeID ExtensionID

// ExtensionFunctionID identifies a function declared by a Substrait extension.
type ExtensionFunctionID ExtensionID

// ExtensionTypeVariationID identifies a type variation declared by a Substrait extension.
type ExtensionTypeVariationID ExtensionID
