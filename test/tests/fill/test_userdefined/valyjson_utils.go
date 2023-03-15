// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_userdefined

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/valyala/bytebufferpool"
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

func valueIsNotNull(v *fastjson.Value) bool {
	return v != nil && v.Type() != fastjson.TypeNull
}

func parseDateTime(s string) (time.Time, error) {
	knownFormats := []string{
		time.RFC3339Nano,
		time.RFC3339,
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
		time.RFC3339Nano,
		time.RFC3339,
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

var intBuf = bytebufferpool.Pool{}

func writeInt64(w io.Writer, i int64) {
	buf := intBuf.Get()
	buf.B = strconv.AppendInt(buf.B[:0], i, 10)
	w.Write(buf.B)
	intBuf.Put(buf)
}

func writeUint64(w io.Writer, i uint64) {
	buf := intBuf.Get()
	buf.B = strconv.AppendUint(buf.B[:0], i, 10)
	w.Write(buf.B)
	intBuf.Put(buf)
}

var fltBuf = bytebufferpool.Pool{}

func writeFloat64(w io.Writer, f float64) {
	buf := fltBuf.Get()
	buf.B = strconv.AppendFloat(buf.B[:0], f, 'f', -1, 64)
	w.Write(buf.B)
	fltBuf.Put(buf)
}

var timeBuf = bytebufferpool.Pool{}

func writeTime(w io.Writer, t time.Time, layout string) {
	buf := timeBuf.Get()
	buf.B = append(buf.B[:0], '"')
	buf.B = t.AppendFormat(buf.B, layout)
	buf.B = append(buf.B, '"')
	w.Write(buf.B)
	timeBuf.Put(buf)
}
