package test

import "time"

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
	List  []int32 `json:"list-i"`
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

	Height    uint32  `json:"height"`
	HeightRef *uint32 `json:"heightRef" default:"443"`
	Weight    uint64  `json:"weight"`
	WeightRef *uint64 `json:"weightRef,omitempty"`

	Bio *Bio `json:"bio,omitempty"`
}

//  valyjson:encode,decode,strict
type Bio struct {
	Description *string    `json:"description,omitempty"`
	Changed     *time.Time `json:"changed,omitempty"`
	Level       *int       `json:"level,omitempty"`
	Name        *int       `json:"name"`
}
