// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package test_map

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fastjson"
)

type Writer interface {
	io.Writer
	io.StringWriter
	Len() int
}

type bufWriter struct {
	buf []*bytebufferpool.ByteBuffer
	br  int // current bucket
}

func (b *bufWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	free := b.ensureSpace(n)
	if free >= n {
		b.buf[b.br].B = append(b.buf[b.br].B, p...)
		return n, nil
	}
	b.buf[b.br].B = append(b.buf[b.br].B, p[:free]...)
	b.br++
	b.buf[b.br].B = append(b.buf[b.br].B, p[free:]...)
	return n, nil
}

func (b *bufWriter) WriteString(s string) (n int, err error) {
	return b.Write(s2b(s))
}

func s2b(s string) (b []byte) {
	strh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh.Data = strh.Data
	sh.Len = strh.Len
	sh.Cap = strh.Len
	return b
}

func (b *bufWriter) Len() (l int) {
	for _, seg := range b.buf {
		l += len(seg.B)
	}
	return
}

func (b *bufWriter) Bytes() []byte {
	out := make([]byte, 0, b.Len())
	for _, buf := range b.buf {
		out = append(out, buf.B...)
	}
	b.Close()
	return out
}

func (b *bufWriter) Close() error {
	for i := range b.buf {
		writeBuf.Put(b.buf[i])
	}
	commonBuffer.Put(b)
	return nil
}

const minBufBlock = 32768

func (b *bufWriter) ensureSpace(minNeededSz int) (free int) {
	if len(b.buf) > 0 {
		free = cap(b.buf[b.br].B) - len(b.buf[b.br].B)
		if free >= minNeededSz {
			return
		}
		minNeededSz -= free
	}
	needSz := minNeededSz
	if needSz < minBufBlock {
		needSz = minBufBlock
	}
	bb := writeBuf.Get()
	if len(bb.B) < minNeededSz {
		bb.B = make([]byte, 0, needSz)
	}
	b.buf = append(b.buf, bb)
	if len(b.buf) == 1 {
		return needSz
	}
	return
}

var writeBuf = bytebufferpool.Pool{}

func valueIsNotNull(v *fastjson.Value) bool {
	return v != nil && v.Type() != fastjson.TypeNull
}

func parseDateTime(s string) (time.Time, error) {
	knownFormats := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02 15:04:05Z07:00",
		"02.01.2006 15:04:05Z07:00",
	}
	for _, format := range knownFormats {
		if value, err := time.Parse(format, s); err == nil {
			return value, nil
		}
	}
	return time.Time{}, fmt.Errorf("can't parse date-time from string '%s'", s)
}

func parseDate(s string) (time.Time, error) {
	knownFormats := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02",
		"02.01.2006",
	}
	for _, format := range knownFormats {
		if value, err := time.Parse(format, s); err == nil {
			return value.Truncate(24 * time.Hour), nil
		}
	}
	return time.Time{}, fmt.Errorf("can't parse date from string '%s'", s)
}

var intBuf = bytebufferpool.Pool{}

func writeInt64(w io.Writer, i int64) {
	buf := intBuf.Get()
	buf.B = strconv.AppendInt(buf.B[:0], i, 10)
	w.Write(buf.B)
	intBuf.Put(buf)
}

func writeUint64(w io.Writer, i uint64) {
	buf := intBuf.Get()
	buf.B = strconv.AppendUint(buf.B[:0], i, 10)
	w.Write(buf.B)
	intBuf.Put(buf)
}

var fltBuf = bytebufferpool.Pool{}

func writeFloat64(w io.Writer, f float64) {
	buf := fltBuf.Get()
	buf.B = strconv.AppendFloat(buf.B[:0], f, 'f', -1, 64)
	w.Write(buf.B)
	fltBuf.Put(buf)
}

var timeBuf = bytebufferpool.Pool{}

func writeTime(w io.Writer, t time.Time, layout string) {
	buf := timeBuf.Get()
	buf.B = append(buf.B[:0], '"')
	buf.B = t.AppendFormat(buf.B, layout)
	buf.B = append(buf.B, '"')
	w.Write(buf.B)
	timeBuf.Put(buf)
}

var stringBuf = bytebufferpool.Pool{}

func writeString(w Writer, s string) {
	w.WriteString(`"`)
	if !hasSpecialChars(s) {
		w.WriteString(s)
		w.WriteString(`"`)
		return
	}
	var (
		buf [128]byte
		idx int
	)
	flush := func() {
		if len(buf) > 0 {
			w.WriteString(string(buf[:idx]))
			idx = 0
		}
	}
	for _, r := range s {
		switch r {

		case '\t':
			flush()
			w.WriteString(`\t`)

		case '\r':
			flush()
			w.WriteString(`\r`)

		case '\n':
			flush()
			w.WriteString(`\n`)

		case '\\':
			flush()
			w.WriteString(`\\`)

		case '"':
			flush()
			w.WriteString(`\"`)

		default:
			if len(buf) >= cap(buf)-2 {
				flush()
			}
			if r < 256 {
				buf[idx] = byte(r & 0xff)
				idx++
			} else {
				buf[idx] = byte(r >> 8)
				idx++
				buf[idx] = byte(r & 0xff)
				idx++
			}
		}
	}
	flush()
	w.WriteString(`"`)
}

func hasSpecialChars(s string) bool {
	if strings.IndexByte(s, '"') >= 0 || strings.IndexByte(s, '\\') >= 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 0x20 {
			return true
		}
	}
	return false
}

type cb struct {
	pool sync.Pool
}

func (c *cb) Get() *bufWriter {
	p := c.pool.Get()
	if p == nil {
		return &bufWriter{}
	}
	return p.(*bufWriter)
}

func (c *cb) Put(w *bufWriter) {
	w.br = 0
	w.buf = w.buf[:0]
	c.pool.Put(w)
}

var commonBuffer = cb{}
