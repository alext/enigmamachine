package enigmamachine_test

import (
	"testing"

	em "github.com/alext/enigmamachine"
)

func TestRotorTranslateLetter(t *testing.T) {
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
			input:        'R',
			innerExpects: 'B',
			innerReturns: 'F',
			expected:     'W',
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

	// 1 2 3 4 5 6 7 8 9 1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2 2
	//                   0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6
	// A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	// E K M F L G D Q V Z N T O W Y H X U S P A I B R C J
	for _, ex := range examples {
		inner := &stubLetterTranslator{translateReturns: ex.innerReturns}
		rotor, err := em.NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ_R", ex.ringSetting, inner)
		if err != nil {
			t.Errorf("[ring: %d, pos: %c] Unexpected error creating rotor: %s", ex.ringSetting, ex.position, err.Error())
			continue
		}
		rotor.SetPosition(ex.position)
		actual := rotor.TranslateLetter(ex.input)
		if ex.innerExpects != 0 && ex.innerExpects != inner.translateParam {
			t.Errorf("ring: %d, pos: %c, input: %c: inner wants: %c, got: %c", ex.ringSetting, ex.position, ex.input, ex.innerExpects, inner.translateParam)
		}
		if actual != ex.expected {
			t.Errorf("ring: %d, pos: %c, input: %c: want: %c, got: %c", ex.ringSetting, ex.position, ex.input, ex.expected, actual)
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
		"EKEFLGDQVZNTOWYHXUSPAIBRCJ_N", // Duplicate entry
	}
	for i, spec := range errorExamples {
		_, err := em.NewRotor(spec, 1, nil)
		if err == nil {
			t.Errorf("[%d] spec: %s, want error, got no error", i, spec)
			continue
		}
	}

	validExamples := []em.RotorSpec{
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ_N",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ_",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	}
	for i, spec := range validExamples {
		_, err := em.NewRotor(spec, 1, nil)
		if err != nil {
			t.Errorf("[%d] spec: %s, expected no error, got: '%s'", i, spec, err.Error())
		}
	}

	ringSettingExamples := map[int]bool{
		-1: false,
		0:  false,
		1:  true,
		5:  true,
		26: true,
		27: false,
	}
	for ringSetting, expectValid := range ringSettingExamples {
		_, err := em.NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ_N", ringSetting, nil)
		if expectValid && err != nil {
			t.Errorf("[ring %d] expected valid, got error: %s", ringSetting, err.Error())
		}
		if !expectValid && err == nil {
			t.Errorf("[ring %d] expected invalid, got no error", ringSetting)
		}
	}
}
