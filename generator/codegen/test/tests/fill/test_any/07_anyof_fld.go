package test_any

import (
	"encoding/json"
	"fmt"
)

// TestAllOf01 tests allOf
//json:custom,decode
type TestAllOf01 struct {
	TestAllOfFirstIsOne
	TestAllOfSecond
	TestAllOfThird
}

// TestAllOfFirstIsOne tests oneOf
//json:custom,decode
type TestAllOfFirstIsOne struct {
	OneOf `json:"value"`
}

// OneOf tests oneOf with
type OneOf interface{}

type (
	TestOneOfInteger int64
	TestOneOfString  string
	TestOneOfStruct  struct {
		Class string  `json:"class"`
		Value float64 `json:"width"`
	}
)

// UnmarshalJSON unmarshalls data to one of known structs
func (t *TestAllOfFirstIsOne) UnmarshalJSON(data []byte) error {
	var (
		err error
		v1  struct {
			TestOneOfInteger `json:"value"`
		}
		v2 struct {
			TestOneOfString `json:"value"`
		}
		v3 struct {
			TestOneOfStruct `json:"value"`
		}
	)
	err = json.Unmarshal(data, &v1)
	if err == nil {
		t.OneOf = v1.TestOneOfInteger
		return nil
	}
	err = json.Unmarshal(data, &v2)
	if err == nil {
		t.OneOf = v2.TestOneOfString
		return nil
	}
	err = json.Unmarshal(data, &v3)
	if err == nil {
		t.OneOf = v3.TestOneOfStruct
		return nil
	}
	return fmt.Errorf("can't unmarshal '%s' into one of [TestOneOfInteger, TestOneOfString, TestOneOfStruct]", string(data))
}

func (t *TestAllOfFirstIsOne) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.OneOf)
}

type (
	// TestAllOfSecond tests allOf option
	//json:optional,decode
	TestAllOfSecond struct {
		Comment string `json:"comment"`
		Level   int64  `json:"level,omitempty"`
	}
	// TestAllOfThird tests allOf option
	//json:optional,decode
	TestAllOfThird struct {
		Command string `json:"command,omitempty"`
		Range   int64  `json:"range,omitempty"`
	}
)

// UnmarshalJSON unmarshalls data to all subtypes (inlined structures)
func (t *TestAllOf01) UnmarshalJSON(data []byte) (err error) {
	if err = t.TestAllOfFirstIsOne.UnmarshalJSON(data); err != nil {
		return
	}
	if err = t.TestAllOfSecond.UnmarshalJSON(data); err != nil {
		return
	}
	if err = t.TestAllOfThird.UnmarshalJSON(data); err != nil {
		return
	}
	return nil
}
