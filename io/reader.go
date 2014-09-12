package io

import (
	"bytes"
	"io"
)

func ReaderLen(reader io.Reader) int {
	buffer := new(bytes.Buffer)
	io.Copy(buffer, reader)
	return buffer.Len()
}
