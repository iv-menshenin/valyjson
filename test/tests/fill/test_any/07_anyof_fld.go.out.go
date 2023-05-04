// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_any

import (
	"bytes"
	"fmt"
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
		s.Comment = *(*string)(unsafe.Pointer(&valComment))
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
		s.Command = *(*string)(unsafe.Pointer(&valCommand))
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

var bufDataTestOneOfInteger = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestOneOfInteger) MarshalJSON() ([]byte, error) {
	var result = bufDataTestOneOfInteger.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestOneOfInteger) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	writeInt64(result, int64(*s))
	return nil
}

// IsZero shows whether the object is an empty value.
func (s TestOneOfInteger) IsZero() bool {
	return s == TestOneOfInteger(0)
}

var bufDataTestOneOfString = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestOneOfString) MarshalJSON() ([]byte, error) {
	var result = bufDataTestOneOfString.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestOneOfString) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	writeString(result, string(*s))
	return nil
}

// IsZero shows whether the object is an empty value.
func (s TestOneOfString) IsZero() bool {
	return s == TestOneOfString("")
}

var bufDataTestOneOfStruct = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestOneOfStruct) MarshalJSON() ([]byte, error) {
	var result = bufDataTestOneOfStruct.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestOneOfStruct) MarshalTo(result Writer) error {
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
	if s.Class != "" {
		result.WriteString(`"class":`)
		writeString(result, s.Class)
	} else {
		result.WriteString(`"class":""`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Value != 0 {
		result.WriteString(`"width":`)
		writeFloat64(result, s.Value)
	} else {
		result.WriteString(`"width":0`)
	}
	result.WriteString("}")
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

var bufDataTestAllOfSecond = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestAllOfSecond) MarshalJSON() ([]byte, error) {
	var result = bufDataTestAllOfSecond.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestAllOfSecond) MarshalTo(result Writer) error {
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
	if s.Comment != "" {
		result.WriteString(`"comment":`)
		writeString(result, s.Comment)
	} else {
		result.WriteString(`"comment":""`)
	}
	if s.Level != 0 {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"level":`)
		writeInt64(result, s.Level)
	}
	result.WriteString("}")
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

var bufDataTestAllOfThird = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestAllOfThird) MarshalJSON() ([]byte, error) {
	var result = bufDataTestAllOfThird.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestAllOfThird) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if s.Command != "" {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"command":`)
		writeString(result, s.Command)
	}
	if s.Range != 0 {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"range":`)
		writeInt64(result, s.Range)
	}
	result.WriteString("}")
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
