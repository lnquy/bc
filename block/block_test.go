package block

import "testing"

func TestIsValidHash(t *testing.T) {
	tcs := []struct {
		hash   []byte
		result bool
	}{
		{[]byte{0, 0, 0}, true},
		{[]byte{0, 0, 1}, true},
		{[]byte{0, 0, 128}, true},

		{[]byte{128, 0, 0}, false},
		{[]byte{64, 0, 0}, false},
		{[]byte{32, 0, 0}, false},
		{[]byte{16, 0, 0}, false},
		{[]byte{8, 0, 0}, false},
		{[]byte{4, 0, 0}, false},
		{[]byte{2, 0, 0}, false},
		{[]byte{1, 0, 0}, false},

		{[]byte{0, 128, 0}, false},
		{[]byte{0, 64, 0}, false},
		{[]byte{0, 32, 0}, false},
		{[]byte{0, 16, 0}, false},
		{[]byte{0, 8, 0}, false},
		{[]byte{0, 4, 0}, false},
		{[]byte{0, 2, 0}, false},
		{[]byte{0, 1, 0}, false},

		{[]byte{1, 1, 0}, false},
		{[]byte{128, 128, 0}, false},
	}

	for idx, tc := range tcs {
		if out := isValidHash(tc.hash); out != tc.result {
			t.Errorf("isValidHash #%d: expected %v, got %v", idx, tc.result, out)
		}
	}
}
