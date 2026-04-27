func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	return dfs(grid, visited, 0, 0)
}

func dfs(grid [][]int, visited [][]bool, r, c int) bool {
	m, n := len(grid), len(grid[0])

	if r == m-1 && c == n-1 {
		return true
	}

	visited[r][c] = true

	for _, dir := range getDirections(grid[r][c]) {
		nr := r + dir[0]
		nc := c + dir[1]

		if nr < 0 || nr >= m || nc < 0 || nc >= n {
			continue
		}

		if visited[nr][nc] {
			continue
		}

		if canConnect(grid[nr][nc], -dir[0], -dir[1]) {
			if dfs(grid, visited, nr, nc) {
				return true
			}
		}
	}

	return false
}

func canConnect(streetType int, needR, needC int) bool {
	for _, dir := range getDirections(streetType) {
		if dir[0] == needR && dir[1] == needC {
			return true
		}
	}

	return false
}

func getDirections(streetType int) [][]int {
	switch streetType {
	case 1:
		return [][]int{{0, -1}, {0, 1}} // left, right
	case 2:
		return [][]int{{-1, 0}, {1, 0}} // up, down
	case 3:
		return [][]int{{0, -1}, {1, 0}} // left, down
	case 4:
		return [][]int{{0, 1}, {1, 0}} // right, down
	case 5:
		return [][]int{{0, -1}, {-1, 0}} // left, up
	case 6:
		return [][]int{{0, 1}, {-1, 0}} // right, up
	}

	return [][]int{}
}
