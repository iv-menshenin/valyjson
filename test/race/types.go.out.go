// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package race

import (
	"bytes"
	"fmt"

	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/fastjson"
)

// jsonParserTestStruct used for pooling Parsers for TestStruct JSONs.
var jsonParserTestStruct fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestStruct) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestStruct.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestStruct.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestStruct) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _hasCodes := v.Get("has_codes"); _hasCodes != nil {
		var valHasCodes bool
		valHasCodes, err = _hasCodes.Bool()
		if err != nil {
			return newParsingError("has_codes", err)
		}
		s.HasCodes = valHasCodes
	}
	if _referer := v.Get("referer"); _referer != nil {
		var valReferer []byte
		if valReferer, err = _referer.StringBytes(); err != nil {
			return newParsingError("referer", err)
		}
		s.Referer = string(valReferer)
	}
	if _siteID := v.Get("site_id"); _siteID != nil {
		var valSiteID []byte
		if valSiteID, err = _siteID.StringBytes(); err != nil {
			return newParsingError("site_id", err)
		}
		s.SiteID = string(valSiteID)
	}
	if _uRL := v.Get("url"); _uRL != nil {
		var valURL []byte
		if valURL, err = _uRL.StringBytes(); err != nil {
			return newParsingError("url", err)
		}
		s.URL = string(valURL)
	}
	if _userID := v.Get("user_id"); _userID != nil {
		var valUserID []byte
		if valUserID, err = _userID.StringBytes(); err != nil {
			return newParsingError("user_id", err)
		}
		s.UserID = string(valUserID)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestStruct) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [5]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'h', 'a', 's', '_', 'c', 'o', 'd', 'e', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', 'e', 'r', 'e', 'r'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'i', 't', 'e', '_', 'i', 'd'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 'r', 'l'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 's', 'e', 'r', '_', 'i', 'd'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserTestStructTyped used for pooling Parsers for TestStructTyped JSONs.
var jsonParserTestStructTyped fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *TestStructTyped) UnmarshalJSON(data []byte) error {
	parser := jsonParserTestStructTyped.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserTestStructTyped.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *TestStructTyped) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _hasCodes := v.Get("has_codes"); _hasCodes != nil {
		var valHasCodes bool
		valHasCodes, err = _hasCodes.Bool()
		if err != nil {
			return newParsingError("has_codes", err)
		}
		s.HasCodes = valHasCodes
	}
	if _referer := v.Get("referer"); _referer != nil {
		var valReferer []byte
		if valReferer, err = _referer.StringBytes(); err != nil {
			return newParsingError("referer", err)
		}
		s.Referer = string(valReferer)
	}
	if _siteID := v.Get("site_id"); _siteID != nil {
		var valSiteID []byte
		if valSiteID, err = _siteID.StringBytes(); err != nil {
			return newParsingError("site_id", err)
		}
		s.SiteID = string(valSiteID)
	}
	if _uRL := v.Get("url"); _uRL != nil {
		var valURL []byte
		if valURL, err = _uRL.StringBytes(); err != nil {
			return newParsingError("url", err)
		}
		s.URL = URL(valURL)
	}
	if _userID := v.Get("user_id"); _userID != nil {
		var valUserID []byte
		if valUserID, err = _userID.StringBytes(); err != nil {
			return newParsingError("user_id", err)
		}
		s.UserID = string(valUserID)
	}
	return nil
}

// validate checks for correct data structure
func (s *TestStructTyped) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [5]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'h', 'a', 's', '_', 'c', 'o', 'd', 'e', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', 'e', 'r', 'e', 'r'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'i', 't', 'e', '_', 'i', 'd'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 'r', 'l'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 's', 'e', 'r', '_', 'i', 'd'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestStruct) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestStruct) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if s.HasCodes {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"has_codes":true`)
		wantComma = true
	}
	if s.Referer != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"referer":`)
		result.String(s.Referer)
		wantComma = true
	}
	if s.SiteID != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"site_id":`)
		result.String(s.SiteID)
		wantComma = true
	}
	if s.URL != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"url":`)
		result.String(s.URL)
		wantComma = true
	}
	if s.UserID != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"user_id":`)
		result.String(s.UserID)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestStruct) IsZero() bool {
	if s.HasCodes != false {
		return false
	}
	if s.Referer != "" {
		return false
	}
	if s.SiteID != "" {
		return false
	}
	if s.URL != "" {
		return false
	}
	if s.UserID != "" {
		return false
	}
	return true
}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *TestStructTyped) MarshalJSON() ([]byte, error) {
	var result jwriter.Writer
	if err := s.MarshalTo(&result); err != nil {
		return nil, err
	}
	return result.BuildBytes()
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *TestStructTyped) MarshalTo(result *jwriter.Writer) error {
	if s == nil {
		result.RawString("null")
		return nil
	}
	var (
		err       error
		wantComma bool
	)
	result.RawByte('{')
	if s.HasCodes {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"has_codes":true`)
		wantComma = true
	}
	if s.Referer != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"referer":`)
		result.String(s.Referer)
		wantComma = true
	}
	if s.SiteID != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"site_id":`)
		result.String(s.SiteID)
		wantComma = true
	}
	if s.URL != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"url":`)
		result.String(string(s.URL))
		wantComma = true
	}
	if s.UserID != "" {
		if wantComma {
			result.RawByte(',')
		}
		result.RawString(`"user_id":`)
		result.String(s.UserID)
		wantComma = true
	}
	result.RawByte('}')
	err = result.Error
	return err
}

// IsZero shows whether the object is an empty value.
func (s TestStructTyped) IsZero() bool {
	if s.HasCodes != false {
		return false
	}
	if s.Referer != "" {
		return false
	}
	if s.SiteID != "" {
		return false
	}
	if s.URL != "" {
		return false
	}
	if s.UserID != "" {
		return false
	}
	return true
}
