// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package vjson

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"unsafe"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/fastjson"
)

// jsonParserPerson used for pooling Parsers for Person JSONs.
var jsonParserPerson fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Person) UnmarshalJSON(data []byte) error {
	parser := jsonParserPerson.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserPerson.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Person) FillFromJSON(v *fastjson.Value) (err error) {
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
	if _surname := v.Get("surname"); _surname != nil {
		var valSurname []byte
		if valSurname, err = _surname.StringBytes(); err != nil {
			return newParsingError("surname", err)
		}
		s.Surname = string(valSurname)
	}
	if _middle := v.Get("middle"); valueIsNotNull(_middle) {
		var valMiddle []byte
		if valMiddle, err = _middle.StringBytes(); err != nil {
			return newParsingError("middle", err)
		}
		s.Middle = (*string)(unsafe.Pointer(&valMiddle))
	}
	if _dOB := v.Get("dob"); valueIsNotNull(_dOB) {
		b, err := _dOB.StringBytes()
		if err != nil {
			return newParsingError("dob", err)
		}
		valDOB, err := parseDateTime(string(b))
		if err != nil {
			return newParsingError("dob", err)
		}
		s.DOB = new(time.Time)
		*s.DOB = time.Time(valDOB)
	}
	if _passport := v.Get("passport"); valueIsNotNull(_passport) {
		var valPassport Passport
		err = valPassport.FillFromJSON(_passport)
		if err != nil {
			return newParsingError("passport", err)
		}
		s.Passport = new(Passport)
		*s.Passport = Passport(valPassport)
	}
	if _tables := v.Get("tables"); _tables != nil {
		o, err := _tables.Object()
		if err != nil {
			return newParsingError("tables", err)
		}
		valTables := s.Tables
		if valTables == nil {
			valTables = make(map[string]TableOf, o.Len())
		}
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value TableOf
			err = value.FillFromJSON(v)
			if err != nil {
				err = newParsingError(string(key), err)
			} else {
				valTables[string(key)] = TableOf(value)
			}
		})
		if err != nil {
			return newParsingError("tables", err)
		}
		s.Tables = MapTable(valTables)
	}
	return nil
}

// validate checks for correct data structure
func (s *Person) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [6]int
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
		if bytes.Equal(key, []byte{'m', 'i', 'd', 'd', 'l', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'o', 'b'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'a', 's', 's', 'p', 'o', 'r', 't'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'b', 'l', 'e', 's'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
	})
	return err
}

// jsonParserPassport used for pooling Parsers for Passport JSONs.
var jsonParserPassport fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Passport) UnmarshalJSON(data []byte) error {
	parser := jsonParserPassport.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserPassport.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Passport) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _number := v.Get("number"); _number != nil {
		var valNumber []byte
		if valNumber, err = _number.StringBytes(); err != nil {
			return newParsingError("number", err)
		}
		s.Number = string(valNumber)
	}
	if _dateDoc := v.Get("dateDoc"); _dateDoc != nil {
		b, err := _dateDoc.StringBytes()
		if err != nil {
			return newParsingError("dateDoc", err)
		}
		valDateDoc, err := parseDateTime(string(b))
		if err != nil {
			return newParsingError("dateDoc", err)
		}
		s.DateDoc = valDateDoc
	}
	return nil
}

// validate checks for correct data structure
func (s *Passport) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'n', 'u', 'm', 'b', 'e', 'r'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', 'D', 'o', 'c'}) {
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

// jsonParserTableOf used for pooling Parsers for TableOf JSONs.
var jsonParserTableOf fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TableOf) UnmarshalJSON(data []byte) error {
	parser := jsonParserTableOf.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTableOf.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TableOf) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _tableName := v.Get("tableName"); _tableName != nil {
		var valTableName []byte
		if valTableName, err = _tableName.StringBytes(); err != nil {
			return newParsingError("tableName", err)
		}
		s.TableName = string(valTableName)
	}
	if _tables := v.Get("tables"); valueIsNotNull(_tables) {
		var listA []*fastjson.Value
		listA, err = _tables.Array()
		if err != nil {
			return newParsingError("tables", err)
		}
		valTables := s.Tables[:0]
		if l := len(listA); cap(valTables) < l || (l == 0 && s.Tables == nil) {
			valTables = make([]*Table, 0, len(listA))
		}
		for _key, _val := range listA {
			valTables = valTables[:len(valTables)+1]
			if !valueIsNotNull(_val) {
				valTables[len(valTables)-1] = nil
				continue
			}
			var elem Table
			err = elem.FillFromJSON(_val)
			if err != nil {
				err = newParsingError(strconv.Itoa(_key), err)
				break
			}
			newElem := Table(elem)
			valTables[_key] = &newElem
		}
		if err != nil {
			return newParsingError("tables", err)
		}
		s.Tables = valTables
	}
	return nil
}

// validate checks for correct data structure
func (s *TableOf) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'b', 'l', 'e', 'N', 'a', 'm', 'e'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'b', 'l', 'e', 's'}) {
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

// jsonParserTable used for pooling Parsers for Table JSONs.
var jsonParserTable fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Table) UnmarshalJSON(data []byte) error {
	parser := jsonParserTable.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTable.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Table) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _counter := v.Get("counter"); _counter != nil {
		var valCounter int
		valCounter, err = _counter.Int()
		if err != nil {
			return newParsingError("counter", err)
		}
		s.Counter = valCounter
	}
	if _assessments := v.Get("assessments"); valueIsNotNull(_assessments) {
		var listA []*fastjson.Value
		listA, err = _assessments.Array()
		if err != nil {
			return newParsingError("assessments", err)
		}
		valAssessments := s.Assessments[:0]
		if l := len(listA); cap(valAssessments) < l || (l == 0 && s.Assessments == nil) {
			valAssessments = make([]int, 0, len(listA))
		}
		for _key, _val := range listA {
			valAssessments = valAssessments[:len(valAssessments)+1]
			var elem int
			elem, err = _val.Int()
			if err != nil {
				err = newParsingError(strconv.Itoa(_key), err)
				break
			}
			valAssessments[_key] = int(elem)
		}
		if err != nil {
			return newParsingError("assessments", err)
		}
		s.Assessments = valAssessments
	}
	if _time := v.Get("time"); _time != nil {
		b, err := _time.StringBytes()
		if err != nil {
			return newParsingError("time", err)
		}
		valTime, err := parseDateTime(string(b))
		if err != nil {
			return newParsingError("time", err)
		}
		s.Time = valTime
	}
	if _avg := v.Get("avg"); _avg != nil {
		var valAvg float64
		valAvg, err = _avg.Float64()
		if err != nil {
			return newParsingError("avg", err)
		}
		s.Avg = valAvg
	}
	if _tags := v.Get("tags"); valueIsNotNull(_tags) {
		var listA []*fastjson.Value
		listA, err = _tags.Array()
		if err != nil {
			return newParsingError("tags", err)
		}
		valTags := s.Tags[:0]
		if l := len(listA); cap(valTags) < l || (l == 0 && s.Tags == nil) {
			valTags = make([]Tag, 0, len(listA))
		}
		for _key, _val := range listA {
			valTags = valTags[:len(valTags)+1]
			var elem = &valTags[len(valTags)-1]
			err = elem.FillFromJSON(_val)
			if err != nil {
				err = newParsingError(strconv.Itoa(_key), err)
				break
			}
		}
		if err != nil {
			return newParsingError("tags", err)
		}
		s.Tags = valTags
	}
	return nil
}

// validate checks for correct data structure
func (s *Table) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [5]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'u', 'n', 't', 'e', 'r'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'a', 's', 's', 'e', 's', 's', 'm', 'e', 'n', 't', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'i', 'm', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'a', 'v', 'g'}) {
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
		err = fmt.Errorf("unexpected field '%s'", string(key))
	})
	return err
}

// jsonParserTag used for pooling Parsers for Tag JSONs.
var jsonParserTag fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Tag) UnmarshalJSON(data []byte) error {
	parser := jsonParserTag.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTag.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Tag) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _tagName := v.Get("tagName"); _tagName != nil {
		var valTagName []byte
		if valTagName, err = _tagName.StringBytes(); err != nil {
			return newParsingError("tagName", err)
		}
		s.TagName = string(valTagName)
	}
	if _tagValue := v.Get("tagValue"); _tagValue != nil {
		var valTagValue []byte
		if valTagValue, err = _tagValue.StringBytes(); err != nil {
			return newParsingError("tagValue", err)
		}
		s.TagValue = string(valTagValue)
	}
	return nil
}

// validate checks for correct data structure
func (s *Tag) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 'N', 'a', 'm', 'e'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 'V', 'a', 'l', 'u', 'e'}) {
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

// jsonParserMapTable used for pooling Parsers for MapTable JSONs.
var jsonParserMapTable fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *MapTable) UnmarshalJSON(data []byte) error {
	parser := jsonParserMapTable.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserMapTable.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *MapTable) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	o, err := v.Object()
	if err != nil {
		return err
	}
	*s = make(map[string]TableOf, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value TableOf
		err = value.FillFromJSON(v)
		if err != nil {
			err = newParsingError(string(key), err)
			return
		}
		(*s)[string(key)] = TableOf(value)
	})
	return err
}

// jsonParserMapInt64 used for pooling Parsers for MapInt64 JSONs.
var jsonParserMapInt64 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *MapInt64) UnmarshalJSON(data []byte) error {
	parser := jsonParserMapInt64.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserMapInt64.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *MapInt64) FillFromJSON(v *fastjson.Value) (err error) {
	if v.Type() == fastjson.TypeNull {
		return nil
	}
	o, err := v.Object()
	if err != nil {
		return err
	}
	*s = make(map[string]int64, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value int64
		value, err = v.Int64()
		if err != nil {
			err = newParsingError(string(key), err)
			return
		}
		(*s)[string(key)] = int64(value)
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Person) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Person) MarshalTo(result *jwriter.Writer) error {
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
	if s.Surname != "" {
		result.RawString(`"surname":`)
		result.String(s.Surname)
		wantComma = true
	} else {
		result.RawString(`"surname":""`)
		wantComma = true
	}
	if s.Middle != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"middle":`)
		result.String(*s.Middle)
		wantComma = true
	}
	if s.DOB != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"dob":`)
		writeTime(result, *s.DOB, time.RFC3339Nano)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Passport != nil {
		result.RawString(`"passport":`)
		if err = s.Passport.MarshalTo(result); err != nil {
			return fmt.Errorf(`can't marshal "passport" attribute: %w`, err)
		}
		wantComma = true
	} else {
		result.RawString(`"passport":null`)
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Tables != nil {
		wantComma = true
		result.RawString(`"tables":{`)
		var wantComma bool
		for _k, _v := range s.Tables {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			result.String(_k)
			result.RawByte(':')
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "tables" attribute %q: %w`, _k, err)
			}
		}
		result.RawByte('}')
	} else {
		wantComma = true
		result.RawString(`"tables":null`)
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s Person) IsZero() bool {
	if s.Name != "" {
		return false
	}
	if s.Surname != "" {
		return false
	}
	if s.Middle != nil {
		return false
	}
	if s.DOB != nil {
		return false
	}
	if s.Passport != nil {
		return false
	}
	if s.Tables != nil {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Person) Reset() {
	s.Name = ""
	s.Surname = ""
	s.Middle = nil
	s.DOB = nil
	s.Passport = nil
	if len(s.Tables) > 10000 {
		s.Tables = nil
	} else {
		for key := range s.Tables {
			delete(s.Tables, key)
		}
	}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Passport) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Passport) MarshalTo(result *jwriter.Writer) error {
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
	if s.Number != "" {
		result.RawString(`"number":`)
		result.String(s.Number)
		wantComma = true
	} else {
		result.RawString(`"number":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if !s.DateDoc.IsZero() {
		result.RawString(`"dateDoc":`)
		writeTime(result, s.DateDoc, time.RFC3339Nano)
		wantComma = true
	} else {
		result.RawString(`"dateDoc":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s Passport) IsZero() bool {
	if s.Number != "" {
		return false
	}
	if !s.DateDoc.IsZero() {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Passport) Reset() {
	s.Number = ""
	s.DateDoc = time.Time{}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TableOf) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TableOf) MarshalTo(result *jwriter.Writer) error {
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
	if s.TableName != "" {
		result.RawString(`"tableName":`)
		result.String(s.TableName)
		wantComma = true
	} else {
		result.RawString(`"tableName":""`)
		wantComma = true
	}
	if s.Tables != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"tables":[`)
		var wantComma bool
		for _k, _v := range s.Tables {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "tables" item at position %d: %w`, _k, err)
			}
		}
		result.RawByte(']')
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TableOf) IsZero() bool {
	if s.TableName != "" {
		return false
	}
	if s.Tables != nil {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TableOf) Reset() {
	s.TableName = ""
	for i := range s.Tables {
		s.Tables[i] = nil
	}
	s.Tables = s.Tables[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Table) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Table) MarshalTo(result *jwriter.Writer) error {
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
	if s.Counter != 0 {
		result.RawString(`"counter":`)
		result.Int64(int64(s.Counter))
		wantComma = true
	} else {
		result.RawString(`"counter":0`)
		wantComma = true
	}
	if s.Assessments != nil {
		if wantComma {
			result.RawByte(',')
		}
		wantComma = true
		result.RawString(`"assessments":[`)
		var wantComma bool
		for _k, _v := range s.Assessments {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			result.Int64(int64(_v))
		}
		result.RawByte(']')
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if !s.Time.IsZero() {
		result.RawString(`"time":`)
		writeTime(result, s.Time, time.RFC3339Nano)
		wantComma = true
	} else {
		result.RawString(`"time":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Avg != 0 {
		result.RawString(`"avg":`)
		result.Float64(s.Avg)
		wantComma = true
	} else {
		result.RawString(`"avg":0`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.Tags != nil {
		wantComma = true
		result.RawString(`"tags":[`)
		var wantComma bool
		for _k, _v := range s.Tags {
			if wantComma {
				result.RawByte(',')
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "tags" item at position %d: %w`, _k, err)
			}
		}
		result.RawByte(']')
		wantComma = true
	} else {
		result.RawString(`"tags":null`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s Table) IsZero() bool {
	if s.Counter != 0 {
		return false
	}
	if s.Assessments != nil {
		return false
	}
	if !s.Time.IsZero() {
		return false
	}
	if s.Avg != 0 {
		return false
	}
	if s.Tags != nil {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Table) Reset() {
	s.Counter = 0
	for i := range s.Assessments {
		s.Assessments[i] = 0
	}
	s.Assessments = s.Assessments[:0]
	s.Time = time.Time{}
	s.Avg = 0
	for i := range s.Tags {
		s.Tags[i].Reset()
	}
	s.Tags = s.Tags[:0]
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Tag) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Tag) MarshalTo(result *jwriter.Writer) error {
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
	if s.TagName != "" {
		result.RawString(`"tagName":`)
		result.String(s.TagName)
		wantComma = true
	} else {
		result.RawString(`"tagName":""`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.TagValue != "" {
		result.RawString(`"tagValue":`)
		result.String(s.TagValue)
		wantComma = true
	} else {
		result.RawString(`"tagValue":""`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s Tag) IsZero() bool {
	if s.TagName != "" {
		return false
	}
	if s.TagValue != "" {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *Tag) Reset() {
	s.TagName = ""
	s.TagValue = ""
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *MapTable) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *MapTable) MarshalTo(result *jwriter.Writer) error {
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
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "MapTable" attribute %q: %w`, _k, err)
		}
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s MapTable) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s MapTable) Reset() {
	for k, v := range s {
		v.Reset()
		s[k] = v
	}
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *MapInt64) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *MapInt64) MarshalTo(result *jwriter.Writer) error {
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
		result.Int64(int64(_v))
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s MapInt64) IsZero() bool {
	return len(s) == 0
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s MapInt64) Reset() {
	for k, v := range s {
		v = 0
		s[k] = v
	}
}
