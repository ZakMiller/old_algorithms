package unique

func IsUnique(s string) bool {
	m := make(map[rune]bool)
	for _, r := range s {
		_, ok := m[r]
		if ok {
			return false
		}
		m[r] = true
	}
	return true
}
