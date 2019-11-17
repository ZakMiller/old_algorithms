package lower

import "strings"

func toLower(s string) string {
	b := strings.Builder{}

	for i := 0; i < len(s); i++ {
		v := int(s[i])
		if v >= 65 && v <= 91 {
			b.WriteByte(byte(v + 32))
		} else {
			b.WriteByte(byte(v))
		}
	}
	return b.String()
}
