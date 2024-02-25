// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_string

import (
	"bytes"
	"fmt"
	"unsafe"

	"github.com/mailru/easyjson/jwriter"
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestStr01) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _field := v.Get("field"); _field != nil {
		var valField []byte
		if valField, err = _field.StringBytes(); err != nil {
			return newParsingError("field", err)
		}
		s.Field = string(valField)
	}
	if _fieldRef := v.Get("fieldRef"); valueIsNotNull(_fieldRef) {
		var valFieldRef []byte
		if valFieldRef, err = _fieldRef.StringBytes(); err != nil {
			return newParsingError("fieldRef", err)
		}
		s.FieldRef = (*string)(unsafe.Pointer(&valFieldRef))
	}
	if _defRef := v.Get("defRef"); valueIsNotNull(_defRef) {
		var valDefRef []byte
		if valDefRef, err = _defRef.StringBytes(); err != nil {
			return newParsingError("defRef", err)
		}
		s.DefRef = (*string)(unsafe.Pointer(&valDefRef))
	} else {
		if _defRef == nil {
			var __DefRef string = "default"
			s.DefRef = &__DefRef
		}
	}
	return nil
}

// validate checks for correct data structure
func (s *TestStr01) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', 'R', 'e', 'f'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 'f', 'R', 'e', 'f'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestStr02) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _field := v.Get("field"); _field != nil {
		var valField []byte
		if valField, err = _field.StringBytes(); err != nil {
			return newParsingError("field", err)
		}
		s.Field = string(valField)
	}
	if _fieldRef := v.Get("fieldRef"); valueIsNotNull(_fieldRef) {
		var valFieldRef []byte
		if valFieldRef, err = _fieldRef.StringBytes(); err != nil {
			return newParsingError("fieldRef", err)
		}
		s.FieldRef = (*string)(unsafe.Pointer(&valFieldRef))
	}
	if _string := v.Get("string"); _string != nil {
		var valString []byte
		if valString, err = _string.StringBytes(); err != nil {
			return newParsingError("string", err)
		}
		s.String = FieldValueString(valString)
	} else {
		s.String = "value-foo-bar"
	}
	return nil
}

// validate checks for correct data structure
func (s *TestStr02) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', 'R', 'e', 'f'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 't', 'r', 'i', 'n', 'g'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestStr01) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestStr01) MarshalTo(result *jwriter.Writer) error {
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
	if s.Field != "" {
		result.RawString(`"field":`)
		result.String(s.Field)
		wantComma = true
	} else {
		result.RawString(`"field":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.FieldRef != nil {
		result.RawString(`"fieldRef":`)
		result.String(*s.FieldRef)
		wantComma = true
	} else {
		result.RawString(`"fieldRef":null`)
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.DefRef != nil {
		result.RawString(`"defRef":`)
		result.String(*s.DefRef)
		wantComma = true
	} else {
		result.RawString(`"defRef":null`)
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestStr01) IsZero() bool {
	if s.Field != "" {
		return false
	}
	if s.FieldRef != nil {
		return false
	}
	if s.DefRef != nil {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestStr01) Reset() {
	s.Field = ""
	s.FieldRef = nil
	s.DefRef = nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestStr02) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestStr02) MarshalTo(result *jwriter.Writer) error {
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
	if s.Field != "" {
		result.RawString(`"field":`)
		result.String(s.Field)
		wantComma = true
	} else {
		result.RawString(`"field":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.FieldRef != nil {
		result.RawString(`"fieldRef":`)
		result.String(*s.FieldRef)
		wantComma = true
	} else {
		result.RawString(`"fieldRef":null`)
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.String != "" {
		result.RawString(`"string":`)
		result.String(string(s.String))
		wantComma = true
	} else {
		result.RawString(`"string":""`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestStr02) IsZero() bool {
	if s.Field != "" {
		return false
	}
	if s.FieldRef != nil {
		return false
	}
	if s.String != "" {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestStr02) Reset() {
	s.Field = ""
	s.FieldRef = nil
	s.String = FieldValueString("")
}
