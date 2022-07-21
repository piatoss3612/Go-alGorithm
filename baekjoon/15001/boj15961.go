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
	dishes     []int     // N개의 회전 초밥
	slide      []int     // 연속해서 먹는 k개의 접시
	kind       [3001]int // 초밥의 접시 번호 카운팅
	N, d, k, c int
)

// 메모리: 37824KB
// 시간: 424ms
// 슬라이딩 윈도우
// 왜 80%, 90%에서 틀리나 했는데 뒤쪽에서 다시 앞쪽 인덱스 번호로 연결되는 부분을 확인하지 않아서 였다
// 다른 사람들은 O(N*k)로도 푸는 것 같은데... 요즘 왜 이렇게 열받지?
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	N, d, k, c = scanInt(), scanInt(), scanInt(), scanInt()
	dishes = make([]int, N)
	slide = make([]int, k)

	duplicated := 0 // 중복되는 접시의 수
	ans := 0

	// 슬라이딩 윈도우 기법을 사용해 0번째 접시부터 k-1번째 접시까지 먹는 경우부터 시작해
	// N-1번째 접시에서 0, 1, 2, ... k-2번째 접시까지 먹는 경우까지 따져봐야 한다
	for i := 0; i < N+k-1; i++ {
		turn := i % k // slide의 값을 변경할 인덱스

		// 슬라이드에서 제거할 접시의 번호가 0이 아닌 경우
		// 접시 번호에 해당하는 접시의 개수를 감소시키고
		// 접시의 개수가 1보다 크다면 중복된 접시가 제거된 것이므로
		// duplicated값을 1만큼 감소시킨다
		if slide[turn] != 0 {
			kind[slide[turn]]--
			if kind[slide[turn]] >= 1 {
				duplicated--
			}
		}

		// 슬라이드에 새로운 접시가 추가되는 경우는
		// i가 N보다 작은 경우와 N 이상인 경우로 나뉜다
		if i < N {
			dishes[i] = scanInt()
			slide[turn] = dishes[i]
		} else {
			slide[turn] = dishes[i%N]
		}

		// 추가된 접시의 번호에 해당하는 접시 개수를 증가시키고
		// 접시 개수가 1보다 크다면 중복된 접시이므로
		// duplicated값을 1만큼 증가시킨다
		kind[slide[turn]]++
		if kind[slide[turn]] > 1 {
			duplicated++
		}

		// 연속해서 먹은 접시의 가짓수가 k개 미만인 경우
		if i < k-1 {
			continue
		}

		// 쿠폰으로 먹는 접시는 1번 행사로 먹은 k개의 접시들과 연속해서 먹지 않아도 된다
		// 연속해서 먹은 접시에 쿠폰으로 먹을 수 있는 접시가 포함되어 있지 않다면 가짓수를 1증가시킨다
		if kind[c] == 0 {
			ans = max(ans, k-duplicated+1)
		} else {
			ans = max(ans, k-duplicated)
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
