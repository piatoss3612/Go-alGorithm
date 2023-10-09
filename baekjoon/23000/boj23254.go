package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M    int
	score   []int
)

type Option struct {
	subject  int // 시험 과목 번호
	increase int // 시간당 점수 증가량
	remain   int // 기본 점수를 제외하고 최대 얻을 수 있는 점수
}

// 우선순위 큐 정의 및 인터페이스 구현
type Study []*Option

func (s Study) Len() int { return len(s) }
func (s Study) Less(i, j int) bool {
	if s[i].increase == s[j].increase {
		return s[i].remain > s[j].remain
	}
	return s[i].increase > s[j].increase
}
func (s Study) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s *Study) Push(x interface{}) {
	*s = append(*s, x.(*Option))
}
func (s *Study) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

// 메모리: 	14148KB
// 시간: 316ms
// 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	score = make([]int, M+1)
	// 기본 점수 입력
	for i := 1; i <= M; i++ {
		score[i] = scanInt()
	}

	study := &Study{}
	heap.Init(study)

	// 우선순위 큐에 과목 번호, 남은 점수, 점수 증가량 정보 입력
	for i := 1; i <= M; i++ {
		opt := Option{
			subject:  i,
			increase: scanInt(),
			remain:   100 - score[i],
		}
		heap.Push(study, &opt)
	}

	time := N * 24
	for time > 0 && len(*study) > 0 {
		// 그리디 알고리즘:
		// 점수 증가량이 가장 큰 과목부터 공부함으로써 최종 성적의 최댓값을 구할 수 있다
		x := heap.Pop(study).(*Option)

		q := x.remain / x.increase // 남은 점수를 점수 증가량으로 나눈 값

		if q > 0 {
			// 1. q가 0보다 큰 경우는 q시간 만큼 공부를 할 수 있다는 의미

			// 남은 시간이 q보다 작은 경우는 q를 남은 시간으로 변경
			if time < q {
				q = time
			}

			x.remain -= q * x.increase         // 남은 점수 갱신
			score[x.subject] += q * x.increase // q시간 만큼 공부함으로써 증가한 점수 갱신
			time -= q                          // 남은 시간 갱신

			if x.remain > 0 {
				// 얻을 수 있는 점수가 아직 남아있는 경우
				// 우선순위 큐로 x를 다시 집어넣는다
				heap.Push(study, x)
			}
		} else {
			// 2. q가 0인 경우는 x.remain이 x.increase보다 작아서 시간을 채워서 공부할 수 없다는 의미

			/*
				x.remain이 k일 때, 우선순위 큐에 k보다 작은 점수 증가량을 가진 과목만 남아 있는 경우에
				이 3점을 가져가는 것이 최종 성적의 최댓값에 유의미한 영향을 미치므로
				점수 증가량을 남은 점수로 갱신하고 우선순위 큐에 다시 넣어줘야 한다
			*/

			/*
				앞에서 분명 x.remain이 0보다 큰 경우에만 우선순위 큐에 다시 집어넣었는데
				divideByZero 런타임에러가 발생해서 x.remain이 0인지 체크
			*/
			if x.remain != 0 {
				x.increase = x.remain
				heap.Push(study, x)
			}
		}
	}

	ans := 0
	for i := 1; i <= M; i++ {
		ans += score[i]
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
