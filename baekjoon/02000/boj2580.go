package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	board      [9][9]int
	numInRows  [9][10]bool
	numInCols  [9][10]bool
	numInBoxes [9][10]bool
	blanks     [][2]int
)

// 난이도: Gold 4
// 메모리: 960KB
// 시간: 84ms
// 분류: 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = scanInt()
			if board[i][j] == 0 {
				blanks = append(blanks, [2]int{i, j})
			} else {
				numInRows[i][board[i][j]] = true
				numInCols[j][board[i][j]] = true
				numInBoxes[boxNum(i, j)][board[i][j]] = true
			}
		}
	}
}

func Solve() {
	fill(0)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fprintf(writer, "%d ", board[i][j])
		}
		fmt.Fprintln(writer)
	}
}

func fill(pos int) bool {
	if pos == len(blanks) {
		return true
	}

	r, c := blanks[pos][0], blanks[pos][1]

	for i := 1; i <= 9; i++ {
		if canFill(r, c, i) {
			toggle(r, c, i)
			board[r][c] = i

			if fill(pos + 1) {
				return true
			}

			board[r][c] = 0
			toggle(r, c, i)
		}
	}

	return false
}

func boxNum(r, c int) int {
	return r/3 + (c/3)*3
}

func canFill(r, c, n int) bool {
	return !numInRows[r][n] &&
		!numInCols[c][n] && !numInBoxes[boxNum(r, c)][n]
}

func toggle(r, c, n int) {
	bn := boxNum(r, c)
	numInRows[r][n] = !numInRows[r][n]
	numInCols[c][n] = !numInCols[c][n]
	numInBoxes[bn][n] = !numInBoxes[bn][n]
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
