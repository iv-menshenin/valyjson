// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_map

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/valyala/fastjson"
)

// jsonParserTestMap01 used for pooling Parsers for TestMap01 JSONs.
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestMap01) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _tags := v.Get("tags"); valueIsNotNull(_tags) {
		o, err := _tags.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.tags' value: %w", objPath, err)
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
			return fmt.Errorf("error parsing '%s.tags' value: %w", objPath, err)
		}
		s.Tags = valTags
	}
	if _properties := v.Get("properties"); valueIsNotNull(_properties) {
		o, err := _properties.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.properties' value: %w", objPath, err)
		}
		var valProperties = make(map[string]Property, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value Property
			err = value.FillFromJSON(v, objPath+".properties")
			if err == nil {
				valProperties[string(key)] = Property(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%s.properties' value: %w", objPath, err)
		}
		s.Properties = valProperties
	}
	if _keyTypedProperties := v.Get("key_typed_properties"); valueIsNotNull(_keyTypedProperties) {
		o, err := _keyTypedProperties.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.key_typed_properties' value: %w", objPath, err)
		}
		var valKeyTypedProperties = make(map[Key]Property, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value Property
			err = value.FillFromJSON(v, objPath+".key_typed_properties")
			if err == nil {
				valKeyTypedProperties[Key(key)] = Property(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%s.key_typed_properties' value: %w", objPath, err)
		}
		s.KeyTypedProperties = valKeyTypedProperties
	}
	if _integerVal := v.Get("integerVal"); valueIsNotNull(_integerVal) {
		o, err := _integerVal.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.integerVal' value: %w", objPath, err)
		}
		var valIntegerVal = make(map[Key]int32, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value int
			value, err = v.Int()
			if err == nil {
				valIntegerVal[Key(key)] = int32(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%s.integerVal' value: %w", objPath, err)
		}
		s.IntegerVal = valIntegerVal
	}
	if _floatVal := v.Get("floatVal"); valueIsNotNull(_floatVal) {
		o, err := _floatVal.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.floatVal' value: %w", objPath, err)
		}
		var valFloatVal = make(map[Key]float64, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value float64
			value, err = v.Float64()
			if err == nil {
				valFloatVal[Key(key)] = float64(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%s.floatVal' value: %w", objPath, err)
		}
		s.FloatVal = valFloatVal
	}
	if _uintVal := v.Get("uintVal"); valueIsNotNull(_uintVal) {
		o, err := _uintVal.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.uintVal' value: %w", objPath, err)
		}
		var valUintVal = make(map[Key]*uint16, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			if v.Type() == fastjson.TypeNull {
				valUintVal[Key(key)] = nil
				return
			}
			var value uint
			value, err = v.Uint()
			if err == nil {
				valUintVal[Key(key)] = (*uint16)(unsafe.Pointer(&value))
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%s.uintVal' value: %w", objPath, err)
		}
		s.UintVal = valUintVal
	}
	if _boolVal := v.Get("bool"); valueIsNotNull(_boolVal) {
		o, err := _boolVal.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.bool' value: %w", objPath, err)
		}
		var valBoolVal = make(map[Key]bool, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value bool
			value, err = v.Bool()
			if err == nil {
				valBoolVal[Key(key)] = bool(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%s.bool' value: %w", objPath, err)
		}
		s.BoolVal = valBoolVal
	}
	if _typedVal := v.Get("typed-val"); valueIsNotNull(_typedVal) {
		o, err := _typedVal.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.typed-val' value: %w", objPath, err)
		}
		var valTypedVal = make(map[Key]Val, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value uint64
			value, err = v.Uint64()
			if err == nil {
				valTypedVal[Key(key)] = Val(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%s.typed-val' value: %w", objPath, err)
		}
		s.TypedVal = valTypedVal
	}
	return nil
}

// validate checks for correct data structure
func (s *TestMap01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [8]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'p', 'e', 'r', 't', 'i', 'e', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'k', 'e', 'y', '_', 't', 'y', 'p', 'e', 'd', '_', 'p', 'r', 'o', 'p', 'e', 'r', 't', 'i', 'e', 's'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', 'e', 'g', 'e', 'r', 'V', 'a', 'l'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'l', 'o', 'a', 't', 'V', 'a', 'l'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 'i', 'n', 't', 'V', 'a', 'l'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'b', 'o', 'o', 'l'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'y', 'p', 'e', 'd', '-', 'v', 'a', 'l'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// jsonParserProperty used for pooling Parsers for Property JSONs.
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Property) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _name := v.Get("name"); _name != nil {
		var valName []byte
		if valName, err = _name.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.name' value: %w", objPath, err)
		}
		s.Name = *(*string)(unsafe.Pointer(&valName))
	}
	if _value := v.Get("value"); _value != nil {
		var valValue []byte
		if valValue, err = _value.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.value' value: %w", objPath, err)
		}
		s.Value = *(*string)(unsafe.Pointer(&valValue))
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'v', 'a', 'l', 'u', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap01) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *TestMap01) MarshalAppend(dst []byte) ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	var (
		err    error
		buf    = make([]byte, 0, 128)
		result = bytes.NewBuffer(dst)
	)
	result.WriteRune('{')
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Tags != nil {
		buf = buf[:0]
		result.WriteString(`"tags":{`)
		var _filled bool
		for _k, _v := range s.Tags {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(_k)
			result.WriteString(`":`)
			buf = marshalString(buf[:0], _v)
			result.Write(buf)
		}
		result.WriteRune('}')
	} else {
		result.WriteString(`"tags":null`)
	}
	if s.Properties != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		buf = buf[:0]
		result.WriteString(`"properties":{`)
		var _filled bool
		for _k, _v := range s.Properties {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(_k)
			result.WriteString(`":`)
			buf, err = _v.MarshalAppend(buf[:0])
			if err != nil {
				return nil, fmt.Errorf(`can't marshal "properties" attribute %q: %w`, _k, err)
			}
			result.Write(buf)
		}
		result.WriteRune('}')
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.KeyTypedProperties != nil {
		buf = buf[:0]
		result.WriteString(`"key_typed_properties":{`)
		var _filled bool
		for _k, _v := range s.KeyTypedProperties {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(string(_k))
			result.WriteString(`":`)
			buf, err = _v.MarshalAppend(buf[:0])
			if err != nil {
				return nil, fmt.Errorf(`can't marshal "key_typed_properties" attribute %q: %w`, _k, err)
			}
			result.Write(buf)
		}
		result.WriteRune('}')
	} else {
		result.WriteString(`"key_typed_properties":null`)
	}
	if s.IntegerVal != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		buf = buf[:0]
		result.WriteString(`"integerVal":{`)
		var _filled bool
		for _k, _v := range s.IntegerVal {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(string(_k))
			result.WriteString(`":`)
			buf = strconv.AppendInt(buf[:0], int64(_v), 10)
			result.Write(buf)
		}
		result.WriteRune('}')
	}
	if s.FloatVal != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		buf = buf[:0]
		result.WriteString(`"floatVal":{`)
		var _filled bool
		for _k, _v := range s.FloatVal {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(string(_k))
			result.WriteString(`":`)
			buf = strconv.AppendFloat(buf[:0], float64(_v), 'f', -1, 64)
			result.Write(buf)
		}
		result.WriteRune('}')
	}
	if s.UintVal != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		buf = buf[:0]
		result.WriteString(`"uintVal":{`)
		var _filled bool
		for _k, _v := range s.UintVal {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(string(_k))
			result.WriteString(`":`)
			if _v == nil {
				buf = append(buf[:0], 'n', 'u', 'l', 'l')
			} else {
				buf = strconv.AppendUint(buf[:0], uint64(*_v), 10)
			}
			result.Write(buf)
		}
		result.WriteRune('}')
	}
	if s.BoolVal != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		buf = buf[:0]
		result.WriteString(`"bool":{`)
		var _filled bool
		for _k, _v := range s.BoolVal {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(string(_k))
			result.WriteString(`":`)
			if _v {
				result.WriteString("true")
			} else {
				result.WriteString("false")
			}
			result.Write(buf)
		}
		result.WriteRune('}')
	}
	if s.TypedVal != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		buf = buf[:0]
		result.WriteString(`"typed-val":{`)
		var _filled bool
		for _k, _v := range s.TypedVal {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(string(_k))
			result.WriteString(`":`)
			buf = strconv.AppendUint(buf[:0], uint64(_v), 10)
			result.Write(buf)
		}
		result.WriteRune('}')
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Property) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *Property) MarshalAppend(dst []byte) ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	var (
		err    error
		buf    = make([]byte, 0, 128)
		result = bytes.NewBuffer(dst)
	)
	result.WriteRune('{')
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Name != "" {
		result.WriteString(`"name":`)
		buf = marshalString(buf[:0], s.Name)
		result.Write(buf)
	} else {
		result.WriteString(`"name":""`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Value != "" {
		result.WriteString(`"value":`)
		buf = marshalString(buf[:0], s.Value)
		result.Write(buf)
	} else {
		result.WriteString(`"value":""`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}
