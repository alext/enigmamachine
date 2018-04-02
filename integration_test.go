package enigmamachine_test

import "testing"

func TestRealExamples(t *testing.T) {
	realExamples := []struct {
		reflector     Reflector
		rotors        []Rotor
		ringPositions []int
		plugboard     []string
		positions     []rune
		input         string
		expected      string
	}{
		// Examples taken from http://wiki.franklinheath.co.uk/index.php/Enigma/Sample_Messages
		{
			// Enigma Instruction Manual 1930
			reflector:     ReflectorA,
			rotors:        []Rotor{RotorII, RotorI, RotorIII},
			ringPositions: []int{24, 13, 22},
			plugboard:     []string{"AM", "FI", "NV", "PS", "TU", "WZ"},
			positions:     []rune{'A', 'B', 'L'},
			input:         "GCDSE AHUGW TQGRK VLFGX UCALX VYMIG MMNMF DXTGN VHVRM MEVOU YFZSL RHDRR XFJWC FHUHM UNZEF RDISI KBGPM YVXUZ",
			expected:      "FEIND LIQEI NFANT ERIEK OLONN EBEOB AQTET XANFA NGSUE DAUSG ANGBA ERWAL DEXEN DEDRE IKMOS TWAER TSNEU STADT",
			// German: Feindliche Infanterie Kolonne beobachtet. Anfang Südausgang Bärwalde. Ende 3km ostwärts Neustadt.
			// English: Enemy infantry column was observed. Beginning [at] southern exit [of] Baerwalde. Ending 3km east of Neustadt.
		},
		{
			// Operation Barbarossa, 1941
			reflector:     ReflectorB,
			rotors:        []Rotor{RotorII, RotorIV, RotorV},
			ringPositions: []int{2, 21, 12},
			plugboard:     []string{"AV", "BS", "CG", "DL", "FU", "HZ", "IN", "KM", "OW", "RX"},
			positions:     []rune{'B', 'S', 'A'},
			input:         "EDPUD NRGYS ZRCXN UYTPO MRMBO FKTBZ REZKM LXLVE FGUEY SIOZV EQMIK UBPMM YLKLT TDEIS MDICA GYKUA CTCDO MOHWX MUUIA UBSTS LRNBZ SZWNR FXWFY SSXJZ VIJHI DISHP RKLKA YUPAD TXQSP INQMA TLPIF SVKDA SCTAC DPBOP VHJK-",
			expected:      "AUFKL XABTE ILUNG XVONX KURTI NOWAX KURTI NOWAX NORDW ESTLX SEBEZ XSEBE ZXUAF FLIEG ERSTR ASZER IQTUN GXDUB ROWKI XDUBR OWKIX OPOTS CHKAX OPOTS CHKAX UMXEI NSAQT DREIN ULLXU HRANG ETRET ENXAN GRIFF XINFX RGTX-",
			// German: Aufklärung abteilung von Kurtinowa nordwestlich Sebez [auf] Fliegerstraße in Richtung Dubrowki, Opotschka. Um 18:30 Uhr angetreten angriff. Infanterie Regiment 3 geht langsam aber sicher vorwärts. 17:06 Uhr röm eins InfanterieRegiment 3 auf Fliegerstraße mit Anfang 16km ostwärts Kamenec.
			// English: Reconnaissance division from Kurtinowa north-west of Sebezh on the flight corridor towards Dubrowki, Opochka. Attack begun at 18:30 hours. Infantry Regiment 3 goes slowly but surely forwards. 17:06 hours [Roman numeral I?] Infantry Regiment 3 on the flight corridor starting 16 km east of Kamenec.
		},
		{
			// Scharnhorst (Konteradmiral Erich Bey), 1943
			reflector:     ReflectorB,
			rotors:        []Rotor{RotorIII, RotorVI, RotorVIII},
			ringPositions: []int{1, 8, 13},
			plugboard:     []string{"AN", "EZ", "HK", "IJ", "LR", "MQ", "OT", "PV", "SW", "UX"},
			positions:     []rune{'U', 'Z', 'V'},
			input:         "YKAE NZAP MSCH ZBFO CUVM RMDP YCOF HADZ IZME FXTH FLOL PZLF GGBO TGOX GRET DWTJ IQHL MXVJ WKZU ASTR",
			expected:      "STEU EREJ TANA FJOR DJAN STAN DORT QUAA ACCC VIER NEUN NEUN ZWOF AHRT ZWON ULSM XXSC HARN HORS THCO",
			// German: Steuere Tanafjord an. Standort Quadrat AC4992, fahrt 20sm. Scharnhorst. [hco - padding?]
			// English: Heading for Tanafjord. Position is square AC4992, speed 20 knots. Scharnhorst.
		},
		{
			// Enigma M4: U-264 (Kapitänleutnant Hartwig Looks), 1942
			reflector:     ReflectorBthin,
			rotors:        []Rotor{RotorBeta, RotorII, RotorIV, RotorI},
			ringPositions: []int{1, 1, 1, 22},
			plugboard:     []string{"AT", "BL", "DF", "GJ", "HM", "NW", "OP", "QY", "RZ", "VX"},
			positions:     []rune{'V', 'J', 'N', 'A'},
			input:         "NCZW VUSX PNYM INHZ XMQX SFWX WLKJ AHSH NMCO CCAK UQPM KCSM HKSE INJU SBLK IOSX CKUB HMLL XCSJ USRR DVKO HULX WCCB GVLI YXEO AHXR HKKF VDRE WEZL XOBA FGYU JQUK GRTV UKAM EURB VEKS UHHV OYHA BCJW MAKL FKLM YFVN RIZR VVRT KOFD ANJM OLBG FFLE OPRG TFLV RHOW OPBE KVWM UQFM PWPA RMFH AGKX IIBG",
			expected:      "VONV ONJL OOKS JHFF TTTE INSE INSD REIZ WOYY QNNS NEUN INHA LTXX BEIA NGRI FFUN TERW ASSE RGED RUEC KTYW ABOS XLET ZTER GEGN ERST ANDN ULAC HTDR EINU LUHR MARQ UANT ONJO TANE UNAC HTSE YHSD REIY ZWOZ WONU LGRA DYAC HTSM YSTO SSEN ACHX EKNS VIER MBFA ELLT YNNN NNNO OOVI ERYS ICHT EINS NULL",
			// German: Von Von 'Looks' F T 1132/19 Inhalt: Bei Angriff unter Wasser gedrückt, Wasserbomben. Letzter Gegnerstandort 08:30 Uhr Marine Quadrat AJ9863, 220 Grad, 8sm, stosse nach. 14mb fällt, NNO 4, Sicht 10.
			// English: From Looks, radio-telegram 1132/19 contents: Forced to submerge under attack, depth charges. Last enemy location 08:30 hours, sea square AJ9863, following 220 degrees, 8 knots. [Pressure] 14 millibars falling, [wind] north-north-east 4, visibility 10.
		},
	}
	for i, ex := range realExamples {
		machine := New(MachineSetup{
			Reflector:    ex.reflector,
			Rotors:       ex.rotors,
			RingPosition: ex.ringPositions,
			Plugboard:    ex.plugboard,
		})
		machine.SetPositions(ex.positions)
		actual := machine.Translate(ex.input)
		if actual != ex.expected {
			t.Errorf("[%d] want: %s, got: %s for input %s", i, ex.expected, actual, ex.input)
		}
	}
}
