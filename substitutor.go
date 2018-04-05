package enigmamachine

import "fmt"

type substitutor map[rune]rune

func newSwappingSubstitutor(mapping string) (substitutor, error) {
	s := make(substitutor)
	var first rune
	for _, r := range mapping {
		if r == ' ' {
			continue
		}
		if r < 'A' || r > 'Z' {
			return nil, fmt.Errorf("invalid character %c in mapping", r)
		}
		if first == 0 {
			first = r // store to match with pair
			continue
		}
		if _, found := s[first]; found {
			return nil, fmt.Errorf("duplicate character %c in mapping", first)
		}
		s[first] = r
		if _, found := s[r]; found {
			return nil, fmt.Errorf("duplicate character %c in mapping", r)
		}
		s[r] = first
		first = 0
	}
	if first != 0 {
		return nil, fmt.Errorf("unpaired character %c in mapping", first)
	}
	return s, nil
}

func newBidirectionalSubstitutors(mapping string) (forward, reverse substitutor, err error) {
	forward, reverse = make(substitutor), make(substitutor)
	for i, r := range mapping {
		if _, found := reverse[r]; found {
			return nil, nil, fmt.Errorf("Duplicate entry '%c' in mapping", r)
		}
		c := rune('A' + i)
		forward[c] = r
		reverse[r] = c
	}
	return forward, reverse, nil
}

func (s substitutor) substitute(input rune) rune {
	r, ok := s[input]
	if ok {
		return r
	}
	return input
}
