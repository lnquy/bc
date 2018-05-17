package block

import "testing"

func TestIsValidHash(t *testing.T) {
	tcs := []struct {
		hash   [32]byte
		result bool
	}{
		{[32]byte{0, 0, 1}, true},
		{[32]byte{0, 0, 0}, true},
		{[32]byte{0, 1, 0}, false},
		{[32]byte{1, 0, 0}, false},
		{[32]byte{2, 1, 0}, false},
	}

	for idx, tc := range tcs {
		if out := isValidHash(tc.hash); out != tc.result {
			t.Errorf("isValidHash #%d: expected %v, got %v", idx, tc.result, out)
		}
	}
}
