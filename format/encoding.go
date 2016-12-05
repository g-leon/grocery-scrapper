package format

import "bytes"

// Encoding replaces unsafe characters with their string representation
func Encoding(b []byte) []byte {
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	return b
}
