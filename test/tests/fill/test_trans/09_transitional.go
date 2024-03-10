package test_trans

import "fill/test_extr"

// TestTransitional test transitional json generation
//
//json:json
type TestTransitional TestTransitionalElem

// TestTransitionalElem test transitional json generation
//
//json:json
type TestTransitionalElem struct {
	TestField int64 `json:"test-field"`
}

// TestExternalNested test transitional json generation
//
//json:json
type TestExternalNested test_extr.ExternalNested
