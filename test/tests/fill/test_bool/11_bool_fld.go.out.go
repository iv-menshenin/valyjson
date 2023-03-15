// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_bool

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
)

// jsonParserTestBool01 used for pooling Parsers for TestBool01 JSONs.
var jsonParserTestBool01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestBool01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestBool01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestBool01.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestBool01) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _bool := v.Get("bl"); _bool != nil {
		var valBool bool
		valBool, err = _bool.Bool()
		if err != nil {
			return fmt.Errorf("error parsing '%s.bl' value: %w", objPath, err)
		}
		s.Bool = valBool
	}
	if _blMaybe := v.Get("mb"); _blMaybe != nil {
		var valBlMaybe bool
		valBlMaybe, err = _blMaybe.Bool()
		if err != nil {
			return fmt.Errorf("error parsing '%s.mb' value: %w", objPath, err)
		}
		s.BlMaybe = valBlMaybe
	}
	if _refBool := v.Get("refBool"); valueIsNotNull(_refBool) {
		var valRefBool bool
		valRefBool, err = _refBool.Bool()
		if err != nil {
			return fmt.Errorf("error parsing '%s.refBool' value: %w", objPath, err)
		}
		s.RefBool = &valRefBool
	}
	if _refMaybe := v.Get("refMaybe"); valueIsNotNull(_refMaybe) {
		var valRefMaybe bool
		valRefMaybe, err = _refMaybe.Bool()
		if err != nil {
			return fmt.Errorf("error parsing '%s.refMaybe' value: %w", objPath, err)
		}
		s.RefMaybe = &valRefMaybe
	}
	if _defBool := v.Get("defBool"); _defBool != nil {
		var valDefBool bool
		valDefBool, err = _defBool.Bool()
		if err != nil {
			return fmt.Errorf("error parsing '%s.defBool' value: %w", objPath, err)
		}
		s.DefBool = valDefBool
	} else {
		s.DefBool = true
	}
	return nil
}

// validate checks for correct data structure
func (s *TestBool01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [5]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'b', 'l'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'m', 'b'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', 'B', 'o', 'o', 'l'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', 'M', 'a', 'y', 'b', 'e'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 'f', 'B', 'o', 'o', 'l'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestBool01) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestBool01) MarshalTo(result Writer) error {
	if s == nil {
		writeString(result, "null")
		return nil
	}
	var err error
	result.Write([]byte{'{'})
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Bool {
		result.WriteString(`"bl":true`)
	} else {
		result.WriteString(`"bl":false`)
	}
	if s.BlMaybe {
		if result.Len() > 1 {
			result.Write([]byte{','})
		}
		result.WriteString(`"mb":true`)
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.RefBool != nil {
		if *s.RefBool {
			result.WriteString(`"refBool":true`)
		} else {
			result.WriteString(`"refBool":false`)
		}
	} else {
		result.WriteString(`"refBool":null`)
	}
	if s.RefMaybe != nil {
		if result.Len() > 1 {
			result.Write([]byte{','})
		}
		if *s.RefMaybe {
			result.WriteString(`"refMaybe":true`)
		} else {
			result.WriteString(`"refMaybe":false`)
		}
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.DefBool {
		result.WriteString(`"defBool":true`)
	} else {
		result.WriteString(`"defBool":false`)
	}
	result.Write([]byte{'}'})
	return err
}
