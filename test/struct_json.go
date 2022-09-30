// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

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
		valFilter := filter.String()
		s.Filter = valFilter
	} else {
		return fmt.Errorf("required element '%sfilter' is missing", objPath)
	}
	if limit := v.Get("limit"); limit != nil {
		var valLimit int
		valLimit, err = limit.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
		s.Limit = valLimit
	}
	if offset := v.Get("offset"); offset != nil {
		var valOffset int
		valOffset, err = offset.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%soffset' value: %w", objPath, err)
		}
		s.Offset = valOffset
	} else {
		s.Offset = 100
	}
	if nested := v.Get("nested"); nested != nil {
		var valNested Nested
		err = valNested.FillFromJson(nested, objPath+"nested.")
		if err != nil {
			return fmt.Errorf("error parsing '%snested' value: %w", objPath, err)
		}
		s.Nested = valNested
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
	if list := v.Get("list-i"); list != nil {
		var listA []*fastjson.Value
		listA, err = list.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%slist-i' value: %w", objPath, err)
		}
		valList := make([]int32, 0, len(listA))
		for _, listElem := range listA {
			var elem int
			elem, err = listElem.Int()
			if err != nil {
				break
			}
			valList = append(valList, int32(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%slist-i' value: %w", objPath, err)
		}
		s.List = valList
	}
	if count := v.Get("count"); count != nil {
		var valCount int64
		valCount, err = count.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%scount' value: %w", objPath, err)
		}
		s.Count = &valCount
	}
	if cross := v.Get("cross"); cross != nil {
		var valCross int64
		valCross, err = cross.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%scross' value: %w", objPath, err)
		}
		s.Cross = &valCross
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
		if bytes.Equal(key, []byte{'l', 'i', 's', 't', '-', 'i'}) {
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
		valName := name.String()
		s.Name = valName
	}
	if surname := v.Get("surname"); surname != nil {
		if surname.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", surname.Type())
			return fmt.Errorf("error parsing '%ssurname' value: %w", objPath, err)
		}
		valSurname := surname.String()
		s.Surname = valSurname
	}
	if rate64 := v.Get("rate64"); rate64 != nil {
		var valRate64 float64
		valRate64, err = rate64.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%srate64' value: %w", objPath, err)
		}
		s.Rate64 = valRate64
	} else {
		s.Rate64 = 1
	}
	if rate32 := v.Get("rate32"); rate32 != nil {
		var valRate32 float64
		valRate32, err = rate32.Float64()
		if err != nil {
			return fmt.Errorf("error parsing '%srate32' value: %w", objPath, err)
		}
		s.Rate32 = float32(valRate32)
	} else {
		s.Rate32 = 1
	}
	if height := v.Get("height"); height != nil {
		var valHeight uint
		valHeight, err = height.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sheight' value: %w", objPath, err)
		}
		s.Height = uint32(valHeight)
	}
	if heightref := v.Get("heightRef"); heightref != nil {
		var valHeightRef uint
		valHeightRef, err = heightref.Uint()
		if err != nil {
			return fmt.Errorf("error parsing '%sheightRef' value: %w", objPath, err)
		}
		s.HeightRef = new(uint32)
		*s.HeightRef = uint32(valHeightRef)
	} else {
		var xHeightRef uint32 = 443
		s.HeightRef = &xHeightRef
	}
	if weight := v.Get("weight"); weight != nil {
		var valWeight uint64
		valWeight, err = weight.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sweight' value: %w", objPath, err)
		}
		s.Weight = valWeight
	}
	if weightref := v.Get("weightRef"); weightref != nil {
		var valWeightRef uint64
		valWeightRef, err = weightref.Uint64()
		if err != nil {
			return fmt.Errorf("error parsing '%sweightRef' value: %w", objPath, err)
		}
		s.WeightRef = new(uint64)
		*s.WeightRef = uint64(valWeightRef)
	}
	if bio := v.Get("bio"); bio != nil {
		var valBio Bio
		err = valBio.FillFromJson(bio, objPath+"bio.")
		if err != nil {
			return fmt.Errorf("error parsing '%sbio' value: %w", objPath, err)
		}
		s.Bio = new(Bio)
		*s.Bio = Bio(valBio)
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
		if bytes.Equal(key, []byte{'b', 'i', 'o'}) {
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
// jsonParserBioused for pooling Parsers for Bio JSONs.
var jsonParserBio fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Bio) UnmarshalJSON(data []byte) error {
	parser := jsonParserBio.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserBio.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *Bio) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	// only if there is a strict rules
	if err = s.validate(v, ""); err != nil {
		return err
	}
	if description := v.Get("description"); description != nil {
		if description.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", description.Type())
			return fmt.Errorf("error parsing '%sdescription' value: %w", objPath, err)
		}
		valDescription := description.String()
		s.Description = new(string)
		*s.Description = string(valDescription)
	}
	if changed := v.Get("changed"); changed != nil {
		valChanged, err := time.Parse(time.RFC3339, changed.String())
		if err != nil {
			return fmt.Errorf("error parsing '%schanged' value: %w", objPath, err)
		}
		s.Changed = new(time.Time)
		*s.Changed = time.Time(valChanged)
	}
	if level := v.Get("level"); level != nil {
		var valLevel int
		valLevel, err = level.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slevel' value: %w", objPath, err)
		}
		s.Level = &valLevel
	}
	if name := v.Get("name"); name != nil {
		var valName int
		valName, err = name.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sname' value: %w", objPath, err)
		}
		s.Name = &valName
	}
	return nil
}

// validate checks for correct data structure
func (s *Bio) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 's', 'c', 'r', 'i', 'p', 't', 'i', 'o', 'n'}) {
			return
		}
		if bytes.Equal(key, []byte{'c', 'h', 'a', 'n', 'g', 'e', 'd'}) {
			return
		}
		if bytes.Equal(key, []byte{'l', 'e', 'v', 'e', 'l'}) {
			return
		}
		if bytes.Equal(key, []byte{'n', 'a', 'm', 'e'}) {
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
	b = marshalString(s.Filter, buf[:0])
	result.Write(b)
	result.WriteString("\"limit\":")
	b = strconv.AppendInt(buf[:0], int64(s.Limit), 10)
	result.Write(b)
	result.WriteString("\"offset\":")
	b = strconv.AppendInt(buf[:0], int64(s.Offset), 10)
	result.Write(b)
	result.WriteString("\"nested\":")
	b, err = s.Nested.MarshalAppend(buf[:0])
	if err != nil {
		return nil, err
	}
	result.Write(b)
	result.WriteRune('}')
	return result.Bytes(), err
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
	if s.Count != nil {
		count := *s.Count
		result.WriteString("\"count\":")
		b = strconv.AppendInt(buf[:0], count, 10)
		result.Write(b)
	} else {
		result.WriteString("\"count\":null")
	}
	if s.Cross != nil {
		cross := *s.Cross
		result.WriteString("\"cross\":")
		b = strconv.AppendInt(buf[:0], cross, 10)
		result.Write(b)
	} else {
		result.WriteString("\"cross\":null")
	}
	result.WriteRune('}')
	return result.Bytes(), err
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
	b = marshalString(s.Name, buf[:0])
	result.Write(b)
	result.WriteString("\"surname\":")
	b = marshalString(s.Surname, buf[:0])
	result.Write(b)
	result.WriteString("\"rate64\":")
	b = strconv.AppendFloat(buf[:0], float64(s.Rate64), 'f', -1, 64)
	result.Write(b)
	result.WriteString("\"rate32\":")
	b = strconv.AppendFloat(buf[:0], float64(s.Rate32), 'f', -1, 64)
	result.Write(b)
	result.WriteString("\"height\":")
	b = strconv.AppendUint(buf[:0], uint64(s.Height), 10)
	result.Write(b)
	if s.HeightRef != nil {
		heightref := *s.HeightRef
		result.WriteString("\"heightRef\":")
		b = strconv.AppendUint(buf[:0], uint64(heightref), 10)
		result.Write(b)
	} else {
		result.WriteString("\"heightRef\":443")
	}
	result.WriteString("\"weight\":")
	b = strconv.AppendUint(buf[:0], s.Weight, 10)
	result.Write(b)
	if s.WeightRef != nil {
		weightref := *s.WeightRef
		result.WriteString("\"weightRef\":")
		b = strconv.AppendUint(buf[:0], weightref, 10)
		result.Write(b)
	}
	if s.Bio != nil {
		bio := *s.Bio
		result.WriteString("\"bio\":")
		b, err = bio.MarshalAppend(buf[:0])
		if err != nil {
			return nil, err
		}
		result.Write(b)
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// MarshalJSON serializes the structure with all its values into JSON format
func (s *Bio) MarshalJSON() ([]byte, error) {
	var buf [128]byte
	return s.MarshalAppend(buf[:0])
}

// MarshalAppend serializes all fields of the structure using a buffer
func (s *Bio) MarshalAppend(dst []byte) ([]byte, error) {
	var result = bytes.NewBuffer(dst)
	var (
		b	[]byte
		buf	[128]byte
		err	error
	)
	result.WriteRune('{')
	if s.Description != nil {
		description := *s.Description
		result.WriteString("\"description\":")
		b = marshalString(description, buf[:0])
		result.Write(b)
	}
	if s.Level != nil {
		level := *s.Level
		result.WriteString("\"level\":")
		b = strconv.AppendInt(buf[:0], int64(level), 10)
		result.Write(b)
	}
	if s.Name != nil {
		name := *s.Name
		result.WriteString("\"name\":")
		b = strconv.AppendInt(buf[:0], int64(name), 10)
		result.Write(b)
	} else {
		result.WriteString("\"name\":null")
	}
	result.WriteRune('}')
	return result.Bytes(), err
}

// valueIsNotNull allows you to determine if the value is contained in a Json structure.
// Checks if the structure itself or the value is Null.
func valueIsNotNull(v *fastjson.Value) bool {
	return v != nil && v.Type() != fastjson.TypeNull
}
