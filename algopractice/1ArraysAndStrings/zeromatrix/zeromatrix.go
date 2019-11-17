package zeromatrix

type present struct{}

func getRowsAndColsToZeroOut(m [][]int) (map[int]present, map[int]present) {
	rowsToZeroOut := make(map[int]present)
	colsToZeroOut := make(map[int]present)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 0 {
				rowsToZeroOut[i] = present{}
				colsToZeroOut[j] = present{}
			}
		}
	}
	return rowsToZeroOut, colsToZeroOut
}

func ZeroOut(m [][]int) [][]int {
	rowsToZero, colsToZero := getRowsAndColsToZeroOut(m)

	for i := range rowsToZero {
		m[i] = make([]int, len(m[0]))
	}

	for i := 0; i < len(m); i++ {
		for j := range colsToZero {
			m[i][j] = 0
		}
	}
	return m
}
