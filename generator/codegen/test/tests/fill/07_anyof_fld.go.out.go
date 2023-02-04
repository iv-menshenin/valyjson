// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package testo

import (
	"bytes"
	"fmt"

	"github.com/valyala/fastjson"
)

// jsonParserTestAllOfSecondused for pooling Parsers for TestAllOfSecond JSONs.
var jsonParserTestAllOfSecond fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestAllOfSecond) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestAllOfSecond.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestAllOfSecond.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestAllOfSecond) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if comment := v.Get("comment"); comment != nil {
		var valComment []byte
		if valComment, err = comment.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%scomment' value: %w", objPath, err)
		}
		s.Comment = string(valComment)
	}
	if level := v.Get("level"); level != nil {
		var valLevel int64
		valLevel, err = level.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%slevel' value: %w", objPath, err)
		}
		s.Level = valLevel
	}
	return nil
}

// validate checks for correct data structure
func (s *TestAllOfSecond) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'm', 'm', 'e', 'n', 't'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'l', 'e', 'v', 'e', 'l'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}

// jsonParserTestAllOfThirdused for pooling Parsers for TestAllOfThird JSONs.
var jsonParserTestAllOfThird fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestAllOfThird) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestAllOfThird.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestAllOfThird.Put(parser)
	return s.FillFromJson(v, "")
}

// FillFromJson recursively fills the fields with fastjson.Value
func (s *TestAllOfThird) FillFromJson(v *fastjson.Value, objPath string) (err error) {
	if err = s.validate(v, objPath); err != nil {
		return err
	}
	if command := v.Get("command"); command != nil {
		var valCommand []byte
		if valCommand, err = command.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%scommand' value: %w", objPath, err)
		}
		s.Command = string(valCommand)
	}
	if _range := v.Get("range"); _range != nil {
		var valRange int64
		valRange, err = _range.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%srange' value: %w", objPath, err)
		}
		s.Range = valRange
	}
	return nil
}

// validate checks for correct data structure
func (s *TestAllOfThird) validate(v *fastjson.Value, objPath string) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'm', 'm', 'a', 'n', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'a', 'n', 'g', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = fmt.Errorf("the '%s%s' field appears in the object twice", objPath, string(key))
			}
			return
		}
	})
	return err
}
