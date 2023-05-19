// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_userdefined

import (
	"bytes"
	"fmt"
	"math"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"

	"fill/test_userdefined/userdefined"
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
	if _refInt32 := v.Get("r_int32"); valueIsNotNull(_refInt32) {
		var valRefInt32 int
		valRefInt32, err = _refInt32.Int()
		if err != nil {
			return newParsingError("r_int32", err)
		}
		if valRefInt32 > math.MaxInt32 {
			return newParsingError("r_int32", fmt.Errorf("%d exceeds maximum for data type int32", valRefInt32))
		}
		s.RefInt32 = new(DefinedInt32)
		*s.RefInt32 = DefinedInt32(valRefInt32)
	}
	if _refInt64 := v.Get("r_int64"); valueIsNotNull(_refInt64) {
		var valRefInt64 int64
		valRefInt64, err = _refInt64.Int64()
		if err != nil {
			return newParsingError("r_int64", err)
		}
		s.RefInt64 = new(DefinedInt64)
		*s.RefInt64 = DefinedInt64(valRefInt64)
	}
	if _refFloat32 := v.Get("r_float32"); valueIsNotNull(_refFloat32) {
		var valRefFloat32 float64
		valRefFloat32, err = _refFloat32.Float64()
		if err != nil {
			return newParsingError("r_float32", err)
		}
		if valRefFloat32 > math.MaxFloat32 {
			return newParsingError("r_float32", fmt.Errorf("%f exceeds maximum for data type float32", valRefFloat32))
		}
		s.RefFloat32 = new(DefinedFloat32)
		*s.RefFloat32 = DefinedFloat32(valRefFloat32)
	}
	if _refFloat64 := v.Get("r_float64"); valueIsNotNull(_refFloat64) {
		var valRefFloat64 float64
		valRefFloat64, err = _refFloat64.Float64()
		if err != nil {
			return newParsingError("r_float64", err)
		}
		s.RefFloat64 = new(DefinedFloat64)
		*s.RefFloat64 = DefinedFloat64(valRefFloat64)
	}
	if _refString := v.Get("r_string"); valueIsNotNull(_refString) {
		var valRefString []byte
		if valRefString, err = _refString.StringBytes(); err != nil {
			return newParsingError("r_string", err)
		}
		s.RefString = new(DefinedString)
		*s.RefString = DefinedString(valRefString)
	}
	if _refBool := v.Get("r_bool"); valueIsNotNull(_refBool) {
		var valRefBool bool
		valRefBool, err = _refBool.Bool()
		if err != nil {
			return newParsingError("r_bool", err)
		}
		s.RefBool = new(DefinedBool)
		*s.RefBool = DefinedBool(valRefBool)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestUserDefined) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [12]int
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
		if bytes.Equal(key, []byte{'r', '_', 'i', 'n', 't', '3', '2'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', '_', 'i', 'n', 't', '6', '4'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', '_', 'f', 'l', 'o', 'a', 't', '3', '2'}) {
			checkFields[8]++
			if checkFields[8] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', '_', 'f', 'l', 'o', 'a', 't', '6', '4'}) {
			checkFields[9]++
			if checkFields[9] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', '_', 's', 't', 'r', 'i', 'n', 'g'}) {
			checkFields[10]++
			if checkFields[10] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', '_', 'b', 'o', 'o', 'l'}) {
			checkFields[11]++
			if checkFields[11] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserDefinedFieldAsUserDefined used for pooling Parsers for DefinedFieldAsUserDefined JSONs.
var jsonParserDefinedFieldAsUserDefined fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *DefinedFieldAsUserDefined) UnmarshalJSON(data []byte) error {
	parser := jsonParserDefinedFieldAsUserDefined.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserDefinedFieldAsUserDefined.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *DefinedFieldAsUserDefined) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _status := v.Get("status"); _status != nil {
		var valStatus []byte
		if valStatus, err = _status.StringBytes(); err != nil {
			return newParsingError("status", err)
		}
		s.Status = *(*userdefined.DefinedFieldAsUserDefinedStatus)(unsafe.Pointer(&valStatus))
	}
	if _time := v.Get("time"); _time != nil {
		b, err := _time.StringBytes()
		if err != nil {
			return newParsingError("time", err)
		}
		valTime, err := parseDateTime(string(b))
		if err != nil {
			return newParsingError("time", err)
		}
		s.Time = valTime
	}
	return nil
}

// validate checks for correct data structure
func (s *DefinedFieldAsUserDefined) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'s', 't', 'a', 't', 'u', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'i', 'm', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
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
	if s.RefInt32 != nil {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"r_int32":`)
		writeInt64(result, int64(*s.RefInt32))
		wantComma = true
	}
	if s.RefInt64 != nil {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"r_int64":`)
		writeInt64(result, int64(*s.RefInt64))
		wantComma = true
	}
	if s.RefFloat32 != nil {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"r_float32":`)
		writeFloat64(result, float64(*s.RefFloat32))
		wantComma = true
	}
	if s.RefFloat64 != nil {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"r_float64":`)
		writeFloat64(result, float64(*s.RefFloat64))
		wantComma = true
	}
	if s.RefString != nil {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"r_string":`)
		writeString(result, string(*s.RefString))
		wantComma = true
	}
	if s.RefBool != nil {
		if wantComma {
			result.WriteString(",")
		}
		if *s.RefBool {
			result.WriteString(`"r_bool":true`)
		} else {
			result.WriteString(`"r_bool":false`)
		}
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
	if s.RefInt32 != nil {
		return false
	}
	if s.RefInt64 != nil {
		return false
	}
	if s.RefFloat32 != nil {
		return false
	}
	if s.RefFloat64 != nil {
		return false
	}
	if s.RefString != nil {
		return false
	}
	if s.RefBool != nil {
		return false
	}
	return true
}

var bufDataDefinedFieldAsUserDefined = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *DefinedFieldAsUserDefined) MarshalJSON() ([]byte, error) {
	var result = bufDataDefinedFieldAsUserDefined.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *DefinedFieldAsUserDefined) MarshalTo(result Writer) error {
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
	if s.Status != "" {
		result.WriteString(`"status":`)
		writeString(result, string(s.Status))
		wantComma = true
	} else {
		result.WriteString(`"status":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if !s.Time.IsZero() {
		result.WriteString(`"time":`)
		writeTime(result, s.Time, time.RFC3339Nano)
		wantComma = true
	} else {
		result.WriteString(`"time":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s DefinedFieldAsUserDefined) IsZero() bool {
	if s.Status != "" {
		return false
	}
	if !s.Time.IsZero() {
		return false
	}
	return true
}
