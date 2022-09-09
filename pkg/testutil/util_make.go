// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"fmt"
	"time"

	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
)

// All vectors generated by the Make Function, their memory is not allocated through the memory pool
// if you want to generate a vector in memory pool, use NewFunction to instead of MakeFunction.
var (
	MakeBoolVector = func(values []bool) *vector.Vector {
		return makeVector(values, nil, boolType)
	}

	MakeBooleanlVector = func(values []bool, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, boolType)
	}

	MakeInt64Vector = func(values []int64, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, int64Type)
	}

	MakeInt32Vector = func(values []int32, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, int32Type)
	}

	MakeInt16Vector = func(values []int16, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, int16Type)
	}

	MakeInt8Vector = func(values []int8, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, int8Type)
	}

	MakeUint64Vector = func(values []uint64, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, uint64Type)
	}

	MakeUint32Vector = func(values []uint32, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, uint32Type)
	}

	MakeUint16Vector = func(values []uint16, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, uint16Type)
	}

	MakeUint8Vector = func(values []uint8, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, uint8Type)
	}

	MakeFloat32Vector = func(values []float32, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, float32Type)
	}

	MakeFloat64Vector = func(values []float64, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, float64Type)
	}

	MakeCharVector = func(values []string, nsp []uint64) *vector.Vector {
		return makeStringVector(values, nsp, charType)
	}

	MakeVarcharVector = func(values []string, nsp []uint64) *vector.Vector {
		return makeStringVector(values, nsp, varcharType)
	}

	MakeTextVector = func(values []string, nsp []uint64) *vector.Vector {
		return makeStringVector(values, nsp, textType)
	}

	MakeDecimal64Vector = func(values []int64, nsp []uint64, _ types.Type) *vector.Vector {
		vec := vector.New(decimal64Type)
		cols := make([]types.Decimal64, len(values))
		if nsp == nil {
			for i, v := range values {
				cols[i] = types.Decimal64_FromInt64(v)
			}
		} else {
			for _, n := range nsp {
				nulls.Add(vec.Nsp, n)
			}
			for i, v := range values {
				if nulls.Contains(vec.Nsp, uint64(i)) {
					continue
				}
				cols[i] = types.Decimal64_FromInt64(v)
			}
		}
		vec.Col = cols
		return vec
	}

	MakeDecimal128Vector = func(values []int64, nsp []uint64, _ types.Type) *vector.Vector {
		vec := vector.New(decimal128Type)
		cols := make([]types.Decimal128, len(values))
		if nsp == nil {
			for i, v := range values {
				d := types.InitDecimal128(v)
				cols[i] = d
			}
		} else {
			for _, n := range nsp {
				nulls.Add(vec.Nsp, n)
			}
			for i, v := range values {
				if nulls.Contains(vec.Nsp, uint64(i)) {
					continue
				}
				d := types.InitDecimal128(v)
				cols[i] = d
			}
		}
		vec.Col = cols
		return vec
	}

	MakeDateVector = func(values []string, nsp []uint64) *vector.Vector {
		ds := make([]types.Date, len(values))
		ns := nulls.Build(len(values), nsp...)
		for i, s := range values {
			if nulls.Contains(ns, uint64(i)) {
				continue
			}
			d, err := types.ParseDate(s)
			if err != nil {
				panic(err)
			}
			ds[i] = d
		}
		return vector.NewWithFixed(types.T_date.ToType(), ds, ns, nil)
	}

	MakeDateTimeVector = func(values []string, nsp []uint64) *vector.Vector {
		ds := make([]types.Datetime, len(values))
		ns := nulls.Build(len(values), nsp...)
		for i, s := range values {
			if nulls.Contains(ns, uint64(i)) {
				continue
			}
			d, err := types.ParseDatetime(s, 6)
			if err != nil {
				panic(err)
			}
			ds[i] = d
		}
		return vector.NewWithFixed(types.T_datetime.ToType(), ds, ns, nil)
	}

	MakeTimeStampVector = func(values []string, nsp []uint64) *vector.Vector {
		ds := make([]types.Timestamp, len(values))
		ns := nulls.Build(len(values), nsp...)
		for i, s := range values {
			if nulls.Contains(ns, uint64(i)) {
				continue
			}
			d, err := types.ParseTimestamp(time.Local, s, 6)
			if err != nil {
				panic(err)
			}
			ds[i] = d
		}
		return vector.NewWithFixed(types.T_timestamp.ToType(), ds, ns, nil)
	}
)

// functions to make a scalar vector for test.
var (
	MakeScalarNull = func(typ types.T, length int) *vector.Vector {
		vec := NewProc().AllocConstNullVector(typ.ToType(), length)
		return vec
	}

	MakeScalarBool = func(v bool, length int) *vector.Vector {
		return makeScalar(v, length, boolType)
	}

	MakeScalarInt64 = func(v int64, length int) *vector.Vector {
		return makeScalar(v, length, int64Type)
	}

	MakeScalarInt32 = func(v int32, length int) *vector.Vector {
		return makeScalar(v, length, int32Type)
	}

	MakeScalarInt16 = func(v int16, length int) *vector.Vector {
		return makeScalar(v, length, int16Type)
	}

	MakeScalarInt8 = func(v int8, length int) *vector.Vector {
		return makeScalar(v, length, int8Type)
	}

	MakeScalarUint64 = func(v uint64, length int) *vector.Vector {
		return makeScalar(v, length, uint64Type)
	}

	MakeScalarUint3 = func(v uint32, length int) *vector.Vector {
		return makeScalar(v, length, uint32Type)
	}

	MakeScalarUint16 = func(v uint16, length int) *vector.Vector {
		return makeScalar(v, length, uint16Type)
	}

	MakeScalarUint8 = func(v uint8, length int) *vector.Vector {
		return makeScalar(v, length, uint8Type)
	}

	MakeScalarFloat32 = func(v float32, length int) *vector.Vector {
		return makeScalar(v, length, float32Type)
	}

	MakeScalarFloat64 = func(v float64, length int) *vector.Vector {
		return makeScalar(v, length, float64Type)
	}

	MakeScalarChar = func(value string, length int) *vector.Vector {
		return makeScalarString(value, length, charType)
	}

	MakeScalarVarchar = func(value string, length int) *vector.Vector {
		return makeScalarString(value, length, varcharType)
	}

	MakeTextVarchar = func(value string, length int) *vector.Vector {
		return makeScalarString(value, length, textType)
	}

	MakeScalarDate = func(value string, length int) *vector.Vector {
		d, err := types.ParseDate(value)
		if err != nil {
			panic(err)
		}
		return vector.NewConstFixed(dateType, length, d)
	}

	MakeScalarDateTime = func(value string, length int) *vector.Vector {
		d, err := types.ParseDatetime(value, 6)
		if err != nil {
			panic(err)
		}
		return vector.NewConstFixed(datetimeType, length, d)
	}

	MakeScalarTimeStamp = func(value string, length int) *vector.Vector {
		d, err := types.ParseTimestamp(time.Local, value, 6)
		if err != nil {
			panic(err)
		}
		return vector.NewConstFixed(timestampType, length, d)
	}

	MakeScalarDecimal64 = func(v int64, length int, _ types.Type) *vector.Vector {
		return vector.NewConstFixed(decimal64Type, length, types.InitDecimal64(v))
	}

	MakeScalarDecimal128 = func(v uint64, length int, _ types.Type) *vector.Vector {
		return vector.NewConstFixed(decimal64Type, length, types.InitDecimal128UsingUint(v))
	}

	MakeScalarDecimal128ByFloat64 = func(v float64, length int, _ types.Type) *vector.Vector {
		val := fmt.Sprintf("%f", v)
		_, scale, err := types.ParseStringToDecimal128WithoutTable(val)
		if err != nil {
			panic(err)
		}
		dec128Val, err := types.ParseStringToDecimal128(val, 34, scale)
		if err != nil {
			panic(err)
		}
		return vector.NewConstFixed(decimal128Type, length, dec128Val)
	}
)

func makeVector[T types.FixedSizeT](values []T, nsp []uint64, typ types.Type) *vector.Vector {
	ns := nulls.Build(len(values), nsp...)
	return vector.NewWithFixed(typ, values, ns, nil)
}

func makeStringVector(values []string, nsp []uint64, typ types.Type) *vector.Vector {
	if nsp == nil {
		return vector.NewWithStrings(typ, values, nil, nil)
	} else {
		vnsp := nulls.Build(len(values), nsp...)
		return vector.NewWithStrings(typ, values, vnsp, nil)
	}
}

func makeScalar[T types.FixedSizeT](value T, length int, typ types.Type) *vector.Vector {
	return vector.NewConstFixed(typ, length, value)
}

func makeScalarString(value string, length int, typ types.Type) *vector.Vector {
	return vector.NewConstString(typ, length, value)
}
