package exprxfunc

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value string
		prev *node
	}
)

func NewStack() *Stack {
	return &Stack{nil,0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Top() string {
	if s.length == 0 {
		return Empty
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(value string) {
	n := &node{value,s.top}
	s.top = n
	s.length++
}
