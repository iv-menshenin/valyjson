// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package race

import (
	"fmt"
	"io"
	"time"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/valyala/fastjson"
)

type BufWriter interface {
	Size() int
	DumpTo(io.Writer) (int, error)
	ReadCloser() (io.ReadCloser, error)
	String(string)
	RawByte(byte)
	RawString(string)
	Base64Bytes([]byte)
	Uint8(uint8)
	Uint16(uint16)
	Uint32(uint32)
	Uint(uint)
	Uint64(uint64)
	Int8(int8)
	Int16(int16)
	Int32(int32)
	Int(int)
	Int64(int64)
	Uint8Str(uint8)
	Uint16Str(uint16)
	Uint32Str(uint32)
	UintStr(uint)
	Uint64Str(uint64)
	UintptrStr(uintptr)
	Int8Str(int8)
	Int16Str(int16)
	Int32Str(int32)
	IntStr(int)
	Int64Str(int64)
	Float32(float32)
	Float32Str(float32)
	Float64(float64)
	Float64Str(float64)
	Bool(bool)
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

func writeTime(w BufWriter, t time.Time, layout string) {
	var buf = make([]byte, 0, 32)
	buf = append(buf, '"')
	buf = t.AppendFormat(buf, layout)
	buf = append(buf, '"')
	w.RawString(*(*string)(unsafe.Pointer(&buf)))
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
		err:  errors.WithStack(err),
	}
}

func (p parsingError) WrapPath(objPath string) error {
	return parsingError{
		path: objPath + "." + p.path,
		err:  p.err,
	}
}

func (p parsingError) Error() string {
	return fmt.Sprintf("error parsing '%s': %s", p.path, p.err.Error())
}

func (p parsingError) Unwrap() error {
	return p.err
}

func unpackObject(data []byte, err error) ([]byte, error) {
	if err != nil {
		return data, err
	}
	l := len(data)
	if l < 2 {
		return data, nil
	}
	if data[0] != '{' || data[l-1] != '}' {
		return data, fmt.Errorf("expected object")
	}
	return data[1 : l-1], nil
}
