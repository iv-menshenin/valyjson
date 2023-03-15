// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_extr

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"

	"fill/test_any"
	"fill/test_string"
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *External) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _test01 := v.Get("test1"); _test01 != nil {
		var valTest01 test_any.TestAllOfSecond
		err = valTest01.FillFromJSON(_test01, objPath+".test1")
		if err != nil {
			return fmt.Errorf("error parsing '%s.test1' value: %w", objPath, err)
		}
		s.Test01 = valTest01
	}
	if _test02 := v.Get("test2"); _test02 != nil {
		var valTest02 test_string.TestStr01
		err = valTest02.FillFromJSON(_test02, objPath+".test2")
		if err != nil {
			return fmt.Errorf("error parsing '%s.test2' value: %w", objPath, err)
		}
		s.Test02 = valTest02
	}
	return nil
}

// validate checks for correct data structure
func (s *External) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 's', 't', '2'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *External) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *External) MarshalTo(result Writer) error {
	if s == nil {
		writeString(result, "null")
		return nil
	}
	var err error
	result.Write([]byte{'{'})
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if err = s.Test01.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if err = s.Test02.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
	}
	result.Write([]byte{'}'})
	return err
}
