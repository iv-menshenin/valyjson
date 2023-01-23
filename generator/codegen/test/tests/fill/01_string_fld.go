package testo

//json:optional
type TestStr01 struct {
	Field    string  `json:"field"`
	FieldRef *string `json:"fieldRef"`
}

//json:strict
type TestStr02 struct {
	Field    string  `json:"field"`
	FieldRef *string `json:"fieldRef"`
}
