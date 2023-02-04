// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package testo

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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestTransitional) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	return (*TestTransitionalElem)(s).FillFromJson(v, objPath)
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestTransitionalElem) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if testfield := v.Get("test-field"); testfield != nil {
		var valTestField int64
		valTestField, err = testfield.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%stest-field' value: %w", objPath, err)
		}
		s.TestField = valTestField
	}
	return nil
}

// validate checks for correct data structure
func (s *TestTransitionalElem) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}
