package enigmamachine_test

import (
	"strings"
	"testing"

	em "github.com/alext/enigmamachine"
)

func TestReflectorTranslate(t *testing.T) {
	examples := []struct {
		input    rune
		expected rune
	}{
		{input: 'B', expected: 'R'},
		{input: 'L', expected: 'G'},
	}

	reflector, err := em.NewReflector("AY BR CU DH EQ FS GL IP JX KN MO TZ VW")
	if err != nil {
		t.Fatalf("Unexpected error creating reflector: %s", err.Error())
	}
	for i, ex := range examples {
		actual := reflector.Translate(ex.input)
		if actual != ex.expected {
			t.Errorf("[%d] input: %c, want: %c, got: %c", i, ex.input, ex.expected, actual)
		}
	}
}

func TestReflectorSetup(t *testing.T) {
	errorExamples := []struct {
		config         string
		errorSubstring string
	}{
		{
			config:         "AY BR cu DH EQ FS GL IP JX KN MO TZ VW",
			errorSubstring: "invalid character",
		},
		{
			config:         "AY BR CB DH EQ FS GL IP JX KN MO TZ VW",
			errorSubstring: "duplicate character",
		},
		{
			config:         "AY BR CU DH EQ FS GL IP JX KN MO TZ",
			errorSubstring: "incomplete mapping",
		},
	}
	for i, ex := range errorExamples {
		_, err := em.NewReflector(ex.config)
		if err == nil {
			t.Errorf("[%d] config: %s, want error containing '%s', got no error", i, ex.config, ex.errorSubstring)
			continue
		}
		if !strings.Contains(err.Error(), ex.errorSubstring) {
			t.Errorf("[%d] config: %s, want error containing '%s', got: '%s'", i, ex.config, ex.errorSubstring, err.Error())
		}
	}

	_, err := em.NewReflector("AB CD EF GH IJ KL MN OP QR ST UV WX YZ")
	if err != nil {
		t.Errorf("good config, expected no error, got: '%s'", err.Error())
	}
}
