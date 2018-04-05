package enigmamachine

import "fmt"

type ReflectorSpec string

type Reflector struct {
	substitutor
}

func NewReflector(mapping ReflectorSpec) (r Reflector, err error) {
	r = Reflector{}
	r.substitutor, err = newSubstitutor(string(mapping))
	if err == nil && len(r.substitutor) != 26 {
		err = fmt.Errorf("incomplete mapping (length: %d)", len(r.substitutor))
	}
	return r, err
}

func (r Reflector) Translate(input rune) rune {
	return r.substitute(input)
}
