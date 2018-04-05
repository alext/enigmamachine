package enigmamachine

type Plugboard struct {
	substitutor
	next Translator
}

func NewPlugboard(config string, next Translator) (p Plugboard, err error) {
	p = Plugboard{next: next}
	p.substitutor, err = newSwappingSubstitutor(config)
	return p, err
}

func (p Plugboard) Translate(input rune) rune {
	r := p.substitute(input)
	r = p.next.Translate(r)
	return p.substitute(r)
}
