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
	MAX     = 100001
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := scanInt(), scanInt()
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}

	l, r := 0, 0    // 두 포인터
	sum := input[0] // l~r까지의 부분합을 저장
	ans := MAX
	for l <= r && r < n {
		// l~r까지의 부분합이 m보다 큰 경우
		if sum >= m {
			// ans를 최솟값으로 갱신
			ans = min(r-l+1, ans)
			// l을 1만큼 늘려야 하므로 현재 l값을 인덱스로 가지는 입력값을 부분합에서 빼준다
			sum -= input[l]
			l += 1
		} else {
			// r이 입력값의 마지막 인덱스에 도달한 경우
			if r == n-1 {
				// l을 1늘려주고 부분합을 수정한다
				sum -= input[l]
				l += 1
				continue
			}
			// r을 1늘리고 부분합에 r값을 인덱스로 가지는 입력값을 더해준다
			r += 1
			sum += input[r]
		}
	}

	if ans == MAX {
		fmt.Fprintln(writer, 0)
		return
	}
	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
