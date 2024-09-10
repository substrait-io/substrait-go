package parameter_types

import "fmt"

// LeafIntParamConcreteType represents a single integer concrete parameter for a concrete type
// Example: VARCHAR(6) -> 6 is an LeafIntParamConcreteType
// DECIMAL<P, 0> --> 0 Is an LeafIntParamConcreteType but P not
type LeafIntParamConcreteType int32

func (m LeafIntParamConcreteType) IsCompatible(o LeafParameter) bool {
	if t, ok := o.(LeafIntParamConcreteType); ok {
		return t == m
	}
	return false
}

func (m LeafIntParamConcreteType) String() string {
	return fmt.Sprintf("%d", m)
}

func (m LeafIntParamConcreteType) ToProtoVal() int32 {
	return int32(m)
}
