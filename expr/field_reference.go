// SPDX-License-Identifier: Apache-2.0

package expr

import (
	"fmt"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
	"golang.org/x/exp/slices"
)

type RootRefType interface {
	isRootRef()
}

var RootReference RootRefType

type OuterReference uint32

func (OuterReference) isRootRef() {}

type ReferenceSegment interface {
	Reference
	fmt.Stringer
	GetChild() ReferenceSegment
	GetType(types.Type) (types.Type, error)
	ToProto() *proto.Expression_ReferenceSegment
	Equals(ReferenceSegment) bool
}

func RefSegmentFromProto(p *proto.Expression_ReferenceSegment) ReferenceSegment {
	if p == nil {
		return nil
	}

	switch seg := p.ReferenceType.(type) {
	case *proto.Expression_ReferenceSegment_MapKey_:
		return &MapKeyRef{
			MapKey: LiteralFromProto(seg.MapKey.MapKey),
			Child:  RefSegmentFromProto(seg.MapKey.Child),
		}
	case *proto.Expression_ReferenceSegment_StructField_:
		return &StructFieldRef{
			Field: seg.StructField.Field,
			Child: RefSegmentFromProto(seg.StructField.Child),
		}
	case *proto.Expression_ReferenceSegment_ListElement_:
		return &ListElementRef{
			Offset: seg.ListElement.Offset,
			Child:  RefSegmentFromProto(seg.ListElement.Child),
		}
	}

	return nil
}

func FlattenRefSegments(refs ...ReferenceSegment) ReferenceSegment {
	if len(refs) == 0 {
		return nil
	}

	if len(refs) == 1 {
		return refs[0]
	}

	out, cur := refs[0], refs[0]
	for _, r := range refs[1:] {
		switch parent := cur.(type) {
		case *MapKeyRef:
			parent.Child = r
		case *StructFieldRef:
			parent.Child = r
		case *ListElementRef:
			parent.Child = r
		}
		cur = r
	}

	return out
}

type MapKeyRef struct {
	MapKey Literal
	Child  ReferenceSegment
}

func NewMapKeyRef(key Literal) *MapKeyRef { return &MapKeyRef{MapKey: key} }

func (r *MapKeyRef) String() string {
	var c string
	if r.Child != nil {
		c = r.Child.String()
	}
	return ".[" + r.MapKey.String() + "]" + c
}

func (r *MapKeyRef) ToProto() *proto.Expression_ReferenceSegment {
	var c *proto.Expression_ReferenceSegment
	if r.Child != nil {
		c = r.Child.ToProto()
	}

	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_MapKey_{
			MapKey: &proto.Expression_ReferenceSegment_MapKey{
				MapKey: r.MapKey.ToProtoLiteral(),
				Child:  c,
			},
		},
	}
}

func (r *MapKeyRef) GetType(parentType types.Type) (types.Type, error) {
	mt, ok := parentType.(*types.MapType)
	if !ok {
		return nil, substraitgo.ErrInvalidType
	}

	if !r.MapKey.GetType().Equals(mt.Key) {
		return nil, substraitgo.ErrInvalidType
	}

	if r.Child != nil {
		return r.Child.GetType(mt.Value)
	}

	return mt.Value, nil
}
func (r *MapKeyRef) GetChild() ReferenceSegment { return r.Child }
func (r *MapKeyRef) Equals(rhs ReferenceSegment) bool {
	if rhs, ok := rhs.(*MapKeyRef); ok {
		if !r.MapKey.Equals(rhs.MapKey) {
			return false
		}

		if r.Child == rhs.Child {
			// both point to the same object or both are nil
			return true
		}

		if (r.Child == nil && rhs.Child != nil) ||
			(r.Child != nil && rhs.Child == nil) {
			return false
		}

		return r.Child.Equals(rhs.Child)
	}
	return false
}

func (*MapKeyRef) isRefType() {}

type StructFieldRef struct {
	Field int32
	Child ReferenceSegment
}

func NewStructFieldRef(field int32) *StructFieldRef { return &StructFieldRef{Field: field} }

func (r *StructFieldRef) String() string {
	var c string
	if r.Child != nil {
		c = r.Child.String()
	}

	return fmt.Sprintf(".field(%d)%s", r.Field, c)
}

func (r *StructFieldRef) GetType(parentType types.Type) (types.Type, error) {
	st, ok := parentType.(*types.StructType)
	if !ok {
		return nil, substraitgo.ErrInvalidType
	}

	if len(st.Types) < int(r.Field) {
		return nil, substraitgo.ErrInvalidType
	}

	if r.Child != nil {
		return r.Child.GetType(st.Types[r.Field])
	}

	return st.Types[r.Field], nil
}

func (r *StructFieldRef) ToProto() *proto.Expression_ReferenceSegment {
	var c *proto.Expression_ReferenceSegment
	if r.Child != nil {
		c = r.Child.ToProto()
	}

	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_StructField_{
			StructField: &proto.Expression_ReferenceSegment_StructField{
				Field: r.Field,
				Child: c,
			},
		},
	}
}

func (r *StructFieldRef) GetChild() ReferenceSegment { return r.Child }
func (r *StructFieldRef) Equals(rhs ReferenceSegment) bool {
	if rhs, ok := rhs.(*StructFieldRef); ok {
		if r.Field != rhs.Field {
			return false
		}

		if r.Child == rhs.Child {
			// both point to the same object or both are nil
			return true
		}

		if (r.Child == nil && rhs.Child != nil) ||
			(r.Child != nil && rhs.Child == nil) {
			return false
		}

		return r.Child.Equals(rhs.Child)
	}

	return false
}

func (*StructFieldRef) isRefType() {}

type ListElementRef struct {
	Offset int32
	Child  ReferenceSegment
}

func NewListElemRef(offset int32) *ListElementRef { return &ListElementRef{Offset: offset} }

func (r *ListElementRef) String() string {
	var c string
	if r.Child != nil {
		c = r.Child.String()
	}
	return fmt.Sprintf(".[%d]%s", r.Offset, c)
}

func (r *ListElementRef) GetType(parentType types.Type) (types.Type, error) {
	lt, ok := parentType.(*types.ListType)
	if !ok {
		return nil, substraitgo.ErrInvalidType
	}

	if r.Child != nil {
		return r.Child.GetType(lt.Type)
	}
	return lt.Type, nil
}

func (r *ListElementRef) ToProto() *proto.Expression_ReferenceSegment {
	var c *proto.Expression_ReferenceSegment
	if r.Child != nil {
		c = r.Child.ToProto()
	}

	return &proto.Expression_ReferenceSegment{
		ReferenceType: &proto.Expression_ReferenceSegment_ListElement_{
			ListElement: &proto.Expression_ReferenceSegment_ListElement{
				Offset: r.Offset,
				Child:  c,
			},
		},
	}
}

func (r *ListElementRef) GetChild() ReferenceSegment { return r.Child }
func (r *ListElementRef) Equals(rhs ReferenceSegment) bool {
	if rhs, ok := rhs.(*ListElementRef); ok {
		if r.Offset != rhs.Offset {
			return false
		}

		if r.Child == rhs.Child {
			// both point to the same object or both are nil
			return true
		}

		if (r.Child == nil && rhs.Child != nil) ||
			(r.Child != nil && rhs.Child == nil) {
			return false
		}

		return r.Child.Equals(rhs.Child)
	}

	return false
}
func (*ListElementRef) isRefType() {}

// MaskExpression is a reference that takes an existing subtype and
// selectively removes fields from it. For example, one might initially
// have an inner struct with 100 fields, but a particular operation only
// needs to interact with 2 of those 100 fields. In this situation, one
// would use a mask expression to eliminate the 98 fields that are not
// relevant to the rest of the operations pipeline.
//
// Note that this does not fundamentally alter the structure of data
// beyond the elimination of unnecessary elements.
type MaskExpression struct {
	sel              MaskStructSelect
	maintainSingular bool
}

func (*MaskExpression) isRefType() {}
func (e *MaskExpression) ToProto() *proto.Expression_MaskExpression {
	return &proto.Expression_MaskExpression{
		Select:                 e.sel.toProtoStructSelect(),
		MaintainSingularStruct: e.maintainSingular,
	}
}

func (e *MaskExpression) MaintainSingularStruct() bool {
	return e.maintainSingular
}

func (e *MaskExpression) Select() MaskStructSelect {
	return slices.Clone(e.sel)
}

func MaskExpressionFromProto(p *proto.Expression_MaskExpression) *MaskExpression {
	sel := make(MaskStructSelect, len(p.Select.StructItems))
	for i, item := range p.Select.StructItems {
		sel[i].field = item.Field
		if item.Child != nil {
			sel[i].child = maskSelectFromProto(item.Child)
		}
	}
	return &MaskExpression{sel: sel, maintainSingular: p.MaintainSingularStruct}
}

func maskSelectFromProto(p *proto.Expression_MaskExpression_Select) MaskSelect {
	switch s := p.Type.(type) {
	case *proto.Expression_MaskExpression_Select_Struct:
		items := make(MaskStructSelect, len(s.Struct.StructItems))
		for i, item := range s.Struct.StructItems {
			items[i].field = item.Field
			if item.Child != nil {
				items[i].child = maskSelectFromProto(item.Child)
			}
		}
		return items
	case *proto.Expression_MaskExpression_Select_List:
		selection := make([]MaskListSelectItem, len(s.List.Selection))
		for i, sel := range s.List.Selection {
			switch s := sel.Type.(type) {
			case *proto.Expression_MaskExpression_ListSelect_ListSelectItem_Item:
				selection[i] = (*MaskListElement)(s.Item)
			case *proto.Expression_MaskExpression_ListSelect_ListSelectItem_Slice:
				selection[i] = (*MaskListSlice)(s.Slice)
			}
		}
		return &MaskListSelect{
			selection: selection,
			child:     maskSelectFromProto(s.List.Child),
		}
	case *proto.Expression_MaskExpression_Select_Map:
		var ret MaskMapSelect
		if s.Map.Child != nil {
			ret.child = maskSelectFromProto(s.Map.Child)
		}

		switch sk := s.Map.Select.(type) {
		case *proto.Expression_MaskExpression_MapSelect_Expression:
			ret.key = sk.Expression.MapKeyExpression
			ret.kind = MapSelectExpr
		case *proto.Expression_MaskExpression_MapSelect_Key:
			ret.key = sk.Key.MapKey
			ret.kind = MapSelectKey
		}
		return &ret
	}
	panic("unimplemented mask select type")
}

type MaskSelect interface {
	ToProto() *proto.Expression_MaskExpression_Select
}

type MaskStructSelect []MaskStructItem

func (m MaskStructSelect) toProtoStructSelect() *proto.Expression_MaskExpression_StructSelect {
	items := make([]*proto.Expression_MaskExpression_StructItem, len(m))
	for i, item := range m {
		items[i] = item.ToProto()
	}
	return &proto.Expression_MaskExpression_StructSelect{
		StructItems: items,
	}
}

func (m MaskStructSelect) ToProto() *proto.Expression_MaskExpression_Select {
	return &proto.Expression_MaskExpression_Select{
		Type: &proto.Expression_MaskExpression_Select_Struct{
			Struct: m.toProtoStructSelect(),
		},
	}
}

type MaskStructItem struct {
	field int32
	child MaskSelect
}

func (m *MaskStructItem) Field() int32      { return m.field }
func (m *MaskStructItem) Child() MaskSelect { return m.child }
func (m *MaskStructItem) ToProto() *proto.Expression_MaskExpression_StructItem {
	return &proto.Expression_MaskExpression_StructItem{
		Field: m.field,
		Child: m.child.ToProto(),
	}
}

type MaskListSelect struct {
	selection []MaskListSelectItem
	child     MaskSelect
}

func (m *MaskListSelect) ToProto() *proto.Expression_MaskExpression_Select {
	selection := make([]*proto.Expression_MaskExpression_ListSelect_ListSelectItem, len(m.selection))
	for i, s := range m.selection {
		selection[i] = s.ToProto()
	}

	return &proto.Expression_MaskExpression_Select{
		Type: &proto.Expression_MaskExpression_Select_List{
			List: &proto.Expression_MaskExpression_ListSelect{
				Selection: selection,
				Child:     m.child.ToProto(),
			},
		},
	}
}

func (m *MaskListSelect) Child() MaskSelect { return m.child }
func (m *MaskListSelect) Selection() []MaskListSelectItem {
	return slices.Clone(m.selection)
}

type MaskListSelectItem interface {
	ToProto() *proto.Expression_MaskExpression_ListSelect_ListSelectItem
}

type MaskListElement proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListElement

func (m *MaskListElement) GetField() int32 {
	return m.Field
}

func (m *MaskListElement) ToProto() *proto.Expression_MaskExpression_ListSelect_ListSelectItem {
	return &proto.Expression_MaskExpression_ListSelect_ListSelectItem{
		Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Item{
			Item: (*proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListElement)(m),
		},
	}
}

type MaskListSlice proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListSlice

func (m *MaskListSlice) GetBounds() (start, end int32) {
	return m.Start, m.End
}

func (m *MaskListSlice) ToProto() *proto.Expression_MaskExpression_ListSelect_ListSelectItem {
	return &proto.Expression_MaskExpression_ListSelect_ListSelectItem{
		Type: &proto.Expression_MaskExpression_ListSelect_ListSelectItem_Slice{
			Slice: (*proto.Expression_MaskExpression_ListSelect_ListSelectItem_ListSlice)(m),
		},
	}
}

type MapSelectKind int8

const (
	MapSelectKey MapSelectKind = iota
	MapSelectExpr
)

type MaskMapSelect struct {
	child MaskSelect
	kind  MapSelectKind
	key   string
}

func (m *MaskMapSelect) KeyKind() MapSelectKind { return m.kind }
func (m *MaskMapSelect) Key() string            { return m.key }

func (m *MaskMapSelect) Child() MaskSelect {
	return m.child
}

func (m *MaskMapSelect) ToProto() *proto.Expression_MaskExpression_Select {
	ret := &proto.Expression_MaskExpression_Select_Map{
		Map: &proto.Expression_MaskExpression_MapSelect{
			Child: m.child.ToProto(),
		},
	}

	if m.kind == MapSelectKey {
		ret.Map.Select = &proto.Expression_MaskExpression_MapSelect_Key{
			Key: &proto.Expression_MaskExpression_MapSelect_MapKey{
				MapKey: m.key,
			},
		}
	} else {
		ret.Map.Select = &proto.Expression_MaskExpression_MapSelect_Expression{
			Expression: &proto.Expression_MaskExpression_MapSelect_MapKeyExpression{
				MapKeyExpression: m.key,
			},
		}
	}
	return &proto.Expression_MaskExpression_Select{
		Type: ret,
	}
}

type Reference interface {
	isRefType()
}

type FieldReference struct {
	Reference Reference
	Root      RootRefType

	knownType types.Type
}

func NewRootFieldRef(ref Reference, baseSchema *types.StructType) (*FieldReference, error) {
	return NewFieldRef(RootReference, ref, baseSchema)
}

func NewFieldRef(root RootRefType, ref Reference, baseSchema *types.StructType) (*FieldReference, error) {
	if ref != nil && root == RootReference && baseSchema == nil {
		return nil, fmt.Errorf("%w: must provide the base schema to create a root field ref",
			substraitgo.ErrInvalidExpr)
	}

	switch rt := ref.(type) {
	case ReferenceSegment:
		var rootType types.Type
		if root == RootReference {
			rootType = baseSchema
		} else if rootExpr, ok := root.(Expression); ok {
			rootType = rootExpr.GetType()
		} else {
			return nil, fmt.Errorf("%w: unknown root reference type %v",
				substraitgo.ErrInvalidExpr, root)
		}

		typ, err := rt.GetType(rootType)
		if err != nil {
			return nil, fmt.Errorf("error resolving ref type: %w", err)
		}
		return &FieldReference{
			Reference: ref,
			Root:      root,
			knownType: typ,
		}, nil
	case *MaskExpression:
	}

	return nil, substraitgo.ErrNotImplemented
}

func (*FieldReference) isRootRef() {}

func (f *FieldReference) String() string {
	var b strings.Builder
	if rootExpr, ok := f.Root.(Expression); ok {
		b.WriteString("[root:(")
		b.WriteString(rootExpr.String())
		b.WriteString(")]")
	} else if outerRef, ok := f.Root.(OuterReference); ok {
		fmt.Fprintf(&b, "[outerRef:%d]", outerRef)
	}

	var typ string
	if f.knownType != nil {
		typ = " => " + f.knownType.String()
	}
	return b.String() + f.Reference.(ReferenceSegment).String() + typ
}

func (f *FieldReference) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Value{Value: f.ToProto()},
	}
}
func (f *FieldReference) ToProtoFieldRef() *proto.Expression_FieldReference {
	ret := &proto.Expression_FieldReference{}
	switch r := f.Reference.(type) {
	case ReferenceSegment:
		ret.ReferenceType = &proto.Expression_FieldReference_DirectReference{
			DirectReference: r.ToProto()}
	case *MaskExpression:
		ret.ReferenceType = &proto.Expression_FieldReference_MaskedReference{
			MaskedReference: r.ToProto(),
		}
	}

	if f.Root != RootReference {
		switch r := f.Root.(type) {
		case Expression:
			ret.RootType = &proto.Expression_FieldReference_Expression{
				Expression: r.ToProto(),
			}
		case OuterReference:
			ret.RootType = &proto.Expression_FieldReference_OuterReference_{
				OuterReference: &proto.Expression_FieldReference_OuterReference{
					StepsOut: uint32(r),
				},
			}
		}
	} else {
		ret.RootType = &proto.Expression_FieldReference_RootReference_{
			RootReference: &proto.Expression_FieldReference_RootReference{},
		}
	}

	return ret
}

func (f *FieldReference) ToProto() *proto.Expression {
	return &proto.Expression{
		RexType: &proto.Expression_Selection{
			Selection: f.ToProtoFieldRef(),
		},
	}
}

func (f *FieldReference) Equals(rhs Expression) bool {
	if rhs, ok := rhs.(*FieldReference); ok {
		switch root := f.Root.(type) {
		case OuterReference:
			rhsRoot, ok := rhs.Root.(OuterReference)
			if !ok {
				return false
			}

			if rhsRoot != root {
				return false
			}
		case Expression:
			rhsExpr, ok := rhs.Root.(Expression)
			if !ok {
				return false
			}

			if !root.Equals(rhsExpr) {
				return false
			}
		default:
			if rhs.Root != RootReference {
				return false
			}
		}

		switch ref := f.Reference.(type) {
		case ReferenceSegment:
			rhs, ok := rhs.Reference.(ReferenceSegment)
			if !ok {
				return false
			}

			return ref.Equals(rhs)
		case *MaskExpression:
			rhs, ok := rhs.Reference.(*MaskExpression)
			if !ok {
				return false
			}

			return ref == rhs
		}
	}

	return false
}

func (f *FieldReference) GetType() types.Type {
	return f.knownType
}

func (f *FieldReference) Visit(v VisitFunc) Expression {
	return f
}

func (*FieldReference) IsScalar() bool { return true }

func FieldReferenceFromProto(p *proto.Expression_FieldReference, baseSchema types.Type, reg ExtensionRegistry) (*FieldReference, error) {
	var (
		ref       Reference
		root      RootRefType
		knownType types.Type
		err       error
	)

	switch rt := p.RootType.(type) {
	case *proto.Expression_FieldReference_Expression:
		if root, err = ExprFromProto(rt.Expression, baseSchema, reg); err != nil {
			return nil, err
		}
	case *proto.Expression_FieldReference_OuterReference_:
		root = OuterReference(rt.OuterReference.StepsOut)
	case *proto.Expression_FieldReference_RootReference_:
		root = RootReference
	}

	switch rt := p.ReferenceType.(type) {
	case *proto.Expression_FieldReference_DirectReference:
		refseg := RefSegmentFromProto(rt.DirectReference)
		if root == RootReference && baseSchema != nil {
			knownType, err = refseg.GetType(baseSchema)
			if err != nil {
				return nil, err
			}
		} else if rootExpr, ok := root.(Expression); ok {
			knownType, err = refseg.GetType(rootExpr.GetType())
			if err != nil {
				return nil, err
			}
		}

		ref = refseg

	case *proto.Expression_FieldReference_MaskedReference:
		ref = MaskExpressionFromProto(rt.MaskedReference)
	}

	return &FieldReference{Root: root, Reference: ref, knownType: knownType}, nil
}
