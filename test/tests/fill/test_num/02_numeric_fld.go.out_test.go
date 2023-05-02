package test_num

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NumStruct01_Unmarshal(t *testing.T) {
	t.Run("empty_struct", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Zero(t, num.IntFld)
		require.Zero(t, num.IntFld8)
		require.Zero(t, num.IntFld16)
		require.Zero(t, num.IntFld32)
		require.Zero(t, num.IntFld64)
		require.Zero(t, num.UintFld)
		require.Zero(t, num.UintFld8)
		require.NotZero(t, num.UintFld16)
		require.Zero(t, num.UintFld32)
		require.Zero(t, num.UintFld64)
		require.Zero(t, num.FloatFld32)
		require.Zero(t, num.FloatFld64)
		require.Nil(t, num.RefIntFld)
		require.Nil(t, num.RefIntFld8)
		require.Nil(t, num.RefIntFld16)
		require.NotNil(t, num.RefIntFld32)
		require.Nil(t, num.RefIntFld64)
		require.Nil(t, num.RefUintFld)
		require.Nil(t, num.RefUintFld8)
		require.Nil(t, num.RefUintFld16)
		require.Nil(t, num.RefUintFld32)
		require.Nil(t, num.RefUintFld64)
		require.Nil(t, num.RefFloatFld32)
		require.Nil(t, num.RefFloatFld64)
	})
	t.Run("fulfilled_struct", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 12345, "int_fld8": 25, "int_fld16": 32767, "int_fld32": 12345, "int_fld64": 12345, "Uint_fld": 12345, "Uint_fld8": 12, "Uint_fld16": 12345, "Uint_fld32": 12345, "Uint_fld64": 12345, "fl23": 12345, "fl64": 12345, "ref_int_fld": 12345, "ref_int_fld8": 33, "ref_int_fld16": 12345, "ref_int_fld32": 12345, "ref_int_fld64": 12345, "ref_Uint_fld": 12345, "ref_Uint_fld8": 45, "ref_Uint_fld16": 12345, "ref_Uint_fld32": 12345, "ref_Uint_fld64": 12345, "ref_fl23": 12345, "ref_fl64": 12345}`))
		require.NoError(t, err)
		require.Equal(t, num.IntFld, 12345)
		require.Equal(t, num.IntFld8, int8(25))
		require.Equal(t, num.UintFld8, uint8(12))
		require.NotNil(t, *num.RefIntFld8)
		require.Equal(t, *num.RefIntFld8, int8(33))
		require.NotNil(t, *num.RefUintFld8)
		require.Equal(t, *num.RefUintFld8, uint8(45))
	})
	t.Run("null_filled_part", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 12345, "int_fld8": 25, "int_fld16": 32767, "int_fld32": 12345, "int_fld64": 12345, "Uint_fld": 12345, "Uint_fld8": 12, "Uint_fld16": 12345, "Uint_fld32": 12345, "Uint_fld64": 12345, "fl23": 12345, "fl64": 12345, "ref_int_fld": null, "ref_int_fld8": null, "ref_int_fld16": null, "ref_int_fld64": null, "ref_Uint_fld": null, "ref_Uint_fld8": null, "ref_Uint_fld16": null, "ref_Uint_fld32": null, "ref_Uint_fld64": null, "ref_fl23": null, "ref_fl64": null}`))
		require.NoError(t, err)
		require.Nil(t, num.RefIntFld8)
		require.Nil(t, num.RefUintFld8)
		require.Nil(t, num.RefIntFld16)
		require.Nil(t, num.RefUintFld16)
		require.NotNil(t, num.RefIntFld32)
		require.Nil(t, num.RefUintFld32)
		require.Nil(t, num.RefIntFld64)
		require.Nil(t, num.RefUintFld64)
	})
	t.Run("overflow-int8", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld8": 98765}`))
		require.Error(t, err)
	})
	t.Run("overflow-uint8", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld8": 345}`))
		require.Error(t, err)
	})
	t.Run("overflow-int16", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld16": 45633}`))
		require.Error(t, err)
	})
	t.Run("overflow-uint16", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld16": 65536}`))
		require.Error(t, err)
	})
	t.Run("overflow-int32", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld32": 2147483648}`))
		require.Error(t, err)
	})
	t.Run("overflow-uint32", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld32": 4294967296}`))
		require.Error(t, err)
	})
	t.Run("overflow-float32", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "fl23": 3.40282346638528859811704183484516925440e+38}`))
		require.NoError(t, err)
	})
	t.Run("overflow-float32", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "fl23": 4.40282346638528859811704183484516925440e+38}`))
		require.Error(t, err)
	})
	t.Run("nan-int8", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld8": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-uint8", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld8": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-int16", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld16": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-uint16", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld16": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-int32", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld32": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-int64", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld64": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-uint32", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld32": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-uint64", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld64": "0"}`))
		require.Error(t, err)
	})
	t.Run("float-uint32", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld32": 23.56}`))
		require.Error(t, err)
	})
	t.Run("negative-uint", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"Uint_fld": -12}`))
		require.Error(t, err)
	})
	t.Run("inf-float64", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "fl64": 1.8976931348623157E308}`))
		require.NoError(t, err)
		require.Equal(t, num.FloatFld64, math.Inf(1))
	})
	t.Run("double-field", func(t *testing.T) {
		var num NumStruct01
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld": 1200}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "the 'int_fld' field appears in the object twice")
	})
}

func Test_NumStruct01_Marshal(t *testing.T) {
	t.Run("filled-nonref", func(t *testing.T) {
		var test = NumStruct01{
			IntFld:     -12,
			IntFld8:    78,
			IntFld16:   -133,
			IntFld32:   65536,
			IntFld64:   -1,
			UintFld:    0,
			UintFld8:   255,
			UintFld16:  6780,
			UintFld32:  4294967295,
			UintFld64:  4294967296,
			FloatFld32: 123.00,
			FloatFld64: -456.123,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		var expected = `{
"int_fld": -12, "int_fld8": 78, "int_fld16": -133, "int_fld32": 65536, "int_fld64": -1,
"Uint_fld": 0, "Uint_fld8": 255, "Uint_fld16": 6780, "Uint_fld32": 4294967295, "Uint_fld64": 4294967296, 
"fl23": 123.00, "fl64": -456.123,
"ref_int_fld": null, "ref_int_fld16": null, "ref_int_fld32": null, "ref_int_fld64": null,
"ref_Uint_fld": null, "ref_Uint_fld16": null, "ref_Uint_fld32": null, "ref_Uint_fld64": null,
"ref_fl23": null, "ref_fl64": null}`
		require.JSONEq(t, expected, string(b))
	})
	t.Run("filled-refs", func(t *testing.T) {
		var test = NumStruct01{
			RefIntFld:     ref(133),
			RefIntFld8:    ref(int8(-4)),
			RefIntFld16:   ref(int16(0)),
			RefIntFld32:   ref(int32(2147483647)),
			RefIntFld64:   ref(int64(-999)),
			RefUintFld:    ref(uint(0)),
			RefUintFld8:   ref(uint8(254)),
			RefUintFld16:  ref(uint16(65535)),
			RefUintFld32:  ref(uint32(65537)),
			RefUintFld64:  ref(uint64(1)),
			RefFloatFld32: ref(float32(-1)),
			RefFloatFld64: ref(float64(-1)),
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		var expected = `{
"int_fld": 0, "int_fld16": 0, "int_fld32": 0, "int_fld64": 0,
"Uint_fld": 0, "Uint_fld16": 0, "Uint_fld32": 0, "Uint_fld64": 0, 
"fl23": 0, "fl64": 0,
"ref_int_fld": 133, "ref_int_fld8": -4, "ref_int_fld16": 0, "ref_int_fld32": 2147483647, "ref_int_fld64": -999,
"ref_Uint_fld": 0, "ref_Uint_fld8": 254, "ref_Uint_fld16": 65535, "ref_Uint_fld32": 65537, "ref_Uint_fld64": 1,
"ref_fl23": -1, "ref_fl64": -1}`
		require.JSONEq(t, expected, string(b))
	})
}

func ref[T any](a T) *T {
	return &a
}

func Test_NumStruct02_Unmarshal(t *testing.T) {
	t.Run("empty_struct", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Zero(t, num.IntFld)
		require.Zero(t, num.IntFld8)
		require.Zero(t, num.IntFld16)
		require.NotZero(t, num.IntFld32)
		require.Zero(t, num.IntFld64)
		require.Zero(t, num.UintFld)
		require.Zero(t, num.UintFld8)
		require.Zero(t, num.UintFld16)
		require.Zero(t, num.UintFld32)
		require.Zero(t, num.UintFld64)
		require.Zero(t, num.FloatFld32)
		require.Zero(t, num.FloatFld64)
		require.Nil(t, num.RefIntFld)
		require.Nil(t, num.RefIntFld8)
		require.Nil(t, num.RefIntFld16)
		require.Nil(t, num.RefIntFld32)
		require.Nil(t, num.RefIntFld64)
		require.Nil(t, num.RefUintFld)
		require.Nil(t, num.RefUintFld8)
		require.Nil(t, num.RefUintFld16)
		require.Nil(t, num.RefUintFld32)
		require.Nil(t, num.RefUintFld64)
		require.NotNil(t, num.RefFloatFld32)
		require.Nil(t, num.RefFloatFld64)
	})
	t.Run("fulfilled_struct", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 12345, "int_fld8": 25, "int_fld16": 32767, "int_fld32": 12345, "int_fld64": 12345, "Uint_fld": 12345, "Uint_fld8": 12, "Uint_fld16": 12345, "Uint_fld32": 12345, "Uint_fld64": 12345, "fl23": 12345, "fl64": 12345, "ref_int_fld": 12345, "ref_int_fld8": 33, "ref_int_fld16": 12345, "ref_int_fld32": 12345, "ref_int_fld64": 12345, "ref_Uint_fld": 12345, "ref_Uint_fld8": 45, "ref_Uint_fld16": 12345, "ref_Uint_fld32": 12345, "ref_Uint_fld64": 12345, "ref_fl23": 12345, "ref_fl64": 12345}`))
		require.NoError(t, err)
		require.Equal(t, num.IntFld, 12345)
		require.Equal(t, num.IntFld8, int8(25))
		require.Equal(t, num.UintFld8, uint8(12))
		require.NotNil(t, *num.RefIntFld8)
		require.Equal(t, *num.RefIntFld8, int8(33))
		require.NotNil(t, *num.RefUintFld8)
		require.Equal(t, *num.RefUintFld8, uint8(45))
	})
	t.Run("null_filled_part", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 12345, "int_fld8": 25, "int_fld16": 32767, "int_fld32": 12345, "int_fld64": 12345, "Uint_fld": 12345, "Uint_fld8": 12, "Uint_fld16": 12345, "Uint_fld32": 12345, "Uint_fld64": 12345, "fl23": 12345, "fl64": 12345, "ref_int_fld": null, "ref_int_fld8": null, "ref_int_fld16": null, "ref_int_fld32": null, "ref_int_fld64": null, "ref_Uint_fld": null, "ref_Uint_fld8": null, "ref_Uint_fld16": null, "ref_Uint_fld32": null, "ref_Uint_fld64": null, "ref_fl23": null, "ref_fl64": null}`))
		require.NoError(t, err)
		require.Nil(t, num.RefIntFld8)
		require.Nil(t, num.RefUintFld8)
		require.Nil(t, num.RefIntFld16)
		require.Nil(t, num.RefUintFld16)
		require.Nil(t, num.RefIntFld32)
		require.Nil(t, num.RefUintFld32)
		require.Nil(t, num.RefIntFld64)
		require.Nil(t, num.RefUintFld64)
	})
	t.Run("overflow-int8", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld8": 98765}`))
		require.Error(t, err)
	})
	t.Run("overflow-uint8", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld8": 345}`))
		require.Error(t, err)
	})
	t.Run("overflow-int16", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld16": 45633}`))
		require.Error(t, err)
	})
	t.Run("overflow-uint16", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld16": 65536}`))
		require.Error(t, err)
	})
	t.Run("overflow-int32", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld32": 2147483648}`))
		require.Error(t, err)
	})
	t.Run("overflow-uint32", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld32": 4294967296}`))
		require.Error(t, err)
	})
	t.Run("overflow-float32", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "fl23": 3.40282346638528859811704183484516925440e+38}`))
		require.NoError(t, err)
	})
	t.Run("overflow-float32", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "fl23": 4.40282346638528859811704183484516925440e+38}`))
		require.Error(t, err)
	})
	t.Run("nan-int8", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld8": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-uint8", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld8": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-int16", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld16": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-uint16", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld16": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-int32", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld32": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-int64", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld64": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-uint32", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld32": "0"}`))
		require.Error(t, err)
	})
	t.Run("nan-uint64", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld64": "0"}`))
		require.Error(t, err)
	})
	t.Run("float-uint32", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "Uint_fld32": 23.56}`))
		require.Error(t, err)
	})
	t.Run("negative-uint", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"Uint_fld": -12}`))
		require.Error(t, err)
	})
	t.Run("inf-float64", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "fl64": 1.8976931348623157E308}`))
		require.NoError(t, err)
		require.Equal(t, num.FloatFld64, math.Inf(1))
	})
	t.Run("double-field", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld": 1200}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "the 'int_fld' field appears in the object twice")
	})
	t.Run("extra-field", func(t *testing.T) {
		var num NumStruct02
		err := num.UnmarshalJSON([]byte(`{"int_fld": 1200, "int_fld_extr": 1200}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "unexpected field 'int_fld_extr'")
	})
}

func TestNumStruct_IsZero(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var test1 NumStruct01
		require.True(t, test1.IsZero())
		var test2 NumStruct02
		require.True(t, test2.IsZero())
	})
	t.Run("not zero", func(t *testing.T) {
		require.False(t, NumStruct01{
			IntFld: 1,
		}.IsZero())
		require.False(t, NumStruct01{
			IntFld8: 2,
		}.IsZero())
		require.False(t, NumStruct01{
			IntFld16: 3,
		}.IsZero())
		require.False(t, NumStruct01{
			IntFld32: 4,
		}.IsZero())
		require.False(t, NumStruct01{
			IntFld64: 5,
		}.IsZero())
		require.False(t, NumStruct01{
			UintFld: 6,
		}.IsZero())
		require.False(t, NumStruct01{
			UintFld8: 7,
		}.IsZero())
		require.False(t, NumStruct01{
			UintFld16: 8,
		}.IsZero())
		require.False(t, NumStruct01{
			UintFld32: 9,
		}.IsZero())
		require.False(t, NumStruct01{
			UintFld64: 10,
		}.IsZero())
		require.False(t, NumStruct01{
			FloatFld32: 11,
		}.IsZero())
		require.False(t, NumStruct01{
			FloatFld64: 12,
		}.IsZero())
		require.False(t, NumStruct01{
			RefIntFld: ref(0),
		}.IsZero())
		require.False(t, NumStruct01{
			RefIntFld8: ref(int8(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefIntFld16: ref(int16(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefIntFld32: ref(int32(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefIntFld64: ref(int64(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefUintFld: ref(uint(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefUintFld8: ref(uint8(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefUintFld16: ref(uint16(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefUintFld32: ref(uint32(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefUintFld64: ref(uint64(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefFloatFld32: ref(float32(0)),
		}.IsZero())
		require.False(t, NumStruct01{
			RefFloatFld64: ref(float64(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			IntFld: 1,
		}.IsZero())
		require.False(t, NumStruct02{
			IntFld8: 2,
		}.IsZero())
		require.False(t, NumStruct02{
			IntFld16: 3,
		}.IsZero())
		require.False(t, NumStruct02{
			IntFld32: 4,
		}.IsZero())
		require.False(t, NumStruct02{
			IntFld64: 5,
		}.IsZero())
		require.False(t, NumStruct02{
			UintFld: 6,
		}.IsZero())
		require.False(t, NumStruct02{
			UintFld8: 7,
		}.IsZero())
		require.False(t, NumStruct02{
			UintFld16: 8,
		}.IsZero())
		require.False(t, NumStruct02{
			UintFld32: 9,
		}.IsZero())
		require.False(t, NumStruct02{
			UintFld64: 10,
		}.IsZero())
		require.False(t, NumStruct02{
			FloatFld32: 11,
		}.IsZero())
		require.False(t, NumStruct02{
			FloatFld64: 12,
		}.IsZero())
		require.False(t, NumStruct02{
			RefIntFld: ref(0),
		}.IsZero())
		require.False(t, NumStruct02{
			RefIntFld8: ref(int8(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefIntFld16: ref(int16(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefIntFld32: ref(int32(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefIntFld64: ref(int64(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefUintFld: ref(uint(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefUintFld8: ref(uint8(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefUintFld16: ref(uint16(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefUintFld32: ref(uint32(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefUintFld64: ref(uint64(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefFloatFld32: ref(float32(0)),
		}.IsZero())
		require.False(t, NumStruct02{
			RefFloatFld64: ref(float64(0)),
		}.IsZero())
	})
}

func Test_NumStruct02_Marshal(t *testing.T) {
	t.Run("filled-nonref", func(t *testing.T) {
		var test = NumStruct02{
			IntFld:     -12,
			IntFld8:    78,
			IntFld16:   -133,
			IntFld32:   65536,
			IntFld64:   -1,
			UintFld:    50,
			UintFld8:   255,
			UintFld16:  6780,
			UintFld32:  4294967295,
			UintFld64:  4294967296,
			FloatFld32: 123.00,
			FloatFld64: -456.123,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		var expected = `{
"int_fld": -12, "int_fld8": 78, "int_fld16": -133, "int_fld32": 65536, "int_fld64": -1,
"Uint_fld": 50, "Uint_fld8": 255, "Uint_fld16": 6780, "Uint_fld32": 4294967295, "Uint_fld64": 4294967296, 
"fl23": 123.00, "fl64": -456.123,
"ref_int_fld": null, "ref_int_fld16": null, "ref_int_fld32": null, "ref_int_fld64": null,
"ref_Uint_fld": null, "ref_Uint_fld16": null, "ref_Uint_fld32": null, "ref_Uint_fld64": null,
"ref_fl23": null, "ref_fl64": null}`
		require.JSONEq(t, expected, string(b))
	})
	t.Run("filled-refs", func(t *testing.T) {
		var test = NumStruct02{
			RefIntFld:     ref(133),
			RefIntFld8:    ref(int8(-4)),
			RefIntFld16:   ref(int16(0)),
			RefIntFld32:   ref(int32(2147483647)),
			RefIntFld64:   ref(int64(-999)),
			RefUintFld:    ref(uint(0)),
			RefUintFld8:   ref(uint8(254)),
			RefUintFld16:  ref(uint16(65535)),
			RefUintFld32:  ref(uint32(65537)),
			RefUintFld64:  ref(uint64(1)),
			RefFloatFld32: ref(float32(-1)),
			RefFloatFld64: ref(float64(-1)),
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		var expected = `{
"int_fld": 0, "int_fld16": 0, "int_fld32": 0, "int_fld64": 0,
"Uint_fld": 0, "Uint_fld16": 0, "Uint_fld32": 0, "Uint_fld64": 0, 
"fl23": 0, "fl64": 0,
"ref_int_fld": 133, "ref_int_fld8": -4, "ref_int_fld16": 0, "ref_int_fld32": 2147483647, "ref_int_fld64": -999,
"ref_Uint_fld": 0, "ref_Uint_fld8": 254, "ref_Uint_fld16": 65535, "ref_Uint_fld32": 65537, "ref_Uint_fld64": 1,
"ref_fl23": -1, "ref_fl64": -1}`
		require.JSONEq(t, expected, string(b))
	})
}
