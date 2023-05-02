// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_userdefined

import (
	"bytes"
	"fmt"
	"math"
	"unsafe"

	"github.com/valyala/fastjson"
)

// jsonParserTestUserDefined used for pooling Parsers for TestUserDefined JSONs.
var jsonParserTestUserDefined fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestUserDefined) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestUserDefined.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestUserDefined.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestUserDefined) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _int32 := v.Get("f_int32"); _int32 != nil {
		var valInt32 int
		valInt32, err = _int32.Int()
		if err != nil {
			return newParsingError("f_int32", err)
		}
		if valInt32 > math.MaxInt32 {
			return newParsingError("f_int32", fmt.Errorf("%d exceeds maximum for data type int32", valInt32))
		}
		s.Int32 = DefinedInt32(valInt32)
	} else {
		s.Int32 = 32
	}
	if _int64 := v.Get("f_int64"); _int64 != nil {
		var valInt64 int64
		valInt64, err = _int64.Int64()
		if err != nil {
			return newParsingError("f_int64", err)
		}
		s.Int64 = DefinedInt64(valInt64)
	}
	if _float32 := v.Get("f_float32"); _float32 != nil {
		var valFloat32 float64
		valFloat32, err = _float32.Float64()
		if err != nil {
			return newParsingError("f_float32", err)
		}
		if valFloat32 > math.MaxFloat32 {
			return newParsingError("f_float32", fmt.Errorf("%f exceeds maximum for data type float32", valFloat32))
		}
		s.Float32 = DefinedFloat32(valFloat32)
	} else {
		s.Float32 = 123.01
	}
	if _float64 := v.Get("f_float64"); _float64 != nil {
		var valFloat64 float64
		valFloat64, err = _float64.Float64()
		if err != nil {
			return newParsingError("f_float64", err)
		}
		s.Float64 = DefinedFloat64(valFloat64)
	}
	if _string := v.Get("f_string"); _string != nil {
		var valString []byte
		if valString, err = _string.StringBytes(); err != nil {
			return newParsingError("f_string", err)
		}
		s.String = *(*DefinedString)(unsafe.Pointer(&valString))
	} else {
		s.String = "default_string"
	}
	if _bool := v.Get("f_bool"); _bool != nil {
		var valBool bool
		valBool, err = _bool.Bool()
		if err != nil {
			return newParsingError("f_bool", err)
		}
		s.Bool = DefinedBool(valBool)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestUserDefined) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [6]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'i', 'n', 't', '3', '2'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'i', 'n', 't', '6', '4'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'f', 'l', 'o', 'a', 't', '3', '2'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'f', 'l', 'o', 'a', 't', '6', '4'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 's', 't', 'r', 'i', 'n', 'g'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'b', 'o', 'o', 'l'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

var bufDataTestUserDefined = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestUserDefined) MarshalJSON() ([]byte, error) {
	var result = bufDataTestUserDefined.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestUserDefined) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if wantComma {
		result.WriteString(",")
	}
	if s.Int32 != 0 {
		result.WriteString(`"f_int32":`)
		writeInt64(result, int64(s.Int32))
		wantComma = true
	} else {
		result.WriteString(`"f_int32":0`)
		wantComma = true
	}
	if s.Int64 != 0 {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"f_int64":`)
		writeInt64(result, int64(s.Int64))
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Float32 != 0 {
		result.WriteString(`"f_float32":`)
		writeFloat64(result, float64(s.Float32))
		wantComma = true
	} else {
		result.WriteString(`"f_float32":0`)
		wantComma = true
	}
	if s.Float64 != 0 {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"f_float64":`)
		writeFloat64(result, float64(s.Float64))
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.String != "" {
		result.WriteString(`"f_string":`)
		writeString(result, string(s.String))
		wantComma = true
	} else {
		result.WriteString(`"f_string":""`)
		wantComma = true
	}
	if s.Bool {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"f_bool":true`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestUserDefined) IsZero() bool {
	if s.Int32 != 0 {
		return false
	}
	if s.Int64 != 0 {
		return false
	}
	if s.Float32 != 0 {
		return false
	}
	if s.Float64 != 0 {
		return false
	}
	if s.String != "" {
		return false
	}
	if s.Bool != false {
		return false
	}
	return true
}
