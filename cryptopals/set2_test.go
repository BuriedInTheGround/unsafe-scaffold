package cryptopals

import (
	"bytes"
	"crypto/aes"
	"math"
	"testing"
)

func TestChallenge9(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")

	want := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	if got := padPKCS7(input, 20); !bytes.Equal(got, want) {
		t.Fatalf("wrong padding; want %q, got %q", want, got)
	}

	want = []byte("YELLOW SUBMARINE\x01")
	if got := padPKCS7(input, 17); !bytes.Equal(got, want) {
		t.Fatalf("wrong padding; want %q, got %q", want, got)
	}

	want = []byte("YELLOW SUBMARINE\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10")
	if got := padPKCS7(input, 16); !bytes.Equal(got, want) {
		t.Fatalf("wrong padding; want %q, got %q", want, got)
	}
}

func TestChallenge10(t *testing.T) {
	b, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Fatalf("cannot initialize aes cipher; err = %v", err)
	}

	pt := []byte("Yellow Submarine has perfect len")
	ct := encryptECB(pt, b)
	if !bytes.Equal(decryptECB(ct, b), pt) {
		t.Fatalf("encryptECB does not satisfy the consistency property")
	}

	input := base64ToByteSlice(t, readFromFile("testdata/10.txt"))
	iv := bytes.Repeat([]byte{0}, b.BlockSize())
	message := decryptCBC(iv, input, b)
	t.Logf("decrypted message =\n%s", message)
}

func TestChallenge11(t *testing.T) {
	input := bytes.Repeat([]byte("magic"), 64)

	var ecb, cbc int
	for i := 0; i < 2000; i++ {
		out := encryptWithRandomKey(input)
		if detectECB(out) {
			ecb++
		} else {
			cbc++
		}
	}

	if math.Abs(float64(ecb-cbc)) > 100 {
		t.Fatalf("ECB and CBC encryption times differ by more than 5%%; got ecb=%d, cbc=%d", ecb, cbc)
	}
}

func TestChallenge12(t *testing.T) {
	want := base64ToByteSlice(t, `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`)
	oracle := newECBSuffixOracle(want)
	got := recoverECBSuffix(oracle)
	t.Logf("secret suffix recovered =\n%s", got)
	if !bytes.Equal(got, want) {
		t.Fatalf("wrong suffix; want %q, got %q", want, got)
	}
}
