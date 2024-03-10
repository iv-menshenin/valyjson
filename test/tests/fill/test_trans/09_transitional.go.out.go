// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_trans

import (
	"bytes"
	"fmt"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/fastjson"

	"fill/test_extr"
)

// jsonParserTestTransitional used for pooling Parsers for TestTransitional JSONs.
var jsonParserTestTransitional fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestTransitional) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestTransitional.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestTransitional.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestTransitional) FillFromJSON(v *fastjson.Value) (err error) {
	var _val TestTransitionalElem
	err = _val.FillFromJSON(v)
	if err != nil {
		return err
	}
	*s = TestTransitional(_val)
	return nil
}

// jsonParserTestTransitionalElem used for pooling Parsers for TestTransitionalElem JSONs.
var jsonParserTestTransitionalElem fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestTransitionalElem) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestTransitionalElem.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestTransitionalElem.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestTransitionalElem) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _testField := v.Get("test-field"); _testField != nil {
		var valTestField int64
		valTestField, err = _testField.Int64()
		if err != nil {
			return newParsingError("test-field", err)
		}
		s.TestField = valTestField
	}
	return nil
}

// validate checks for correct data structure
func (s *TestTransitionalElem) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 's', 't', '-', 'f', 'i', 'e', 'l', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserTestExternalNested used for pooling Parsers for TestExternalNested JSONs.
var jsonParserTestExternalNested fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestExternalNested) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestExternalNested.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestExternalNested.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestExternalNested) FillFromJSON(v *fastjson.Value) (err error) {
	var _val test_extr.ExternalNested
	err = _val.FillFromJSON(v)
	if err != nil {
		return err
	}
	*s = TestExternalNested(_val)
	return nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestTransitional) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestTransitional) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	return (*TestTransitionalElem)(s).MarshalTo(result)
}

// IsZero shows whether the object is an empty value.
func (s TestTransitional) IsZero() bool {
	return TestTransitionalElem(s).IsZero()
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestTransitional) Reset() {
	var tmp = (*TestTransitionalElem)(s)
	tmp.Reset()
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestTransitionalElem) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestTransitionalElem) MarshalTo(result *jwriter.Writer) error {
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
	if s.TestField != 0 {
		result.RawString(`"test-field":`)
		result.Int64(s.TestField)
		wantComma = true
	} else {
		result.RawString(`"test-field":0`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestTransitionalElem) IsZero() bool {
	if s.TestField != 0 {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestTransitionalElem) Reset() {
	s.TestField = 0
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestExternalNested) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestExternalNested) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	return (*test_extr.ExternalNested)(s).MarshalTo(result)
}

// IsZero shows whether the object is an empty value.
func (s TestExternalNested) IsZero() bool {
	return test_extr.ExternalNested(s).IsZero()
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestExternalNested) Reset() {
	var tmp = (*test_extr.ExternalNested)(s)
	tmp.Reset()
}
