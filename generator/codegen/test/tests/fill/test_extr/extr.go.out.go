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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *External) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _test01 := v.Get("test1"); _test01 != nil {
		var valTest01 test_any.TestAllOfSecond
		err = valTest01.FillFromJson(_test01, objPath+"test1.")
		if err != nil {
			return fmt.Errorf("error parsing '%stest1' value: %w", objPath, err)
		}
		s.Test01 = valTest01
	}
	if _test02 := v.Get("test2"); _test02 != nil {
		var valTest02 test_string.TestStr01
		err = valTest02.FillFromJson(_test02, objPath+"test2.")
		if err != nil {
			return fmt.Errorf("error parsing '%stest2' value: %w", objPath, err)
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
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 's', 't', '2'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}
