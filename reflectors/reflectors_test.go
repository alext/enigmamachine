package reflectors_test

import (
	"testing"

	"github.com/alext/enigmamachine"
	"github.com/alext/enigmamachine/reflectors"
)

func TestReflectors(t *testing.T) {
	examples := []struct {
		reflector enigmamachine.Reflector
		inputs    []rune
		expected  []rune
	}{
		{
			reflector: reflectors.A,
			inputs:    []rune{'A', 'F', 'Q'},
			expected:  []rune{'E', 'L', 'O'},
		},
		{
			reflector: reflectors.B,
			inputs:    []rune{'A', 'F', 'Q'},
			expected:  []rune{'Y', 'S', 'E'},
		},
		{
			reflector: reflectors.C,
			inputs:    []rune{'A', 'K', 'Q'},
			expected:  []rune{'F', 'R', 'T'},
		},
		{
			reflector: reflectors.Bthin,
			inputs:    []rune{'A', 'F', 'P'},
			expected:  []rune{'E', 'U', 'M'},
		},
		{
			reflector: reflectors.Cthin,
			inputs:    []rune{'A', 'K', 'Q'},
			expected:  []rune{'R', 'H', 'Z'},
		},
	}
	for i, ex := range examples {
		for j, input := range ex.inputs {
			actual := ex.reflector.Translate(input)
			if actual != ex.expected[j] {
				t.Errorf("[%d] input: %c, want %c, got: %c", i, input, ex.expected[j], actual)
			}
		}
	}
}
