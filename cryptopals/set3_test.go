package cryptopals

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"math/big"
	"strings"
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

func TestChallenge18(t *testing.T) {
	b, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Fatalf("cannot initialize cipher; err = %v", err)
	}

	// Test consistency equation `D(key, nonce, E(key, nonce, msg)) = msg`.
	want := []byte("sollicitudin aliquam ultrices sagittis orci a scelerisque purus semper eget duis at tellus at urna")
	intermediate := encryptCTR(bytes.Repeat([]byte{42}, aes.BlockSize/2), want, b)
	got := encryptCTR(bytes.Repeat([]byte{42}, aes.BlockSize/2), intermediate, b)
	if !bytes.Equal(got, want) {
		t.Fatalf("wrong D(k, n, E(k, n, m)) result; want %q, got %q", want, got)
	}

	// Decrypt the message.
	input := base64ToByteSlice(t, "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==")
	decryptedMessage := encryptCTR(bytes.Repeat([]byte{0}, aes.BlockSize/2), input, b)
	t.Logf("%q", decryptedMessage)
}

func TestChallenge19(t *testing.T) {
	data := readFromFile("testdata/19.txt")
	inputs := strings.Split(data, "\n")

	encrypt := newCTRFixedNonce()
	var ciphertexts [][]byte
	for _, in := range inputs {
		ciphertexts = append(ciphertexts, encrypt(base64ToByteSlice(t, in)))
	}

	scoring := generateScoringFromCorpus(corpus)

	// This keystream (and the subsequent plaintexts) aren't perfect, in
	// particular for the last few bytes. This is because the guess-and-score
	// tecnique is applied only to single letters and isn't to i-grams.
	keystream := findCTRFixedNonceKeystreamWithSubstitution(ciphertexts, scoring)

	for i, c := range ciphertexts {
		plaintext := xor(c, keystream[:len(c)])
		t.Logf("plaintext #%d ~= %q", i+1, plaintext)
	}
}
