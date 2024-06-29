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
	N       int
	lanes   []Lane
)

type Lane struct {
	S, E, C int
}

// 23740번: 버스 노선 개편하기
// hhttps://www.acmicpc.net/problem/23740
// 난이도: 골드 5
// 메모리: 29448 KB
// 시간: 184 ms
// 분류: 정렬, 스위핑
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	lanes = make([]Lane, N)
	for i := 0; i < N; i++ {
		lane := Lane{
			S: scanInt(),
			E: scanInt(),
			C: scanInt(),
		}
		lanes[i] = lane
	}

	sort.Slice(lanes, func(i, j int) bool {
		if lanes[i].S == lanes[j].S {
			return lanes[i].E < lanes[j].E
		}

		return lanes[i].S < lanes[j].S
	})
}

func Solve() {
	newLanes := make([]Lane, 0)
	s, e, c := lanes[0].S, lanes[0].E, lanes[0].C

	for i := 1; i < N; i++ {
		// 겹치는 경우
		if lanes[i].S <= e {
			e = max(e, lanes[i].E)
			c = min(c, lanes[i].C)
			continue
		}

		newLanes = append(newLanes, Lane{s, e, c})
		s, e, c = lanes[i].S, lanes[i].E, lanes[i].C
	}

	newLanes = append(newLanes, Lane{s, e, c})

	fmt.Fprintln(writer, len(newLanes))
	for _, lane := range newLanes {
		fmt.Fprintln(writer, lane.S, lane.E, lane.C)
	}
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
