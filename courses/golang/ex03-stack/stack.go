package stack

type Stack struct {
	data []int
}

func New() *Stack {
	return &Stack{
		data: make([]int, 5),
	}
}

func (s *Stack) Push(elem int) {
	s.data = append(s.data, elem)
}

func (s *Stack) Pop() (toRet int) {
	toRet = s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return
}
