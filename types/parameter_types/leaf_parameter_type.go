package parameter_types

// LeafParameter represents a parameter type
// parameter can of concrete (38) or abstract type (P)
// or another parameterized type like VARCHAR<"L1">
type LeafParameter interface {
	// IsCompatible is type compatible with other
	// compatible is other can be used in place of this type
	IsCompatible(other LeafParameter) bool
	String() string
}

// AbstractParameterType represents a parameter type which is abstract.
// it can be a leaf parameter (LeafIntParamAbstractType)
// or another abstract type like "DECIMAL<P, S>"
type AbstractParameterType interface {
	GetAbstractParamName() string
}
