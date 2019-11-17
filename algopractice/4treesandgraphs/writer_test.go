package treesandgraphs

import "testing"

func TestWriter(t *testing.T) {
	w := &writer{}
	w.store(1)
	w.store(2)
	w.store(3)
	expected := `1
2 3
`
	if expected != w.String() {
		t.Errorf("writer error wanted \"%v\" got \"%v\"", expected, w.String())
	}
}

func TestWriterFour(t *testing.T) {
	w := &writer{}
	w.store(1)
	w.store(2)
	w.store(3)
	w.store(4)
	w.store(5)
	w.store(6)
	w.store(7)
	expected := `1
2 3
4 5 6 7
`
	if expected != w.String() {
		t.Errorf("writer error wanted \"%v\" got \"%v\"", expected, w.String())
	}
}

func TestIsEndOfLine(t *testing.T) {
	tests := []struct {
		i      int
		n      int
		wanted bool
	}{
		{0, 1, true},
		{1, 2, true},
		{1, 3, false},
		{2, 3, true},
		{3, 4, true},
		{3, 5, false},
		{3, 6, false},
		{6, 7, true},
	}

	for _, test := range tests {
		got := isEndOfLine(test.i, test.n)
		if got != test.wanted {
			t.Errorf("isEndOfLine(%v, %v) error got %v wanted %v", test.i, test.n, got, test.wanted)
		}
	}
}
