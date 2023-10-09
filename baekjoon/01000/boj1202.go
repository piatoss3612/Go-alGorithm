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
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, K       int
	gems       *Gems       // 보석 정보
	candidates *Candidates // 가방에 담을 수 있는 무게보다 작거나 같은 무게의 보석들의 가격
	bags       []int       // 가방의 무게 한도
)

type Gem struct {
	weight int
	value  int
}

// 우선순위 큐 정의 및 인터페이스 구현
type Gems []*Gem

func (g Gems) Len() int { return len(g) }
func (g Gems) Less(i, j int) bool {
	return g[i].weight < g[j].weight // pop 연산을 실행할 경우, 보석의 무게가 적게 나가는 것을 먼저 꺼내온다
}
func (g Gems) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
func (g *Gems) Push(x interface{}) {
	*g = append(*g, x.(*Gem))
}
func (g *Gems) Pop() interface{} {
	old := *g
	n := len(old)
	x := old[n-1]
	*g = old[:n-1]
	return x
}

// 최대 힙 정의 및 인터페이스 구현
type Candidates []int

func (c Candidates) Len() int { return len(c) }
func (c Candidates) Less(i, j int) bool {
	return c[i] > c[j] // pop 연산을 실행할 경우, 가방에 들어갈 수 있는 보석들 중 가장 높은 가격을 가진 보석을 꺼내온다
}
func (c Candidates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c *Candidates) Push(x interface{}) {
	*c = append(*c, x.(int))
}
func (c *Candidates) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

// 난이도: Gold 2
// 메모리: 	23056KB
// 시간: 648ms
// 그리디 알고리즘, 우선순위 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()

	gems = new(Gems)
	candidates = new(Candidates)
	heap.Init(gems)
	heap.Init(candidates)

	// 보석 정보 입력
	for i := 1; i <= N; i++ {
		heap.Push(gems, &Gem{scanInt(), scanInt()})
	}

	bags = make([]int, K)

	// 가방의 무게 한도 입력
	for i := 0; i < K; i++ {
		bags[i] = scanInt()
	}

	// 가방의 무게 한도를 오름차순으로 정렬
	sort.Ints(bags)

	ans := 0

	// 그리디 알고리즘: 상덕이가 훔칠 수 있는 보석 가격의 합의 최댓값 구하기
	// 1. i번 가방에 담을 수 있는 무게 한도 bags[i] 이하의 무게를 가진 보석들을 후보자로 선출하고
	// 2. 후보자들 중 가장 가치가 큰 보석을 i번 가방에 담는다
	// 이 때, 가방은 가장 작은 것부터 시작하여 오름차순으로 진행함으로써
	// i+1번째 가방에 담을 수 있는 후보자들이 i번 가방의 후보자들을 포함하도록 한다
	// 3. 1과 2를 반복한다
	for i := 0; i < K; i++ {
		for len(*gems) > 0 {
			g := heap.Pop(gems).(*Gem)

			// 보석의 무게가 가방의 무게 한도보다 작거나 같은 경우
			if g.weight <= bags[i] {
				heap.Push(candidates, g.value) // 후보 목록에 추가
			} else {
				// 보석의 무게가 가방의 무게 한도보다 큰 경우
				heap.Push(gems, g) // 보석을 우선순위 큐에 되돌려 놓고 반복문 종료
				break
			}
		}

		// i번 가방에 담을 수 있는 보석들(candidates) 중 가장 가치가 큰 보석을 담는다
		// i번 가방의 후보자들은 없을 수도 있다 (len(*candidates) = 0)
		if len(*candidates) > 0 {
			ans += heap.Pop(candidates).(int)
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
