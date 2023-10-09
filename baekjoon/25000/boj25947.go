package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N, B, A int
	arr     []int
	sum     []int
)

// 난이도: Silver 1
// 메모리: 4120KB
// 시간: 44ms
// 분류: 정렬, 누적 합, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	// 연속으로 선물을 골라야 한다는 조건이 없으므로 오름차순으로 정렬 후 누적합니 구한다.
	N, B, A = scanInt(), scanInt(), scanInt()
	arr = make([]int, N)
	sum = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}
	sort.Ints(arr)

	sum[0] = arr[0]
	for i := 1; i < N; i++ {
		sum[i] = sum[i-1] + arr[i]
	}
}

func Solve() {
	ans := 0
	for i := 0; i < N; i++ {
		// 0~i번째 선물을 고른다

		// 고른 선물의 개수가 A개 이하인 경우
		if i <= A-1 {
			// A개 이하인 경우는 모든 선물을 할인받을 수 있으므로 할인받은 금액이 B보다 작거나 같은지만 확인하면 된다.
			if sum[i]/2 <= B {
				ans = i + 1
			} else {
				break
			}
		} else { // 고른 선물의 개수가 A개 초과인 경우
			// A개 초과인 경우는 금액이 가장 비싼 A개의 선물을 할인받는다.
			// 따라서 (선물 가격의 합) - (가장 비싼 A개의 선물 가격의 합) / 2 가 B보다 작거나 같은지 확인하면 된다.
			if sum[i]-(sum[i]-sum[i-A])/2 <= B {
				ans = i + 1
			} else {
				break
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", ans) // 정답 출력
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
