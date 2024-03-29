// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_extr

import (
	"bytes"
	"fmt"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/fastjson"

	"fill/test_string"
)

// jsonParserExternal used for pooling Parsers for External JSONs.
var jsonParserExternal fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *External) UnmarshalJSON(data []byte) error {
	parser := jsonParserExternal.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserExternal.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *External) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _test01 := v.Get("test1"); _test01 != nil {
		var valTest01 = &s.Test01
		err = valTest01.FillFromJSON(_test01)
		if err != nil {
			return newParsingError("test1", err)
		}
	}
	if _test02 := v.Get("test2"); _test02 != nil {
		var valTest02 = &s.Test02
		err = valTest02.FillFromJSON(_test02)
		if err != nil {
			return newParsingError("test2", err)
		}
	}
	return nil
}

// validate checks for correct data structure
func (s *External) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 's', 't', '1'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 's', 't', '2'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserExternalStructSlice used for pooling Parsers for ExternalStructSlice JSONs.
var jsonParserExternalStructSlice fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *ExternalStructSlice) UnmarshalJSON(data []byte) error {
	parser := jsonParserExternalStructSlice.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserExternalStructSlice.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *ExternalStructSlice) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	a, err := v.Array()
	if err != nil {
		return err
	}
	*s = make([]test_string.TestStr01, len(a))
	for i, v := range a {
		var value test_string.TestStr01
		err = value.FillFromJSON(v)
		if err != nil {
			return newParsingError(fmt.Sprintf("%d", i), err)
		}
		(*s)[i] = test_string.TestStr01(value)
	}
	return nil
}

// jsonParserExternalStringSlice used for pooling Parsers for ExternalStringSlice JSONs.
var jsonParserExternalStringSlice fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *ExternalStringSlice) UnmarshalJSON(data []byte) error {
	parser := jsonParserExternalStringSlice.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserExternalStringSlice.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *ExternalStringSlice) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	a, err := v.Array()
	if err != nil {
		return err
	}
	*s = make([]test_string.FieldValueString, len(a))
	for i, v := range a {
		var value []byte
		value, err = v.StringBytes()
		if err != nil {
			return newParsingError(fmt.Sprintf("%d", i), err)
		}
		(*s)[i] = test_string.FieldValueString(value)
	}
	return nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *External) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *External) MarshalTo(result *jwriter.Writer) error {
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
	result.RawString(`"test1":`)
	if err = s.Test01.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "test1" attribute: %w`, err)
	}
	wantComma = true
	if !s.Test02.IsZero() {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"test2":`)
		if err = s.Test02.MarshalTo(result); err != nil {
			return fmt.Errorf(`can't marshal "test2" attribute: %w`, err)
		}
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s External) IsZero() bool {
	if !s.Test01.IsZero() {
		return false
	}
	if !s.Test02.IsZero() {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *External) Reset() {
	s.Test01.Reset()
	s.Test02.Reset()
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *ExternalStructSlice) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *ExternalStructSlice) MarshalTo(result *jwriter.Writer) error {
	if s == nil || *s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('[')
	for _k, _v := range *s {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		_k = _k
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "ExternalStructSlice" value at position %d: %w`, _k, err)
		}
	}
	result.RawByte(']')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s ExternalStructSlice) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *ExternalStructSlice) Reset() {
	for i := range *s {
		(*s)[i].Reset()
	}
	*s = (*s)[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *ExternalStringSlice) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *ExternalStringSlice) MarshalTo(result *jwriter.Writer) error {
	if s == nil || *s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('[')
	for _k, _v := range *s {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		_k = _k
		result.String(string(_v))
	}
	result.RawByte(']')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s ExternalStringSlice) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *ExternalStringSlice) Reset() {
	*s = (*s)[:0]
}
