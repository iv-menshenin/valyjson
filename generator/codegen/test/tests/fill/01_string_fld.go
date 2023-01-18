package testo

//valyjson:optional
type TestStr01 struct {
	Field    string  `json:"field"`
	FieldRef *string `json:"fieldRef"`
}

//valyjson:strict
type TestStr02 struct {
	Field    string  `json:"field"`
	FieldRef *string `json:"fieldRef"`
}
