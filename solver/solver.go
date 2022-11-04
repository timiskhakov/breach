package solver

type point struct {
	x, y int
}

type Solver struct {
}

func (s *Solver) Solve(matrix [][]byte, sequence []byte) {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == sequence[0] {
				isHor := true
				if i == 0 {
					isHor = false
				}

				s := stack{}
				dfs(matrix, sequence, &s, point{i, j}, isHor)

				if s.Len() == len(sequence) {
					return
				}

				count++
			}
		}
	}
}

func dfs(matrix [][]byte, sequence []byte, s *stack, p point, isHor bool) {
	s.Push(p)
	if s.Len() == len(sequence) {
		return
	}

	if isHor {
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
