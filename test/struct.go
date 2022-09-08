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

//  valyjson:encode,decode
type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	// DOB time.Time `json:"dob"`
	Rate64 float64 `json:"rate64" default:"1"`
	Rate32 float32 `json:"rate32" default:"1"`

	Height uint32 `json:"height"`
	Weight uint64 `json:"weight"`
}
