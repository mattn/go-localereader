package localereader

import (
	"io"
)

func NewReader(r io.Reader) io.Reader {
	return newReader(r)
}
