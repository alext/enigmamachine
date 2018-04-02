package enigmamachine

type Machine struct{}

type MachineSetup struct {
	Reflector     Reflector
	Rotors        []Rotor
	RingPositions []int
	Plugboard     []string
}

func New(s MachineSetup) *Machine {
	return &Machine{}
}

func (m *Machine) SetPositions(positions []rune) {}

func (m Machine) Translate(input string) string {
	return ""
}
