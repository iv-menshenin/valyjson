// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_time

import (
	"bytes"
	"fmt"
	"time"

	"github.com/mailru/easyjson/jwriter"
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestTime01) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _dateBegin := v.Get("date_begin"); _dateBegin != nil {
		b, err := _dateBegin.StringBytes()
		if err != nil {
			return newParsingError("date_begin", err)
		}
		valDateBegin, err := parseDateTime(string(b))
		if err != nil {
			return newParsingError("date_begin", err)
		}
		s.DateBegin = valDateBegin
	}
	if _dateCustom := v.Get("date_custom"); _dateCustom != nil {
		b, err := _dateCustom.StringBytes()
		if err != nil {
			return newParsingError("date_custom", err)
		}
		valDateCustom, err := time.Parse("2006.01.02", string(b))
		if err != nil {
			return newParsingError("date_custom", err)
		}
		s.DateCustom = valDateCustom
	}
	if _dateEnd := v.Get("date_end"); valueIsNotNull(_dateEnd) {
		b, err := _dateEnd.StringBytes()
		if err != nil {
			return newParsingError("date_end", err)
		}
		valDateEnd, err := parseDateTime(string(b))
		if err != nil {
			return newParsingError("date_end", err)
		}
		s.DateEnd = new(time.Time)
		*s.DateEnd = time.Time(valDateEnd)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestTime01) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'c', 'u', 's', 't', 'o', 'm'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'e', 'n', 'd'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestTime2) FillFromJSON(v *fastjson.Value) (err error) {
	b, err := v.StringBytes()
	if err != nil {
		return newParsingError("", err)
	}
	_val, err := parseDateTime(string(b))
	if err != nil {
		return err
	}
	*s = TestTime2(_val)
	return nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestTime01) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestTime01) MarshalTo(result *jwriter.Writer) error {
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
	if !s.DateBegin.IsZero() {
		result.RawString(`"date_begin":`)
		writeTime(result, s.DateBegin, time.RFC3339Nano)
		wantComma = true
	} else {
		result.RawString(`"date_begin":"0001-01-01T00:00:00Z"`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if !s.DateCustom.IsZero() {
		result.RawString(`"date_custom":`)
		writeTime(result, s.DateCustom, "2006.01.02")
		wantComma = true
	} else {
		result.RawString(`"date_custom":"0001.01.01"`)
		wantComma = true
	}
	if s.DateEnd != nil {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"date_end":`)
		writeTime(result, *s.DateEnd, time.RFC3339Nano)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestTime01) Reset() {
	s.DateBegin = time.Time{}
	s.DateCustom = time.Time{}
	s.DateEnd = nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestTime2) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestTime2) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	_time, err := time.Time(*s).MarshalText()
	if err == nil {
		result.RawByte('"')
		result.Buffer.AppendBytes(_time)
		result.RawByte('"')
	}
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestTime2) IsZero() bool {
	return time.Time(s).IsZero()
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestTime2) Reset() {
	var tmp time.Time
	tmp = time.Time{}
	*s = TestTime2(tmp)
}
