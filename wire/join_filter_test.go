// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/v9/expr"
	"github.com/substrait-io/substrait-go/v9/plan"
	"github.com/substrait-io/substrait-go/v9/wire"
)

// nilFilterInputs builds two single-column scans to use as join inputs.
func nilFilterInputs() (plan.Rel, plan.Rel) {
	b := plan.NewBuilderDefault()
	schema := nilFilterSchema()
	return b.NamedScan([]string{"L"}, schema), b.NamedScan([]string{"R"}, schema)
}

// A nil post-join filter must be omitted from the encoded proto, not replaced
// by the default true literal the read-time getter substitutes. See the
// aggregate-measure test for why the round-trip cannot catch this.
func TestJoinRelNilPostJoinFilterOmitted(t *testing.T) {
	left, right := nilFilterInputs()
	rel := plan.NewJoinRel(left, right, plan.JoinTypeInner, expr.NewPrimitiveLiteral(true, false), nil, plan.RelCommon{}, nil)

	assert.Nil(t, rel.RawPostJoinFilter(), "raw accessor must preserve nil")
	assert.NotNil(t, rel.PostJoinFilter(), "getter substitutes the default true literal")
	assert.Nil(t, wire.RelToProto(rel).GetJoin().GetPostJoinFilter(),
		"a nil post-join filter must be absent from the encoded proto")
}

func TestHashJoinRelNilPostJoinFilterOmitted(t *testing.T) {
	left, right := nilFilterInputs()
	rel := plan.NewHashJoinRel(left, right, nil, plan.HashMergeInner, nil, plan.RelCommon{}, nil)

	assert.Nil(t, rel.RawPostJoinFilter())
	assert.NotNil(t, rel.PostJoinFilter())
	assert.Nil(t, wire.RelToProto(rel).GetHashJoin().GetPostJoinFilter(),
		"a nil post-join filter must be absent from the encoded proto")
}

func TestMergeJoinRelNilPostJoinFilterOmitted(t *testing.T) {
	left, right := nilFilterInputs()
	rel := plan.NewMergeJoinRel(left, right, nil, plan.HashMergeInner, nil, plan.RelCommon{}, nil)

	assert.Nil(t, rel.RawPostJoinFilter())
	assert.NotNil(t, rel.PostJoinFilter())
	assert.Nil(t, wire.RelToProto(rel).GetMergeJoin().GetPostJoinFilter(),
		"a nil post-join filter must be absent from the encoded proto")
}
