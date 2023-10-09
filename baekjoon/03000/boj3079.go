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
	gate    []int // 각 심사대에서 심사를 마치는데 걸리는 시간
)

// 난이도: Gold 5
// 메모리: 3336KB
// 시간: 80ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	gate = make([]int, N)
	for i := 0; i < N; i++ {
		gate[i] = scanInt()
	}
}

func Solve() {
	l, r := 1, 1000000000000000000 // 1 <= N*Tk <= 10**18
	for l <= r {
		mid := (l + r) / 2 // 매개 변수 mid를 상근이와 친구들이 모두 심사대를 통과하는데 필요한 시간이라고 가정

		passed := 0 // 심사대를 통과한 사람의 수

		// mid 시간 안에 동시에 통과할 수 있는 사람들의 수 구하기
		// mid / gate[i] (=각 심사대에서 걸리는 시간)을 누적해서 더한다
		for i := 0; i < N; i++ {
			passed += mid / gate[i]
			// mid 시간 안에 충분히 M명 이상이 통과할 수 있는 경우
			if passed >= M {
				break
			}
		}

		if passed >= M {
			// mid 시간 안에 충분히 M명 이상이 통과할 수 있는 경우
			// 더 짧은 시간에 조건을 만족할 수 있는지 확인
			r = mid - 1
		} else {
			// mid 시간 안에  M명이 통과할 수 없는 경우
			//
			l = mid + 1
		}
	}

	// 마지막으로 변경된 r은 M명 이상이 심사대를 통과할 수 있는 가장 짧은 시간 t에서 1을 뺀 값이며
	// 마지막으로 변경된 l값은 마지막으로 변경된 r보다 항상 1만큼 크므로 l값은 t와 동일
	// 즉, l은 M명 이상이 심사대를 통과할 수 있는 가장 짧은 시간이다
	fmt.Fprintln(writer, l)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
