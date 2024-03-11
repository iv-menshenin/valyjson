package test_extr

import (
	"fill/test_any"
	"fill/test_string"
)

type (
	//json:optional
	External struct {
		Test01 test_any.TestAllOfSecond `json:"test1"`
		Test02 test_string.TestStr01    `json:"test2,omitempty"`
	}
	//json:json
	ExternalStructSlice []test_string.TestStr01
	//json:json
	ExternalStringSlice []test_string.FieldValueString
)
