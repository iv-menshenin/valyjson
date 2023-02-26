package test_map

// TestMap01 tests maps
//json:strict
type TestMap01 struct {
	Tags               map[string]string   `json:"tags"`
	Properties         map[string]Property `json:"properties,omitempty"`
	KeyTypedProperties map[Key]Property    `json:"key_typed_properties"`
	IntegerVal         map[Key]int32       `json:"integerVal,omitempty"`
	FloatVal           map[Key]float64     `json:"floatVal,omitempty"`
	UintVal            map[Key]uint16      `json:"uintVal,omitempty"`
}

// Property tests properties
//json:strict
type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Key string
