package linkedlist

import (
	"strconv"
)

func Sum(one *LinkedList, two *LinkedList) *LinkedList {
	oneString := one.String()
	twoString := two.String()
	oneInt, _ := strconv.Atoi(oneString)
	twoInt, _ := strconv.Atoi(twoString)
	returnString := strconv.Itoa(oneInt + twoInt)
	returnRunes := []rune(returnString)
	returnArray := make([]string, len(returnRunes))
	for i, r := range returnRunes {
		returnArray[i] = string(r)
	}
	return createLinkedList(returnArray...)
}

// Assumes no UTF8, should just be numbers.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func SumReversed(one *LinkedList, two *LinkedList) *LinkedList {
	oneString := reverse(one.String())
	twoString := reverse(two.String())
	oneInt, _ := strconv.Atoi(oneString)
	twoInt, _ := strconv.Atoi(twoString)
	returnString := reverse(strconv.Itoa(oneInt + twoInt))
	returnRunes := []rune(returnString)
	returnArray := make([]string, len(returnRunes))
	for i, r := range returnRunes {
		returnArray[i] = string(r)
	}
	return createLinkedList(returnArray...)
}
