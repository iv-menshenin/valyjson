package benchmark

import (
	"bytes"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bufWriter_Write(t *testing.T) {
	t.Parallel()
	t.Run("write", func(t *testing.T) {
		t.Parallel()
		var w bufWriter
		var ex []byte
		for n := 0; n < 256; n++ {
			wr := bytes.Repeat([]byte{byte(n)}, 68)
			n, err := w.Write(wr)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if n != 68 {
				t.Errorf("expected 68 writed bytes, got: %d", n)
			}
			ex = append(ex, wr...)
		}
		require.Equal(t, ex, w.Bytes())
	})
	t.Run("allocations", func(t *testing.T) {
		var bx [256]byte
		n := testing.AllocsPerRun(1000, func() {
			var w = bufDataXLStruct.Get()
			var x = 16384
			for x > 0 {
				pie := 255
				if pie > x {
					pie = x
				}
				x -= pie
				w.Write(bx[:pie])
			}
			w.Close()
		})
		require.EqualValues(t, 0, n)
	})
}

// Benchmark_testWriter-8                     217779              5426 ns/op               0 B/op          0 allocs/op
func Benchmark_testWriter(b *testing.B) {
	var pool cb
	var bx [256]byte
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var w = pool.Get()
		var x = 65535
		for x > 0 {
			pie := 255
			if pie > x {
				pie = x
			}
			x -= pie
			w.Write(bx[:pie])
		}
		w.Close()
	}
}

func Test_writeString(t *testing.T) {
	t.Run("ascii", func(t *testing.T) {
		var b = bytes.NewBufferString("")
		writeString(b, "te\tst\n")
		require.Equal(t, `"te\tst\n"`, b.String())
	})
	t.Run("double", func(t *testing.T) {
		var b = bytes.NewBufferString("")
		writeString(b, "съешь еще этих \\горьких апельсинов\n")
		require.Equal(t, `"съешь еще этих \\горьких апельсинов\n"`, b.String())
	})
	t.Run("china", func(t *testing.T) {
		var b = bytes.NewBufferString("")
		writeString(b, "去吧，离开这里。\n")
		require.Equal(t, `"去吧，离开这里。\n"`, b.String())
	})
}

func Test_bufWriter_ensureSpace(t *testing.T) {
	t.Run("take bigger first", func(t *testing.T) {
		var w = bufWriter{
			br: 0,
			sz: 128,
		}
		_, err := w.Write(make([]byte, 129))
		require.NoError(t, err)
	})
	t.Run("take bigger second", func(t *testing.T) {
		var w = bufWriter{
			br: 0,
			sz: 128,
		}
		_, err := w.Write(make([]byte, 64))
		require.NoError(t, err)
		_, err = w.Write(make([]byte, 290))
		require.NoError(t, err)
	})
}

func Test_testWriter_Parallelism(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	var etalon []byte
	for r := 0; r < 64; r++ {
		for x := byte(0); x < 254; x++ {
			etalon = append(etalon, bytes.Repeat([]byte{x}, 250)...)
		}
	}
	var ch = make(chan struct{}, 64)
	for n := 0; n < 255; n++ {
		wg.Add(1)
		go func() {
			ch <- struct{}{}
			defer func() {
				<-ch
			}()
			defer wg.Done()
			var w bufWriter
			r := bytes.NewReader(etalon)
			for {
				var buf = make([]byte, 215)
				read, err := r.Read(buf)
				if err == io.EOF {
					break
				}
				require.NoError(t, err)
				_, err = w.Write(buf[:read])
				require.NoError(t, err)
			}
			require.Equal(t, etalon, w.Bytes())
		}()
	}
	wg.Wait()
}
