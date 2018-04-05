package enigmamachine

import "fmt"

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

func (m *Machine) SetPositions(positions []rune) {}

func (m Machine) Translate(input string) string {
	return ""
}
