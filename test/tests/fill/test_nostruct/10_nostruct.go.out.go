// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_nostruct

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/valyala/fastjson"

	"fill/test_extr"
)

// jsonParserTestMap10 used for pooling Parsers for TestMap10 JSONs.
var jsonParserTestMap10 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap10) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap10.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap10.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *TestMap10) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	o, err := v.Object()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make(map[string]int64, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value int64
		value, err = v.Int64()
		if err != nil {
			err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
			return
		}
		(*s)[string(key)] = int64(value)
	})
	return err
}

// jsonParserTestMap11 used for pooling Parsers for TestMap11 JSONs.
var jsonParserTestMap11 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap11) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap11.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap11.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *TestMap11) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	o, err := v.Object()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make(map[string]test_extr.External, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		var value test_extr.External
		err = value.FillFromJSON(v, objPath+".")
		if err != nil {
			err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
			return
		}
		(*s)[string(key)] = test_extr.External(value)
	})
	return err
}

// jsonParserTestMap11Ref used for pooling Parsers for TestMap11Ref JSONs.
var jsonParserTestMap11Ref fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestMap11Ref) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestMap11Ref.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestMap11Ref.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON recursively fills the keys with fastjson.Value
func (s *TestMap11Ref) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	o, err := v.Object()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make(map[string]*test_extr.External, o.Len())
	o.Visit(func(key []byte, v *fastjson.Value) {
		if err != nil {
			return
		}
		if v.Type() == fastjson.TypeNull {
			(*s)[string(key)] = nil
			return
		}
		var value test_extr.External
		err = value.FillFromJSON(v, objPath+".")
		if err != nil {
			err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
			return
		}
		(*s)[string(key)] = (*test_extr.External)(unsafe.Pointer(&value))
	})
	return err
}

// jsonParserTestSlice12 used for pooling Parsers for TestSlice12 JSONs.
var jsonParserTestSlice12 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice12) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice12.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice12.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *TestSlice12) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	a, err := v.Array()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make([]int64, len(a))
	for i, v := range a {
		var value int64
		value, err = v.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%s.[%d]' value: %w", objPath, i, err)
		}
		(*s)[i] = int64(value)
	}
	return nil
}

// jsonParserTestSlice13 used for pooling Parsers for TestSlice13 JSONs.
var jsonParserTestSlice13 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice13) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice13.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice13.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *TestSlice13) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	a, err := v.Array()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	*s = make([]test_extr.External, len(a))
	for i, v := range a {
		var value test_extr.External
		err = value.FillFromJSON(v, objPath+".")
		if err != nil {
			return fmt.Errorf("error parsing '%s.[%d]' value: %w", objPath, i, err)
		}
		(*s)[i] = test_extr.External(value)
	}
	return nil
}

// jsonParserTestSlice14 used for pooling Parsers for TestSlice14 JSONs.
var jsonParserTestSlice14 fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestSlice14) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestSlice14.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestSlice14.Put(parser)
	return s.FillFromJSON(v, "(root)")
}

// FillFromJSON fills the array with the values recognized from fastjson.Value
func (s *TestSlice14) FillFromJSON(v *fastjson.Value, objPath string) (err error) {
	a, err := v.Array()
	if err != nil {
		return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
	}
	if len(*s) != len(a) {
		return fmt.Errorf("error parsing '%s', expected %d elements, got %d", objPath, len(*s), len(a))
	}
	for i, v := range a {
		b, err := v.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.' value: %w", objPath, err)
		}
		value, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.[%d]' value: %w", objPath, i, err)
		}
		(*s)[i] = time.Time(value)
	}
	return nil
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap10) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestMap10) MarshalTo(result Writer) error {
	if s == nil || *s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.Write([]byte{'{'})
	for _k, _v := range *s {
		if wantComma {
			result.Write([]byte{','})
		}
		wantComma = true
		result.Write([]byte{'"'})
		result.WriteString(string(_k))
		result.WriteString(`":`)
		writeInt64(result, int64(_v))
	}
	result.Write([]byte{'}'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap11) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestMap11) MarshalTo(result Writer) error {
	if s == nil || *s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.Write([]byte{'{'})
	for _k, _v := range *s {
		if wantComma {
			result.Write([]byte{','})
		}
		wantComma = true
		result.Write([]byte{'"'})
		result.WriteString(string(_k))
		result.WriteString(`":`)
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "TestMap11" attribute %q: %w`, _k, err)
		}
	}
	result.Write([]byte{'}'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestMap11Ref) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestMap11Ref) MarshalTo(result Writer) error {
	if s == nil || *s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.Write([]byte{'{'})
	for _k, _v := range *s {
		if wantComma {
			result.Write([]byte{','})
		}
		wantComma = true
		result.Write([]byte{'"'})
		result.WriteString(string(_k))
		result.WriteString(`":`)
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "TestMap11Ref" attribute %q: %w`, _k, err)
		}
	}
	result.Write([]byte{'}'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSlice12) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice12) MarshalTo(result Writer) error {
	if s == nil || *s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.Write([]byte{'['})
	for _k, _v := range *s {
		if wantComma {
			result.Write([]byte{','})
		}
		wantComma = true
		_k = _k
		writeInt64(result, int64(_v))
	}
	result.Write([]byte{']'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSlice13) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice13) MarshalTo(result Writer) error {
	if s == nil || *s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.Write([]byte{'['})
	for _k, _v := range *s {
		if wantComma {
			result.Write([]byte{','})
		}
		wantComma = true
		_k = _k
		err = _v.MarshalTo(result)
		if err != nil {
			return fmt.Errorf(`can't marshal "TestSlice13" value at position %d: %w`, _k, err)
		}
	}
	result.Write([]byte{']'})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestSlice14) MarshalJSON() ([]byte, error) {
	var result = commonBuffer.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestSlice14) MarshalTo(result Writer) error {
	if s == nil {
		result.WriteString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.Write([]byte{'['})
	for _k, _v := range *s {
		if wantComma {
			result.Write([]byte{','})
		}
		wantComma = true
		_k = _k
		writeTime(result, _v, time.RFC3339Nano)
	}
	result.Write([]byte{']'})
	return err
}
