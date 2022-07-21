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
	inp     [100001]int
	n       int
)

// 메모리: 2500KB
// 시간: 24ms
// 6549번과 동일한 풀이: 분할 정복
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
	}

	fmt.Fprintln(writer, DAC(1, n))
}

func DAC(left, right int) int {
	if left == right {
		return inp[left]
	}

	mid := (left + right) / 2
	ret := max(DAC(left, mid), DAC(mid+1, right))

	lo, hi := mid, mid+1
	minHeight := min(inp[lo], inp[hi])
	ret = max(ret, minHeight*2)

	for lo > left || hi < right {
		if hi < right && (lo == left || inp[lo-1] < inp[hi+1]) {
			hi += 1
			minHeight = min(minHeight, inp[hi])
		} else {
			lo -= 1
			minHeight = min(minHeight, inp[lo])
		}
		ret = max(ret, (hi-lo+1)*minHeight)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
