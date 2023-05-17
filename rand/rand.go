package rand

import (
	"bytes"
	"crypto/rand"
	"io"
)

var buf *bytes.Buffer

func Read(b []byte) (int, error) {
	if buf == nil {
		return rand.Read(b)
	}

	return io.ReadFull(buf, b)
}

func Reset() {
	buf = nil
}

func Set(values []byte) {
	buf = bytes.NewBuffer(values)
}
