package test

import "bytes"

// Struct contains all fields for struct
//  valyjson:encode,decode,strict
type Struct struct {
	Filter string `json:"filter,required"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty" default:"100"`

	Nested Nested `json:"nested"`
}

// MarshalJSON implements json.Marshaler
func (s *Struct) MarshalJSON() ([]byte, error) {
	var out = bytes.NewBuffer(make([]byte, 0, 128))
	for _, r := range s.Filter {
		var err error
		switch r {

		case '\t':
			_, err = out.WriteString(`\t`)

		case '\r':
			_, err = out.WriteString(`\r`)

		case '\n':
			_, err = out.WriteString(`\n`)

		case '\\':
			_, err = out.WriteString(`\\`)

		case '"':
			_, err = out.WriteString(`\"`)

		default:
			_, err = out.WriteRune(r)
		}
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
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

	Height    uint32  `json:"height"`
	HeightRef *uint32 `json:"heightRef"`
	Weight    uint64  `json:"weight"`
	WeightRef *uint64 `json:"weightRef"`
}
