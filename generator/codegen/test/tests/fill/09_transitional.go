package testo

// TestTransitional test transitional json generation
//json:transit
type TestTransitional TestTransitionalElem

// TestTransitionalElem test transitional json generation
//json:optional
type TestTransitionalElem struct {
	TestField int64 `json:"test-field"`
}
