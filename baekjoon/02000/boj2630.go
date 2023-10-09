package bj2630

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	table      [][]int
	bCnt, wCnt int = 0, 0
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	table = make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
		for j := 0; j < n; j++ {
			table[i][j] = scanInt()
		}
	}

	Search(0, 0, n)

	fmt.Fprintln(writer, wCnt)
	fmt.Fprintln(writer, bCnt)
}

func Search(x, y, n int) {
	cntOne := 0
	cntZero := 0
	for i := x; i < x+n; i++ {
		for j := y; j < y+n; j++ {
			if table[i][j] == 1 {
				cntOne += 1
			} else {
				cntZero += 1
			}
		}
	}
	if cntOne == 0 {
		wCnt += 1
	} else if cntZero == 0 {
		bCnt += 1
	} else {
		Search(x, y, n/2)
		Search(x+n/2, y, n/2)
		Search(x, y+n/2, n/2)
		Search(x+n/2, y+n/2, n/2)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
