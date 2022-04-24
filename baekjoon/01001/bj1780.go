package bj1780

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner                  = bufio.NewScanner(os.Stdin)
	writer                   = bufio.NewWriter(os.Stdout)
	table                    [][]int
	moneCnt, zeroCnt, oneCnt int = 0, 0, 0
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
	fmt.Fprintln(writer, moneCnt)
	fmt.Fprintln(writer, zeroCnt)
	fmt.Fprintln(writer, oneCnt)
}

func Search(x, y, n int) {
	cntOne := 0
	cntZero := 0
	cntMone := 0
	for i := x; i < x+n; i++ {
		for j := y; j < y+n; j++ {
			if table[i][j] == 1 {
				cntOne += 1
			} else if table[i][j] == 0 {
				cntZero += 1
			} else {
				cntMone += 1
			}
		}
	}
	if cntZero == 0 && cntMone == 0 {
		oneCnt += 1
	} else if cntMone == 0 && cntOne == 0 {
		zeroCnt += 1
	} else if cntOne == 0 && cntZero == 0 {
		moneCnt += 1
	} else {
		Search(x, y, n/3)             // 1
		Search(x, y+n/3, n/3)         // 2
		Search(x, y+(n/3)*2, n/3)     // 3
		Search(x+n/3, y, n/3)         // 4
		Search(x+n/3, y+n/3, n/3)     // 5
		Search(x+n/3, y+(n/3)*2, n/3) // 6
		Search(x+(n/3)*2, y, n/3)     // 7
		Search(x+(n/3)*2, y+n/3, n/3) // 8
		Search(x+(n/3)*2, y+(n/3)*2, n/3)
	}
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

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
