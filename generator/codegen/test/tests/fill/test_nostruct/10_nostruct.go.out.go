// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_nostruct

import (
	"fmt"
	"time"

	"github.com/valyala/fastjson"

	"fill/test_extr"
)

// jsonParserTestMap10 used for pooling Parsers for TestMap10 JSONs.
var jsonParserTestMap10 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap10) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap10.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap10.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the keys with fastjson.Value
func (s *TestMap10) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	o, err := v.Object()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make(map[string]int64, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value int64
		value, err = v.Int64()
		if err != nil {
			err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
			return
		}
		(*s)[string(key)] = int64(value)
	})
	return err
}

// jsonParserTestMap11 used for pooling Parsers for TestMap11 JSONs.
var jsonParserTestMap11 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap11) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap11.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap11.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the keys with fastjson.Value
func (s *TestMap11) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	o, err := v.Object()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make(map[string]test_extr.External, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value test_extr.External
		err = value.FillFromJson(v, objPath+".")
		if err != nil {
			err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
			return
		}
		(*s)[string(key)] = test_extr.External(value)
	})
	return err
}

// jsonParserTestSlice12 used for pooling Parsers for TestSlice12 JSONs.
var jsonParserTestSlice12 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice12) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice12.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice12.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson fills the array with the values recognized from fastjson.Value
func (s *TestSlice12) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	a, err := v.Array()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make([]int64, len(a))
	for i, v := range a {
		var value int64
		value, err = v.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%s[%d]' value: %w", objPath, i, err)
		}
		(*s)[i] = int64(value)
	}
	return nil
}

// jsonParserTestSlice13 used for pooling Parsers for TestSlice13 JSONs.
var jsonParserTestSlice13 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice13) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice13.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice13.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson fills the array with the values recognized from fastjson.Value
func (s *TestSlice13) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	a, err := v.Array()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make([]test_extr.External, len(a))
	for i, v := range a {
		var value test_extr.External
		err = value.FillFromJson(v, objPath+".")
		if err != nil {
			return fmt.Errorf("error parsing '%s[%d]' value: %w", objPath, i, err)
		}
		(*s)[i] = test_extr.External(value)
	}
	return nil
}

// jsonParserTestSlice14 used for pooling Parsers for TestSlice14 JSONs.
var jsonParserTestSlice14 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice14) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice14.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice14.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson fills the array with the values recognized from fastjson.Value
func (s *TestSlice14) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	a, err := v.Array()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	if len(*s) != len(a) {
		return fmt.Errorf("error parsing '%s', expected %d elements, got %d", objPath, len(*s), len(a))
	}
	for i, v := range a {
		b, err := v.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
		}
		value, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s[%d]' value: %w", objPath, i, err)
		}
		(*s)[i] = time.Time(value)
	}
	return nil
}
