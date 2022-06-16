package main

import (
	"bufio"
	"container/heap"
	"fmt"
	_ "math"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

// 과제 정보
type Homework struct {
	deadline, score int
	index           int
}

type Schedule []*Homework // 점수의 합이 최대가 되는 과제들을 스케줄링할 우선순위 큐

// 정렬 인터페이스 구현
func (s Schedule) Len() int { return len(s) }
func (s Schedule) Less(i, j int) bool {
	if s[i].score == s[j].score {
		return s[i].deadline < s[j].deadline
	}
	return s[i].score < s[j].score
}
func (s Schedule) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].index, s[j].index = i, j
}

// 우선순위 큐 푸시, 팝 메서드
func (s *Schedule) Push(x interface{}) {
	*s = append(*s, x.(*Homework))
}
func (s *Schedule) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

// 메모리: 956KB
// 시간: 4ms
// Go 언어에서 우선순위 큐를 사용하기 위해 필요한 인터페이스의 구현에 익숙해지기 위한 반복 연습
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	inp := make([]*Homework, n)
	for i := 0; i < n; i++ {
		inp[i] = &Homework{scanInt(), scanInt(), i}
	}

	sort.Slice(inp, func(i, j int) bool {
		if inp[i].deadline == inp[j].deadline {
			return inp[i].score > inp[j].score
		}
		return inp[i].deadline < inp[j].deadline
	})

	res := &Schedule{}
	heap.Init(res)

	for i := 0; i < n; i++ {
		if len(*res) < inp[i].deadline {
			heap.Push(res, inp[i])
		} else if len(*res) == inp[i].deadline {
			y := heap.Pop(res).(*Homework)
			if y.score < inp[i].score {
				heap.Push(res, inp[i])
			} else {
				heap.Push(res, y)
			}
		}
	}

	total := 0
	for len(*res) > 0 {
		total += heap.Pop(res).(*Homework).score
	}

	fmt.Fprintln(writer, total)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
