package cryptopals

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"unicode"
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
	guessOneMoreByte := func(knownBytes, previousBlock, targetBlock []byte) []byte {
		if len(knownBytes) >= aes.BlockSize {
			panic("attackCBCPaddingOracle: wrong input length for guessOneMoreByte")
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
				return guessedPlaintext
			}
		}

		// (*) If every guess is skipped, it means that an already valid
		// padding was found, specifically of length len(guessedPlaintext). So,
		// set the byte to be guessed to this padding value.
		guessedPlaintext[0] = byte(len(guessedPlaintext))

		return guessedPlaintext
	}

	var plaintext []byte
	for i := 0; i < len(ciphertext)/aes.BlockSize-1; i++ {
		var knownBytes []byte
		for j := 0; j < aes.BlockSize; j++ {
			previousBlock := ciphertext[i*aes.BlockSize : (i+1)*aes.BlockSize]
			targetBlock := ciphertext[(i+1)*aes.BlockSize : (i+2)*aes.BlockSize]
			knownBytes = guessOneMoreByte(knownBytes, previousBlock, targetBlock)
		}
		plaintext = append(plaintext, knownBytes...)
	}

	stripped, err := validateAndStripPKCS7(plaintext)
	for err == nil {
		plaintext = stripped
		stripped, err = validateAndStripPKCS7(plaintext)
	}
	return plaintext
}

func encryptCTR(nonce, in []byte, b cipher.Block) []byte {
	if len(nonce) != b.BlockSize()/2 {
		panic("encryptCTR: wrong nonce length")
	}
	keystream := make([]byte, b.BlockSize())

	counter := uint64(0)
	counterBytes := make([]byte, b.BlockSize()/2)
	binary.LittleEndian.PutUint64(counterBytes, counter)
	b.Encrypt(keystream, append(nonce, counterBytes...))

	res := make([]byte, len(in))
	for i := range in {
		res[i] = in[i] ^ keystream[i%b.BlockSize()]
		if (i+1)%b.BlockSize() == 0 {
			counter++
			counterBytes := make([]byte, b.BlockSize()/2)
			binary.LittleEndian.PutUint64(counterBytes, counter)
			b.Encrypt(keystream, append(nonce, counterBytes...))
		}
	}

	return res
}

func newCTRFixedNonce() (
	encrypt func(in []byte) []byte,
) {
	key := randomBytes(aes.BlockSize)
	b, err := aes.NewCipher(key)
	if err != nil {
		panic("newCTRFixedNonce: cannot initialize cipher")
	}
	nonce := bytes.Repeat([]byte{0}, b.BlockSize()/2)
	encrypt = func(in []byte) []byte {
		return encryptCTR(nonce, in, b)
	}
	return
}

func findCTRFixedNonceKeystreamWithSubstitution(ciphertexts [][]byte, scoring map[rune]float64) []byte {
	// Generate a subset of scoring containing only uppercase characters. This
	// will be used to guess the first byte of the keystream.
	upperScoring := make(map[rune]float64)
	for r, s := range scoring {
		if !unicode.IsUpper(r) {
			continue
		}
		upperScoring[r] = s
	}

	// Find the longest ciphertext.
	var longest []byte
	maxLength := 0
	for _, c := range ciphertexts {
		if len(c) > maxLength {
			longest = c
			maxLength = len(c)
		}
	}

	// Initialize the keystream with the same length as the longest ciphertext.
	keystream := make([]byte, maxLength)

	for i := 0; i < maxLength; i++ {
		// column contains all bytes XORed with the same keystream byte.
		var column []byte
		for _, c := range ciphertexts {
			if len(c) <= i {
				continue
			}
			column = append(column, c[i])
		}

		// Find the most probable value for the i-th keystream byte.
		var bestByteGuess byte
		maxScore := float64(0)
		for b := 0; b < 256; b++ {
			byteGuess := longest[i] ^ byte(b)
			t := xor(column, bytes.Repeat([]byte{byteGuess}, len(column)))
			var score float64
			if i == 0 {
				score = scoreByteSlice(t, upperScoring)
			} else {
				score = scoreByteSlice(t, scoring)
			}
			if score > maxScore {
				bestByteGuess = byte(b)
				maxScore = score
			}
		}

		// Set the i-byte keystream byte using the best scoring byte guess.
		keystream[i] = longest[i] ^ bestByteGuess
	}

	return keystream
}
