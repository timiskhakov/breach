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

func (s *Solver) dfs(matrix [][]byte, seq []byte, node Point, hrz bool) {
	s.stack.push(node)
	if s.stack.len() == len(seq) {
		return
	}

	if hrz {
		for i := 0; i < len(matrix[0]); i++ {
			candidate := Point{node.X, i}
			if s.stack.len() < len(seq) && matrix[node.X][i] == seq[s.stack.len()] && !s.stack.exists(candidate) {
				s.dfs(matrix, seq, candidate, false)
			}
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			candidate := Point{i, node.Y}
			if s.stack.len() < len(seq) && matrix[i][node.Y] == seq[s.stack.len()] && !s.stack.exists(candidate) {
				s.dfs(matrix, seq, candidate, true)
			}
		}
	}

	if s.stack.len() < len(seq) {
		s.stack.pop()
	}
}
