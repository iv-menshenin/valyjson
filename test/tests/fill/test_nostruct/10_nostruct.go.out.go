// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_nostruct

import (
	"fmt"
	"time"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/fastjson"

	"fill/test_extr"
)

// jsonParserTestMap10 used for pooling Parsers for TestMap10 JSONs.
var jsonParserTestMap10 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap10) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap10.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap10.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *TestMap10) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	o, err := v.Object()
	if err != nil {
		return err
	}
	*s = make(map[string]int64, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value int64
		value, err = v.Int64()
		if err != nil {
			err = newParsingError(string(key), err)
			return
		}
		(*s)[string(key)] = int64(value)
	})
	return err
}

// jsonParserTestMap11 used for pooling Parsers for TestMap11 JSONs.
var jsonParserTestMap11 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap11) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap11.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap11.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *TestMap11) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	o, err := v.Object()
	if err != nil {
		return err
	}
	*s = make(map[string]test_extr.External, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value test_extr.External
		err = value.FillFromJSON(v)
		if err != nil {
			err = newParsingError(string(key), err)
			return
		}
		(*s)[string(key)] = test_extr.External(value)
	})
	return err
}

// jsonParserTestMap11Ref used for pooling Parsers for TestMap11Ref JSONs.
var jsonParserTestMap11Ref fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap11Ref) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap11Ref.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap11Ref.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *TestMap11Ref) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	o, err := v.Object()
	if err != nil {
		return err
	}
	*s = make(map[string]*test_extr.External, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		if v.Type() == fastjson.TypeNull {
			(*s)[string(key)] = nil
			return
		}
		var value test_extr.External
		err = value.FillFromJSON(v)
		if err != nil {
			err = newParsingError(string(key), err)
			return
		}
		(*s)[string(key)] = &value
	})
	return err
}

// jsonParserTestSlice12 used for pooling Parsers for TestSlice12 JSONs.
var jsonParserTestSlice12 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice12) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice12.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice12.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *TestSlice12) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	a, err := v.Array()
	if err != nil {
		return err
	}
	*s = make([]int64, len(a))
	for i, v := range a {
		var value int64
		value, err = v.Int64()
		if err != nil {
			return newParsingError(fmt.Sprintf("%d", i), err)
		}
		(*s)[i] = int64(value)
	}
	return nil
}

// jsonParserTestSlice13 used for pooling Parsers for TestSlice13 JSONs.
var jsonParserTestSlice13 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice13) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice13.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice13.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *TestSlice13) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	a, err := v.Array()
	if err != nil {
		return err
	}
	*s = make([]test_extr.External, len(a))
	for i, v := range a {
		var value test_extr.External
		err = value.FillFromJSON(v)
		if err != nil {
			return newParsingError(fmt.Sprintf("%d", i), err)
		}
		(*s)[i] = test_extr.External(value)
	}
	return nil
}

// jsonParserTestSlice14 used for pooling Parsers for TestSlice14 JSONs.
var jsonParserTestSlice14 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice14) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice14.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice14.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *TestSlice14) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	a, err := v.Array()
	if err != nil {
		return err
	}
	if len(*s) != len(a) {
		return fmt.Errorf("expected %d elements, got %d", len(*s), len(a))
	}
	for i, v := range a {
		b, err := v.StringBytes()
		if err != nil {
			return newParsingError("", err)
		}
		value, err := parseDateTime(string(b))
		if err != nil {
			return newParsingError(fmt.Sprintf("%d", i), err)
		}
		(*s)[i] = time.Time(value)
	}
	return nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap10) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestMap10) MarshalTo(result *jwriter.Writer) error {
	if s == nil || *s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	for _k, _v := range *s {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.String(string(_k))
		result.RawByte(':')
		result.Int64(int64(_v))
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestMap10) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s TestMap10) Reset() {
	for k, v := range s {
		v = 0
		s[k] = v
	}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap11) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestMap11) MarshalTo(result *jwriter.Writer) error {
	if s == nil || *s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	for _k, _v := range *s {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.String(string(_k))
		result.RawByte(':')
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "TestMap11" attribute %q: %w`, _k, err)
		}
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestMap11) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s TestMap11) Reset() {
	for k, v := range s {
		v.Reset()
		s[k] = v
	}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap11Ref) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestMap11Ref) MarshalTo(result *jwriter.Writer) error {
	if s == nil || *s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	for _k, _v := range *s {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.String(string(_k))
		result.RawByte(':')
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "TestMap11Ref" attribute %q: %w`, _k, err)
		}
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestMap11Ref) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s TestMap11Ref) Reset() {
	for k, v := range s {
		v = nil
		s[k] = v
	}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSlice12) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice12) MarshalTo(result *jwriter.Writer) error {
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
		result.Int64(int64(_v))
	}
	result.RawByte(']')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestSlice12) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestSlice12) Reset() {
	*s = (*s)[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSlice13) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice13) MarshalTo(result *jwriter.Writer) error {
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
			return fmt.Errorf(`can't marshal "TestSlice13" value at position %d: %w`, _k, err)
		}
	}
	result.RawByte(']')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestSlice13) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestSlice13) Reset() {
	for i := range *s {
		(*s)[i].Reset()
	}
	*s = (*s)[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSlice14) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice14) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
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
		writeTime(result, _v, time.RFC3339Nano)
	}
	result.RawByte(']')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestSlice14) IsZero() bool {
	for _, _v := range s {
		if _v.IsZero() {
			return false
		}
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestSlice14) Reset() {
	*s = TestSlice14{}
}
