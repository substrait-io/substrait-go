// SPDX-License-Identifier: Apache-2.0

package leaf_parameters

import "fmt"

// LeafParameter represents a parameter type
// parameter can of concrete (38) or abstract type (P)
// or another parameterized type like VARCHAR<"L1">
type LeafParameter interface {
	// IsCompatible is type compatible with other
	// compatible is other can be used in place of this type
	IsCompatible(other LeafParameter) bool
	fmt.Stringer
}
