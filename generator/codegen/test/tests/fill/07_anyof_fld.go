package testo

import (
	"encoding/json"
	"fmt"
)

// TestAnyOf01 tests allOf
//json:strict
type TestAnyOf01 struct {
	TestAnyOfUplevel
	TestAllOfUplevel1
	TestAllOfUplevel2
}

type TestAnyOfUplevel struct {
	Any `json:"value"`
}

type Any interface{}

type (
	TestAnyOfInteger int64
	TestAnyOfString  string
	TestAnyOfStruct  struct {
		Class string  `json:"class"`
		Value float64 `json:"width"`
	}
)

func (t *TestAnyOfUplevel) UnmarshalJSON(data []byte) error {
	var (
		err error
		v1  struct {
			TestAnyOfInteger `json:"value"`
		}
		v2 struct {
			TestAnyOfString `json:"value"`
		}
		v3 struct {
			TestAnyOfStruct `json:"value"`
		}
	)
	err = json.Unmarshal(data, &v1)
	if err == nil {
		t.Any = v1.TestAnyOfInteger
		return nil
	}
	err = json.Unmarshal(data, &v2)
	if err == nil {
		t.Any = v2.TestAnyOfString
		return nil
	}
	err = json.Unmarshal(data, &v3)
	if err == nil {
		t.Any = v3.TestAnyOfStruct
		return nil
	}
	return fmt.Errorf("can't unmarshal '%s' into any of [TestAnyOfInteger, TestAnyOfString]", string(data))
}

func (t *TestAnyOfUplevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Any)
}

type (
	TestAllOfUplevel1 struct {
		Comment string `json:"comment"`
		Level   int64  `json:"level,omitempty"`
	}
	TestAllOfUplevel2 struct {
		Command string `json:"command,omitempty"`
		Range   int64  `json:"range,omitempty"`
	}
)

func (t *TestAnyOf01) UnmarshalJSON(data []byte) (err error) {
	if err = json.Unmarshal(data, &t.TestAnyOfUplevel); err != nil {
		return
	}
	if err = json.Unmarshal(data, &t.TestAllOfUplevel1); err != nil {
		return
	}
	if err = json.Unmarshal(data, &t.TestAllOfUplevel2); err != nil {
		return
	}
	return nil
}
