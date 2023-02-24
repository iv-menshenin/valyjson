// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_string

import (
	"bytes"
	"fmt"
	"unsafe"

	"github.com/valyala/fastjson"
)

// jsonParserTestStr01 used for pooling Parsers for TestStr01 JSONs.
var jsonParserTestStr01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestStr01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestStr01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestStr01.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestStr01) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field := v.Get("field"); _field != nil {
		var valField []byte
		if valField, err = _field.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.Field = *(*string)(unsafe.Pointer(&valField))
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
func (s *TestStr01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', 'R', 'e', 'f'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}

// jsonParserTestStr02 used for pooling Parsers for TestStr02 JSONs.
var jsonParserTestStr02 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestStr02) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestStr02.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestStr02.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestStr02) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field := v.Get("field"); _field != nil {
		var valField []byte
		if valField, err = _field.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.Field = *(*string)(unsafe.Pointer(&valField))
	}
	if _fieldRef := v.Get("fieldRef"); valueIsNotNull(_fieldRef) {
		var valFieldRef []byte
		if valFieldRef, err = _fieldRef.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sfieldRef' value: %w", objPath, err)
		}
		s.FieldRef = new(string)
		*s.FieldRef = string(valFieldRef)
	}
	if _string := v.Get("string"); _string != nil {
		var valString []byte
		if valString, err = _string.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sstring' value: %w", objPath, err)
		}
		s.String = *(*FieldValueString)(unsafe.Pointer(&valString))
	}
	return nil
}

// validate checks for correct data structure
func (s *TestStr02) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [3]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', 'R', 'e', 'f'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 't', 'r', 'i', 'n', 'g'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}
