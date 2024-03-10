// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_uuid

import (
	"bytes"
	"fmt"

	"github.com/google/uuid"
	"github.com/mailru/easyjson/jwriter"
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestUUID) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestUUID) MarshalTo(result *jwriter.Writer) error {
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
	if buf, err := s.UUID.MarshalText(); err != nil {
		return newParsingError("uuid", err)
	} else {
		result.RawString(`"uuid":"`)
		result.Buffer.AppendBytes(buf)
		result.RawByte('"')
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestUUID) IsZero() bool {
	if s.UUID != uuid.Nil {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestUUID) Reset() {
	s.UUID = uuid.Nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *InheritUUID2) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *InheritUUID2) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	return (*InheritUUID)(s).MarshalTo(result)
}

// IsZero shows whether the object is an empty value.
func (s InheritUUID2) IsZero() bool {
	return s == InheritUUID2(uuid.Nil)
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *InheritUUID2) Reset() {
	var tmp = (*InheritUUID)(s)
	*tmp = InheritUUID(uuid.Nil)
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *InheritUUID) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *InheritUUID) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	_uuid, err := uuid.UUID(*s).MarshalText()
	if err == nil {
		result.RawByte('"')
		result.Buffer.AppendBytes(_uuid)
		result.RawByte('"')
	}
	return err
}

// IsZero shows whether the object is an empty value.
func (s InheritUUID) IsZero() bool {
	return s == InheritUUID(uuid.Nil)
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *InheritUUID) Reset() {
	var tmp = (*uuid.UUID)(s)
	*tmp = uuid.Nil
}
