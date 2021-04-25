package cryptopals

import (
	"bytes"
	"crypto/aes"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestChallenge13(t *testing.T) {
	objWant := map[string]string{
		"foo": "bar",
		"baz": "qux",
		"zap": "zazzle",
	}
	objGot := parseCookie("foo=bar&baz=qux&zap=zazzle")
	if !cmp.Equal(objGot, objWant) {
		t.Fatalf("wrong cookie parsing; want %#v, got %#v", objWant, objGot)
	}

	strWant := "email=foo@bar.com&uid=10&role=user"
	if got := profileFor("foo@bar.com"); got != strWant {
		t.Fatalf("wrong profile cookie encoding; want %q, got %q", strWant, got)
	}

	generateCookie, isAdmin := newECBCutAndPasteOracles()

	if isAdmin(generateCookie("foo@bar.com")) {
		t.Fatalf("generated cookie cannot be for role admin")
	}

	if !isAdmin(makeAdminECBCookie(generateCookie)) {
		t.Fatalf("made up cookie is not for role admin")
	}
}

func TestChallenge14(t *testing.T) {
	want := base64ToByteSlice(t, `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`)
	oracle := newECBSuffixOracleWithPrefix(want)
	got := recoverECBSuffixWithPrefix(oracle)
	t.Logf("secret suffix recovered =\n%s", got)
	if !bytes.Equal(got, want) {
		t.Fatalf("wrong suffix; want %q, got %q", want, got)
	}
}

func TestChallenge15(t *testing.T) {
	input := []byte("ICE ICE BABY\x04\x04\x04\x04")
	want := []byte("ICE ICE BABY")
	got, err := validateAndStripPKCS7(input)
	if err != nil {
		t.Fatalf("expected no error; got %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("wrong padding strip; want %q, got %q", want, got)
	}

	input = []byte("ICE ICE BABY\x05\x05\x05\x05")
	_, err = validateAndStripPKCS7(input)
	if err == nil {
		t.Fatalf("expected error")
	}

	input = []byte("ICE ICE BABY\x01\x02\x03\x04")
	_, err = validateAndStripPKCS7(input)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestChallenge16(t *testing.T) {
	generateCookie, isAdmin := newCBCBitflipOracles()

	input := "iceicebaby;admin=true;"
	if isAdmin(generateCookie(input)) {
		t.Fatalf("this would be too easy! try again")
	}

	if !isAdmin(makeAdminCBCCookie(generateCookie)) {
		t.Fatalf("generated cookie is not for admin")
	}
}
