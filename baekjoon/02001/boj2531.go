package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, d, k, c int
	eaten      []int
	dishes     []int
)

// 난이도: Silver 1
// 메모리: 1324KB
// 시간: 8ms
// 분류: 브루트포스 알고리즘, 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, d, k, c = scanInt(), scanInt(), scanInt(), scanInt()
	eaten = make([]int, d+1)
	dishes = make([]int, N+k)
	for i := 1; i <= N; i++ {
		dishes[i] = scanInt()
	}
	for i := N + 1; i < N+k; i++ {
		dishes[i] = dishes[i-N]
	}
}

func Solve() {
	successive := 0 // 연속으로 선택한 초밥 접시의 수
	noDupCnt := 0   // 초밥 종류의 중복없이 선택한 접시의 개수
	ans := 0        // 연속으로 k개의 초밥 접시를 선택했을 때, 초밥의 가짓수의 최댓값

	l, r := 1, 1 // 두 포인터
	for l < N {
		// 연속으로 k개의 초밥 접시를 선택한 경우
		if successive == k {
			// 쿠폰으로 먹을 수 있는 초밥을 선택하지 않은 경우
			if eaten[c] == 0 {
				ans = max(ans, noDupCnt+1)
			} else {
				ans = max(ans, noDupCnt)
			}

			// 포인터 l을 이동하기 전에
			// l에 있는 초밥을 선택하지 않은 것으로 처리
			eaten[dishes[l]]--
			if eaten[dishes[l]] == 0 {
				noDupCnt--
			}
			successive--
			l++
		}

		// 선택한 초밥 접시의 개수가 k보다 작은 경우
		for successive < k {
			// 포인터 r이 가리키는 초밥 접시 선택
			eaten[dishes[r]]++
			if eaten[dishes[r]] == 1 {
				noDupCnt++
			}
			successive++
			r++ // 포이터 r 이동
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
