package enigmamachine

// PlugboardSpec describes the setup of a plugboard. It is a string containing
// zero or more pairs of letters (optionally separated by spaces) to be swapped
// by the plugboard. Example:
//
//  "AP BR CM FZ GJ IL NT OV QS WX"
type PlugboardSpec string

// Plugboard represents the plugboard component (Steckerbrett in German) of the
// EnigmaMachine. It swaps given pairs of letters before and after the main
// rotor scrambling unit.
type Plugboard struct {
	substitutor
	next LetterTranslator
}

// NewPlugboard constructs a new plugboard. config describes the letter pairs
// to be swapped before and after they are passed to the next component.
func NewPlugboard(config PlugboardSpec, next LetterTranslator) (p Plugboard, err error) {
	p = Plugboard{next: next}
	p.substitutor, err = newSwappingSubstitutor(string(config))
	return p, err
}

// TranslateLetter translates a single letter before and after passing it to
// the next component.
func (p Plugboard) TranslateLetter(input rune) rune {
	r := p.substitute(input)
	r = p.next.TranslateLetter(r)
	return p.substitute(r)
}
