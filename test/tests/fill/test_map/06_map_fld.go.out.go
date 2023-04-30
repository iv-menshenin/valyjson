// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_map

import (
	"bytes"
	"fmt"
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestMap01) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _tags := v.Get("tags"); valueIsNotNull(_tags) {
		o, err := _tags.Object()
		if err != nil {
			return newParsingError("tags", err)
		}
		var valTags = make(map[string]string, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value []byte
			value, err = v.StringBytes()
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valTags[string(key)] = string(value)
			}
		})
		if err != nil {
			return newParsingError("tags", err)
		}
		s.Tags = valTags
	}
	if _properties := v.Get("properties"); valueIsNotNull(_properties) {
		o, err := _properties.Object()
		if err != nil {
			return newParsingError("properties", err)
		}
		var valProperties = make(map[string]Property, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value Property
			err = value.FillFromJSON(v)
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valProperties[string(key)] = Property(value)
			}
		})
		if err != nil {
			return newParsingError("properties", err)
		}
		s.Properties = valProperties
	}
	if _keyTypedProperties := v.Get("key_typed_properties"); valueIsNotNull(_keyTypedProperties) {
		o, err := _keyTypedProperties.Object()
		if err != nil {
			return newParsingError("key_typed_properties", err)
		}
		var valKeyTypedProperties = make(map[Key]Property, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value Property
			err = value.FillFromJSON(v)
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valKeyTypedProperties[Key(key)] = Property(value)
			}
		})
		if err != nil {
			return newParsingError("key_typed_properties", err)
		}
		s.KeyTypedProperties = valKeyTypedProperties
	}
	if _integerVal := v.Get("integerVal"); valueIsNotNull(_integerVal) {
		o, err := _integerVal.Object()
		if err != nil {
			return newParsingError("integerVal", err)
		}
		var valIntegerVal = make(map[Key]int32, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value int
			value, err = v.Int()
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valIntegerVal[Key(key)] = int32(value)
			}
		})
		if err != nil {
			return newParsingError("integerVal", err)
		}
		s.IntegerVal = valIntegerVal
	}
	if _floatVal := v.Get("floatVal"); valueIsNotNull(_floatVal) {
		o, err := _floatVal.Object()
		if err != nil {
			return newParsingError("floatVal", err)
		}
		var valFloatVal = make(map[Key]float64, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value float64
			value, err = v.Float64()
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valFloatVal[Key(key)] = float64(value)
			}
		})
		if err != nil {
			return newParsingError("floatVal", err)
		}
		s.FloatVal = valFloatVal
	}
	if _uintVal := v.Get("uintVal"); valueIsNotNull(_uintVal) {
		o, err := _uintVal.Object()
		if err != nil {
			return newParsingError("uintVal", err)
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
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valUintVal[Key(key)] = (*uint16)(unsafe.Pointer(&value))
			}
		})
		if err != nil {
			return newParsingError("uintVal", err)
		}
		s.UintVal = valUintVal
	}
	if _boolVal := v.Get("bool"); valueIsNotNull(_boolVal) {
		o, err := _boolVal.Object()
		if err != nil {
			return newParsingError("bool", err)
		}
		var valBoolVal = make(map[Key]bool, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value bool
			value, err = v.Bool()
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valBoolVal[Key(key)] = bool(value)
			}
		})
		if err != nil {
			return newParsingError("bool", err)
		}
		s.BoolVal = valBoolVal
	}
	if _typedVal := v.Get("typed-val"); valueIsNotNull(_typedVal) {
		o, err := _typedVal.Object()
		if err != nil {
			return newParsingError("typed-val", err)
		}
		var valTypedVal = make(map[Key]Val, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value uint64
			value, err = v.Uint64()
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valTypedVal[Key(key)] = Val(value)
			}
		})
		if err != nil {
			return newParsingError("typed-val", err)
		}
		s.TypedVal = valTypedVal
	}
	return nil
}

// validate checks for correct data structure
func (s *TestMap01) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'p', 'e', 'r', 't', 'i', 'e', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'k', 'e', 'y', '_', 't', 'y', 'p', 'e', 'd', '_', 'p', 'r', 'o', 'p', 'e', 'r', 't', 'i', 'e', 's'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', 'e', 'g', 'e', 'r', 'V', 'a', 'l'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'l', 'o', 'a', 't', 'V', 'a', 'l'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 'i', 'n', 't', 'V', 'a', 'l'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'b', 'o', 'o', 'l'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'y', 'p', 'e', 'd', '-', 'v', 'a', 'l'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Property) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _name := v.Get("name"); _name != nil {
		var valName []byte
		if valName, err = _name.StringBytes(); err != nil {
			return newParsingError("name", err)
		}
		s.Name = *(*string)(unsafe.Pointer(&valName))
	}
	if _value := v.Get("value"); _value != nil {
		var valValue []byte
		if valValue, err = _value.StringBytes(); err != nil {
			return newParsingError("value", err)
		}
		s.Value = *(*string)(unsafe.Pointer(&valValue))
	}
	return nil
}

// validate checks for correct data structure
func (s *Property) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'v', 'a', 'l', 'u', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
	})
	return err
}

var bufDataTestMap01 = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap01) MarshalJSON() ([]byte, error) {
	var result = bufDataTestMap01.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestMap01) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if s.Tags != nil {
		wantComma = true
		result.WriteString(`"tags":{`)
		var wantComma bool
		for _k, _v := range s.Tags {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(_k)
			result.WriteString(`":`)
			writeString(result, _v)
		}
		result.WriteString("}")
	} else {
		wantComma = true
		result.WriteString(`"tags":null`)
	}
	if s.Properties != nil {
		result.WriteString(",")
		wantComma = true
		result.WriteString(`"properties":{`)
		var wantComma bool
		for _k, _v := range s.Properties {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(_k)
			result.WriteString(`":`)
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "properties" attribute %q: %w`, _k, err)
			}
		}
		result.WriteString("}")
	}
	result.WriteString(",")
	if s.KeyTypedProperties != nil {
		wantComma = true
		result.WriteString(`"key_typed_properties":{`)
		var wantComma bool
		for _k, _v := range s.KeyTypedProperties {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(string(_k))
			result.WriteString(`":`)
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "key_typed_properties" attribute %q: %w`, _k, err)
			}
		}
		result.WriteString("}")
	} else {
		wantComma = true
		result.WriteString(`"key_typed_properties":null`)
	}
	if s.IntegerVal != nil {
		result.WriteString(",")
		wantComma = true
		result.WriteString(`"integerVal":{`)
		var wantComma bool
		for _k, _v := range s.IntegerVal {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(string(_k))
			result.WriteString(`":`)
			writeInt64(result, int64(_v))
		}
		result.WriteString("}")
	}
	if s.FloatVal != nil {
		result.WriteString(",")
		wantComma = true
		result.WriteString(`"floatVal":{`)
		var wantComma bool
		for _k, _v := range s.FloatVal {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(string(_k))
			result.WriteString(`":`)
			writeFloat64(result, float64(_v))
		}
		result.WriteString("}")
	}
	if s.UintVal != nil {
		result.WriteString(",")
		wantComma = true
		result.WriteString(`"uintVal":{`)
		var wantComma bool
		for _k, _v := range s.UintVal {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(string(_k))
			result.WriteString(`":`)
			if _v == nil {
				result.WriteString("null")
			} else {
				writeUint64(result, uint64(*_v))
			}
		}
		result.WriteString("}")
	}
	if s.BoolVal != nil {
		result.WriteString(",")
		wantComma = true
		result.WriteString(`"bool":{`)
		var wantComma bool
		for _k, _v := range s.BoolVal {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(string(_k))
			result.WriteString(`":`)
			if _v {
				result.WriteString("true")
			} else {
				result.WriteString("false")
			}
		}
		result.WriteString("}")
	}
	if s.TypedVal != nil {
		result.WriteString(",")
		wantComma = true
		result.WriteString(`"typed-val":{`)
		var wantComma bool
		for _k, _v := range s.TypedVal {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(string(_k))
			result.WriteString(`":`)
			writeUint64(result, uint64(_v))
		}
		result.WriteString("}")
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestMap01) IsZero() bool {
	if s.Tags != nil {
		return false
	}
	if s.Properties != nil {
		return false
	}
	if s.KeyTypedProperties != nil {
		return false
	}
	if s.IntegerVal != nil {
		return false
	}
	if s.FloatVal != nil {
		return false
	}
	if s.UintVal != nil {
		return false
	}
	if s.BoolVal != nil {
		return false
	}
	if s.TypedVal != nil {
		return false
	}
	return true
}

var bufDataProperty = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Property) MarshalJSON() ([]byte, error) {
	var result = bufDataProperty.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Property) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if s.Name != "" {
		result.WriteString(`"name":`)
		writeString(result, s.Name)
		wantComma = true
	} else {
		result.WriteString(`"name":""`)
		wantComma = true
	}
	result.WriteString(",")
	if s.Value != "" {
		result.WriteString(`"value":`)
		writeString(result, s.Value)
		wantComma = true
	} else {
		result.WriteString(`"value":""`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Property) IsZero() bool {
	if s.Name != "" {
		return false
	}
	if s.Value != "" {
		return false
	}
	return true
}
