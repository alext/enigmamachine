// Package rotors contains definitions for the standard rotors that
// were in use.
package rotors

import "github.com/alext/enigmamachine"

// The standard rotor configurations. The Beta and Gamma rotors were only used
// in the 4-rotor machines and could only go in position 4 (next to the thin
// reflectors), and hence have no notches.
const (
	I     enigmamachine.RotorSpec = "EKMFLGDQVZNTOWYHXUSPAIBRCJ_Q"
	II                            = "AJDKSIRUXBLHWTMCQGZNPYFVOE_E"
	III                           = "BDFHJLCPRTXVZNYEIWGAKMUSQO_V"
	IV                            = "ESOVPZJAYQUIRHXLNFTGKDCMWB_J"
	V                             = "VZBRGITYUPSDNHLXAWMJQOFECK_Z"
	VI                            = "JPGVOUMFYQBENHZRDKASXLICTW_MZ"
	VII                           = "NZJHGRCXMYSWBOUFAIVLPEKQDT_MZ"
	VIII                          = "FKQHTLXOCBJSPDZRAMEWNIUYGV_MZ"
	Beta                          = "LEYJVCNIXWPBQMDRTAKZGFUHOS_"
	Gamma                         = "FSOKANUERHMBTIYCWLQPZXVGJD_"
)
