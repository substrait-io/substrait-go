// SPDX-License-Identifier: Apache-2.0

package functions

import (
	"strconv"
	"strings"

	substraitgo "github.com/substrait-io/substrait-go/v3"
	"github.com/substrait-io/substrait-go/v3/types"
)

var (
	nameToTypeMap  map[string]types.Type
	toShortNameMap map[string]string
)

func init() {
	initTypeMaps()
}

func initTypeMaps() {
	nameToTypeMap = types.GetTypeNameToTypeMap()
	toShortNameMap = make(map[string]string)
	for k := range nameToTypeMap {
		shortName := types.GetShortTypeName(types.TypeName(k))
		if shortName != k {
			toShortNameMap[k] = shortName
		}
	}
	for k, v := range toShortNameMap {
		nameToTypeMap[v] = nameToTypeMap[k]
	}
}

func getTypeFromBaseTypeName(baseType string) (types.Type, error) {
	if typ, ok := nameToTypeMap[baseType]; ok {
		return typ, nil
	}
	return nil, substraitgo.ErrNotFound
}

var substraitEnclosure = &substraitTypeEnclosure{}

func isSupportedType(typeString string) bool {
	_, ok := nameToTypeMap[typeString]
	return ok
}

type typeRegistryImpl struct {
	typeMap map[string]types.Type
}

func NewTypeRegistry() TypeRegistry {
	return &typeRegistryImpl{typeMap: nameToTypeMap}
}

func (t *typeRegistryImpl) GetTypeFromTypeString(typeString string) (types.Type, error) {
	return getTypeFromTypeString(typeString, t.typeMap, substraitEnclosure)
}

func getTypeFromTypeString(typeString string, typeMap map[string]types.Type, enclosure typeEnclosure) (types.Type, error) {
	baseType, parameters, err := extractTypeAndParameters(typeString, enclosure)
	if err != nil {
		return nil, err
	}

	nullable := types.NullabilityRequired
	if strings.HasSuffix(baseType, "?") {
		baseType = baseType[:len(baseType)-1]
		nullable = types.NullabilityNullable
	}
	if typ, ok := typeMap[baseType]; ok {
		if typ, err = getTypeWithParameters(typ, parameters); err != nil {
			return nil, err
		}
		return typ.WithNullability(nullable), nil
	}
	return nil, substraitgo.ErrNotFound
}

func getTypeWithParameters(typ types.Type, parameters []int32) (types.Type, error) {
	switch typ.(type) {
	case *types.DecimalType:
		if len(parameters) != 2 {
			return nil, substraitgo.ErrInvalidType
		}
		return &types.DecimalType{Precision: parameters[0], Scale: parameters[1]}, nil
	case *types.FixedBinaryType, *types.FixedCharType, *types.VarCharType:
		if len(parameters) != 1 {
			return nil, substraitgo.ErrInvalidType
		}
		switch typ.(type) {
		case *types.FixedBinaryType:
			return &types.FixedBinaryType{Length: parameters[0]}, nil
		case *types.FixedCharType:
			return &types.FixedCharType{Length: parameters[0]}, nil
		case *types.VarCharType:
			return &types.VarCharType{Length: parameters[0]}, nil
		}
	default:
		if len(parameters) != 0 {
			return nil, substraitgo.ErrInvalidType
		}
	}
	return typ, nil
}

func extractTypeAndParameters(typeString string, enclosure typeEnclosure) (string, []int32, error) {
	conStart, conEnd := enclosure.containerStart(), enclosure.containerEnd()
	if !strings.Contains(typeString, conStart) || !strings.HasSuffix(typeString, conEnd) {
		return typeString, nil, nil
	}
	baseType := typeString[:strings.Index(typeString, conStart)]
	paramStr := typeString[strings.Index(typeString, conStart)+1 : len(typeString)-len(conEnd)]
	params := strings.Split(paramStr, ",")
	parameters := make([]int32, len(params))
	for i, p := range params {
		intValue, err := strconv.ParseInt(p, 10, 32)
		if err != nil {
			return "", nil, err
		}
		parameters[i] = int32(intValue)
	}
	return baseType, parameters, nil
}

type typeEnclosure interface {
	containerStart() string
	containerEnd() string
}

type substraitTypeEnclosure struct{}

func (t *substraitTypeEnclosure) containerStart() string {
	return "<"
}

func (t *substraitTypeEnclosure) containerEnd() string {
	return ">"
}

type typeInfo struct {
	typ               types.Type
	shortName         string
	localName         string
	supportedAsColumn bool
}

func (ti *typeInfo) getLongName() string {
	switch ti.typ.(type) {
	case types.CompositeType:
		return ti.typ.(types.CompositeType).BaseString()
	}
	return ti.typ.String()
}

func (ti *typeInfo) getLocalTypeString(input types.Type, enclosure typeEnclosure) string {
	if paramType, ok := input.(types.CompositeType); ok {
		return ti.localName + enclosure.containerStart() + paramType.ParameterString() + enclosure.containerEnd()
	}
	return ti.localName
}

type localTypeRegistryImpl struct {
	nameToType      map[string]types.Type
	localNameToType map[string]types.Type
	typeInfoMap     map[string]typeInfo
}

func NewLocalTypeRegistry(typeInfos []typeInfo) LocalTypeRegistry {
	nameToType := make(map[string]types.Type)
	localNameToType := make(map[string]types.Type)
	typeInfoMap := make(map[string]typeInfo)
	for _, ti := range typeInfos {
		nameToType[ti.shortName] = ti.typ
		localNameToType[ti.localName] = ti.typ
		typeInfoMap[ti.shortName] = ti
		longName := ti.getLongName()
		if longName != ti.shortName {
			nameToType[longName] = ti.typ
			typeInfoMap[longName] = ti
		}
	}
	return &localTypeRegistryImpl{
		nameToType:      nameToType,
		localNameToType: localNameToType,
		typeInfoMap:     typeInfoMap,
	}
}

func (t *localTypeRegistryImpl) containerStart() string {
	return "("
}

func (t *localTypeRegistryImpl) containerEnd() string {
	return ")"
}

func (t *localTypeRegistryImpl) GetTypeFromTypeString(typeString string) (types.Type, error) {
	return getTypeFromTypeString(typeString, t.nameToType, substraitEnclosure)
}

func (t *localTypeRegistryImpl) GetSubstraitTypeFromLocalType(localType string) (types.Type, error) {
	return getTypeFromTypeString(localType, t.localNameToType, t)
}

func (t *localTypeRegistryImpl) GetLocalTypeFromSubstraitType(typ types.Type) (string, error) {
	// TODO handle nullable
	name := typ.ShortString()
	if ti, ok := t.typeInfoMap[name]; ok {
		return ti.getLocalTypeString(typ, t), nil
	}
	return "", substraitgo.ErrNotFound
}

func (t *localTypeRegistryImpl) GetSupportedTypes() map[string]types.Type {
	return t.localNameToType
}

func (t *localTypeRegistryImpl) IsTypeSupportedInTables(typ types.Type) bool {
	if ti, ok := t.typeInfoMap[typ.ShortString()]; ok {
		return ti.supportedAsColumn
	}
	return false
}
