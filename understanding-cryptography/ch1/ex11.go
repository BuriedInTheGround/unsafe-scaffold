package main

import "fmt"

// {a, b, a_inverse}
var key = [3]int{7, 22, 15}

const ciphertext = "falszztysyjzyjkywjrztyjztyynaryjkyswarztyegyyj"

func main() {
	for _, b := range []byte(ciphertext) {
		fmt.Printf("%c", decryptByte(int(b)))
	}
	fmt.Printf("\n")

	// output: firstthesentenceandthentheevidencesaidthequeen

	// spaced output: first the sentence and then the evidence said the queen
}

func decryptByte(y int) byte {
	y0 := y - 97

	t := (y0 - key[1]) % 26
	if t < 0 {
		t += 26
	}
	x0 := (key[2] * t) % 26

	return byte(x0) + 97
}
