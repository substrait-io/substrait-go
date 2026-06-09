package types

// ExtensionID identifies an object declared by a Substrait extension.
type ExtensionID struct {
	URN  string
	Name string
}

// ExtensionTypeID identifies a user-defined type declared by a Substrait extension.
type ExtensionTypeID ExtensionID

// ExtensionFunctionID identifies a function declared by a Substrait extension.
type ExtensionFunctionID ExtensionID

// ExtensionTypeVariationID identifies a type variation declared by a Substrait extension.
type ExtensionTypeVariationID ExtensionID
