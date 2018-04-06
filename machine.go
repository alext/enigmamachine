package enigmamachine

import (
	"fmt"
	"strings"
)

type MachineSetup struct {
	Reflector     ReflectorSpec
	Rotors        []RotorSpec
	RingPositions []int
	Plugboard     PlugboardSpec
}

type Machine struct {
	rotors    []*Rotor
	plugboard Plugboard
}

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

func (m *Machine) TranslateLetter(input rune) rune {
	if input < 'A' || input > 'Z' {
		return input
	}
	m.advanceRotors()
	return m.plugboard.TranslateLetter(input)
}

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
