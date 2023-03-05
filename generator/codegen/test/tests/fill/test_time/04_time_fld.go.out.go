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

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestTime01) MarshalJSON() ([]byte, error) {
	var buf [512]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer.
func (s *TestTime01) MarshalAppend(dst []byte) ([]byte, error) {
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
	if !s.DateBegin.IsZero() {
		result.WriteString(`"date_begin":`)
		buf = marshalTime(buf[:0], s.DateBegin, time.RFC3339Nano)
		result.Write(buf)
	} else {
		result.WriteString(`"date_begin":"0000-00-00T00:00:00Z"`)
	}
	if result.Len() > 1 {
		result.WriteRune(',')
	}
	if !s.DateCustom.IsZero() {
		result.WriteString(`"date_custom":`)
		buf = marshalTime(buf[:0], s.DateCustom, time.RFC3339Nano)
		result.Write(buf)
	} else {
		result.WriteString(`"date_custom":"0000-00-00T00:00:00Z"`)
	}
	if s.DateEnd != nil {
		if result.Len() > 1 {
			result.WriteRune(',')
		}
		result.WriteString(`"date_end":`)
		buf = marshalTime(buf[:0], *s.DateEnd, time.RFC3339Nano)
		result.Write(buf)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}
