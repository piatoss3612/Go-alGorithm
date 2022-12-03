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
	N, K    int
	count   [100001]int // count[i]: 값이 i인 원소의 개수
	inp     []int
)

// 난이도: Silver 1
// 메모리: 5168KB
// 시간: 44ms
// 분류: 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, K = scanInt(), scanInt()
	inp = make([]int, N+1)
	for i := 1; i <= N; i++ {
		inp[i] = scanInt()
	}
}

func Solve() {
	l, r := 1, 1 // 두 포인터 l과 r
	ans := 0     // 수열의 최대 길이
	for l <= r && r <= N {
		if count[inp[r]]+1 > K {
			// 포인터 r에 해당하는 원소를 수열에 포함했을 때,
			// 그 개수가 K보다 커지는 경우
			count[inp[l]]-- // 포인터 l에 해당하는 원소 제거
			l++             // 포인터 l을 오른쪽으로 한 칸 이동
		} else {
			count[inp[r]]++       // 포인터 r에 해당하는 원소를 수열에 포함
			ans = max(ans, r-l+1) // 수열의 최대 길이 갱신
			r++                   // 포인터 r을 오른쪽으로 한 칸 이동
		}
	}
	fmt.Fprintln(writer, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
