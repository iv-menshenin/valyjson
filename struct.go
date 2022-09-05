package main

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
)

// Struct contains all fields for struct
//  valyjson:encode,decode,strict
type Struct struct {
	Filter string `json:"filter,required"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty" default:"100"`

	Nested Nested `json:"nested"`
}

var structPool fastjson.ParserPool

func (s *Struct) UnmarshalJSON(data []byte) error {
	parser := structPool.Get()
	defer structPool.Put(parser)

	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	return s.FillFromJson(v, "")
}

func validateStructKeys(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, v *fastjson.Value) {
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
	return err
}

func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// only if there is a strict rules
	if err = validateStructKeys(v, ""); err != nil {
		return err
	}
	if filter := v.Get("filter"); filter != nil {
		if filter.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", v.Type())
			return fmt.Errorf("error parsing '%sfilter' value: %w", objPath, err)
		}
		s.Filter = filter.String()
	} else {
		return fmt.Errorf("the '%sfilter' path is required but ommitted", objPath)
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
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	} else {
		s.Offset = 100
	}
	if nested := v.Get("nested"); nested != nil {
		err = s.Nested.FillFromJson(nested, objPath+"nested.")
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	}
	return nil
}

func (s *Struct) MarshalJSON() ([]byte, error) {
	return nil, nil
}

type Nested struct {
	List  []int64 `json:"list"`
	Count *int64  `json:"count"`
	Cross *int64  `json:"cross"`
}

var nestedPool fastjson.ParserPool

func (s *Nested) UnmarshalJSON(data []byte) error {
	parser := nestedPool.Get()
	defer nestedPool.Put(parser)

	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	// only if there is a strict rules
	if err = validateNestedKeys(v, ""); err != nil {
		return err
	}
	return s.FillFromJson(v, "")
}

func validateNestedKeys(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, v *fastjson.Value) {
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
	return err
}

func (s *Nested) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if list := v.Get("list"); list != nil {
		var listA []*fastjson.Value
		listA, err = list.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%slist' value: %w", objPath, err)
		}
		s.List = make([]int64, 0, len(listA))
		for listElemNum, listElem := range listA {
			var listElemVal int64
			listElemVal, err = listElem.Int64()
			if err != nil {
				return fmt.Errorf("error parsing '%slist[%d]' value: %w", objPath, listElemNum, err)
			}
			s.List = append(s.List, listElemVal)
		}
	}
	if count := v.Get("count"); count != nil {
		var countVal int64
		countVal, err = count.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%scount' value: %w", objPath, err)
		}
		s.Count = &countVal
	}
	if count := v.Get("cross"); valueIsNotNull(count) {
		var countVal int64
		countVal, err = count.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%scross' value: %w", objPath, err)
		}
		s.Count = &countVal
	}
	return nil
}

func valueIsNotNull(v *fastjson.Value) bool {
	return v != nil && v.Type() != fastjson.TypeNull
}
