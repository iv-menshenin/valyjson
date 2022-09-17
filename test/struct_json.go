// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test

import (
	"bytes"
	"fmt"
	"strconv"
	"github.com/valyala/fastjson"
)
// jsonParserStructused for pooling Parsers for Struct JSONs.
var jsonParserStruct fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Struct) UnmarshalJSON(data []byte) error {
	parser := jsonParserStruct.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserStruct.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// only if there is a strict rules
	if err = s.validate(v, ""); err != nil {
		return err
	}
	if filter := v.Get("filter"); filter != nil {
		if filter.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", filter.Type())
			return fmt.Errorf("error parsing '%sfilter' value: %w", objPath, err)
		}
		xFilter := filter.String()
		s.Filter = xFilter
	} else {
		return fmt.Errorf("required element '%sfilter' is missing", objPath)
	}
	if limit := v.Get("limit"); limit != nil {
		var xLimit int
		xLimit, err = limit.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
		s.Limit = xLimit
	}
	if offset := v.Get("offset"); offset != nil {
		var xOffset int
		xOffset, err = offset.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%soffset' value: %w", objPath, err)
		}
		s.Offset = xOffset
	} else {
		s.Offset = 100
	}
	if nested := v.Get("nested"); nested != nil {
		err = s.Nested.FillFromJson(nested, objPath+"nested.")
		if err != nil {
			return fmt.Errorf("error parsing '%snested' value: %w", objPath, err)
		}
	}
	return nil
}

// validate checks for correct data structure
func (s *Struct) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'f', 'i', 'l', 't', 'e', 'r'}) {
			return
		}
		if bytes.Equal(key, []byte{'l', 'i', 'm', 'i', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'o', 'f', 'f', 's', 'e', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'n', 'e', 's', 't', 'e', 'd'}) {
			return
		}
		if objPath == "" {
			err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
		} else {
			err = fmt.Errorf("unexpected field '%s' in the '%s' path", string(key), objPath)
		}
	})
	return nil
}
// jsonParserNestedused for pooling Parsers for Nested JSONs.
var jsonParserNested fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Nested) UnmarshalJSON(data []byte) error {
	parser := jsonParserNested.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserNested.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Nested) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// only if there is a strict rules
	if err = s.validate(v, ""); err != nil {
		return err
	}
	if list := v.Get("list"); list != nil {
		if err != nil {
			return fmt.Errorf("error parsing '%slist' value: %w", objPath, err)
		}
	}
	if count := v.Get("count"); count != nil {
		var xCount int64
		xCount, err = count.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%scount' value: %w", objPath, err)
		}
		s.Count = &xCount
	}
	if cross := v.Get("cross"); cross != nil {
		var xCross int64
		xCross, err = cross.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%scross' value: %w", objPath, err)
		}
		s.Cross = &xCross
	}
	return nil
}

// validate checks for correct data structure
func (s *Nested) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'l', 'i', 's', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'u', 'n', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'c', 'r', 'o', 's', 's'}) {
			return
		}
		if objPath == "" {
			err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
		} else {
			err = fmt.Errorf("unexpected field '%s' in the '%s' path", string(key), objPath)
		}
	})
	return nil
}
// jsonParserPersonused for pooling Parsers for Person JSONs.
var jsonParserPerson fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Person) UnmarshalJSON(data []byte) error {
	parser := jsonParserPerson.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserPerson.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Person) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if name := v.Get("name"); name != nil {
		if name.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", name.Type())
			return fmt.Errorf("error parsing '%sname' value: %w", objPath, err)
		}
		xName := name.String()
		s.Name = xName
	}
	if surname := v.Get("surname"); surname != nil {
		if surname.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", surname.Type())
			return fmt.Errorf("error parsing '%ssurname' value: %w", objPath, err)
		}
		xSurname := surname.String()
		s.Surname = xSurname
	}
	if rate64 := v.Get("rate64"); rate64 != nil {
		var xRate64 float64
		xRate64, err = rate64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%srate64' value: %w", objPath, err)
		}
		s.Rate64 = xRate64
	} else {
		s.Rate64 = 1
	}
	if rate32 := v.Get("rate32"); rate32 != nil {
		var xRate32 float64
		xRate32, err = rate32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%srate32' value: %w", objPath, err)
		}
		s.Rate32 = float32(xRate32)
	} else {
		s.Rate32 = 1
	}
	if height := v.Get("height"); height != nil {
		var xHeight uint
		xHeight, err = height.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sheight' value: %w", objPath, err)
		}
		s.Height = uint32(xHeight)
	}
	if heightref := v.Get("heightRef"); heightref != nil {
		var xHeightRef uint
		xHeightRef, err = heightref.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sheightRef' value: %w", objPath, err)
		}
		s.HeightRef = new(uint32)
		*s.HeightRef = uint32(xHeightRef)
	} else {
		var xHeightRef uint32 = 443
		s.HeightRef = &xHeightRef
	}
	if weight := v.Get("weight"); weight != nil {
		var xWeight uint64
		xWeight, err = weight.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sweight' value: %w", objPath, err)
		}
		s.Weight = xWeight
	}
	if weightref := v.Get("weightRef"); weightref != nil {
		var xWeightRef uint64
		xWeightRef, err = weightref.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sweightRef' value: %w", objPath, err)
		}
		s.WeightRef = &xWeightRef
	}
	return nil
}

// validate checks for correct data structure
func (s *Person) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'n', 'a', 'm', 'e'}) {
			return
		}
		if bytes.Equal(key, []byte{'s', 'u', 'r', 'n', 'a', 'm', 'e'}) {
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 't', 'e', '6', '4'}) {
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 't', 'e', '3', '2'}) {
			return
		}
		if bytes.Equal(key, []byte{'h', 'e', 'i', 'g', 'h', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'h', 'e', 'i', 'g', 'h', 't', 'R', 'e', 'f'}) {
			return
		}
		if bytes.Equal(key, []byte{'w', 'e', 'i', 'g', 'h', 't'}) {
			return
		}
		if bytes.Equal(key, []byte{'w', 'e', 'i', 'g', 'h', 't', 'R', 'e', 'f'}) {
			return
		}
		if objPath == "" {
			err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
		} else {
			err = fmt.Errorf("unexpected field '%s' in the '%s' path", string(key), objPath)
		}
	})
	return nil
}

// MarshalJSON serializes the structure with all its values into JSON format
func (s *Struct) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer
func (s *Struct) MarshalAppend(dst []byte) ([]byte, error) {
	var result = bytes.NewBuffer(dst)
	var (
		b	[]byte
		buf	[128]byte
		err	error
	)
	result.WriteRune('{')
	result.WriteString("\"filter\":")
	b, err = marshalString(s.Filter, buf[:0])
	if err != nil {
		return nil, err
	}
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"limit\":")
	b = strconv.AppendInt(buf[:0], int64(s.Limit), 10)
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"offset\":")
	b = strconv.AppendInt(buf[:0], int64(s.Offset), 10)
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"nested\":")
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteRune('}')
	return result.Bytes(), nil
}

// MarshalJSON serializes the structure with all its values into JSON format
func (s *Nested) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer
func (s *Nested) MarshalAppend(dst []byte) ([]byte, error) {
	var result = bytes.NewBuffer(dst)
	var (
		b	[]byte
		buf	[128]byte
		err	error
	)
	result.WriteRune('{')
	result.WriteString("\"list\":")
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"count\":")
	if s.Count != nil {
		count := *s.Count
		b = strconv.AppendInt(buf[:0], count, 10)
	} else {
		result.WriteString("\"count\":null")
	}
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"cross\":")
	if s.Cross != nil {
		cross := *s.Cross
		b = strconv.AppendInt(buf[:0], cross, 10)
	} else {
		result.WriteString("\"cross\":null")
	}
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteRune('}')
	return result.Bytes(), nil
}

// MarshalJSON serializes the structure with all its values into JSON format
func (s *Person) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer
func (s *Person) MarshalAppend(dst []byte) ([]byte, error) {
	var result = bytes.NewBuffer(dst)
	var (
		b	[]byte
		buf	[128]byte
		err	error
	)
	result.WriteRune('{')
	result.WriteString("\"name\":")
	b, err = marshalString(s.Name, buf[:0])
	if err != nil {
		return nil, err
	}
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"surname\":")
	b, err = marshalString(s.Surname, buf[:0])
	if err != nil {
		return nil, err
	}
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"rate64\":")
	b = strconv.AppendFloat(buf[:0], float64(s.Rate64), 'f', -1, 64)
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"rate32\":")
	b = strconv.AppendFloat(buf[:0], float64(s.Rate32), 'f', -1, 64)
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"height\":")
	b = strconv.AppendUint(buf[:0], uint64(s.Height), 10)
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"heightRef\":")
	if s.HeightRef != nil {
		heightref := *s.HeightRef
		b = strconv.AppendUint(buf[:0], uint64(heightref), 10)
	} else {
		result.WriteString("\"heightRef\":443")
	}
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"weight\":")
	b = strconv.AppendUint(buf[:0], s.Weight, 10)
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteString("\"weightRef\":")
	if s.WeightRef != nil {
		weightref := *s.WeightRef
		b = strconv.AppendUint(buf[:0], weightref, 10)
	}
	if _, err = result.Write(b); err != nil {
		return nil, err
	}
	result.WriteRune('}')
	return result.Bytes(), nil
}

// valueIsNotNull allows you to determine if the value is contained in a Json structure.
// Checks if the structure itself or the value is Null.
func valueIsNotNull(v *fastjson.Value) bool {
	return v != nil && v.Type() != fastjson.TypeNull
}
