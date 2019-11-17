package closesthigher

import "math"

func closestLower(v int) int {
	if v == 0 {
		return 0
	}
	slider := 1
	var seenZero bool
	for slider <= v {
		isZero := v&slider == 0
		if isZero {
			seenZero = true
		} else {
			if seenZero {
				return moveRight(v, slider)
			}
		}
		slider <<= 1
	}
	return v // Can't go lower than e.g. 3, 7 so just return that number
}

func moveRight(v, slider int) int {
	rightOne := slider >> 1
	leftFill := -1 << uint(math.Log2(float64(1+slider)))
	rightZeros := leftFill | slider - 1
	v &= rightZeros // Zero out the slider index
	v |= rightOne   // Add the one to the right of the slider index
	return v
}
