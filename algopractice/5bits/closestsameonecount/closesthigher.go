package closesthigher

import "math"

func closestHigher(v int) int {
	if v == 0 {
		return 0
	}
	slider := 1
	var seenOne bool
	for {
		isOne := v&slider != 0
		if isOne {
			seenOne = true
		} else {
			if seenOne {
				return moveLeft(v, slider)
			}
		}
		slider <<= 1
	}
}

func moveLeft(v, slider int) int {
	leftOne := slider
	leftFill := -1 << uint(math.Log2(float64(slider)))
	rightOnes := leftFill | (slider >> 1) - 1
	v &= rightOnes
	v |= leftOne
	return v
}
