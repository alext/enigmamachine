package enigmamachine_test

import (
	"testing"

	em "github.com/alext/enigmamachine"
)

func TestRotorTranslate(t *testing.T) {
	examples := []struct {
		ringSetting  int
		position     rune
		input        rune
		innerExpects rune
		innerReturns rune
		expected     rune
	}{
		{
			ringSetting:  1,
			position:     'A',
			input:        'B',
			innerExpects: 'K',
			innerReturns: 'L',
			expected:     'E',
		},
		{
			ringSetting:  5,
			position:     'A',
			input:        'B',
			innerExpects: 'V',
			innerReturns: 'F',
			expected:     'A',
		},
		{
			ringSetting:  1,
			position:     'L',
			input:        'B',
			innerExpects: 'D',
			innerReturns: 'L',
			expected:     'C',
		},
		{
			ringSetting:  5,
			position:     'T',
			input:        'B',
			innerExpects: 'I',
			innerReturns: 'L',
			expected:     'F',
		},
	}

	for i, ex := range examples {
		inner := &stubTranslator{translateReturns: ex.innerReturns}
		rotor, err := em.NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ_R", ex.ringSetting, inner)
		if err != nil {
			t.Errorf("[%d] Unexpected error creating rotor: %s", i, err.Error())
			continue
		}
		rotor.SetPosition(ex.position)
		actual := rotor.Translate(ex.input)
		if ex.innerExpects != 0 && ex.innerExpects != inner.translateParam {
			t.Errorf("[%d] input: %c, inner wants: %c, got: %c", i, ex.input, ex.innerExpects, inner.translateParam)
		}
		if actual != ex.expected {
			t.Errorf("[%d] input: %c, want: %c, got: %c", i, ex.input, ex.expected, actual)
		}
	}
}

func TestRotorPosition(t *testing.T) {
	rotor, err := em.NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ_DR", 1, nil)
	if err != nil {
		t.Fatalf("Unexpected error creating rotor: %s", err.Error())
	}
	if rotor.Position() != 'A' {
		t.Errorf("Default position, want: A, got: %c", rotor.Position())
	}
	rotor.SetPosition('G')
	if rotor.Position() != 'G' {
		t.Errorf("Set position, want: G, got: %c", rotor.Position())
	}
	rotor.AdvancePosition()
	if rotor.Position() != 'H' {
		t.Errorf("Advance position, want: H, got: %c", rotor.Position())
	}
	rotor.SetPosition('Z')
	rotor.AdvancePosition()
	if rotor.Position() != 'A' {
		t.Errorf("Advance position wraparound, want: A, got: %c", rotor.Position())
	}
}

func TestRotorAtNotch(t *testing.T) {
	examples := []struct {
		spec        em.RotorSpec
		ringSetting int
		expected    []rune
		notExpected []rune
	}{
		{
			spec:        "EKMFLGDQVZNTOWYHXUSPAIBRCJ_DR",
			expected:    []rune{'D', 'R'},
			notExpected: []rune{'C', 'E', 'S'},
		},
		{
			spec:        "EKMFLGDQVZNTOWYHXUSPAIBRCJ_DR",
			ringSetting: 5, // Should make no difference
			expected:    []rune{'D', 'R'},
			notExpected: []rune{'C', 'E', 'S'},
		},
	}

	for i, ex := range examples {
		if ex.ringSetting == 0 {
			ex.ringSetting = 1
		}
		rotor, err := em.NewRotor(ex.spec, ex.ringSetting, nil)
		if err != nil {
			t.Errorf("[%d] Unexpected error creating rotor: %s", i, err.Error())
			continue
		}
		for _, r := range ex.expected {
			rotor.SetPosition(r)
			if !rotor.AtNotch() {
				t.Errorf("[%d] spec: '%s', position: %c, expected AtNotch to return true", i, ex.spec, r)
			}
		}
		for _, r := range ex.notExpected {
			rotor.SetPosition(r)
			if rotor.AtNotch() {
				t.Errorf("[%d] spec: '%s', position: %c, expected AtNotch to return false", i, ex.spec, r)
			}
		}
	}
}

func TestRotorSetup(t *testing.T) {
	errorExamples := []em.RotorSpec{
		"eKMFLGDQVZNTOWYHXUSPAIBRCJ",   // invalid char
		"EKMFLGDQVZNTOWYHXUSPAIBRC",    // incomplete
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ_z", // invalid notch
		"EKEFLGDQVZNTOWYHXUSPAIBRCJ",   // Duplicate entry
	}
	for i, spec := range errorExamples {
		_, err := em.NewRotor(spec, 1, nil)
		if err == nil {
			t.Errorf("[%d] spec: %s, want error, got no error", i, spec)
			continue
		}
	}

	_, err := em.NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ_N", 1, nil)
	if err != nil {
		t.Errorf("good config, expected no error, got: '%s'", err.Error())
	}
}
