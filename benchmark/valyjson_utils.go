// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.
package benchmark

import (
	"fmt"
	"io"
	"strconv"
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
	ln  int // len
	br  int // current bucket
}

func (b *bufWriter) Write(p []byte) (n int, err error) {
	b.ensureSpace(len(p))
	b.ln += len(p)
	free := cap(b.buf[b.br].B) - len(b.buf[b.br].B)
	if free >= len(p) {
		b.buf[b.br].B = append(b.buf[b.br].B, p...)
		return len(p), nil
	}
	b.buf[b.br].B = append(b.buf[b.br].B, p[:free]...)
	b.br++
	b.buf[b.br].B = append(b.buf[b.br].B, p[free:]...)
	return len(p), nil
}

func (b *bufWriter) WriteString(s string) (n int, err error) {
	return b.Write((*(*[0x7fff0000]byte)(unsafe.Pointer(&s)))[:len(s)])
}

func (b *bufWriter) Len() int {
	return b.ln
}

func (b *bufWriter) Bytes() []byte {
	out := make([]byte, 0, b.ln)
	var expLn = b.ln
	for _, buf := range b.buf {
		if len(buf.B) > expLn {
			out = append(out, buf.B[:expLn]...)
			break
		}
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

const minBufBlock = 65535

func (b *bufWriter) ensureSpace(minNeededSz int) {
	if cap(b.buf) == 0 {
		b.buf = make([]*bytebufferpool.ByteBuffer, 0, 64)
	}
	if len(b.buf) > 0 {
		free := cap(b.buf[b.br].B) - len(b.buf[b.br].B)
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
	w.ln = 0
	w.br = 0
	w.buf = w.buf[:0]
	c.pool.Put(w)
}

var commonBuffer = cb{}
