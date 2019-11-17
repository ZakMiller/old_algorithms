package replace

func getClearer(i, j uint) int {
	left := -1<<j + 1
	right := (1 << i) - 1
	return left | right
}

func clear(n int, i, j uint) int {
	clearer := getClearer(i, j)
	return clearer & n
}
func replace(n, m int, i, j uint) int {
	n = clear(n, i, j)
	m <<= i
	return m | n
}
