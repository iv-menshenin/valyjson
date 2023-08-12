package test_packages

import pack_a "fill/test_packages/pack_b"

//json:optional
type Test01 struct {
	Field pack_a.Test `json:"field"`
}
