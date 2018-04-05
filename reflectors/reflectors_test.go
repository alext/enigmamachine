package reflectors_test

import (
	"testing"

	"github.com/alext/enigmamachine"
	"github.com/alext/enigmamachine/reflectors"
)

func TestReflectors(t *testing.T) {
	specs := []enigmamachine.ReflectorSpec{
		reflectors.A,
		reflectors.B,
		reflectors.C,
		reflectors.Bthin,
		reflectors.Cthin,
	}
	for i, spec := range specs {
		_, err := enigmamachine.NewReflector(spec)
		if err != nil {
			t.Errorf("[%d] spec: %s, expected no error, got: %s", i, spec, err.Error())
		}
	}
}
