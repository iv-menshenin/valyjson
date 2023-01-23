package testo

// TestSlice01 tests nested fields inheritance
//json:strict
type TestSlice01 struct {
	Field    []string `json:"strs"`
	FieldRef []*int   `json:"ints"`
}
