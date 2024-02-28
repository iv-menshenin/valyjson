// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_userdefined

import (
	"bytes"
	"fmt"
	"math"
	"time"

	"github.com/mailru/easyjson/jwriter"
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
		s.String = DefinedString(valString)
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

// jsonParserTestUserDefinedRef used for pooling Parsers for TestUserDefinedRef JSONs.
var jsonParserTestUserDefinedRef fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestUserDefinedRef) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestUserDefinedRef.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestUserDefinedRef.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestUserDefinedRef) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _int32Ref := v.Get("d_int32"); _int32Ref != nil {
		var valInt32Ref int
		valInt32Ref, err = _int32Ref.Int()
		if err != nil {
			return newParsingError("d_int32", err)
		}
		var _ref int32
		_ref = int32(valInt32Ref)
		s.Int32Ref = DefinedRefInt32(&_ref)
	} else {
		var __Int32Ref int32 = 32
		s.Int32Ref = DefinedRefInt32(&__Int32Ref)
	}
	if _int64Ref := v.Get("d_int64"); _int64Ref != nil {
		var valInt64Ref int64
		valInt64Ref, err = _int64Ref.Int64()
		if err != nil {
			return newParsingError("d_int64", err)
		}
		var _ref int64
		_ref = valInt64Ref
		s.Int64Ref = DefinedRefInt64(&_ref)
	}
	if _float32Ref := v.Get("d_float32"); _float32Ref != nil {
		var valFloat32Ref float64
		valFloat32Ref, err = _float32Ref.Float64()
		if err != nil {
			return newParsingError("d_float32", err)
		}
		var _ref float32
		_ref = float32(valFloat32Ref)
		s.Float32Ref = DefinedRefFloat32(&_ref)
	} else {
		var __Float32Ref float32 = 123.01
		s.Float32Ref = DefinedRefFloat32(&__Float32Ref)
	}
	if _float64Ref := v.Get("d_float64"); _float64Ref != nil {
		var valFloat64Ref float64
		valFloat64Ref, err = _float64Ref.Float64()
		if err != nil {
			return newParsingError("d_float64", err)
		}
		var _ref float64
		_ref = valFloat64Ref
		s.Float64Ref = DefinedRefFloat64(&_ref)
	}
	if _stringRef := v.Get("d_string"); _stringRef != nil {
		var valStringRef []byte
		if valStringRef, err = _stringRef.StringBytes(); err != nil {
			return newParsingError("d_string", err)
		}
		if err != nil {
			return newParsingError("d_string", err)
		}
		var _ref string
		_ref = string(valStringRef)
		s.StringRef = DefinedRefString(&_ref)
	} else {
		var __StringRef string = "default_string"
		s.StringRef = DefinedRefString(&__StringRef)
	}
	if _boolRef := v.Get("d_bool"); _boolRef != nil {
		var valBoolRef bool
		valBoolRef, err = _boolRef.Bool()
		if err != nil {
			return newParsingError("d_bool", err)
		}
		var _ref bool
		_ref = valBoolRef
		s.BoolRef = DefinedRefBool(&_ref)
	}
	if _refInt32Ref := v.Get("x_int32"); valueIsNotNull(_refInt32Ref) {
		var valRefInt32Ref int
		valRefInt32Ref, err = _refInt32Ref.Int()
		if err != nil {
			return newParsingError("x_int32", err)
		}
		var _ref int32
		_ref = int32(valRefInt32Ref)
		s.RefInt32Ref = new(DefinedRefInt32)
		*s.RefInt32Ref = DefinedRefInt32(&_ref)
	} else {
		var __RefInt32Ref int32 = 32
		s.RefInt32Ref = new(DefinedRefInt32)
		*s.RefInt32Ref = DefinedRefInt32(&__RefInt32Ref)
	}
	if _refInt64Ref := v.Get("x_int64"); valueIsNotNull(_refInt64Ref) {
		var valRefInt64Ref int64
		valRefInt64Ref, err = _refInt64Ref.Int64()
		if err != nil {
			return newParsingError("x_int64", err)
		}
		var _ref int64
		_ref = valRefInt64Ref
		s.RefInt64Ref = new(DefinedRefInt64)
		*s.RefInt64Ref = DefinedRefInt64(&_ref)
	}
	if _refFloat32Ref := v.Get("x_float32"); valueIsNotNull(_refFloat32Ref) {
		var valRefFloat32Ref float64
		valRefFloat32Ref, err = _refFloat32Ref.Float64()
		if err != nil {
			return newParsingError("x_float32", err)
		}
		var _ref float32
		_ref = float32(valRefFloat32Ref)
		s.RefFloat32Ref = new(DefinedRefFloat32)
		*s.RefFloat32Ref = DefinedRefFloat32(&_ref)
	} else {
		var __RefFloat32Ref float32 = 123.01
		s.RefFloat32Ref = new(DefinedRefFloat32)
		*s.RefFloat32Ref = DefinedRefFloat32(&__RefFloat32Ref)
	}
	if _refFloat64Ref := v.Get("x_float64"); valueIsNotNull(_refFloat64Ref) {
		var valRefFloat64Ref float64
		valRefFloat64Ref, err = _refFloat64Ref.Float64()
		if err != nil {
			return newParsingError("x_float64", err)
		}
		var _ref float64
		_ref = valRefFloat64Ref
		s.RefFloat64Ref = new(DefinedRefFloat64)
		*s.RefFloat64Ref = DefinedRefFloat64(&_ref)
	}
	if _refStringRef := v.Get("x_string"); valueIsNotNull(_refStringRef) {
		var valRefStringRef []byte
		if valRefStringRef, err = _refStringRef.StringBytes(); err != nil {
			return newParsingError("x_string", err)
		}
		if err != nil {
			return newParsingError("x_string", err)
		}
		var _ref string
		_ref = string(valRefStringRef)
		s.RefStringRef = new(DefinedRefString)
		*s.RefStringRef = DefinedRefString(&_ref)
	} else {
		var __RefStringRef string = "default_string"
		s.RefStringRef = new(DefinedRefString)
		*s.RefStringRef = DefinedRefString(&__RefStringRef)
	}
	if _refBoolRef := v.Get("x_bool"); valueIsNotNull(_refBoolRef) {
		var valRefBoolRef bool
		valRefBoolRef, err = _refBoolRef.Bool()
		if err != nil {
			return newParsingError("x_bool", err)
		}
		var _ref bool
		_ref = valRefBoolRef
		s.RefBoolRef = new(DefinedRefBool)
		*s.RefBoolRef = DefinedRefBool(&_ref)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestUserDefinedRef) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [12]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'d', '_', 'i', 'n', 't', '3', '2'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', '_', 'i', 'n', 't', '6', '4'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', '_', 'f', 'l', 'o', 'a', 't', '3', '2'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', '_', 'f', 'l', 'o', 'a', 't', '6', '4'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', '_', 's', 't', 'r', 'i', 'n', 'g'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', '_', 'b', 'o', 'o', 'l'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'x', '_', 'i', 'n', 't', '3', '2'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'x', '_', 'i', 'n', 't', '6', '4'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'x', '_', 'f', 'l', 'o', 'a', 't', '3', '2'}) {
			checkFields[8]++
			if checkFields[8] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'x', '_', 'f', 'l', 'o', 'a', 't', '6', '4'}) {
			checkFields[9]++
			if checkFields[9] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'x', '_', 's', 't', 'r', 'i', 'n', 'g'}) {
			checkFields[10]++
			if checkFields[10] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'x', '_', 'b', 'o', 'o', 'l'}) {
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
		s.Status = userdefined.DefinedFieldAsUserDefinedStatus(valStatus)
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestUserDefined) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestUserDefined) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.Int32 != 0 {
		result.RawString(`"f_int32":`)
		result.Int64(int64(s.Int32))
		wantComma = true
	} else {
		result.RawString(`"f_int32":0`)
		wantComma = true
	}
	if s.Int64 != 0 {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"f_int64":`)
		result.Int64(int64(s.Int64))
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Float32 != 0 {
		result.RawString(`"f_float32":`)
		result.Float64(float64(s.Float32))
		wantComma = true
	} else {
		result.RawString(`"f_float32":0`)
		wantComma = true
	}
	if s.Float64 != 0 {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"f_float64":`)
		result.Float64(float64(s.Float64))
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.String != "" {
		result.RawString(`"f_string":`)
		result.String(string(s.String))
		wantComma = true
	} else {
		result.RawString(`"f_string":""`)
		wantComma = true
	}
	if s.Bool {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"f_bool":true`)
		wantComma = true
	}
	if s.RefInt32 != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"r_int32":`)
		result.Int64(int64(*s.RefInt32))
		wantComma = true
	}
	if s.RefInt64 != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"r_int64":`)
		result.Int64(int64(*s.RefInt64))
		wantComma = true
	}
	if s.RefFloat32 != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"r_float32":`)
		result.Float64(float64(*s.RefFloat32))
		wantComma = true
	}
	if s.RefFloat64 != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"r_float64":`)
		result.Float64(float64(*s.RefFloat64))
		wantComma = true
	}
	if s.RefString != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"r_string":`)
		result.String(string(*s.RefString))
		wantComma = true
	}
	if s.RefBool != nil {
		if wantComma {
			result.RawByte(',')
		}
		if *s.RefBool {
			result.RawString(`"r_bool":true`)
		} else {
			result.RawString(`"r_bool":false`)
		}
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestUserDefined) Reset() {
	s.Int32 = DefinedInt32(0)
	s.Int64 = DefinedInt64(0)
	s.Float32 = DefinedFloat32(0)
	s.Float64 = DefinedFloat64(0)
	s.String = DefinedString("")
	s.Bool = DefinedBool(false)
	s.RefInt32 = nil
	s.RefInt64 = nil
	s.RefFloat32 = nil
	s.RefFloat64 = nil
	s.RefString = nil
	s.RefBool = nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestUserDefinedRef) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestUserDefinedRef) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.Int32Ref != nil {
		result.RawString(`"d_int32":`)
		result.Int64(int64(*s.Int32Ref))
		wantComma = true
	} else {
		result.RawString(`"d_int32":null`)
	}
	if s.Int64Ref != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"d_int64":`)
		result.Int64(int64(*s.Int64Ref))
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Float32Ref != nil {
		result.RawString(`"d_float32":`)
		result.Float64(float64(*s.Float32Ref))
		wantComma = true
	} else {
		result.RawString(`"d_float32":null`)
	}
	if s.Float64Ref != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"d_float64":`)
		result.Float64(float64(*s.Float64Ref))
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.StringRef != nil {
		result.RawString(`"d_string":`)
		result.String(string(*s.StringRef))
		wantComma = true
	} else {
		result.RawString(`"d_string":null`)
	}
	if s.BoolRef != nil {
		if wantComma {
			result.RawByte(',')
		}
		if *s.BoolRef {
			result.RawString(`"d_bool":true`)
		} else {
			result.RawString(`"d_bool":false`)
		}
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.RefInt32Ref != nil && *s.RefInt32Ref != nil {
		result.RawString(`"x_int32":`)
		result.Int64(int64(**s.RefInt32Ref))
		wantComma = true
	} else {
		result.RawString(`"x_int32":null`)
	}
	if s.RefInt64Ref != nil && *s.RefInt64Ref != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"x_int64":`)
		result.Int64(int64(**s.RefInt64Ref))
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.RefFloat32Ref != nil && *s.RefFloat32Ref != nil {
		result.RawString(`"x_float32":`)
		result.Float64(float64(**s.RefFloat32Ref))
		wantComma = true
	} else {
		result.RawString(`"x_float32":null`)
	}
	if s.RefFloat64Ref != nil && *s.RefFloat64Ref != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"x_float64":`)
		result.Float64(float64(**s.RefFloat64Ref))
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.RefStringRef != nil && *s.RefStringRef != nil {
		result.RawString(`"x_string":`)
		result.String(string(**s.RefStringRef))
		wantComma = true
	} else {
		result.RawString(`"x_string":null`)
	}
	if s.RefBoolRef != nil && *s.RefBoolRef != nil {
		if wantComma {
			result.RawByte(',')
		}
		if **s.RefBoolRef {
			result.RawString(`"x_bool":true`)
		} else {
			result.RawString(`"x_bool":false`)
		}
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestUserDefinedRef) IsZero() bool {
	if s.Int32Ref != nil {
		return false
	}
	if s.Int64Ref != nil {
		return false
	}
	if s.Float32Ref != nil {
		return false
	}
	if s.Float64Ref != nil {
		return false
	}
	if s.StringRef != nil {
		return false
	}
	if s.BoolRef != nil {
		return false
	}
	if s.RefInt32Ref != nil {
		return false
	}
	if s.RefInt64Ref != nil {
		return false
	}
	if s.RefFloat32Ref != nil {
		return false
	}
	if s.RefFloat64Ref != nil {
		return false
	}
	if s.RefStringRef != nil {
		return false
	}
	if s.RefBoolRef != nil {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestUserDefinedRef) Reset() {
	s.Int32Ref = DefinedRefInt32(nil)
	s.Int64Ref = DefinedRefInt64(nil)
	s.Float32Ref = DefinedRefFloat32(nil)
	s.Float64Ref = DefinedRefFloat64(nil)
	s.StringRef = DefinedRefString(nil)
	s.BoolRef = DefinedRefBool(nil)
	s.RefInt32Ref = nil
	s.RefInt64Ref = nil
	s.RefFloat32Ref = nil
	s.RefFloat64Ref = nil
	s.RefStringRef = nil
	s.RefBoolRef = nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *DefinedFieldAsUserDefined) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *DefinedFieldAsUserDefined) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.Status != "" {
		result.RawString(`"status":`)
		result.String(string(s.Status))
		wantComma = true
	} else {
		result.RawString(`"status":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if !s.Time.IsZero() {
		result.RawString(`"time":`)
		writeTime(result, s.Time, time.RFC3339Nano)
		wantComma = true
	} else {
		result.RawString(`"time":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *DefinedFieldAsUserDefined) Reset() {
	s.Status = userdefined.DefinedFieldAsUserDefinedStatus("")
	s.Time = time.Time{}
}
