// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package vjson

import (
	"bytes"
	"fmt"
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
	if v.Type() == fastjson.TypeNull {
		return nil
	}
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
	if v.Type() == fastjson.TypeNull {
		return nil
	}
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
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Person) MarshalTo(result Writer) error {
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
		writeString(result, s.Name)
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
		writeString(result, s.Surname)
		wantComma = true
	} else {
		result.WriteString(`"surname":""`)
		wantComma = true
	}
	if s.Middle != nil {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"middle":`)
		writeString(result, *s.Middle)
		wantComma = true
	}
	if s.DOB != nil {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"dob":`)
		writeTime(result, *s.DOB, time.RFC3339Nano)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Passport != nil {
		result.WriteString(`"passport":`)
		if err = s.Passport.MarshalTo(result); err != nil {
			return fmt.Errorf(`can't marshal "passport" attribute: %w`, err)
		}
		wantComma = true
	} else {
		result.WriteString(`"passport":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Tables != nil {
		wantComma = true
		result.WriteString(`"tables":{`)
		var wantComma bool
		for _k, _v := range s.Tables {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			result.WriteString(`"`)
			result.WriteString(_k)
			result.WriteString(`":`)
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "tables" attribute %q: %w`, _k, err)
			}
		}
		result.WriteString("}")
	} else {
		wantComma = true
		result.WriteString(`"tables":null`)
	}
	result.WriteString("}")
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Passport) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Passport) MarshalTo(result Writer) error {
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
	if s.Number != "" {
		result.WriteString(`"number":`)
		writeString(result, s.Number)
		wantComma = true
	} else {
		result.WriteString(`"number":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if !s.DateDoc.IsZero() {
		result.WriteString(`"dateDoc":`)
		writeTime(result, s.DateDoc, time.RFC3339Nano)
		wantComma = true
	} else {
		result.WriteString(`"dateDoc":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	result.WriteString("}")
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TableOf) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TableOf) MarshalTo(result Writer) error {
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
	if s.TableName != "" {
		result.WriteString(`"tableName":`)
		writeString(result, s.TableName)
		wantComma = true
	} else {
		result.WriteString(`"tableName":""`)
		wantComma = true
	}
	if s.Tables != nil {
		if wantComma {
			result.WriteString(",")
		}
		wantComma = true
		result.WriteString(`"tables":[`)
		var wantComma bool
		for _k, _v := range s.Tables {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "tables" item at position %d: %w`, _k, err)
			}
		}
		result.WriteString("]")
	}
	result.WriteString("}")
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Table) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Table) MarshalTo(result Writer) error {
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
	if s.Counter != 0 {
		result.WriteString(`"counter":`)
		writeInt64(result, int64(s.Counter))
		wantComma = true
	} else {
		result.WriteString(`"counter":0`)
		wantComma = true
	}
	if s.Assessments != nil {
		if wantComma {
			result.WriteString(",")
		}
		wantComma = true
		result.WriteString(`"assessments":[`)
		var wantComma bool
		for _k, _v := range s.Assessments {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			writeInt64(result, int64(_v))
		}
		result.WriteString("]")
	}
	if wantComma {
		result.WriteString(",")
	}
	if !s.Time.IsZero() {
		result.WriteString(`"time":`)
		writeTime(result, s.Time, time.RFC3339Nano)
		wantComma = true
	} else {
		result.WriteString(`"time":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Avg != 0 {
		result.WriteString(`"avg":`)
		writeFloat64(result, s.Avg)
		wantComma = true
	} else {
		result.WriteString(`"avg":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Tags != nil {
		wantComma = true
		result.WriteString(`"tags":[`)
		var wantComma bool
		for _k, _v := range s.Tags {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "tags" item at position %d: %w`, _k, err)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"tags":null`)
		wantComma = true
	}
	result.WriteString("}")
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Tag) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Tag) MarshalTo(result Writer) error {
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
	if s.TagName != "" {
		result.WriteString(`"tagName":`)
		writeString(result, s.TagName)
		wantComma = true
	} else {
		result.WriteString(`"tagName":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.TagValue != "" {
		result.WriteString(`"tagValue":`)
		writeString(result, s.TagValue)
		wantComma = true
	} else {
		result.WriteString(`"tagValue":""`)
		wantComma = true
	}
	result.WriteString("}")
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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *MapTable) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *MapTable) MarshalTo(result Writer) error {
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
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "MapTable" attribute %q: %w`, _k, err)
		}
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s MapTable) IsZero() bool {
	return len(s) == 0
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *MapInt64) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *MapInt64) MarshalTo(result Writer) error {
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
		writeInt64(result, int64(_v))
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s MapInt64) IsZero() bool {
	return len(s) == 0
}
