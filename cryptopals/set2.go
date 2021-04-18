package cryptopals

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"math/big"
)

func padPKCS7(in []byte, blockSize int) []byte {
	if blockSize <= 0 {
		panic("the blockSize must be greater than 0")
	}
	if blockSize >= 256 {
		panic("cannot pad higher than 255")
	}
	pad := blockSize - len(in)%blockSize
	res := make([]byte, len(in)+pad)
	copy(res[:len(in)], in[:])
	for i := len(in); i < len(res); i++ {
		res[i] = byte(pad)
	}
	return res
}

func encryptECB(in []byte, b cipher.Block) []byte {
	if len(in)%b.BlockSize() != 0 {
		panic("encryptECB: input length not a multiple of BlockSize")
	}
	res := make([]byte, len(in))
	for i := 0; i < len(in); i += b.BlockSize() {
		b.Encrypt(res[i:i+b.BlockSize()], in[i:i+b.BlockSize()])
	}
	return res
}

func encryptCBC(iv, in []byte, b cipher.Block) []byte {
	if len(iv)%b.BlockSize() != 0 {
		panic("encryptCBC: IV length not a multiple of BlockSize")
	}
	plaintext := padPKCS7(in, b.BlockSize())
	res := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += b.BlockSize() {
		var t []byte
		if i == 0 {
			t = xor(iv, plaintext[i:i+b.BlockSize()])
		} else {
			t = xor(res[i-b.BlockSize():i], plaintext[i:i+b.BlockSize()])
		}
		copy(res[i:i+b.BlockSize()], encryptECB(t, b))
	}
	return res
}

func decryptCBC(iv, in []byte, b cipher.Block) []byte {
	if len(iv)%b.BlockSize() != 0 {
		panic("decryptCBC: IV length not a multiple of BlockSize")
	}
	if len(in)%b.BlockSize() != 0 {
		panic("decryptCBC: input length not a multiple of BlockSize")
	}
	res := make([]byte, len(in))
	for i := 0; i < len(in); i += b.BlockSize() {
		t := decryptECB(in[i:i+b.BlockSize()], b)
		if i == 0 {
			copy(res[i:i+b.BlockSize()], xor(iv, t))
		} else {
			copy(res[i:i+b.BlockSize()], xor(in[i-b.BlockSize():i], t))
		}
	}
	padLen := int(res[len(res)-1])
	if padLen == 0 {
		return res[:len(res)-b.BlockSize()]
	}
	return res[:len(res)-padLen]
}

func randomBytes(n int) []byte {
	res := make([]byte, n)
	b, err := rand.Read(res)
	if err != nil {
		panic("randomBytes: failed to read from random")
	}
	return res[:b]
}

func encryptWithRandomKey(in []byte) []byte {
	key := randomBytes(aes.BlockSize)
	b, err := aes.NewCipher(key)
	if err != nil {
		panic("encryptWithRandomKey: cannot initialize cipher")
	}

	// Append 5-10 random bytes before the plaintext.
	rndBefore, err := rand.Int(rand.Reader, big.NewInt(6))
	panicIfErr(err)
	bytesBefore := int(rndBefore.Int64()) + 5
	in = append(randomBytes(bytesBefore), in...)

	// Append 5-10 random bytes after the plaintext.
	rndAfter, err := rand.Int(rand.Reader, big.NewInt(6))
	panicIfErr(err)
	bytesAfter := int(rndAfter.Int64()) + 5
	in = append(in, randomBytes(bytesAfter)...)

	// Choose which encryption mode to use.
	var encryptWithCBC bool
	coinFlip, err := rand.Int(rand.Reader, big.NewInt(2))
	panicIfErr(err)
	if coinFlip.Int64() > 0 {
		encryptWithCBC = true
	}

	if encryptWithCBC {
		return encryptCBC(randomBytes(b.BlockSize()), in, b)
	}
	return encryptECB(padPKCS7(in, b.BlockSize()), b)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func newECBSuffixOracle(secretSuffix []byte) func([]byte) []byte {
	key := randomBytes(aes.BlockSize)
	b, err := aes.NewCipher(key)
	if err != nil {
		panic("newECBSuffixOracle: cannot initialize cipher")
	}
	return func(in []byte) []byte {
		in = append(in, secretSuffix...)
		return encryptECB(padPKCS7(in, b.BlockSize()), b)
	}
}

func recoverECBSuffix(oracle func([]byte) []byte) []byte {
	blockSize := len(oracle(bytes.Repeat([]byte("A"), 1)))
	for i := 2; i <= 128; i++ {
		ct := oracle(bytes.Repeat([]byte("A"), i))
		if len(ct) > blockSize {
			blockSize = len(ct) - blockSize
			break
		}
	}

	if !detectECB(oracle(bytes.Repeat([]byte("magic"), 64))) {
		panic("recoverECBSuffix: oracle is not using ECB")
	}

	var suffix []byte

	// Iterate j times, up to a maximum of 100 blocks.
	for j := 1; j < 100; j++ {
		for b := 1; b < blockSize; b++ {
			// Evaluate the target, which actually is the last byte of the
			// output of `oracle(known-string || unknown-1-byte)`.
			target := oracle(bytes.Repeat([]byte("A"), j*blockSize-len(suffix)-1))
			target = append(target, suffix...)

			// Store every possible last byte oracle output.
			dict := make(map[string]byte)
			for i := 0; i < 256; i++ {
				payload := bytes.Repeat([]byte("A"), j*blockSize-len(suffix)-1)
				payload = append(payload, suffix...)
				payload = append(payload, byte(i))
				out := oracle(payload)
				dict[string(out[(j-1)*blockSize:j*blockSize])] = byte(i)
			}

			// Find the one byte that match and append it to the recovered
			// suffix. If the matching byte has value less than blockSize (and
			// different from newline) terminate, as that is padding whp.
			suffixByte := dict[string(target[(j-1)*blockSize:j*blockSize])]
			if suffixByte != byte('\n') && suffixByte < byte(blockSize) {
				break
			}
			suffix = append(suffix, suffixByte)
		}
	}

	return suffix
}
