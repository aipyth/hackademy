package downcase

func Downcase(s string) (string, error) {
	var o string
	for _, v := range s {
		if v <= 90 && v >= 65 {
			o = o + string(v+32)
		} else {
			o = o + string(v)
		}
	}
	return o, nil
}
