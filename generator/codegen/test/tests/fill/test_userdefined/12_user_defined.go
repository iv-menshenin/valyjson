package test_userdefined

// TestUserDefined tests inheritance
//json:optional
type TestUserDefined struct {
	Int32   DefinedInt32   `json:"f_int32" default:"32"`
	Int64   DefinedInt64   `json:"f_int64,omitempty"`
	Float32 DefinedFloat32 `json:"f_float32" default:"123.01"`
	Float64 DefinedFloat64 `json:"f_float64,omitempty"`
	String  DefinedString  `json:"f_string" default:"default_string"`
	Bool    DefinedBool    `json:"f_bool,omitempty"`
}

type (
	DefinedInt32   int32
	DefinedInt64   int64
	DefinedFloat32 float32
	DefinedFloat64 float64
	DefinedString  string
	DefinedBool    bool
)
