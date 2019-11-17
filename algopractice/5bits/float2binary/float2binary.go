package float2binary

import (
	"math"
	"strings"
)

func floatToBinary(f float64) string {

	var sum float64
	base := .5
	b := strings.Builder{}
	b.WriteByte('.')
	for base > 1/(math.Pow(2, 32)) && sum != f {
		if f-sum >= base {
			sum += base
			b.WriteRune('1')
		} else {
			b.WriteRune('0')
		}
		base /= 2
	}

	switch {
	case sum == f:
		return b.String()
	default:
		return "ERROR"
	}
}
