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
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	inp     []Usage
	taken   [100001]int // i번째 자리에 사람들이 몇 번 앉았다 일어났는지
	N       int
)

// PC 사용량
type Usage struct {
	start, end int
}

// 메모리: 8896KB
// 시간: 176ms
// 우선순위 큐, 최소 힙
// 왜 해병이 아니고 해군이지?
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	// PC 사용량 입력
	inp = make([]Usage, N)
	for i := 0; i < N; i++ {
		inp[i] = Usage{scanInt(), scanInt()}
	}

	// 문제 조건: 시작 시각이나 종료 시각이 다른 사람과 겹치는 경우는 없다

	// PC 사용 시작 시간을 기준으로 오름차순 정렬
	sort.Slice(inp, func(i, j int) bool {
		return inp[i].start < inp[j].start
	})

	sgb := &SGB{}   // 싸지방 PC 사용 정보 우선순위 큐
	seat := &Seat{} // 앉을 수 있는 자리 번호 최소 힙
	heap.Init(sgb)
	heap.Init(seat)

	need := 0 // 필요한 컴퓨터 수

	for i := 0; i < N; i++ {
		// i번째 사람이 컴퓨터를 사용하기 위해 싸지방에 들어갔을 때
		// 종료 시간이 지난 사람들을 싸지방에서 쫓아낸다
		for sgb.Len() > 0 {
			// temp는 싸지방을 이용하는 사람들 중 사용 종료 시간이 가장 빠른 사람
			temp := heap.Pop(sgb).(*PC)
			// temp의 사용 종료 시간이 i번째 사람의 시작 시간보다 작다면
			if temp.end <= inp[i].start {
				heap.Push(seat, temp.pos) // temp가 사용중인 컴퓨터 자리를 사용할 수 있다
			} else {
				heap.Push(sgb, temp) // 사용 시간이 남은 경우 싸지방에 되돌려 놓는다
				break
			}
		}

		// i번째 사람이 사용할 수 있는 컴퓨터가 없는 경우
		if seat.Len() == 0 {
			need++
			taken[need]++
			heap.Push(sgb, &PC{need, inp[i].end})
		} else {
			// 비어있는 자리 중에서 번호가 가장 작은 자리에 앉는다
			next := heap.Pop(seat).(int)
			taken[next]++
			heap.Push(sgb, &PC{next, inp[i].end})
		}
	}

	fmt.Fprintln(writer, need)
	for i := 1; i <= need; i++ {
		fmt.Fprintf(writer, "%d ", taken[i])
	}
	fmt.Fprintln(writer)
}

// 컴퓨터 사용 정보
type PC struct {
	pos int // 컴퓨터의 위치
	end int // 사용 종료 시간
}

// 싸지방 우선순위 큐
type SGB []*PC

func (s SGB) Len() int { return len(s) }
func (s SGB) Less(i, j int) bool {
	return s[i].end < s[j].end // 컴퓨터 사용 종료 시간을 기준으로 오름차순 정렬
}
func (s SGB) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s *SGB) Push(x interface{}) {
	*s = append(*s, x.(*PC))
}
func (s *SGB) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

// 최소 힙
type Seat []int

func (s Seat) Len() int { return len(s) }
func (s Seat) Less(i, j int) bool {
	// 문제 조건: 모든 사람은 싸지방에 들어왔을 때 비어있는 자리 중에서 번호가 가장 작은 자리에 앉는 것이 규칙
	return s[i] < s[j]
}
func (s Seat) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s *Seat) Push(x interface{}) {
	*s = append(*s, x.(int))
}
func (s *Seat) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
