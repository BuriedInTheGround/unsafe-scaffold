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
