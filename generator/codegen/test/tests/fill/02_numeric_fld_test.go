package testo

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestNumStruct01(t *testing.T) {
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
}

func TestNumStruct02(t *testing.T) {
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
