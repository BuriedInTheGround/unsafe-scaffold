package cryptopals

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"math"
	"math/bits"
	"unicode/utf8"
)

func hexToBase64(hexString string) string {
	b, err := hex.DecodeString(hexString)
	if err != nil {
		panic("hexToBase64: cannot decode hexString")
	}
	return base64.StdEncoding.EncodeToString(b)
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

func findRepeatingXORKey(in []byte, scoring map[rune]float64) []byte {
	var keySize int
	minDistance := math.MaxInt64
	for size := 2; size <= 40; size++ {
		first := in[:size]
		second := in[size : 2*size]
		third := in[2*size : 3*size]
		fourth := in[3*size : 4*size]
		d := hammingDistance(first, second)
		d += hammingDistance(first, third)
		d += hammingDistance(first, fourth)
		d += hammingDistance(second, third)
		d += hammingDistance(second, fourth)
		d += hammingDistance(third, fourth)
		d /= 6    // average
		d /= size // normalize
		if d <= minDistance {
			keySize = size
			minDistance = d
		}
	}

	// Every element contains a portion of the ciphertext that is encrypted
	// with the same single-byte key.
	blocks := make([][]byte, keySize)
	for i := 0; i < len(in); i++ {
		blocks[i%keySize] = append(blocks[i%keySize], in[i])
	}

	key := make([]byte, keySize)
	for i, b := range blocks {
		key[i] = findSingleXORKey(b, scoring)
	}

	return key
}

func hammingDistance(a, b []byte) int {
	var res int
	for _, x := range xor(a, b) {
		res += bits.OnesCount8(x)
	}
	return res
}

func decryptECB(in []byte, b cipher.Block) []byte {
	if len(in)%b.BlockSize() != 0 {
		panic("decryptECB: input length not a multiple of BlockSize")
	}
	res := make([]byte, len(in))
	for i := 0; i < len(in); i += b.BlockSize() {
		b.Decrypt(res[i:i+b.BlockSize()], in[i:i+b.BlockSize()])
	}
	return res
}

func detectECB(in []byte) bool {
	if len(in)%aes.BlockSize > 0 {
		return false // Should this thing panic() here instead?
	}

	foundBlocks := make(map[string]struct{})
	for i := 0; i < len(in); i += aes.BlockSize {
		hs := hex.EncodeToString(in[i : i+aes.BlockSize])
		if _, ok := foundBlocks[hs]; ok {
			return true
		}
		foundBlocks[hs] = struct{}{}
	}

	return false
}
