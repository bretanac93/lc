package validparentheses

type Stack struct {
	store []rune
}

func NewStack() Stack {
	return Stack{
		store: make([]rune, 0),
	}
}

// Push inserts an item into the stack
func (s *Stack) Push(item rune) {
	s.store = append(s.store, item)
}

// Pop returns the item at the top of the stack removing it
func (s *Stack) Pop() rune {
	lastIdx := len(s.store) - 1
	lastElem := s.store[lastIdx]

	s.store = s.store[:lastIdx]

	return lastElem
}

func (s *Stack) Top() rune {
	lastIdx := len(s.store) - 1
	lastElem := s.store[lastIdx]

	return lastElem
}

func (s *Stack) Empty() bool {
	return len(s.store) == 0
}

func isValid(str string) bool {
	s := NewStack()

	// first, index the opening symbols
	for _, item := range str {
		if item == '(' || item == '[' || item == '{' {
			s.Push(item)
		} else {
			if s.Empty() || (s.Top() == '(' && item != ')') || (s.Top() == '{' && item != '}') || (s.Top() == '[' && item != ']') {
				return false
			}
			s.Pop()
		}
	}

	return s.Empty()
}
