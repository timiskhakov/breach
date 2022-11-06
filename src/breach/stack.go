package breach

type stack struct {
	storage []Point
}

func (s *stack) Push(p Point) {
	s.storage = append(s.storage, p)
}

func (s *stack) Pop() {
	if len(s.storage) == 0 {
		return
	}

	s.storage = s.storage[:len(s.storage)-1]
}

func (s *stack) Len() int {
	return len(s.storage)
}

func (s *stack) Exists(p Point) bool {
	for _, v := range s.storage {
		if v == p {
			return true
		}
	}
	
	return false
}

func (s *stack) Slice() []Point {
	return s.storage
}
