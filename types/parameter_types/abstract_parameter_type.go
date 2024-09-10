package parameter_types

// LeafIntParamAbstractType represents an integer parameter for a parameterized type
// Example: VARCHAR(L1) -> L1 is an LeafIntParamAbstractType
// DECIMAL<P, 0> --> P Is an LeafIntParamAbstractType
type LeafIntParamAbstractType string

func (m LeafIntParamAbstractType) IsCompatible(o LeafParameter) bool {
	switch o.(type) {
	case LeafIntParamAbstractType, LeafIntParamConcreteType:
		return true
	default:
		return false
	}
}

func (m LeafIntParamAbstractType) String() string {
	return string(m)
}

func (m LeafIntParamAbstractType) GetAbstractParamName() string {
	return string(m)
}
