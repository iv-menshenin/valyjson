// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_extr

import (
	"bytes"
	"fmt"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/fastjson"
)

// jsonParserExternalNested used for pooling Parsers for ExternalNested JSONs.
var jsonParserExternalNested fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *ExternalNested) UnmarshalJSON(data []byte) error {
	parser := jsonParserExternalNested.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserExternalNested.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *ExternalNested) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _testAllOfSecond := v; _testAllOfSecond != nil {
		var valTestAllOfSecond = &s.TestAllOfSecond
		err = valTestAllOfSecond.FillFromJSON(_testAllOfSecond)
		if err != nil {
			return newParsingError("", err)
		}
	}
	if _testAllOfThird := v; _testAllOfThird != nil {
		var valTestAllOfThird = &s.TestAllOfThird
		err = valTestAllOfThird.FillFromJSON(_testAllOfThird)
		if err != nil {
			return newParsingError("", err)
		}
	}
	if _testStr01 := v; _testStr01 != nil {
		var valTestStr01 = &s.TestStr01
		err = valTestStr01.FillFromJSON(_testStr01)
		if err != nil {
			return newParsingError("", err)
		}
	}
	return nil
}

// validate checks for correct data structure
func (s *ExternalNested) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [3]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *ExternalNested) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *ExternalNested) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if !s.TestAllOfSecond.IsZero() {
		if wantComma {
			result.RawByte(',')
		}
		result.Raw(unpackObject(s.TestAllOfSecond.MarshalJSON()))
		wantComma = true
	}
	if !s.TestAllOfThird.IsZero() {
		if wantComma {
			result.RawByte(',')
		}
		result.Raw(unpackObject(s.TestAllOfThird.MarshalJSON()))
		wantComma = true
	}
	if !s.TestStr01.IsZero() {
		if wantComma {
			result.RawByte(',')
		}
		result.Raw(unpackObject(s.TestStr01.MarshalJSON()))
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s ExternalNested) IsZero() bool {
	if !s.TestAllOfSecond.IsZero() {
		return false
	}
	if !s.TestAllOfThird.IsZero() {
		return false
	}
	if !s.TestStr01.IsZero() {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *ExternalNested) Reset() {
	s.TestAllOfSecond.Reset()
	s.TestAllOfThird.Reset()
	s.TestStr01.Reset()
}
