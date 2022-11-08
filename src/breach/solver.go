package breach

func Solve(matrix [][]byte, seq []byte) []Point {
	s := stack{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == seq[0] {
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

	return s.slice()
}

func dfs(matrix [][]byte, seq []byte, s *stack, p Point, dir bool) {
	s.push(p)
	if s.len() == len(seq) {
		return
	}

	if dir {
		for i := 0; i < len(matrix[0]); i++ {
			candidate := Point{p.X, i}
			if s.len() < len(seq) && matrix[p.X][i] == seq[s.len()] && !s.exists(candidate) {
				dfs(matrix, seq, s, candidate, false)
			}
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			candidate := Point{i, p.Y}
			if s.len() < len(seq) && matrix[i][p.Y] == seq[s.len()] && !s.exists(candidate) {
				dfs(matrix, seq, s, candidate, true)
			}
		}
	}

	if s.len() < len(seq) {
		s.pop()
	}
}
