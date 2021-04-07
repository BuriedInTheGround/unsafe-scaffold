package cryptopals

import (
	"bytes"
	"crypto/aes"
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
