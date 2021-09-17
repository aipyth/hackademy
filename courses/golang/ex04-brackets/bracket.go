package brackets

func Bracket(brackets string) (bool, error) {
	bracketsStack := NewStack()
	matches := func(f, s string) bool {
		switch f + s {
		case "{}", "[]", "()":
			return true
		default:
			return false
		}
	}
	for _, v := range brackets {
		switch v {
		case '{', '[', '(':
			bracketsStack.Push(string(v))
		case '}', ']', ')':
			s := bracketsStack.Pop()
			if !matches(s, string(v)) {
				return false, nil
			}
		}
	}
	return bracketsStack.Empty(), nil
}
