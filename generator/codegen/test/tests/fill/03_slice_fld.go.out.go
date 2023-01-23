// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package testo

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
)

// jsonParserTestSlice01used for pooling Parsers for TestSlice01 JSONs.
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestSlice01) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, ""); err != nil {
		return err
	}
	if field := v.Get("strs"); valueIsNotNull(field) {
		var listA []*fastjson.Value
		listA, err = field.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%sstrs' value: %w", objPath, err)
		}
		valField := make([]string, 0, len(listA))
		for _, listElem := range listA {
			var elem []byte
			if elem, err = listElem.StringBytes(); err != nil {
				return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
			}
			valField = append(valField, string(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%sstrs' value: %w", objPath, err)
		}
		s.Field = valField
	}
	if fieldref := v.Get("ints"); valueIsNotNull(fieldref) {
		var listA []*fastjson.Value
		listA, err = fieldref.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%sints' value: %w", objPath, err)
		}
		valFieldRef := make([]*int, 0, len(listA))
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
			return fmt.Errorf("error parsing '%sints' value: %w", objPath, err)
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
				err = fmt.Errorf("the '%s' field appears in the object twice [%s]", string(key), objPath)
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s' field appears in the object twice [%s]", string(key), objPath)
			}
			return
		}
		if objPath == "" {
			err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
		} else {
			err = fmt.Errorf("unexpected field '%s' in the '%s' path", string(key), objPath)
		}
	})
	return err
}
