package stringbuilder

import (
	"bytes"
)

type StringBuilder struct {
	strings       []string
}

func (sb *StringBuilder) AddString(newString string) {
	sb.strings = append(sb.strings, newString)
}

func (sb *StringBuilder) ToString() string {
	var buffer bytes.Buffer

	for _, str := range sb.strings {
		buffer.WriteString(str)
	}
	return buffer.String()
}