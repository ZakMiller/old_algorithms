package longestpalindrome

import "testing"

func TestLongestPalindromeOld(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"ababa", "ababa"},
		{"abaqbb", "aba"},
		{"qwpratytfpm", "tyt"},
		{"", ""},
		{"p", "p"},
		{"qmdd", "dd"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaacd", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		{"apoiawjfoungakmefoiawefpiowayepoiawfmkmlfajwesfpoaiwefioputetetaoiejfwaopyyyyyy", "yyyyyy"},
	}

	for _, test := range tests {
		got := longestPalindromeOld(test.input)
		if got != test.expected {
			t.Errorf("longestPalidnromeOld(%v) error got %v expected %v", test.input, got, test.expected)
		}
	}
}

func TestLongestPalindromeNew(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"ababa", "ababa"},
		{"abaqbb", "aba"},
		{"qwpratytfpm", "tyt"},
		{"", ""},
		{"p", "p"},
		{"qmdd", "dd"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaacd", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		{"apoiawjfoungakmefoiawefpiowayepoiawfmkmlfajwesfpoaiwefioputetetaoiejfwaopyyyyyy", "yyyyyy"},
	}

	for _, test := range tests {
		got := longestPalindromeNew(test.input)
		if got != test.expected {
			t.Errorf("longestPalidnromeNew(%v) error got %v expected %v", test.input, got, test.expected)
		}
	}
}

func BenchmarkLongestOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tests := []struct {
			input, expected string
		}{
			{"ababa", "ababa"},
			{"abaqbb", "aba"},
			{"qwpratytfpm", "tyt"},
			{"", ""},
			{"p", "p"},
			{"qmdd", "dd"},
			{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaacd", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			{"apoiawjfoungakmefoiawefpiowayepoiawfmkmlfajwesfpoaiwefioputetetaoiejfwaopyyyyyy", "yyyyyy"},
		}
		for _, test := range tests {
			longestPalindromeOld(test.input)
		}
	}
}

func BenchmarkLongestNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tests := []struct {
			input, expected string
		}{
			{"ababa", "ababa"},
			{"abaqbb", "aba"},
			{"qwpratytfpm", "tyt"},
			{"", ""},
			{"p", "p"},
			{"qmdd", "dd"},
			{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaacd", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			{"apoiawjfoungakmefoiawefpiowayepoiawfmkmlfajwesfpoaiwefioputetetaoiejfwaopyyyyyy", "yyyyyy"},
		}
		for _, test := range tests {
			longestPalindromeNew(test.input)
		}
	}
}

func BenchmarkLongestNewSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tests := []struct {
			input, expected string
		}{
			{"ababa", "ababa"},
			{"abaqbb", "aba"},
			{"qwpratytfpm", "tyt"},
			{"", ""},
			{"p", "p"},
			{"qmdd", "dd"},
			{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaacd", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			{"apoiawjfoungakmefoiawefpiowayepoiawfmkmlfajwesfpoaiwefioputetetaoiejfwaopyyyyyy", "yyyyyy"},
		}
		for _, test := range tests {
			longestPalindrome(test.input)
		}
	}
}
