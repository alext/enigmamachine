package enigmamachine

type Machine struct{}

type MachineSetup struct {
	Reflector     ReflectorSpec
	Rotors        []RotorSpec
	RingPositions []int
	Plugboard     PlugboardSpec
}

func New(s MachineSetup) *Machine {
	return &Machine{}
}

func (m *Machine) SetPositions(positions []rune) {}

func (m Machine) Translate(input string) string {
	return ""
}
