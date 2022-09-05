// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test

import (
	"bytes"
	"fmt"
	"github.com/valyala/fastjson"
)

// UnmarshalJSON implements json.Unmarshaler
func (s *Struct) UnmarshalJSON(data []byte) error {
	parser := structPool.Get()
	defer structPool.Put(parser)
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
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
		s.Filter = filter.String()
	}
	if limit := v.Get("limit"); limit != nil {
		s.Limit, err = limit.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	}
	if offset := v.Get("offset"); offset != nil {
		s.Offset, err = offset.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%soffset' value: %w", objPath, err)
		}
	} else {
		s.Offset = 100
	}
	if nested := v.Get("nested"); nested != nil {
		err = s.Nested.FillFromJson(nested, objPath+"nested.")
		if err != nil {
			return fmt.Errorf("error parsing '%snested' value: %w", objPath, err)
		}
	}
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
}
