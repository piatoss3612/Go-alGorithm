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
	N, C    int
	house   []int
)

// 난이도: Gold 4
// 메모리: 6008KB
// 시간: 68ms
// 분류: 이분 탐색, 매개 변수 탐색
// 이 문제 맞은 적이 없는데 왜 업로드되어 있었는지 모르겠네요
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, C = scanInt(), scanInt()
	house = make([]int, N+1)
	for i := 1; i <= N; i++ {
		house[i] = scanInt()
	}
	sort.Ints(house) // 오름차순 정렬
}

func Solve() {
	l, r := 0, 1000000000
	for l <= r {
		m := (l + r) / 2
		// 인접한 공유기 사이의 최대 거리를 m으로 잡았을 때 공유기의 개수가 C개 이상인 경우와 아닌 경우
		if routers(m) {
			l = m + 1 // 최대 거리를 늘려서 재탐색
		} else {
			r = m - 1 // 최대 거리를 줄여서 재탐색
		}
	}
	fmt.Fprintln(writer, r) // upper bound 출력
}

func routers(expected int) bool {
	cnt := 1
	prev := house[1]

	for i := 2; i <= N; i++ {
		if house[i]-prev >= expected {
			cnt += 1
			prev = house[i]
		}

		if cnt >= C {
			return true
		}
	}

	return cnt >= C
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
