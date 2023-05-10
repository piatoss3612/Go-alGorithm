package main

// Runtime: 0ms
// Memory Usage: 2.1 MB
// Time complexity: O(n^2) -> 1 <= n <= 20
// Space complexity: O(n^2)
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	var visited [21][21]bool
	moves := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	isValid := func(y, x int) bool { return y >= 0 && y < n && x >= 0 && x < n }

	num := 1
	y, x := 0, 0
	direction := 0

	ans[y][x] = num
	visited[y][x] = true

	for num < n*n {
		var ny, nx int

		ny, nx = y+moves[direction][0], x+moves[direction][1]

		if isValid(ny, nx) && !visited[ny][nx] {
			num += 1
			ans[ny][nx] = num
			visited[ny][nx] = true
			y, x = ny, nx
			continue
		}

		direction = (direction + 1) % 4
		ny, nx = y+moves[direction][0], x+moves[direction][1]

		if isValid(ny, nx) && !visited[ny][nx] {
			num += 1
			ans[ny][nx] = num
			visited[ny][nx] = true
			y, x = ny, nx
			continue
		}

		break
	}

	return ans
}
