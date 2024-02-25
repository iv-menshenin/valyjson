// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_bool

import (
	"bytes"
	"fmt"

	"github.com/mailru/easyjson/jwriter"
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestBool01) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _bool := v.Get("bl"); _bool != nil {
		var valBool bool
		valBool, err = _bool.Bool()
		if err != nil {
			return newParsingError("bl", err)
		}
		s.Bool = valBool
	}
	if _blMaybe := v.Get("mb"); _blMaybe != nil {
		var valBlMaybe bool
		valBlMaybe, err = _blMaybe.Bool()
		if err != nil {
			return newParsingError("mb", err)
		}
		s.BlMaybe = valBlMaybe
	}
	if _refBool := v.Get("refBool"); valueIsNotNull(_refBool) {
		var valRefBool bool
		valRefBool, err = _refBool.Bool()
		if err != nil {
			return newParsingError("refBool", err)
		}
		s.RefBool = &valRefBool
	}
	if _refMaybe := v.Get("refMaybe"); valueIsNotNull(_refMaybe) {
		var valRefMaybe bool
		valRefMaybe, err = _refMaybe.Bool()
		if err != nil {
			return newParsingError("refMaybe", err)
		}
		s.RefMaybe = &valRefMaybe
	}
	if _defBool := v.Get("defBool"); _defBool != nil {
		var valDefBool bool
		valDefBool, err = _defBool.Bool()
		if err != nil {
			return newParsingError("defBool", err)
		}
		s.DefBool = valDefBool
	} else {
		s.DefBool = true
	}
	return nil
}

// validate checks for correct data structure
func (s *TestBool01) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'m', 'b'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', 'B', 'o', 'o', 'l'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', 'M', 'a', 'y', 'b', 'e'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 'f', 'B', 'o', 'o', 'l'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestBool02) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _i := v.Get("i"); _i != nil {
		var valI bool
		valI, err = _i.Bool()
		if err != nil {
			return newParsingError("i", err)
		}
		s.I = TestInhBool(valI)
	}
	if _x := v.Get("x"); _x != nil {
		var valX bool
		valX, err = _x.Bool()
		if err != nil {
			return newParsingError("x", err)
		}
		s.X = TestInhBool(valX)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestBool02) validate(v *fastjson.Value) error {
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
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'x'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
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
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestInhBool) FillFromJSON(v *fastjson.Value) (err error) {
	var _val bool
	_val, err = v.Bool()
	if err != nil {
		return err
	}
	*s = TestInhBool(_val)
	return nil
}

// jsonParserTestBool03 used for pooling Parsers for TestBool03 JSONs.
var jsonParserTestBool03 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestBool03) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestBool03.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestBool03.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestBool03) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _required := v.Get("required"); _required != nil {
		var valRequired bool
		valRequired, err = _required.Bool()
		if err != nil {
			return newParsingError("required", err)
		}
		s.Required = valRequired
	}
	return nil
}

// validate checks for correct data structure
func (s *TestBool03) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'q', 'u', 'i', 'r', 'e', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
	})
	return err
}

// jsonParserTestBool04 used for pooling Parsers for TestBool04 JSONs.
var jsonParserTestBool04 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestBool04) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestBool04.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestBool04.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestBool04) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _required := v.Get("required"); _required != nil {
		var valRequired bool
		valRequired, err = _required.Bool()
		if err != nil {
			return newParsingError("required", err)
		}
		s.Required = valRequired
	} else {
		return newParsingError("required", fmt.Errorf("required element '%s' is missing", "required"))
	}
	return nil
}

// validate checks for correct data structure
func (s *TestBool04) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'q', 'u', 'i', 'r', 'e', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
	})
	return err
}

// jsonParserTestBool05 used for pooling Parsers for TestBool05 JSONs.
var jsonParserTestBool05 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestBool05) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestBool05.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestBool05.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestBool05) FillFromJSON(v *fastjson.Value) (err error) {
	// strict rules
	if err = s.validate(v); err != nil {
		return err
	}
	if _required := v.Get("required"); _required != nil {
		var valRequired bool
		valRequired, err = _required.Bool()
		if err != nil {
			return newParsingError("required", err)
		}
		s.Required = valRequired
	}
	return nil
}

// validate checks for correct data structure
func (s *TestBool05) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'q', 'u', 'i', 'r', 'e', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s'", string(key))
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestBool01) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestBool01) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.Bool {
		result.RawString(`"bl":true`)
		wantComma = true
	} else {
		result.RawString(`"bl":false`)
		wantComma = true
	}
	if s.BlMaybe {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"mb":true`)
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.RefBool != nil {
		if *s.RefBool {
			result.RawString(`"refBool":true`)
		} else {
			result.RawString(`"refBool":false`)
		}
		wantComma = true
	} else {
		result.RawString(`"refBool":null`)
	}
	if s.RefMaybe != nil {
		if wantComma {
			result.RawByte(',')
		}
		if *s.RefMaybe {
			result.RawString(`"refMaybe":true`)
		} else {
			result.RawString(`"refMaybe":false`)
		}
		wantComma = true
	}
	if wantComma {
		result.RawByte(',')
	}
	if s.DefBool {
		result.RawString(`"defBool":true`)
		wantComma = true
	} else {
		result.RawString(`"defBool":false`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestBool01) Reset() {
	s.Bool = false
	s.BlMaybe = false
	s.RefBool = nil
	s.RefMaybe = nil
	s.DefBool = false
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestBool02) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestBool02) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.I {
		result.RawString(`"i":true`)
		wantComma = true
	} else {
		result.RawString(`"i":false`)
		wantComma = true
	}
	if s.X {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"x":true`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
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

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestBool02) Reset() {
	s.I = TestInhBool(false)
	s.X = TestInhBool(false)
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestInhBool) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestInhBool) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	if *s {
		result.RawString("true")
	} else {
		result.RawString("false")
	}
	return nil
}

// IsZero shows whether the object is an empty value.
func (s TestInhBool) IsZero() bool {
	return s == TestInhBool(false)
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestInhBool) Reset() {
	var tmp bool
	tmp = false
	*s = TestInhBool(tmp)
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestBool03) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestBool03) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.Required {
		result.RawString(`"required":true`)
		wantComma = true
	} else {
		result.RawString(`"required":false`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestBool03) IsZero() bool {
	if s.Required != false {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestBool03) Reset() {
	s.Required = false
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestBool04) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestBool04) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if wantComma {
		result.RawByte(',')
	}
	if s.Required {
		result.RawString(`"required":true`)
		wantComma = true
	} else {
		result.RawString(`"required":false`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestBool04) IsZero() bool {
	if s.Required != false {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestBool04) Reset() {
	s.Required = false
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestBool05) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestBool05) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if s.Required {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"required":true`)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestBool05) IsZero() bool {
	if s.Required != false {
		return false
	}
	return true
}

// Reset resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.
func (s *TestBool05) Reset() {
	s.Required = false
}
