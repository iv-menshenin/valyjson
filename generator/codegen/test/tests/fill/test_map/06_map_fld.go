package test_map

// TestMap01 tests maps
//json:strict,decode
type TestMap01 struct {
	Tags               map[string]string   `json:"tags"`
	Properties         map[string]Property `json:"properties"`
	KeyTypedProperties map[Key]Property    `json:"key_typed_properties"`
}

// Property tests properties
//json:strict,decode
type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Key string
