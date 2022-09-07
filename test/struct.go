package test

// Struct contains all fields for struct
//  valyjson:encode,decode,strict
type Struct struct {
	Filter string `json:"filter,required"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty" default:"100"`

	Nested Nested `json:"nested"`
}

//  valyjson:encode,decode,strict
type Nested struct {
	List  []int64 `json:"list"`
	Count *int64  `json:"count"`
	Cross *int64  `json:"cross"`
}
