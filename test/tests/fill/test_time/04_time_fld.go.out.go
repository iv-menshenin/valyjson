// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_time

import (
	"bytes"
	"fmt"
	"time"

	"github.com/valyala/fastjson"
)

// jsonParserTestTime01 used for pooling Parsers for TestTime01 JSONs.
var jsonParserTestTime01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestTime01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestTime01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestTime01.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestTime01) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _dateBegin := v.Get("date_begin"); _dateBegin != nil {
		b, err := _dateBegin.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.date_begin' value: %w", objPath, err)
		}
		valDateBegin, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.date_begin' value: %w", objPath, err)
		}
		s.DateBegin = valDateBegin
	}
	if _dateCustom := v.Get("date_custom"); _dateCustom != nil {
		b, err := _dateCustom.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.date_custom' value: %w", objPath, err)
		}
		valDateCustom, err := time.Parse("2006.01.02", string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.date_custom' value: %w", objPath, err)
		}
		s.DateCustom = valDateCustom
	}
	if _dateEnd := v.Get("date_end"); valueIsNotNull(_dateEnd) {
		b, err := _dateEnd.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.date_end' value: %w", objPath, err)
		}
		valDateEnd, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.date_end' value: %w", objPath, err)
		}
		s.DateEnd = new(time.Time)
		*s.DateEnd = time.Time(valDateEnd)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestTime01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [3]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'b', 'e', 'g', 'i', 'n'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'c', 'u', 's', 't', 'o', 'm'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'e', 'n', 'd'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// jsonParserTestTime2 used for pooling Parsers for TestTime2 JSONs.
var jsonParserTestTime2 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestTime2) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestTime2.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestTime2.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestTime2) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	b, err := v.StringBytes()
	if err != nil {
		return fmt.Errorf("error parsing '%s.' value: %w", objPath, err)
	}
	_val, err := parseDateTime(string(b))
	if err != nil {
		return err
	}
	*s = TestTime2(_val)
	return nil
}

var bufDataTestTime01 = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestTime01) MarshalJSON() ([]byte, error) {
	var result = bufDataTestTime01.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestTime01) MarshalTo(result Writer) error {
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
	if !s.DateBegin.IsZero() {
		result.WriteString(`"date_begin":`)
		writeTime(result, s.DateBegin, time.RFC3339Nano)
		wantComma = true
	} else {
		result.WriteString(`"date_begin":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if !s.DateCustom.IsZero() {
		result.WriteString(`"date_custom":`)
		writeTime(result, s.DateCustom, "2006.01.02")
		wantComma = true
	} else {
		result.WriteString(`"date_custom":"0001.01.01"`)
		wantComma = true
	}
	if s.DateEnd != nil {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"date_end":`)
		writeTime(result, *s.DateEnd, time.RFC3339Nano)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestTime01) IsZero() bool {
	if !s.DateBegin.IsZero() {
		return false
	}
	if !s.DateCustom.IsZero() {
		return false
	}
	if s.DateEnd != nil {
		return false
	}
	return true
}

var bufDataTestTime2 = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestTime2) MarshalJSON() ([]byte, error) {
	var result = bufDataTestTime2.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestTime2) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	_time, err := time.Time(*s).MarshalText()
	if err == nil {
		result.Write(_time)
	}
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestTime2) IsZero() bool {
	return time.Time(s).IsZero()
}
