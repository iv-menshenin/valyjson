package main

import "testing"

func Benchmark_FastJson(b *testing.B) {
	b.ReportAllocs()
	var data = []byte(j)
	for i := 0; i < b.N; i++ {
		var s Struct
		if err := s.UnmarshalJSON(data); err != nil {
			b.Error(err)
		}
	}
}
