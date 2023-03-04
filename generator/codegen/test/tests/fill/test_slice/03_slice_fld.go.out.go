// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_slice

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
)

// jsonParserTestSlice01 used for pooling Parsers for TestSlice01 JSONs.
var jsonParserTestSlice01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice01.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestSlice01) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field := v.Get("strs"); valueIsNotNull(_field) {
		var listA []*fastjson.Value
		listA, err = _field.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%s.strs' value: %w", objPath, err)
		}
		valField := s.Field[:0]
		if l := len(listA); cap(valField) < l || (l == 0 && s.Field == nil) {
			valField = make([]string, 0, len(listA))
		}
		for _, listElem := range listA {
			var elem []byte
			if elem, err = listElem.StringBytes(); err != nil {
				return fmt.Errorf("error parsing '%s.' value: %w", objPath, err)
			}
			valField = append(valField, string(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%s.strs' value: %w", objPath, err)
		}
		s.Field = valField
	}
	if _fieldRef := v.Get("ints"); valueIsNotNull(_fieldRef) {
		var listA []*fastjson.Value
		listA, err = _fieldRef.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%s.ints' value: %w", objPath, err)
		}
		valFieldRef := s.FieldRef[:0]
		if l := len(listA); cap(valFieldRef) < l || (l == 0 && s.FieldRef == nil) {
			valFieldRef = make([]*int, 0, len(listA))
		}
		for _, listElem := range listA {
			if !valueIsNotNull(listElem) {
				valFieldRef = append(valFieldRef, nil)
				continue
			}
			var elem int
			elem, err = listElem.Int()
			if err != nil {
				break
			}
			newElem := int(elem)
			valFieldRef = append(valFieldRef, &newElem)
		}
		if err != nil {
			return fmt.Errorf("error parsing '%s.ints' value: %w", objPath, err)
		}
		s.FieldRef = valFieldRef
	}
	return nil
}

// validate checks for correct data structure
func (s *TestSlice01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'s', 't', 'r', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}
