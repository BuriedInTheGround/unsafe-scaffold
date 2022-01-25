package main

import (
	"fmt"
	"strconv"
)

type bit uint

type lfsr struct {
	coefficients []bit
	state        []bit
}

func newLFSR(coeffs []bit, iv []bit) *lfsr {
	return &lfsr{
		coefficients: coeffs,
		state:        iv,
	}
}

func (r *lfsr) checkInternals() error {
	if len(r.coefficients) != len(r.state) {
		return fmt.Errorf("invalid state/coefficients lengths")
	}
	return nil
}

func (r *lfsr) shift() {
	if err := r.checkInternals(); err != nil {
		panic(err)
	}

	var temp bit
	for i, c := range r.coefficients {
		temp ^= r.state[i] & c
	}

	for i := len(r.state) - 1; i >= 1; i-- {
		r.state[i] = r.state[i-1]
	}

	r.state[0] = temp
}

func (r *lfsr) output() bit {
	if err := r.checkInternals(); err != nil {
		panic(err)
	}

	return r.state[len(r.state)-1]
}

var mapping = map[int]rune{
	0:  'A',
	1:  'B',
	2:  'C',
	3:  'D',
	4:  'E',
	5:  'F',
	6:  'G',
	7:  'H',
	8:  'I',
	9:  'J',
	10: 'K',
	11: 'L',
	12: 'M',
	13: 'N',
	14: 'O',
	15: 'P',
	16: 'Q',
	17: 'R',
	18: 'S',
	19: 'T',
	20: 'U',
	21: 'V',
	22: 'W',
	23: 'X',
	24: 'Y',
	25: 'Z',
	26: '0',
	27: '1',
	28: '2',
	29: '3',
	30: '4',
	31: '5',
}

func main() {
	ciphertext := []bit{
		0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0,
		1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1,
		1, 1, 0, 0, 0, 0, 0, 0, 1,
	}
	keystream := make([]bit, len(ciphertext))
	plaintext := make([]bit, len(ciphertext))

	coeffs := []bit{0, 0, 0, 0, 1, 1}
	iv := []bit{1, 1, 1, 1, 1, 1}
	reg := newLFSR(coeffs, iv)

	fmt.Printf("Key Stream: ")
	for i := 0; i < len(keystream); i++ {
		keystream[i] = reg.output()
		fmt.Printf("%d", keystream[i])
		reg.shift()
	}
	fmt.Printf("\n")

	binStrings := make([]string, len(plaintext)/5)
	var k int
	for i := 0; i < len(plaintext); i++ {
		plaintext[i] = ciphertext[i] ^ keystream[i]
		if i%5 == 0 && i > 0 {
			k++
		}
		binStrings[k] += fmt.Sprintf("%d", plaintext[i])
	}

	fmt.Printf("Plaintext: ")
	for _, s := range binStrings {
		value, _ := strconv.ParseUint(s, 2, 5)
		fmt.Printf("%c", mapping[int(value)])
	}
	fmt.Printf("\n")
}
