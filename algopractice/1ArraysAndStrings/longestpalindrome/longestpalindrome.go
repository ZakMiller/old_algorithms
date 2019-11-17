package longestpalindrome

import (
	"container/list"
)

type indices struct {
	start, end int
}

func addSinglesAndDoubles(s string, m map[indices]struct{}) indices {
	longest := indices{0, 0}
	for i := 0; i < len(s); i++ {
		m[indices{i, i}] = struct{}{}
		if i < len(s)-1 {
			if s[i] == s[i+1] {
				m[indices{i, i + 1}] = struct{}{}
				longest = indices{i, i + 1}
			}
		}
	}
	return longest
}

func longestPalindromeOld(s string) string {
	if len(s) == 0 {
		return ""
	}
	isP := make(map[indices]struct{})
	longest := addSinglesAndDoubles(s, isP)

	for strLen := 3; strLen <= len(s); strLen++ {
		for i := 0; i <= len(s)-strLen; i++ {
			start := i
			end := i + strLen - 1
			if _, ok := isP[indices{start + 1, end - 1}]; ok && s[start] == s[end] {
				isP[indices{start, end}] = struct{}{}
				longest = indices{start, end}
			}
		}
	}

	return s[longest.start : longest.end+1]
}

func palindromes(s string) *list.List {
	l := list.New()
	for i := 0; i < len(s); i++ {
		l.PushBack(&indices{i, i})
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			l.PushBack(&indices{i, i + 1})
		}
	}
	return l
}
func longestPalindromeNew(s string) string {
	if len(s) == 0 {
		return ""
	}
	var longest *indices
	q := palindromes(s)
	for q.Len() > 0 {
		element := q.Front()
		q.Remove(element)
		i := element.Value.(*indices)
		longest = i
		if i.start > 0 && i.end < len(s)-1 && s[i.start-1] == s[i.end+1] {
			q.PushBack(&indices{i.start - 1, i.end + 1})
		}
	}
	return s[longest.start : longest.end+1]
}

func palindromesSlice(s string) []*indices {
	l := make([]*indices, 0)
	for i := 0; i < len(s); i++ {
		l = append(l, &indices{i, i})
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			l = append(l, &indices{i, i + 1})
		}
	}
	return l
}
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var longest *indices
	q := palindromesSlice(s)
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		longest = i
		if i.start > 0 && i.end < len(s)-1 && s[i.start-1] == s[i.end+1] {
			q = append(q, &indices{i.start - 1, i.end + 1})
		}
	}
	return s[longest.start : longest.end+1]
}
