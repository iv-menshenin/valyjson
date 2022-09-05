package test

import (
	"github.com/valyala/fastjson"
)

// Struct contains all fields for struct
//  valyjson:encode,decode,strict
type Struct struct {
	Filter string `json:"filter,required"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty" default:"100"`

	Nested Nested `json:"nested"`
}

type Nested struct {
	List  []int64 `json:"list"`
	Count *int64  `json:"count"`
	Cross *int64  `json:"cross"`
}

func valueIsNotNull(v *fastjson.Value) bool {
	return v != nil && v.Type() != fastjson.TypeNull
}

var structPool fastjson.ParserPool
var nestedPool fastjson.ParserPool
