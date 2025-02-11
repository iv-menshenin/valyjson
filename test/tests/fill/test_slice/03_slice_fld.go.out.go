// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_slice

import (
	"bytes"
	"encoding/base64"
	"errors"
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
		valField := s.Field
		if l := len(listA); cap(valField) < l || (l == 0 && s.Field == nil) {
			valField = make([]string, 0, len(listA))
		} else {
			valField = s.Field[:0]
		}
		for _key1, _val1 := range listA {
			valField = valField[:len(valField)+1]
			var _elem1 []byte
			if _elem1, err = _val1.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_key1), err)
			}
			valField[_key1] = string(_elem1)
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
		valFieldRef := s.FieldRef
		if l := len(listA); cap(valFieldRef) < l || (l == 0 && s.FieldRef == nil) {
			valFieldRef = make([]*int, 0, len(listA))
		} else {
			valFieldRef = s.FieldRef[:0]
		}
		for _key1, _val1 := range listA {
			valFieldRef = valFieldRef[:len(valFieldRef)+1]
			if !valueIsNotNull(_val1) {
				valFieldRef[len(valFieldRef)-1] = nil
				continue
			}
			var _elem1 int
			_elem1, err = _val1.Int()
			if err != nil {
				err = newParsingError(strconv.Itoa(_key1), err)
				break
			}
			newElem := int(_elem1)
			valFieldRef[_key1] = &newElem
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
		valExcluded := s.Excluded
		if l := len(listA); cap(valExcluded) < l || (l == 0 && s.Excluded == nil) {
			valExcluded = make([]FieldValueString, 0, len(listA))
		} else {
			valExcluded = s.Excluded[:0]
		}
		for _key1, _val1 := range listA {
			valExcluded = valExcluded[:len(valExcluded)+1]
			var _elem1 []byte
			if _elem1, err = _val1.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_key1), err)
			}
			if err != nil {
				err = newParsingError(strconv.Itoa(_key1), err)
				break
			}
			valExcluded[_key1] = FieldValueString(_elem1)
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
		for _key1, _val1 := range listA {
			var _tmp1 []byte
			if _tmp1, err = _val1.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_key1), err)
			}
			if err != nil {
				err = newParsingError(strconv.Itoa(_key1), err)
				break
			}
			valIncluded[_key1] = FieldValueString(_tmp1)
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

// jsonParserTestSliceSlice used for pooling Parsers for TestSliceSlice JSONs.
var jsonParserTestSliceSlice fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSliceSlice) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSliceSlice.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSliceSlice.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestSliceSlice) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _fieldStr := v.Get("strs"); valueIsNotNull(_fieldStr) {
		var listA []*fastjson.Value
		listA, err = _fieldStr.Array()
		if err != nil {
			return newParsingError("strs", err)
		}
		valFieldStr := s.FieldStr
		if l := len(listA); cap(valFieldStr) < l || (l == 0 && s.FieldStr == nil) {
			valFieldStr = make([][]InnerString, 0, len(listA))
		} else {
			valFieldStr = s.FieldStr[:0]
		}
		for _key1, _val1 := range listA {
			valFieldStr = valFieldStr[:len(valFieldStr)+1]
			if !valueIsNotNull(_val1) {
				valFieldStr[len(valFieldStr)-1] = nil
				continue
			}
			var listA []*fastjson.Value
			listA, err = _val1.Array()
			if err != nil {
				return newParsingError("", err)
			}
			_elem1 := valFieldStr[len(valFieldStr)-1]
			if l := len(listA); cap(_elem1) < l || (l == 0 && valFieldStr[len(valFieldStr)-1] == nil) {
				_elem1 = make([]InnerString, 0, len(listA))
			} else {
				_elem1 = valFieldStr[len(valFieldStr)-1][:0]
			}
			for _key2, _val2 := range listA {
				_elem1 = _elem1[:len(_elem1)+1]
				var _elem2 []byte
				if _elem2, err = _val2.StringBytes(); err != nil {
					return newParsingError(strconv.Itoa(_key2), err)
				}
				if err != nil {
					err = newParsingError(strconv.Itoa(_key2), err)
					break
				}
				_elem1[_key2] = InnerString(_elem2)
			}
			if err != nil {
				err = newParsingError(strconv.Itoa(_key1), err)
				break
			}
			valFieldStr[_key1] = []InnerString(_elem1)
		}
		if err != nil {
			return newParsingError("strs", err)
		}
		s.FieldStr = valFieldStr
	}
	if _fieldInt := v.Get("ints"); valueIsNotNull(_fieldInt) {
		var listA []*fastjson.Value
		listA, err = _fieldInt.Array()
		if err != nil {
			return newParsingError("ints", err)
		}
		valFieldInt := s.FieldInt
		if l := len(listA); cap(valFieldInt) < l || (l == 0 && s.FieldInt == nil) {
			valFieldInt = make([][]int, 0, len(listA))
		} else {
			valFieldInt = s.FieldInt[:0]
		}
		for _key1, _val1 := range listA {
			valFieldInt = valFieldInt[:len(valFieldInt)+1]
			if !valueIsNotNull(_val1) {
				valFieldInt[len(valFieldInt)-1] = nil
				continue
			}
			var listA []*fastjson.Value
			listA, err = _val1.Array()
			if err != nil {
				return newParsingError("", err)
			}
			_elem1 := valFieldInt[len(valFieldInt)-1]
			if l := len(listA); cap(_elem1) < l || (l == 0 && valFieldInt[len(valFieldInt)-1] == nil) {
				_elem1 = make([]int, 0, len(listA))
			} else {
				_elem1 = valFieldInt[len(valFieldInt)-1][:0]
			}
			for _key2, _val2 := range listA {
				_elem1 = _elem1[:len(_elem1)+1]
				var _elem2 int
				_elem2, err = _val2.Int()
				if err != nil {
					err = newParsingError(strconv.Itoa(_key2), err)
					break
				}
				_elem1[_key2] = int(_elem2)
			}
			if err != nil {
				err = newParsingError(strconv.Itoa(_key1), err)
				break
			}
			valFieldInt[_key1] = []int(_elem1)
		}
		if err != nil {
			return newParsingError("ints", err)
		}
		s.FieldInt = valFieldInt
	}
	return nil
}

// validate checks for correct data structure
func (s *TestSliceSlice) validate(v *fastjson.Value) error {
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

// jsonParserBytes used for pooling Parsers for Bytes JSONs.
var jsonParserBytes fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Bytes) UnmarshalJSON(data []byte) error {
	parser := jsonParserBytes.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserBytes.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the byteslice with the values recognized from fastjson.Value
func (s *Bytes) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() != fastjson.TypeNull {
		// slice of bytes in JSON format is implemented via BASE64 string
		b, err := v.StringBytes()
		if err != nil {
			return err
		}
		*s = make([]byte, base64.StdEncoding.DecodedLen(len(b)))
		n, err := base64.StdEncoding.Decode(*s, b)
		if err != nil {
			return err
		}
		*s = (*s)[:n]
	}
	return
}

// jsonParserArrBytes used for pooling Parsers for ArrBytes JSONs.
var jsonParserArrBytes fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *ArrBytes) UnmarshalJSON(data []byte) error {
	parser := jsonParserArrBytes.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserArrBytes.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the byteslice with the values recognized from fastjson.Value
func (s *ArrBytes) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() != fastjson.TypeNull {
		// slice of bytes in JSON format is implemented via BASE64 string
		b, err := v.StringBytes()
		if err != nil {
			return err
		}
		n, err := base64.StdEncoding.Decode((*s)[:], b)
		if err != nil {
			return err
		}
		if n != len(*s) {
			return errors.New("incomplete data")
		}
	}
	return
}

// jsonParserBytesInStruct used for pooling Parsers for BytesInStruct JSONs.
var jsonParserBytesInStruct fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *BytesInStruct) UnmarshalJSON(data []byte) error {
	parser := jsonParserBytesInStruct.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserBytesInStruct.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *BytesInStruct) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _line := v.Get("line"); _line != nil {
		var valLine []byte
		if valLine, err = _line.StringBytes(); err != nil {
			return newParsingError("line", err)
		}
		s.Line = string(valLine)
	}
	if _ifBinary := v.Get("if-binary"); _ifBinary != nil {
		var valIfBinary []byte
		if _ifBinary.Type() != fastjson.TypeNull {
			// slice of bytes in JSON format is implemented via BASE64 string
			b, err := _ifBinary.StringBytes()
			if err != nil {
				return err
			}
			valIfBinary = make([]byte, base64.StdEncoding.DecodedLen(len(b)))
			n, err := base64.StdEncoding.Decode(valIfBinary, b)
			if err != nil {
				return err
			}
			valIfBinary = valIfBinary[:n]
		}
		if err != nil {
			return newParsingError("if-binary", err)
		}
		s.IfBinary = Bytes(valIfBinary)
	}
	if _binary := v.Get("binary"); _binary != nil {
		var valBinary []byte
		if _binary.Type() != fastjson.TypeNull {
			// slice of bytes in JSON format is implemented via BASE64 string
			b, err := _binary.StringBytes()
			if err != nil {
				return err
			}
			valBinary = make([]byte, base64.StdEncoding.DecodedLen(len(b)))
			n, err := base64.StdEncoding.Decode(valBinary, b)
			if err != nil {
				return err
			}
			valBinary = valBinary[:n]
		}
		if err != nil {
			return newParsingError("binary", err)
		}
		s.Binary = Bytes(valBinary)
	}
	if _arrBinary := v.Get("arrbinary"); _arrBinary != nil {
		var valArrBinary [8]byte
		if _arrBinary.Type() != fastjson.TypeNull {
			// slice of bytes in JSON format is implemented via BASE64 string
			b, err := _arrBinary.StringBytes()
			if err != nil {
				return err
			}
			n, err := base64.StdEncoding.Decode(valArrBinary[:], b)
			if err != nil {
				return err
			}
			if n != len(valArrBinary) {
				return errors.New("incomplete data")
			}
		}
		if err != nil {
			return newParsingError("arrbinary", err)
		}
		s.ArrBinary = ArrBytes(valArrBinary)
	}
	if _ifArrBinary := v.Get("if-arrbinary"); _ifArrBinary != nil {
		var valIfArrBinary [8]byte
		if _ifArrBinary.Type() != fastjson.TypeNull {
			// slice of bytes in JSON format is implemented via BASE64 string
			b, err := _ifArrBinary.StringBytes()
			if err != nil {
				return err
			}
			n, err := base64.StdEncoding.Decode(valIfArrBinary[:], b)
			if err != nil {
				return err
			}
			if n != len(valIfArrBinary) {
				return errors.New("incomplete data")
			}
		}
		if err != nil {
			return newParsingError("if-arrbinary", err)
		}
		s.IfArrBinary = ArrBytes(valIfArrBinary)
	}
	return nil
}

// validate checks for correct data structure
func (s *BytesInStruct) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [5]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'l', 'i', 'n', 'e'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'f', '-', 'b', 'i', 'n', 'a', 'r', 'y'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'b', 'i', 'n', 'a', 'r', 'y'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'a', 'r', 'r', 'b', 'i', 'n', 'a', 'r', 'y'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'f', '-', 'a', 'r', 'r', 'b', 'i', 'n', 'a', 'r', 'y'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestSlice01) Reset() {
	for i := range s.Field {
		s.Field[i] = ""
	}
	s.Field = s.Field[:0]
	for i := range s.FieldRef {
		s.FieldRef[i] = nil
	}
	s.FieldRef = s.FieldRef[:0]
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestSlice03) Reset() {
	s.Data = 0
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestSlice02) Reset() {
	for i := range *s {
		(*s)[i].Reset()
	}
	*s = (*s)[:0]
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *CampaignSites) Reset() {
	for i := range s.Excluded {
		s.Excluded[i] = FieldValueString("")
	}
	s.Excluded = s.Excluded[:0]
	s.Included = [5]FieldValueString{}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSliceSlice) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSliceSlice) MarshalTo(result *jwriter.Writer) error {
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
	if s.FieldStr != nil {
		wantComma = true
		result.RawString(`"strs":[`)
		var wantComma bool
		for _k, _v := range s.FieldStr {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			result.RawByte('[')
			for _k1, _v1 := range _v {
				result.String(string(_v1))
				if len(_v)-1 > _k1 {
					result.RawByte(',')
				}
			}
			result.RawByte(']')
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
	if s.FieldInt != nil {
		wantComma = true
		result.RawString(`"ints":[`)
		var wantComma bool
		for _k, _v := range s.FieldInt {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			result.RawByte('[')
			for _k1, _v1 := range _v {
				result.Int64(int64(_v1))
				if len(_v)-1 > _k1 {
					result.RawByte(',')
				}
			}
			result.RawByte(']')
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
func (s TestSliceSlice) IsZero() bool {
	if s.FieldStr != nil {
		return false
	}
	if s.FieldInt != nil {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestSliceSlice) Reset() {
	for i := range s.FieldStr {
		for j := range s.FieldStr[i] {
			s.FieldStr[i][j] = InnerString("")
		}
		s.FieldStr[i] = s.FieldStr[i][:0]
	}
	s.FieldStr = s.FieldStr[:0]
	for i := range s.FieldInt {
		for j := range s.FieldInt[i] {
			s.FieldInt[i][j] = 0
		}
		s.FieldInt[i] = s.FieldInt[i][:0]
	}
	s.FieldInt = s.FieldInt[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Bytes) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Bytes) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	if *s == nil {
		result.RawString("null")
		return nil
	}
	// slice of bytes in JSON format is implemented via BASE64 string
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(*s)))
	base64.StdEncoding.Encode(buf, *s)
	result.String(string(buf))
	return nil
}

// IsZero shows whether the object is an empty value.
func (s Bytes) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Bytes) Reset() {
	*s = (*s)[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *ArrBytes) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *ArrBytes) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	// slice of bytes in JSON format is implemented via BASE64 string
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(*s)))
	base64.StdEncoding.Encode(buf, (*s)[:])
	result.String(string(buf))
	return nil
}

// IsZero shows whether the object is an empty value.
func (s ArrBytes) IsZero() bool {
	for _, _v := range s {
		if _v != 0 {
			return false
		}
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *ArrBytes) Reset() {
	*s = ArrBytes{}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *BytesInStruct) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *BytesInStruct) MarshalTo(result *jwriter.Writer) error {
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
	if s.Line != "" {
		result.RawString(`"line":`)
		result.String(s.Line)
		wantComma = true
	} else {
		result.RawString(`"line":""`)
		wantComma = true
	}
	if s.IfBinary != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"if-binary":`)
		// slice of bytes in JSON format is implemented via BASE64 string
		buf := make([]byte, base64.StdEncoding.EncodedLen(len(s.IfBinary)))
		base64.StdEncoding.Encode(buf, s.IfBinary)
		result.String(string(buf))
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Binary != nil {
		wantComma = true
		result.RawString(`"binary":`)
		// slice of bytes in JSON format is implemented via BASE64 string
		buf := make([]byte, base64.StdEncoding.EncodedLen(len(s.Binary)))
		base64.StdEncoding.Encode(buf, s.Binary)
		result.String(string(buf))
	} else {
		result.RawString(`"binary":null`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.ArrBinary != [8]byte{} {
		wantComma = true
		result.RawString(`"arrbinary":`)
		// slice of bytes in JSON format is implemented via BASE64 string
		buf := make([]byte, base64.StdEncoding.EncodedLen(len(s.ArrBinary[:])))
		base64.StdEncoding.Encode(buf, s.ArrBinary[:])
		result.String(string(buf))
	} else {
		result.RawString(`"arrbinary":null`)
		wantComma = true
	}
	if s.IfArrBinary != [8]byte{} {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"if-arrbinary":`)
		// slice of bytes in JSON format is implemented via BASE64 string
		buf := make([]byte, base64.StdEncoding.EncodedLen(len(s.IfArrBinary[:])))
		base64.StdEncoding.Encode(buf, s.IfArrBinary[:])
		result.String(string(buf))
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s BytesInStruct) IsZero() bool {
	if s.Line != "" {
		return false
	}
	if s.IfBinary != nil {
		return false
	}
	if s.Binary != nil {
		return false
	}
	if !s.ArrBinary.IsZero() {
		return false
	}
	if !s.IfArrBinary.IsZero() {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *BytesInStruct) Reset() {
	s.Line = ""
	for i := range s.IfBinary {
		s.IfBinary[i] = 0
	}
	s.IfBinary = s.IfBinary[:0]
	for i := range s.Binary {
		s.Binary[i] = 0
	}
	s.Binary = s.Binary[:0]
	s.ArrBinary = [8]byte{}
	s.IfArrBinary = [8]byte{}
}
