// Package enigmamachine is a library that simulates an Enigma Machine
// (http://en.wikipedia.org/wiki/Enigma_machine). It currently simulates both
// the Enigma I/M3 and the Enigma M4. It also allows for specifying your own
// custom rotor/reflector configurations.
package enigmamachine

// LetterTranslator is the interface that all the individual components of the
// Enigma Machine implement. It represents a component that can take a letter,
// and perform substitutions on it as part of the encryption/decryption process.
type LetterTranslator interface {
	TranslateLetter(rune) rune
}
