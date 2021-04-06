package cryptopals

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"unicode/utf8"
)

func hexToBase64(hexString string) string {
	b, err := hex.DecodeString(hexString)
	if err != nil {
		panic("hexToBase64: cannot decode hexString")
	}
	return base64.RawStdEncoding.EncodeToString(b)
}

func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor: length mismatch")
	}
	res := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] ^ b[i]
	}
	return res
}

func singleXOR(in []byte, key byte) []byte {
	return xor(in, bytes.Repeat([]byte{key}, len(in)))
}

func findSingleXORKey(in []byte, scoring map[rune]float64) byte {
	key := byte(0)
	score := scoreByteSlice(singleXOR(in, key), scoring)
	for i := 1; i < 256; i++ {
		k := byte(i)
		s := scoreByteSlice(singleXOR(in, k), scoring)
		if s > score {
			key = k
			score = s
		}
	}
	return key
}

func scoreByteSlice(in []byte, scoring map[rune]float64) float64 {
	var score float64
	input := string(in)
	for _, c := range input {
		score += scoring[c]
	}
	return score / float64(utf8.RuneCountInString(input))
}

func generateScoringFromCorpus(corpus string) map[rune]float64 {
	res := make(map[rune]float64)
	for _, c := range corpus {
		res[c]++
	}
	corpusLen := float64(utf8.RuneCountInString(corpus))
	for c, score := range res {
		res[c] = score / corpusLen
	}
	return res
}

func repeatingXOR(in, key []byte) []byte {
	res := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		res[i] = in[i] ^ key[i%len(key)]
	}
	return res
}
