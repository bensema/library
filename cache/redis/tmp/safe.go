// +build appengine

package tmp

func String(b []byte) string {
	return string(b)
}

func Bytes(s string) []byte {
	return []byte(s)
}
