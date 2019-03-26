# Human-Friendly IDs
HuFID has been created with one purpose:  to help with sharing of a unique piece of info between two endpoints via human input with a possibility of correctness validation of this input. Example of usage: verbal sharing of PSK or initial tokens.

## Format
An ID is a string with groups of 5 symbols with a possible separator (-) between groups. Every group encodes 25 bits (20 bits for the last group).
Example: SVXVC-2T9E2-75VKJ-YCL3E-GX7C6

## Alphabet
0123456789ACEFGHJKLMNPQRSTUVWXYZ
Separator: - (can be any position except the last)

## Normalization
- All uppercase symbols
- Special translation for extra symbols:
```
    O -> 0
    B -> 8
    D -> 0
    I -> 1
```
- Separator after every 5 symbols (except the last)

## Checksum
The last symbol is the validation symbol (checksum).
### Checksum calculation
```
    check[0] := 0x74
    check[i] := symbol[i] xor (check[i-1]+1)
```
where symbol[i] represents ascii value of symbol in position i

### Checksum validation
The last symbol must equal representation of last 5 bits from check[i-1] in the alphabet,
where i is the position of the last symbol.
