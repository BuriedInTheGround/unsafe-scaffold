<style type="text/css">
    ol ol { list-style-type: lower-alpha; }
</style>

# Chapter 1

## Exercise 1.1

*Ciphertext*

```
  lrvmnir bpr sumvbwvr jx bpr lmiwv yjeryrkbi jx qmbm wi
bpr xjvni mkd ymibrut jx irhx wi bpr riirkvr jx
ymbinlmtmipw utn qmumbr dj w ipmhh but bj rhnvwdmbr bpr
yjeryrkbi jx bpr qmbm mvvjudwko bj yt wkbrusurbmbwjk
lmird jx xjubt trmui jx ibndt

  wb wi kjb mk rmit bmiq bj rashmwk rmvp yjeryrkb mkd wbi
iwokwxwvmkvr mkd ijyr ynib urymwk nkrashmwkrd bj ower m
vjyshrbr rashmkmbwjk jkr cjnhd pmer bj lr fnmhwxwrd mkd
wkiswurd bj invp mk rabrkb bpmb pr vjnhd urmvp bpr ibmbr
jx rkhwopbrkrd ywkd vmsmlhr jx urvjokwgwko ijnkdhrii
ijnkd mkd ipmsrhrii ipmsr w dj kjb drry ytirhx bpr xwkmh
mnbpjuwbt lnb yt rasruwrkvr cwbp qmbm pmi hrxb kj djnlb
bpmb bpr xjhhjcwko wi bpr sujsru msshwvmbwjk mkd
wkbrusurbmbwjk w jxxru yt bprjuwri wk bpr pjsr bpmb bpr
riirkvr jx jqwkmcmk qmumbr cwhh urymwk wkbmvb
```

*Plaintext*

```
  because the practice of the basic movements of kata is
the focus and mastery of self is the essence of
matsubayashi ryu karate do i shall try to elucidate the
movements of the kata according to my interpretation
based of forty years of study

  it is not an easy task to explain each movement and its
significance and some must remain unexplained to give a
complete explanation one would have to be qualified and
inspired to such an extent that he could reach the state
of enlightened mind capable of recognizing soundless
sound and shapeless shape i do not deem myself the final
authority but my experience with kata has left no doubt
that the following is the proper application and
interpretation i offer my theories in the hope that the
essence of okinawan karate will remain intact
```

The plaintext was recovered in two steps: first, through letter frequency
analysis, a roughly good decryption was extracted; then, by hand, looking at
the previous result, wrong character associations were fixed.

*Who wrote the text?* **Shōshin Nagamine**

## Exercise 1.2

*Ciphertext*

```
xultpaajcxitltlxaarpjhtiwtgxktghidhipxciwtvgtpilpitghlxiwiwtxgqadds
```

*Plaintext*

```
ifweallunitewewillcausetheriverstostainthegreatwaterswiththeirblood
```

which, with spaces added, becomes

```
if we all unite we will cause the rivers to stain the great waters with their blood
```

The plaintext was recovered via letter frequency analysis: the letter `t` have
the higher density (19%), and this means that it should correspond with the
most frequent letter of the English alphabet, which is `e` (with 12.7%). So,
the shift applied, that is the key, is **15**, the distance between the two
letters.

*Who wrote this message?* **Tecumseh** (Tecumseh's Speech to the Osages)

## Exercise 1.3

### 1.

- Total cost of one ASIC = $50 + 100% * $50 = $100
- **Number of ASICs within budget =** $1'000'000 / $100 = **10'000**
- Total keys per second = 10'000 * 5 * 10^8 = 5 * 10^12
- Seconds in a year = 60 * 60 * 24 * 365.25 = 31'557'600
- Number of key attempts on average = 2^128 / 2 = 2^127 (which is, halfway through)
- **Average key search time =** 2^127 / (5 * 10^12 * 31'557'600) = **1.078 * 10^18 years**

### 2.

- Seconds in an hour = 60 * 60 = 3600
- Number of Moore iterations for a 24h key search time = log2(2^127 / (5 * 10^12 * 3600 * 24)) = 68.416
- **Number of years for a 24h key search time =** 68.416 * (18 / 12) = **102.6 years**

## Exercise 1.4

1. **Size of the key space =** 128^8 = **2^56**
2. **Key length =** 7 * 8 = **56 bits**
3. **Key length with only the 26 lowercase letters =** log2(26^8) = **37.6 bits**
4. **Number of required characters for key length of 128 bits =**
    1. [7-bit characters] 128 / 7 = **18.286 characters**
    2. [26 lowercase letters] log26(2^128) = **27.231 characters**

## Exercise 1.5

1. *Calculation:* **15 * 29 mod 13 =** 2 * 3 mod 13 = **6 mod 13**\
   *Proof:* 15 * 29 = 435 = 6 mod 13 indeed 13 divides (435 - 6)

2. *Calculation:* **2 * 29 mod 13 =** 2 * 3 mod 13 =  **6 mod 13**\
   *Proof:* 2 * 29 = 58 = 6 mod 13 indeed 13 divides (58 - 6)

3. *Calculation:* **2 * 3 mod 13 = 6 mod 13**\
   *Proof:* 2 * 3 = 6 = 6 mod 13 indeed 13 divides (6 - 6)

4. *Calculation:* **-11 * 3 mod 13 =** 2 * 3 mod 13 = **6 mod 13**\
   *Proof:* -11 * 3 = -33 = 6 mod 13 indeed 13 divides (-33 - 6)

## Exercise 1.6

1. **1/5 mod 13 = 8** indeed 8 * 5 mod 13 = 1 mod 13
2. **1/5 mod 7 = 3** indeed 3 * 5 mod 7 = 1 mod 7
3. **3 * 2/5 mod 7 =** 3 * 2 * 3 mod 7 = **4** (using 2.)

## Exercise 1.7

### 1.

Addition and multiplication tables for Z4

|   +   | 0 | 1 | 2 | 3 |
|:-----:|:-:|:-:|:-:|:-:|
| **0** | 0 | 1 | 2 | 3 |
| **1** | 1 | 2 | 3 | 0 |
| **2** | 2 | 3 | 0 | 1 |
| **3** | 3 | 0 | 1 | 2 |

|   ×   | 0 | 1 | 2 | 3 |
|:-----:|:-:|:-:|:-:|:-:|
| **0** | 0 | 0 | 0 | 0 |
| **1** | 0 | 1 | 2 | 3 |
| **2** | 0 | 2 | 0 | 2 |
| **3** | 0 | 3 | 2 | 1 |

### 2.

Addition and multiplication tables for Z5

|   +   | 0 | 1 | 2 | 3 | 4 |
|:-----:|:-:|:-:|:-:|:-:|:-:|
| **0** | 0 | 1 | 2 | 3 | 4 |
| **1** | 1 | 2 | 3 | 4 | 0 |
| **2** | 2 | 3 | 4 | 0 | 1 |
| **3** | 3 | 4 | 0 | 1 | 2 |
| **4** | 4 | 0 | 1 | 2 | 3 |

|   ×   | 0 | 1 | 2 | 3 | 4 |
|:-----:|:-:|:-:|:-:|:-:|:-:|
| **0** | 0 | 0 | 0 | 0 | 0 |
| **1** | 0 | 1 | 2 | 3 | 4 |
| **2** | 0 | 2 | 4 | 1 | 3 |
| **3** | 0 | 3 | 1 | 4 | 2 |
| **4** | 0 | 4 | 3 | 2 | 1 |

### 3.

Addition and multiplication tables for Z6

|   +   | 0 | 1 | 2 | 3 | 4 | 5 |
|:-----:|:-:|:-:|:-:|:-:|:-:|:-:|
| **0** | 0 | 1 | 2 | 3 | 4 | 5 |
| **1** | 1 | 2 | 3 | 4 | 5 | 0 |
| **2** | 2 | 3 | 4 | 5 | 0 | 1 |
| **3** | 3 | 4 | 5 | 0 | 1 | 2 |
| **4** | 4 | 5 | 0 | 1 | 2 | 3 |
| **5** | 5 | 0 | 1 | 2 | 3 | 4 |

|   ×   | 0 | 1 | 2 | 3 | 4 | 5 |
|:-----:|:-:|:-:|:-:|:-:|:-:|:-:|
| **0** | 0 | 0 | 0 | 0 | 0 | 0 |
| **1** | 0 | 1 | 2 | 3 | 4 | 5 |
| **2** | 0 | 2 | 4 | 0 | 2 | 4 |
| **3** | 0 | 3 | 0 | 3 | 0 | 3 |
| **4** | 0 | 4 | 2 | 0 | 4 | 2 |
| **5** | 0 | 5 | 4 | 3 | 2 | 1 |

### 4.

The nonzero elements without a multiplicative inverse are:
- { 2 } in Z4
- { 2, 3, 4 } in Z6

In Z5 every nonzero element has a multiplicative inverse because the number 5
is prime, indeed every nonzero element is coprime with 5.

## Exercise 1.8

**Inverse of 5:**
- **in Z11 = 9** indeed 9 * 5 = 45 = 1 mod 11
- **in Z12 = 5** indeed 5 * 5 = 25 = 1 mod 12
- **in Z13 = 8** indeed 8 * 5 = 40 = 1 mod 13

## Exercise 1.9

### 1.

```
x = 3^2 mod 13
  = 9 mod 13
```

### 2.

```
x = 7^2 mod 13
  = 49 mod 13
  = 10 mod 13
```

### 3.

```
x = 3^10 mod 13
  = 3^3 * 3^3 * 3^3 * 3 mod 13
  = 1 * 1 * 1 * 3 mod 13
  = 3 mod 13
```

### 4.

```
x = 7^100 mod 13
  = (7^2)^50 mod 13
  = 10^50 mod 13
  = (10^2)^25 mod 13
  = 100^25 mod 13
  = 9^25 mod 13
  = (9^2)^12 * 9 mod 13
  = 81^12 * 9 mod 13
  = 3^12 * 9 mod 13
  = (3^3)^4 * 9 mod 13
  = 1^4 * 9 mod 13
  = 9 mod 13
```

### 5.

`7^x = 11 mod 13`

By trial and error, a valid solution is `x = 5`.

## Exercise 1.10

Integers in 0 <= n < m relatively prime to
- m = 4: { 1, 3 }
- m = 5: { 1, 2, 3, 4 }
- m = 9: { 1, 2, 4, 5, 7, 8 }
- m = 26: { 1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, 25 }

Euler phi function:
- phi(4) = 2
- phi(5) = 4
- phi(9) = 6
- phi(26) = 12

## Exercise 1.11

*Ciphertext:*

```
falszztysyjzyjkywjrztyjztyynaryjkyswarztyegyyj
```

*Plaintext:*

```
firstthesentenceandthentheevidencesaidthequeen
```

which, with spaces added, becomes

```
first the sentence and then the evidence said the queen
```

The plaintext was recovered by applying the Affine Cipher decryption algorithm
with the given key `k = (7, 22)`.

*Who wrote the line?* **Lewis Carroll in "Alice's Adventures in Wonderland".**

## Exercise 1.12

### 1.

Let the key `k = (a, b)` with `gcd(a, 30) = 1`.

Encryption equation:
```
Ek(x) = (a * x) + b  mod 30
```

Decryption equation:
```
Dk(y) = a^{-1} * (y - b)  mod 30
```

### 2.

**Size of the key space =** phi(30) * 30 = 8 * 30 = **240**

### 3.

key: (a = 17, b = 1)  -->  a^{-1} = 23, indeed 23 * 17 = 1 mod 30

*Ciphertext:*

```
ä u ß w ß
```

*Plaintext:*

```
f r o d o
```

### 4.

*From which village does the plaintext come?* **A village of The Shire.**

## Exercise 1.13

Let `x1 = 0`, `x2 = 1`.

Then, from Alice, Oscar obtains `y1 = (a * 0) + b = b` and `y2 = (a * 1) + b =
a + b`. The value of y1 is the second half of the key, and from the value of
y2, subtracting y1, Oscar gets also the first part of the key.

More generally,
```
a = (x1 - x2)^{-1} * (y1 - y2)  mod m
b = y2 - (a * x2)  mod m
```

with the constraint that `gcd((x1 - x2), m) = 1`.

## Exercise 1.14

### 1.

Given the two affine ciphers

```
ek1 = a1 * x + b1
ek2 = a2 * x + b2
```

we have the combined cipher as

```
ek2(ek1) = a2 * [a1 * x + b1] + b2
         = a2 * a1 * x + a2 * b1 + b2
```

Let

```
a3 = a2 * a1
b3 = a2 * b1 + b2
```

As a result we have a single affine cipher

```
ek3 = a3 * x + b3
```

### 2.

With

```
a1 = 3, b1 = 5
a2 = 11, b2 = 7
```

we have

```
a3 = 11 * 3 = 33
b3 = 11 * 5 + 7 = 62
```

### 3.

Applying the ciphers subsequentially.

```
ek1('K') = 3 * ord('K') + 5
         = 3 * 10 + 5
         = 35
         = 9 mod 26
         = 'J'

ek2('J') = 11 * ord('J') + 7
         = 11 * 9 + 7
         = 106
         = 2 mod 26
         = 'C'
```

Applying the third cipher.

```
ek3('K') = 33 * ord('K') + 62
         = 33 * 10 + 62
         = 392
         = 2 mod 26
         = 'C'
```

The verification succeded, as the result is the same.

### 4.

If an exhaustive key search is used against a ciphertext encrypted with a
double affine ciphertext, it will take the same amount of time as if it was
against a ciphertext encrypted with a single affine ciphertext. This result
comes from the fact that a double affine cipher can always be reduced to a
single one, as shown before. Indeed, the key space does not increase.
