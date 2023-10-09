package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	N, M      int
	intervals []Interval
)

type Interval struct {
	from, to int
}

// 난이도: Gold 3
// 메모리: 5868KB
// 시간: 108ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	intervals = make([]Interval, M)
	for i := 0; i < M; i++ {
		intervals[i] = Interval{scanInt(), scanInt()}
	}

	// 문제에서 정렬되어 있다는 표현이 없으므로 정렬 필수, 문제 조건에 따라 중복 구간은 존재하지 않음
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].from < intervals[j].from
	})
}

func Solve() {
	l, r := 1, 1000000000000000000
	for l <= r {
		m := (l + r) / 2
		if Check(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	fmt.Fprintln(writer, r) // 최댓값 (upper bound) 출력
}

// 사회적 거리의 하한이 expected일 때
// N마리 이상의 소들을 사회적 거리를 유지하면서 인터벌 위에 배치할 수 있는지 여부 확인
func Check(expected int) bool {
	prev := 0 // 마지막으로 소를 배치한 위치
	cnt := 0  // 배치한 소의 수

	for i := 0; i < M; i++ {
		itv := intervals[i]

		temp := prev + expected // 새롭게 소를 배치할 예상 위치

		// 소를 처음 배치하거나 예상 위치가 i번째 인터벌 위가 아닌 경우
		if prev == 0 || temp < itv.from {
			prev = itv.from
			cnt++
			// 예상 위치가 i번째 인터벌 위에 있는 경우
		} else if temp >= itv.from && temp <= itv.to {
			prev = temp
			cnt++
			// i번째 인터벌에 소를 배치할 수 없는 경우
		} else {
			continue
		}

		// 소를 배치하고 남은 구간에 얼마나 많은 소를 배치할 수 있는지 확인
		dist := itv.to - prev
		move := dist / expected
		prev += move * expected
		cnt += move

		// 이미 소를 N마리 이상 배치한 경우
		if cnt >= N {
			return true
		}
	}

	// 소를 N마리 이상 배치할 수 없는 경우
	return false
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
