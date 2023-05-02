// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_trans

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
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
