// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_inh

import (
	"bytes"
	"fmt"
	"math"
	"time"

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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestInh01) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _testInh02 := v.Get("injected"); _testInh02 != nil {
		var valTestInh02 TestInh02
		err = valTestInh02.FillFromJSON(_testInh02, objPath+".injected")
		if err != nil {
			return fmt.Errorf("error parsing '%s.injected' value: %w", objPath, err)
		}
		s.TestInh02 = TestInh02(valTestInh02)
	}
	if _int16 := v.Get("int_16"); _int16 != nil {
		var valInt16 int
		valInt16, err = _int16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.int_16' value: %w", objPath, err)
		}
		if valInt16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%s.int_16' value %d exceeds maximum for data type int16", objPath, valInt16)
		}
		s.Int16 = int16(valInt16)
	}
	if _random := v.Get("random"); _random != nil {
		var valRandom int
		valRandom, err = _random.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.random' value: %w", objPath, err)
		}
		s.Random = valRandom
	}
	if _dateBegin := v.Get("date_begin"); _dateBegin != nil {
		b, err := _dateBegin.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.date_begin' value: %w", objPath, err)
		}
		valDateBegin, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.date_begin' value: %w", objPath, err)
		}
		s.DateBegin = valDateBegin
	}
	if _nested1 := v.Get("nested1"); _nested1 != nil {
		var valNested1 TestInh03
		err = valNested1.FillFromJSON(_nested1, objPath+".nested1")
		if err != nil {
			return fmt.Errorf("error parsing '%s.nested1' value: %w", objPath, err)
		}
		s.Nested1 = TestInh03(valNested1)
	}
	if _nested2 := v.Get("nested2"); valueIsNotNull(_nested2) {
		var valNested2 TestInh03
		err = valNested2.FillFromJSON(_nested2, objPath+".nested2")
		if err != nil {
			return fmt.Errorf("error parsing '%s.nested2' value: %w", objPath, err)
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 't', '_', '1', '6'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 'n', 'd', 'o', 'm'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'e', '_', 'b', 'e', 'g', 'i', 'n'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'n', 'e', 's', 't', 'e', 'd', '1'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'n', 'e', 's', 't', 'e', 'd', '2'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestInh02) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _int32 := v.Get("int_32"); _int32 != nil {
		var valInt32 int
		valInt32, err = _int32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.int_32' value: %w", objPath, err)
		}
		if valInt32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%s.int_32' value %d exceeds maximum for data type int32", objPath, valInt32)
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
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
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestInh03) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _int16 := v.Get("int_16"); _int16 != nil {
		var valInt16 int
		valInt16, err = _int16.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.int_16' value: %w", objPath, err)
		}
		if valInt16 > math.MaxInt16 {
			return fmt.Errorf("error parsing '%s.int_16' value %d exceeds maximum for data type int16", objPath, valInt16)
		}
		s.Int16 = int16(valInt16)
	}
	if _random := v.Get("random"); _random != nil {
		var valRandom int
		valRandom, err = _random.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.random' value: %w", objPath, err)
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
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 'n', 'd', 'o', 'm'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// jsonParserTestNested01 used for pooling Parsers for TestNested01 JSONs.
var jsonParserTestNested01 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestNested01) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestNested01.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestNested01.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestNested01) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field32 := v.Get("field_32"); _field32 != nil {
		var valField32 int
		valField32, err = _field32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field_32' value: %w", objPath, err)
		}
		if valField32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%s.field_32' value %d exceeds maximum for data type int32", objPath, valField32)
		}
		s.Field32 = int32(valField32)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestNested01) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', '_', '3', '2'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// jsonParserTestNested02 used for pooling Parsers for TestNested02 JSONs.
var jsonParserTestNested02 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestNested02) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestNested02.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestNested02.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestNested02) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field32 := v.Get("field_32"); _field32 != nil {
		var valField32 int
		valField32, err = _field32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field_32' value: %w", objPath, err)
		}
		if valField32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%s.field_32' value %d exceeds maximum for data type int32", objPath, valField32)
		}
		s.Field32 = int32(valField32)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestNested02) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', '_', '3', '2'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// jsonParserTestNested03 used for pooling Parsers for TestNested03 JSONs.
var jsonParserTestNested03 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestNested03) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestNested03.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestNested03.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestNested03) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	// strict rules
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if _field32 := v.Get("field_32"); _field32 != nil {
		var valField32 int
		valField32, err = _field32.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field_32' value: %w", objPath, err)
		}
		if valField32 > math.MaxInt32 {
			return fmt.Errorf("error parsing '%s.field_32' value %d exceeds maximum for data type int32", objPath, valField32)
		}
		s.Field32 = int32(valField32)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestNested03) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'e', 'l', 'd', '_', '3', '2'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s.%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		err = fmt.Errorf("unexpected field '%s.%s'", objPath, string(key))
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestInh01) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestInh01) MarshalTo(result Writer) error {
	if s == nil {
		writeString(result, "null")
		return nil
	}
	var err error
	result.Write([]byte{'{'})
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if err = s.TestInh02.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Int16 != 0 {
		result.WriteString(`"int_16":`)
		writeInt64(result, int64(s.Int16))
	} else {
		result.WriteString(`"int_16":0`)
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Random != 0 {
		result.WriteString(`"random":`)
		writeInt64(result, int64(s.Random))
	} else {
		result.WriteString(`"random":0`)
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if !s.DateBegin.IsZero() {
		result.WriteString(`"date_begin":`)
		writeTime(result, s.DateBegin, time.RFC3339Nano)
	} else {
		result.WriteString(`"date_begin":"0000-00-00T00:00:00Z"`)
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if err = s.Nested1.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Nested2 != nil {
		if err = s.Nested2.MarshalTo(result); err != nil {
			return fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
		}
	} else {
		result.WriteString(`"nested2":null`)
	}
	result.Write([]byte{'}'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestInh02) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestInh02) MarshalTo(result Writer) error {
	if s == nil {
		writeString(result, "null")
		return nil
	}
	var err error
	result.Write([]byte{'{'})
	if s.Int32 != 0 {
		if result.Len() > 1 {
			result.Write([]byte{','})
		}
		result.WriteString(`"int_32":`)
		writeInt64(result, int64(s.Int32))
	}
	result.Write([]byte{'}'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestInh03) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestInh03) MarshalTo(result Writer) error {
	if s == nil {
		writeString(result, "null")
		return nil
	}
	var err error
	result.Write([]byte{'{'})
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Int16 != 0 {
		result.WriteString(`"int_16":`)
		writeInt64(result, int64(s.Int16))
	} else {
		result.WriteString(`"int_16":0`)
	}
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Random != 0 {
		result.WriteString(`"random":`)
		writeInt64(result, int64(s.Random))
	} else {
		result.WriteString(`"random":0`)
	}
	result.Write([]byte{'}'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestNested01) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestNested01) MarshalTo(result Writer) error {
	if s == nil {
		writeString(result, "null")
		return nil
	}
	var err error
	result.Write([]byte{'{'})
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Field32 != 0 {
		result.WriteString(`"field_32":`)
		writeInt64(result, int64(s.Field32))
	} else {
		result.WriteString(`"field_32":0`)
	}
	result.Write([]byte{'}'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestNested02) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestNested02) MarshalTo(result Writer) error {
	if s == nil {
		writeString(result, "null")
		return nil
	}
	var err error
	result.Write([]byte{'{'})
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Field32 != 0 {
		result.WriteString(`"field_32":`)
		writeInt64(result, int64(s.Field32))
	} else {
		result.WriteString(`"field_32":0`)
	}
	result.Write([]byte{'}'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestNested03) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestNested03) MarshalTo(result Writer) error {
	if s == nil {
		writeString(result, "null")
		return nil
	}
	var err error
	result.Write([]byte{'{'})
	if result.Len() > 1 {
		result.Write([]byte{','})
	}
	if s.Field32 != 0 {
		result.WriteString(`"field_32":`)
		writeInt64(result, int64(s.Field32))
	} else {
		result.WriteString(`"field_32":0`)
	}
	result.Write([]byte{'}'})
	return err
}
