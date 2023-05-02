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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestUUID) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _uUID := v.Get("uuid"); _uUID != nil {
		var valUUID uuid.UUID
		b, err := _uUID.StringBytes()
		if err != nil {
			return newParsingError("uuid", err)
		}
		valUUID, err = uuid.ParseBytes(b)
		if err != nil {
			return newParsingError("uuid", err)
		}
		s.UUID = valUUID
	}
	return nil
}

// validate checks for correct data structure
func (s *TestUUID) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserInheritUUID2 used for pooling Parsers for InheritUUID2 JSONs.
var jsonParserInheritUUID2 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *InheritUUID2) UnmarshalJSON(data []byte) error {
	parser := jsonParserInheritUUID2.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserInheritUUID2.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *InheritUUID2) FillFromJSON(v *fastjson.Value) (err error) {
	var _val uuid.UUID
	b, err := v.StringBytes()
	if err != nil {
		return newParsingError("", err)
	}
	_val, err = uuid.ParseBytes(b)
	if err != nil {
		return err
	}
	*s = InheritUUID2(_val)
	return nil
}

// jsonParserInheritUUID used for pooling Parsers for InheritUUID JSONs.
var jsonParserInheritUUID fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *InheritUUID) UnmarshalJSON(data []byte) error {
	parser := jsonParserInheritUUID.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserInheritUUID.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *InheritUUID) FillFromJSON(v *fastjson.Value) (err error) {
	var _val uuid.UUID
	b, err := v.StringBytes()
	if err != nil {
		return newParsingError("", err)
	}
	_val, err = uuid.ParseBytes(b)
	if err != nil {
		return err
	}
	*s = InheritUUID(_val)
	return nil
}

var bufDataTestUUID = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestUUID) MarshalJSON() ([]byte, error) {
	var result = bufDataTestUUID.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestUUID) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if wantComma {
		result.WriteString(",")
	}
	if buf, err := s.UUID.MarshalText(); err != nil {
		return newParsingError("uuid", err)
	} else {
		result.WriteString(`"uuid":"`)
		result.Write(buf)
		result.WriteString(`"`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestUUID) IsZero() bool {
	if s.UUID != uuid.Nil {
		return false
	}
	return true
}

var bufDataInheritUUID2 = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *InheritUUID2) MarshalJSON() ([]byte, error) {
	var result = bufDataInheritUUID2.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *InheritUUID2) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	return (*InheritUUID)(s).MarshalTo(result)
}

// IsZero shows whether the object is an empty value.
func (s InheritUUID2) IsZero() bool {
	return s == InheritUUID2(uuid.Nil)
}

var bufDataInheritUUID = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *InheritUUID) MarshalJSON() ([]byte, error) {
	var result = bufDataInheritUUID.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *InheritUUID) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	_uuid, err := uuid.UUID(*s).MarshalText()
	if err == nil {
		result.Write(_uuid)
	}
	return err
}

// IsZero shows whether the object is an empty value.
func (s InheritUUID) IsZero() bool {
	return s == InheritUUID(uuid.Nil)
}
