package benchmark

import (
	"testing"

	"github.com/mailru/easyjson/jwriter"
)

// BenchmarkEJ_Unmarshal_M-8                   	   10000	    107585 ns/op	 121.06 MB/s	    9512 B/op	     126 allocs/op
// BenchmarkEJ_Unmarshal_S-8                   	  240164	      4583 ns/op	  39.93 MB/s	     416 B/op	      13 allocs/op

func BenchmarkVJ_Unmarshal_M(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		err := s.UnmarshalJSON(largeStructText)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkVJ_Unmarshal_S(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallStructText)))

	for i := 0; i < b.N; i++ {
		var s Entities
		err := s.UnmarshalJSON(smallStructText)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkVJ_Marshal_M(b *testing.B) {
	b.ReportAllocs()
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := largeStructData.MarshalJSON()
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkVJ_Marshal_L(b *testing.B) {
	b.ReportAllocs()
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := xlStructData.MarshalJSON()
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkVJ_Marshal_M_Parallel(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeStructText)))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := largeStructData.MarshalJSON()
			if err != nil {
				b.Error(err)
			}
		}
	})
}

// BenchmarkEJ_Marshal_L_Parallel-8           10178     117615 ns/op    3798.93 MB/s      457667 B/op         27 allocs/op
func BenchmarkVJ_Marshal_L_Parallel(b *testing.B) {
	b.ReportAllocs()
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := xlStructData.MarshalJSON()
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}

func BenchmarkVJ_Marshal_L_ToWriter_Parallel(b *testing.B) {
	b.ReportAllocs()
	out := &DummyWriter{}
	b.RunParallel(func(pb *testing.PB) {
		var l int64
		for pb.Next() {
			w := jwriter.Writer{}

			err := xlStructData.MarshalTo(&w)
			if err != nil {
				b.Error(w.Error)
			}
			l = int64(w.Size())
			w.DumpTo(out)
		}
		if l > 0 {
			b.SetBytes(l)
		}
	})
}

func BenchmarkVJ_Marshal_S(b *testing.B) {
	b.ReportAllocs()
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := smallStructData.MarshalJSON()
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkVJ_Marshal_S_Parallel(b *testing.B) {
	b.ReportAllocs()
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := smallStructData.MarshalJSON()
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}
