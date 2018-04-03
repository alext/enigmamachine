package enigmamachine

import "fmt"

type Reflector struct {
	substitutor
}

func NewReflector(mapping string) (r Reflector, err error) {
	r = Reflector{}
	r.substitutor, err = newSubstitutor(mapping)
	if err == nil && len(r.substitutor) != 26 {
		err = fmt.Errorf("incomplete mapping (length: %d)", len(r.substitutor))
	}
	return r, err
}

func MustNewReflector(mapping string) Reflector {
	r, err := NewReflector(mapping)
	if err != nil {
		panic(err)
	}
	return r
}

func (r Reflector) Translate(input rune) rune {
	return r.substitute(input)
}
