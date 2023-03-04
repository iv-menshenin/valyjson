// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_string

import (
	"bytes"
	"fmt"
	"time"

	"github.com/valyala/fastjson"
)

func marshalString(buf []byte, s string) []byte {
	var out = bytes.NewBuffer(buf)
	out.WriteRune('"')
	for _, r := range s {
		switch r {

		case '\t':
			out.WriteString(`\t`)

		case '\r':
			out.WriteString(`\r`)

		case '\n':
			out.WriteString(`\n`)

		case '\\':
			out.WriteString(`\\`)

		case '"':
			out.WriteString(`\"`)

		default:
			out.WriteRune(r)
		}
	}
	out.WriteRune('"')
	return out.Bytes()
}

func marshalTime(t time.Time, layout string, b []byte) []byte {
	return t.AppendFormat(b, layout)
}

func valueIsNotNull(v *fastjson.Value) bool {
	return v != nil && v.Type() != fastjson.TypeNull
}

func parseDateTime(s string) (time.Time, error) {
	knownFormats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02 15:04:05Z07:00",
		"02.01.2006 15:04:05Z07:00",
	}
	for _, format := range knownFormats {
		if value, err := time.Parse(format, s); err == nil {
			return value, nil
		}
	}
	return time.Time{}, fmt.Errorf("can't parse date-time from string '%s'", s)
}

func parseDate(s string) (time.Time, error) {
	knownFormats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02",
		"02.01.2006",
	}
	for _, format := range knownFormats {
		if value, err := time.Parse(format, s); err == nil {
			return value.Truncate(24 * time.Hour), nil
		}
	}
	return time.Time{}, fmt.Errorf("can't parse date from string '%s'", s)
}
