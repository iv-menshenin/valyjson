// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_nested

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"
)

// jsonParserRoot used for pooling Parsers for Root JSONs.
var jsonParserRoot fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Root) UnmarshalJSON(data []byte) error {
	parser := jsonParserRoot.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserRoot.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Root) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _meta := v.Get("meta"); _meta != nil {
		var valMeta Meta
		err = valMeta.FillFromJSON(_meta)
		if err != nil {
			return newParsingError("meta", err)
		}
		s.Meta = Meta(valMeta)
	}
	if _data := v.Get("data"); _data != nil {
		var listA []*fastjson.Value
		listA, err = _data.Array()
		if err != nil {
			return newParsingError("data", err)
		}
		valData := s.Data[:0]
		if l := len(listA); cap(valData) < l || (l == 0 && s.Data == nil) {
			valData = make([]Middle, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem Middle
			err = elem.FillFromJSON(listElem)
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valData = append(valData, Middle(elem))
		}
		if err != nil {
			return newParsingError("data", err)
		}
		s.Data = Middles(valData)
	}
	return nil
}

// validate checks for correct data structure
func (s *Root) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'m', 'e', 't', 'a'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'a'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserMiddle used for pooling Parsers for Middle JSONs.
var jsonParserMiddle fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Middle) UnmarshalJSON(data []byte) error {
	parser := jsonParserMiddle.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserMiddle.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Middle) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _name := v.Get("name"); _name != nil {
		var valName []byte
		if valName, err = _name.StringBytes(); err != nil {
			return newParsingError("name", err)
		}
		s.Name = *(*UserName)(unsafe.Pointer(&valName))
	}
	if _surname := v.Get("surname"); _surname != nil {
		var valSurname []byte
		if valSurname, err = _surname.StringBytes(); err != nil {
			return newParsingError("surname", err)
		}
		s.Surname = *(*UserSurname)(unsafe.Pointer(&valSurname))
	}
	if _patname := v.Get("patname"); valueIsNotNull(_patname) {
		var valPatname []byte
		if valPatname, err = _patname.StringBytes(); err != nil {
			return newParsingError("patname", err)
		}
		s.Patname = new(UserPatname)
		*s.Patname = UserPatname(valPatname)
	}
	if _dateOfBorn := v.Get("dateOfBorn"); _dateOfBorn != nil {
		b, err := _dateOfBorn.StringBytes()
		if err != nil {
			return newParsingError("dateOfBorn", err)
		}
		valDateOfBorn, err := parseDateTime(string(b))
		if err != nil {
			return newParsingError("dateOfBorn", err)
		}
		s.DateOfBorn = valDateOfBorn
	}
	if _tags := v.Get("tags"); _tags != nil {
		o, err := _tags.Object()
		if err != nil {
			return newParsingError("tags", err)
		}
		var valTags = make(map[TagName]TagValue, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value []byte
			value, err = v.StringBytes()
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valTags[TagName(key)] = TagValue(value)
			}
		})
		if err != nil {
			return newParsingError("tags", err)
		}
		s.Tags = Tags(valTags)
	}
	return nil
}

// validate checks for correct data structure
func (s *Middle) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [5]int
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
		if bytes.Equal(key, []byte{'s', 'u', 'r', 'n', 'a', 'm', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'a', 't', 'n', 'a', 'm', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', 'O', 'f', 'B', 'o', 'r', 'n'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 's'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserMiddles used for pooling Parsers for Middles JSONs.
var jsonParserMiddles fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Middles) UnmarshalJSON(data []byte) error {
	parser := jsonParserMiddles.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserMiddles.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *Middles) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	a, err := v.Array()
	if err != nil {
		return err
	}
	*s = make([]Middle, len(a))
	for i, v := range a {
		var value Middle
		err = value.FillFromJSON(v)
		if err != nil {
			return newParsingError(fmt.Sprintf("%d", i), err)
		}
		(*s)[i] = Middle(value)
	}
	return nil
}

// jsonParserMeta used for pooling Parsers for Meta JSONs.
var jsonParserMeta fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Meta) UnmarshalJSON(data []byte) error {
	parser := jsonParserMeta.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserMeta.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Meta) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _count := v.Get("count"); _count != nil {
		var valCount int
		valCount, err = _count.Int()
		if err != nil {
			return newParsingError("count", err)
		}
		s.Count = valCount
	}
	return nil
}

// validate checks for correct data structure
func (s *Meta) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'u', 'n', 't'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserPersonal used for pooling Parsers for Personal JSONs.
var jsonParserPersonal fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Personal) UnmarshalJSON(data []byte) error {
	parser := jsonParserPersonal.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserPersonal.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Personal) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _name := v.Get("name"); _name != nil {
		var valName []byte
		if valName, err = _name.StringBytes(); err != nil {
			return newParsingError("name", err)
		}
		s.Name = *(*UserName)(unsafe.Pointer(&valName))
	}
	if _surname := v.Get("surname"); _surname != nil {
		var valSurname []byte
		if valSurname, err = _surname.StringBytes(); err != nil {
			return newParsingError("surname", err)
		}
		s.Surname = *(*UserSurname)(unsafe.Pointer(&valSurname))
	}
	if _patname := v.Get("patname"); valueIsNotNull(_patname) {
		var valPatname []byte
		if valPatname, err = _patname.StringBytes(); err != nil {
			return newParsingError("patname", err)
		}
		s.Patname = new(UserPatname)
		*s.Patname = UserPatname(valPatname)
	}
	return nil
}

// validate checks for correct data structure
func (s *Personal) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [3]int
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
		if bytes.Equal(key, []byte{'s', 'u', 'r', 'n', 'a', 'm', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'a', 't', 'n', 'a', 'm', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserTags used for pooling Parsers for Tags JSONs.
var jsonParserTags fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Tags) UnmarshalJSON(data []byte) error {
	parser := jsonParserTags.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTags.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *Tags) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	o, err := v.Object()
	if err != nil {
		return err
	}
	*s = make(map[TagName]TagValue, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value []byte
		value, err = v.StringBytes()
		if err != nil {
			err = newParsingError(string(key), err)
			return
		}
		(*s)[TagName(key)] = TagValue(value)
	})
	return err
}

var bufDataRoot = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Root) MarshalJSON() ([]byte, error) {
	var result = bufDataRoot.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Root) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if wantComma {
		result.WriteString(",")
	}
	result.WriteString(`"meta":`)
	if err = s.Meta.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "meta" attribute: %w`, err)
	}
	wantComma = true
	if wantComma {
		result.WriteString(",")
	}
	if s.Data != nil {
		wantComma = true
		result.WriteString(`"data":[`)
		var wantComma bool
		for _k, _v := range s.Data {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "data" item at position %d: %w`, _k, err)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"data":null`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Root) IsZero() bool {
	if !s.Meta.IsZero() {
		return false
	}
	if s.Data != nil {
		return false
	}
	return true
}

var bufDataMiddle = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Middle) MarshalJSON() ([]byte, error) {
	var result = bufDataMiddle.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Middle) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if wantComma {
		result.WriteString(",")
	}
	if s.Name != "" {
		result.WriteString(`"name":`)
		writeString(result, string(s.Name))
		wantComma = true
	} else {
		result.WriteString(`"name":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Surname != "" {
		result.WriteString(`"surname":`)
		writeString(result, string(s.Surname))
		wantComma = true
	} else {
		result.WriteString(`"surname":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Patname != nil {
		result.WriteString(`"patname":`)
		writeString(result, string(*s.Patname))
		wantComma = true
	} else {
		result.WriteString(`"patname":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if !s.DateOfBorn.IsZero() {
		result.WriteString(`"dateOfBorn":`)
		writeTime(result, s.DateOfBorn, time.RFC3339Nano)
		wantComma = true
	} else {
		result.WriteString(`"dateOfBorn":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
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
			result.WriteString(string(_k))
			result.WriteString(`":`)
			writeString(result, string(_v))
		}
		result.WriteString("}")
	} else {
		wantComma = true
		result.WriteString(`"tags":null`)
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Middle) IsZero() bool {
	if s.Name != "" {
		return false
	}
	if s.Surname != "" {
		return false
	}
	if s.Patname != nil {
		return false
	}
	if !s.DateOfBorn.IsZero() {
		return false
	}
	if s.Tags != nil {
		return false
	}
	return true
}

var bufDataMiddles = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Middles) MarshalJSON() ([]byte, error) {
	var result = bufDataMiddles.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Middles) MarshalTo(result Writer) error {
	if s == nil || *s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("[")
	for _k, _v := range *s {
		if wantComma {
			result.WriteString(",")
		}
		wantComma = true
		_k = _k
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "Middles" value at position %d: %w`, _k, err)
		}
	}
	result.WriteString("]")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Middles) IsZero() bool {
	return len(s) == 0
}

var bufDataMeta = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Meta) MarshalJSON() ([]byte, error) {
	var result = bufDataMeta.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Meta) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if wantComma {
		result.WriteString(",")
	}
	if s.Count != 0 {
		result.WriteString(`"count":`)
		writeInt64(result, int64(s.Count))
		wantComma = true
	} else {
		result.WriteString(`"count":0`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Meta) IsZero() bool {
	if s.Count != 0 {
		return false
	}
	return true
}

var bufDataPersonal = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Personal) MarshalJSON() ([]byte, error) {
	var result = bufDataPersonal.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Personal) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if wantComma {
		result.WriteString(",")
	}
	if s.Name != "" {
		result.WriteString(`"name":`)
		writeString(result, string(s.Name))
		wantComma = true
	} else {
		result.WriteString(`"name":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Surname != "" {
		result.WriteString(`"surname":`)
		writeString(result, string(s.Surname))
		wantComma = true
	} else {
		result.WriteString(`"surname":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Patname != nil {
		result.WriteString(`"patname":`)
		writeString(result, string(*s.Patname))
		wantComma = true
	} else {
		result.WriteString(`"patname":null`)
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Personal) IsZero() bool {
	if s.Name != "" {
		return false
	}
	if s.Surname != "" {
		return false
	}
	if s.Patname != nil {
		return false
	}
	return true
}

var bufDataTags = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Tags) MarshalJSON() ([]byte, error) {
	var result = bufDataTags.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Tags) MarshalTo(result Writer) error {
	if s == nil || *s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	for _k, _v := range *s {
		if wantComma {
			result.WriteString(",")
		}
		wantComma = true
		result.WriteString(`"`)
		result.WriteString(string(_k))
		result.WriteString(`":`)
		writeString(result, string(_v))
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Tags) IsZero() bool {
	return len(s) == 0
}
