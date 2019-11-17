package unique

import "testing"

func TestIsUniqueTrue(t *testing.T) {
	if IsUnique("a") == false {
		t.Error("IsUnique returned false")
	}
}

func TestIsUniqueTrueLonger(t *testing.T) {
	if IsUnique("abcdefg") == false {
		t.Error("IsUnique returned false")
	}
}

func TestIsUniqueFalse(t *testing.T) {
	if IsUnique("aa") == true {
		t.Error("IsUnique returned true")
	}
}
