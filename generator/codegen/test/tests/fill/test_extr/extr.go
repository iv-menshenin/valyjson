package test_extr

import (
	"fill/test_any"
	"fill/test_string"
)

type (
	//json:optional
	External struct {
		Test01 test_any.TestAllOf01  `json:",inline"`
		Test02 test_string.TestStr01 `json:",inline"`
	}
)
