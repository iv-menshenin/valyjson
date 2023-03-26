package test_any

import (
	"bytes"
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
	//json:encode
	TestOneOfInteger int64
	//json:encode
	TestOneOfString string
	//json:encode
	TestOneOfStruct struct {
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

func (t *TestAllOfFirstIsOne) MarshalTo(result *bufWriter) error {
	if t == nil {
		_, err := result.WriteString("null")
		return err
	}
	if t1, ok := t.OneOf.(TestOneOfInteger); ok {
		return t1.MarshalTo(result)
	}
	if t1, ok := t.OneOf.(TestOneOfString); ok {
		return t1.MarshalTo(result)
	}
	if t1, ok := t.OneOf.(TestOneOfStruct); ok {
		return t1.MarshalTo(result)
	}
	return fmt.Errorf("unknown data type %T", t.OneOf)
}

type (
	// TestAllOfSecond tests allOf option
	//json:optional
	TestAllOfSecond struct {
		Comment string `json:"comment"`
		Level   int64  `json:"level,omitempty"`
	}
	// TestAllOfThird tests allOf option
	//json:optional
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

func (t *TestAllOf01) MarshalJSON() ([]byte, error) {
	var err error
	var result = commonBuffer.Get()
	result.WriteString("{")

	var result1 = commonBuffer.Get()
	if err = t.TestAllOfFirstIsOne.MarshalTo(result1); err != nil {
		return nil, err
	}
	result.WriteString(`"value":`)
	_, err = result.Write(result1.Bytes())

	var result2 = commonBuffer.Get()
	if err = t.TestAllOfSecond.MarshalTo(result2); err != nil {
		return nil, err
	}
	if b := result2.Bytes(); len(b) > 2 && !bytes.Equal(b, []byte{'n', 'u', 'l', 'l'}) {
		result.WriteString(",")
		_, err = result.Write(b[1 : len(b)-1])
		if err != nil {
			return nil, err
		}
	}

	var result3 = commonBuffer.Get()
	if err = t.TestAllOfThird.MarshalTo(result3); err != nil {
		return nil, err
	}
	if b := result3.Bytes(); len(b) > 2 && !bytes.Equal(b, []byte{'n', 'u', 'l', 'l'}) {
		result.WriteString(",")
		_, err = result.Write(b[1 : len(b)-1])
		if err != nil {
			return nil, err
		}
	}

	result.WriteString("}")
	return result.Bytes(), nil
}
