// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_slice

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/fastjson"
)

// jsonParserTestSlice01 used for pooling Parsers for TestSlice01 JSONs.
var jsonParserTestSlice01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice01.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestSlice01) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _field := v.Get("strs"); valueIsNotNull(_field) {
		var listA []*fastjson.Value
		listA, err = _field.Array()
		if err != nil {
			return newParsingError("strs", err)
		}
		valField := s.Field[:0]
		if l := len(listA); cap(valField) < l || (l == 0 && s.Field == nil) {
			valField = make([]string, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem []byte
			if elem, err = listElem.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_elemNum), err)
			}
			valField = append(valField, string(elem))
		}
		if err != nil {
			return newParsingError("strs", err)
		}
		s.Field = valField
	}
	if _fieldRef := v.Get("ints"); valueIsNotNull(_fieldRef) {
		var listA []*fastjson.Value
		listA, err = _fieldRef.Array()
		if err != nil {
			return newParsingError("ints", err)
		}
		valFieldRef := s.FieldRef[:0]
		if l := len(listA); cap(valFieldRef) < l || (l == 0 && s.FieldRef == nil) {
			valFieldRef = make([]*int, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			if !valueIsNotNull(listElem) {
				valFieldRef = append(valFieldRef, nil)
				continue
			}
			var elem int
			elem, err = listElem.Int()
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			newElem := int(elem)
			valFieldRef = append(valFieldRef, &newElem)
		}
		if err != nil {
			return newParsingError("ints", err)
		}
		s.FieldRef = valFieldRef
	}
	return nil
}

// validate checks for correct data structure
func (s *TestSlice01) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'s', 't', 'r', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserTestSlice03 used for pooling Parsers for TestSlice03 JSONs.
var jsonParserTestSlice03 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice03) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice03.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice03.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestSlice03) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _data := v.Get("data"); _data != nil {
		var valData int64
		valData, err = _data.Int64()
		if err != nil {
			return newParsingError("data", err)
		}
		s.Data = valData
	}
	return nil
}

// validate checks for correct data structure
func (s *TestSlice03) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'a'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserTestSlice02 used for pooling Parsers for TestSlice02 JSONs.
var jsonParserTestSlice02 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice02) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice02.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice02.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *TestSlice02) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	a, err := v.Array()
	if err != nil {
		return err
	}
	*s = make([]TestSlice03, len(a))
	for i, v := range a {
		var value TestSlice03
		err = value.FillFromJSON(v)
		if err != nil {
			return newParsingError(fmt.Sprintf("%d", i), err)
		}
		(*s)[i] = TestSlice03(value)
	}
	return nil
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
		var listA []*fastjson.Value
		listA, err = _excluded.Array()
		if err != nil {
			return newParsingError("excluded", err)
		}
		valExcluded := s.Excluded[:0]
		if l := len(listA); cap(valExcluded) < l || (l == 0 && s.Excluded == nil) {
			valExcluded = make([]FieldValueString, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem []byte
			if elem, err = listElem.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_elemNum), err)
			}
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valExcluded = append(valExcluded, FieldValueString(elem))
		}
		if err != nil {
			return newParsingError("excluded", err)
		}
		s.Excluded = valExcluded
	}
	if _included := v.Get("included"); valueIsNotNull(_included) {
		var listA []*fastjson.Value
		listA, err = _included.Array()
		if err != nil {
			return newParsingError("included", err)
		}
		var valIncluded [5]FieldValueString
		if len(listA) != 5 {
			return newParsingError("included", fmt.Errorf("array len mismatch"))
		}
		for _elemNum, listElem := range listA {
			var elem []byte
			if elem, err = listElem.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_elemNum), err)
			}
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valIncluded[_elemNum] = FieldValueString(elem)
		}
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
func (s *TestSlice01) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice01) MarshalTo(result *jwriter.Writer) error {
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
	if s.Field != nil {
		wantComma = true
		result.RawString(`"strs":[`)
		var wantComma bool
		for _k, _v := range s.Field {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			result.String(_v)
		}
		result.RawByte(']')
		wantComma = true
	} else {
		result.RawString(`"strs":null`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.FieldRef != nil {
		wantComma = true
		result.RawString(`"ints":[`)
		var wantComma bool
		for _k, _v := range s.FieldRef {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			if _v == nil {
				result.RawString("null")
			} else {
				result.Int64(int64(*_v))
			}
		}
		result.RawByte(']')
		wantComma = true
	} else {
		result.RawString(`"ints":null`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestSlice01) IsZero() bool {
	if s.Field != nil {
		return false
	}
	if s.FieldRef != nil {
		return false
	}
	return true
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSlice03) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice03) MarshalTo(result *jwriter.Writer) error {
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
	if s.Data != 0 {
		result.RawString(`"data":`)
		result.Int64(s.Data)
		wantComma = true
	} else {
		result.RawString(`"data":0`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestSlice03) IsZero() bool {
	if s.Data != 0 {
		return false
	}
	return true
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSlice02) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice02) MarshalTo(result *jwriter.Writer) error {
	if s == nil || *s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('[')
	for _k, _v := range *s {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		_k = _k
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "TestSlice02" value at position %d: %w`, _k, err)
		}
	}
	result.RawByte(']')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestSlice02) IsZero() bool {
	return len(s) == 0
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
		result.RawString(`"excluded":[`)
		var wantComma bool
		for _k, _v := range s.Excluded {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			result.String(string(_v))
		}
		result.RawByte(']')
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	wantComma = true
	result.RawString(`"included":[`)
	wantComma = false
	for _k, _v := range s.Included {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		_k = _k
		result.String(string(_v))
	}
	result.RawByte(']')
	wantComma = true
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s CampaignSites) IsZero() bool {
	return false
}
