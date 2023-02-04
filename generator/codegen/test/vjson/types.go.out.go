// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package vjson

import (
	"bytes"
	"fmt"
	"time"

	"github.com/valyala/fastjson"
)

// jsonParserPersonused for pooling Parsers for Person JSONs.
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Person) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if name := v.Get("name"); name != nil {
		var valName []byte
		if valName, err = name.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%sname' value: %w", objPath, err)
		}
		s.Name = string(valName)
	}
	if surname := v.Get("surname"); surname != nil {
		var valSurname []byte
		if valSurname, err = surname.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%ssurname' value: %w", objPath, err)
		}
		s.Surname = string(valSurname)
	}
	if middle := v.Get("middle"); valueIsNotNull(middle) {
		var valMiddle []byte
		if valMiddle, err = middle.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%smiddle' value: %w", objPath, err)
		}
		s.Middle = new(string)
		*s.Middle = string(valMiddle)
	}
	if dob := v.Get("dob"); valueIsNotNull(dob) {
		b, err := dob.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%sdob' value: %w", objPath, err)
		}
		valDOB, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%sdob' value: %w", objPath, err)
		}
		s.DOB = new(time.Time)
		*s.DOB = time.Time(valDOB)
	}
	if passport := v.Get("passport"); valueIsNotNull(passport) {
		var valPassport Passport
		err = valPassport.FillFromJson(passport, objPath+"passport.")
		if err != nil {
			return fmt.Errorf("error parsing '%spassport' value: %w", objPath, err)
		}
		s.Passport = new(Passport)
		*s.Passport = Passport(valPassport)
	}
	if tables := v.Get("tables"); valueIsNotNull(tables) {
		o, err := tables.Object()
		if err != nil {
			return fmt.Errorf("error parsing '%stables' value: %w", objPath, err)
		}
		var valTables = make(map[string]TableOf, o.Len())
		o.Visit(func(key []byte, v *fastjson.Value) {
			if err != nil {
				return
			}
			var value TableOf
			err = value.FillFromJson(v, objPath+"tables.")
			if err == nil {
				valTables[string(key)] = TableOf(value)
			}
		})
		if err != nil {
			return fmt.Errorf("error parsing '%stables' value: %w", objPath, err)
		}
		s.Tables = valTables
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
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'u', 'r', 'n', 'a', 'm', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'m', 'i', 'd', 'd', 'l', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'o', 'b'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'a', 's', 's', 'p', 'o', 'r', 't'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'b', 'l', 'e', 's'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// jsonParserPassportused for pooling Parsers for Passport JSONs.
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Passport) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if number := v.Get("number"); number != nil {
		var valNumber []byte
		if valNumber, err = number.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%snumber' value: %w", objPath, err)
		}
		s.Number = string(valNumber)
	}
	if datedoc := v.Get("dateDoc"); datedoc != nil {
		b, err := datedoc.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%sdateDoc' value: %w", objPath, err)
		}
		valDateDoc, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%sdateDoc' value: %w", objPath, err)
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
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', 'D', 'o', 'c'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// jsonParserTableOfused for pooling Parsers for TableOf JSONs.
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TableOf) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if tablename := v.Get("tableName"); tablename != nil {
		var valTableName []byte
		if valTableName, err = tablename.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%stableName' value: %w", objPath, err)
		}
		s.TableName = string(valTableName)
	}
	if tables := v.Get("tables"); valueIsNotNull(tables) {
		var listA []*fastjson.Value
		listA, err = tables.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%stables' value: %w", objPath, err)
		}
		valTables := make([]*Table, 0, len(listA))
		for _, listElem := range listA {
			if !valueIsNotNull(listElem) {
				valTables = append(valTables, nil)
				continue
			}
			var elem Table
			err = elem.FillFromJson(listElem, objPath+".")
			if err != nil {
				break
			}
			newElem := Table(elem)
			valTables = append(valTables, &newElem)
		}
		if err != nil {
			return fmt.Errorf("error parsing '%stables' value: %w", objPath, err)
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
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'b', 'l', 'e', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// jsonParserTableused for pooling Parsers for Table JSONs.
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Table) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if counter := v.Get("counter"); counter != nil {
		var valCounter int
		valCounter, err = counter.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%scounter' value: %w", objPath, err)
		}
		s.Counter = valCounter
	}
	if assessments := v.Get("assessments"); valueIsNotNull(assessments) {
		var listA []*fastjson.Value
		listA, err = assessments.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%sassessments' value: %w", objPath, err)
		}
		valAssessments := make([]int, 0, len(listA))
		for _, listElem := range listA {
			var elem int
			elem, err = listElem.Int()
			if err != nil {
				break
			}
			valAssessments = append(valAssessments, int(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%sassessments' value: %w", objPath, err)
		}
		s.Assessments = valAssessments
	}
	if time := v.Get("time"); time != nil {
		b, err := time.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%stime' value: %w", objPath, err)
		}
		valTime, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%stime' value: %w", objPath, err)
		}
		s.Time = valTime
	}
	if avg := v.Get("avg"); avg != nil {
		var valAvg float64
		valAvg, err = avg.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%savg' value: %w", objPath, err)
		}
		s.Avg = valAvg
	}
	if tags := v.Get("tags"); valueIsNotNull(tags) {
		var listA []*fastjson.Value
		listA, err = tags.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%stags' value: %w", objPath, err)
		}
		valTags := make([]Tag, 0, len(listA))
		for _, listElem := range listA {
			var elem Tag
			err = elem.FillFromJson(listElem, objPath+".")
			if err != nil {
				break
			}
			valTags = append(valTags, Tag(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%stags' value: %w", objPath, err)
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
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'a', 's', 's', 'e', 's', 's', 'm', 'e', 'n', 't', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'i', 'm', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'a', 'v', 'g'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 's'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// jsonParserTagused for pooling Parsers for Tag JSONs.
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Tag) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if tagname := v.Get("tagName"); tagname != nil {
		var valTagName []byte
		if valTagName, err = tagname.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%stagName' value: %w", objPath, err)
		}
		s.TagName = string(valTagName)
	}
	if tagvalue := v.Get("tagValue"); tagvalue != nil {
		var valTagValue []byte
		if valTagValue, err = tagvalue.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%stagValue' value: %w", objPath, err)
		}
		s.TagValue = string(valTagValue)
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
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'a', 'g', 'V', 'a', 'l', 'u', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}
