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
	seg     [2100000]int
	n, q    int
)

// 메모리: 32296KB
// 시간: 108ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, q = scanInt(), scanInt()

	var a, b, c int
	for i := 1; i <= q; i++ {
		a, b, c = scanInt(), scanInt(), scanInt()
		if a == 1 {
			Update(b, c, 1, n, 1)
		} else {
			fmt.Fprintln(writer, Sum(b, c, 1, n, 1))
		}
	}
}

func Sum(left, right, nodeLeft, nodeRight, node int) int {
	if right < nodeLeft || nodeRight < left {
		return 0
	}

	if left <= nodeLeft && nodeRight <= right {
		return seg[node]
	}

	mid := (nodeLeft + nodeRight) / 2
	return Sum(left, right, nodeLeft, mid, node*2) + Sum(left, right, mid+1, nodeRight, node*2+1)
}

// 수입, 지출 내용을 추가한다는 점에 주의!
func Update(index, newVal, left, right, node int) int {
	if left > index || right < index {
		return seg[node]
	}

	if left == right {
		seg[node] += newVal // 새로 입력된 값을 추가
		return seg[node]
	}

	mid := (left + right) / 2
	seg[node] = Update(index, newVal, left, mid, node*2) + Update(index, newVal, mid+1, right, node*2+1)
	return seg[node]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
