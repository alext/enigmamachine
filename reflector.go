package enigmamachine

import "fmt"

// ReflectorSpec describes the configuration of a reflector. It is a string
// containing 13 pairs of letters (optionally space separated) that the
// reflector will swap. Example:
//
//  "AE BJ CM DZ FL GY HX IV KW NR OQ PU ST"
type ReflectorSpec string

// Reflector represents the reflector component of the EnigmaMachine (German:
// Umkehrwalze, meaning 'reversal rotor'). This is the innermost component
// which swaps configured pairs of letters.
type Reflector struct {
	substitutor
}

// NewReflector constructs a new reflector that will swap letters as described
// by the given ReflectorSpec.
func NewReflector(mapping ReflectorSpec) (r Reflector, err error) {
	r = Reflector{}
	r.substitutor, err = newSwappingSubstitutor(string(mapping))
	if err == nil && len(r.substitutor) != 26 {
		err = fmt.Errorf("incomplete mapping (length: %d)", len(r.substitutor))
	}
	return r, err
}

// TranslateLetter performs a substitution of the given letter.
func (r Reflector) TranslateLetter(input rune) rune {
	return r.substitute(input)
}
