# Chapter 3

## Exercise 3.1

### 1.

```
S1(x1) = 1110
S1(x2) = 0000

S1(x1) xor S1(x2) = 1110

S1(x1 xor x2) = 0000 != 1110 q.e.d.
```

### 2.

```
S1(x1) = 1101
S1(x2) = 0100

S1(x1) xor S1(x2) = 1001

S1(x1 xor x2) = 1000 != 1001 q.e.d.
```

### 3.

```
S1(x1) = 0110
S1(x2) = 1100

S1(x1) xor S1(x2) = 1010

S1(x1 xor x2) = 1101 != 1010 q.e.d.
```

## Exercise 3.2

From the `IP` and `IP^{-1}` permutation tables we have the following
correspondences.

```
IP[1] = 58
IP[2] = 50
IP[3] = 42
IP[4] = 34
IP[5] = 26

IP^{-1}[58] = 1
IP^{-1}[50] = 2
IP^{-1}[42] = 3
IP^{-1}[34] = 4
IP^{-1}[26] = 5
```

As a result, the following equation holds true.

```
x[1:5] = IP^{-1}(IP(x[1:5]))
```

## Exercise 3.3

With regard to the *Key Schedule* steps, all the permutations doesn't change
the state.

Inside the *f*-function is the same, until the S-boxes are reached. The results
of the S-boxes and the subsequent permutation P are the following.

```
  S1   S2   S3   S4   S5   S6   S7   S8
   🠗    🠗    🠗    🠗    🠗    🠗    🠗    🠗
 1110 1111 1010 0111 0010 1100 0100 1101

                   P
                   🠗
 1101 1000 1101 1000 1101 1011 1011 1100
```

XORing with zeros in the Feistel network keeps the result the same. Then, half
zeros are copied from R0 to L1.

Finally, the result of the first round of DES with 80 bits of zero as for the
plaintext, and 56 bits also zero for the key, is the following.

```
L1 = 0000 0000 0000 0000 0000 0000 0000 0000 = 0x00000000
R1 = 1101 1000 1101 1000 1101 1011 1011 1100 = 0xD8D8DBBC
```

## Exercise 3.4

The *f*-function gives the same result as the previous exercise, because the
round key with all bits to 1 is XORed with the expanded right half of the
plaintext, with all bits to 1 too.

XORing with ones in the Feistel network negate the result. Then, half ones are
copied from R0 to L1.

So the result of the first round of DES with both the key and the plaintext
with all bits set to 1 is the following.

```
L1 = 1111 1111 1111 1111 1111 1111 1111 1111 = 0xFFFFFFFF
R1 = 0010 0111 0010 0111 0010 0100 0100 0011 = 0x27272443
```

## Exercise 3.5

```
x   = 00000000 00000000 00000000 00000000 00000000 00000000 00000000 10000000
key = 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
```

First, the initial permutation must be applied.

```
IP(x) = 00000000 00000000 00000000 00000000 10000000 00000000 00000000 00000000
```

Then, the right half is expanded.

```
E(R0) = E(10000000 00000000 00000000 00000000)
      = 010000 000000 000000 000000 000000 000000 000000 000001
```

XORing with the round key leaves the expanded half unchanged. Then, the value
go through the S-boxes and the P permutation.

```
  S1   S2   S3   S4   S5   S6   S7   S8
   🠗    🠗    🠗    🠗    🠗    🠗    🠗    🠗
 0011 1111 1010 0111 0010 1100 0100 0001

                   P
                   🠗
 1101 0000 0101 1000 0101 1011 1001 1110
```

XORing with zeros in the Feistel network keeps the result the same. Then, the
unmodified right half bits are copied from R0 to L1. The final result after the
first round is the following.

```
L1 = 1000 0000 0000 0000 0000 0000 0000 0000 = 0x00000000
R1 = 1101 0000 0101 1000 0101 1011 1001 1110 = 0xD0585B9E
```

### 1.

Within the first round, two S-boxes, S1 and S8, got a different input compared
to the case when an all-zero plaintext is provided.

### 2.

According to the S-box design criteria, for one single bit difference two
output bits must change. In this case, a total of 4 bits must have changed.

### 3.

The output of the first round, as stated above, is the following.
```
L1 = 1000 0000 0000 0000 0000 0000 0000 0000 = 0x80000000
R1 = 1101 0000 0101 1000 0101 1011 1001 1110 = 0xD0585B9E
```

### 4.

After the first round, a total of six bits have actually changed compared to
the case when the plaintext is all zero. One bit within the left half, and five
in the left half, three of which by the first S-box and the remaining two by
the last S-box.

## Exercise 3.6

### 1.

After applying the PC-1 permutation, the bit flip (which is at position 1 in
the original key) results at position 8.

```
Round 1
C0                    = 0000 0001 0000 0000 0000 0000 0000
C1 (rotate left by 1) = 0000 0010 0000 0000 0000 0000 0000
PC-2 puts the flipped bit at position 20 of the round key. S-box 4 affected.

Round 2
C1                    = 0000 0010 0000 0000 0000 0000 0000
C2 (rotate left by 1) = 0000 0100 0000 0000 0000 0000 0000
PC-2 puts the flipped bit at position 10 of the round key. S-box 2 affected.

Round 3
C2                    = 0000 0100 0000 0000 0000 0000 0000
C3 (rotate left by 2) = 0001 0000 0000 0000 0000 0000 0000
PC-2 puts the flipped bit at position 16 of the round key. S-box 3 affected.

Round 4
C3                    = 0001 0000 0000 0000 0000 0000 0000
C4 (rotate left by 2) = 0100 0000 0000 0000 0000 0000 0000
PC-2 puts the flipped bit at position 24 of the round key. S-box 4 affected.

Round 5
C4                    = 0100 0000 0000 0000 0000 0000 0000
C5 (rotate left by 2) = 0000 0000 0000 0000 0000 0000 0001
PC-2 puts the flipped bit at position 8 of the round key. S-box 2 affected.

Round 6
C5                    = 0000 0000 0000 0000 0000 0000 0001
C6 (rotate left by 2) = 0000 0000 0000 0000 0000 0000 0100
PC-2 puts the flipped bit at position 17 of the round key. S-box 3 affected.

Round 7
C6                    = 0000 0000 0000 0000 0000 0000 0100
C7 (rotate left by 2) = 0000 0000 0000 0000 0000 0001 0000
PC-2 puts the flipped bit at position 4 of the round key. S-box 1 affected.

Round 8
C7                    = 0000 0000 0000 0000 0000 0001 0000
C8 (rotate left by 2) = 0000 0000 0000 0000 0000 0100 0000
PC-2 discards the flipped bit.

Round 9
C8                    = 0000 0000 0000 0000 0000 0100 0000
C9 (rotate left by 1) = 0000 0000 0000 0000 0000 1000 0000
PC-2 puts the flipped bit at position 11 of the round key. S-box 2 affected.

Round 10
C9                     = 0000 0000 0000 0000 0000 1000 0000
C10 (rotate left by 2) = 0000 0000 0000 0000 0010 0000 0000
PC-2 puts the flipped bit at position 14 of the round key. S-box 3 affected.

Round 11
C10                    = 0000 0000 0000 0000 0010 0000 0000
C11 (rotate left by 2) = 0000 0000 0000 0000 1000 0000 0000
PC-2 puts the flipped bit at position 2 of the round key. S-box 1 affected.

Round 12
C11                    = 0000 0000 0000 0000 1000 0000 0000
C12 (rotate left by 2) = 0000 0000 0000 0010 0000 0000 0000
PC-2 puts the flipped bit at position 9 of the round key. S-box 2 affected.

Round 13
C12                    = 0000 0000 0000 0010 0000 0000 0000
C13 (rotate left by 2) = 0000 0000 0000 1000 0000 0000 0000
PC-2 puts the flipped bit at position 23 of the round key. S-box 4 affected.

Round 14
C13                    = 0000 0000 0000 1000 0000 0000 0000
C14 (rotate left by 2) = 0000 0000 0010 0000 0000 0000 0000
PC-2 puts the flipped bit at position 3 of the round key. S-box 1 affected.

Round 15
C14                    = 0000 0000 0010 0000 0000 0000 0000
C15 (rotate left by 2) = 0000 0000 1000 0000 0000 0000 0000
PC-2 discards the flipped bit.

Round 16
C15                    = 0000 0000 1000 0000 0000 0000 0000
C16 (rotate left by 1) = 0000 0001 0000 0000 0000 0000 0000
PC-2 puts the flipped bit at position 18 of the round key. S-box 3 affected.
```

As a result, during DES encryption, with the said bit flip in the key:
- S-box 1 is affected in rounds 7, 11, 14
- S-box 2 is affected in rounds 2, 5, 9, 12
- S-box 3 is affected in rounds 3, 6, 10, 16
- S-box 4 is affected in rounds 1, 4, 13

### 2.

During DES decryption the key schedule algorithm is inverted, so we have that:
- S-box 1 is affected in rounds 3, 6, 10
- S-box 2 is affected in rounds 5, 8, 12, 15
- S-box 3 is affected in rounds 1, 7, 11, 14
- S-box 4 is affected in rounds 4, 13, 16

## Exercise 3.7

### 1.

For a key to be a _weak key_, the extracted subkeys must satisfy the following
equations.
```
k_1 = k_16
k_2 = k_15
k_3 = k_14
k_4 = k_13
k_5 = k_12
k_6 = k_11
k_7 = k_10
k_8 = k_9
```

The above equations can also be written as
```
k_{1+i} = k_{16-i}    for i = 0, 1, ..., 7
```

### 2.

The four weak DES keys are
```
Kw1 = 0000000000000000000000000000 0000000000000000000000000000
Kw2 = 0000000000000000000000000000 1111111111111111111111111111
Kw3 = 1111111111111111111111111111 0000000000000000000000000000
Kw4 = 1111111111111111111111111111 1111111111111111111111111111
```
because rotating the left or the right half leaves them unchanged.

### 3.

```
P_k = 1 / 2^56

P_kw = 4 * P_k
     = 4 / 2^56
     = 4 / (2^2 * 2^54)
     = 1 / 2^54
     ≈ 5.55 * 10^-17
```

The likelihood that a random selected key is weak is `P_kw = 1 / 2^54`.

## Exercise 3.8

### 1.

From the XOR truth table we have that `A ⊕ 1 = A'`.

Then, we can do
```
    A' ⊕ B' = (A ⊕ 1) ⊕ (B ⊕ 1)
            = (A ⊕ B) ⊕ (1 ⊕ 1)
            = A ⊕ B ∎
```

Also,
```
    A' ⊕ B = (A ⊕ 1) ⊕ B
           = (A ⊕ B) ⊕ 1
           = (A ⊕ B)' ∎
```

### 2.

The `PC-1` permutation is linear because it only move the bits around and one
bit does not affect another one, so we can do
```
    PC-1(k') = PC-1(k ⊕ 1)
             = PC-1(k) ⊕ 1
             = (PC-1(k))' ∎
```

### 3.

As for the previous point, the `LS` rotations are linear, and we have
```
    LSᵢ(C'ᵢ₋₁) = LSᵢ(Cᵢ₋₁ ⊕ 1)
               = LSᵢ(Cᵢ₋₁) ⊕ 1
               = (LSᵢ(Cᵢ₋₁))' ∎
```

### 4.

Using the key `k'`, from 2 we have that
```
    PC-1(k') = (C₀ || D₀)'
             = C'₀ || D'₀
```

Then, using 3, we have that
```
    LSᵢ(C'₀) || LSᵢ(D'₀) = C'₁ || D'₁
                         = (C₁ || D₁)'
```

Finally, since for `PC-2` holds the same property as for `PC-1`, we have that
```
    PC-2((C₁ || D₁)') = (PC-2(C₁ || D₁))'
                      = k'₁
```

This reasoning continues in the same fashion for the calculation of all 16
round keys. We have proven that from `k'` the subkeys `k'ᵢ` are generated. ∎

### 5.

The `IP` is a bitwise permutation where every bit only changes position without
affecting the others, so this is linear and we have that
```
    IP(x') = IP(x ⊕ 1)
           = IP(x) ⊕ 1
           = (IP(x))' ∎
```

### 6.

The same goes for the `E` box, which is just a special type of permutation in
which some bits are used more than once on the output. No bit affects another.

```
    E(R'ᵢ) = E(Rᵢ ⊕ 1)
           = E(Rᵢ) ⊕ 1
           = (E(Rᵢ))' ∎
```

### 7.

Regarding the `f`-function, we first have that
```
    E(R'ᵢ₋₁) ⊕ k'ᵢ = (E(Rᵢ₋₁))' ⊕ k'ᵢ
                   = (E(Rᵢ₋₁) ⊕ 1) ⊕ (kᵢ ⊕ 1)
                   = (E(Rᵢ₋₁) ⊕ kᵢ) ⊕ (1 ⊕ 1)
                   = E(Rᵢ₋₁) ⊕ kᵢ
```
and then the result goes through the S-boxes. Since the input to the S-boxes
when having `R'ᵢ₋₁` and `k'ᵢ` is the same as when we have `Rᵢ₋₁` and `kᵢ`, the
overall output of the `f`-function does not change.

So, after the round `i`, using `R'ᵢ₋₁`, `L'ᵢ₋₁` and `k'ᵢ`, we have that
```
    L'ᵢ₋₁ ⊕ f(R'ᵢ₋₁, k'ᵢ) = L'ᵢ₋₁ ⊕ f(Rᵢ₋₁, kᵢ)
                          = (Lᵢ₋₁ ⊕ 1) ⊕ f(Rᵢ₋₁, kᵢ)
                          = (Lᵢ₋₁ ⊕ f(Rᵢ₋₁, kᵢ)) ⊕ 1
                          = Rᵢ ⊕ 1
                          = R'ᵢ ∎
```

### 8.

Using `x'` as plaintext and `k'` as key, we obtain the following.

```
    IP(x') = (L₀ || R₀)'
           = L'₀ || R'₀
```

And the two halves after the first round are
```
    R'₀ = R₀ ⊕ 1
        = L₁ ⊕ 1
        = L'₁
```
and, from 7,
```
    L'₀ ⊕ f(R'₀, k'₁) = R'₁
```

Proceeding in the same fashion for all 16 rounds we would finally have `L'₁₆`
and `R'₁₆`.

Then, the prove in 5 holds true also for `IP⁻¹`, therefore
```
    IP⁻¹(R'₁₆ || L'₁₆) = IP⁻¹((R₁₆ || L₁₆)')
                       = (IP⁻¹(R₁₆ || L₁₆))'
                       = y'
```

And we have proven that `y' = DESₖ'(x')`. ∎

## Exercise 3.9

Using an exhaustive key search, in the worst-case scenario we have to test all
`2⁵⁶` keys, while on average we just need to test half of them, which means
`2⁵⁵` keys.

## Exercise 3.10

### 1.

The clock frequency is measured in hertz, so the expression needs to be
something like
```
[Hz] = 1/[s] = [bit]/[s] × [bit]
```

To encrypt a single block we need to do 16 rounds, and one round requires one
clock cycle, so to encrypt a block 16 clock cycles are required.

The block size of DES is 64 bits, so we encrypt 64 bits at a time.

Hence, the expression for the clock rate (or frequency) is as follows, where
`f` is the frequency in hertz and `r` is the data rate in bits per second.
```
f = r / 64 * 16
```

The dimensional analysis follows.
```
[Hz] = [bit]/[s] / [bit] * 1
     = 1/[s] ∎
```

### 2.

For a data rate of 1 Gbit/s a clock frequency of 250 MHz is required.

For a data rate of 8 Gbit/s a clock frequency of 2 GHz is required.
