// SPDX-License-Identifier: Apache-2.0

package leaf_parameters

import "fmt"

// ConcreteIntParam represents a single integer concrete parameter for a concrete type
// Example: VARCHAR(6) -> 6 is an ConcreteIntParam
// DECIMAL<P, 0> --> 0 Is an ConcreteIntParam but P not
type ConcreteIntParam int32

func NewConcreteIntParam(v int32) LeafParameter {
	m := ConcreteIntParam(v)
	return &m
}

func (m *ConcreteIntParam) IsCompatible(o LeafParameter) bool {
	if t, ok := o.(*ConcreteIntParam); ok {
		return t == m
	}
	return false
}

func (m *ConcreteIntParam) String() string {
	return fmt.Sprintf("%d", *m)
}
