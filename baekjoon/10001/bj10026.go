package bj10026

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	RGB        [][]string
	visitedRGB [][]bool
	visitedRB  [][]bool
	n          int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	RGB = make([][]string, n)
	visitedRGB = make([][]bool, n)
	visitedRB = make([][]bool, n)
	for i := 0; i < n; i++ {
		RGB[i] = strings.Split(scanString(), "")
		visitedRGB[i] = make([]bool, n)
		visitedRB[i] = make([]bool, n)
	}

	cntRGB, cntRB := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !visitedRGB[i][j] {
				BFS_RGB(i, j)
				cntRGB += 1
			}
			if !visitedRB[i][j] {
				BFS_RB(i, j)
				cntRB += 1
			}
		}
	}
	fmt.Fprintf(writer, "%d %d\n", cntRGB, cntRB)
}

func BFS_RGB(x, y int) {
	visitedRGB[x][y] = true

	if validRGB(x-1, y) && RGB[x-1][y] == RGB[x][y] {
		BFS_RGB(x-1, y)
	}

	if validRGB(x, y-1) && RGB[x][y-1] == RGB[x][y] {
		BFS_RGB(x, y-1)
	}

	if validRGB(x, y+1) && RGB[x][y+1] == RGB[x][y] {
		BFS_RGB(x, y+1)
	}

	if validRGB(x+1, y) && RGB[x+1][y] == RGB[x][y] {
		BFS_RGB(x+1, y)
	}
}

func BFS_RB(x, y int) {
	visitedRB[x][y] = true

	if validRB(x-1, y) {
		if isRG(x, y) && isRG(x-1, y) {
			BFS_RB(x-1, y)
		} else if RGB[x-1][y] == RGB[x][y] {
			BFS_RB(x-1, y)
		}
	}

	if validRB(x, y-1) {
		if isRG(x, y) && isRG(x, y-1) {
			BFS_RB(x, y-1)
		} else if RGB[x][y-1] == RGB[x][y] {
			BFS_RB(x, y-1)
		}
	}

	if validRB(x, y+1) {
		if isRG(x, y) && isRG(x, y+1) {
			BFS_RB(x, y+1)
		} else if RGB[x][y+1] == RGB[x][y] {
			BFS_RB(x, y+1)
		}
	}

	if validRB(x+1, y) {
		if isRG(x, y) && isRG(x+1, y) {
			BFS_RB(x+1, y)
		} else if RGB[x+1][y] == RGB[x][y] {
			BFS_RB(x+1, y)
		}
	}
}

func isRG(x, y int) bool {
	if RGB[x][y] == "R" || RGB[x][y] == "G" {
		return true
	}
	return false
}

func validIdx(v int) bool {
	if v >= 0 && v < n {
		return true
	}
	return false
}

func validRGB(x, y int) bool {
	if validIdx(x) && validIdx(y) && !visitedRGB[x][y] {
		return true
	}
	return false
}

func validRB(x, y int) bool {
	if validIdx(x) && validIdx(y) && !visitedRB[x][y] {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
