package test_slice

// TestSlice01 tests struct with array fields
//json:json
type TestSlice01 struct {
	Field    []string `json:"strs"`
	FieldRef []*int   `json:"ints"`
}

type (
	// TestSlice02 tests type inherited slice definition
	//json:json
	TestSlice02 []TestSlice03

	// TestSlice03 supplies test for TestSlice02 (no tags!)
	TestSlice03 struct {
		Data int64 `json:"data"`
	}
)
