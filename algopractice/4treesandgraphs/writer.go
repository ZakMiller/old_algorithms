package treesandgraphs

import (
	"fmt"
	"strconv"
	"strings"
)

type writer struct {
	total    int
	rowTotal int
	items    []int
}

func (w *writer) store(val int) {
	w.items = append(w.items, val)
}

func isEndOfLine(i int, n int) bool {
	oneBased := i + 1
	isOneLessThanPowerOfTwo := oneBased&(oneBased+1) == 0
	lastElementOnLine := isOneLessThanPowerOfTwo
	lastElementInList := i == n-1
	return lastElementOnLine || lastElementInList

}
func (w *writer) String() string {
	b := strings.Builder{}

	for i := 0; i < len(w.items); i++ {
		isEndOfLine := isEndOfLine(i, len(w.items))
		if !isEndOfLine {
			b.WriteString(fmt.Sprintf("%v ", strconv.Itoa(w.items[i])))
		} else {
			b.WriteString(fmt.Sprintf("%v\n", strconv.Itoa(w.items[i])))
		}
	}
	return b.String()
}
