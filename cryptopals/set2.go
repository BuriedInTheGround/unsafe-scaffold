package cryptopals

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
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

func parseCookie(cookie string) map[string]string {
	res := make(map[string]string)
	for _, elem := range strings.Split(cookie, "&") {
		keyvalue := strings.Split(elem, "=")
		res[keyvalue[0]] = keyvalue[1]
	}
	return res
}

func profileFor(email string) string {
	if strings.ContainsAny(email, "&=") {
		panic("profileFor: & and = characters not allowed")
	}
	cookie := "email="
	cookie += email
	cookie += "&"
	cookie += "uid=10&"
	cookie += "role=user"
	return cookie
}

func newECBCutAndPasteOracles() (
	generateCookie func(email string) string,
	isAdmin func(cookie string) bool,
) {
	key := randomBytes(aes.BlockSize)
	b, err := aes.NewCipher(key)
	if err != nil {
		panic("newECBCutAndPasteOracles: cannot initialize cipher")
	}
	generateCookie = func(email string) string {
		in := []byte(profileFor(email))
		return string(encryptECB(padPKCS7(in, b.BlockSize()), b))
	}
	isAdmin = func(cookie string) bool {
		plain := decryptECB([]byte(cookie), b)
		padLen := int(plain[len(plain)-1])
		if padLen == 0 {
			cookie = string(plain[:len(plain)-b.BlockSize()])
		} else {
			cookie = string(plain[:len(plain)-padLen])
		}
		obj := parseCookie(cookie)
		if role, ok := obj["role"]; ok {
			return role == "admin"
		}
		return false
	}
	return
}

func makeAdminECBCookie(generateCookie func(email string) string) string {
	// Cookie content:
	//     BLOCK 0          BLOCK 1          BLOCK 2          BLOCK 3
	// email=foo+x@bar. adminPPPPPPPPPPP com&uid=10&role= userQQQQQQQQQQQQ (Q=12)
	// 1234567890123456 1234567890123456 1234567890123456 1234567890123456

	padding := string(bytes.Repeat([]byte{11}, 11))
	cookie := generateCookie("foo+x@bar.admin" + padding + "com")

	getCookieBlock := func(n int) string {
		return cookie[16*n : 16*(n+1)]
	}

	// Admin cookie content:
	// email=foo+x@bar. com&uid=10&role= adminPPPPPPPPPPP (P=11)
	// 1234567890123456 1234567890123456 1234567890123456

	adminCookie := getCookieBlock(0)
	adminCookie += getCookieBlock(2)
	adminCookie += getCookieBlock(1)

	// Forged cookie:
	// email=foo+x@bar.com&uid=10&role=admin
	// {
	//   email: 'foo+x@bar.com',
	//   uid: 10,
	//   role: 'admin'
	// }

	return adminCookie
}

func newECBSuffixOracleWithPrefix(secretSuffix []byte) func([]byte) []byte {
	key := randomBytes(aes.BlockSize)
	b, err := aes.NewCipher(key)
	if err != nil {
		panic("newECBSuffixOracleWithPrefix: cannot initialize cipher")
	}
	rndCount, _ := rand.Int(rand.Reader, big.NewInt(100))
	prefixLen := int(rndCount.Int64())
	return func(in []byte) []byte {
		in = append(randomBytes(prefixLen), in...)
		in = append(in, secretSuffix...)
		return encryptECB(padPKCS7(in, b.BlockSize()), b)
	}
}

func recoverECBSuffixWithPrefix(oracle func([]byte) []byte) []byte {
	blockSize := 16

	repeatingBlocks := func(in []byte) (count int, lastRepBlock int) {
		count = 0
		foundBlocks := make(map[string]struct{})
		for i := 0; i < len(in); i += blockSize {
			hs := hex.EncodeToString(in[i : i+blockSize])
			if _, ok := foundBlocks[hs]; ok {
				count++
				lastRepBlock = i/blockSize + 1
			}
			foundBlocks[hs] = struct{}{}
		}
		return
	}

	var attackBytes int
	var bytesUntilSuffix int
	initRepBlocks, _ := repeatingBlocks(oracle(bytes.Repeat([]byte{0}, 48)))
	for b := 49; b < 128; b++ {
		blocks, num := repeatingBlocks(oracle(bytes.Repeat([]byte{0}, b)))
		if blocks > initRepBlocks {
			attackBytes = b
			bytesUntilSuffix = num * blockSize
			break
		}
	}

	var suffix []byte

	// Iterate j times, up to a maximum of 100 blocks.
	for j := 1; j < 100; j++ {
		for b := 1; b < blockSize; b++ {
			// Evaluate the target, which actually is the last byte of the
			// output of `oracle(random-prefix || known-string || unknown-1-byte)`.
			target := oracle(bytes.Repeat([]byte("A"), attackBytes+j*blockSize-len(suffix)-1))
			target = append(target, suffix...)

			// Store every possible last byte oracle output.
			dict := make(map[string]byte)
			for i := 0; i < 256; i++ {
				payload := bytes.Repeat([]byte("A"), attackBytes+j*blockSize-len(suffix)-1)
				payload = append(payload, suffix...)
				payload = append(payload, byte(i))
				out := oracle(payload)
				dict[string(out[bytesUntilSuffix+(j-1)*blockSize:bytesUntilSuffix+j*blockSize])] = byte(i)
			}

			// Find the one byte that match and append it to the recovered
			// suffix. If the matching byte has value less than blockSize (and
			// different from newline) terminate, as that is padding whp.
			suffixByte := dict[string(target[bytesUntilSuffix+(j-1)*blockSize:bytesUntilSuffix+j*blockSize])]
			if suffixByte != byte('\n') && suffixByte < byte(blockSize) {
				break
			}
			suffix = append(suffix, suffixByte)
		}
	}

	return suffix
}

func validateAndStripPKCS7(in []byte) ([]byte, error) {
	// We specify a block size to be able to handle the \x00 pad.
	blockSize := 16

	// Treat the last padding byte as the validating one, so use it to get the
	// paddding length. If the pad is \x00 set the padding length to be the
	// same as the block size.
	pad := in[len(in)-1]
	padLength := int(pad)
	if pad == 0 {
		padLength = blockSize
	}

	// Check that the input is long enough to fit the padding.
	if len(in) < padLength {
		return []byte{}, fmt.Errorf("invalid PKCS#7 padding")
	}

	// Check that the padding content is entirely the same as the last byte.
	for i := 2; i <= padLength; i++ {
		if in[len(in)-i] != pad {
			return []byte{}, fmt.Errorf("invalid PKCS#7 padding")
		}
	}

	// Strip the padding and return with no error.
	return in[:len(in)-padLength], nil
}

func newCBCBitflipOracles() (
	generateCookie func(userdata string) string,
	isAdmin func(cookie string) bool,
) {
	key := randomBytes(aes.BlockSize)
	iv := randomBytes(aes.BlockSize)
	b, err := aes.NewCipher(key)
	if err != nil {
		panic("newCBCBitflipOracles: cannot initialize cipher")
	}
	generateCookie = func(userdata string) string {
		prefix := []byte("comment1=cooking%20MCs;userdata=")
		suffix := []byte(";comment2=%20like%20a%20pound%20of%20bacon")
		quotedIn := bytes.ReplaceAll([]byte(userdata), []byte(";"), []byte("%3B"))
		quotedIn = bytes.ReplaceAll(quotedIn, []byte("="), []byte("%3D"))
		in := append(prefix, quotedIn...)
		in = append(in, suffix...)
		return string(encryptCBC(iv, padPKCS7(in, b.BlockSize()), b))
	}
	isAdmin = func(in string) bool {
		plaintext := decryptCBC(iv, []byte(in), b)
		plainCookie, err := validateAndStripPKCS7(plaintext)
		if err != nil {
			return false
		}
		return strings.Contains(string(plainCookie), ";admin=true;")
	}
	return
}

func makeAdminCBCCookie(generateCookie func(userdata string) string) string {
	// Prefix:
	// comment1=cooking %20MCs;userdata=
	// 1234567890123456 1234567890123456

	// Suffix:
	// ;comment2=%20lik e%20a%20pound%20 of%20baconPPPPPP  (P = 6)
	// 1234567890123456 1234567890123456 1234567890123456

	// Injected userdata:
	// XXXXXXXXXXXXXXXX XXXXXXadminXtrue
	// 1234567890123456 1234567890123456

	cookie := generateCookie("XXXXXXXXXXXXXXXXXXXXXXadminXtrue")

	// ord('X') ^ ord(';') = 88 ^ 59 = 99 = ord('c')
	// ord('X') ^ ord('=') = 88 ^ 61 = 101 = ord('e')

	adminCookie := []byte(cookie)

	// Alteration ciphertext of injected userdata:
	//      ???????????????? ????????????????
	//  xor 00000c00000e0000 0000000000000000
	//               ... decrypt ...
	//    = ???????????????? ?????;admin=true
	//      1234567890123456 1234567890123456

	adminCookie[2*aes.BlockSize+5] ^= byte('c')
	adminCookie[2*aes.BlockSize+11] ^= byte('e')

	return string(adminCookie)
}
