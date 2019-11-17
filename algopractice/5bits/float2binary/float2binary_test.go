package float2binary

import "testing"

func TestFloatToBinary(t *testing.T) {
	tests := []struct {
		f      float64
		binary string
	}{
		{.72, "ERROR"},
		{.5, ".1"},
		{.25, ".01"},
		{.75, ".11"},
	}
	for _, test := range tests {
		got := floatToBinary(test.f)
		if got != test.binary {
			t.Errorf("floatToBinary(%v) error got %v wanted %v", test.f, got, test.binary)
		}
	}
}
