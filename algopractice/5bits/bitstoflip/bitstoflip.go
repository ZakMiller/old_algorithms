package bitstoflip

func bitsToFlip(a, b int) int {
	var n int
	slider := 1
	for slider <= a || slider <= b {
		aIsOne := a&slider != 0
		bIsOne := b&slider != 0
		if aIsOne != bIsOne {
			n++
		}
		slider <<= 1
	}
	return n
}
