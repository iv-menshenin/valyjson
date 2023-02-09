// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_inh

import (
	"bytes"
	"fmt"
	"math"

	"github.com/valyala/fastjson"
)

// jsonParserTestInh01 used for pooling Parsers for TestInh01 JSONs.
var jsonParserTestInh01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestInh01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestInh01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestInh01.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestInh01) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _testInh02 := v.Get("injected"); _testInh02 != nil {
		var valTestInh02 TestInh02
		err = valTestInh02.FillFromJson(_testInh02, objPath+"injected.")
		if err != nil {
			return fmt.Errorf("error parsing '%sinjected' value: %w", objPath, err)
		}
		s.TestInh02 = TestInh02(valTestInh02)
	}
	if _int16 := v.Get("int_16"); _int16 != nil {
		var valInt16 int
		valInt16, err = _int16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_16' value: %w", objPath, err)
		}
		if valInt16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sint_16' value %d exceeds maximum for data type int16", objPath, valInt16)
		}
		s.Int16 = int16(valInt16)
	}
	if _random := v.Get("random"); _random != nil {
		var valRandom int
		valRandom, err = _random.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%srandom' value: %w", objPath, err)
		}
		s.Random = valRandom
	}
	if _dateBegin := v.Get("date_begin"); _dateBegin != nil {
		b, err := _dateBegin.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%sdate_begin' value: %w", objPath, err)
		}
		valDateBegin, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%sdate_begin' value: %w", objPath, err)
		}
		s.DateBegin = valDateBegin
	}
	if _nested1 := v.Get("nested1"); _nested1 != nil {
		var valNested1 TestInh03
		err = valNested1.FillFromJson(_nested1, objPath+"nested1.")
		if err != nil {
			return fmt.Errorf("error parsing '%snested1' value: %w", objPath, err)
		}
		s.Nested1 = TestInh03(valNested1)
	}
	if _nested2 := v.Get("nested2"); valueIsNotNull(_nested2) {
		var valNested2 TestInh03
		err = valNested2.FillFromJson(_nested2, objPath+"nested2.")
		if err != nil {
			return fmt.Errorf("error parsing '%snested2' value: %w", objPath, err)
		}
		s.Nested2 = new(TestInh03)
		*s.Nested2 = TestInh03(valNested2)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestInh01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [6]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 'j', 'e', 'c', 't', 'e', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', '1', '6'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 'n', 'd', 'o', 'm'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'b', 'e', 'g', 'i', 'n'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'n', 'e', 's', 't', 'e', 'd', '1'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'n', 'e', 's', 't', 'e', 'd', '2'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// jsonParserTestInh02 used for pooling Parsers for TestInh02 JSONs.
var jsonParserTestInh02 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestInh02) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestInh02.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestInh02.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestInh02) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _int32 := v.Get("int_32"); _int32 != nil {
		var valInt32 int
		valInt32, err = _int32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_32' value: %w", objPath, err)
		}
		if valInt32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%sint_32' value %d exceeds maximum for data type int32", objPath, valInt32)
		}
		s.Int32 = int32(valInt32)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestInh02) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', '3', '2'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// jsonParserTestInh03 used for pooling Parsers for TestInh03 JSONs.
var jsonParserTestInh03 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestInh03) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestInh03.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestInh03.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestInh03) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _int16 := v.Get("int_16"); _int16 != nil {
		var valInt16 int
		valInt16, err = _int16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sint_16' value: %w", objPath, err)
		}
		if valInt16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sint_16' value %d exceeds maximum for data type int16", objPath, valInt16)
		}
		s.Int16 = int16(valInt16)
	}
	if _random := v.Get("random"); _random != nil {
		var valRandom int
		valRandom, err = _random.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%srandom' value: %w", objPath, err)
		}
		s.Random = valRandom
	}
	return nil
}

// validate checks for correct data structure
func (s *TestInh03) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', '1', '6'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 'n', 'd', 'o', 'm'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}

// jsonParserTestInh04 used for pooling Parsers for TestInh04 JSONs.
var jsonParserTestInh04 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestInh04) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestInh04.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestInh04.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestInh04) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _fooBar := v.Get("foo-bar"); _fooBar != nil {
		var valFooBar int
		valFooBar, err = _fooBar.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sfoo-bar' value: %w", objPath, err)
		}
		if valFooBar > math.MaxInt16 {
			return fmt.Errorf("error parsing '%sfoo-bar' value %d exceeds maximum for data type int16", objPath, valFooBar)
		}
		s.FooBar = int16(valFooBar)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestInh04) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'o', 'o', '-', 'b', 'a', 'r'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	})
	return err
}
