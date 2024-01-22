// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_map

import (
	"bytes"
	"fmt"
	"unsafe"

	"github.com/mailru/easyjson/jwriter"
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
		s.Name = string(valName)
	}
	if _value := v.Get("value"); _value != nil {
		var valValue []byte
		if valValue, err = _value.StringBytes(); err != nil {
			return newParsingError("value", err)
		}
		s.Value = string(valValue)
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

// jsonParserCampaignSites used for pooling Parsers for CampaignSites JSONs.
var jsonParserCampaignSites fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *CampaignSites) UnmarshalJSON(data []byte) error {
	parser := jsonParserCampaignSites.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserCampaignSites.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *CampaignSites) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _excluded := v.Get("excluded"); valueIsNotNull(_excluded) {
		o, err := _excluded.Object()
		if err != nil {
			return newParsingError("excluded", err)
		}
		var valExcluded = make(map[string]FieldValueString, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value []byte
			value, err = v.StringBytes()
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valExcluded[string(key)] = FieldValueString(value)
			}
		})
		if err != nil {
			return newParsingError("excluded", err)
		}
		s.Excluded = valExcluded
	}
	if _included := v.Get("included"); valueIsNotNull(_included) {
		o, err := _included.Object()
		if err != nil {
			return newParsingError("included", err)
		}
		var valIncluded = make(map[FieldValueString]string, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value []byte
			value, err = v.StringBytes()
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valIncluded[FieldValueString(key)] = string(value)
			}
		})
		if err != nil {
			return newParsingError("included", err)
		}
		s.Included = valIncluded
	}
	return nil
}

// validate checks for correct data structure
func (s *CampaignSites) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'e', 'x', 'c', 'l', 'u', 'd', 'e', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 'c', 'l', 'u', 'd', 'e', 'd'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap01) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestMap01) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.Tags != nil {
		wantComma = true
		result.RawString(`"tags":{`)
		var wantComma bool
		for _k, _v := range s.Tags {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(_k)
			result.RawByte(':')
			result.String(_v)
		}
		result.RawByte('}')
	} else {
		wantComma = true
		result.RawString(`"tags":null`)
	}
	if s.Properties != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"properties":{`)
		var wantComma bool
		for _k, _v := range s.Properties {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(_k)
			result.RawByte(':')
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "properties" attribute %q: %w`, _k, err)
			}
		}
		result.RawByte('}')
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.KeyTypedProperties != nil {
		wantComma = true
		result.RawString(`"key_typed_properties":{`)
		var wantComma bool
		for _k, _v := range s.KeyTypedProperties {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(string(_k))
			result.RawByte(':')
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "key_typed_properties" attribute %q: %w`, _k, err)
			}
		}
		result.RawByte('}')
	} else {
		wantComma = true
		result.RawString(`"key_typed_properties":null`)
	}
	if s.IntegerVal != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"integerVal":{`)
		var wantComma bool
		for _k, _v := range s.IntegerVal {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(string(_k))
			result.RawByte(':')
			writeInt64(result, int64(_v))
		}
		result.RawByte('}')
	}
	if s.FloatVal != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"floatVal":{`)
		var wantComma bool
		for _k, _v := range s.FloatVal {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(string(_k))
			result.RawByte(':')
			writeFloat64(result, float64(_v))
		}
		result.RawByte('}')
	}
	if s.UintVal != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"uintVal":{`)
		var wantComma bool
		for _k, _v := range s.UintVal {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(string(_k))
			result.RawByte(':')
			if _v == nil {
				result.RawString("null")
			} else {
				writeUint64(result, uint64(*_v))
			}
		}
		result.RawByte('}')
	}
	if s.BoolVal != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"bool":{`)
		var wantComma bool
		for _k, _v := range s.BoolVal {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(string(_k))
			result.RawByte(':')
			if _v {
				result.RawString("true")
			} else {
				result.RawString("false")
			}
		}
		result.RawByte('}')
	}
	if s.TypedVal != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"typed-val":{`)
		var wantComma bool
		for _k, _v := range s.TypedVal {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(string(_k))
			result.RawByte(':')
			writeUint64(result, uint64(_v))
		}
		result.RawByte('}')
	}
	result.RawByte('}')
	err = result.Error
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Property) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Property) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.Name != "" {
		result.RawString(`"name":`)
		result.String(s.Name)
		wantComma = true
	} else {
		result.RawString(`"name":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Value != "" {
		result.RawString(`"value":`)
		result.String(s.Value)
		wantComma = true
	} else {
		result.RawString(`"value":""`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *CampaignSites) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *CampaignSites) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if s.Excluded != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"excluded":{`)
		var wantComma bool
		for _k, _v := range s.Excluded {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(_k)
			result.RawByte(':')
			result.String(string(_v))
		}
		result.RawByte('}')
	}
	if s.Included != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"included":{`)
		var wantComma bool
		for _k, _v := range s.Included {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(string(_k))
			result.RawByte(':')
			result.String(_v)
		}
		result.RawByte('}')
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s CampaignSites) IsZero() bool {
	if s.Excluded != nil {
		return false
	}
	if s.Included != nil {
		return false
	}
	return true
}
