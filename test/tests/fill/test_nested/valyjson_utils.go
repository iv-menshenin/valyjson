// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_nested

import (
	"fmt"
	"time"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fastjson"
)

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

func writeInt64(w *jwriter.Writer, i int64) {
	w.Int64(i)
}

func writeUint64(w *jwriter.Writer, i uint64) {
	w.Uint64(i)
}

var fltBuf = bytebufferpool.Pool{}

func writeFloat64(w *jwriter.Writer, f float64) {
	w.Float64(f)
}

var timeBuf = bytebufferpool.Pool{}

func writeTime(w *jwriter.Writer, t time.Time, layout string) {
	buf := timeBuf.Get()
	buf.B = append(buf.B[:0], '"')
	buf.B = t.AppendFormat(buf.B, layout)
	buf.B = append(buf.B, '"')
	w.Buffer.AppendBytes(buf.B)
	timeBuf.Put(buf)
}

type parsingError struct {
	path string
	err  error
}

func newParsingError(objPath string, err error) error {
	if err == nil {
		return nil
	}
	type wrapper interface {
		WrapPath(string) error
	}
	if w, ok := err.(wrapper); ok {
		return w.WrapPath(objPath)
	}
	return parsingError{
		path: objPath,
		err:  err,
	}
}

func (p parsingError) WrapPath(objPath string) error {
	return parsingError{
		path: objPath + "." + p.path,
		err:  p.err,
	}
}

func (p parsingError) Error() string {
	return fmt.Sprintf("error parsing '%s': %+v", p.path, p.err)
}
