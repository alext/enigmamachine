package enigmamachine

type PlugboardSpec string

type Plugboard struct {
	substitutor
	next LetterTranslator
}

func NewPlugboard(config PlugboardSpec, next LetterTranslator) (p Plugboard, err error) {
	p = Plugboard{next: next}
	p.substitutor, err = newSwappingSubstitutor(string(config))
	return p, err
}

func (p Plugboard) TranslateLetter(input rune) rune {
	r := p.substitute(input)
	r = p.next.TranslateLetter(r)
	return p.substitute(r)
}
