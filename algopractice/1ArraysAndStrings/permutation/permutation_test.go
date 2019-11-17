package permutation

import "testing"

func TestIsPermutationTrue(t *testing.T) {
	if IsPermutation("abcde", "eabdc") == false {
		t.Error("correct permutation returned false")
	}
}

func TestIsPermutationFalse(t *testing.T) {
	if IsPermutation("abcdc", "eabdc") == true {
		t.Error("not permutation returned true")
	}
}

func TestIsPermutationDifferentLengths(t *testing.T) {
	if IsPermutation("abcdee", "eabdc") == true {
		t.Error("different lengths returned true")
	}
}

func TestIsPermutationDuplicates(t *testing.T) {
	if IsPermutation("aaabbbccc", "abcabcabc") == false {
		t.Error("≈valid permutation with duplicates returned false")
	}
}

func TestIsPermutationDuplicatesFalse(t *testing.T) {
	if IsPermutation("aaaabb", "bbbbaa") == true {
		t.Error("invalid permutation with duplicates returned true")
	}
}