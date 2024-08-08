package functions

import (
	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/types"
	"strconv"
	"strings"
)

var nameToTypeMap = map[string]types.Type{
	"boolean":       &types.BooleanType{},
	"i8":            &types.Int8Type{},
	"i16":           &types.Int16Type{},
	"i32":           &types.Int32Type{},
	"i64":           &types.Int64Type{},
	"fp32":          &types.Float32Type{},
	"fp64":          &types.Float64Type{},
	"binary":        &types.BinaryType{},
	"string":        &types.StringType{},
	"timestamp":     &types.TimestampType{},
	"date":          &types.DateType{},
	"time":          &types.TimeType{},
	"timestamp_tz":  &types.TimestampTzType{},
	"interval_year": &types.IntervalYearType{},
	"interval_day":  &types.IntervalDayType{},
	"uuid":          &types.UUIDType{},

	"fixedbinary": &types.FixedBinaryType{},
	"fixedchar":   &types.FixedCharType{},
	"varchar":     &types.VarCharType{},
	"decimal":     &types.DecimalType{},

	// short names
	"bool":  &types.BooleanType{},
	"vbin":  &types.BinaryType{},
	"str":   &types.StringType{},
	"ts":    &types.TimestampType{},
	"tstz":  &types.TimestampTzType{},
	"iyear": &types.IntervalYearType{},
	"iday":  &types.IntervalDayType{},

	"fbin":  &types.FixedBinaryType{},
	"fchar": &types.FixedCharType{},
	"vchar": &types.VarCharType{},
	"dec":   &types.DecimalType{},
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
	case types.ParameterizedType:
		return ti.typ.(types.ParameterizedType).BaseString()
	}
	return ti.typ.String()
}

func (ti *typeInfo) getLocalTypeString(input types.Type, enclosure typeEnclosure) string {
	if paramType, ok := input.(types.ParameterizedType); ok {
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

func (t *localTypeRegistryImpl) IsTypeSupportedInTables(typ types.Type) bool {
	if ti, ok := t.typeInfoMap[typ.ShortString()]; ok {
		return ti.supportedAsColumn
	}
	return false
}
