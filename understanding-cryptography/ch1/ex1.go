package main

import "fmt"

// English alphabet letter frequencies.
var alpha = map[rune]float64{
	'a': 0.0817,
	'b': 0.0150,
	'c': 0.0278,
	'd': 0.0425,
	'e': 0.1270,
	'f': 0.0223,
	'g': 0.0202,
	'h': 0.0609,
	'i': 0.0697,
	'j': 0.0015,
	'k': 0.0077,
	'l': 0.0403,
	'm': 0.0241,
	'n': 0.0675,
	'o': 0.0751,
	'p': 0.0193,
	'q': 0.0010,
	'r': 0.0599,
	's': 0.0633,
	't': 0.0906,
	'u': 0.0276,
	'v': 0.0098,
	'w': 0.0236,
	'x': 0.0015,
	'y': 0.0197,
	'z': 0.0007,
}

// Target ciphertext.
const enc = `  lrvmnir bpr sumvbwvr jx bpr lmiwv yjeryrkbi jx qmbm wi
bpr xjvni mkd ymibrut jx irhx wi bpr riirkvr jx
ymbinlmtmipw utn qmumbr dj w ipmhh but bj rhnvwdmbr bpr
yjeryrkbi jx bpr qmbm mvvjudwko bj yt wkbrusurbmbwjk
lmird jx xjubt trmui jx ibndt

  wb wi kjb mk rmit bmiq bj rashmwk rmvp yjeryrkb mkd wbi
iwokwxwvmkvr mkd ijyr ynib urymwk nkrashmwkrd bj ower m
vjyshrbr rashmkmbwjk jkr cjnhd pmer bj lr fnmhwxwrd mkd
wkiswurd bj invp mk rabrkb bpmb pr vjnhd urmvp bpr ibmbr
jx rkhwopbrkrd ywkd vmsmlhr jx urvjokwgwko ijnkdhrii
ijnkd mkd ipmsrhrii ipmsr w dj kjb drry ytirhx bpr xwkmh
mnbpjuwbt lnb yt rasruwrkvr cwbp qmbm pmi hrxb kj djnlb
bpmb bpr xjhhjcwko wi bpr sujsru msshwvmbwjk mkd
wkbrusurbmbwjk w jxxru yt bprjuwri wk bpr pjsr bpmb bpr
riirkvr jx jqwkmcmk qmumbr cwhh urymwk wkbmvb
`

func main() {
	// Build the letter frequency table for the target ciphertext.
	freq := make(map[rune]float64)
	for _, r := range enc {
		freq[r]++
	}
	for k := range freq {
		freq[k] /= float64(len(enc))
	}

	// Remove the space and the carriage return characters from the frequency
	// table, to preserve formatting and focus only on alphabet's characters.
	delete(freq, 32) // <space>
	delete(freq, 13) // '/r'
	delete(freq, 10) // '/n'

	// Build the "decryption key", that is, the characters' association table.
	decKey := make(map[rune]rune)
	for len(freq) > 0 {
		from := keyWithMaxValue(alpha)
		delete(alpha, from)
		to := keyWithMaxValue(freq)
		delete(freq, to)
		decKey[to] = from
	}

	// Manual corrections based on observation.
	decKey['l'] = 'b'
	decKey['n'] = 'u'
	decKey['s'] = 'p'
	decKey['w'] = 'i'
	decKey['j'] = 'o'
	decKey['e'] = 'v'
	decKey['k'] = 'n'
	decKey['x'] = 'f'
	decKey['t'] = 'y'
	decKey['d'] = 'd'
	decKey['q'] = 'k'
	decKey['c'] = 'w'
	decKey['o'] = 'g'
	decKey['h'] = 'l'
	decKey['a'] = 'x'
	decKey['f'] = 'q'
	decKey['g'] = 'z'

	// Print the resulting plaintext by using the built key.
	for _, to := range enc {
		if from, ok := decKey[to]; ok {
			fmt.Printf("%c", from)
		} else {
			fmt.Printf("%c", to)
		}
	}
}

func keyWithMaxValue(x map[rune]float64) rune {
	var max float64 = 0
	var key rune
	for k := range x {
		if x[k] > max {
			max = x[k]
			key = k
		}
	}
	return key
}
