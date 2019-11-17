package hamming

func max(one, two int) int {
	if one > two {
		return one
	}
	return two
}
func hamming(one, two int) int {
	differenceCount := 0
	comparer := 1
	higher := max(one, two)
	for comparer <= higher {
		oneD := one & comparer
		twoD := two & comparer
		if oneD != twoD {
			differenceCount++
		}
		comparer = comparer << 1
	}
	return differenceCount
}
