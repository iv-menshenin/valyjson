package test_extr

import (
	"fill/test_any"
	"fill/test_string"
)

type (
	//json:optional
	ExternalNested struct {
		test_any.TestAllOfSecond
		test_any.TestAllOfThird
		test_string.TestStr01
	}
)
