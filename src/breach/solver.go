package breach

type Solver struct {
	stack stack
}

func (s *Solver) Solve(matrix [][]byte, seq []byte) []Point {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == seq[0] {
				s.dfs(matrix, seq, Point{i, j}, i != 0)
				if s.stack.len() == len(seq) {
					if i != 0 {
						return append([]Point{{0, j}}, s.stack.slice()...)
					}

					return s.stack.slice()
				}
			}
		}
	}

	return s.stack.slice()
}

func (s *Solver) dfs(matrix [][]byte, seq []byte, p Point, dir bool) {
	s.stack.push(p)
	if s.stack.len() == len(seq) {
		return
	}

	if dir {
		for i := 0; i < len(matrix[0]); i++ {
			candidate := Point{p.X, i}
			if s.stack.len() < len(seq) && matrix[p.X][i] == seq[s.stack.len()] && !s.stack.exists(candidate) {
				s.dfs(matrix, seq, candidate, false)
			}
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			candidate := Point{i, p.Y}
			if s.stack.len() < len(seq) && matrix[i][p.Y] == seq[s.stack.len()] && !s.stack.exists(candidate) {
				s.dfs(matrix, seq, candidate, true)
			}
		}
	}

	if s.stack.len() < len(seq) {
		s.stack.pop()
	}
}
