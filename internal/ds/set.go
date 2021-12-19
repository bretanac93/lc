package ds

type Set struct {
	store map[string]struct{}
}

func NewSet() Set {
	return Set{
		store: make(map[string]struct{}),
	}
}

func (s *Set) Add(item string) {
	s.store[item] = struct{}{}
}

func (s *Set) Remove(item string) {
	delete(s.store, item)
}

func (s *Set) Exists(item string) bool {
	_, exists := s.store[item]
	return exists
}
