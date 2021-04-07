package cryptopals

import "crypto/cipher"

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
			t = xor(res[i-b.BlockSize():i], plaintext)
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
	return res[:len(res)-padLen]
}
