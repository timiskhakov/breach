package breach

func Solve(matrix [][]byte, seq []byte) []Point {
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == seq[0] {
			s := stack{}
			dfs(matrix, seq, &s, Point{0, i}, false)
			if s.Len() == len(seq) {
				return s.Slice()
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
		_ = s.Pop()
	}
}
