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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestUserDefined) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _int32 := v.Get("f_int32"); _int32 != nil {
		var valInt32 int
		valInt32, err = _int32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.f_int32' value: %w", objPath, err)
		}
		if valInt32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%s.f_int32' value %d exceeds maximum for data type int32", objPath, valInt32)
		}
		s.Int32 = DefinedInt32(valInt32)
	} else {
		s.Int32 = 32
	}
	if _int64 := v.Get("f_int64"); _int64 != nil {
		var valInt64 int64
		valInt64, err = _int64.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%s.f_int64' value: %w", objPath, err)
		}
		s.Int64 = DefinedInt64(valInt64)
	}
	if _float32 := v.Get("f_float32"); _float32 != nil {
		var valFloat32 float64
		valFloat32, err = _float32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%s.f_float32' value: %w", objPath, err)
		}
		if valFloat32 > math.MaxFloat32 {
			return fmt.Errorf("error parsing '%s.f_float32' value %f exceeds maximum for data type float32", objPath, valFloat32)
		}
		s.Float32 = DefinedFloat32(valFloat32)
	} else {
		s.Float32 = 123.01
	}
	if _float64 := v.Get("f_float64"); _float64 != nil {
		var valFloat64 float64
		valFloat64, err = _float64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%s.f_float64' value: %w", objPath, err)
		}
		s.Float64 = DefinedFloat64(valFloat64)
	}
	if _string := v.Get("f_string"); _string != nil {
		var valString []byte
		if valString, err = _string.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.f_string' value: %w", objPath, err)
		}
		s.String = *(*DefinedString)(unsafe.Pointer(&valString))
	} else {
		s.String = "default_string"
	}
	if _bool := v.Get("f_bool"); _bool != nil {
		var valBool bool
		valBool, err = _bool.Bool()
		if err != nil {
			return fmt.Errorf("error parsing '%s.f_bool' value: %w", objPath, err)
		}
		s.Bool = DefinedBool(valBool)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestUserDefined) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'i', 'n', 't', '6', '4'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'f', 'l', 'o', 'a', 't', '3', '2'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'f', 'l', 'o', 'a', 't', '6', '4'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 's', 't', 'r', 'i', 'n', 'g'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', '_', 'b', 'o', 'o', 'l'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestUserDefined) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
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
	result.Write([]byte{'{'})
	if wantComma {
		result.Write([]byte{','})
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
			result.Write([]byte{','})
		}
		result.WriteString(`"f_int64":`)
		writeInt64(result, int64(s.Int64))
		wantComma = true
	}
	if wantComma {
		result.Write([]byte{','})
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
			result.Write([]byte{','})
		}
		result.WriteString(`"f_float64":`)
		writeFloat64(result, float64(s.Float64))
		wantComma = true
	}
	if wantComma {
		result.Write([]byte{','})
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
			result.Write([]byte{','})
		}
		result.WriteString(`"f_bool":true`)
		wantComma = true
	}
	result.Write([]byte{'}'})
	return err
}
