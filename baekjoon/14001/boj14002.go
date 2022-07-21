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
	n       int
	input   []int
	dp      [1001]int
	choice  [1001]int
)

// 메모리: 940KB
// 시간: 40ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	input = make([]int, n+1)
	for i := 1; i <= n; i++ {
		input[i] = scanInt()
	}

	// 0에서 시작하면 반복문의 필요없이 한 번만 호출하면 된다
	res := rec(0)
	fmt.Fprintln(writer, res)

	idx := choice[0]
	for idx != -1 {
		fmt.Fprintf(writer, "%d ", input[idx])
		idx = choice[idx]
	}
	fmt.Fprintln(writer)
}

func rec(start int) int {
	ret := &dp[start]
	if *ret != 0 {
		return *ret
	}

	if start == 0 {
		*ret = -1 // 인덱스 0의 값은 포함되지 않으므로 오히려 -1로 초기화한다
	}
	next := -1 // 다음에 올 최적의 수의 인덱스를 저장

	for i := start + 1; i <= n; i++ {
		if input[start] < input[i] {
			temp := rec(i) + 1
			if temp > *ret {
				*ret = temp
				next = i
			}
		}
	}
	choice[start] = next
	return *ret
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
