package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	stations []station
	N, L, P  int
)

type station struct {
	dist int
	oil  int
}

// 메모리: 1324KB
// 시간: 12ms
// 그리디 알고리즘, 최대 힙
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	N = scanInt()
	stations = make([]station, N)
	for i := 0; i < N; i++ {
		stations[i] = station{scanInt(), scanInt()}
	}

	// 정글은 일직선이고, 성경이의 트럭과 주유소도 모두 일직선 위에 있다. 주유소는 모두 성경이의 트럭을 기준으로 오른쪽에 있다.
	// 조건에 따르면 특정 위치에 오직 하나의 주유소만 존재한다는 것을 알 수 있다
	// 따라서 입력값을 거리를 기준으로 오름차순으로 정렬해준다
	sort.Slice(stations, func(i, j int) bool {
		return stations[i].dist < stations[j].dist
	})

	L, P = scanInt(), scanInt()

	supply := &Supply{}
	heap.Init(supply) // 최대 힙 초기화

	count := 0
	idx := 0

	/*
		그리디 알고리즘,
		마을에 도달하기 위해 주유소에 멈추는 횟수의 최솟값은 어떻게 구할 수 있을까?

		1. 현재 사용할 수 있는 연료로 이동할 수 있는 한계까지 이동
		2. 이동하는 과정에서 발견한 각각의 주유소에서 충전할 수 있는 연료를 최대 힙에 저장
		3. 연료가 부족하다면 최대 힙에 저장된 연료를 꺼내오고 방문 횟수를 1늘린다
		4. 1~3번 반복
	*/

	// 이동 가능한 거리(총 연료) < 마을까지의 거리
	for P < L {
		// 1, 2번 실행
		for idx < N && stations[idx].dist <= P {
			heap.Push(supply, stations[idx].oil)
			idx++
		}

		// 예외 처리: 마을에 도착하지 못했는데 연료는 부족하고 들릴 수 있는 주유소도 없다면
		if supply.Len() == 0 {
			fmt.Fprintln(writer, -1)
			return
		}

		// 3번 실행
		P += heap.Pop(supply).(int)
		count++
	}

	fmt.Fprintln(writer, count)
}

// 최대 힙
type Supply []int

func (s Supply) Len() int { return len(s) }
func (s Supply) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s Supply) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s *Supply) Push(x interface{}) {
	*s = append(*s, x.(int))
}
func (s *Supply) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
