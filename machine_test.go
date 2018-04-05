package enigmamachine_test

import (
	"testing"

	em "github.com/alext/enigmamachine"
	"github.com/alext/enigmamachine/reflectors"
	"github.com/alext/enigmamachine/rotors"
)

func TestMachineSetup(t *testing.T) {
	examples := []struct {
		setup       em.MachineSetup
		expectValid bool
	}{
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.A,
				Rotors:        []em.RotorSpec{rotors.I, rotors.II, rotors.III},
				RingPositions: []int{1, 1, 1},
				Plugboard:     "AB FE",
			},
			expectValid: true,
		},
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.Bthin,
				Rotors:        []em.RotorSpec{rotors.Beta, rotors.I, rotors.II, rotors.III},
				RingPositions: []int{1, 1, 1, 1},
				Plugboard:     "AB FE",
			},
			expectValid: true,
		},
		{
			setup: em.MachineSetup{
				Reflector:     "AB CD EF GH IJ KL MN OP QR", // incomplete
				Rotors:        []em.RotorSpec{rotors.I, rotors.II, rotors.III},
				RingPositions: []int{1, 1, 1},
				Plugboard:     "AB FE",
			},
			expectValid: false,
		},
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.A,
				Rotors:        []em.RotorSpec{"ABCDEFGH", rotors.II, rotors.III}, // incomplete spec
				RingPositions: []int{1, 1, 1},
				Plugboard:     "AB FE",
			},
			expectValid: false,
		},
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.A,
				Rotors:        []em.RotorSpec{rotors.I, rotors.II, rotors.III},
				RingPositions: []int{1, 1, 27}, // invalid
				Plugboard:     "AB FE",
			},
			expectValid: false,
		},
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.A,
				Rotors:        []em.RotorSpec{rotors.I, rotors.II, rotors.III},
				RingPositions: []int{1, 1}, // missing 3rd
				Plugboard:     "AB FE",
			},
			expectValid: false,
		},
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.A,
				Rotors:        []em.RotorSpec{rotors.I, rotors.II, rotors.III},
				RingPositions: []int{1, 1}, // missing 3rd
				Plugboard:     "AB FE",
			},
			expectValid: false,
		},
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.A,
				Rotors:        []em.RotorSpec{rotors.Beta, rotors.I, rotors.II, rotors.III},
				RingPositions: []int{1, 1, 1}, // missing 4th
				Plugboard:     "AB FE",
			},
			expectValid: false,
		},
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.A,
				Rotors:        []em.RotorSpec{rotors.I, rotors.II, rotors.III},
				RingPositions: []int{1, 1, 1},
				Plugboard:     "AB FB", // conflicting
			},
			expectValid: false,
		},
		{
			setup: em.MachineSetup{
				Reflector:     reflectors.A,
				Rotors:        []em.RotorSpec{rotors.I, rotors.II}, // only 2
				RingPositions: []int{1, 1},
				Plugboard:     "AB FE",
			},
			expectValid: false,
		},
	}
	for i, ex := range examples {
		_, err := em.New(ex.setup)
		if ex.expectValid && err != nil {
			t.Errorf("[%d] expected valid, got error: %s", i, err.Error())
		}
		if !ex.expectValid && err == nil {
			t.Errorf("[%d] expected invalid, got no error", i)
		}
	}
}
