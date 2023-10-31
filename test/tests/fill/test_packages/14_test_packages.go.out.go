// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_packages

import (
	"bytes"
	"fmt"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/fastjson"

	pack_a "fill/test_packages/pack_b"
)

// jsonParserTest01 used for pooling Parsers for Test01 JSONs.
var jsonParserTest01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Test01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTest01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTest01.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Test01) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _field := v.Get("field"); _field != nil {
		var valField int64
		valField, err = _field.Int64()
		if err != nil {
			return newParsingError("field", err)
		}
		s.Field = pack_a.Test(valField)
	}
	return nil
}

// validate checks for correct data structure
func (s *Test01) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Test01) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Test01) MarshalTo(result *jwriter.Writer) error {
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
	if s.Field != 0 {
		result.RawString(`"field":`)
		writeInt64(result, int64(s.Field))
		wantComma = true
	} else {
		result.RawString(`"field":0`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s Test01) IsZero() bool {
	if s.Field != 0 {
		return false
	}
	return true
}
