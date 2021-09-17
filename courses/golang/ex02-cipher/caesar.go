package cipher

type Caesar struct{}

const shift = 3

func NewCaesar() *Caesar {
	return &Caesar{}
}

func (*Caesar) Encode(source string) string {
	sourceCleaned := lettersOnly(source)
	var encoded string
	var offset rune
	for _, v := range sourceCleaned {
		// calculate offset from 'a'
		offset = ((v - 'a') + shift) % alphabetSize
		// and write gotten character
		encoded += string(rune('a' + offset))
	}
	return encoded
}

func (*Caesar) Decode(source string) string {
	var decoded string
	var offset rune
	for _, v := range source {
		offset = ((v - 'a') - shift) % alphabetSize
		if offset < 0 {
			offset = alphabetSize + offset
		}
		decoded += string(rune('a' + offset))
	}
	return decoded
}
