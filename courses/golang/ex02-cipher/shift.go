package cipher

type Shift struct {
	shift int
}

func NewShift(shift int) *Shift {
	if shift == 0 || shift >= alphabetSize || shift <= -alphabetSize {
		return nil
	}
	return &Shift{shift}
}

func (s *Shift) Encode(source string) string {
	sourceCleaned := lettersOnly(source)
	var encoded string
	var offset int
	for _, v := range sourceCleaned {
		// calculate offset from 'a'
		offset = (int(v-'a') + s.shift) % alphabetSize
		if offset < 0 {
			offset = alphabetSize + offset
		}
		// and write gotten character
		encoded += string(rune('a' + offset))
	}
	return encoded
}

func (s *Shift) Decode(source string) string {
	var decoded string
	var offset int
	for _, v := range source {
		offset = (int(v-'a') - s.shift) % alphabetSize
		if offset < 0 {
			offset = alphabetSize + offset
		}
		decoded += string(rune('a' + offset))
	}
	return decoded
}
