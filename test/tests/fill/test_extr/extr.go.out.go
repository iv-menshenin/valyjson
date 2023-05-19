// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_extr

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
)

// jsonParserExternal used for pooling Parsers for External JSONs.
var jsonParserExternal fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *External) UnmarshalJSON(data []byte) error {
	parser := jsonParserExternal.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserExternal.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *External) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _test01 := v.Get("test1"); _test01 != nil {
		var valTest01 test_any.TestAllOfSecond
		err = valTest01.FillFromJSON(_test01)
		if err != nil {
			return newParsingError("test1", err)
		}
		s.Test01 = valTest01
	}
	if _test02 := v.Get("test2"); _test02 != nil {
		var valTest02 test_string.TestStr01
		err = valTest02.FillFromJSON(_test02)
		if err != nil {
			return newParsingError("test2", err)
		}
		s.Test02 = valTest02
	}
	return nil
}

// validate checks for correct data structure
func (s *External) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 's', 't', '1'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 's', 't', '2'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

var bufDataExternal = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *External) MarshalJSON() ([]byte, error) {
	var result = bufDataExternal.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *External) MarshalTo(result Writer) error {
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
	result.WriteString(`"test1":`)
	if err = s.Test01.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "test1" attribute: %w`, err)
	}
	wantComma = true
	if !s.Test02.IsZero() {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"test2":`)
		if err = s.Test02.MarshalTo(result); err != nil {
			return fmt.Errorf(`can't marshal "test2" attribute: %w`, err)
		}
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s External) IsZero() bool {
	if !s.Test01.IsZero() {
		return false
	}
	if !s.Test02.IsZero() {
		return false
	}
	return true
}
