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

1.\
*Calculation:* **15 * 29 mod 13 =** 2 * 3 mod 13 = **6 mod 13**\
*Proof:* 15 * 29 = 435 = 6 mod 13 indeed 13 divides (435 - 6)

2.\
*Calculation:* **2 * 29 mod 13 =** 2 * 3 mod 13 =  **6 mod 13**\
*Proof:* 2 * 29 = 58 = 6 mod 13 indeed 13 divides (58 - 6)

3.\
*Calculation:* **2 * 3 mod 13 = 6 mod 13**\
*Proof:* 2 * 3 = 6 = 6 mod 13 indeed 13 divides (6 - 6)

4.\
*Calculation:* **-11 * 3 mod 13 =** 2 * 3 mod 13 = **6 mod 13**\
*Proof:* -11 * 3 = -33 = 6 mod 13 indeed 13 divides (-33 - 6)
