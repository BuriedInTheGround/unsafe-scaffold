package cryptopals

import (
	"bytes"
	"crypto/aes"
)

func newCBCPaddingOracles(target []byte) (
	generateCiphertext func() []byte,
	isPaddingValid func(ciphertext []byte) bool,
) {
	key := randomBytes(aes.BlockSize)
	b, err := aes.NewCipher(key)
	if err != nil {
		panic("newCBCPaddingOracles: cannot initialize cipher")
	}
	generateCiphertext = func() []byte {
		iv := randomBytes(b.BlockSize())
		plaintext := padPKCS7(target, b.BlockSize())
		ciphertext := encryptCBC(iv, plaintext, b)
		return append(iv, ciphertext...)
	}
	isPaddingValid = func(in []byte) bool {
		iv := in[:b.BlockSize()]
		plaintext := decryptCBC(iv, in, b)
		_, err := validateAndStripPKCS7(plaintext)
		return err == nil
	}
	return
}

func attackCBCPaddingOracle(ciphertext []byte, isPaddingValid func([]byte) bool) []byte {
	guessNextByte := func(knownBytes, previousBlock, targetBlock []byte) ([]byte, bool) {
		if len(knownBytes) >= aes.BlockSize {
			panic("attackCBCPaddingOracle: wrong input length for guessNextByte")
		}

		// Initialize a payload block that will be modified and used as IV for
		// the targetBlock.
		payload := make([]byte, len(previousBlock))

		// Set the already guessed bytes of the plaintext for the current
		// block.
		guessedPlaintext := append([]byte{0}, knownBytes...)

		// Iterate for every possible byte guess.
		for g := 0; g < 256; g++ {
			// Initially re-set the payload to the same content as the real IV
			// of the targetBlock.
			copy(payload, previousBlock)

			// Set the byte to guess to the current byte guess.
			guessedPlaintext[0] = byte(g)

			// Xor the payload with the plaintext guess.
			for i := 0; i < len(guessedPlaintext); i++ {
				payload[aes.BlockSize-i-1] ^= guessedPlaintext[len(guessedPlaintext)-i-1]
			}

			// Xor the payload with a valid padding of length
			// len(guessedPlaintext).
			for i := 0; i < len(guessedPlaintext); i++ {
				payload[aes.BlockSize-i-1] ^= byte(len(guessedPlaintext))
			}

			// When the guessed byte is the same as the pad byte nothing
			// changes. This case is then skipped because when don't want to
			// test the original ciphertext, it would be useless. (*)
			if bytes.Equal(payload[:aes.BlockSize], previousBlock) {
				continue
			}

			// If the guess evaluates to one with a valid padding, the guess is
			// correct and we return.
			if isPaddingValid(append(payload, targetBlock...)) {
				return guessedPlaintext, false
			}
		}

		// (*) If every guess is skipped, it means that an already valid
		// padding was found, specifically of length len(guessedPlaintext). So,
		// set the byte to be guessed to this padding value.
		guessedPlaintext[0] = byte(len(guessedPlaintext))

		// If the length of the guessed plaintext is equals to the block size,
		// than the entire block contains only padding. Return by signaling
		// this info.
		return guessedPlaintext, len(guessedPlaintext) == aes.BlockSize
	}

	var plaintext []byte
	var isEmptyBlock bool
	for i := 0; i < len(ciphertext)/aes.BlockSize-1; i++ {
		var knownBytes []byte
		for j := 0; j < aes.BlockSize; j++ {
			previousBlock := ciphertext[i*aes.BlockSize : (i+1)*aes.BlockSize]
			targetBlock := ciphertext[(i+1)*aes.BlockSize : (i+2)*aes.BlockSize]
			knownBytes, isEmptyBlock = guessNextByte(knownBytes, previousBlock, targetBlock)
		}
		if isEmptyBlock {
			break
		}
		plaintext = append(plaintext, knownBytes...)
	}

	stripped, err := validateAndStripPKCS7(plaintext)
	if err == nil {
		return stripped
	}

	return plaintext
}
