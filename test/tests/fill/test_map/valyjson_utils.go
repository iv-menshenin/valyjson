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
}

type bufWriter struct {
	buf []*bytebufferpool.ByteBuffer
	br  int // current bucket
	sz  int
	cb  *cb
}

func (b *bufWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	free := b.ensureSpace(n)
	if free >= n {
		b.buf[b.br].B = append(b.buf[b.br].B, p...)
		return n, nil
	}
	for len(p) > 0 {
		sz := cap(b.buf[b.br].B) - len(b.buf[b.br].B)
		if sz >= len(p) {
			b.buf[b.br].B = append(b.buf[b.br].B, p...)
			break
		}
		b.buf[b.br].B = append(b.buf[b.br].B, p[:sz]...)
		p = p[sz:]
		b.br++
	}
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

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func (b *bufWriter) Len() (l int) {
	for _, seg := range b.buf {
		l += len(seg.B)
	}
	return
}

func (b *bufWriter) Bytes() []byte {
	var sz int
	out := make([]byte, 0, b.Len())
	for _, buf := range b.buf {
		sz += len(buf.B)
		out = append(out, buf.B...)
	}
	b.sz = tuneBuf(b.sz, sz)
	b.Close()
	return out
}

func tuneBuf(old, cur int) int {
	var nw = cur
	if old > 0 {
		nw = (cur + old) / 2
	}
	const kib = 1024
	switch {
	case nw > 22*kib:
		return 32 * kib
	case nw > 16*kib:
		return 22 * kib
	case nw > 10*kib:
		return 16 * kib
	case nw > 8*kib:
		return 10 * kib
	case nw > 4*kib:
		return 8 * kib
	case nw > 2*kib:
		return 4 * kib
	case nw > kib:
		return 2 * kib
	case nw > 512:
		return kib
	case nw > 256:
		return 512
	case nw > 128:
		return 256
	default:
		return 128
	}
}

func (b *bufWriter) Close() error {
	for i := range b.buf {
		writeBuf.Put(b.buf[i])
	}
	if b.cb != nil {
		b.cb.Put(b)
	}
	return nil
}

const defBufBlock = 8192

func (b *bufWriter) ensureSpace(minNeededSz int) (currBlockFree int) {
	if len(b.buf) > 0 {
		currBlockFree = cap(b.buf[b.br].B) - len(b.buf[b.br].B)
		if currBlockFree >= minNeededSz {
			return currBlockFree
		}
		minNeededSz -= currBlockFree
	}
	for {
		bb := writeBuf.Get()
		if cap(bb.B) == 0 {
			if b.sz == 0 {
				b.sz = defBufBlock
			}
			if minNeededSz < b.sz {
				minNeededSz = b.sz
			}
			bb.B = make([]byte, 0, minNeededSz)
		}
		if len(b.buf) == 0 {
			currBlockFree = cap(bb.B)
		}
		b.buf = append(b.buf, bb)
		if cap(bb.B) >= minNeededSz {
			return currBlockFree
		}
		minNeededSz -= cap(bb.B)
	}
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
	var buf = stringBuf.Get()
	defer stringBuf.Put(buf)
	flush := func() {
		if buf.Len() > 0 {
			buf.WriteTo(w)
			buf.B = buf.B[:0]
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
			buf.WriteString(string(r))
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
		return &bufWriter{
			cb: c,
		}
	}
	b := p.(*bufWriter)
	b.cb = c
	return b
}

func (c *cb) Put(w *bufWriter) {
	w.br = 0
	w.buf = w.buf[:0]
	c.pool.Put(w)
}
