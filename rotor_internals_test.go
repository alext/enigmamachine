package enigmamachine

import "testing"

func TestRotorLetterOffset(t *testing.T) {
	examples := []struct {
		input    rune
		offset   int
		expected rune
	}{
		{
			input:    'A',
			offset:   4,
			expected: 'E',
		},
		{
			input:    'L',
			offset:   -5,
			expected: 'G',
		},
		{
			input:    'R',
			offset:   10,
			expected: 'B',
		},
		{
			input:    'B',
			offset:   -4,
			expected: 'X',
		},
		{
			input:    'R',
			offset:   36,
			expected: 'B',
		},
		{
			input:    'B',
			offset:   -30,
			expected: 'X',
		},
	}

	for _, ex := range examples {
		actual := offsetLetter(ex.input, ex.offset)
		if actual != ex.expected {
			t.Errorf("input: %c, offset: %d: want: %c, got: %c", ex.input, ex.offset, ex.expected, actual)
		}
	}
}
