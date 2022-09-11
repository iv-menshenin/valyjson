// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package {{ .Package }}

import (
    "bytes"
)

func marshalString(s string, b []byte) ([]byte, error) {
	var out = bytes.NewBuffer(b)
	if _, err := out.WriteRune('"'); err != nil {
		return nil, err
	}
	for _, r := range s {
		var err error
		switch r {

		case '\t':
			_, err = out.WriteString(`\t`)

		case '\r':
			_, err = out.WriteString(`\r`)

		case '\n':
			_, err = out.WriteString(`\n`)

		case '\\':
			_, err = out.WriteString(`\\`)

		case '"':
			_, err = out.WriteString(`\"`)

		default:
			_, err = out.WriteRune(r)
		}
		if err != nil {
			return nil, err
		}
	}
	if _, err := out.WriteRune('"'); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}