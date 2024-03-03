package test_slice

// TestSlice01 tests struct with array fields
//
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

	//json:optional
	CampaignSites struct {
		Excluded []FieldValueString  `json:"excluded,omitempty"`
		Included [5]FieldValueString `json:"included,omitempty"`
	}
	FieldValueString string
)

// TestSliceSlice tests struct with slice of slice
//
//json:json
type TestSliceSlice struct {
	FieldStr [][]string `json:"strs"`
	FieldInt [][]int    `json:"ints"`
}
