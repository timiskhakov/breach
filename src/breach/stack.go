package breach

type stack struct {
	storage []Point
}

func (s *stack) push(p Point) {
	s.storage = append(s.storage, p)
}

func (s *stack) pop() {
	if len(s.storage) == 0 {
		return
	}

	s.storage = s.storage[:len(s.storage)-1]
}

func (s *stack) len() int {
	return len(s.storage)
}

func (s *stack) exists(p Point) bool {
	for _, v := range s.storage {
		if v == p {
			return true
		}
	}

	return false
}

func (s *stack) slice() []Point {
	return s.storage
}
