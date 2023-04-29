package main

// 1st solution - switch statement
// Runtime: 7ms
// Memory Usage: 3.2 MB
// Time complexity: O(n) -> n is the length of moves
// Space complexity: O(1)
func judgeCircle(moves string) bool {
	var u, d, l, r int

	for _, c := range moves {
		switch c {
		case 'U':
			u += 1
		case 'D':
			d += 1
		case 'L':
			l += 1
		case 'R':
			r += 1
		}
	}

	return (u == d) && (l == r)
}

// 2nd solution - map
// Runtime: 6ms
// Memory Usage: 3.2 MB
// Time complexity: O(n)
// Space complexity: O(1)
func judgeCircle2(moves string) bool {
	dir := map[byte]int{}

	for i := 0; i < len(moves); i++ {
		dir[moves[i]] += 1
	}

	return (dir['U'] == dir['D']) && (dir['L'] == dir['R'])
}
