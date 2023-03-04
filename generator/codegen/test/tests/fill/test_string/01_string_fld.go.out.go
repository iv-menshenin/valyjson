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
	return s.FillFromJSON(v, "")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestStr01) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
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
		s.FieldRef = (*string)(unsafe.Pointer(&valFieldRef))
	}
	if _defRef := v.Get("defRef"); valueIsNotNull(_defRef) {
		var valDefRef []byte
		if valDefRef, err = _defRef.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sdefRef' value: %w", objPath, err)
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
		if bytes.Equal(key, []byte{'d', 'e', 'f', 'R', 'e', 'f'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
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
	return s.FillFromJSON(v, "")
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
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.Field = *(*string)(unsafe.Pointer(&valField))
	}
	if _fieldRef := v.Get("fieldRef"); valueIsNotNull(_fieldRef) {
		var valFieldRef []byte
		if valFieldRef, err = _fieldRef.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sfieldRef' value: %w", objPath, err)
		}
		s.FieldRef = (*string)(unsafe.Pointer(&valFieldRef))
	}
	if _string := v.Get("string"); _string != nil {
		var valString []byte
		if valString, err = _string.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sstring' value: %w", objPath, err)
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestStr01) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *TestStr01) MarshalAppend(dst []byte) ([]byte, error) {
	var result = bytes.NewBuffer(dst)
	var (
		err error
		buf = make([]byte, 0, 128)
	)
	result.WriteRune('{')
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Field != "" {
		result.WriteString(`"field":`)
		buf = marshalString(buf[:0], s.Field)
		result.Write(buf)
	} else {
		result.WriteString(`"field":""`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.FieldRef != nil {
		result.WriteString(`"fieldRef":`)
		buf = marshalString(buf[:0], *s.FieldRef)
		result.Write(buf)
	} else {
		result.WriteString(`"fieldRef":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.DefRef != nil {
		result.WriteString(`"defRef":`)
		buf = marshalString(buf[:0], *s.DefRef)
		result.Write(buf)
	} else {
		result.WriteString(`"defRef":null`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestStr02) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *TestStr02) MarshalAppend(dst []byte) ([]byte, error) {
	var result = bytes.NewBuffer(dst)
	var (
		err error
		buf = make([]byte, 0, 128)
	)
	result.WriteRune('{')
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Field != "" {
		result.WriteString(`"field":`)
		buf = marshalString(buf[:0], s.Field)
		result.Write(buf)
	} else {
		result.WriteString(`"field":""`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.FieldRef != nil {
		result.WriteString(`"fieldRef":`)
		buf = marshalString(buf[:0], *s.FieldRef)
		result.Write(buf)
	} else {
		result.WriteString(`"fieldRef":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.String != "" {
		result.WriteString(`"string":`)
		buf = marshalString(buf[:0], string(s.String))
		result.Write(buf)
	} else {
		result.WriteString(`"string":""`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}
