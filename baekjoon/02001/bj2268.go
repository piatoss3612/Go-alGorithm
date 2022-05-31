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
	seg     [21000000]int
	n, m    int
)

// 메모리: 31268KB
// 시간: 792ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	var a, b, c int
	for i := 1; i <= m; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()

		if a == 0 {
			// b가 c보다 큰 경우에는 c~b 구간합을 구한다
			if b > c {
				fmt.Fprintln(writer, Sum(c, b, 1, n, 1))
			} else {
				fmt.Fprintln(writer, Sum(b, c, 1, n, 1))
			}
		} else {
			Modify(b, c, 1, n, 1)
		}
	}
}

func Sum(left, right, nodeLeft, nodeRight, node int) int {
	if left > nodeRight || right < nodeLeft {
		return 0
	}

	if left <= nodeLeft && nodeRight <= right {
		return seg[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	// 여기서 인덱스를 잘못 지정해서 런타임 에러가 발생했다
	return Sum(left, right, nodeLeft, mid, node*2) + Sum(left, right, mid+1, nodeRight, node*2+1)
}

func Modify(index, newVal, left, right, node int) {
	if index < left || index > right {
		return
	}

	if left == right {
		seg[node] = newVal
		return
	}

	mid := (left + right) / 2
	Modify(index, newVal, left, mid, node*2)
	Modify(index, newVal, mid+1, right, node*2+1)
	seg[node] = seg[node*2] + seg[node*2+1]
	return
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
