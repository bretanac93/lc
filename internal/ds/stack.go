package ds

type Stack struct {
	store []interface{}
}

func NewStack() *Stack {
	return &Stack{
		store: make([]interface{}, 0),
	}
}

// Push inserts an item into the stack
func (s *Stack) Push(item interface{}) {
	s.store = append(s.store, item)
}

// Pop returns the item at the top of the stack removing it
func (s *Stack) Pop() interface{} {
	lastIdx := len(s.store) - 1
	lastElem := s.store[lastIdx]

	s.store = s.store[:lastIdx]

	return lastElem
}

func (s *Stack) Top() interface{} {
	lastIdx := len(s.store) - 1
	lastElem := s.store[lastIdx]

	return lastElem
}

func (s *Stack) Empty() bool {
	return len(s.store) == 0
}
