package test_slice

// TestSlice01 tests nested fields inheritance
//json:json
type TestSlice01 struct {
	Field    []string `json:"strs"`
	FieldRef []*int   `json:"ints"`
}
