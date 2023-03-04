// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_uuid

import (
	"bytes"
	"fmt"

	"github.com/google/uuid"
	"github.com/valyala/fastjson"
)

// jsonParserTestUUID used for pooling Parsers for TestUUID JSONs.
var jsonParserTestUUID fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestUUID) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestUUID.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestUUID.Put(parser)
	return s.FillFromJSON(v, "")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestUUID) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _uUID := v.Get("uuid"); _uUID != nil {
		var valUUID uuid.UUID
		b, err := _uUID.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%suuid' value: %w", objPath, err)
		}
		valUUID, err = uuid.ParseBytes(b)
		if err != nil {
			return fmt.Errorf("error parsing '%suuid' value: %w", objPath, err)
		}
		s.UUID = valUUID
	}
	return nil
}

// validate checks for correct data structure
func (s *TestUUID) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'u', 'u', 'i', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}
