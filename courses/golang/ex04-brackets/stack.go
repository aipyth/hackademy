package brackets

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{
		data: make([]string, 0, 5),
	}
}

func (s *Stack) Push(element string) {
	s.data = append(s.data, element)
}

func (s *Stack) Pop() string {
	if s.Empty() {
		return ""
	}
	toRet := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return toRet
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}
