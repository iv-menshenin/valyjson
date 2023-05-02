// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package benchmark

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/valyala/fastjson"
)

// jsonParserSearchMetadata used for pooling Parsers for SearchMetadata JSONs.
var jsonParserSearchMetadata fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *SearchMetadata) UnmarshalJSON(data []byte) error {
	parser := jsonParserSearchMetadata.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserSearchMetadata.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *SearchMetadata) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _completedIn := v.Get("completed_in"); _completedIn != nil {
		var valCompletedIn float64
		valCompletedIn, err = _completedIn.Float64()
		if err != nil {
			return newParsingError("completed_in", err)
		}
		s.CompletedIn = valCompletedIn
	}
	if _count := v.Get("count"); _count != nil {
		var valCount int
		valCount, err = _count.Int()
		if err != nil {
			return newParsingError("count", err)
		}
		s.Count = valCount
	}
	if _maxID := v.Get("max_id"); _maxID != nil {
		var valMaxID int64
		valMaxID, err = _maxID.Int64()
		if err != nil {
			return newParsingError("max_id", err)
		}
		s.MaxID = valMaxID
	}
	if _maxIDStr := v.Get("max_id_str"); _maxIDStr != nil {
		var valMaxIDStr []byte
		if valMaxIDStr, err = _maxIDStr.StringBytes(); err != nil {
			return newParsingError("max_id_str", err)
		}
		s.MaxIDStr = *(*string)(unsafe.Pointer(&valMaxIDStr))
	}
	if _nextResults := v.Get("next_results"); _nextResults != nil {
		var valNextResults []byte
		if valNextResults, err = _nextResults.StringBytes(); err != nil {
			return newParsingError("next_results", err)
		}
		s.NextResults = *(*string)(unsafe.Pointer(&valNextResults))
	}
	if _query := v.Get("query"); _query != nil {
		var valQuery []byte
		if valQuery, err = _query.StringBytes(); err != nil {
			return newParsingError("query", err)
		}
		s.Query = *(*string)(unsafe.Pointer(&valQuery))
	}
	if _refreshURL := v.Get("refresh_url"); _refreshURL != nil {
		var valRefreshURL []byte
		if valRefreshURL, err = _refreshURL.StringBytes(); err != nil {
			return newParsingError("refresh_url", err)
		}
		s.RefreshURL = *(*string)(unsafe.Pointer(&valRefreshURL))
	}
	if _sinceID := v.Get("since_id"); _sinceID != nil {
		var valSinceID int64
		valSinceID, err = _sinceID.Int64()
		if err != nil {
			return newParsingError("since_id", err)
		}
		s.SinceID = valSinceID
	}
	if _sinceIDStr := v.Get("since_id_str"); _sinceIDStr != nil {
		var valSinceIDStr []byte
		if valSinceIDStr, err = _sinceIDStr.StringBytes(); err != nil {
			return newParsingError("since_id_str", err)
		}
		s.SinceIDStr = *(*string)(unsafe.Pointer(&valSinceIDStr))
	}
	return nil
}

// validate checks for correct data structure
func (s *SearchMetadata) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [9]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'm', 'p', 'l', 'e', 't', 'e', 'd', '_', 'i', 'n'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'u', 'n', 't'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'m', 'a', 'x', '_', 'i', 'd'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'m', 'a', 'x', '_', 'i', 'd', '_', 's', 't', 'r'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'n', 'e', 'x', 't', '_', 'r', 'e', 's', 'u', 'l', 't', 's'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'q', 'u', 'e', 'r', 'y'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 'f', 'r', 'e', 's', 'h', '_', 'u', 'r', 'l'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'i', 'n', 'c', 'e', '_', 'i', 'd'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'i', 'n', 'c', 'e', '_', 'i', 'd', '_', 's', 't', 'r'}) {
			checkFields[8]++
			if checkFields[8] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserHashtag used for pooling Parsers for Hashtag JSONs.
var jsonParserHashtag fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Hashtag) UnmarshalJSON(data []byte) error {
	parser := jsonParserHashtag.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserHashtag.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Hashtag) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _indices := v.Get("indices"); valueIsNotNull(_indices) {
		var listA []*fastjson.Value
		listA, err = _indices.Array()
		if err != nil {
			return newParsingError("indices", err)
		}
		valIndices := s.Indices[:0]
		if l := len(listA); cap(valIndices) < l || (l == 0 && s.Indices == nil) {
			valIndices = make([]int, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem int
			elem, err = listElem.Int()
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valIndices = append(valIndices, int(elem))
		}
		if err != nil {
			return newParsingError("indices", err)
		}
		s.Indices = valIndices
	}
	if _text := v.Get("text"); _text != nil {
		var valText []byte
		if valText, err = _text.StringBytes(); err != nil {
			return newParsingError("text", err)
		}
		s.Text = *(*string)(unsafe.Pointer(&valText))
	}
	return nil
}

// validate checks for correct data structure
func (s *Hashtag) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 'd', 'i', 'c', 'e', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 'x', 't'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserEntities used for pooling Parsers for Entities JSONs.
var jsonParserEntities fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Entities) UnmarshalJSON(data []byte) error {
	parser := jsonParserEntities.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserEntities.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Entities) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _hashtags := v.Get("hashtags"); valueIsNotNull(_hashtags) {
		var listA []*fastjson.Value
		listA, err = _hashtags.Array()
		if err != nil {
			return newParsingError("hashtags", err)
		}
		valHashtags := s.Hashtags[:0]
		if l := len(listA); cap(valHashtags) < l || (l == 0 && s.Hashtags == nil) {
			valHashtags = make([]Hashtag, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem Hashtag
			err = elem.FillFromJSON(listElem)
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valHashtags = append(valHashtags, Hashtag(elem))
		}
		if err != nil {
			return newParsingError("hashtags", err)
		}
		s.Hashtags = valHashtags
	}
	if _urls := v.Get("urls"); valueIsNotNull(_urls) {
		var listA []*fastjson.Value
		listA, err = _urls.Array()
		if err != nil {
			return newParsingError("urls", err)
		}
		valUrls := s.Urls[:0]
		if l := len(listA); cap(valUrls) < l || (l == 0 && s.Urls == nil) {
			valUrls = make([]*string, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			if !valueIsNotNull(listElem) {
				valUrls = append(valUrls, nil)
				continue
			}
			var elem []byte
			if elem, err = listElem.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_elemNum), err)
			}
			newElem := string(elem)
			valUrls = append(valUrls, &newElem)
		}
		if err != nil {
			return newParsingError("urls", err)
		}
		s.Urls = valUrls
	}
	if _userMentions := v.Get("user_mentions"); valueIsNotNull(_userMentions) {
		var listA []*fastjson.Value
		listA, err = _userMentions.Array()
		if err != nil {
			return newParsingError("user_mentions", err)
		}
		valUserMentions := s.UserMentions[:0]
		if l := len(listA); cap(valUserMentions) < l || (l == 0 && s.UserMentions == nil) {
			valUserMentions = make([]*string, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			if !valueIsNotNull(listElem) {
				valUserMentions = append(valUserMentions, nil)
				continue
			}
			var elem []byte
			if elem, err = listElem.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_elemNum), err)
			}
			newElem := string(elem)
			valUserMentions = append(valUserMentions, &newElem)
		}
		if err != nil {
			return newParsingError("user_mentions", err)
		}
		s.UserMentions = valUserMentions
	}
	return nil
}

// validate checks for correct data structure
func (s *Entities) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [3]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'h', 'a', 's', 'h', 't', 'a', 'g', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 'r', 'l', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 's', 'e', 'r', '_', 'm', 'e', 'n', 't', 'i', 'o', 'n', 's'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserUserEntityDescription used for pooling Parsers for UserEntityDescription JSONs.
var jsonParserUserEntityDescription fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *UserEntityDescription) UnmarshalJSON(data []byte) error {
	parser := jsonParserUserEntityDescription.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserUserEntityDescription.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *UserEntityDescription) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _urls := v.Get("urls"); valueIsNotNull(_urls) {
		var listA []*fastjson.Value
		listA, err = _urls.Array()
		if err != nil {
			return newParsingError("urls", err)
		}
		valUrls := s.Urls[:0]
		if l := len(listA); cap(valUrls) < l || (l == 0 && s.Urls == nil) {
			valUrls = make([]*string, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			if !valueIsNotNull(listElem) {
				valUrls = append(valUrls, nil)
				continue
			}
			var elem []byte
			if elem, err = listElem.StringBytes(); err != nil {
				return newParsingError(strconv.Itoa(_elemNum), err)
			}
			newElem := string(elem)
			valUrls = append(valUrls, &newElem)
		}
		if err != nil {
			return newParsingError("urls", err)
		}
		s.Urls = valUrls
	}
	return nil
}

// validate checks for correct data structure
func (s *UserEntityDescription) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'u', 'r', 'l', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserURL used for pooling Parsers for URL JSONs.
var jsonParserURL fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *URL) UnmarshalJSON(data []byte) error {
	parser := jsonParserURL.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserURL.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *URL) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _expandedURL := v.Get("expanded_url"); valueIsNotNull(_expandedURL) {
		var valExpandedURL []byte
		if valExpandedURL, err = _expandedURL.StringBytes(); err != nil {
			return newParsingError("expanded_url", err)
		}
		s.ExpandedURL = (*string)(unsafe.Pointer(&valExpandedURL))
	}
	if _indices := v.Get("indices"); valueIsNotNull(_indices) {
		var listA []*fastjson.Value
		listA, err = _indices.Array()
		if err != nil {
			return newParsingError("indices", err)
		}
		valIndices := s.Indices[:0]
		if l := len(listA); cap(valIndices) < l || (l == 0 && s.Indices == nil) {
			valIndices = make([]int, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem int
			elem, err = listElem.Int()
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valIndices = append(valIndices, int(elem))
		}
		if err != nil {
			return newParsingError("indices", err)
		}
		s.Indices = valIndices
	}
	if _uRL := v.Get("url"); _uRL != nil {
		var valURL []byte
		if valURL, err = _uRL.StringBytes(); err != nil {
			return newParsingError("url", err)
		}
		s.URL = *(*string)(unsafe.Pointer(&valURL))
	}
	return nil
}

// validate checks for correct data structure
func (s *URL) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [3]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'e', 'x', 'p', 'a', 'n', 'd', 'e', 'd', '_', 'u', 'r', 'l'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', 'd', 'i', 'c', 'e', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 'r', 'l'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserUserEntityURL used for pooling Parsers for UserEntityURL JSONs.
var jsonParserUserEntityURL fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *UserEntityURL) UnmarshalJSON(data []byte) error {
	parser := jsonParserUserEntityURL.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserUserEntityURL.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *UserEntityURL) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _urls := v.Get("urls"); valueIsNotNull(_urls) {
		var listA []*fastjson.Value
		listA, err = _urls.Array()
		if err != nil {
			return newParsingError("urls", err)
		}
		valUrls := s.Urls[:0]
		if l := len(listA); cap(valUrls) < l || (l == 0 && s.Urls == nil) {
			valUrls = make([]URL, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem URL
			err = elem.FillFromJSON(listElem)
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valUrls = append(valUrls, URL(elem))
		}
		if err != nil {
			return newParsingError("urls", err)
		}
		s.Urls = valUrls
	}
	return nil
}

// validate checks for correct data structure
func (s *UserEntityURL) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'u', 'r', 'l', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserUserEntities used for pooling Parsers for UserEntities JSONs.
var jsonParserUserEntities fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *UserEntities) UnmarshalJSON(data []byte) error {
	parser := jsonParserUserEntities.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserUserEntities.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *UserEntities) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _description := v.Get("description"); _description != nil {
		var valDescription UserEntityDescription
		err = valDescription.FillFromJSON(_description)
		if err != nil {
			return newParsingError("description", err)
		}
		s.Description = UserEntityDescription(valDescription)
	}
	if _uRL := v.Get("url"); _uRL != nil {
		var valURL UserEntityURL
		err = valURL.FillFromJSON(_uRL)
		if err != nil {
			return newParsingError("url", err)
		}
		s.URL = UserEntityURL(valURL)
	}
	return nil
}

// validate checks for correct data structure
func (s *UserEntities) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 's', 'c', 'r', 'i', 'p', 't', 'i', 'o', 'n'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 'r', 'l'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserUser used for pooling Parsers for User JSONs.
var jsonParserUser fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *User) UnmarshalJSON(data []byte) error {
	parser := jsonParserUser.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserUser.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *User) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _contributorsEnabled := v.Get("contributors_enabled"); _contributorsEnabled != nil {
		var valContributorsEnabled bool
		valContributorsEnabled, err = _contributorsEnabled.Bool()
		if err != nil {
			return newParsingError("contributors_enabled", err)
		}
		s.ContributorsEnabled = valContributorsEnabled
	}
	if _createdAt := v.Get("created_at"); _createdAt != nil {
		var valCreatedAt []byte
		if valCreatedAt, err = _createdAt.StringBytes(); err != nil {
			return newParsingError("created_at", err)
		}
		s.CreatedAt = *(*string)(unsafe.Pointer(&valCreatedAt))
	}
	if _defaultProfile := v.Get("default_profile"); _defaultProfile != nil {
		var valDefaultProfile bool
		valDefaultProfile, err = _defaultProfile.Bool()
		if err != nil {
			return newParsingError("default_profile", err)
		}
		s.DefaultProfile = valDefaultProfile
	}
	if _defaultProfileImage := v.Get("default_profile_image"); _defaultProfileImage != nil {
		var valDefaultProfileImage bool
		valDefaultProfileImage, err = _defaultProfileImage.Bool()
		if err != nil {
			return newParsingError("default_profile_image", err)
		}
		s.DefaultProfileImage = valDefaultProfileImage
	}
	if _description := v.Get("description"); _description != nil {
		var valDescription []byte
		if valDescription, err = _description.StringBytes(); err != nil {
			return newParsingError("description", err)
		}
		s.Description = *(*string)(unsafe.Pointer(&valDescription))
	}
	if _entities := v.Get("entities"); _entities != nil {
		var valEntities UserEntities
		err = valEntities.FillFromJSON(_entities)
		if err != nil {
			return newParsingError("entities", err)
		}
		s.Entities = UserEntities(valEntities)
	}
	if _favouritesCount := v.Get("favourites_count"); _favouritesCount != nil {
		var valFavouritesCount int
		valFavouritesCount, err = _favouritesCount.Int()
		if err != nil {
			return newParsingError("favourites_count", err)
		}
		s.FavouritesCount = valFavouritesCount
	}
	if _followRequestSent := v.Get("follow_request_sent"); valueIsNotNull(_followRequestSent) {
		var valFollowRequestSent []byte
		if valFollowRequestSent, err = _followRequestSent.StringBytes(); err != nil {
			return newParsingError("follow_request_sent", err)
		}
		s.FollowRequestSent = (*string)(unsafe.Pointer(&valFollowRequestSent))
	}
	if _followersCount := v.Get("followers_count"); _followersCount != nil {
		var valFollowersCount int
		valFollowersCount, err = _followersCount.Int()
		if err != nil {
			return newParsingError("followers_count", err)
		}
		s.FollowersCount = valFollowersCount
	}
	if _following := v.Get("following"); valueIsNotNull(_following) {
		var valFollowing []byte
		if valFollowing, err = _following.StringBytes(); err != nil {
			return newParsingError("following", err)
		}
		s.Following = (*string)(unsafe.Pointer(&valFollowing))
	}
	if _friendsCount := v.Get("friends_count"); _friendsCount != nil {
		var valFriendsCount int
		valFriendsCount, err = _friendsCount.Int()
		if err != nil {
			return newParsingError("friends_count", err)
		}
		s.FriendsCount = valFriendsCount
	}
	if _geoEnabled := v.Get("geo_enabled"); _geoEnabled != nil {
		var valGeoEnabled bool
		valGeoEnabled, err = _geoEnabled.Bool()
		if err != nil {
			return newParsingError("geo_enabled", err)
		}
		s.GeoEnabled = valGeoEnabled
	}
	if _iD := v.Get("id"); _iD != nil {
		var valID int
		valID, err = _iD.Int()
		if err != nil {
			return newParsingError("id", err)
		}
		s.ID = valID
	}
	if _iDStr := v.Get("id_str"); _iDStr != nil {
		var valIDStr []byte
		if valIDStr, err = _iDStr.StringBytes(); err != nil {
			return newParsingError("id_str", err)
		}
		s.IDStr = *(*string)(unsafe.Pointer(&valIDStr))
	}
	if _isTranslator := v.Get("is_translator"); _isTranslator != nil {
		var valIsTranslator bool
		valIsTranslator, err = _isTranslator.Bool()
		if err != nil {
			return newParsingError("is_translator", err)
		}
		s.IsTranslator = valIsTranslator
	}
	if _lang := v.Get("lang"); _lang != nil {
		var valLang []byte
		if valLang, err = _lang.StringBytes(); err != nil {
			return newParsingError("lang", err)
		}
		s.Lang = *(*string)(unsafe.Pointer(&valLang))
	}
	if _listedCount := v.Get("listed_count"); _listedCount != nil {
		var valListedCount int
		valListedCount, err = _listedCount.Int()
		if err != nil {
			return newParsingError("listed_count", err)
		}
		s.ListedCount = valListedCount
	}
	if _location := v.Get("location"); _location != nil {
		var valLocation []byte
		if valLocation, err = _location.StringBytes(); err != nil {
			return newParsingError("location", err)
		}
		s.Location = *(*string)(unsafe.Pointer(&valLocation))
	}
	if _name := v.Get("name"); _name != nil {
		var valName []byte
		if valName, err = _name.StringBytes(); err != nil {
			return newParsingError("name", err)
		}
		s.Name = *(*string)(unsafe.Pointer(&valName))
	}
	if _notifications := v.Get("notifications"); valueIsNotNull(_notifications) {
		var valNotifications []byte
		if valNotifications, err = _notifications.StringBytes(); err != nil {
			return newParsingError("notifications", err)
		}
		s.Notifications = (*string)(unsafe.Pointer(&valNotifications))
	}
	if _profileBackgroundColor := v.Get("profile_background_color"); _profileBackgroundColor != nil {
		var valProfileBackgroundColor []byte
		if valProfileBackgroundColor, err = _profileBackgroundColor.StringBytes(); err != nil {
			return newParsingError("profile_background_color", err)
		}
		s.ProfileBackgroundColor = *(*string)(unsafe.Pointer(&valProfileBackgroundColor))
	}
	if _profileBackgroundImageURL := v.Get("profile_background_image_url"); _profileBackgroundImageURL != nil {
		var valProfileBackgroundImageURL []byte
		if valProfileBackgroundImageURL, err = _profileBackgroundImageURL.StringBytes(); err != nil {
			return newParsingError("profile_background_image_url", err)
		}
		s.ProfileBackgroundImageURL = *(*string)(unsafe.Pointer(&valProfileBackgroundImageURL))
	}
	if _profileBackgroundImageURLHTTPS := v.Get("profile_background_image_url_https"); _profileBackgroundImageURLHTTPS != nil {
		var valProfileBackgroundImageURLHTTPS []byte
		if valProfileBackgroundImageURLHTTPS, err = _profileBackgroundImageURLHTTPS.StringBytes(); err != nil {
			return newParsingError("profile_background_image_url_https", err)
		}
		s.ProfileBackgroundImageURLHTTPS = *(*string)(unsafe.Pointer(&valProfileBackgroundImageURLHTTPS))
	}
	if _profileBackgroundTile := v.Get("profile_background_tile"); _profileBackgroundTile != nil {
		var valProfileBackgroundTile bool
		valProfileBackgroundTile, err = _profileBackgroundTile.Bool()
		if err != nil {
			return newParsingError("profile_background_tile", err)
		}
		s.ProfileBackgroundTile = valProfileBackgroundTile
	}
	if _profileImageURL := v.Get("profile_image_url"); _profileImageURL != nil {
		var valProfileImageURL []byte
		if valProfileImageURL, err = _profileImageURL.StringBytes(); err != nil {
			return newParsingError("profile_image_url", err)
		}
		s.ProfileImageURL = *(*string)(unsafe.Pointer(&valProfileImageURL))
	}
	if _profileImageURLHTTPS := v.Get("profile_image_url_https"); _profileImageURLHTTPS != nil {
		var valProfileImageURLHTTPS []byte
		if valProfileImageURLHTTPS, err = _profileImageURLHTTPS.StringBytes(); err != nil {
			return newParsingError("profile_image_url_https", err)
		}
		s.ProfileImageURLHTTPS = *(*string)(unsafe.Pointer(&valProfileImageURLHTTPS))
	}
	if _profileLinkColor := v.Get("profile_link_color"); _profileLinkColor != nil {
		var valProfileLinkColor []byte
		if valProfileLinkColor, err = _profileLinkColor.StringBytes(); err != nil {
			return newParsingError("profile_link_color", err)
		}
		s.ProfileLinkColor = *(*string)(unsafe.Pointer(&valProfileLinkColor))
	}
	if _profileSidebarBorderColor := v.Get("profile_sidebar_border_color"); _profileSidebarBorderColor != nil {
		var valProfileSidebarBorderColor []byte
		if valProfileSidebarBorderColor, err = _profileSidebarBorderColor.StringBytes(); err != nil {
			return newParsingError("profile_sidebar_border_color", err)
		}
		s.ProfileSidebarBorderColor = *(*string)(unsafe.Pointer(&valProfileSidebarBorderColor))
	}
	if _profileSidebarFillColor := v.Get("profile_sidebar_fill_color"); _profileSidebarFillColor != nil {
		var valProfileSidebarFillColor []byte
		if valProfileSidebarFillColor, err = _profileSidebarFillColor.StringBytes(); err != nil {
			return newParsingError("profile_sidebar_fill_color", err)
		}
		s.ProfileSidebarFillColor = *(*string)(unsafe.Pointer(&valProfileSidebarFillColor))
	}
	if _profileTextColor := v.Get("profile_text_color"); _profileTextColor != nil {
		var valProfileTextColor []byte
		if valProfileTextColor, err = _profileTextColor.StringBytes(); err != nil {
			return newParsingError("profile_text_color", err)
		}
		s.ProfileTextColor = *(*string)(unsafe.Pointer(&valProfileTextColor))
	}
	if _profileUseBackgroundImage := v.Get("profile_use_background_image"); _profileUseBackgroundImage != nil {
		var valProfileUseBackgroundImage bool
		valProfileUseBackgroundImage, err = _profileUseBackgroundImage.Bool()
		if err != nil {
			return newParsingError("profile_use_background_image", err)
		}
		s.ProfileUseBackgroundImage = valProfileUseBackgroundImage
	}
	if _protected := v.Get("protected"); _protected != nil {
		var valProtected bool
		valProtected, err = _protected.Bool()
		if err != nil {
			return newParsingError("protected", err)
		}
		s.Protected = valProtected
	}
	if _screenName := v.Get("screen_name"); _screenName != nil {
		var valScreenName []byte
		if valScreenName, err = _screenName.StringBytes(); err != nil {
			return newParsingError("screen_name", err)
		}
		s.ScreenName = *(*string)(unsafe.Pointer(&valScreenName))
	}
	if _showAllInlineMedia := v.Get("show_all_inline_media"); _showAllInlineMedia != nil {
		var valShowAllInlineMedia bool
		valShowAllInlineMedia, err = _showAllInlineMedia.Bool()
		if err != nil {
			return newParsingError("show_all_inline_media", err)
		}
		s.ShowAllInlineMedia = valShowAllInlineMedia
	}
	if _statusesCount := v.Get("statuses_count"); _statusesCount != nil {
		var valStatusesCount int
		valStatusesCount, err = _statusesCount.Int()
		if err != nil {
			return newParsingError("statuses_count", err)
		}
		s.StatusesCount = valStatusesCount
	}
	if _timeZone := v.Get("time_zone"); _timeZone != nil {
		var valTimeZone []byte
		if valTimeZone, err = _timeZone.StringBytes(); err != nil {
			return newParsingError("time_zone", err)
		}
		s.TimeZone = *(*string)(unsafe.Pointer(&valTimeZone))
	}
	if _uRL := v.Get("url"); valueIsNotNull(_uRL) {
		var valURL []byte
		if valURL, err = _uRL.StringBytes(); err != nil {
			return newParsingError("url", err)
		}
		s.URL = (*string)(unsafe.Pointer(&valURL))
	}
	if _utcOffset := v.Get("utc_offset"); _utcOffset != nil {
		var valUtcOffset int
		valUtcOffset, err = _utcOffset.Int()
		if err != nil {
			return newParsingError("utc_offset", err)
		}
		s.UtcOffset = valUtcOffset
	}
	if _verified := v.Get("verified"); _verified != nil {
		var valVerified bool
		valVerified, err = _verified.Bool()
		if err != nil {
			return newParsingError("verified", err)
		}
		s.Verified = valVerified
	}
	return nil
}

// validate checks for correct data structure
func (s *User) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [39]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'n', 't', 'r', 'i', 'b', 'u', 't', 'o', 'r', 's', '_', 'e', 'n', 'a', 'b', 'l', 'e', 'd'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'c', 'r', 'e', 'a', 't', 'e', 'd', '_', 'a', 't'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 'f', 'a', 'u', 'l', 't', '_', 'p', 'r', 'o', 'f', 'i', 'l', 'e'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 'f', 'a', 'u', 'l', 't', '_', 'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'i', 'm', 'a', 'g', 'e'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'d', 'e', 's', 'c', 'r', 'i', 'p', 't', 'i', 'o', 'n'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'e', 'n', 't', 'i', 't', 'i', 'e', 's'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'a', 'v', 'o', 'u', 'r', 'i', 't', 'e', 's', '_', 'c', 'o', 'u', 'n', 't'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'o', 'l', 'l', 'o', 'w', '_', 'r', 'e', 'q', 'u', 'e', 's', 't', '_', 's', 'e', 'n', 't'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'o', 'l', 'l', 'o', 'w', 'e', 'r', 's', '_', 'c', 'o', 'u', 'n', 't'}) {
			checkFields[8]++
			if checkFields[8] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'o', 'l', 'l', 'o', 'w', 'i', 'n', 'g'}) {
			checkFields[9]++
			if checkFields[9] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'r', 'i', 'e', 'n', 'd', 's', '_', 'c', 'o', 'u', 'n', 't'}) {
			checkFields[10]++
			if checkFields[10] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'g', 'e', 'o', '_', 'e', 'n', 'a', 'b', 'l', 'e', 'd'}) {
			checkFields[11]++
			if checkFields[11] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'd'}) {
			checkFields[12]++
			if checkFields[12] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'd', '_', 's', 't', 'r'}) {
			checkFields[13]++
			if checkFields[13] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 's', '_', 't', 'r', 'a', 'n', 's', 'l', 'a', 't', 'o', 'r'}) {
			checkFields[14]++
			if checkFields[14] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'l', 'a', 'n', 'g'}) {
			checkFields[15]++
			if checkFields[15] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'l', 'i', 's', 't', 'e', 'd', '_', 'c', 'o', 'u', 'n', 't'}) {
			checkFields[16]++
			if checkFields[16] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'l', 'o', 'c', 'a', 't', 'i', 'o', 'n'}) {
			checkFields[17]++
			if checkFields[17] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'n', 'a', 'm', 'e'}) {
			checkFields[18]++
			if checkFields[18] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'n', 'o', 't', 'i', 'f', 'i', 'c', 'a', 't', 'i', 'o', 'n', 's'}) {
			checkFields[19]++
			if checkFields[19] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'b', 'a', 'c', 'k', 'g', 'r', 'o', 'u', 'n', 'd', '_', 'c', 'o', 'l', 'o', 'r'}) {
			checkFields[20]++
			if checkFields[20] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'b', 'a', 'c', 'k', 'g', 'r', 'o', 'u', 'n', 'd', '_', 'i', 'm', 'a', 'g', 'e', '_', 'u', 'r', 'l'}) {
			checkFields[21]++
			if checkFields[21] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'b', 'a', 'c', 'k', 'g', 'r', 'o', 'u', 'n', 'd', '_', 'i', 'm', 'a', 'g', 'e', '_', 'u', 'r', 'l', '_', 'h', 't', 't', 'p', 's'}) {
			checkFields[22]++
			if checkFields[22] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'b', 'a', 'c', 'k', 'g', 'r', 'o', 'u', 'n', 'd', '_', 't', 'i', 'l', 'e'}) {
			checkFields[23]++
			if checkFields[23] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'i', 'm', 'a', 'g', 'e', '_', 'u', 'r', 'l'}) {
			checkFields[24]++
			if checkFields[24] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'i', 'm', 'a', 'g', 'e', '_', 'u', 'r', 'l', '_', 'h', 't', 't', 'p', 's'}) {
			checkFields[25]++
			if checkFields[25] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'l', 'i', 'n', 'k', '_', 'c', 'o', 'l', 'o', 'r'}) {
			checkFields[26]++
			if checkFields[26] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 's', 'i', 'd', 'e', 'b', 'a', 'r', '_', 'b', 'o', 'r', 'd', 'e', 'r', '_', 'c', 'o', 'l', 'o', 'r'}) {
			checkFields[27]++
			if checkFields[27] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 's', 'i', 'd', 'e', 'b', 'a', 'r', '_', 'f', 'i', 'l', 'l', '_', 'c', 'o', 'l', 'o', 'r'}) {
			checkFields[28]++
			if checkFields[28] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 't', 'e', 'x', 't', '_', 'c', 'o', 'l', 'o', 'r'}) {
			checkFields[29]++
			if checkFields[29] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 'f', 'i', 'l', 'e', '_', 'u', 's', 'e', '_', 'b', 'a', 'c', 'k', 'g', 'r', 'o', 'u', 'n', 'd', '_', 'i', 'm', 'a', 'g', 'e'}) {
			checkFields[30]++
			if checkFields[30] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'r', 'o', 't', 'e', 'c', 't', 'e', 'd'}) {
			checkFields[31]++
			if checkFields[31] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'c', 'r', 'e', 'e', 'n', '_', 'n', 'a', 'm', 'e'}) {
			checkFields[32]++
			if checkFields[32] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'h', 'o', 'w', '_', 'a', 'l', 'l', '_', 'i', 'n', 'l', 'i', 'n', 'e', '_', 'm', 'e', 'd', 'i', 'a'}) {
			checkFields[33]++
			if checkFields[33] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 't', 'a', 't', 'u', 's', 'e', 's', '_', 'c', 'o', 'u', 'n', 't'}) {
			checkFields[34]++
			if checkFields[34] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'i', 'm', 'e', '_', 'z', 'o', 'n', 'e'}) {
			checkFields[35]++
			if checkFields[35] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 'r', 'l'}) {
			checkFields[36]++
			if checkFields[36] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 't', 'c', '_', 'o', 'f', 'f', 's', 'e', 't'}) {
			checkFields[37]++
			if checkFields[37] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'v', 'e', 'r', 'i', 'f', 'i', 'e', 'd'}) {
			checkFields[38]++
			if checkFields[38] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserStatusMetadata used for pooling Parsers for StatusMetadata JSONs.
var jsonParserStatusMetadata fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *StatusMetadata) UnmarshalJSON(data []byte) error {
	parser := jsonParserStatusMetadata.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserStatusMetadata.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *StatusMetadata) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _isoLanguageCode := v.Get("iso_language_code"); _isoLanguageCode != nil {
		var valIsoLanguageCode []byte
		if valIsoLanguageCode, err = _isoLanguageCode.StringBytes(); err != nil {
			return newParsingError("iso_language_code", err)
		}
		s.IsoLanguageCode = *(*string)(unsafe.Pointer(&valIsoLanguageCode))
	}
	if _resultType := v.Get("result_type"); _resultType != nil {
		var valResultType []byte
		if valResultType, err = _resultType.StringBytes(); err != nil {
			return newParsingError("result_type", err)
		}
		s.ResultType = *(*string)(unsafe.Pointer(&valResultType))
	}
	return nil
}

// validate checks for correct data structure
func (s *StatusMetadata) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'i', 's', 'o', '_', 'l', 'a', 'n', 'g', 'u', 'a', 'g', 'e', '_', 'c', 'o', 'd', 'e'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 's', 'u', 'l', 't', '_', 't', 'y', 'p', 'e'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserStatus used for pooling Parsers for Status JSONs.
var jsonParserStatus fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *Status) UnmarshalJSON(data []byte) error {
	parser := jsonParserStatus.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserStatus.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *Status) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _contributors := v.Get("contributors"); valueIsNotNull(_contributors) {
		var valContributors []byte
		if valContributors, err = _contributors.StringBytes(); err != nil {
			return newParsingError("contributors", err)
		}
		s.Contributors = (*string)(unsafe.Pointer(&valContributors))
	}
	if _coordinates := v.Get("coordinates"); valueIsNotNull(_coordinates) {
		var valCoordinates []byte
		if valCoordinates, err = _coordinates.StringBytes(); err != nil {
			return newParsingError("coordinates", err)
		}
		s.Coordinates = (*string)(unsafe.Pointer(&valCoordinates))
	}
	if _createdAt := v.Get("created_at"); _createdAt != nil {
		var valCreatedAt []byte
		if valCreatedAt, err = _createdAt.StringBytes(); err != nil {
			return newParsingError("created_at", err)
		}
		s.CreatedAt = *(*string)(unsafe.Pointer(&valCreatedAt))
	}
	if _entities := v.Get("entities"); _entities != nil {
		var valEntities Entities
		err = valEntities.FillFromJSON(_entities)
		if err != nil {
			return newParsingError("entities", err)
		}
		s.Entities = Entities(valEntities)
	}
	if _favorited := v.Get("favorited"); _favorited != nil {
		var valFavorited bool
		valFavorited, err = _favorited.Bool()
		if err != nil {
			return newParsingError("favorited", err)
		}
		s.Favorited = valFavorited
	}
	if _geo := v.Get("geo"); valueIsNotNull(_geo) {
		var valGeo []byte
		if valGeo, err = _geo.StringBytes(); err != nil {
			return newParsingError("geo", err)
		}
		s.Geo = (*string)(unsafe.Pointer(&valGeo))
	}
	if _iD := v.Get("id"); _iD != nil {
		var valID int64
		valID, err = _iD.Int64()
		if err != nil {
			return newParsingError("id", err)
		}
		s.ID = valID
	}
	if _iDStr := v.Get("id_str"); _iDStr != nil {
		var valIDStr []byte
		if valIDStr, err = _iDStr.StringBytes(); err != nil {
			return newParsingError("id_str", err)
		}
		s.IDStr = *(*string)(unsafe.Pointer(&valIDStr))
	}
	if _inReplyToScreenName := v.Get("in_reply_to_screen_name"); valueIsNotNull(_inReplyToScreenName) {
		var valInReplyToScreenName []byte
		if valInReplyToScreenName, err = _inReplyToScreenName.StringBytes(); err != nil {
			return newParsingError("in_reply_to_screen_name", err)
		}
		s.InReplyToScreenName = (*string)(unsafe.Pointer(&valInReplyToScreenName))
	}
	if _inReplyToStatusID := v.Get("in_reply_to_status_id"); valueIsNotNull(_inReplyToStatusID) {
		var valInReplyToStatusID []byte
		if valInReplyToStatusID, err = _inReplyToStatusID.StringBytes(); err != nil {
			return newParsingError("in_reply_to_status_id", err)
		}
		s.InReplyToStatusID = (*string)(unsafe.Pointer(&valInReplyToStatusID))
	}
	if _inReplyToStatusIDStr := v.Get("in_reply_to_status_id_str"); valueIsNotNull(_inReplyToStatusIDStr) {
		var valInReplyToStatusIDStr []byte
		if valInReplyToStatusIDStr, err = _inReplyToStatusIDStr.StringBytes(); err != nil {
			return newParsingError("in_reply_to_status_id_str", err)
		}
		s.InReplyToStatusIDStr = (*string)(unsafe.Pointer(&valInReplyToStatusIDStr))
	}
	if _inReplyToUserID := v.Get("in_reply_to_user_id"); valueIsNotNull(_inReplyToUserID) {
		var valInReplyToUserID []byte
		if valInReplyToUserID, err = _inReplyToUserID.StringBytes(); err != nil {
			return newParsingError("in_reply_to_user_id", err)
		}
		s.InReplyToUserID = (*string)(unsafe.Pointer(&valInReplyToUserID))
	}
	if _inReplyToUserIDStr := v.Get("in_reply_to_user_id_str"); valueIsNotNull(_inReplyToUserIDStr) {
		var valInReplyToUserIDStr []byte
		if valInReplyToUserIDStr, err = _inReplyToUserIDStr.StringBytes(); err != nil {
			return newParsingError("in_reply_to_user_id_str", err)
		}
		s.InReplyToUserIDStr = (*string)(unsafe.Pointer(&valInReplyToUserIDStr))
	}
	if _metadata := v.Get("metadata"); _metadata != nil {
		var valMetadata StatusMetadata
		err = valMetadata.FillFromJSON(_metadata)
		if err != nil {
			return newParsingError("metadata", err)
		}
		s.Metadata = StatusMetadata(valMetadata)
	}
	if _place := v.Get("place"); valueIsNotNull(_place) {
		var valPlace []byte
		if valPlace, err = _place.StringBytes(); err != nil {
			return newParsingError("place", err)
		}
		s.Place = (*string)(unsafe.Pointer(&valPlace))
	}
	if _retweetCount := v.Get("retweet_count"); _retweetCount != nil {
		var valRetweetCount int
		valRetweetCount, err = _retweetCount.Int()
		if err != nil {
			return newParsingError("retweet_count", err)
		}
		s.RetweetCount = valRetweetCount
	}
	if _retweeted := v.Get("retweeted"); _retweeted != nil {
		var valRetweeted bool
		valRetweeted, err = _retweeted.Bool()
		if err != nil {
			return newParsingError("retweeted", err)
		}
		s.Retweeted = valRetweeted
	}
	if _source := v.Get("source"); _source != nil {
		var valSource []byte
		if valSource, err = _source.StringBytes(); err != nil {
			return newParsingError("source", err)
		}
		s.Source = *(*string)(unsafe.Pointer(&valSource))
	}
	if _text := v.Get("text"); _text != nil {
		var valText []byte
		if valText, err = _text.StringBytes(); err != nil {
			return newParsingError("text", err)
		}
		s.Text = *(*string)(unsafe.Pointer(&valText))
	}
	if _truncated := v.Get("truncated"); _truncated != nil {
		var valTruncated bool
		valTruncated, err = _truncated.Bool()
		if err != nil {
			return newParsingError("truncated", err)
		}
		s.Truncated = valTruncated
	}
	if _user := v.Get("user"); _user != nil {
		var valUser User
		err = valUser.FillFromJSON(_user)
		if err != nil {
			return newParsingError("user", err)
		}
		s.User = User(valUser)
	}
	return nil
}

// validate checks for correct data structure
func (s *Status) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [21]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'n', 't', 'r', 'i', 'b', 'u', 't', 'o', 'r', 's'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'c', 'o', 'o', 'r', 'd', 'i', 'n', 'a', 't', 'e', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'c', 'r', 'e', 'a', 't', 'e', 'd', '_', 'a', 't'}) {
			checkFields[2]++
			if checkFields[2] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'e', 'n', 't', 'i', 't', 'i', 'e', 's'}) {
			checkFields[3]++
			if checkFields[3] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'f', 'a', 'v', 'o', 'r', 'i', 't', 'e', 'd'}) {
			checkFields[4]++
			if checkFields[4] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'g', 'e', 'o'}) {
			checkFields[5]++
			if checkFields[5] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'd'}) {
			checkFields[6]++
			if checkFields[6] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'd', '_', 's', 't', 'r'}) {
			checkFields[7]++
			if checkFields[7] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', '_', 'r', 'e', 'p', 'l', 'y', '_', 't', 'o', '_', 's', 'c', 'r', 'e', 'e', 'n', '_', 'n', 'a', 'm', 'e'}) {
			checkFields[8]++
			if checkFields[8] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', '_', 'r', 'e', 'p', 'l', 'y', '_', 't', 'o', '_', 's', 't', 'a', 't', 'u', 's', '_', 'i', 'd'}) {
			checkFields[9]++
			if checkFields[9] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', '_', 'r', 'e', 'p', 'l', 'y', '_', 't', 'o', '_', 's', 't', 'a', 't', 'u', 's', '_', 'i', 'd', '_', 's', 't', 'r'}) {
			checkFields[10]++
			if checkFields[10] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', '_', 'r', 'e', 'p', 'l', 'y', '_', 't', 'o', '_', 'u', 's', 'e', 'r', '_', 'i', 'd'}) {
			checkFields[11]++
			if checkFields[11] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'i', 'n', '_', 'r', 'e', 'p', 'l', 'y', '_', 't', 'o', '_', 'u', 's', 'e', 'r', '_', 'i', 'd', '_', 's', 't', 'r'}) {
			checkFields[12]++
			if checkFields[12] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'m', 'e', 't', 'a', 'd', 'a', 't', 'a'}) {
			checkFields[13]++
			if checkFields[13] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'p', 'l', 'a', 'c', 'e'}) {
			checkFields[14]++
			if checkFields[14] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 't', 'w', 'e', 'e', 't', '_', 'c', 'o', 'u', 'n', 't'}) {
			checkFields[15]++
			if checkFields[15] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'r', 'e', 't', 'w', 'e', 'e', 't', 'e', 'd'}) {
			checkFields[16]++
			if checkFields[16] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 'o', 'u', 'r', 'c', 'e'}) {
			checkFields[17]++
			if checkFields[17] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'e', 'x', 't'}) {
			checkFields[18]++
			if checkFields[18] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'t', 'r', 'u', 'n', 'c', 'a', 't', 'e', 'd'}) {
			checkFields[19]++
			if checkFields[19] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'u', 's', 'e', 'r'}) {
			checkFields[20]++
			if checkFields[20] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserLargeStruct used for pooling Parsers for LargeStruct JSONs.
var jsonParserLargeStruct fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *LargeStruct) UnmarshalJSON(data []byte) error {
	parser := jsonParserLargeStruct.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserLargeStruct.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *LargeStruct) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _searchMetadata := v.Get("search_metadata"); _searchMetadata != nil {
		var valSearchMetadata SearchMetadata
		err = valSearchMetadata.FillFromJSON(_searchMetadata)
		if err != nil {
			return newParsingError("search_metadata", err)
		}
		s.SearchMetadata = SearchMetadata(valSearchMetadata)
	}
	if _statuses := v.Get("statuses"); valueIsNotNull(_statuses) {
		var listA []*fastjson.Value
		listA, err = _statuses.Array()
		if err != nil {
			return newParsingError("statuses", err)
		}
		valStatuses := s.Statuses[:0]
		if l := len(listA); cap(valStatuses) < l || (l == 0 && s.Statuses == nil) {
			valStatuses = make([]Status, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem Status
			err = elem.FillFromJSON(listElem)
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valStatuses = append(valStatuses, Status(elem))
		}
		if err != nil {
			return newParsingError("statuses", err)
		}
		s.Statuses = valStatuses
	}
	return nil
}

// validate checks for correct data structure
func (s *LargeStruct) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [2]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'s', 'e', 'a', 'r', 'c', 'h', '_', 'm', 'e', 't', 'a', 'd', 'a', 't', 'a'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
		if bytes.Equal(key, []byte{'s', 't', 'a', 't', 'u', 's', 'e', 's'}) {
			checkFields[1]++
			if checkFields[1] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

// jsonParserXLStruct used for pooling Parsers for XLStruct JSONs.
var jsonParserXLStruct fastjson.ParserPool

// UnmarshalJSON implements json.Unmarshaler
func (s *XLStruct) UnmarshalJSON(data []byte) error {
	parser := jsonParserXLStruct.Get()
	// parses data containing JSON
	v, err := parser.ParseBytes(data)
	if err != nil {
		return err
	}
	defer jsonParserXLStruct.Put(parser)
	return s.FillFromJSON(v)
}

// FillFromJSON recursively fills the fields with fastjson.Value
func (s *XLStruct) FillFromJSON(v *fastjson.Value) (err error) {
	if err = s.validate(v); err != nil {
		return err
	}
	if _data := v.Get("data"); valueIsNotNull(_data) {
		var listA []*fastjson.Value
		listA, err = _data.Array()
		if err != nil {
			return newParsingError("data", err)
		}
		valData := s.Data[:0]
		if l := len(listA); cap(valData) < l || (l == 0 && s.Data == nil) {
			valData = make([]LargeStruct, 0, len(listA))
		}
		for _elemNum, listElem := range listA {
			var elem LargeStruct
			err = elem.FillFromJSON(listElem)
			if err != nil {
				err = newParsingError(strconv.Itoa(_elemNum), err)
				break
			}
			valData = append(valData, LargeStruct(elem))
		}
		if err != nil {
			return newParsingError("data", err)
		}
		s.Data = valData
	}
	return nil
}

// validate checks for correct data structure
func (s *XLStruct) validate(v *fastjson.Value) error {
	o, err := v.Object()
	if err != nil {
		return err
	}
	var checkFields [1]int
	o.Visit(func(key []byte, _ *fastjson.Value) {
		if err != nil {
			return
		}
		if bytes.Equal(key, []byte{'d', 'a', 't', 'a'}) {
			checkFields[0]++
			if checkFields[0] > 1 {
				err = newParsingError(string(key), fmt.Errorf("the '%s' field appears in the object twice", string(key)))
			}
			return
		}
	})
	return err
}

var bufDataSearchMetadata = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *SearchMetadata) MarshalJSON() ([]byte, error) {
	var result = bufDataSearchMetadata.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *SearchMetadata) MarshalTo(result Writer) error {
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
	if s.CompletedIn != 0 {
		result.WriteString(`"completed_in":`)
		writeFloat64(result, s.CompletedIn)
		wantComma = true
	} else {
		result.WriteString(`"completed_in":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Count != 0 {
		result.WriteString(`"count":`)
		writeInt64(result, int64(s.Count))
		wantComma = true
	} else {
		result.WriteString(`"count":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.MaxID != 0 {
		result.WriteString(`"max_id":`)
		writeInt64(result, s.MaxID)
		wantComma = true
	} else {
		result.WriteString(`"max_id":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.MaxIDStr != "" {
		result.WriteString(`"max_id_str":`)
		writeString(result, s.MaxIDStr)
		wantComma = true
	} else {
		result.WriteString(`"max_id_str":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.NextResults != "" {
		result.WriteString(`"next_results":`)
		writeString(result, s.NextResults)
		wantComma = true
	} else {
		result.WriteString(`"next_results":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Query != "" {
		result.WriteString(`"query":`)
		writeString(result, s.Query)
		wantComma = true
	} else {
		result.WriteString(`"query":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.RefreshURL != "" {
		result.WriteString(`"refresh_url":`)
		writeString(result, s.RefreshURL)
		wantComma = true
	} else {
		result.WriteString(`"refresh_url":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.SinceID != 0 {
		result.WriteString(`"since_id":`)
		writeInt64(result, s.SinceID)
		wantComma = true
	} else {
		result.WriteString(`"since_id":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.SinceIDStr != "" {
		result.WriteString(`"since_id_str":`)
		writeString(result, s.SinceIDStr)
		wantComma = true
	} else {
		result.WriteString(`"since_id_str":""`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s SearchMetadata) IsZero() bool {
	if s.CompletedIn != 0 {
		return false
	}
	if s.Count != 0 {
		return false
	}
	if s.MaxID != 0 {
		return false
	}
	if s.MaxIDStr != "" {
		return false
	}
	if s.NextResults != "" {
		return false
	}
	if s.Query != "" {
		return false
	}
	if s.RefreshURL != "" {
		return false
	}
	if s.SinceID != 0 {
		return false
	}
	if s.SinceIDStr != "" {
		return false
	}
	return true
}

var bufDataHashtag = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Hashtag) MarshalJSON() ([]byte, error) {
	var result = bufDataHashtag.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Hashtag) MarshalTo(result Writer) error {
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
	if s.Indices != nil {
		wantComma = true
		result.WriteString(`"indices":[`)
		var wantComma bool
		for _k, _v := range s.Indices {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			writeInt64(result, int64(_v))
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"indices":null`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Text != "" {
		result.WriteString(`"text":`)
		writeString(result, s.Text)
		wantComma = true
	} else {
		result.WriteString(`"text":""`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Hashtag) IsZero() bool {
	if s.Indices != nil {
		return false
	}
	if s.Text != "" {
		return false
	}
	return true
}

var bufDataEntities = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Entities) MarshalJSON() ([]byte, error) {
	var result = bufDataEntities.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Entities) MarshalTo(result Writer) error {
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
	if s.Hashtags != nil {
		wantComma = true
		result.WriteString(`"hashtags":[`)
		var wantComma bool
		for _k, _v := range s.Hashtags {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "hashtags" item at position %d: %w`, _k, err)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"hashtags":null`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Urls != nil {
		wantComma = true
		result.WriteString(`"urls":[`)
		var wantComma bool
		for _k, _v := range s.Urls {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			if _v == nil {
				result.WriteString("null")
			} else {
				writeString(result, *_v)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"urls":null`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.UserMentions != nil {
		wantComma = true
		result.WriteString(`"user_mentions":[`)
		var wantComma bool
		for _k, _v := range s.UserMentions {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			if _v == nil {
				result.WriteString("null")
			} else {
				writeString(result, *_v)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"user_mentions":null`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Entities) IsZero() bool {
	if s.Hashtags != nil {
		return false
	}
	if s.Urls != nil {
		return false
	}
	if s.UserMentions != nil {
		return false
	}
	return true
}

var bufDataUserEntityDescription = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *UserEntityDescription) MarshalJSON() ([]byte, error) {
	var result = bufDataUserEntityDescription.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *UserEntityDescription) MarshalTo(result Writer) error {
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
	if s.Urls != nil {
		wantComma = true
		result.WriteString(`"urls":[`)
		var wantComma bool
		for _k, _v := range s.Urls {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			if _v == nil {
				result.WriteString("null")
			} else {
				writeString(result, *_v)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"urls":null`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s UserEntityDescription) IsZero() bool {
	if s.Urls != nil {
		return false
	}
	return true
}

var bufDataURL = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *URL) MarshalJSON() ([]byte, error) {
	var result = bufDataURL.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *URL) MarshalTo(result Writer) error {
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
	if s.ExpandedURL != nil {
		result.WriteString(`"expanded_url":`)
		writeString(result, *s.ExpandedURL)
		wantComma = true
	} else {
		result.WriteString(`"expanded_url":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Indices != nil {
		wantComma = true
		result.WriteString(`"indices":[`)
		var wantComma bool
		for _k, _v := range s.Indices {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			writeInt64(result, int64(_v))
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"indices":null`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.URL != "" {
		result.WriteString(`"url":`)
		writeString(result, s.URL)
		wantComma = true
	} else {
		result.WriteString(`"url":""`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s URL) IsZero() bool {
	if s.ExpandedURL != nil {
		return false
	}
	if s.Indices != nil {
		return false
	}
	if s.URL != "" {
		return false
	}
	return true
}

var bufDataUserEntityURL = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *UserEntityURL) MarshalJSON() ([]byte, error) {
	var result = bufDataUserEntityURL.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *UserEntityURL) MarshalTo(result Writer) error {
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
	if s.Urls != nil {
		wantComma = true
		result.WriteString(`"urls":[`)
		var wantComma bool
		for _k, _v := range s.Urls {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "urls" item at position %d: %w`, _k, err)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"urls":null`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s UserEntityURL) IsZero() bool {
	if s.Urls != nil {
		return false
	}
	return true
}

var bufDataUserEntities = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *UserEntities) MarshalJSON() ([]byte, error) {
	var result = bufDataUserEntities.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *UserEntities) MarshalTo(result Writer) error {
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
	result.WriteString(`"description":`)
	if err = s.Description.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "description" attribute: %w`, err)
	}
	wantComma = true
	if wantComma {
		result.WriteString(",")
	}
	result.WriteString(`"url":`)
	if err = s.URL.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "url" attribute: %w`, err)
	}
	wantComma = true
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s UserEntities) IsZero() bool {
	if !s.Description.IsZero() {
		return false
	}
	if !s.URL.IsZero() {
		return false
	}
	return true
}

var bufDataUser = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *User) MarshalJSON() ([]byte, error) {
	var result = bufDataUser.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *User) MarshalTo(result Writer) error {
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
	if s.ContributorsEnabled {
		result.WriteString(`"contributors_enabled":true`)
		wantComma = true
	} else {
		result.WriteString(`"contributors_enabled":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.CreatedAt != "" {
		result.WriteString(`"created_at":`)
		writeString(result, s.CreatedAt)
		wantComma = true
	} else {
		result.WriteString(`"created_at":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.DefaultProfile {
		result.WriteString(`"default_profile":true`)
		wantComma = true
	} else {
		result.WriteString(`"default_profile":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.DefaultProfileImage {
		result.WriteString(`"default_profile_image":true`)
		wantComma = true
	} else {
		result.WriteString(`"default_profile_image":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Description != "" {
		result.WriteString(`"description":`)
		writeString(result, s.Description)
		wantComma = true
	} else {
		result.WriteString(`"description":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	result.WriteString(`"entities":`)
	if err = s.Entities.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "entities" attribute: %w`, err)
	}
	wantComma = true
	if wantComma {
		result.WriteString(",")
	}
	if s.FavouritesCount != 0 {
		result.WriteString(`"favourites_count":`)
		writeInt64(result, int64(s.FavouritesCount))
		wantComma = true
	} else {
		result.WriteString(`"favourites_count":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.FollowRequestSent != nil {
		result.WriteString(`"follow_request_sent":`)
		writeString(result, *s.FollowRequestSent)
		wantComma = true
	} else {
		result.WriteString(`"follow_request_sent":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.FollowersCount != 0 {
		result.WriteString(`"followers_count":`)
		writeInt64(result, int64(s.FollowersCount))
		wantComma = true
	} else {
		result.WriteString(`"followers_count":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Following != nil {
		result.WriteString(`"following":`)
		writeString(result, *s.Following)
		wantComma = true
	} else {
		result.WriteString(`"following":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.FriendsCount != 0 {
		result.WriteString(`"friends_count":`)
		writeInt64(result, int64(s.FriendsCount))
		wantComma = true
	} else {
		result.WriteString(`"friends_count":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.GeoEnabled {
		result.WriteString(`"geo_enabled":true`)
		wantComma = true
	} else {
		result.WriteString(`"geo_enabled":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ID != 0 {
		result.WriteString(`"id":`)
		writeInt64(result, int64(s.ID))
		wantComma = true
	} else {
		result.WriteString(`"id":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.IDStr != "" {
		result.WriteString(`"id_str":`)
		writeString(result, s.IDStr)
		wantComma = true
	} else {
		result.WriteString(`"id_str":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.IsTranslator {
		result.WriteString(`"is_translator":true`)
		wantComma = true
	} else {
		result.WriteString(`"is_translator":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Lang != "" {
		result.WriteString(`"lang":`)
		writeString(result, s.Lang)
		wantComma = true
	} else {
		result.WriteString(`"lang":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ListedCount != 0 {
		result.WriteString(`"listed_count":`)
		writeInt64(result, int64(s.ListedCount))
		wantComma = true
	} else {
		result.WriteString(`"listed_count":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Location != "" {
		result.WriteString(`"location":`)
		writeString(result, s.Location)
		wantComma = true
	} else {
		result.WriteString(`"location":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Name != "" {
		result.WriteString(`"name":`)
		writeString(result, s.Name)
		wantComma = true
	} else {
		result.WriteString(`"name":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Notifications != nil {
		result.WriteString(`"notifications":`)
		writeString(result, *s.Notifications)
		wantComma = true
	} else {
		result.WriteString(`"notifications":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileBackgroundColor != "" {
		result.WriteString(`"profile_background_color":`)
		writeString(result, s.ProfileBackgroundColor)
		wantComma = true
	} else {
		result.WriteString(`"profile_background_color":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileBackgroundImageURL != "" {
		result.WriteString(`"profile_background_image_url":`)
		writeString(result, s.ProfileBackgroundImageURL)
		wantComma = true
	} else {
		result.WriteString(`"profile_background_image_url":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileBackgroundImageURLHTTPS != "" {
		result.WriteString(`"profile_background_image_url_https":`)
		writeString(result, s.ProfileBackgroundImageURLHTTPS)
		wantComma = true
	} else {
		result.WriteString(`"profile_background_image_url_https":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileBackgroundTile {
		result.WriteString(`"profile_background_tile":true`)
		wantComma = true
	} else {
		result.WriteString(`"profile_background_tile":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileImageURL != "" {
		result.WriteString(`"profile_image_url":`)
		writeString(result, s.ProfileImageURL)
		wantComma = true
	} else {
		result.WriteString(`"profile_image_url":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileImageURLHTTPS != "" {
		result.WriteString(`"profile_image_url_https":`)
		writeString(result, s.ProfileImageURLHTTPS)
		wantComma = true
	} else {
		result.WriteString(`"profile_image_url_https":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileLinkColor != "" {
		result.WriteString(`"profile_link_color":`)
		writeString(result, s.ProfileLinkColor)
		wantComma = true
	} else {
		result.WriteString(`"profile_link_color":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileSidebarBorderColor != "" {
		result.WriteString(`"profile_sidebar_border_color":`)
		writeString(result, s.ProfileSidebarBorderColor)
		wantComma = true
	} else {
		result.WriteString(`"profile_sidebar_border_color":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileSidebarFillColor != "" {
		result.WriteString(`"profile_sidebar_fill_color":`)
		writeString(result, s.ProfileSidebarFillColor)
		wantComma = true
	} else {
		result.WriteString(`"profile_sidebar_fill_color":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileTextColor != "" {
		result.WriteString(`"profile_text_color":`)
		writeString(result, s.ProfileTextColor)
		wantComma = true
	} else {
		result.WriteString(`"profile_text_color":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ProfileUseBackgroundImage {
		result.WriteString(`"profile_use_background_image":true`)
		wantComma = true
	} else {
		result.WriteString(`"profile_use_background_image":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Protected {
		result.WriteString(`"protected":true`)
		wantComma = true
	} else {
		result.WriteString(`"protected":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ScreenName != "" {
		result.WriteString(`"screen_name":`)
		writeString(result, s.ScreenName)
		wantComma = true
	} else {
		result.WriteString(`"screen_name":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ShowAllInlineMedia {
		result.WriteString(`"show_all_inline_media":true`)
		wantComma = true
	} else {
		result.WriteString(`"show_all_inline_media":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.StatusesCount != 0 {
		result.WriteString(`"statuses_count":`)
		writeInt64(result, int64(s.StatusesCount))
		wantComma = true
	} else {
		result.WriteString(`"statuses_count":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.TimeZone != "" {
		result.WriteString(`"time_zone":`)
		writeString(result, s.TimeZone)
		wantComma = true
	} else {
		result.WriteString(`"time_zone":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.URL != nil {
		result.WriteString(`"url":`)
		writeString(result, *s.URL)
		wantComma = true
	} else {
		result.WriteString(`"url":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.UtcOffset != 0 {
		result.WriteString(`"utc_offset":`)
		writeInt64(result, int64(s.UtcOffset))
		wantComma = true
	} else {
		result.WriteString(`"utc_offset":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Verified {
		result.WriteString(`"verified":true`)
		wantComma = true
	} else {
		result.WriteString(`"verified":false`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s User) IsZero() bool {
	if s.ContributorsEnabled != false {
		return false
	}
	if s.CreatedAt != "" {
		return false
	}
	if s.DefaultProfile != false {
		return false
	}
	if s.DefaultProfileImage != false {
		return false
	}
	if s.Description != "" {
		return false
	}
	if !s.Entities.IsZero() {
		return false
	}
	if s.FavouritesCount != 0 {
		return false
	}
	if s.FollowRequestSent != nil {
		return false
	}
	if s.FollowersCount != 0 {
		return false
	}
	if s.Following != nil {
		return false
	}
	if s.FriendsCount != 0 {
		return false
	}
	if s.GeoEnabled != false {
		return false
	}
	if s.ID != 0 {
		return false
	}
	if s.IDStr != "" {
		return false
	}
	if s.IsTranslator != false {
		return false
	}
	if s.Lang != "" {
		return false
	}
	if s.ListedCount != 0 {
		return false
	}
	if s.Location != "" {
		return false
	}
	if s.Name != "" {
		return false
	}
	if s.Notifications != nil {
		return false
	}
	if s.ProfileBackgroundColor != "" {
		return false
	}
	if s.ProfileBackgroundImageURL != "" {
		return false
	}
	if s.ProfileBackgroundImageURLHTTPS != "" {
		return false
	}
	if s.ProfileBackgroundTile != false {
		return false
	}
	if s.ProfileImageURL != "" {
		return false
	}
	if s.ProfileImageURLHTTPS != "" {
		return false
	}
	if s.ProfileLinkColor != "" {
		return false
	}
	if s.ProfileSidebarBorderColor != "" {
		return false
	}
	if s.ProfileSidebarFillColor != "" {
		return false
	}
	if s.ProfileTextColor != "" {
		return false
	}
	if s.ProfileUseBackgroundImage != false {
		return false
	}
	if s.Protected != false {
		return false
	}
	if s.ScreenName != "" {
		return false
	}
	if s.ShowAllInlineMedia != false {
		return false
	}
	if s.StatusesCount != 0 {
		return false
	}
	if s.TimeZone != "" {
		return false
	}
	if s.URL != nil {
		return false
	}
	if s.UtcOffset != 0 {
		return false
	}
	if s.Verified != false {
		return false
	}
	return true
}

var bufDataStatusMetadata = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *StatusMetadata) MarshalJSON() ([]byte, error) {
	var result = bufDataStatusMetadata.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *StatusMetadata) MarshalTo(result Writer) error {
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
	if s.IsoLanguageCode != "" {
		result.WriteString(`"iso_language_code":`)
		writeString(result, s.IsoLanguageCode)
		wantComma = true
	} else {
		result.WriteString(`"iso_language_code":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ResultType != "" {
		result.WriteString(`"result_type":`)
		writeString(result, s.ResultType)
		wantComma = true
	} else {
		result.WriteString(`"result_type":""`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s StatusMetadata) IsZero() bool {
	if s.IsoLanguageCode != "" {
		return false
	}
	if s.ResultType != "" {
		return false
	}
	return true
}

var bufDataStatus = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *Status) MarshalJSON() ([]byte, error) {
	var result = bufDataStatus.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *Status) MarshalTo(result Writer) error {
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
	if s.Contributors != nil {
		result.WriteString(`"contributors":`)
		writeString(result, *s.Contributors)
		wantComma = true
	} else {
		result.WriteString(`"contributors":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Coordinates != nil {
		result.WriteString(`"coordinates":`)
		writeString(result, *s.Coordinates)
		wantComma = true
	} else {
		result.WriteString(`"coordinates":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.CreatedAt != "" {
		result.WriteString(`"created_at":`)
		writeString(result, s.CreatedAt)
		wantComma = true
	} else {
		result.WriteString(`"created_at":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	result.WriteString(`"entities":`)
	if err = s.Entities.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "entities" attribute: %w`, err)
	}
	wantComma = true
	if wantComma {
		result.WriteString(",")
	}
	if s.Favorited {
		result.WriteString(`"favorited":true`)
		wantComma = true
	} else {
		result.WriteString(`"favorited":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Geo != nil {
		result.WriteString(`"geo":`)
		writeString(result, *s.Geo)
		wantComma = true
	} else {
		result.WriteString(`"geo":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.ID != 0 {
		result.WriteString(`"id":`)
		writeInt64(result, s.ID)
		wantComma = true
	} else {
		result.WriteString(`"id":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.IDStr != "" {
		result.WriteString(`"id_str":`)
		writeString(result, s.IDStr)
		wantComma = true
	} else {
		result.WriteString(`"id_str":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.InReplyToScreenName != nil {
		result.WriteString(`"in_reply_to_screen_name":`)
		writeString(result, *s.InReplyToScreenName)
		wantComma = true
	} else {
		result.WriteString(`"in_reply_to_screen_name":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.InReplyToStatusID != nil {
		result.WriteString(`"in_reply_to_status_id":`)
		writeString(result, *s.InReplyToStatusID)
		wantComma = true
	} else {
		result.WriteString(`"in_reply_to_status_id":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.InReplyToStatusIDStr != nil {
		result.WriteString(`"in_reply_to_status_id_str":`)
		writeString(result, *s.InReplyToStatusIDStr)
		wantComma = true
	} else {
		result.WriteString(`"in_reply_to_status_id_str":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.InReplyToUserID != nil {
		result.WriteString(`"in_reply_to_user_id":`)
		writeString(result, *s.InReplyToUserID)
		wantComma = true
	} else {
		result.WriteString(`"in_reply_to_user_id":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.InReplyToUserIDStr != nil {
		result.WriteString(`"in_reply_to_user_id_str":`)
		writeString(result, *s.InReplyToUserIDStr)
		wantComma = true
	} else {
		result.WriteString(`"in_reply_to_user_id_str":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	result.WriteString(`"metadata":`)
	if err = s.Metadata.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "metadata" attribute: %w`, err)
	}
	wantComma = true
	if wantComma {
		result.WriteString(",")
	}
	if s.Place != nil {
		result.WriteString(`"place":`)
		writeString(result, *s.Place)
		wantComma = true
	} else {
		result.WriteString(`"place":null`)
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.RetweetCount != 0 {
		result.WriteString(`"retweet_count":`)
		writeInt64(result, int64(s.RetweetCount))
		wantComma = true
	} else {
		result.WriteString(`"retweet_count":0`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Retweeted {
		result.WriteString(`"retweeted":true`)
		wantComma = true
	} else {
		result.WriteString(`"retweeted":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Source != "" {
		result.WriteString(`"source":`)
		writeString(result, s.Source)
		wantComma = true
	} else {
		result.WriteString(`"source":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Text != "" {
		result.WriteString(`"text":`)
		writeString(result, s.Text)
		wantComma = true
	} else {
		result.WriteString(`"text":""`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	if s.Truncated {
		result.WriteString(`"truncated":true`)
		wantComma = true
	} else {
		result.WriteString(`"truncated":false`)
		wantComma = true
	}
	if wantComma {
		result.WriteString(",")
	}
	result.WriteString(`"user":`)
	if err = s.User.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "user" attribute: %w`, err)
	}
	wantComma = true
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s Status) IsZero() bool {
	if s.Contributors != nil {
		return false
	}
	if s.Coordinates != nil {
		return false
	}
	if s.CreatedAt != "" {
		return false
	}
	if !s.Entities.IsZero() {
		return false
	}
	if s.Favorited != false {
		return false
	}
	if s.Geo != nil {
		return false
	}
	if s.ID != 0 {
		return false
	}
	if s.IDStr != "" {
		return false
	}
	if s.InReplyToScreenName != nil {
		return false
	}
	if s.InReplyToStatusID != nil {
		return false
	}
	if s.InReplyToStatusIDStr != nil {
		return false
	}
	if s.InReplyToUserID != nil {
		return false
	}
	if s.InReplyToUserIDStr != nil {
		return false
	}
	if !s.Metadata.IsZero() {
		return false
	}
	if s.Place != nil {
		return false
	}
	if s.RetweetCount != 0 {
		return false
	}
	if s.Retweeted != false {
		return false
	}
	if s.Source != "" {
		return false
	}
	if s.Text != "" {
		return false
	}
	if s.Truncated != false {
		return false
	}
	if !s.User.IsZero() {
		return false
	}
	return true
}

var bufDataLargeStruct = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *LargeStruct) MarshalJSON() ([]byte, error) {
	var result = bufDataLargeStruct.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *LargeStruct) MarshalTo(result Writer) error {
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
	result.WriteString(`"search_metadata":`)
	if err = s.SearchMetadata.MarshalTo(result); err != nil {
		return fmt.Errorf(`can't marshal "search_metadata" attribute: %w`, err)
	}
	wantComma = true
	if wantComma {
		result.WriteString(",")
	}
	if s.Statuses != nil {
		wantComma = true
		result.WriteString(`"statuses":[`)
		var wantComma bool
		for _k, _v := range s.Statuses {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "statuses" item at position %d: %w`, _k, err)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"statuses":null`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s LargeStruct) IsZero() bool {
	if !s.SearchMetadata.IsZero() {
		return false
	}
	if s.Statuses != nil {
		return false
	}
	return true
}

var bufDataXLStruct = cb{}

// MarshalJSON serializes the structure with all its values into JSON format.
func (s *XLStruct) MarshalJSON() ([]byte, error) {
	var result = bufDataXLStruct.Get()
	err := s.MarshalTo(result)
	return result.Bytes(), err
}

// MarshalTo serializes all fields of the structure using a buffer.
func (s *XLStruct) MarshalTo(result Writer) error {
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
	if s.Data != nil {
		wantComma = true
		result.WriteString(`"data":[`)
		var wantComma bool
		for _k, _v := range s.Data {
			if wantComma {
				result.WriteString(",")
			}
			wantComma = true
			_k = _k
			err = _v.MarshalTo(result)
			if err != nil {
				return fmt.Errorf(`can't marshal "data" item at position %d: %w`, _k, err)
			}
		}
		result.WriteString("]")
	} else {
		result.WriteString(`"data":null`)
		wantComma = true
	}
	result.WriteString("}")
	return err
}

// IsZero shows whether the object is an empty value.
func (s XLStruct) IsZero() bool {
	if s.Data != nil {
		return false
	}
	return true
}
