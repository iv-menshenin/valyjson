package test_nostruct

import (
	"fill/test_extr"
	"time"
)

type (
	// TestMap10 tests maps
	//json:strict
	TestMap10 map[string]int64

	// TestMap11 tests maps
	//json:strict
	TestMap11 map[string]test_extr.External

	// TestMap11Ref tests maps
	//json:strict
	TestMap11Ref map[string]*test_extr.External

	// TestSlice12 tests maps
	//json:strict
	TestSlice12 []int64

	// TestSlice13 tests maps
	//json:strict
	TestSlice13 []test_extr.External

	// TestSlice14 tests maps
	//json:strict
	TestSlice14 [2]time.Time

	// TODO []map[string][]int64
)
