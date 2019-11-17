package sortedsquares

import "sort"

func sortNums(numbers []int) {

	sort.Slice(numbers, func(i, j int) bool {
		valI, valJ := numbers[i], numbers[j]
		if valI < 0 {
			valI *= -1
		}
		if valJ < 0 {
			valJ *= -1
		}
		return valI < valJ
	})
}
func sortedSquares(numbers []int) []int {
	sortNums(numbers)
	for i := range numbers {
		numbers[i] = numbers[i] * numbers[i]
	}
	return numbers
}
