package solver

type point struct {
	x, y int
}

func Solve(matrix [][]byte, sequence []byte) []point {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == sequence[0] {
				s := stack{}
				dfs(matrix, sequence, &s, point{i, j}, i != 0)
				if s.Len() == len(sequence) {
					return s.Slice()
				}
			}
		}
	}

	return []point{}
}

func dfs(matrix [][]byte, sequence []byte, s *stack, p point, hrz bool) {
	s.Push(p)
	if s.Len() == len(sequence) {
		return
	}

	if hrz {
		for i := 0; i < len(matrix[0]); i++ {
			if s.Len() < len(sequence) && matrix[p.x][i] == sequence[s.Len()] && !s.Exists(point{p.x, i}) {
				dfs(matrix, sequence, s, point{p.x, i}, false)
			}
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			if s.Len() < len(sequence) && matrix[i][p.y] == sequence[s.Len()] && !s.Exists(point{i, p.y}) {
				dfs(matrix, sequence, s, point{i, p.y}, true)
			}
		}
	}

	if s.Len() < len(sequence) {
		_ = s.Pop()
	}
}
