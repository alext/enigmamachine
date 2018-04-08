package enigmamachine

import (
	"fmt"
	"strings"
)

// MachineSetup describes the setup of an Enigma Machine. This describes the
// Reflector, Rotors and Plugboard setup. This passed to New to describe the
// setup of a new Machine instance.
type MachineSetup struct {
	Reflector     ReflectorSpec
	Rotors        []RotorSpec
	RingPositions []int
	Plugboard     PlugboardSpec
}

// Machine represents an Enigma Machine.
type Machine struct {
	rotors    []*Rotor
	plugboard Plugboard
}

// New constructs an enigma machine with given MachineSetup. err will be
// non-nil if the given configuration is invalid.
func New(s MachineSetup) (m *Machine, err error) {
	m = &Machine{}
	t, err := NewReflector(s.Reflector)
	if err != nil {
		return nil, fmt.Errorf("Reflector error: %s", err.Error())
	}
	var next LetterTranslator = t
	if len(s.Rotors) < 3 {
		return nil, fmt.Errorf("Minimum 3 rotors required, have %d", len(s.Rotors))
	}
	if len(s.Rotors) != len(s.RingPositions) {
		return nil, fmt.Errorf("RingPosition mismatch: %d rotors and %d ring positions", len(s.Rotors), len(s.RingPositions))
	}
	for i, rs := range s.Rotors {
		r, err := NewRotor(rs, s.RingPositions[i], next)
		if err != nil {
			return nil, fmt.Errorf("Rotor error: %s", err.Error())
		}
		m.rotors = append(m.rotors, r)
		next = r
	}
	m.plugboard, err = NewPlugboard(s.Plugboard, next)
	if err != nil {
		return nil, fmt.Errorf("Plugboard error: %s", err.Error())
	}
	return m, nil
}

// SetPositions sets the positions of the rotors to the given letters.
func (m *Machine) SetPositions(positions []rune) {
	for i, pos := range positions {
		if i >= len(m.rotors) {
			return
		}
		m.rotors[i].SetPosition(pos)
	}
}

func (m *Machine) advanceRotors() {
	l := len(m.rotors)
	if m.rotors[l-2].AtNotch() {
		m.rotors[l-3].AdvancePosition()
	}
	if m.rotors[l-1].AtNotch() || m.rotors[l-2].AtNotch() {
		m.rotors[l-2].AdvancePosition()
	}
	m.rotors[l-1].AdvancePosition()
}

// TranslateLetter performs a translation (encryption or decryption) of a
// single letter. This effectively simulates the pressing of a single key on
// the Enigma Machine.
func (m *Machine) TranslateLetter(input rune) rune {
	if input < 'A' || input > 'Z' {
		return input
	}
	m.advanceRotors()
	return m.plugboard.TranslateLetter(input)
}

// TranslateString runs the given string through the EnigmaMachine and returns
// the result. Any characters in the input that are not the uppercase letters
// 'A' - 'Z' are returned in the output unchanged.
func (m *Machine) TranslateString(input string) (string, error) {
	var out strings.Builder
	for _, ch := range input {
		_, err := out.WriteRune(m.TranslateLetter(ch))
		if err != nil {
			return "", err
		}
	}
	return out.String(), nil
}
