// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_num

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/valyala/fastjson"
)

// jsonParserNumStruct01 used for pooling Parsers for NumStruct01 JSONs.
var jsonParserNumStruct01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *NumStruct01) UnmarshalJSON(data []byte) error {
	parser := jsonParserNumStruct01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserNumStruct01.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *NumStruct01) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _intFld := v.Get("int_fld"); _intFld != nil {
		var valIntFld int
		valIntFld, err = _intFld.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld' value: %w", objPath, err)
		}
		s.IntFld = valIntFld
	}
	if _intFld8 := v.Get("int_fld8"); _intFld8 != nil {
		var valIntFld8 int
		valIntFld8, err = _intFld8.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld8' value: %w", objPath, err)
		}
		if valIntFld8 > math.MaxInt8 {
			return fmt.Errorf("error parsing '%sint_fld8' value %d exceeds maximum for data type int8", objPath, valIntFld8)
		}
		s.IntFld8 = int8(valIntFld8)
	}
	if _intFld16 := v.Get("int_fld16"); _intFld16 != nil {
		var valIntFld16 int
		valIntFld16, err = _intFld16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld16' value: %w", objPath, err)
		}
		if valIntFld16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sint_fld16' value %d exceeds maximum for data type int16", objPath, valIntFld16)
		}
		s.IntFld16 = int16(valIntFld16)
	}
	if _intFld32 := v.Get("int_fld32"); _intFld32 != nil {
		var valIntFld32 int
		valIntFld32, err = _intFld32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld32' value: %w", objPath, err)
		}
		if valIntFld32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sint_fld32' value %d exceeds maximum for data type int32", objPath, valIntFld32)
		}
		s.IntFld32 = int32(valIntFld32)
	}
	if _intFld64 := v.Get("int_fld64"); _intFld64 != nil {
		var valIntFld64 int64
		valIntFld64, err = _intFld64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld64' value: %w", objPath, err)
		}
		s.IntFld64 = valIntFld64
	}
	if _uintFld := v.Get("Uint_fld"); _uintFld != nil {
		var valUintFld uint
		valUintFld, err = _uintFld.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld' value: %w", objPath, err)
		}
		s.UintFld = valUintFld
	}
	if _uintFld8 := v.Get("Uint_fld8"); _uintFld8 != nil {
		var valUintFld8 uint
		valUintFld8, err = _uintFld8.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld8' value: %w", objPath, err)
		}
		if valUintFld8 > math.MaxUint8 {
			return fmt.Errorf("error parsing '%sUint_fld8' value %d exceeds maximum for data type uint8", objPath, valUintFld8)
		}
		s.UintFld8 = uint8(valUintFld8)
	}
	if _uintFld16 := v.Get("Uint_fld16"); _uintFld16 != nil {
		var valUintFld16 uint
		valUintFld16, err = _uintFld16.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld16' value: %w", objPath, err)
		}
		if valUintFld16 > math.MaxUint16 {
			return fmt.Errorf("error parsing '%sUint_fld16' value %d exceeds maximum for data type uint16", objPath, valUintFld16)
		}
		s.UintFld16 = uint16(valUintFld16)
	} else {
		s.UintFld16 = 333
	}
	if _uintFld32 := v.Get("Uint_fld32"); _uintFld32 != nil {
		var valUintFld32 uint
		valUintFld32, err = _uintFld32.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld32' value: %w", objPath, err)
		}
		if valUintFld32 > math.MaxUint32 {
			return fmt.Errorf("error parsing '%sUint_fld32' value %d exceeds maximum for data type uint32", objPath, valUintFld32)
		}
		s.UintFld32 = uint32(valUintFld32)
	}
	if _uintFld64 := v.Get("Uint_fld64"); _uintFld64 != nil {
		var valUintFld64 uint64
		valUintFld64, err = _uintFld64.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld64' value: %w", objPath, err)
		}
		s.UintFld64 = valUintFld64
	}
	if _floatFld32 := v.Get("fl23"); _floatFld32 != nil {
		var valFloatFld32 float64
		valFloatFld32, err = _floatFld32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfl23' value: %w", objPath, err)
		}
		if valFloatFld32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%sfl23' value %f exceeds maximum for data type float32", objPath, valFloatFld32)
		}
		s.FloatFld32 = float32(valFloatFld32)
	}
	if _floatFld64 := v.Get("fl64"); _floatFld64 != nil {
		var valFloatFld64 float64
		valFloatFld64, err = _floatFld64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfl64' value: %w", objPath, err)
		}
		s.FloatFld64 = valFloatFld64
	}
	if _refIntFld := v.Get("ref_int_fld"); valueIsNotNull(_refIntFld) {
		var valRefIntFld int
		valRefIntFld, err = _refIntFld.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld' value: %w", objPath, err)
		}
		s.RefIntFld = &valRefIntFld
	}
	if _refIntFld8 := v.Get("ref_int_fld8"); valueIsNotNull(_refIntFld8) {
		var valRefIntFld8 int
		valRefIntFld8, err = _refIntFld8.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld8' value: %w", objPath, err)
		}
		if valRefIntFld8 > math.MaxInt8 {
			return fmt.Errorf("error parsing '%sref_int_fld8' value %d exceeds maximum for data type int8", objPath, valRefIntFld8)
		}
		s.RefIntFld8 = new(int8)
		*s.RefIntFld8 = int8(valRefIntFld8)
	}
	if _refIntFld16 := v.Get("ref_int_fld16"); valueIsNotNull(_refIntFld16) {
		var valRefIntFld16 int
		valRefIntFld16, err = _refIntFld16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld16' value: %w", objPath, err)
		}
		if valRefIntFld16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sref_int_fld16' value %d exceeds maximum for data type int16", objPath, valRefIntFld16)
		}
		s.RefIntFld16 = new(int16)
		*s.RefIntFld16 = int16(valRefIntFld16)
	}
	if _refIntFld32 := v.Get("ref_int_fld32"); valueIsNotNull(_refIntFld32) {
		var valRefIntFld32 int
		valRefIntFld32, err = _refIntFld32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld32' value: %w", objPath, err)
		}
		if valRefIntFld32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sref_int_fld32' value %d exceeds maximum for data type int32", objPath, valRefIntFld32)
		}
		s.RefIntFld32 = new(int32)
		*s.RefIntFld32 = int32(valRefIntFld32)
	} else {
		if _refIntFld32 == nil {
			var __RefIntFld32 int32 = 456
			s.RefIntFld32 = &__RefIntFld32
		}
	}
	if _refIntFld64 := v.Get("ref_int_fld64"); valueIsNotNull(_refIntFld64) {
		var valRefIntFld64 int64
		valRefIntFld64, err = _refIntFld64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld64' value: %w", objPath, err)
		}
		s.RefIntFld64 = &valRefIntFld64
	}
	if _refUintFld := v.Get("ref_Uint_fld"); valueIsNotNull(_refUintFld) {
		var valRefUintFld uint
		valRefUintFld, err = _refUintFld.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld' value: %w", objPath, err)
		}
		s.RefUintFld = new(uint)
		*s.RefUintFld = uint(valRefUintFld)
	}
	if _refUintFld8 := v.Get("ref_Uint_fld8"); valueIsNotNull(_refUintFld8) {
		var valRefUintFld8 uint
		valRefUintFld8, err = _refUintFld8.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld8' value: %w", objPath, err)
		}
		if valRefUintFld8 > math.MaxUint8 {
			return fmt.Errorf("error parsing '%sref_Uint_fld8' value %d exceeds maximum for data type uint8", objPath, valRefUintFld8)
		}
		s.RefUintFld8 = new(uint8)
		*s.RefUintFld8 = uint8(valRefUintFld8)
	}
	if _refUintFld16 := v.Get("ref_Uint_fld16"); valueIsNotNull(_refUintFld16) {
		var valRefUintFld16 uint
		valRefUintFld16, err = _refUintFld16.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld16' value: %w", objPath, err)
		}
		if valRefUintFld16 > math.MaxUint16 {
			return fmt.Errorf("error parsing '%sref_Uint_fld16' value %d exceeds maximum for data type uint16", objPath, valRefUintFld16)
		}
		s.RefUintFld16 = new(uint16)
		*s.RefUintFld16 = uint16(valRefUintFld16)
	}
	if _refUintFld32 := v.Get("ref_Uint_fld32"); valueIsNotNull(_refUintFld32) {
		var valRefUintFld32 uint
		valRefUintFld32, err = _refUintFld32.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld32' value: %w", objPath, err)
		}
		if valRefUintFld32 > math.MaxUint32 {
			return fmt.Errorf("error parsing '%sref_Uint_fld32' value %d exceeds maximum for data type uint32", objPath, valRefUintFld32)
		}
		s.RefUintFld32 = new(uint32)
		*s.RefUintFld32 = uint32(valRefUintFld32)
	}
	if _refUintFld64 := v.Get("ref_Uint_fld64"); valueIsNotNull(_refUintFld64) {
		var valRefUintFld64 uint64
		valRefUintFld64, err = _refUintFld64.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld64' value: %w", objPath, err)
		}
		s.RefUintFld64 = new(uint64)
		*s.RefUintFld64 = uint64(valRefUintFld64)
	}
	if _refFloatFld32 := v.Get("ref_fl23"); valueIsNotNull(_refFloatFld32) {
		var valRefFloatFld32 float64
		valRefFloatFld32, err = _refFloatFld32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_fl23' value: %w", objPath, err)
		}
		if valRefFloatFld32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%sref_fl23' value %f exceeds maximum for data type float32", objPath, valRefFloatFld32)
		}
		s.RefFloatFld32 = new(float32)
		*s.RefFloatFld32 = float32(valRefFloatFld32)
	}
	if _refFloatFld64 := v.Get("ref_fl64"); valueIsNotNull(_refFloatFld64) {
		var valRefFloatFld64 float64
		valRefFloatFld64, err = _refFloatFld64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_fl64' value: %w", objPath, err)
		}
		s.RefFloatFld64 = &valRefFloatFld64
	}
	return nil
}

// validate checks for correct data structure
func (s *NumStruct01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [24]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd', '8'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd', '1', '6'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd', '3', '2'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd', '6', '4'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '8'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '1', '6'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '3', '2'}) {
			checkFields[8]++
			if checkFields[8] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '6', '4'}) {
			checkFields[9]++
			if checkFields[9] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'l', '2', '3'}) {
			checkFields[10]++
			if checkFields[10] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'l', '6', '4'}) {
			checkFields[11]++
			if checkFields[11] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd'}) {
			checkFields[12]++
			if checkFields[12] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd', '8'}) {
			checkFields[13]++
			if checkFields[13] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd', '1', '6'}) {
			checkFields[14]++
			if checkFields[14] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd', '3', '2'}) {
			checkFields[15]++
			if checkFields[15] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd', '6', '4'}) {
			checkFields[16]++
			if checkFields[16] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd'}) {
			checkFields[17]++
			if checkFields[17] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '8'}) {
			checkFields[18]++
			if checkFields[18] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '1', '6'}) {
			checkFields[19]++
			if checkFields[19] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '3', '2'}) {
			checkFields[20]++
			if checkFields[20] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '6', '4'}) {
			checkFields[21]++
			if checkFields[21] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'f', 'l', '2', '3'}) {
			checkFields[22]++
			if checkFields[22] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'f', 'l', '6', '4'}) {
			checkFields[23]++
			if checkFields[23] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}

// jsonParserNumStruct02 used for pooling Parsers for NumStruct02 JSONs.
var jsonParserNumStruct02 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *NumStruct02) UnmarshalJSON(data []byte) error {
	parser := jsonParserNumStruct02.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserNumStruct02.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *NumStruct02) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _intFld := v.Get("int_fld"); _intFld != nil {
		var valIntFld int
		valIntFld, err = _intFld.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld' value: %w", objPath, err)
		}
		s.IntFld = valIntFld
	}
	if _intFld8 := v.Get("int_fld8"); _intFld8 != nil {
		var valIntFld8 int
		valIntFld8, err = _intFld8.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld8' value: %w", objPath, err)
		}
		if valIntFld8 > math.MaxInt8 {
			return fmt.Errorf("error parsing '%sint_fld8' value %d exceeds maximum for data type int8", objPath, valIntFld8)
		}
		s.IntFld8 = int8(valIntFld8)
	}
	if _intFld16 := v.Get("int_fld16"); _intFld16 != nil {
		var valIntFld16 int
		valIntFld16, err = _intFld16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld16' value: %w", objPath, err)
		}
		if valIntFld16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sint_fld16' value %d exceeds maximum for data type int16", objPath, valIntFld16)
		}
		s.IntFld16 = int16(valIntFld16)
	}
	if _intFld32 := v.Get("int_fld32"); _intFld32 != nil {
		var valIntFld32 int
		valIntFld32, err = _intFld32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld32' value: %w", objPath, err)
		}
		if valIntFld32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sint_fld32' value %d exceeds maximum for data type int32", objPath, valIntFld32)
		}
		s.IntFld32 = int32(valIntFld32)
	} else {
		s.IntFld32 = 16
	}
	if _intFld64 := v.Get("int_fld64"); _intFld64 != nil {
		var valIntFld64 int64
		valIntFld64, err = _intFld64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld64' value: %w", objPath, err)
		}
		s.IntFld64 = valIntFld64
	}
	if _uintFld := v.Get("Uint_fld"); _uintFld != nil {
		var valUintFld uint
		valUintFld, err = _uintFld.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld' value: %w", objPath, err)
		}
		s.UintFld = valUintFld
	}
	if _uintFld8 := v.Get("Uint_fld8"); _uintFld8 != nil {
		var valUintFld8 uint
		valUintFld8, err = _uintFld8.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld8' value: %w", objPath, err)
		}
		if valUintFld8 > math.MaxUint8 {
			return fmt.Errorf("error parsing '%sUint_fld8' value %d exceeds maximum for data type uint8", objPath, valUintFld8)
		}
		s.UintFld8 = uint8(valUintFld8)
	}
	if _uintFld16 := v.Get("Uint_fld16"); _uintFld16 != nil {
		var valUintFld16 uint
		valUintFld16, err = _uintFld16.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld16' value: %w", objPath, err)
		}
		if valUintFld16 > math.MaxUint16 {
			return fmt.Errorf("error parsing '%sUint_fld16' value %d exceeds maximum for data type uint16", objPath, valUintFld16)
		}
		s.UintFld16 = uint16(valUintFld16)
	}
	if _uintFld32 := v.Get("Uint_fld32"); _uintFld32 != nil {
		var valUintFld32 uint
		valUintFld32, err = _uintFld32.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld32' value: %w", objPath, err)
		}
		if valUintFld32 > math.MaxUint32 {
			return fmt.Errorf("error parsing '%sUint_fld32' value %d exceeds maximum for data type uint32", objPath, valUintFld32)
		}
		s.UintFld32 = uint32(valUintFld32)
	}
	if _uintFld64 := v.Get("Uint_fld64"); _uintFld64 != nil {
		var valUintFld64 uint64
		valUintFld64, err = _uintFld64.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld64' value: %w", objPath, err)
		}
		s.UintFld64 = valUintFld64
	}
	if _floatFld32 := v.Get("fl23"); _floatFld32 != nil {
		var valFloatFld32 float64
		valFloatFld32, err = _floatFld32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfl23' value: %w", objPath, err)
		}
		if valFloatFld32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%sfl23' value %f exceeds maximum for data type float32", objPath, valFloatFld32)
		}
		s.FloatFld32 = float32(valFloatFld32)
	}
	if _floatFld64 := v.Get("fl64"); _floatFld64 != nil {
		var valFloatFld64 float64
		valFloatFld64, err = _floatFld64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfl64' value: %w", objPath, err)
		}
		s.FloatFld64 = valFloatFld64
	}
	if _refIntFld := v.Get("ref_int_fld"); valueIsNotNull(_refIntFld) {
		var valRefIntFld int
		valRefIntFld, err = _refIntFld.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld' value: %w", objPath, err)
		}
		s.RefIntFld = &valRefIntFld
	}
	if _refIntFld8 := v.Get("ref_int_fld8"); valueIsNotNull(_refIntFld8) {
		var valRefIntFld8 int
		valRefIntFld8, err = _refIntFld8.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld8' value: %w", objPath, err)
		}
		if valRefIntFld8 > math.MaxInt8 {
			return fmt.Errorf("error parsing '%sref_int_fld8' value %d exceeds maximum for data type int8", objPath, valRefIntFld8)
		}
		s.RefIntFld8 = new(int8)
		*s.RefIntFld8 = int8(valRefIntFld8)
	}
	if _refIntFld16 := v.Get("ref_int_fld16"); valueIsNotNull(_refIntFld16) {
		var valRefIntFld16 int
		valRefIntFld16, err = _refIntFld16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld16' value: %w", objPath, err)
		}
		if valRefIntFld16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sref_int_fld16' value %d exceeds maximum for data type int16", objPath, valRefIntFld16)
		}
		s.RefIntFld16 = new(int16)
		*s.RefIntFld16 = int16(valRefIntFld16)
	}
	if _refIntFld32 := v.Get("ref_int_fld32"); valueIsNotNull(_refIntFld32) {
		var valRefIntFld32 int
		valRefIntFld32, err = _refIntFld32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld32' value: %w", objPath, err)
		}
		if valRefIntFld32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sref_int_fld32' value %d exceeds maximum for data type int32", objPath, valRefIntFld32)
		}
		s.RefIntFld32 = new(int32)
		*s.RefIntFld32 = int32(valRefIntFld32)
	}
	if _refIntFld64 := v.Get("ref_int_fld64"); valueIsNotNull(_refIntFld64) {
		var valRefIntFld64 int64
		valRefIntFld64, err = _refIntFld64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld64' value: %w", objPath, err)
		}
		s.RefIntFld64 = &valRefIntFld64
	}
	if _refUintFld := v.Get("ref_Uint_fld"); valueIsNotNull(_refUintFld) {
		var valRefUintFld uint
		valRefUintFld, err = _refUintFld.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld' value: %w", objPath, err)
		}
		s.RefUintFld = new(uint)
		*s.RefUintFld = uint(valRefUintFld)
	}
	if _refUintFld8 := v.Get("ref_Uint_fld8"); valueIsNotNull(_refUintFld8) {
		var valRefUintFld8 uint
		valRefUintFld8, err = _refUintFld8.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld8' value: %w", objPath, err)
		}
		if valRefUintFld8 > math.MaxUint8 {
			return fmt.Errorf("error parsing '%sref_Uint_fld8' value %d exceeds maximum for data type uint8", objPath, valRefUintFld8)
		}
		s.RefUintFld8 = new(uint8)
		*s.RefUintFld8 = uint8(valRefUintFld8)
	}
	if _refUintFld16 := v.Get("ref_Uint_fld16"); valueIsNotNull(_refUintFld16) {
		var valRefUintFld16 uint
		valRefUintFld16, err = _refUintFld16.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld16' value: %w", objPath, err)
		}
		if valRefUintFld16 > math.MaxUint16 {
			return fmt.Errorf("error parsing '%sref_Uint_fld16' value %d exceeds maximum for data type uint16", objPath, valRefUintFld16)
		}
		s.RefUintFld16 = new(uint16)
		*s.RefUintFld16 = uint16(valRefUintFld16)
	}
	if _refUintFld32 := v.Get("ref_Uint_fld32"); valueIsNotNull(_refUintFld32) {
		var valRefUintFld32 uint
		valRefUintFld32, err = _refUintFld32.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld32' value: %w", objPath, err)
		}
		if valRefUintFld32 > math.MaxUint32 {
			return fmt.Errorf("error parsing '%sref_Uint_fld32' value %d exceeds maximum for data type uint32", objPath, valRefUintFld32)
		}
		s.RefUintFld32 = new(uint32)
		*s.RefUintFld32 = uint32(valRefUintFld32)
	}
	if _refUintFld64 := v.Get("ref_Uint_fld64"); valueIsNotNull(_refUintFld64) {
		var valRefUintFld64 uint64
		valRefUintFld64, err = _refUintFld64.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld64' value: %w", objPath, err)
		}
		s.RefUintFld64 = new(uint64)
		*s.RefUintFld64 = uint64(valRefUintFld64)
	}
	if _refFloatFld32 := v.Get("ref_fl23"); valueIsNotNull(_refFloatFld32) {
		var valRefFloatFld32 float64
		valRefFloatFld32, err = _refFloatFld32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_fl23' value: %w", objPath, err)
		}
		if valRefFloatFld32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%sref_fl23' value %f exceeds maximum for data type float32", objPath, valRefFloatFld32)
		}
		s.RefFloatFld32 = new(float32)
		*s.RefFloatFld32 = float32(valRefFloatFld32)
	} else {
		if _refFloatFld32 == nil {
			var __RefFloatFld32 float32 = 1.234
			s.RefFloatFld32 = &__RefFloatFld32
		}
	}
	if _refFloatFld64 := v.Get("ref_fl64"); valueIsNotNull(_refFloatFld64) {
		var valRefFloatFld64 float64
		valRefFloatFld64, err = _refFloatFld64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_fl64' value: %w", objPath, err)
		}
		s.RefFloatFld64 = &valRefFloatFld64
	}
	return nil
}

// validate checks for correct data structure
func (s *NumStruct02) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [24]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd', '8'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd', '1', '6'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd', '3', '2'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', 'f', 'l', 'd', '6', '4'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '8'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '1', '6'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '3', '2'}) {
			checkFields[8]++
			if checkFields[8] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '6', '4'}) {
			checkFields[9]++
			if checkFields[9] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'l', '2', '3'}) {
			checkFields[10]++
			if checkFields[10] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'l', '6', '4'}) {
			checkFields[11]++
			if checkFields[11] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd'}) {
			checkFields[12]++
			if checkFields[12] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd', '8'}) {
			checkFields[13]++
			if checkFields[13] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd', '1', '6'}) {
			checkFields[14]++
			if checkFields[14] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd', '3', '2'}) {
			checkFields[15]++
			if checkFields[15] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'i', 'n', 't', '_', 'f', 'l', 'd', '6', '4'}) {
			checkFields[16]++
			if checkFields[16] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd'}) {
			checkFields[17]++
			if checkFields[17] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '8'}) {
			checkFields[18]++
			if checkFields[18] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '1', '6'}) {
			checkFields[19]++
			if checkFields[19] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '3', '2'}) {
			checkFields[20]++
			if checkFields[20] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'U', 'i', 'n', 't', '_', 'f', 'l', 'd', '6', '4'}) {
			checkFields[21]++
			if checkFields[21] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'f', 'l', '2', '3'}) {
			checkFields[22]++
			if checkFields[22] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', '_', 'f', 'l', '6', '4'}) {
			checkFields[23]++
			if checkFields[23] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *NumStruct01) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s NumStruct01) MarshalAppend(dst []byte) ([]byte, error) {
	var result = bytes.NewBuffer(dst)
	var (
		err error
		buf = make([]byte, 0, 128)
	)
	result.WriteRune('{')
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld != 0 {
		result.WriteString(`"int_fld":`)
		buf = strconv.AppendInt(buf[:0], int64(s.IntFld), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld":0`)
	}
	if s.IntFld8 != 0 {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"int_fld8":`)
		buf = strconv.AppendInt(buf[:0], int64(s.IntFld8), 10)
		result.Write(buf)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld16 != 0 {
		result.WriteString(`"int_fld16":`)
		buf = strconv.AppendInt(buf[:0], int64(s.IntFld16), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld16":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld32 != 0 {
		result.WriteString(`"int_fld32":`)
		buf = strconv.AppendInt(buf[:0], int64(s.IntFld32), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld32":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld64 != 0 {
		result.WriteString(`"int_fld64":`)
		buf = strconv.AppendInt(buf[:0], s.IntFld64, 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld64":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld != 0 {
		result.WriteString(`"Uint_fld":`)
		buf = strconv.AppendUint(buf[:0], uint64(s.UintFld), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld":0`)
	}
	if s.UintFld8 != 0 {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"Uint_fld8":`)
		buf = strconv.AppendUint(buf[:0], uint64(s.UintFld8), 10)
		result.Write(buf)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld16 != 0 {
		result.WriteString(`"Uint_fld16":`)
		buf = strconv.AppendUint(buf[:0], uint64(s.UintFld16), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld16":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld32 != 0 {
		result.WriteString(`"Uint_fld32":`)
		buf = strconv.AppendUint(buf[:0], uint64(s.UintFld32), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld32":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld64 != 0 {
		result.WriteString(`"Uint_fld64":`)
		buf = strconv.AppendUint(buf[:0], s.UintFld64, 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld64":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.FloatFld32 != 0 {
		result.WriteString(`"fl23":`)
		buf = strconv.AppendFloat(buf[:0], float64(s.FloatFld32), 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"fl23":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.FloatFld64 != 0 {
		result.WriteString(`"fl64":`)
		buf = strconv.AppendFloat(buf[:0], s.FloatFld64, 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"fl64":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld != nil {
		result.WriteString(`"ref_int_fld":`)
		buf = strconv.AppendInt(buf[:0], int64(*s.RefIntFld), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld":null`)
	}
	if s.RefIntFld8 != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"ref_int_fld8":`)
		buf = strconv.AppendInt(buf[:0], int64(*s.RefIntFld8), 10)
		result.Write(buf)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld16 != nil {
		result.WriteString(`"ref_int_fld16":`)
		buf = strconv.AppendInt(buf[:0], int64(*s.RefIntFld16), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld16":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld32 != nil {
		result.WriteString(`"ref_int_fld32":`)
		buf = strconv.AppendInt(buf[:0], int64(*s.RefIntFld32), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld32":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld64 != nil {
		result.WriteString(`"ref_int_fld64":`)
		buf = strconv.AppendInt(buf[:0], *s.RefIntFld64, 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld64":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld != nil {
		result.WriteString(`"ref_Uint_fld":`)
		buf = strconv.AppendUint(buf[:0], uint64(*s.RefUintFld), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld":null`)
	}
	if s.RefUintFld8 != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"ref_Uint_fld8":`)
		buf = strconv.AppendUint(buf[:0], uint64(*s.RefUintFld8), 10)
		result.Write(buf)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld16 != nil {
		result.WriteString(`"ref_Uint_fld16":`)
		buf = strconv.AppendUint(buf[:0], uint64(*s.RefUintFld16), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld16":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld32 != nil {
		result.WriteString(`"ref_Uint_fld32":`)
		buf = strconv.AppendUint(buf[:0], uint64(*s.RefUintFld32), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld32":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld64 != nil {
		result.WriteString(`"ref_Uint_fld64":`)
		buf = strconv.AppendUint(buf[:0], *s.RefUintFld64, 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld64":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefFloatFld32 != nil {
		result.WriteString(`"ref_fl23":`)
		buf = strconv.AppendFloat(buf[:0], float64(*s.RefFloatFld32), 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_fl23":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefFloatFld64 != nil {
		result.WriteString(`"ref_fl64":`)
		buf = strconv.AppendFloat(buf[:0], *s.RefFloatFld64, 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_fl64":null`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *NumStruct02) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s NumStruct02) MarshalAppend(dst []byte) ([]byte, error) {
	var result = bytes.NewBuffer(dst)
	var (
		err error
		buf = make([]byte, 0, 128)
	)
	result.WriteRune('{')
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld != 0 {
		result.WriteString(`"int_fld":`)
		buf = strconv.AppendInt(buf[:0], int64(s.IntFld), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld8 != 0 {
		result.WriteString(`"int_fld8":`)
		buf = strconv.AppendInt(buf[:0], int64(s.IntFld8), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld8":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld16 != 0 {
		result.WriteString(`"int_fld16":`)
		buf = strconv.AppendInt(buf[:0], int64(s.IntFld16), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld16":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld32 != 0 {
		result.WriteString(`"int_fld32":`)
		buf = strconv.AppendInt(buf[:0], int64(s.IntFld32), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld32":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.IntFld64 != 0 {
		result.WriteString(`"int_fld64":`)
		buf = strconv.AppendInt(buf[:0], s.IntFld64, 10)
		result.Write(buf)
	} else {
		result.WriteString(`"int_fld64":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld != 0 {
		result.WriteString(`"Uint_fld":`)
		buf = strconv.AppendUint(buf[:0], uint64(s.UintFld), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld8 != 0 {
		result.WriteString(`"Uint_fld8":`)
		buf = strconv.AppendUint(buf[:0], uint64(s.UintFld8), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld8":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld16 != 0 {
		result.WriteString(`"Uint_fld16":`)
		buf = strconv.AppendUint(buf[:0], uint64(s.UintFld16), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld16":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld32 != 0 {
		result.WriteString(`"Uint_fld32":`)
		buf = strconv.AppendUint(buf[:0], uint64(s.UintFld32), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld32":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.UintFld64 != 0 {
		result.WriteString(`"Uint_fld64":`)
		buf = strconv.AppendUint(buf[:0], s.UintFld64, 10)
		result.Write(buf)
	} else {
		result.WriteString(`"Uint_fld64":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.FloatFld32 != 0 {
		result.WriteString(`"fl23":`)
		buf = strconv.AppendFloat(buf[:0], float64(s.FloatFld32), 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"fl23":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.FloatFld64 != 0 {
		result.WriteString(`"fl64":`)
		buf = strconv.AppendFloat(buf[:0], s.FloatFld64, 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"fl64":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld != nil {
		result.WriteString(`"ref_int_fld":`)
		buf = strconv.AppendInt(buf[:0], int64(*s.RefIntFld), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld8 != nil {
		result.WriteString(`"ref_int_fld8":`)
		buf = strconv.AppendInt(buf[:0], int64(*s.RefIntFld8), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld8":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld16 != nil {
		result.WriteString(`"ref_int_fld16":`)
		buf = strconv.AppendInt(buf[:0], int64(*s.RefIntFld16), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld16":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld32 != nil {
		result.WriteString(`"ref_int_fld32":`)
		buf = strconv.AppendInt(buf[:0], int64(*s.RefIntFld32), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld32":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefIntFld64 != nil {
		result.WriteString(`"ref_int_fld64":`)
		buf = strconv.AppendInt(buf[:0], *s.RefIntFld64, 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_int_fld64":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld != nil {
		result.WriteString(`"ref_Uint_fld":`)
		buf = strconv.AppendUint(buf[:0], uint64(*s.RefUintFld), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld8 != nil {
		result.WriteString(`"ref_Uint_fld8":`)
		buf = strconv.AppendUint(buf[:0], uint64(*s.RefUintFld8), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld8":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld16 != nil {
		result.WriteString(`"ref_Uint_fld16":`)
		buf = strconv.AppendUint(buf[:0], uint64(*s.RefUintFld16), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld16":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld32 != nil {
		result.WriteString(`"ref_Uint_fld32":`)
		buf = strconv.AppendUint(buf[:0], uint64(*s.RefUintFld32), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld32":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefUintFld64 != nil {
		result.WriteString(`"ref_Uint_fld64":`)
		buf = strconv.AppendUint(buf[:0], *s.RefUintFld64, 10)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_Uint_fld64":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefFloatFld32 != nil {
		result.WriteString(`"ref_fl23":`)
		buf = strconv.AppendFloat(buf[:0], float64(*s.RefFloatFld32), 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_fl23":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.RefFloatFld64 != nil {
		result.WriteString(`"ref_fl64":`)
		buf = strconv.AppendFloat(buf[:0], *s.RefFloatFld64, 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"ref_fl64":null`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}
