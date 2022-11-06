package breach

func Solve(matrix [][]byte, seq []byte) []Point {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == seq[0] {
				s := stack{}
				dfs(matrix, seq, &s, Point{i, j}, i != 0)
				if s.Len() == len(seq) {
					if i != 0 {
						return append([]Point{{0, j}}, s.Slice()...)
					}

					return s.Slice()
				}
			}
		}
	}

	return []Point{}
}

func dfs(matrix [][]byte, seq []byte, s *stack, p Point, hrz bool) {
	s.Push(p)
	if s.Len() == len(seq) {
		return
	}

	if hrz {
		for i := 0; i < len(matrix[0]); i++ {
			candidate := Point{p.x, i}
			if s.Len() < len(seq) && matrix[p.x][i] == seq[s.Len()] && !s.Exists(candidate) {
				dfs(matrix, seq, s, candidate, false)
			}
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			candidate := Point{i, p.y}
			if s.Len() < len(seq) && matrix[i][p.y] == seq[s.Len()] && !s.Exists(candidate) {
				dfs(matrix, seq, s, candidate, true)
			}
		}
	}

	if s.Len() < len(seq) {
		s.Pop()
	}
}
