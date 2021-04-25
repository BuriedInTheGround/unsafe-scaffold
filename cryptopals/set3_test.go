package cryptopals

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"testing"
)

func TestChallenge17(t *testing.T) {
	targets := []string{
		"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
		"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
		"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
		"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
		"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
		"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
		"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
		"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
		"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
		"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
	}

	// Choose a random target and create the oracle functions.
	bigIdx, _ := rand.Int(rand.Reader, big.NewInt(10))
	idx := int(bigIdx.Int64())
	target := base64ToByteSlice(t, targets[idx])
	generateCiphertext, isPaddingValid := newCBCPaddingOracles(target)

	// Generate the ciphertext, then recover the plaintext by exploiting the
	// padding oracle.
	ciphertext := generateCiphertext()
	recoveredPlaintext := attackCBCPaddingOracle(ciphertext, isPaddingValid)

	if !bytes.Equal(recoveredPlaintext, target) {
		t.Fatalf("wrong plaintext recovered; want %q, got %q", target, recoveredPlaintext)
	}
}
