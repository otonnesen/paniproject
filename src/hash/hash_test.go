package hash

import "testing"

func TestHashFromNumber(t *testing.T) {
	tests := []struct {
		id       int
		expected string
	}{
		{
			id:       0,
			expected: "000000",
		},
		{
			id:       1,
			expected: "000001",
		},
		{
			id:       2,
			expected: "000002",
		},
		{
			id:       9,
			expected: "000009",
		},
		{
			id:       10,
			expected: "00000a",
		},
		{
			id:       35,
			expected: "00000z",
		},
		{
			id:       36,
			expected: "000010",
		},
		{
			id:       1296,
			expected: "000100",
		},
		{
			id:       46656,
			expected: "001000",
		},
		{
			id:       1679616,
			expected: "010000",
		},
		{
			id:       1679616 + 1,
			expected: "010001",
		},
		{
			id:       60466176,
			expected: "100000",
		},
		{
			id:       2176782336 - 1,
			expected: "zzzzzz",
		},
		{
			id:       1234567890,
			expected: "kf12oi",
		},
		{
			id:       987654321,
			expected: "gc0uy9",
		},
	}

	for _, i := range tests {
		if actual, _ := HashFromNumber(i.id); actual != i.expected {
			t.Errorf("Incorrect hash for id %v. Expected '%v', got '%v'.", i.id, i.expected, actual)
		}
	}
}

func TestNumberFromHash(t *testing.T) {
	tests := []struct {
		hash     string
		expected int
	}{
		{
			hash:     "000000",
			expected: 0,
		},
		{
			hash:     "000001",
			expected: 1,
		},
		{
			hash:     "000002",
			expected: 2,
		},
		{
			hash:     "000009",
			expected: 9,
		},
		{
			hash:     "00000a",
			expected: 10,
		},
		{
			hash:     "00000z",
			expected: 35,
		},
		{
			hash:     "000010",
			expected: 36,
		},
		{
			hash:     "000100",
			expected: 1296,
		},
		{
			hash:     "001000",
			expected: 46656,
		},
		{
			hash:     "010000",
			expected: 1679616,
		},
		{
			hash:     "010001",
			expected: 1679616 + 1,
		},
		{
			hash:     "100000",
			expected: 60466176,
		},
		{
			hash:     "zzzzzz",
			expected: 2176782336 - 1,
		},
		{
			hash:     "kf12oi",
			expected: 1234567890,
		},
		{
			hash:     "gc0uy9",
			expected: 987654321,
		},
	}

	for _, i := range tests {
		if actual, _ := NumberFromHash(i.hash); actual != i.expected {
			t.Errorf("Incorrect id for hash '%v'. Expected %v, got %v.", i.hash, i.expected, actual)
		}
	}
}
