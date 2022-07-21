package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	fenwick [1000001]int
	inp     [1000001]int
	n, m    int
)

// 메모리: 	34896KB
// 시간: 560ms
// 펜윅 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	var a, b, c int
	for i := 1; i <= m; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()

		if a == 0 {
			if b > c {
				fmt.Fprintln(writer, sum(b)-sum(c-1))
			} else {
				fmt.Fprintln(writer, sum(c)-sum(b-1))
			}
		} else {
			update(b, c-inp[b])
			inp[b] = c
		}
	}
}

func sum(pos int) int {
	ret := 0
	for pos > 0 {
		ret += fenwick[pos]
		pos &= pos - 1
	}
	return ret
}

func update(pos, val int) {
	for pos <= n {
		fenwick[pos] += val
		pos += (pos & -pos)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
