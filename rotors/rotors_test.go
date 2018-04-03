package rotors_test

import (
	"testing"

	"github.com/alext/enigmamachine"
	"github.com/alext/enigmamachine/rotors"
)

func TestRotors(t *testing.T) {
	specs := []enigmamachine.RotorSpec{
		rotors.I,
		rotors.II,
		rotors.III,
		rotors.IV,
		rotors.V,
		rotors.VI,
		rotors.VII,
		rotors.VIII,
		rotors.Beta,
		rotors.Gamma,
	}

	for i, spec := range specs {
		_, err := enigmamachine.NewRotor(spec, 1, nil)
		if err != nil {
			t.Errorf("[%d] spec: %s, expected no error, got: %s", i, spec, err.Error())
		}
	}
}
