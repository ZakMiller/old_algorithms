package permutation

func IsPermutation(one string, two string) bool {
	if len(one) != len(two) {
		return false
	}
	runeCount := make(map[rune]int)
	for _, r := range one {
		runeCount[r]++
	}

	for _, r := range two {
		runeCount[r]--
		if runeCount[r] < 0 {
			return false
		}
	}
	return true
}
