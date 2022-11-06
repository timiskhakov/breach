package breach

func Solve(matrix [][]byte, seq []byte) []Point {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == seq[0] {
				s := stack{}
				dfs(matrix, seq, &s, Point{i, j}, i != 0)
				if s.len() == len(seq) {
					if i != 0 {
						return append([]Point{{0, j}}, s.slice()...)
					}

					return s.slice()
				}
			}
		}
	}

	return []Point{}
}

func dfs(matrix [][]byte, seq []byte, s *stack, p Point, hrz bool) {
	s.push(p)
	if s.len() == len(seq) {
		return
	}

	if hrz {
		for i := 0; i < len(matrix[0]); i++ {
			candidate := Point{p.x, i}
			if s.len() < len(seq) && matrix[p.x][i] == seq[s.len()] && !s.exists(candidate) {
				dfs(matrix, seq, s, candidate, false)
			}
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			candidate := Point{i, p.y}
			if s.len() < len(seq) && matrix[i][p.y] == seq[s.len()] && !s.exists(candidate) {
				dfs(matrix, seq, s, candidate, true)
			}
		}
	}

	if s.len() < len(seq) {
		s.pop()
	}
}
