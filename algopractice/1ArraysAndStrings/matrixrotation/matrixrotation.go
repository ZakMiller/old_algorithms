package matrixrotation

func Rotate(m [][]uint32) [][]uint32 {
	newMatrix := make([][]uint32, len(m))
	for i := range newMatrix {
		newMatrix[i] = make([]uint32, len(m))
	}
	l := len(m)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			newY := l - 1 - i
			newX := j
			newMatrix[newX][newY] = m[i][j]
		}
	}
	return newMatrix
}

func RotateInPlace(m [][]uint32) {
	for i := 0; i < len(m)/2; i++ {
		for j := i; j < len(m[0])-i-1; j++ {
			top := m[i][j]
			right := m[j][len(m)-1-i]
			bottom := m[len(m)-1-i][len(m)-1-j]
			left := m[len(m)-1-j][i]

			m[j][len(m)-1-i] = top
			m[len(m)-1-i][len(m)-1-j] = right
			m[len(m)-1-j][i] = bottom
			m[i][j] = left
		}
	}
}
