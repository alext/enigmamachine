package enigmamachine_test

import (
	"fmt"

	"github.com/alext/enigmamachine"
	"github.com/alext/enigmamachine/reflectors"
	"github.com/alext/enigmamachine/rotors"
)

func ExampleNew() {
	machine, _ := enigmamachine.New(enigmamachine.MachineSetup{
		Reflector:     reflectors.B,
		Rotors:        []enigmamachine.RotorSpec{rotors.I, rotors.II, rotors.III},
		RingPositions: []int{10, 14, 21},
		Plugboard:     "AP BR CM FZ GJ IL NT OV QS WX",
	})
	machine.SetPositions([]rune{'V', 'Q', 'Q'})

	result, _ := machine.TranslateString("TESTING TESTING")
	fmt.Println(result)
	// Output: HKDHDZF PDBBZXK
}

func ExampleNew_custom() {
	machine, _ := enigmamachine.New(enigmamachine.MachineSetup{
		Reflector: "AF BV CP DJ EI GO HY KR LZ MX NW TQ SU",
		Rotors: []enigmamachine.RotorSpec{
			"BDFHJLCPRTXVZNYEIWGAKMUSQO_V",
			"JPGVOUMFYQBENHZRDKASXLICTW_MZ",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ_AHT",
		},
		RingPositions: []int{10, 14, 21},
		Plugboard:     "AP BR CM FZ GJ IL NT OV QS WX",
	})
	machine.SetPositions([]rune{'V', 'Q', 'Q'})

	result, _ := machine.TranslateString("TESTING TESTING")
	fmt.Println(result)
	// Output: KPNGAST GRNJDVY
}
