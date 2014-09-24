package sha1

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func HashString(source io.Reader) string {
	hasher := sha1.New()
	io.Copy(hasher, source)
	bs := hasher.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
