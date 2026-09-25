package expr

import (
	"fmt"
	"strings"

	"github.com/substrait-io/substrait-go/v9/types"
)

// Lambda represents a lambda expression with parameters and a body.
type Lambda struct {
	Parameters *types.StructType // The formal lambda parameters, required to have NULLABILITY_REQUIRED
	Body       Expression
}

func (l *Lambda) String() string {
	var b strings.Builder
	b.WriteString("(")
	for i, t := range l.Parameters.Types {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "$%d: %s", i, t)
	}
	b.WriteString(") -> ")
	b.WriteString(l.Body.String())
	return b.String()
}

func (l *Lambda) isRootRef() {}

func (l *Lambda) IsScalar() bool {
	return false
}

func (l *Lambda) GetType() types.Type {
	return &types.FuncType{
		Nullability:    types.NullabilityRequired,
		ParameterTypes: l.Parameters.Types,
		ReturnType:     l.Body.GetType(),
	}
}

func (l *Lambda) Equals(other Expression) bool {
	rhs, ok := other.(*Lambda)
	if !ok {
		return false
	}
	return l.Parameters.Equals(rhs.Parameters) && l.Body.Equals(rhs.Body)
}

func (l *Lambda) Visit(visit VisitFunc) Expression {
	newBody := visit(l.Body)
	if newBody == l.Body {
		return l
	}
	return &Lambda{Parameters: l.Parameters, Body: newBody}
}
