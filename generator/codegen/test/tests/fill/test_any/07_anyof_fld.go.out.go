// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_any

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/valyala/fastjson"
)

// jsonParserTestAllOfSecond used for pooling Parsers for TestAllOfSecond JSONs.
var jsonParserTestAllOfSecond fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestAllOfSecond) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestAllOfSecond.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestAllOfSecond.Put(parser)
	return s.FillFromJSON(v, "")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestAllOfSecond) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _comment := v.Get("comment"); _comment != nil {
		var valComment []byte
		if valComment, err = _comment.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%scomment' value: %w", objPath, err)
		}
		s.Comment = *(*string)(unsafe.Pointer(&valComment))
	}
	if _level := v.Get("level"); _level != nil {
		var valLevel int64
		valLevel, err = _level.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%slevel' value: %w", objPath, err)
		}
		s.Level = valLevel
	}
	return nil
}

// validate checks for correct data structure
func (s *TestAllOfSecond) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'm', 'm', 'e', 'n', 't'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'l', 'e', 'v', 'e', 'l'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}

// jsonParserTestAllOfThird used for pooling Parsers for TestAllOfThird JSONs.
var jsonParserTestAllOfThird fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestAllOfThird) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestAllOfThird.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestAllOfThird.Put(parser)
	return s.FillFromJSON(v, "")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestAllOfThird) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _command := v.Get("command"); _command != nil {
		var valCommand []byte
		if valCommand, err = _command.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%scommand' value: %w", objPath, err)
		}
		s.Command = *(*string)(unsafe.Pointer(&valCommand))
	}
	if _range := v.Get("range"); _range != nil {
		var valRange int64
		valRange, err = _range.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%srange' value: %w", objPath, err)
		}
		s.Range = valRange
	}
	return nil
}

// validate checks for correct data structure
func (s *TestAllOfThird) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'm', 'm', 'a', 'n', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 'n', 'g', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestAllOfSecond) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *TestAllOfSecond) MarshalAppend(dst []byte) ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	var (
		err    error
		buf    = make([]byte, 0, 128)
		result = bytes.NewBuffer(dst)
	)
	result.WriteRune('{')
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Comment != "" {
		result.WriteString(`"comment":`)
		buf = marshalString(buf[:0], s.Comment)
		result.Write(buf)
	} else {
		result.WriteString(`"comment":""`)
	}
	if s.Level != 0 {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"level":`)
		buf = strconv.AppendInt(buf[:0], s.Level, 10)
		result.Write(buf)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestAllOfThird) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *TestAllOfThird) MarshalAppend(dst []byte) ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	var (
		err    error
		buf    = make([]byte, 0, 128)
		result = bytes.NewBuffer(dst)
	)
	result.WriteRune('{')
	if s.Command != "" {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"command":`)
		buf = marshalString(buf[:0], s.Command)
		result.Write(buf)
	}
	if s.Range != 0 {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"range":`)
		buf = strconv.AppendInt(buf[:0], s.Range, 10)
		result.Write(buf)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}
