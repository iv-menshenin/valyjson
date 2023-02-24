// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_extr

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *External) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field := v.Get("field"); _field != nil {
		var valField []byte
		if valField, err = _field.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.Field = string(valField)
	}
	if _fieldRef := v.Get("fieldRef"); valueIsNotNull(_fieldRef) {
		var valFieldRef []byte
		if valFieldRef, err = _fieldRef.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sfieldRef' value: %w", objPath, err)
		}
		s.FieldRef = new(string)
		*s.FieldRef = string(valFieldRef)
	}
	return nil
}

// validate checks for correct data structure
func (s *External) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [5]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', 'R', 'e', 'f'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}
