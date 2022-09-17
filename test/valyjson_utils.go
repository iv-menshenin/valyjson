// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test

import (
    "bytes"
	"time"
)

func marshalString(s string, b []byte) []byte {
	var out = bytes.NewBuffer(b)
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
