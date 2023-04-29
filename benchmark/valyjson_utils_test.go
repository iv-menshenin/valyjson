package benchmark

import (
	"bytes"
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

func Benchmark_testWiter(b *testing.B) {
	var w bufWriter
	var bx [256]byte
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
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
