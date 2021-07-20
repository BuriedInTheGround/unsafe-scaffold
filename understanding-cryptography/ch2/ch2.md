# Chapter 2

## Exercise 2.1

### 1.

Let *xi*, *yi*, *si* be **in {0, 1, ..., 25}**.

The key stream bits *si* are random integers in Z26.

*Encryption:* **yi = E(xi) = xi + si mod 26**
*Decryption:* **xi = D(yi) = yi - si mod 26**

### 2.

*Ciphertext:* `bsaspp kkuosp`\
*Key stream:* `rsidpy dkawoy` (the book has a mistake, the last key letter is indeed `y`, not `a`)

- Ciphertext to integers: ` 1 18  0 18 15 15   10 10 20 14 18 15`
- Key stream to integers: `17 18  8  3 15 24    3 10  0 22 14 24`
- Plaintext integers:     `10  0 18 15  0 17    7  0 20 18  4 17`

*Plaintext:* `kaspar hauser`

### 3.

*How was the young man murdered?* **He was stabbed.**

## Exercise 2.2

Using the OTP, with the said CD-ROM with capacity of 1Gbyte to store the key,
has the following implications:
- once the every bit of the key stream contained into the CD is used for
  encryption, a completely new key must be generated;
- the CD containing the key must be kept in a safe place, even after its use
  for encrypting;
- CD-ROMs with keys must be distributed using a secure transport method, such
  as a trusted courier;
- every party that need to decrypt messages from someone must have a copy of
  the CD with the key;
- when the key is no longer necessary, the CDs containing it must be destroyed
  at both parties;
- every new key stream must be generated using a TRNG.

## Exercise 2.3

Since every block of data of 16 bytes (128 bits) is encrypted with the same
key, it is possible to take the bits encrypted with the same key bit and apply
letter frequency analysis to decrypt that part of content as if it was
encrypted using a shift cipher.

Once 16 bytes of plaintext are recovered, it is possible to find the entire key
stream by XORing the first 128 bits of plaintext with the first 128 bits of
ciphertext.

## Exercise 2.4

An exhaustive key search (brute-force attack) is not possible against an OTP
system because it's not possible to learn anything about the plaintext given
only the ciphertext (the probability of a bit being 1 or 0 is 50% in both
cases), therefore, even if one goes through every possible key, it cannot
determine whether a found plaintext is the correct one or not.

## Exercise 2.5

### 1.

| clk | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|
|  0  |  1  |  0  |  0  |
|  1  |  1  |  1  |  0  |
|  2  |  1  |  1  |  1  |
|  3  |  0  |  1  |  1  |
|  4  |  1  |  0  |  1  |
|  5  |  0  |  1  |  0  |
|  6  |  0  |  0  |  1  |
|  7  |  1  |  0  |  0  |

*Generated sequence:* **0011101** (period of length 7)

### 2.

| clk | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|
|  0  |  0  |  1  |  1  |
|  1  |  1  |  0  |  1  |
|  2  |  0  |  1  |  0  |
|  3  |  0  |  0  |  1  |
|  4  |  1  |  0  |  0  |
|  5  |  1  |  1  |  0  |
|  6  |  1  |  1  |  1  |
|  7  |  0  |  1  |  1  |

*Generated sequence:* **1101001** (period of length 7)

### 3.

The two sequences have the same period length. More than that, the second
sequence is the left rotated version by 3 positions of the first one.

## Exercise 2.6

Let's call *m* the period length in bits. To attack the described cipher, it is
necessary to know *m* bits of plaintext and the corresponding ciphertext (in
this case, between 18.75 bytes and 25 bytes). Once the attacker knows the *m*
bits of plaintext and ciphertext, he can XOR the two to retrieve the key
stream. It is important to check whether the correct number of bits was being
used, by looking for a repetition in the bits of the key stream. If everything
is fine, the key found can be used to decrypt the entire ciphertext.

## Exercise 2.7

Primitive polynomial:
```
1 + x + x^3 + x^4 + x^8
```

Initialization vector:
```
(FF)16 = (1111 1111)2
```

| clk | FF7 | FF6 | FF5 | FF4 | FF3 | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|  0  |  1  |  1  |  1  |  1  |  1  |  1  |  1  |  1  |
|  1  |  0  |  1  |  1  |  1  |  1  |  1  |  1  |  1  |
|  2  |  0  |  0  |  1  |  1  |  1  |  1  |  1  |  1  |
|  3  |  0  |  0  |  0  |  1  |  1  |  1  |  1  |  1  |
|  4  |  0  |  0  |  0  |  0  |  1  |  1  |  1  |  1  |
|  5  |  1  |  0  |  0  |  0  |  0  |  1  |  1  |  1  |
|  6  |  0  |  1  |  0  |  0  |  0  |  0  |  1  |  1  |
|  7  |  0  |  0  |  1  |  0  |  0  |  0  |  0  |  1  |
|  8  |  1  |  0  |  0  |  1  |  0  |  0  |  0  |  0  |
|  9  |  1  |  1  |  0  |  0  |  1  |  0  |  0  |  0  |
| 10  |  1  |  1  |  1  |  0  |  0  |  1  |  0  |  0  |
| 11  |  0  |  1  |  1  |  1  |  0  |  0  |  1  |  0  |
| 12  |  0  |  0  |  1  |  1  |  1  |  0  |  0  |  1  |
| 13  |  1  |  0  |  0  |  1  |  1  |  1  |  0  |  0  |
| 14  |  0  |  1  |  0  |  0  |  1  |  1  |  1  |  0  |
| 15  |  0  |  0  |  1  |  0  |  0  |  1  |  1  |  1  |

First two output bytes of the LFSR:
```
(1001 0000 1111 1111)2 = (90FF)16
```
