package bj1992

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	tree    [][]string
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	tree = make([][]string, n)
	for i := 0; i < n; i++ {
		tree[i] = strings.Split(scanString(), "")
	}
	quadTree(0, 0, n)
	fmt.Fprintln(writer)
}

func quadTree(x, y, n int) {
	tmpZero, tmpOne := 0, 0
	for i := x; i < x+n; i++ {
		for j := y; j < y+n; j++ {
			if tree[i][j] == "1" {
				tmpOne += 1
			} else {
				tmpZero += 1
			}
		}
	}
	if tmpZero == 0 {
		fmt.Fprint(writer, 1)
	} else if tmpOne == 0 {
		fmt.Fprint(writer, 0)
	} else {
		fmt.Fprint(writer, "(")
		quadTree(x, y, n/2)
		quadTree(x, y+n/2, n/2)
		quadTree(x+n/2, y, n/2)
		quadTree(x+n/2, y+n/2, n/2)
		fmt.Fprint(writer, ")")
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
