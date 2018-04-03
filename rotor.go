package enigmamachine

import "regexp"

type RotorSpec string

var validRotorSpec = regexp.MustCompile("^[A-Z]{26}_[A-Z]*$")

func (rs RotorSpec) Validate() error {
	return nil
}

type Rotor struct {
	next Translator
}

func NewRotor(config RotorSpec, ringSetting int, next Translator) (Rotor, error) {
	return Rotor{next: next}, nil
}

func (r *Rotor) Position() rune {
	return 'A'
}

func (r *Rotor) SetPosition(pos rune) {
}

func (r *Rotor) AdvancePosition() {
}

func (r Rotor) AtNotch() bool {
	return false
}

func (r Rotor) Translate(input rune) rune {
	return 'A'
}
