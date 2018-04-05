package enigmamachine

type LetterTranslator interface {
	TranslateLetter(rune) rune
}
