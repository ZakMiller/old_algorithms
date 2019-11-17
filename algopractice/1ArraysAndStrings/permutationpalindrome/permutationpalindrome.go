package permutationpalindrome

type present struct{}
func IsPermutationPalindrome(s string) bool {
	m := make(map[rune]struct{})
	for _, r := range s {
		_, ok := m[r]
		if ok {
			delete(m, r)
		} else {

			m[r] = present{}
		}
	}

	n := 0
	for range m {
		n++
	}
	validEven := len(s) % 2 == 0 && n == 0
	validOdd := len(s) %2 == 1 && n == 1
	return validEven || validOdd

}