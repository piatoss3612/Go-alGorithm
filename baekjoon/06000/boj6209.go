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
	D, N, M int
	island  []int
)

// 난이도: Gold 3
// 메모리: 2032KB
// 시간: 24ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	D, N, M = scanInt(), scanInt(), scanInt()
	island = make([]int, N+1)
	for i := 0; i < N; i++ {
		island[i] = scanInt()
	}
	island[N] = D     // 탈출구 위치 추가
	sort.Ints(island) // 오름차순 정렬
}

func Solve() {
	// 점프할 수 있는 최소거리의 최댓값 X => 항상 점프하는 거리가 X 이상이어야 한다는 의미

	l, r := 1, 1000000000
	for l <= r {
		m := (l + r) / 2
		// 점프할 수 있는 최소거리의 최댓값이 m일 때 돌섬에서 탈출할 수 있는 경우
		if CanEscape(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	fmt.Fprintln(writer, r) // 점프할 수 있는 최소거리의 최댓값(upper bound) 출력
}

func CanEscape(expected int) bool {
	from := 0    // 현재 위치
	removed := 0 // 제거된 돌섬의 개수

	for i := 0; i <= N; i++ {
		// 현재 위치에서 i번째 돌섬까지의 거리가 expected보다 작은 경우
		if island[i]-from < expected {
			removed++ // i번째 돌섬 제거
			// 제거된 돌섬의 개수가 M개를 초과한 경우: 탈출구에 도달할 수 없다
			if removed > M {
				return false
			}
			continue
		}
		from = island[i] // i번째 돌섬으로 이동하여 현재 위치 갱신
	}
	// 제거된 돌섬의 개수가 M개 이하인 경우
	// 제거된 돌섬의 개수가 M이 되도록 임의의 돌섬을 제거하면
	// 점프할 수 있는 최소거리의 최댓값은 expected 이상이 되므로 항상 탈출할 수 있다
	return true
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
