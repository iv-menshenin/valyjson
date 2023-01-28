// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package testo

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
)

// jsonParserTestMap01used for pooling Parsers for TestMap01 JSONs.
var jsonParserTestMap01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap01.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestMap01) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if tags := v.Get("tags"); valueIsNotNull(tags) {
		o, err := tags.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%stags' value: %w", objPath, err)
		}
		var valTags = make(map[string]string, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value []byte
			value, err = v.StringBytes()
			if err == nil {
				valTags[string(key)] = string(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%stags' value: %w", objPath, err)
		}
		s.Tags = valTags
	}
	if properties := v.Get("properties"); valueIsNotNull(properties) {
		o, err := properties.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%sproperties' value: %w", objPath, err)
		}
		var valProperties = make(map[string]Property, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value Property
			err = value.FillFromJson(v, objPath+"properties.")
			if err == nil {
				valProperties[string(key)] = Property(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%sproperties' value: %w", objPath, err)
		}
		s.Properties = valProperties
	}
	if keytypedproperties := v.Get("key_typed_properties"); valueIsNotNull(keytypedproperties) {
		o, err := keytypedproperties.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%skey_typed_properties' value: %w", objPath, err)
		}
		var valKeyTypedProperties = make(map[Key]Property, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value Property
			err = value.FillFromJson(v, objPath+"key_typed_properties.")
			if err == nil {
				valKeyTypedProperties[Key(key)] = Property(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%skey_typed_properties' value: %w", objPath, err)
		}
		s.KeyTypedProperties = valKeyTypedProperties
	}
	return nil
}

// validate checks for correct data structure
func (s *TestMap01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [3]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'p', 'e', 'r', 't', 'i', 'e', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'k', 'e', 'y', '_', 't', 'y', 'p', 'e', 'd', '_', 'p', 'r', 'o', 'p', 'e', 'r', 't', 'i', 'e', 's'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// jsonParserPropertyused for pooling Parsers for Property JSONs.
var jsonParserProperty fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Property) UnmarshalJSON(data []byte) error {
	parser := jsonParserProperty.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserProperty.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Property) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if name := v.Get("name"); name != nil {
		var valName []byte
		if valName, err = name.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sname' value: %w", objPath, err)
		}
		s.Name = string(valName)
	}
	if value := v.Get("value"); value != nil {
		var valValue []byte
		if valValue, err = value.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%svalue' value: %w", objPath, err)
		}
		s.Value = string(valValue)
	}
	return nil
}

// validate checks for correct data structure
func (s *Property) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'n', 'a', 'm', 'e'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'v', 'a', 'l', 'u', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}
