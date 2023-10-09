package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	sudoku   [9][9]int // 스도쿠 판 (0이면 빈칸)
	numInRow [9][10]bool // numInRow[i][j] = i번째 행에 숫자 j가 있는지 여부
	numInCol [9][10]bool // numInCol[i][j] = i번째 열에 숫자 j가 있는지 여부
	numInBox [9][10]bool // numInBox[i][j] = i번째 박스에 숫자 j가 있는지 여부
	blanks   [][2]int // 빈칸의 좌표 (행, 열)
)

// 난이도: Gold 4
// 메모리: 940KB
// 시간: 156ms
// 분류: 구현, 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	for i := 0; i < 9; i++ {
		line := scanString()
		for j := 0; j < 9; j++ {
			x := int(line[j] - '0')
			sudoku[i][j] = x

			// 빈칸이면 blanks에 추가하고, 아니면 numInRow, numInCol, numInBox 갱신
			if x == 0 {
				blanks = append(blanks, [2]int{i, j})
			} else {
				numInRow[i][x] = true
				numInCol[j][x] = true
				numInBox[getBoxNum(i, j)][x] = true
			}
		}
	}
}

func Solve() {
	fill(0) // 백트래킹으로 빈칸을 채움

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fprintf(writer, "%d", sudoku[i][j])
		}
		fmt.Fprintln(writer)
	}
}

func fill(target int) bool {
	// 빈칸을 모두 채웠으면 true 반환
	if target == len(blanks) {
		return true
	}

	row, col := blanks[target][0], blanks[target][1] // target번째 빈칸의 좌표

	for i := 1; i <= 9; i++ {
		// i를 채울 수 있으면 채우고 다음 빈칸을 채움
		if canFill(row, col, i) {
			toggle(row, col, i)
			sudoku[row][col] = i

			// 다음 빈칸을 채울 수 있으면 true 반환
			if fill(target + 1) {
				return true
			}

			// 다음 빈칸을 채울 수 없으면 i를 지우고 다음 i를 채움
			sudoku[row][col] = 0
			toggle(row, col, i)
		}
	}

	return false
}

// row, col에 num을 채울 수 있는지 여부 반환
func canFill(row, col, num int) bool {
	return !numInRow[row][num] && !numInCol[col][num] && !numInBox[getBoxNum(row, col)][num]
}

// row, col에 num을 채우거나 지움
func toggle(row, col, num int) {
	numInRow[row][num] = !numInRow[row][num]
	numInCol[col][num] = !numInCol[col][num]
	numInBox[getBoxNum(row, col)][num] = !numInBox[getBoxNum(row, col)][num]
}

// row, col이 속한 박스 번호 반환
func getBoxNum(row, col int) int {
	return (row/3)*3 + (col / 3)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
