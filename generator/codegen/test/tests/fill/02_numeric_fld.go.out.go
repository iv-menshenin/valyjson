// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package testo

import (
	"bytes"
	"fmt"
	"math"

	"github.com/valyala/fastjson"
)

// jsonParserNumStruct01used for pooling Parsers for NumStruct01 JSONs.
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
	if intfld := v.Get("int_fld"); intfld != nil {
		var valIntFld int
		valIntFld, err = intfld.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld' value: %w", objPath, err)
		}
		s.IntFld = valIntFld
	}
	if intfld8 := v.Get("int_fld8"); intfld8 != nil {
		var valIntFld8 int
		valIntFld8, err = intfld8.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld8' value: %w", objPath, err)
		}
		if valIntFld8 > math.MaxInt8 {
			return fmt.Errorf("error parsing '%sint_fld8' value %d exceeds maximum for data type int8", objPath, valIntFld8)
		}
		s.IntFld8 = int8(valIntFld8)
	}
	if intfld16 := v.Get("int_fld16"); intfld16 != nil {
		var valIntFld16 int
		valIntFld16, err = intfld16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld16' value: %w", objPath, err)
		}
		if valIntFld16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sint_fld16' value %d exceeds maximum for data type int16", objPath, valIntFld16)
		}
		s.IntFld16 = int16(valIntFld16)
	}
	if intfld32 := v.Get("int_fld32"); intfld32 != nil {
		var valIntFld32 int
		valIntFld32, err = intfld32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld32' value: %w", objPath, err)
		}
		if valIntFld32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sint_fld32' value %d exceeds maximum for data type int32", objPath, valIntFld32)
		}
		s.IntFld32 = int32(valIntFld32)
	}
	if intfld64 := v.Get("int_fld64"); intfld64 != nil {
		var valIntFld64 int64
		valIntFld64, err = intfld64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld64' value: %w", objPath, err)
		}
		s.IntFld64 = valIntFld64
	}
	if uintfld := v.Get("Uint_fld"); uintfld != nil {
		var valUintFld uint
		valUintFld, err = uintfld.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld' value: %w", objPath, err)
		}
		s.UintFld = valUintFld
	}
	if uintfld8 := v.Get("Uint_fld8"); uintfld8 != nil {
		var valUintFld8 uint
		valUintFld8, err = uintfld8.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld8' value: %w", objPath, err)
		}
		if valUintFld8 > math.MaxUint8 {
			return fmt.Errorf("error parsing '%sUint_fld8' value %d exceeds maximum for data type uint8", objPath, valUintFld8)
		}
		s.UintFld8 = uint8(valUintFld8)
	}
	if uintfld16 := v.Get("Uint_fld16"); uintfld16 != nil {
		var valUintFld16 uint
		valUintFld16, err = uintfld16.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld16' value: %w", objPath, err)
		}
		if valUintFld16 > math.MaxUint16 {
			return fmt.Errorf("error parsing '%sUint_fld16' value %d exceeds maximum for data type uint16", objPath, valUintFld16)
		}
		s.UintFld16 = uint16(valUintFld16)
	}
	if uintfld32 := v.Get("Uint_fld32"); uintfld32 != nil {
		var valUintFld32 uint
		valUintFld32, err = uintfld32.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld32' value: %w", objPath, err)
		}
		if valUintFld32 > math.MaxUint32 {
			return fmt.Errorf("error parsing '%sUint_fld32' value %d exceeds maximum for data type uint32", objPath, valUintFld32)
		}
		s.UintFld32 = uint32(valUintFld32)
	}
	if uintfld64 := v.Get("Uint_fld64"); uintfld64 != nil {
		var valUintFld64 uint64
		valUintFld64, err = uintfld64.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld64' value: %w", objPath, err)
		}
		s.UintFld64 = valUintFld64
	}
	if floatfld32 := v.Get("fl23"); floatfld32 != nil {
		var valFloatFld32 float64
		valFloatFld32, err = floatfld32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfl23' value: %w", objPath, err)
		}
		if valFloatFld32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%sfl23' value %f exceeds maximum for data type float32", objPath, valFloatFld32)
		}
		s.FloatFld32 = float32(valFloatFld32)
	}
	if floatfld64 := v.Get("fl64"); floatfld64 != nil {
		var valFloatFld64 float64
		valFloatFld64, err = floatfld64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfl64' value: %w", objPath, err)
		}
		s.FloatFld64 = valFloatFld64
	}
	if refintfld := v.Get("ref_int_fld"); valueIsNotNull(refintfld) {
		var valRefIntFld int
		valRefIntFld, err = refintfld.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld' value: %w", objPath, err)
		}
		s.RefIntFld = &valRefIntFld
	}
	if refintfld8 := v.Get("ref_int_fld8"); valueIsNotNull(refintfld8) {
		var valRefIntFld8 int
		valRefIntFld8, err = refintfld8.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld8' value: %w", objPath, err)
		}
		if valRefIntFld8 > math.MaxInt8 {
			return fmt.Errorf("error parsing '%sref_int_fld8' value %d exceeds maximum for data type int8", objPath, valRefIntFld8)
		}
		s.RefIntFld8 = new(int8)
		*s.RefIntFld8 = int8(valRefIntFld8)
	}
	if refintfld16 := v.Get("ref_int_fld16"); valueIsNotNull(refintfld16) {
		var valRefIntFld16 int
		valRefIntFld16, err = refintfld16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld16' value: %w", objPath, err)
		}
		if valRefIntFld16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sref_int_fld16' value %d exceeds maximum for data type int16", objPath, valRefIntFld16)
		}
		s.RefIntFld16 = new(int16)
		*s.RefIntFld16 = int16(valRefIntFld16)
	}
	if refintfld32 := v.Get("ref_int_fld32"); valueIsNotNull(refintfld32) {
		var valRefIntFld32 int
		valRefIntFld32, err = refintfld32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld32' value: %w", objPath, err)
		}
		if valRefIntFld32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sref_int_fld32' value %d exceeds maximum for data type int32", objPath, valRefIntFld32)
		}
		s.RefIntFld32 = new(int32)
		*s.RefIntFld32 = int32(valRefIntFld32)
	}
	if refintfld64 := v.Get("ref_int_fld64"); valueIsNotNull(refintfld64) {
		var valRefIntFld64 int64
		valRefIntFld64, err = refintfld64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld64' value: %w", objPath, err)
		}
		s.RefIntFld64 = &valRefIntFld64
	}
	if refuintfld := v.Get("ref_Uint_fld"); valueIsNotNull(refuintfld) {
		var valRefUintFld uint
		valRefUintFld, err = refuintfld.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld' value: %w", objPath, err)
		}
		s.RefUintFld = new(uint)
		*s.RefUintFld = uint(valRefUintFld)
	}
	if refuintfld8 := v.Get("ref_Uint_fld8"); valueIsNotNull(refuintfld8) {
		var valRefUintFld8 uint
		valRefUintFld8, err = refuintfld8.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld8' value: %w", objPath, err)
		}
		if valRefUintFld8 > math.MaxUint8 {
			return fmt.Errorf("error parsing '%sref_Uint_fld8' value %d exceeds maximum for data type uint8", objPath, valRefUintFld8)
		}
		s.RefUintFld8 = new(uint8)
		*s.RefUintFld8 = uint8(valRefUintFld8)
	}
	if refuintfld16 := v.Get("ref_Uint_fld16"); valueIsNotNull(refuintfld16) {
		var valRefUintFld16 uint
		valRefUintFld16, err = refuintfld16.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld16' value: %w", objPath, err)
		}
		if valRefUintFld16 > math.MaxUint16 {
			return fmt.Errorf("error parsing '%sref_Uint_fld16' value %d exceeds maximum for data type uint16", objPath, valRefUintFld16)
		}
		s.RefUintFld16 = new(uint16)
		*s.RefUintFld16 = uint16(valRefUintFld16)
	}
	if refuintfld32 := v.Get("ref_Uint_fld32"); valueIsNotNull(refuintfld32) {
		var valRefUintFld32 uint
		valRefUintFld32, err = refuintfld32.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld32' value: %w", objPath, err)
		}
		if valRefUintFld32 > math.MaxUint32 {
			return fmt.Errorf("error parsing '%sref_Uint_fld32' value %d exceeds maximum for data type uint32", objPath, valRefUintFld32)
		}
		s.RefUintFld32 = new(uint32)
		*s.RefUintFld32 = uint32(valRefUintFld32)
	}
	if refuintfld64 := v.Get("ref_Uint_fld64"); valueIsNotNull(refuintfld64) {
		var valRefUintFld64 uint64
		valRefUintFld64, err = refuintfld64.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld64' value: %w", objPath, err)
		}
		s.RefUintFld64 = new(uint64)
		*s.RefUintFld64 = uint64(valRefUintFld64)
	}
	if reffloatfld32 := v.Get("ref_fl23"); valueIsNotNull(reffloatfld32) {
		var valRefFloatFld32 float64
		valRefFloatFld32, err = reffloatfld32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_fl23' value: %w", objPath, err)
		}
		if valRefFloatFld32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%sref_fl23' value %f exceeds maximum for data type float32", objPath, valRefFloatFld32)
		}
		s.RefFloatFld32 = new(float32)
		*s.RefFloatFld32 = float32(valRefFloatFld32)
	}
	if reffloatfld64 := v.Get("ref_fl64"); valueIsNotNull(reffloatfld64) {
		var valRefFloatFld64 float64
		valRefFloatFld64, err = reffloatfld64.Float64()
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

// jsonParserNumStruct02used for pooling Parsers for NumStruct02 JSONs.
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
	if intfld := v.Get("int_fld"); intfld != nil {
		var valIntFld int
		valIntFld, err = intfld.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld' value: %w", objPath, err)
		}
		s.IntFld = valIntFld
	}
	if intfld8 := v.Get("int_fld8"); intfld8 != nil {
		var valIntFld8 int
		valIntFld8, err = intfld8.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld8' value: %w", objPath, err)
		}
		if valIntFld8 > math.MaxInt8 {
			return fmt.Errorf("error parsing '%sint_fld8' value %d exceeds maximum for data type int8", objPath, valIntFld8)
		}
		s.IntFld8 = int8(valIntFld8)
	}
	if intfld16 := v.Get("int_fld16"); intfld16 != nil {
		var valIntFld16 int
		valIntFld16, err = intfld16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld16' value: %w", objPath, err)
		}
		if valIntFld16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sint_fld16' value %d exceeds maximum for data type int16", objPath, valIntFld16)
		}
		s.IntFld16 = int16(valIntFld16)
	}
	if intfld32 := v.Get("int_fld32"); intfld32 != nil {
		var valIntFld32 int
		valIntFld32, err = intfld32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld32' value: %w", objPath, err)
		}
		if valIntFld32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sint_fld32' value %d exceeds maximum for data type int32", objPath, valIntFld32)
		}
		s.IntFld32 = int32(valIntFld32)
	}
	if intfld64 := v.Get("int_fld64"); intfld64 != nil {
		var valIntFld64 int64
		valIntFld64, err = intfld64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_fld64' value: %w", objPath, err)
		}
		s.IntFld64 = valIntFld64
	}
	if uintfld := v.Get("Uint_fld"); uintfld != nil {
		var valUintFld uint
		valUintFld, err = uintfld.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld' value: %w", objPath, err)
		}
		s.UintFld = valUintFld
	}
	if uintfld8 := v.Get("Uint_fld8"); uintfld8 != nil {
		var valUintFld8 uint
		valUintFld8, err = uintfld8.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld8' value: %w", objPath, err)
		}
		if valUintFld8 > math.MaxUint8 {
			return fmt.Errorf("error parsing '%sUint_fld8' value %d exceeds maximum for data type uint8", objPath, valUintFld8)
		}
		s.UintFld8 = uint8(valUintFld8)
	}
	if uintfld16 := v.Get("Uint_fld16"); uintfld16 != nil {
		var valUintFld16 uint
		valUintFld16, err = uintfld16.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld16' value: %w", objPath, err)
		}
		if valUintFld16 > math.MaxUint16 {
			return fmt.Errorf("error parsing '%sUint_fld16' value %d exceeds maximum for data type uint16", objPath, valUintFld16)
		}
		s.UintFld16 = uint16(valUintFld16)
	}
	if uintfld32 := v.Get("Uint_fld32"); uintfld32 != nil {
		var valUintFld32 uint
		valUintFld32, err = uintfld32.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld32' value: %w", objPath, err)
		}
		if valUintFld32 > math.MaxUint32 {
			return fmt.Errorf("error parsing '%sUint_fld32' value %d exceeds maximum for data type uint32", objPath, valUintFld32)
		}
		s.UintFld32 = uint32(valUintFld32)
	}
	if uintfld64 := v.Get("Uint_fld64"); uintfld64 != nil {
		var valUintFld64 uint64
		valUintFld64, err = uintfld64.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sUint_fld64' value: %w", objPath, err)
		}
		s.UintFld64 = valUintFld64
	}
	if floatfld32 := v.Get("fl23"); floatfld32 != nil {
		var valFloatFld32 float64
		valFloatFld32, err = floatfld32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfl23' value: %w", objPath, err)
		}
		if valFloatFld32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%sfl23' value %f exceeds maximum for data type float32", objPath, valFloatFld32)
		}
		s.FloatFld32 = float32(valFloatFld32)
	}
	if floatfld64 := v.Get("fl64"); floatfld64 != nil {
		var valFloatFld64 float64
		valFloatFld64, err = floatfld64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfl64' value: %w", objPath, err)
		}
		s.FloatFld64 = valFloatFld64
	}
	if refintfld := v.Get("ref_int_fld"); valueIsNotNull(refintfld) {
		var valRefIntFld int
		valRefIntFld, err = refintfld.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld' value: %w", objPath, err)
		}
		s.RefIntFld = &valRefIntFld
	}
	if refintfld8 := v.Get("ref_int_fld8"); valueIsNotNull(refintfld8) {
		var valRefIntFld8 int
		valRefIntFld8, err = refintfld8.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld8' value: %w", objPath, err)
		}
		if valRefIntFld8 > math.MaxInt8 {
			return fmt.Errorf("error parsing '%sref_int_fld8' value %d exceeds maximum for data type int8", objPath, valRefIntFld8)
		}
		s.RefIntFld8 = new(int8)
		*s.RefIntFld8 = int8(valRefIntFld8)
	}
	if refintfld16 := v.Get("ref_int_fld16"); valueIsNotNull(refintfld16) {
		var valRefIntFld16 int
		valRefIntFld16, err = refintfld16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld16' value: %w", objPath, err)
		}
		if valRefIntFld16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sref_int_fld16' value %d exceeds maximum for data type int16", objPath, valRefIntFld16)
		}
		s.RefIntFld16 = new(int16)
		*s.RefIntFld16 = int16(valRefIntFld16)
	}
	if refintfld32 := v.Get("ref_int_fld32"); valueIsNotNull(refintfld32) {
		var valRefIntFld32 int
		valRefIntFld32, err = refintfld32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld32' value: %w", objPath, err)
		}
		if valRefIntFld32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sref_int_fld32' value %d exceeds maximum for data type int32", objPath, valRefIntFld32)
		}
		s.RefIntFld32 = new(int32)
		*s.RefIntFld32 = int32(valRefIntFld32)
	}
	if refintfld64 := v.Get("ref_int_fld64"); valueIsNotNull(refintfld64) {
		var valRefIntFld64 int64
		valRefIntFld64, err = refintfld64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_int_fld64' value: %w", objPath, err)
		}
		s.RefIntFld64 = &valRefIntFld64
	}
	if refuintfld := v.Get("ref_Uint_fld"); valueIsNotNull(refuintfld) {
		var valRefUintFld uint
		valRefUintFld, err = refuintfld.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld' value: %w", objPath, err)
		}
		s.RefUintFld = new(uint)
		*s.RefUintFld = uint(valRefUintFld)
	}
	if refuintfld8 := v.Get("ref_Uint_fld8"); valueIsNotNull(refuintfld8) {
		var valRefUintFld8 uint
		valRefUintFld8, err = refuintfld8.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld8' value: %w", objPath, err)
		}
		if valRefUintFld8 > math.MaxUint8 {
			return fmt.Errorf("error parsing '%sref_Uint_fld8' value %d exceeds maximum for data type uint8", objPath, valRefUintFld8)
		}
		s.RefUintFld8 = new(uint8)
		*s.RefUintFld8 = uint8(valRefUintFld8)
	}
	if refuintfld16 := v.Get("ref_Uint_fld16"); valueIsNotNull(refuintfld16) {
		var valRefUintFld16 uint
		valRefUintFld16, err = refuintfld16.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld16' value: %w", objPath, err)
		}
		if valRefUintFld16 > math.MaxUint16 {
			return fmt.Errorf("error parsing '%sref_Uint_fld16' value %d exceeds maximum for data type uint16", objPath, valRefUintFld16)
		}
		s.RefUintFld16 = new(uint16)
		*s.RefUintFld16 = uint16(valRefUintFld16)
	}
	if refuintfld32 := v.Get("ref_Uint_fld32"); valueIsNotNull(refuintfld32) {
		var valRefUintFld32 uint
		valRefUintFld32, err = refuintfld32.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld32' value: %w", objPath, err)
		}
		if valRefUintFld32 > math.MaxUint32 {
			return fmt.Errorf("error parsing '%sref_Uint_fld32' value %d exceeds maximum for data type uint32", objPath, valRefUintFld32)
		}
		s.RefUintFld32 = new(uint32)
		*s.RefUintFld32 = uint32(valRefUintFld32)
	}
	if refuintfld64 := v.Get("ref_Uint_fld64"); valueIsNotNull(refuintfld64) {
		var valRefUintFld64 uint64
		valRefUintFld64, err = refuintfld64.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_Uint_fld64' value: %w", objPath, err)
		}
		s.RefUintFld64 = new(uint64)
		*s.RefUintFld64 = uint64(valRefUintFld64)
	}
	if reffloatfld32 := v.Get("ref_fl23"); valueIsNotNull(reffloatfld32) {
		var valRefFloatFld32 float64
		valRefFloatFld32, err = reffloatfld32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%sref_fl23' value: %w", objPath, err)
		}
		if valRefFloatFld32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%sref_fl23' value %f exceeds maximum for data type float32", objPath, valRefFloatFld32)
		}
		s.RefFloatFld32 = new(float32)
		*s.RefFloatFld32 = float32(valRefFloatFld32)
	}
	if reffloatfld64 := v.Get("ref_fl64"); valueIsNotNull(reffloatfld64) {
		var valRefFloatFld64 float64
		valRefFloatFld64, err = reffloatfld64.Float64()
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
