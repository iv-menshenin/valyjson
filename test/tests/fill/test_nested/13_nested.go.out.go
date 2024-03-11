// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_nested

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/mailru/easyjson/jwriter"
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
		var valMeta = &s.Meta
		err = valMeta.FillFromJSON(_meta)
		if err != nil {
			return newParsingError("meta", err)
		}
	}
	if _data := v.Get("data"); _data != nil {
		var listA []*fastjson.Value
		listA, err = _data.Array()
		if err != nil {
			return newParsingError("data", err)
		}
		valData := s.Data
		if l := len(listA); cap(valData) < l || (l == 0 && s.Data == nil) {
			valData = make([]Middle, 0, len(listA))
		} else {
			valData = s.Data[:0]
		}
		for _key1, _val1 := range listA {
			valData = valData[:len(valData)+1]
			var _elem1 = &valData[len(valData)-1]
			err = _elem1.FillFromJSON(_val1)
			if err != nil {
				err = newParsingError(strconv.Itoa(_key1), err)
				break
			}
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
		s.Name = UserName(valName)
	}
	if _surname := v.Get("surname"); _surname != nil {
		var valSurname []byte
		if valSurname, err = _surname.StringBytes(); err != nil {
			return newParsingError("surname", err)
		}
		s.Surname = UserSurname(valSurname)
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
		valTags := s.Tags
		if valTags == nil {
			valTags = make(map[TagName]TagValue, o.Len())
		}
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
		s.Name = UserName(valName)
	}
	if _surname := v.Get("surname"); _surname != nil {
		var valSurname []byte
		if valSurname, err = _surname.StringBytes(); err != nil {
			return newParsingError("surname", err)
		}
		s.Surname = UserSurname(valSurname)
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

// jsonParserCustomEvent used for pooling Parsers for CustomEvent JSONs.
var jsonParserCustomEvent fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *CustomEvent) UnmarshalJSON(data []byte) error {
	parser := jsonParserCustomEvent.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserCustomEvent.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *CustomEvent) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _wRRetry := v; _wRRetry != nil {
		var valWRRetry = &s.WRRetry
		err = valWRRetry.FillFromJSON(_wRRetry)
		if err != nil {
			return newParsingError("", err)
		}
	}
	return nil
}

// validate checks for correct data structure
func (s *CustomEvent) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserWRRetry used for pooling Parsers for WRRetry JSONs.
var jsonParserWRRetry fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *WRRetry) UnmarshalJSON(data []byte) error {
	parser := jsonParserWRRetry.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserWRRetry.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *WRRetry) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _wRRetry := v.Get("WR-Retry"); _wRRetry != nil {
		var valWRRetry int
		valWRRetry, err = _wRRetry.Int()
		if err != nil {
			return newParsingError("WR-Retry", err)
		}
		s.WRRetry = valWRRetry
	}
	return nil
}

// validate checks for correct data structure
func (s *WRRetry) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'W', 'R', '-', 'R', 'e', 't', 'r', 'y'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Root) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Root) MarshalTo(result *jwriter.Writer) error {
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
	result.RawString(`"meta":`)
	if err = s.Meta.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "meta" attribute: %w`, err)
	}
	wantComma = true
	if wantComma {
		result.RawByte(',')
	}
	if s.Data != nil {
		wantComma = true
		result.RawString(`"data":[`)
		var wantComma bool
		for _k, _v := range s.Data {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "data" item at position %d: %w`, _k, err)
			}
		}
		result.RawByte(']')
		wantComma = true
	} else {
		result.RawString(`"data":null`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Root) Reset() {
	s.Meta.Reset()
	for i := range s.Data {
		s.Data[i].Reset()
	}
	s.Data = s.Data[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Middle) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Middle) MarshalTo(result *jwriter.Writer) error {
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
		result.String(string(s.Name))
		wantComma = true
	} else {
		result.RawString(`"name":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Surname != "" {
		result.RawString(`"surname":`)
		result.String(string(s.Surname))
		wantComma = true
	} else {
		result.RawString(`"surname":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Patname != nil {
		result.RawString(`"patname":`)
		result.String(string(*s.Patname))
		wantComma = true
	} else {
		result.RawString(`"patname":null`)
	}
	if wantComma {
		result.RawByte(',')
	}
	if !s.DateOfBorn.IsZero() {
		result.RawString(`"dateOfBorn":`)
		writeTime(result, s.DateOfBorn, time.RFC3339Nano)
		wantComma = true
	} else {
		result.RawString(`"dateOfBorn":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
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
			result.String(string(_k))
			result.RawByte(':')
			result.String(string(_v))
		}
		result.RawByte('}')
	} else {
		wantComma = true
		result.RawString(`"tags":null`)
	}
	result.RawByte('}')
	err = result.Error
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Middle) Reset() {
	s.Name = UserName("")
	s.Surname = UserSurname("")
	s.Patname = nil
	s.DateOfBorn = time.Time{}
	if len(s.Tags) > 10000 {
		s.Tags = nil
	} else {
		for key := range s.Tags {
			delete(s.Tags, key)
		}
	}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Middles) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Middles) MarshalTo(result *jwriter.Writer) error {
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
			return fmt.Errorf(`can't marshal "Middles" value at position %d: %w`, _k, err)
		}
	}
	result.RawByte(']')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s Middles) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Middles) Reset() {
	for i := range *s {
		(*s)[i].Reset()
	}
	*s = (*s)[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Meta) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Meta) MarshalTo(result *jwriter.Writer) error {
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
	if s.Count != 0 {
		result.RawString(`"count":`)
		result.Int64(int64(s.Count))
		wantComma = true
	} else {
		result.RawString(`"count":0`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s Meta) IsZero() bool {
	if s.Count != 0 {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Meta) Reset() {
	s.Count = 0
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Personal) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Personal) MarshalTo(result *jwriter.Writer) error {
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
		result.String(string(s.Name))
		wantComma = true
	} else {
		result.RawString(`"name":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Surname != "" {
		result.RawString(`"surname":`)
		result.String(string(s.Surname))
		wantComma = true
	} else {
		result.RawString(`"surname":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Patname != nil {
		result.RawString(`"patname":`)
		result.String(string(*s.Patname))
		wantComma = true
	} else {
		result.RawString(`"patname":null`)
	}
	result.RawByte('}')
	err = result.Error
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Personal) Reset() {
	s.Name = UserName("")
	s.Surname = UserSurname("")
	s.Patname = nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Tags) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Tags) MarshalTo(result *jwriter.Writer) error {
	if s == nil || *s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	for _k, _v := range *s {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.String(string(_k))
		result.RawByte(':')
		result.String(string(_v))
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s Tags) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s Tags) Reset() {
	for k, v := range s {
		v = TagValue("")
		s[k] = v
	}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *CustomEvent) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *CustomEvent) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if !s.WRRetry.IsZero() {
		if wantComma {
			result.RawByte(',')
		}
		result.Raw(unpackObject(s.WRRetry.MarshalJSON()))
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s CustomEvent) IsZero() bool {
	if !s.WRRetry.IsZero() {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *CustomEvent) Reset() {
	s.WRRetry.Reset()
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *WRRetry) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *WRRetry) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if s.WRRetry != 0 {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"WR-Retry":`)
		result.Int64(int64(s.WRRetry))
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s WRRetry) IsZero() bool {
	if s.WRRetry != 0 {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *WRRetry) Reset() {
	s.WRRetry = 0
}
