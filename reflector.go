package enigmamachine

import "fmt"

type Reflector map[rune]rune

func NewReflector(mapping string) (Reflector, error) {
	r := make(Reflector, 26)
	err := r.populateMapping(mapping)
	return r, err
}

func MustNewReflector(mapping string) Reflector {
	r, err := NewReflector(mapping)
	if err != nil {
		panic(err)
	}
	return r
}

func (r Reflector) populateMapping(mapping string) error {
	var first rune
	for _, l := range mapping {
		if l == ' ' {
			continue
		}
		if l < 'A' || l > 'Z' {
			return fmt.Errorf("invalid character %c in mapping", l)
		}
		if first == 0 {
			first = l // store to match with pair
			continue
		}
		if _, found := r[first]; found {
			return fmt.Errorf("duplicate character %c in mapping", first)
		}
		r[first] = l
		if _, found := r[l]; found {
			return fmt.Errorf("duplicate character %c in mapping", l)
		}
		r[l] = first
		first = 0
	}
	if len(r) != 26 {
		return fmt.Errorf("incomplete mapping (length: %d)", len(r))
	}
	return nil
}

func (r Reflector) Translate(input rune) rune {
	return r[input]
}
