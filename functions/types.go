package functions

import (
	substraitgo "github.com/substrait-io/substrait-go"
	"github.com/substrait-io/substrait-go/types"
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

	// short names
	"bool":  &types.BooleanType{},
	"vbin":  &types.BinaryType{},
	"str":   &types.StringType{},
	"ts":    &types.TimestampType{},
	"tstz":  &types.TimestampTzType{},
	"iyear": &types.IntervalYearType{},
	"iday":  &types.IntervalDayType{},
}

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
	return getTypeFromTypeString(typeString, t.typeMap)
}

func getTypeFromTypeString(typeString string, typeMap map[string]types.Type) (types.Type, error) {
	// TODO handle parameterized types
	nullable := types.NullabilityRequired
	if strings.HasSuffix(typeString, "?") {
		typeString = typeString[:len(typeString)-1]
		nullable = types.NullabilityNullable
	}
	if typ, ok := typeMap[typeString]; ok {
		return typ.WithNullability(nullable), nil
	}
	return nil, substraitgo.ErrNotFound
}

type localTypeRegistryImpl struct {
	nameToType      map[string]types.Type
	localNameToType map[string]types.Type
	typeInfoMap     map[string]typeInfo
}

type typeInfo struct {
	typ               types.Type
	shortName         string
	localName         string
	supportedAsColumn bool
}

func NewLocalTypeRegistry(typeInfos []typeInfo) LocalTypeRegistry {
	nameToType := make(map[string]types.Type)
	localNameToType := make(map[string]types.Type)
	typeInfoMap := make(map[string]typeInfo)
	for _, ti := range typeInfos {
		nameToType[ti.shortName] = ti.typ
		localNameToType[ti.localName] = ti.typ
		typeInfoMap[ti.shortName] = ti
		longName := ti.typ.String()
		if longName != ti.shortName {
			// add long name if it is different from short name
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

func (t *localTypeRegistryImpl) GetTypeFromTypeString(typeString string) (types.Type, error) {
	return getTypeFromTypeString(typeString, t.nameToType)
}

func (t *localTypeRegistryImpl) GetSubstraitTypeFromLocalType(localType string) (types.Type, error) {
	return getTypeFromTypeString(localType, t.localNameToType)
}

func (t *localTypeRegistryImpl) GetLocalTypeFromSubstraitType(typ types.Type) (string, error) {
	// TODO check if nullable needs to be handled
	name := typ.ShortString()
	if ti, ok := t.typeInfoMap[name]; ok {
		return ti.localName, nil
	}
	return "", substraitgo.ErrNotFound
}

func (t *localTypeRegistryImpl) IsTypeSupportedInTables(typ types.Type) bool {
	if ti, ok := t.typeInfoMap[typ.ShortString()]; ok {
		return ti.supportedAsColumn
	}
	return false
}
