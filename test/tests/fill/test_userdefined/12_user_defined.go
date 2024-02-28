package test_userdefined

import "fill/test_userdefined/userdefined"

// TestUserDefined tests inheritance
//
//json:optional
type TestUserDefined struct {
	Int32   DefinedInt32   `json:"f_int32" default:"32"`
	Int64   DefinedInt64   `json:"f_int64,omitempty"`
	Float32 DefinedFloat32 `json:"f_float32" default:"123.01"`
	Float64 DefinedFloat64 `json:"f_float64,omitempty"`
	String  DefinedString  `json:"f_string" default:"default_string"`
	Bool    DefinedBool    `json:"f_bool,omitempty"`

	RefInt32   *DefinedInt32   `json:"r_int32,omitempty"`
	RefInt64   *DefinedInt64   `json:"r_int64,omitempty"`
	RefFloat32 *DefinedFloat32 `json:"r_float32,omitempty"`
	RefFloat64 *DefinedFloat64 `json:"r_float64,omitempty"`
	RefString  *DefinedString  `json:"r_string,omitempty"`
	RefBool    *DefinedBool    `json:"r_bool,omitempty"`
}

//json:optional
type TestUserDefinedRef struct {
	Int32Ref   DefinedRefInt32   `json:"d_int32" default:"32"`
	Int64Ref   DefinedRefInt64   `json:"d_int64,omitempty"`
	Float32Ref DefinedRefFloat32 `json:"d_float32" default:"123.01"`
	Float64Ref DefinedRefFloat64 `json:"d_float64,omitempty"`
	StringRef  DefinedRefString  `json:"d_string" default:"default_string"`
	BoolRef    DefinedRefBool    `json:"d_bool,omitempty"`

	RefInt32Ref   *DefinedRefInt32   `json:"x_int32" default:"32"`
	RefInt64Ref   *DefinedRefInt64   `json:"x_int64,omitempty"`
	RefFloat32Ref *DefinedRefFloat32 `json:"x_float32" default:"123.01"`
	RefFloat64Ref *DefinedRefFloat64 `json:"x_float64,omitempty"`
	RefStringRef  *DefinedRefString  `json:"x_string" default:"default_string"`
	RefBoolRef    *DefinedRefBool    `json:"x_bool,omitempty"`
}

type (
	DefinedInt32   int32
	DefinedInt64   int64
	DefinedFloat32 float32
	DefinedFloat64 float64
	DefinedString  string
	DefinedBool    bool

	DefinedRefInt32   *int32
	DefinedRefInt64   *int64
	DefinedRefFloat32 *float32
	DefinedRefFloat64 *float64
	DefinedRefString  *string
	DefinedRefBool    *bool
)

type (
	//json:json
	DefinedFieldAsUserDefined struct {
		userdefined.DefinedFieldAsUserDefined1 `json:",inline"`
		userdefined.DefinedFieldAsUserDefined2 `json:",inline"`
	}
)
