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
	N, M    int
	moves   []Move // 수상택시가 되돌아가야 하는 구간의 정보
)

type Move struct {
	start, end int
}

// 난이도: Gold 3
// 메모리: 27244KB
// 시간: 220ms
// 분류: 스위핑
// 회고:
// 1. 입력값이 정렬되어 있는 상태라는 보장이 없는데 정렬없이 입력 순서대로 문제를 풀려고 해서 틀림
// 2. 구간이 끝나는 지점(e)의 값을 최댓값으로 갱신하지 않아 틀림
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()

	// N명의 사람이 수상택시를 타는 위치 s와 목적지 e를 입력
	for i := 1; i <= N; i++ {
		s, e := scanInt(), scanInt()
		// 시작 위치 0에서 M까지 반드시 이동해야 하는 거리 M을 제외하고
		// 되돌아가야 하는 구간의 이동 거리만을 고려해야 한다
		// 연산의 편의를 위해 구간의 시작 위치와 끝나는 위치를 반전시켜 moves 슬라이스에 집어넣는다
		if s > e {
			moves = append(moves, Move{e, s})
		}
	}

	// 되돌아가야 하는 구간 정보 정렬
	// 1. 시작 위치가 가장 작은 구간을 기준으로 오름차순 정렬
	// 2. 시작 위치가 같다면 끝나는 위치가 가장 작은 구간이 먼저 오도록 정렬
	sort.Slice(moves, func(i, j int) bool {
		if moves[i].start == moves[j].start {
			return moves[i].end < moves[j].end
		}
		return moves[i].start < moves[j].start
	})

	// 이동 거리의 최솟값을 구하기 위해 스위핑 기법을 사용

	s, e := 0, 0 // 구간의 시작과 끝
	ans := 0     // 단방향 거리의 누적합

	for _, move := range moves {
		if move.start <= e {
			// s~e와 move.start~move.end 두 구간 사이에 겹치는 구간이 존재하는 경우
			// 구간이 끝나는 위치값을 갱신
			e = max(e, move.end)
		} else {
			// 겹치는 구간이 없는 경우
			// 기존 구간의 단방향 거리 누적
			ans += e - s
			//  시작 위치와 목적지를 move 구간으로 새롭게 갱신
			s = move.start
			e = move.end
		}
	}

	ans += e - s                  // 마지막에 갱신한 구간의 거리 누적
	fmt.Fprintln(writer, ans*2+M) // 이동 거리의 최솟값 = 누적된 단방향 거리 * 2 + 반드시 이동해야 하는 거리 M
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
