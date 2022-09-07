// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test

import (
	"bytes"
	"fmt"
	"github.com/valyala/fastjson"
)

// jsonParserStructused for pooling Parsers for Struct JSONs.
var jsonParserStruct fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Struct) UnmarshalJSON(data []byte) error {
	parser := jsonParserStruct.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserStruct.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// only if there is a strict rules
	if err = s.validate(v, ""); err != nil {
		return err
	}
	if filter := v.Get("filter"); filter != nil {
		if filter.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", filter.Type())
			return fmt.Errorf("error parsing '%sfilter' value: %w", objPath, err)
		}
		xFilter := filter.String()
		s.Filter = xFilter
	} else {
		return fmt.Errorf("required element '%sfilter' is missing", objPath)
	}
	if limit := v.Get("limit"); limit != nil {
		var xLimit int
		xLimit, err = limit.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
		s.Limit = xLimit
	}
	if offset := v.Get("offset"); offset != nil {
		var xOffset int
		xOffset, err = offset.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%soffset' value: %w", objPath, err)
		}
		s.Offset = xOffset
	} else {
		s.Offset = 100
	}
	if nested := v.Get("nested"); nested != nil {
		err = s.Nested.FillFromJson(nested, objPath+"nested.")
		if err != nil {
			return fmt.Errorf("error parsing '%snested' value: %w", objPath, err)
		}
	}
	return nil
}

// validate checks for correct data structure
func (s *Struct) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'l', 't', 'e', 'r'}) {
			return
		}
		if bytes.Equal(key, []byte{'l', 'i', 'm', 'i', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'o', 'f', 'f', 's', 'e', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'n', 'e', 's', 't', 'e', 'd'}) {
			return
		}
		if objPath == "" {
			err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
		} else {
			err = fmt.Errorf("unexpected field '%s' in the '%s' path", string(key), objPath)
		}
	})
	return nil
}

// jsonParserNestedused for pooling Parsers for Nested JSONs.
var jsonParserNested fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Nested) UnmarshalJSON(data []byte) error {
	parser := jsonParserNested.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserNested.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Nested) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// only if there is a strict rules
	if err = s.validate(v, ""); err != nil {
		return err
	}
	if list := v.Get("list"); list != nil {
		if err != nil {
			return fmt.Errorf("error parsing '%slist' value: %w", objPath, err)
		}
	}
	if count := v.Get("count"); count != nil {
		var xCount int64
		xCount, err = count.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%scount' value: %w", objPath, err)
		}
		s.Count = &xCount
	}
	if cross := v.Get("cross"); cross != nil {
		var xCross int64
		xCross, err = cross.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%scross' value: %w", objPath, err)
		}
		s.Cross = &xCross
	}
	return nil
}

// validate checks for correct data structure
func (s *Nested) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'l', 'i', 's', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'u', 'n', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'c', 'r', 'o', 's', 's'}) {
			return
		}
		if objPath == "" {
			err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
		} else {
			err = fmt.Errorf("unexpected field '%s' in the '%s' path", string(key), objPath)
		}
	})
	return nil
}

// valueIsNotNull allows you to determine if the value is contained in a Json structure.
// Checks if the structure itself or the value is Null.
func valueIsNotNull(v *fastjson.Value) bool {
	return v != nil && v.Type() != fastjson.TypeNull
}
