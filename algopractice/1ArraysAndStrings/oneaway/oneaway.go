package oneaway

func OneAway(first string, second string) bool {
	lenFirst := len(first)
	lenSecond := len(second)

	isInsertion := lenSecond == lenFirst + 1
	if isInsertion {
		i, j := 0, 0

		for i < lenFirst {
			if first[i] == second[j] {
				i++
				j++
			} else {
					if j > i {
						return false
					} else {
						j++
					}
			}
		}
		return true
	}

	isDeletion := lenSecond == lenFirst - 1
	if isDeletion {
		i, j := 0, 0

		for j < lenSecond {
			if first[i] == second[j] {
				i++
				j++
			} else {
				if j < i {
					return false
				} else {
					i++
				}
			}
		}
		return true
	}

	isReplace := lenSecond == lenFirst
	if isReplace {
		replaced := false
		for i := range first {
			if first[i] != second[i] {
				if replaced {
					return false
				} else {
					replaced = true
				}
			}
		}
		return true
	}
	return false
}