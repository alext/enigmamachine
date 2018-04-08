package enigmamachine

import (
	"fmt"
	"regexp"
	"strings"
)

// RotorSpec describes the mapping and notches of a rotor. It is a string
// consisting of 26 letters describing the forward mapping, followed by an
// underscore, and any notch positions. Example:
//
//  "EKMFLGDQVZNTOWYHXUSPAIBRCJ_Q"
type RotorSpec string

var validRotorSpec = regexp.MustCompile("^[A-Z]{26}(?:_[A-Z]*)?$")

// Validate tests whether a RotorSpec is well-formed.
func (rs RotorSpec) Validate() error {
	if !validRotorSpec.MatchString(string(rs)) {
		return fmt.Errorf("invalid rotor spec")
	}
	return nil
}

func (rs RotorSpec) split() (string, string) {
	parts := strings.SplitN(string(rs), "_", 2)
	if len(parts) != 2 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

func (rs RotorSpec) parse() (forward, reverse substitutor, notches string, err error) {
	err = rs.Validate()
	if err != nil {
		return nil, nil, "", err
	}
	mapping, notches := rs.split()
	forward, reverse, err = newBidirectionalSubstitutors(mapping)
	return forward, reverse, notches, err
}

// Rotor represents a single rotor (Walzen in German) in the EnigmaMachine.
type Rotor struct {
	next           LetterTranslator
	forward        substitutor
	reverse        substitutor
	notches        string
	ringOffset     int
	positionOffset int
}

// NewRotor constructs a new rotor instance with the mapping and notches
// described by the given RotorSpec, and the given ringSetting. The next
// component in the sequence (either another rotor, or the reflector) is given
// in the next param.
func NewRotor(spec RotorSpec, ringSetting int, next LetterTranslator) (*Rotor, error) {
	forward, reverse, notches, err := spec.parse()
	if err != nil {
		return nil, err
	}
	if ringSetting < 1 || ringSetting > 26 {
		return nil, fmt.Errorf("Invalid ring setting %d", ringSetting)
	}
	r := &Rotor{
		next:           next,
		forward:        forward,
		reverse:        reverse,
		notches:        notches,
		ringOffset:     ringSetting - 1,
		positionOffset: 0,
	}
	return r, nil
}

// Position returns the current position of the rotor.
func (r *Rotor) Position() rune {
	return rune('A' + r.positionOffset)
}

// SetPosition sets the position of the rotor.
func (r *Rotor) SetPosition(pos rune) {
	r.positionOffset = int(pos - 'A')
}

// AdvancePosition advances the rotor position by a single step.
func (r *Rotor) AdvancePosition() {
	r.positionOffset = (r.positionOffset + 1) % 26
}

// AtNotch returns whether the rotor is currently in a notch position.
func (r *Rotor) AtNotch() bool {
	return strings.ContainsRune(r.notches, r.Position())
}

// TranslateLetter performs a forward substitution on the given letter, passes
// it to the next component, and then performs the reverse substitution on the
// result.
func (r *Rotor) TranslateLetter(input rune) rune {
	c := r.substitute(input, r.forward)
	c = r.next.TranslateLetter(c)
	c = r.substitute(c, r.reverse)
	return c
}

func (r *Rotor) substitute(input rune, sub substitutor) rune {
	input = offsetLetter(input, r.positionOffset-r.ringOffset)
	res := sub.substitute(input)
	res = offsetLetter(res, r.ringOffset-r.positionOffset)
	return res
}

func offsetLetter(letter rune, offset int) rune {
	offset = (int(letter) - 'A' + offset) % 26
	if offset < 0 {
		offset += 26
	}
	return rune('A' + offset)
}
