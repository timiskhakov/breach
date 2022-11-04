package solver

import "errors"

type stack struct {
	storage []point
}

func (s *stack) Push(p point) {
	s.storage = append(s.storage, p)
}

func (s *stack) Pop() error {
	if len(s.storage) == 0 {
		return errors.New("stack is empty")
	}
	s.storage = s.storage[:len(s.storage)-1]
	return nil
}

func (s *stack) Len() int {
	return len(s.storage)
}

func (s *stack) Exists(p point) bool {
	for _, v := range s.storage {
		if v == p {
			return true
		}
	}
	return false
}

func (s *stack) Iter() chan point {
	c := make(chan point)
	go func() {
		for i := 0; i < len(s.storage); i++ {
			c <- s.storage[i]
		}
		close(c)
	}()
	return c
}
