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
	N, M    int
)

// 난이도: Silver 1
// 메모리: 8508KB
// 시간: 80ms
// 분류: 두 포인터, (큐)
// 변수명을 stack이라고 해놓고 큐로 풀었네...ㅎㅎ
func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
}

func Solve() {
	queue := make([]int, 0, N+1)
	total := 0 // 막은 구멍의 총 크기
	ans := 0   // 최대로 막을 수 있는 구멍의 크기
	for i := 1; i <= N; i++ {
		hole := scanInt()

		// 여태 막은 구멍의 총 크기와 다음 구멍의 크기를 더한 값이 M보다 큰 경우
		for total+hole > M && len(queue) > 0 {
			// 여태 막은 구멍 중에서 가장 앞에 있는 구멍 포기
			total -= queue[0]
			queue = queue[1:]
		}

		// 다음 구멍을 막을 여유가 있는 경우
		if total+hole <= M {
			// 햄스터의 몸을 늘려 구멍 막기
			queue = append(queue, hole)
			total += hole
			ans = max(ans, total) // 최댓값 갱신
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
