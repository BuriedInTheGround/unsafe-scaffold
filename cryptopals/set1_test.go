package cryptopals

import (
	"bytes"
	"encoding/hex"
	"os"
	"strings"
	"testing"
)

func TestChallenge1(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if got := hexToBase64(input); got != want {
		t.Fatalf("wrong hex to base64 conversion; want %q, got %q", want, got)
	}
}

func TestChallenge2(t *testing.T) {
	firstInput := hexToByteSlice(t, "1c0111001f010100061a024b53535009181c")
	secondInput := hexToByteSlice(t, "686974207468652062756c6c277320657965")
	want := "746865206b696420646f6e277420706c6179"
	if got := hex.EncodeToString(xor(firstInput, secondInput)); got != want {
		t.Fatalf("wrong xor result; want %q, got %q", want, got)
	}
}

func hexToByteSlice(t *testing.T, hexString string) []byte {
	t.Helper()
	b, err := hex.DecodeString(hexString)
	if err != nil {
		t.Fatalf("hexToByteSlice: cannot decode hexString")
	}
	return b
}

var corpus = readFromFile("testdata/the-picture-of-dorian-gray.txt")

func TestChallenge3(t *testing.T) {
	input := hexToByteSlice(t, "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	scoring := generateScoringFromCorpus(corpus)
	key := findSingleXORKey(input, scoring)
	t.Logf("found key = %q", key)
	t.Logf("decrypted message = %q", singleXOR(input, key))
}

func readFromFile(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic("readFromFile: cannot read file")
	}
	return string(b)
}

func TestChallenge4(t *testing.T) {
	data := readFromFile("testdata/4.txt")
	inputs := strings.Split(data, "\n")
	scoring := generateScoringFromCorpus(corpus)

	var inputLine int
	var out []byte
	var score float64
	for i, in := range inputs {
		input := hexToByteSlice(t, in)
		key := findSingleXORKey(input, scoring)
		output := singleXOR(input, key)
		s := scoreByteSlice(output, scoring)
		if s > score {
			inputLine = i + 1
			out = output
			score = s
		}
	}

	t.Logf("found singleXOR encrypt message at line #%d", inputLine)
	t.Logf("encrypted message = %q", inputs[inputLine-1])
	t.Logf("decrypted message = %q", out)
}

func TestChallenge5(t *testing.T) {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	got := repeatingXOR([]byte(input), []byte("ICE"))
	if !bytes.Equal(got, hexToByteSlice(t, want)) {
		t.Fatalf("wrong repeating xor result; want %q, got %q", want, got)
	}
}
