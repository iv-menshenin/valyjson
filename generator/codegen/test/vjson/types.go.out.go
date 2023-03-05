// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package vjson

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"unsafe"

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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Person) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
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
	if _surname := v.Get("surname"); _surname != nil {
		var valSurname []byte
		if valSurname, err = _surname.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.surname' value: %w", objPath, err)
		}
		s.Surname = *(*string)(unsafe.Pointer(&valSurname))
	}
	if _middle := v.Get("middle"); valueIsNotNull(_middle) {
		var valMiddle []byte
		if valMiddle, err = _middle.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.middle' value: %w", objPath, err)
		}
		s.Middle = (*string)(unsafe.Pointer(&valMiddle))
	}
	if _dOB := v.Get("dob"); valueIsNotNull(_dOB) {
		b, err := _dOB.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.dob' value: %w", objPath, err)
		}
		valDOB, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.dob' value: %w", objPath, err)
		}
		s.DOB = new(time.Time)
		*s.DOB = time.Time(valDOB)
	}
	if _passport := v.Get("passport"); valueIsNotNull(_passport) {
		var valPassport Passport
		err = valPassport.FillFromJSON(_passport, objPath+".passport")
		if err != nil {
			return fmt.Errorf("error parsing '%s.passport' value: %w", objPath, err)
		}
		s.Passport = new(Passport)
		*s.Passport = Passport(valPassport)
	}
	if _tables := v.Get("tables"); _tables != nil {
		o, err := _tables.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
		}
		var valTables = make(map[string]TableOf, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value TableOf
			err = value.FillFromJSON(v, objPath+".tables")
			if err == nil {
				valTables[string(key)] = TableOf(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
		}
		s.Tables = MapTable(valTables)
	}
	return nil
}

// validate checks for correct data structure
func (s *Person) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'u', 'r', 'n', 'a', 'm', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'m', 'i', 'd', 'd', 'l', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'o', 'b'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'a', 's', 's', 'p', 'o', 'r', 't'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'b', 'l', 'e', 's'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Passport) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _number := v.Get("number"); _number != nil {
		var valNumber []byte
		if valNumber, err = _number.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.number' value: %w", objPath, err)
		}
		s.Number = *(*string)(unsafe.Pointer(&valNumber))
	}
	if _dateDoc := v.Get("dateDoc"); _dateDoc != nil {
		b, err := _dateDoc.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.dateDoc' value: %w", objPath, err)
		}
		valDateDoc, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.dateDoc' value: %w", objPath, err)
		}
		s.DateDoc = valDateDoc
	}
	return nil
}

// validate checks for correct data structure
func (s *Passport) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', 'D', 'o', 'c'}) {
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TableOf) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _tableName := v.Get("tableName"); _tableName != nil {
		var valTableName []byte
		if valTableName, err = _tableName.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.tableName' value: %w", objPath, err)
		}
		s.TableName = *(*string)(unsafe.Pointer(&valTableName))
	}
	if _tables := v.Get("tables"); valueIsNotNull(_tables) {
		var listA []*fastjson.Value
		listA, err = _tables.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
		}
		valTables := s.Tables[:0]
		if l := len(listA); cap(valTables) < l || (l == 0 && s.Tables == nil) {
			valTables = make([]*Table, 0, len(listA))
		}
		for _, listElem := range listA {
			if !valueIsNotNull(listElem) {
				valTables = append(valTables, nil)
				continue
			}
			var elem Table
			err = elem.FillFromJSON(listElem, objPath+".")
			if err != nil {
				break
			}
			newElem := Table(elem)
			valTables = append(valTables, &newElem)
		}
		if err != nil {
			return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
		}
		s.Tables = valTables
	}
	return nil
}

// validate checks for correct data structure
func (s *TableOf) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'b', 'l', 'e', 's'}) {
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Table) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _counter := v.Get("counter"); _counter != nil {
		var valCounter int
		valCounter, err = _counter.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.counter' value: %w", objPath, err)
		}
		s.Counter = valCounter
	}
	if _assessments := v.Get("assessments"); valueIsNotNull(_assessments) {
		var listA []*fastjson.Value
		listA, err = _assessments.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%s.assessments' value: %w", objPath, err)
		}
		valAssessments := s.Assessments[:0]
		if l := len(listA); cap(valAssessments) < l || (l == 0 && s.Assessments == nil) {
			valAssessments = make([]int, 0, len(listA))
		}
		for _, listElem := range listA {
			var elem int
			elem, err = listElem.Int()
			if err != nil {
				break
			}
			valAssessments = append(valAssessments, int(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%s.assessments' value: %w", objPath, err)
		}
		s.Assessments = valAssessments
	}
	if _time := v.Get("time"); _time != nil {
		b, err := _time.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.time' value: %w", objPath, err)
		}
		valTime, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.time' value: %w", objPath, err)
		}
		s.Time = valTime
	}
	if _avg := v.Get("avg"); _avg != nil {
		var valAvg float64
		valAvg, err = _avg.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%s.avg' value: %w", objPath, err)
		}
		s.Avg = valAvg
	}
	if _tags := v.Get("tags"); valueIsNotNull(_tags) {
		var listA []*fastjson.Value
		listA, err = _tags.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%s.tags' value: %w", objPath, err)
		}
		valTags := s.Tags[:0]
		if l := len(listA); cap(valTags) < l || (l == 0 && s.Tags == nil) {
			valTags = make([]Tag, 0, len(listA))
		}
		for _, listElem := range listA {
			var elem Tag
			err = elem.FillFromJSON(listElem, objPath+".")
			if err != nil {
				break
			}
			valTags = append(valTags, Tag(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%s.tags' value: %w", objPath, err)
		}
		s.Tags = valTags
	}
	return nil
}

// validate checks for correct data structure
func (s *Table) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'a', 's', 's', 'e', 's', 's', 'm', 'e', 'n', 't', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'i', 'm', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'a', 'v', 'g'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 's'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Tag) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _tagName := v.Get("tagName"); _tagName != nil {
		var valTagName []byte
		if valTagName, err = _tagName.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.tagName' value: %w", objPath, err)
		}
		s.TagName = *(*string)(unsafe.Pointer(&valTagName))
	}
	if _tagValue := v.Get("tagValue"); _tagValue != nil {
		var valTagValue []byte
		if valTagValue, err = _tagValue.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.tagValue' value: %w", objPath, err)
		}
		s.TagValue = *(*string)(unsafe.Pointer(&valTagValue))
	}
	return nil
}

// validate checks for correct data structure
func (s *Tag) validate(v *fastjson.Value, objPath string) error {
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 'V', 'a', 'l', 'u', 'e'}) {
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *MapTable) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	o, err := v.Object()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make(map[string]TableOf, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value TableOf
		err = value.FillFromJSON(v, objPath+".")
		if err != nil {
			err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *MapInt64) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	o, err := v.Object()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make(map[string]int64, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value int64
		value, err = v.Int64()
		if err != nil {
			err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
			return
		}
		(*s)[string(key)] = int64(value)
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Person) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *Person) MarshalAppend(dst []byte) ([]byte, error) {
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
	if s.Surname != "" {
		result.WriteString(`"surname":`)
		buf = marshalString(buf[:0], s.Surname)
		result.Write(buf)
	} else {
		result.WriteString(`"surname":""`)
	}
	if s.Middle != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"middle":`)
		buf = marshalString(buf[:0], *s.Middle)
		result.Write(buf)
	}
	if !s.DOB.IsZero() {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"dob":`)
		buf = marshalTime(buf[:0], s.DOB, time.RFC3339Nano)
		result.Write(buf)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Passport != nil {
		if buf, err = s.Passport.MarshalAppend(buf[:0]); err != nil {
			return nil, fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
		} else {
			result.WriteString(`"passport":`)
			result.Write(buf)
		}
	} else {
		result.WriteString(`"passport":null`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Tables != nil {
		buf = buf[:0]
		result.WriteString(`"tables":{`)
		var _filled bool
		for _k, _v := range s.Tables {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			result.WriteRune('"')
			result.WriteString(_k)
			result.WriteString(`":`)
			buf, err = _v.MarshalAppend(buf[:0])
			if err != nil {
				return nil, fmt.Errorf(`can't marshal "tables" attribute %q: %w`, _k, err)
			}
			result.Write(buf)
		}
		result.WriteRune('}')
	} else {
		result.WriteString(`"tables":null`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Passport) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *Passport) MarshalAppend(dst []byte) ([]byte, error) {
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
	if s.Number != "" {
		result.WriteString(`"number":`)
		buf = marshalString(buf[:0], s.Number)
		result.Write(buf)
	} else {
		result.WriteString(`"number":""`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if !s.DateDoc.IsZero() {
		result.WriteString(`"dateDoc":`)
		buf = marshalTime(buf[:0], s.DateDoc, time.RFC3339Nano)
		result.Write(buf)
	} else {
		result.WriteString(`"dateDoc":"0000-00-00T00:00:00Z"`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TableOf) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *TableOf) MarshalAppend(dst []byte) ([]byte, error) {
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
	if s.TableName != "" {
		result.WriteString(`"tableName":`)
		buf = marshalString(buf[:0], s.TableName)
		result.Write(buf)
	} else {
		result.WriteString(`"tableName":""`)
	}
	if s.Tables != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		buf = buf[:0]
		result.WriteString(`"tables":[`)
		var _filled bool
		for _k, _v := range s.Tables {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			buf, err = _v.MarshalAppend(buf[:0])
			if err != nil {
				return nil, fmt.Errorf(`can't marshal "tables" item at position %d: %w`, _k, err)
			}
			result.Write(buf)
		}
		result.WriteRune(']')
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Table) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *Table) MarshalAppend(dst []byte) ([]byte, error) {
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
	if s.Counter != 0 {
		result.WriteString(`"counter":`)
		buf = strconv.AppendInt(buf[:0], int64(s.Counter), 10)
		result.Write(buf)
	} else {
		result.WriteString(`"counter":0`)
	}
	if s.Assessments != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		buf = buf[:0]
		result.WriteString(`"assessments":[`)
		var _filled bool
		for _k, _v := range s.Assessments {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			buf = strconv.AppendInt(buf[:0], int64(_v), 10)
			result.Write(buf)
		}
		result.WriteRune(']')
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if !s.Time.IsZero() {
		result.WriteString(`"time":`)
		buf = marshalTime(buf[:0], s.Time, time.RFC3339Nano)
		result.Write(buf)
	} else {
		result.WriteString(`"time":"0000-00-00T00:00:00Z"`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Avg != 0 {
		result.WriteString(`"avg":`)
		buf = strconv.AppendFloat(buf[:0], s.Avg, 'f', -1, 64)
		result.Write(buf)
	} else {
		result.WriteString(`"avg":0`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.Tags != nil {
		buf = buf[:0]
		result.WriteString(`"tags":[`)
		var _filled bool
		for _k, _v := range s.Tags {
			if _filled {
				result.WriteRune(',')
			}
			_filled = true
			buf, err = _v.MarshalAppend(buf[:0])
			if err != nil {
				return nil, fmt.Errorf(`can't marshal "tags" item at position %d: %w`, _k, err)
			}
			result.Write(buf)
		}
		result.WriteRune(']')
	} else {
		result.WriteString(`"tags":null`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Tag) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *Tag) MarshalAppend(dst []byte) ([]byte, error) {
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
	if s.TagName != "" {
		result.WriteString(`"tagName":`)
		buf = marshalString(buf[:0], s.TagName)
		result.Write(buf)
	} else {
		result.WriteString(`"tagName":""`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if s.TagValue != "" {
		result.WriteString(`"tagValue":`)
		buf = marshalString(buf[:0], s.TagValue)
		result.Write(buf)
	} else {
		result.WriteString(`"tagValue":""`)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *MapTable) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *MapTable) MarshalAppend(dst []byte) ([]byte, error) {
	if s == nil || *s == nil {
		return []byte("null"), nil
	}
	var (
		err     error
		_filled bool
		buf     = make([]byte, 0, 128)
		result  = bytes.NewBuffer(dst)
	)
	result.WriteRune('{')
	for _k, _v := range *s {
		if _filled {
			result.WriteRune(',')
		}
		_filled = true
		result.WriteRune('"')
		result.WriteString(string(_k))
		result.WriteString(`":`)
		buf, err = _v.MarshalAppend(buf[:0])
		if err != nil {
			return nil, fmt.Errorf(`can't marshal "MapTable" attribute %q: %w`, _k, err)
		}
		result.Write(buf)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *MapInt64) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *MapInt64) MarshalAppend(dst []byte) ([]byte, error) {
	if s == nil || *s == nil {
		return []byte("null"), nil
	}
	var (
		err     error
		_filled bool
		buf     = make([]byte, 0, 128)
		result  = bytes.NewBuffer(dst)
	)
	result.WriteRune('{')
	for _k, _v := range *s {
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
	return result.Bytes(), err
}
