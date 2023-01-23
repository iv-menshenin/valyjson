package testo

//json:optional
type NumStruct01 struct {
	IntFld        int      `json:"int_fld"`
	IntFld8       int8     `json:"int_fld8"`
	IntFld16      int16    `json:"int_fld16"`
	IntFld32      int32    `json:"int_fld32"`
	IntFld64      int64    `json:"int_fld64"`
	UintFld       uint     `json:"Uint_fld"`
	UintFld8      uint8    `json:"Uint_fld8"`
	UintFld16     uint16   `json:"Uint_fld16"`
	UintFld32     uint32   `json:"Uint_fld32"`
	UintFld64     uint64   `json:"Uint_fld64"`
	FloatFld32    float32  `json:"fl23"`
	FloatFld64    float64  `json:"fl64"`
	RefIntFld     *int     `json:"ref_int_fld"`
	RefIntFld8    *int8    `json:"ref_int_fld8"`
	RefIntFld16   *int16   `json:"ref_int_fld16"`
	RefIntFld32   *int32   `json:"ref_int_fld32"`
	RefIntFld64   *int64   `json:"ref_int_fld64"`
	RefUintFld    *uint    `json:"ref_Uint_fld"`
	RefUintFld8   *uint8   `json:"ref_Uint_fld8"`
	RefUintFld16  *uint16  `json:"ref_Uint_fld16"`
	RefUintFld32  *uint32  `json:"ref_Uint_fld32"`
	RefUintFld64  *uint64  `json:"ref_Uint_fld64"`
	RefFloatFld32 *float32 `json:"ref_fl23"`
	RefFloatFld64 *float64 `json:"ref_fl64"`
}

//json:strict
type NumStruct02 struct {
	IntFld        int      `json:"int_fld"`
	IntFld8       int8     `json:"int_fld8"`
	IntFld16      int16    `json:"int_fld16"`
	IntFld32      int32    `json:"int_fld32"`
	IntFld64      int64    `json:"int_fld64"`
	UintFld       uint     `json:"Uint_fld"`
	UintFld8      uint8    `json:"Uint_fld8"`
	UintFld16     uint16   `json:"Uint_fld16"`
	UintFld32     uint32   `json:"Uint_fld32"`
	UintFld64     uint64   `json:"Uint_fld64"`
	FloatFld32    float32  `json:"fl23"`
	FloatFld64    float64  `json:"fl64"`
	RefIntFld     *int     `json:"ref_int_fld"`
	RefIntFld8    *int8    `json:"ref_int_fld8"`
	RefIntFld16   *int16   `json:"ref_int_fld16"`
	RefIntFld32   *int32   `json:"ref_int_fld32"`
	RefIntFld64   *int64   `json:"ref_int_fld64"`
	RefUintFld    *uint    `json:"ref_Uint_fld"`
	RefUintFld8   *uint8   `json:"ref_Uint_fld8"`
	RefUintFld16  *uint16  `json:"ref_Uint_fld16"`
	RefUintFld32  *uint32  `json:"ref_Uint_fld32"`
	RefUintFld64  *uint64  `json:"ref_Uint_fld64"`
	RefFloatFld32 *float32 `json:"ref_fl23"`
	RefFloatFld64 *float64 `json:"ref_fl64"`
}
