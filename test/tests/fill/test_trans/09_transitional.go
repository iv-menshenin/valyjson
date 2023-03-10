package test_trans

// TestTransitional test transitional json generation
//json:transit,decode
type TestTransitional TestTransitionalElem

// TestTransitionalElem test transitional json generation
//json:optional,decode
type TestTransitionalElem struct {
	TestField int64 `json:"test-field"`
}
