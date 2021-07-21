# Chapter 2

## Exercise 2.1

### 1.

Let *xi*, *yi*, *si* be **in {0, 1, ..., 25}**.

The key stream bits *si* are random integers in Z26.

*Encryption:* **yi = E(xi) = xi + si mod 26**\
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

## Exercise 2.8

### 1.

```
x^4 + x + 1
```

From Table 2.3 (pag. 45), this one is a primitive (hence irreducible)
polynomial.

| clk | FF3 | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|:---:|
|  0  |  1  |  1  |  1  |  1  |
|  1  |  0  |  1  |  1  |  1  |
|  2  |  0  |  0  |  1  |  1  |
|  3  |  0  |  0  |  0  |  1  |
|  4  |  1  |  0  |  0  |  0  |
|  5  |  0  |  1  |  0  |  0  |
|  6  |  0  |  0  |  1  |  0  |
|  7  |  1  |  0  |  0  |  1  |
|  8  |  1  |  1  |  0  |  0  |
|  9  |  0  |  1  |  1  |  0  |
| 10  |  1  |  0  |  1  |  1  |
| 12  |  0  |  1  |  0  |  1  |
| 13  |  1  |  0  |  1  |  0  |
| 14  |  1  |  1  |  0  |  1  |
| 15  |  1  |  1  |  1  |  0  |
| 16  |  1  |  1  |  1  |  1  |

As a result, all rotations of the sequence `111100010011010` are generated by
`x^4 + x + 1`.

### 2.

```
x^4 + x^2 + 1
```

| clk | FF3 | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|:---:|
|  0  |  1  |  1  |  1  |  1  |
|  1  |  0  |  1  |  1  |  1  |
|  2  |  0  |  0  |  1  |  1  |
|  3  |  1  |  0  |  0  |  1  |
|  4  |  1  |  1  |  0  |  0  |
|  5  |  1  |  1  |  1  |  0  |
|  6  |  1  |  1  |  1  |  1  |

All rotations of the sequence `111100` are generated by `x^4 + x^2 + 1`.

| clk | FF3 | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|:---:|
|  0  |  1  |  0  |  0  |  0  |
|  1  |  0  |  1  |  0  |  0  |
|  2  |  1  |  0  |  1  |  0  |
|  3  |  0  |  1  |  0  |  1  |
|  4  |  0  |  0  |  1  |  0  |
|  5  |  0  |  0  |  0  |  1  |
|  6  |  1  |  0  |  0  |  0  |


All rotations of the sequence `000101` are generated by `x^4 + x^2 + 1`.

| clk | FF3 | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|:---:|
|  0  |  0  |  1  |  1  |  0  |
|  1  |  1  |  0  |  1  |  1  |
|  2  |  1  |  1  |  0  |  1  |
|  3  |  0  |  1  |  1  |  0  |

All rotations of the sequence `011` are generated by `x^4 + x^2 + 1`.

As a result, this polynomial is reducible, indeed the sequence length depends
on the initial values of the register.

### 3.

```
x^4 + x^3 + x^2 + x + 1
```

| clk | FF3 | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|:---:|
|  0  |  1  |  1  |  1  |  1  |
|  1  |  0  |  1  |  1  |  1  |
|  2  |  1  |  0  |  1  |  1  |
|  3  |  1  |  1  |  0  |  1  |
|  4  |  1  |  1  |  1  |  0  |
|  5  |  1  |  1  |  1  |  1  |


All rotations of the sequence `11110` are generated by `x^4 + x^3 + x^2 + x + 1`.


| clk | FF3 | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|:---:|
|  0  |  1  |  0  |  0  |  0  |
|  1  |  1  |  1  |  0  |  0  |
|  2  |  0  |  1  |  1  |  0  |
|  3  |  0  |  0  |  1  |  1  |
|  4  |  0  |  0  |  0  |  1  |
|  5  |  1  |  0  |  0  |  0  |

All rotations of the sequence `00011` are generated by `x^4 + x^3 + x^2 + x + 1`.

| clk | FF3 | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|:---:|
|  0  |  0  |  1  |  0  |  1  |
|  1  |  0  |  0  |  1  |  0  |
|  2  |  1  |  0  |  0  |  1  |
|  3  |  0  |  1  |  0  |  0  |
|  4  |  1  |  0  |  1  |  0  |
|  5  |  0  |  1  |  0  |  1  |

All rotations of the sequence `10100` are generated by `x^4 + x^3 + x^2 + x + 1`.

As a result, this polynomial is irreducible, because the sequence length is
indipendent of the initial value of the register, but isn't primitive, because
the sequence length is not maximum.

## Exercise 2.9

### 1.

To launch a successful attack, 512 plaintext/ciphertext bit pairs are needed.

### 2.

Let `m = 256`.

Attack steps:
1. Get the said 512 plaintext/ciphertext bit pairs.
2. Calculate the values of `si = xi + yi mod 2` for `i = 0, 1, ..., 2m-1`; that
   is reconstruct the from 512 bits of key stream.
3. Mount a system of linear equations to find the values `pi` of the feedback
   coefficients, described as following.

    ```
    s_{i+m} = sum_{j=0}^{m-1} ( p_j * s_{i+j} ) mod 2

    with s_i, p_j in {0, 1}
         i = 0, 1, ..., m-1
    ```

4. Solve the system of linear equations for the unknowns, which are indeed the
   256 feedback coefficients.
5. Using the values found, build the LFSR with which decrypt the entire
   ciphertext.

### 3.

In this system, the key is the set of feedback coefficients. Since the initial
contents of the LFSR are outputted unaltered, it would be easy to calculate
them, so it doesn't make sense to include them in the key or worse to use only
them as the key.

## Exercise 2.10

*Plaintext:* `1001 0010 0110 1101 1001 0010 0110`\
*Ciphertext:* `1011 1100 0011 0001 0010 1011 0001`

### 1.

*Key stream:*
```
XOR  1001 0010 0110 1101 1001 0010 0110
     1011 1100 0011 0001 0010 1011 0001
   ------------------------------------
     0010 1110 0101 1100 1011 1001 0111
```

The key stream is the repetition of the sequence `0010111`, that has length
equals to 7, hence the degree of the LFSR is `m = log2(7 + 1) = 3`.

### 2.

The initialization vector is composed of the first m values of the key stream
(in the reverse order), so in this case it is `IV = 100`.

### 3.

The first 2m key stream bits are `001011`.

The linear equations are

```
s3 = p0*s0 + p1*s1 + p2*s2  mod 2
s4 = p0*s1 + p1*s2 + p2*s3  mod 2
s5 = p0*s2 + p1*s3 + p2*s4  mod 2
```

substituting the known values we have

```
0 = p2  mod 2
1 = p1  mod 2
1 = p0 + p2  mod 2
```

so the results are

```
p2 = 0
p1 = 1
p0 = 1
```

and the corresponding polynomial is `x^3 + x + 1`.

### 4.

```
 +----------------------------------⊕<---------------+
 |                                  ^                |
 |    +--------+       +--------+   |   +--------+   |
 |    |        |       |        |   |   |        |   |
 +--->|   s2   |------>|   s1   |---+-->|   s0   |---+--> si, ..., s1, s0
      |        |       |        |       |        |
      +--------+       +--------+       +--------+
         FF2              FF1              FF0
```

| clk | FF2 | FF1 | FF0 |
|:---:|:---:|:---:|:---:|
|  0  |  1  |  0  |  0  |
|  1  |  0  |  1  |  0  |
|  2  |  1  |  0  |  1  |
|  3  |  1  |  1  |  0  |
|  4  |  1  |  1  |  1  |
|  5  |  0  |  1  |  1  |
|  6  |  0  |  0  |  1  |
|  7  |  1  |  0  |  0  |

The output sequence is `0010111` as expected.

## Exercise 2.11

*Ciphertext:*
```
j5a0edj2b

base 10 = 9 31 0 26 4 3 9 28 1

base 2 = 01001 11111 00000 11010 00100 00011 01001 11100 00001
```

*Header of the plaintext:*
```
WPI

base 10 = 22 15 8

base 2 = 10110 01111 01000
```

*Key stream (partial):*
```
XOR  01001 11111 00000
     10110 01111 01000
   -------------------
     11111 10000 01000
```

### 1.

Given the partial recovered key stream, the initialization vector is `IV = 111111`.

### 2.

The first 2m key stream bits are `111111000001`.

The linear equations are

```
s6 = p0*s0 + p1*s1 + p2*s2 + p3*s3 + p4*s4 + p5*s5  mod 2
s7 = p0*s1 + p1*s2 + p2*s3 + p3*s4 + p4*s5 + p5*s6  mod 2
s8 = p0*s2 + p1*s3 + p2*s4 + p3*s5 + p4*s6 + p5*s7  mod 2
s9 = p0*s3 + p1*s4 + p2*s5 + p3*s6 + p4*s7 + p5*s8  mod 2
s10 = p0*s4 + p1*s5 + p2*s6 + p3*s7 + p4*s8 + p5*s9  mod 2
s11 = p0*s5 + p1*s6 + p2*s7 + p3*s8 + p4*s9 + p5*s10  mod 2
```

substituting the known values we have

s0  s1  s2  s3  s4  s5  s6  s7  s8  s9  s10 s11
1   1   1   1   1   1   0   0   0   0   0   1

```
0 = p0 + p1 + p2 + p3 + p4 + p5  mod 2
0 = p0 + p1 + p2 + p3 + p4  mod 2
0 = p0 + p1 + p2 + p3  mod 2
0 = p0 + p1 + p2  mod 2
0 = p0 + p1  mod 2
1 = p0  mod 2
```

so the results are

```
p0 = 1
p1 = 1
p2 = 0
p3 = 0
p4 = 0
p5 = 0
```

and the corresponding polynomial is `x^6 + x + 1`.

### 3.

See [this](./ex11.go) Go source code.

*Key Stream (complete):* `111111000001000011000101001111010001110010010`\
*Plaintext:* `WPIWOMBAT`

### 4.

The wombats inhabit the forests, mountains and moors of southeastern Australia
and Tasmania.

### 5.

We performed a **known-plaintext attack**.

## Exercise 2.12

Trivium info:

| Register | Length | Feedback Bit | Feedforward Bit | AND Inputs |
|:--------:|:------:|:------------:|:---------------:|:----------:|
|  A       |   93   |  69          |  66             |    91, 92  |
|  B       |   84   |  78          |  69             |    82, 83  |
|  C       |  111   |  87          |  66             |  109, 110  |

Initial state:

```
A = 0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
B = 0000000000000000000000000000000000000000000000000000000000000000000000000000000000
C = 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000111
```

See [this](./ex12.go) Go source code.

The result is the following:
```
First 70 bits during the warm-up phase of Trivium:
01100000000000000000000000000000000000000000000000000000000000000000110
```
