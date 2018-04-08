// Package reflectors contains definitions for the standard reflectors that
// were in use.
package reflectors

import "github.com/alext/enigmamachine"

// The standard reflector configurations. A, B and C are used with the 3-rotor
// models. Bthin and Cthin are used with the 4-rotor enigma.
const (
	A     enigmamachine.ReflectorSpec = "AE BJ CM DZ FL GY HX IV KW NR OQ PU ST"
	B                                 = "AY BR CU DH EQ FS GL IP JX KN MO TZ VW"
	C                                 = "AF BV CP DJ EI GO HY KR LZ MX NW TQ SU"
	Bthin                             = "AE BN CK DQ FU GY HW IJ LO MP RX SZ TV"
	Cthin                             = "AR BD CO EJ FN GT HK IV LM PW QZ SX UY"
)
