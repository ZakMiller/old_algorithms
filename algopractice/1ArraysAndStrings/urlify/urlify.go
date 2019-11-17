package urlify

import "strings"

func Urlify(s string) string {
	b := strings.Builder{}
	for _, r := range s {
		if r == ' ' {
			b.WriteString("%20")
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
