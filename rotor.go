package enigmamachine

import (
	"fmt"
	"regexp"
	"strings"
)

type RotorSpec string

var validRotorSpec = regexp.MustCompile("^[A-Z]{26}(?:_[A-Z]*)?$")

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
	forward, reverse = make(substitutor), make(substitutor)
	for i, r := range mapping {
		if _, found := reverse[r]; found {
			return nil, nil, "", fmt.Errorf("Duplicate entry '%c' in mapping", r)
		}
		c := rune('A' + i)
		forward[c] = r
		reverse[r] = c
	}
	return forward, reverse, notches, nil
}

type Rotor struct {
	next           Translator
	forward        substitutor
	reverse        substitutor
	notches        string
	positionOffset int
}

func NewRotor(spec RotorSpec, ringSetting int, next Translator) (*Rotor, error) {
	forward, reverse, notches, err := spec.parse()
	if err != nil {
		return nil, err
	}
	r := &Rotor{
		next:           next,
		forward:        forward,
		reverse:        reverse,
		notches:        notches,
		positionOffset: 0,
	}
	return r, nil
}

func (r *Rotor) Position() rune {
	return rune('A' + r.positionOffset)
}

func (r *Rotor) SetPosition(pos rune) {
	r.positionOffset = int(pos - 'A')
}

func (r *Rotor) AdvancePosition() {
	r.positionOffset = (r.positionOffset + 1) % 26
}

func (r *Rotor) AtNotch() bool {
	return strings.ContainsRune(r.notches, r.Position())
}

func (r *Rotor) Translate(input rune) rune {
	c := r.substitute(input, r.forward)
	c = r.next.Translate(c)
	c = r.substitute(c, r.reverse)
	return c
}

func (r *Rotor) substitute(input rune, sub substitutor) rune {
	input = offsetLetter(input, r.positionOffset)
	res := sub.substitute(input)
	res = offsetLetter(res, -r.positionOffset)
	return res
}

func offsetLetter(letter rune, offset int) rune {
	offset = (int(letter) - 'A' + offset) % 26
	if offset < 0 {
		offset += 26
	}
	return rune('A' + offset)
}
