package test_inh

import "time"

// TestInh01 tests inheritance
//json:strict
type TestInh01 struct {
	BreakFirst int `json:"breakFirst,omitempty"`
	TestInh02  `json:"injected,omitempty"`
	TestInh03  `json:",inline"`
	DateBegin  time.Time  `json:"date_begin"`
	Nested1    TestInh03  `json:"nested1"`
	Nested2    *TestInh03 `json:"nested2"`
}

// TestInh02 tests inheritance
//json:strict
type TestInh02 struct {
	Int32 int32 `json:"int_32,omitempty"`
}

// TestInh03 tests inheritance
//json:strict
type TestInh03 struct {
	Int16  int16 `json:"int_16"`
	Random int   `json:"random"`
}

// TestNested01 tests inheritance
//json:strict
type TestNested01 struct {
	TestNested02 `json:",inline"`
}

// TestNested02 tests inheritance
//json:strict
type TestNested02 struct {
	TestNested03 `json:",inline"`
}

// TestNested03 tests inheritance
//json:strict
type TestNested03 struct {
	Field32 int32 `json:"field_32"`
}
