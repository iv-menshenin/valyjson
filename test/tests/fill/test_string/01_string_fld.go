package test_string

//json:optional
type TestStr01 struct {
	Field    string  `json:"field"`
	FieldRef *string `json:"fieldRef"`
	DefRef   *string `json:"defRef" default:"default"`
}

//json:strict
type TestStr02 struct {
	Field    string           `json:"field"`
	FieldRef *string          `json:"fieldRef"`
	String   FieldValueString `json:"string" default:"value-foo-bar"`
}

type FieldValueString string
