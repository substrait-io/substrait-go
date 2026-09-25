// SPDX-License-Identifier: Apache-2.0

package wire_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/substrait-io/substrait-go/v9/types"
	"github.com/substrait-io/substrait-go/v9/wire"
)

func TestTypeRoundtrip(t *testing.T) {
	for _, nullable := range []bool{true, false} {
		t.Run(fmt.Sprintf("nullable=%t", nullable), func(t *testing.T) {
			n := NullabilityRequired
			if nullable {
				n = NullabilityNullable
			}

			tests := []Type{
				&BooleanType{Nullability: n},
				&Int8Type{Nullability: n},
				&Int16Type{Nullability: n},
				&Int32Type{Nullability: n},
				&Int64Type{Nullability: n},
				&Float32Type{Nullability: n},
				&Float64Type{Nullability: n},
				&StringType{Nullability: n},
				&BinaryType{Nullability: n},
				&TimeType{Nullability: n},
				&DateType{Nullability: n},
				&TimestampType{Nullability: n},
				&TimestampTzType{Nullability: n},
				&IntervalYearType{Nullability: n},
				&UUIDType{Nullability: n},
				&FixedCharType{Nullability: n, Length: 25},
				&VarCharType{Nullability: n, Length: 35},
				&FixedBinaryType{Nullability: n, Length: 45},
				&IntervalDayType{Nullability: n, Precision: 5},
				&IntervalDayType{Nullability: n, Precision: 0},
				NewIntervalCompoundType().WithPrecision(PrecisionEMinus7Seconds).WithNullability(n),

				&DecimalType{Nullability: n, Precision: 34, Scale: 3},
				&PrecisionTimeType{Nullability: n, Precision: PrecisionEMinus4Seconds},
				&PrecisionTimestampType{Nullability: n, Precision: PrecisionEMinus4Seconds},
				&PrecisionTimestampTzType{PrecisionTimestampType: PrecisionTimestampType{Nullability: n, Precision: PrecisionEMinus5Seconds}},
				&MapType{Nullability: n, Key: &Int8Type{}, Value: &Int16Type{Nullability: n}},
				&ListType{Nullability: n, Type: &TimeType{Nullability: n}},
				&StructType{Nullability: n, Types: []Type{
					&TimeType{Nullability: n}, &TimestampType{Nullability: n},
					&TimestampTzType{Nullability: n}}},
				&UserDefinedType{TypeParameters: []TypeParam{&DataTypeParameter{Type: &Int32Type{}}}, Nullability: n},
				&FuncType{Nullability: n, ParameterTypes: []Type{&Int8Type{}}, ReturnType: &Int16Type{Nullability: n}},
				&FuncType{Nullability: n, ParameterTypes: []Type{&Int32Type{Nullability: n}, &Float64Type{Nullability: n}}, ReturnType: &BooleanType{Nullability: NullabilityNullable}},
				&FuncType{Nullability: n, ParameterTypes: []Type{}, ReturnType: &Int32Type{Nullability: n}},
			}

			for _, tt := range tests {
				t.Run(tt.String(), func(t *testing.T) {
					converted := wire.TypeToProto(tt)
					convertedType := wire.TypeFromProto(converted)
					assert.True(t, tt.Equals(convertedType))
				})
			}
		})
	}
}
