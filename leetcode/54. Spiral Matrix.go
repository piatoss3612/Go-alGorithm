package main

// Runtime: 0ms
// Memory Usage: 2MB
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	total := m * n

	visited := [11][11]bool{}
	move := map[int][2]int{
		0: [2]int{0, 1},  // right
		1: [2]int{1, 0},  // down
		2: [2]int{0, -1}, // left
		3: [2]int{-1, 0}, // up
	}

	valid := func(y, x int) bool {
		return y >= 0 && y < m && x >= 0 && x < n
	}

	result := make([]int, 0, m*n)

	y, x := 0, 0
	result = append(result, matrix[y][x])
	visited[y][x] = true

	direction := 0

	for len(result) < total {
		var ny, nx int

		next := move[direction]
		ny, nx = y+next[0], x+next[1]

		if valid(ny, nx) && !visited[ny][nx] {
			result = append(result, matrix[ny][nx])
			visited[ny][nx] = true
			y, x = ny, nx
			continue
		}

		// if cannot move, change direction
		direction = (direction + 1) % 4
		next = move[direction]
		ny, nx = y+next[0], x+next[1]

		if valid(ny, nx) && !visited[ny][nx] {
			result = append(result, matrix[ny][nx])
			visited[ny][nx] = true
			y, x = ny, nx
			continue
		}

		// if still cannot move, break
		break
	}

	return result
}
