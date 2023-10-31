// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_any

import (
	"bytes"
	"fmt"

	"github.com/mailru/easyjson/jwriter"
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestAllOfSecond) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _comment := v.Get("comment"); _comment != nil {
		var valComment []byte
		if valComment, err = _comment.StringBytes(); err != nil {
			return newParsingError("comment", err)
		}
		s.Comment = string(valComment)
	}
	if _level := v.Get("level"); _level != nil {
		var valLevel int64
		valLevel, err = _level.Int64()
		if err != nil {
			return newParsingError("level", err)
		}
		s.Level = valLevel
	}
	return nil
}

// validate checks for correct data structure
func (s *TestAllOfSecond) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'l', 'e', 'v', 'e', 'l'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestAllOfThird) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _command := v.Get("command"); _command != nil {
		var valCommand []byte
		if valCommand, err = _command.StringBytes(); err != nil {
			return newParsingError("command", err)
		}
		s.Command = string(valCommand)
	}
	if _range := v.Get("range"); _range != nil {
		var valRange int64
		valRange, err = _range.Int64()
		if err != nil {
			return newParsingError("range", err)
		}
		s.Range = valRange
	}
	return nil
}

// validate checks for correct data structure
func (s *TestAllOfThird) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 'n', 'g', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestOneOfInteger) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestOneOfInteger) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	writeInt64(result, int64(*s))
	return nil
}

// IsZero shows whether the object is an empty value.
func (s TestOneOfInteger) IsZero() bool {
	return s == TestOneOfInteger(0)
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestOneOfString) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestOneOfString) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	result.String(string(*s))
	return nil
}

// IsZero shows whether the object is an empty value.
func (s TestOneOfString) IsZero() bool {
	return s == TestOneOfString("")
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestOneOfStruct) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestOneOfStruct) MarshalTo(result *jwriter.Writer) error {
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
	if s.Class != "" {
		result.RawString(`"class":`)
		result.String(s.Class)
		wantComma = true
	} else {
		result.RawString(`"class":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Value != 0 {
		result.RawString(`"width":`)
		writeFloat64(result, s.Value)
		wantComma = true
	} else {
		result.RawString(`"width":0`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestOneOfStruct) IsZero() bool {
	if s.Class != "" {
		return false
	}
	if s.Value != 0 {
		return false
	}
	return true
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestAllOfSecond) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestAllOfSecond) MarshalTo(result *jwriter.Writer) error {
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
	if s.Comment != "" {
		result.RawString(`"comment":`)
		result.String(s.Comment)
		wantComma = true
	} else {
		result.RawString(`"comment":""`)
		wantComma = true
	}
	if s.Level != 0 {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"level":`)
		writeInt64(result, s.Level)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestAllOfSecond) IsZero() bool {
	if s.Comment != "" {
		return false
	}
	if s.Level != 0 {
		return false
	}
	return true
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestAllOfThird) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestAllOfThird) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if s.Command != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"command":`)
		result.String(s.Command)
		wantComma = true
	}
	if s.Range != 0 {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"range":`)
		writeInt64(result, s.Range)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestAllOfThird) IsZero() bool {
	if s.Command != "" {
		return false
	}
	if s.Range != 0 {
		return false
	}
	return true
}
