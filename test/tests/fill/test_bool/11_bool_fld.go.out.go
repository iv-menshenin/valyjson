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

// jsonParserTestBool02 used for pooling Parsers for TestBool02 JSONs.
var jsonParserTestBool02 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestBool02) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestBool02.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestBool02.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestBool02) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _i := v.Get("i"); _i != nil {
		var valI bool
		valI, err = _i.Bool()
		if err != nil {
			return fmt.Errorf("error parsing '%s.i' value: %w", objPath, err)
		}
		s.I = TestInhBool(valI)
	}
	if _x := v.Get("x"); _x != nil {
		var valX bool
		valX, err = _x.Bool()
		if err != nil {
			return fmt.Errorf("error parsing '%s.x' value: %w", objPath, err)
		}
		s.X = TestInhBool(valX)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestBool02) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'i'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'x'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}

// jsonParserTestInhBool used for pooling Parsers for TestInhBool JSONs.
var jsonParserTestInhBool fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestInhBool) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestInhBool.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestInhBool.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestInhBool) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	var _val bool
	_val, err = v.Bool()
	if err != nil {
		return err
	}
	*s = TestInhBool(_val)
	return nil
}

var bufDataTestBool01 = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestBool01) MarshalJSON() ([]byte, error) {
	var result = bufDataTestBool01.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestBool01) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if wantComma {
		result.WriteString(",")
	}
	if s.Bool {
		result.WriteString(`"bl":true`)
		wantComma = true
	} else {
		result.WriteString(`"bl":false`)
		wantComma = true
	}
	if s.BlMaybe {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"mb":true`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.RefBool != nil {
		if *s.RefBool {
			result.WriteString(`"refBool":true`)
		} else {
			result.WriteString(`"refBool":false`)
		}
		wantComma = true
	} else {
		result.WriteString(`"refBool":null`)
	}
	if s.RefMaybe != nil {
		if wantComma {
			result.WriteString(",")
		}
		if *s.RefMaybe {
			result.WriteString(`"refMaybe":true`)
		} else {
			result.WriteString(`"refMaybe":false`)
		}
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.DefBool {
		result.WriteString(`"defBool":true`)
		wantComma = true
	} else {
		result.WriteString(`"defBool":false`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestBool01) IsZero() bool {
	if s.Bool != false {
		return false
	}
	if s.BlMaybe != false {
		return false
	}
	if s.RefBool != nil {
		return false
	}
	if s.RefMaybe != nil {
		return false
	}
	if s.DefBool != false {
		return false
	}
	return true
}

var bufDataTestBool02 = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestBool02) MarshalJSON() ([]byte, error) {
	var result = bufDataTestBool02.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestBool02) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.WriteString("{")
	if wantComma {
		result.WriteString(",")
	}
	if s.I {
		result.WriteString(`"i":true`)
		wantComma = true
	} else {
		result.WriteString(`"i":false`)
		wantComma = true
	}
	if s.X {
		if wantComma {
			result.WriteString(",")
		}
		result.WriteString(`"x":true`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestBool02) IsZero() bool {
	if s.I != false {
		return false
	}
	if s.X != false {
		return false
	}
	return true
}

var bufDataTestInhBool = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestInhBool) MarshalJSON() ([]byte, error) {
	var result = bufDataTestInhBool.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestInhBool) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	if *s {
		result.WriteString("true")
	} else {
		result.WriteString("false")
	}
	return nil
}

// IsZero shows whether the object is an empty value.
func (s TestInhBool) IsZero() bool {
	return s == TestInhBool(false)
}
