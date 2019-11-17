package stringrotation

import (
	"fmt"
	"strings"
)

type connection struct {
	first  uint8
	second uint8
}

func addConnection(first uint8, second uint8, m map[connection]int) {
	c := connection{
		first:  first,
		second: second,
	}
	_, ok := m[c]
	if ok {
		m[c]++
	} else {
		m[c] = 0
	}
}

func removeConnection(first uint8, second uint8, m map[connection]int) bool {
	c := connection{
		first:  first,
		second: second,
	}

	v, ok := m[c]
	if v == 1 {
		delete(m, c)
	} else if ok {
		m[c]--
	} else {
		return false
	}
	return true
}

func getConnections(s string) map[connection]int {
	m := make(map[connection]int)
	for i := 0; i < len(s)-1; i++ {
		addConnection(s[i], s[i+1], m)
	}
	addConnection(s[len(s)-1], s[0], m)
	return m
}

func verifyConnections(m map[connection]int, s string) bool {
	for i := 0; i < len(s)-1; i++ {
		ok := removeConnection(s[i], s[i+1], m)
		if !ok {
			return false
		}
	}
	ok := removeConnection(s[len(s)-1], s[0], m)
	if !ok {
		return false
	}
	return true
}

// What about "abacadae" and "acadabae" ?
func isRotatedString_DoesntWork(one string, two string) bool {
	if len(one) != len(two) {
		return false
	}

	oneConnections := getConnections(one)

	return verifyConnections(oneConnections, two)
}

func IsRotatedString(one string, two string) bool {
	double := fmt.Sprintf("%v%v", one, one)
	return len(one) == len(two) && strings.Contains(double, two)
}
