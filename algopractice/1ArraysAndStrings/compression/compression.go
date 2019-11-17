package compression

import (
	"strconv"
	"strings"
)

func Compress(s string) string {
	b := strings.Builder{}
	n := 1
	for i, r := range s {
		if i + 1 < len(s) && s[i] == s[i + 1] {
			n++
		} else {
			b.WriteRune(r)
			b.WriteString(strconv.Itoa(n))
			n = 1
		}
	}

	bs := b.String()
	if len(bs) < len(s) {
		return bs
	} else {
		return s
	}
}
