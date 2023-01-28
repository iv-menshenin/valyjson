// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package testo

import (
	"bytes"
	"fmt"
	"time"

	"github.com/valyala/fastjson"
)

// jsonParserTestTime01used for pooling Parsers for TestTime01 JSONs.
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
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestTime01) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if datebegin := v.Get("date_begin"); datebegin != nil {
		b, err := datebegin.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%sdate_begin' value: %w", objPath, err)
		}
		valDateBegin, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%sdate_begin' value: %w", objPath, err)
		}
		s.DateBegin = valDateBegin
	}
	if datecustom := v.Get("date_custom"); datecustom != nil {
		b, err := datecustom.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%sdate_custom' value: %w", objPath, err)
		}
		valDateCustom, err := time.Parse("2006.01.02", string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%sdate_custom' value: %w", objPath, err)
		}
		s.DateCustom = valDateCustom
	}
	if dateend := v.Get("date_end"); valueIsNotNull(dateend) {
		b, err := dateend.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%sdate_end' value: %w", objPath, err)
		}
		valDateEnd, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%sdate_end' value: %w", objPath, err)
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
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'c', 'u', 's', 't', 'o', 'm'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'e', 'n', 'd'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}
