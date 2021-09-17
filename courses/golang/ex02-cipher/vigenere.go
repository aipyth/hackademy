package cipher

import "strings"

type Vigenere struct {
	key string
}

func testVigenereKey(k string) bool {
	if len(k) < 2 ||
		strings.Count(k, string(k[0])) == len(k) ||
		lettersOnly(k) != k {
		return false
	}
	return true
}

func NewVigenere(key string) *Vigenere {
	if !testVigenereKey(key) {
		return nil
	}
	return &Vigenere{key}
}

func (v *Vigenere) Encode(source string) string {
	sourceCleaned := lettersOnly(source)
	var encoded string
	var offset int
	keyLength := len(v.key)
	for i, r := range sourceCleaned {
		offset = (letterOffset(r) + letterOffset(rune(v.key[i%keyLength]))) % alphabetSize
		encoded += string(rune('a' + offset))
	}
	return encoded
}

func (v *Vigenere) Decode(source string) string {
	var decoded string
	var offset int
	keyLength := len(v.key)
	for i, r := range source {
		offset = (letterOffset(r) - letterOffset(rune(v.key[i%keyLength]))) % alphabetSize
		if offset < 0 {
			offset += alphabetSize
		}
		decoded += string(rune('a' + offset))
	}
	return decoded
}
