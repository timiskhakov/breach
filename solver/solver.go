package solver

import "fmt"

type point struct {
	x, y int
}

type Solver struct {
}

func (s *Solver) Solve(matrix [][]byte, sequence []byte) {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == sequence[0] {
				fmt.Println(fmt.Sprintf("slv: %d, %d", i, j))

				isHor := true
				if i == 0 {
					isHor = false
				}

				visited := []point{{i, j}}
				dfs(matrix, sequence, visited, point{i, j}, 1, isHor)
				return
			}
		}
	}
}

func dfs(matrix [][]byte, sequence []byte, visited []point, p point, n int, isHor bool) {
	if len(visited) == len(sequence) {
		return
	}

	if isHor {
		for i := 0; i < len(matrix[0]); i++ {
			if matrix[p.x][i] == sequence[n] && !isVisited(visited, point{p.x, i}) {
				fmt.Println(fmt.Sprintf("dfs: %d, %d", p.x, i))
				visited = append(visited, point{p.x, i})
				dfs(matrix, sequence, visited, point{p.x, i}, n+1, false)
				break
			}
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][p.y] == sequence[n] && !isVisited(visited, point{i, p.y}) {
				fmt.Println(fmt.Sprintf("dfs: %d, %d", i, p.y))
				visited = append(visited, point{i, p.y})
				dfs(matrix, sequence, visited, point{i, p.y}, n+1, true)
				break
			}
		}
	}
}

func isVisited(visited []point, p point) bool {
	for _, v := range visited {
		if v == p {
			return true
		}
	}
	return false
}
