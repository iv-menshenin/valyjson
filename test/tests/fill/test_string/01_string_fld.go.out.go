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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestStr01) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field := v.Get("field"); _field != nil {
		var valField []byte
		if valField, err = _field.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.Field = *(*string)(unsafe.Pointer(&valField))
	}
	if _fieldRef := v.Get("fieldRef"); valueIsNotNull(_fieldRef) {
		var valFieldRef []byte
		if valFieldRef, err = _fieldRef.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.fieldRef' value: %w", objPath, err)
		}
		s.FieldRef = (*string)(unsafe.Pointer(&valFieldRef))
	}
	if _defRef := v.Get("defRef"); valueIsNotNull(_defRef) {
		var valDefRef []byte
		if valDefRef, err = _defRef.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.defRef' value: %w", objPath, err)
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
func (s *TestStr01) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', 'R', 'e', 'f'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 'f', 'R', 'e', 'f'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestStr02) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field := v.Get("field"); _field != nil {
		var valField []byte
		if valField, err = _field.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.Field = *(*string)(unsafe.Pointer(&valField))
	}
	if _fieldRef := v.Get("fieldRef"); valueIsNotNull(_fieldRef) {
		var valFieldRef []byte
		if valFieldRef, err = _fieldRef.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.fieldRef' value: %w", objPath, err)
		}
		s.FieldRef = (*string)(unsafe.Pointer(&valFieldRef))
	}
	if _string := v.Get("string"); _string != nil {
		var valString []byte
		if valString, err = _string.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.string' value: %w", objPath, err)
		}
		s.String = *(*FieldValueString)(unsafe.Pointer(&valString))
	} else {
		s.String = "value-foo-bar"
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', 'R', 'e', 'f'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 't', 'r', 'i', 'n', 'g'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestStr01) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestStr01) MarshalTo(result Writer) error {
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
	if s.Field != "" {
		result.WriteString(`"field":`)
		writeString(result, s.Field)
		wantComma = true
	} else {
		result.WriteString(`"field":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.FieldRef != nil {
		result.WriteString(`"fieldRef":`)
		writeString(result, *s.FieldRef)
		wantComma = true
	} else {
		result.WriteString(`"fieldRef":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.DefRef != nil {
		result.WriteString(`"defRef":`)
		writeString(result, *s.DefRef)
		wantComma = true
	} else {
		result.WriteString(`"defRef":null`)
	}
	result.WriteString("}")
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestStr02) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestStr02) MarshalTo(result Writer) error {
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
	if s.Field != "" {
		result.WriteString(`"field":`)
		writeString(result, s.Field)
		wantComma = true
	} else {
		result.WriteString(`"field":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.FieldRef != nil {
		result.WriteString(`"fieldRef":`)
		writeString(result, *s.FieldRef)
		wantComma = true
	} else {
		result.WriteString(`"fieldRef":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.String != "" {
		result.WriteString(`"string":`)
		writeString(result, string(s.String))
		wantComma = true
	} else {
		result.WriteString(`"string":""`)
		wantComma = true
	}
	result.WriteString("}")
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
