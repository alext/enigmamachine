package enigmamachine_test

import (
	"strings"
	"testing"

	em "github.com/alext/enigmamachine"
)

type stubLetterTranslator struct {
	translateParam   rune
	translateReturns rune
}

func (t *stubLetterTranslator) TranslateLetter(input rune) rune {
	t.translateParam = input
	if t.translateReturns != 0 {
		return t.translateReturns
	}
	return input
}

func TestPlugboardTranslateLetter(t *testing.T) {
	examples := []struct {
		input        rune
		innerExpects rune
		innerReturns rune
		expected     rune
	}{
		{
			input:        'A',
			innerExpects: 'F',
			innerReturns: 'G',
			expected:     'D',
		},
		{
			input:        'X',
			innerExpects: 'E',
			innerReturns: 'D',
			expected:     'G',
		},
		{
			input:        'B',
			innerExpects: 'B',
			innerReturns: 'H',
			expected:     'H',
		},
	}

	inner := &stubLetterTranslator{}
	plugboard, err := em.NewPlugboard("AF DG EX", inner)
	if err != nil {
		t.Fatalf("Unexpected error creating plugboard: %s", err.Error())
	}
	for i, ex := range examples {
		inner.translateReturns = ex.innerReturns
		actual := plugboard.TranslateLetter(ex.input)
		if ex.innerExpects != 0 && ex.innerExpects != inner.translateParam {
			t.Errorf("[%d] input: %c, inner wants: %c, got: %c", i, ex.input, ex.innerExpects, inner.translateParam)
		}
		if actual != ex.expected {
			t.Errorf("[%d] input: %c, want: %c, got: %c", i, ex.input, ex.expected, actual)
		}
	}
}

func TestPlugboardSetup(t *testing.T) {
	errorExamples := []struct {
		config         em.PlugboardSpec
		errorSubstring string
	}{
		{
			config:         "AY BR cu",
			errorSubstring: "invalid character",
		},
		{
			config:         "AB CD GC",
			errorSubstring: "duplicate character",
		},
		{
			config:         "AB CD EF G",
			errorSubstring: "unpaired character",
		},
	}
	for i, ex := range errorExamples {
		_, err := em.NewPlugboard(ex.config, nil)
		if err == nil {
			t.Errorf("[%d] config: %s, want error containing '%s', got no error", i, ex.config, ex.errorSubstring)
			continue
		}
		if !strings.Contains(err.Error(), ex.errorSubstring) {
			t.Errorf("[%d] config: %s, want error containing '%s', got: '%s'", i, ex.config, ex.errorSubstring, err.Error())
		}
	}

	_, err := em.NewPlugboard("AB CD EF", nil)
	if err != nil {
		t.Errorf("good config, expected no error, got: '%s'", err.Error())
	}
}
