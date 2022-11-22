package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, M       int
	difficulty []int    // 각 몬스터의 난이도
	defeated   []bool   // 몬스터를 잡았는지 여부
	itemFor    [][]Item // i번 아이템을 얻었을 때, 획득 난이도에 영향을 받는 아이템들
	guide      *Guide   // 정수를 위한 가이드, 우선 순위 큐
	p, a, b, t int
)

type Monster struct {
	number     int // 몬스터 번호
	difficulty int // 몬스터 난이도
}

// 우선 순위 큐
// 난이도가 가장 쉬운 몬스터를 기준으로 정렬
type Guide []Monster

func (g Guide) Len() int { return len(g) }
func (g Guide) Less(i, j int) bool {
	return g[i].difficulty < g[j].difficulty
}
func (g Guide) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g *Guide) Push(x interface{}) {
	*g = append(*g, x.(Monster))
}
func (g *Guide) Pop() interface{} {
	old := *g
	n := len(old)
	x := old[n-1]
	*g = old[:n-1]
	return x
}

type Item struct {
	number        int // 아이템 번호
	difficultyInc int // 아이템이 없는 경우에 증가하는 난이도
}

// 난이도: Gold 3
// 메모리: 43000KB
// 시간: 356ms
// 분류: 우선 순위 큐

// 처음에는 몬스터 번호와 몬스터 정보를 포인터로 매핑(map[int]*Monster)하여 몬스터를 잡을 때마다
// 몬스터의 난이도를 조절하려 했지만, 몬스터 난이도가 변경될 때마다
// 우선 순위 큐의 초기화를 새로해줘야 하는 문제로 인해 시간 초과 발생

// -> 갱신된 몬스터의 번호와 난이도를 우선 순위 큐에 새로 추가하고
// 이미 잡은 몬스터인지 여부를 확인해 줌으로써 아직 잡지 않은 몬스터만 골라 잡으면 된다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	difficulty = make([]int, N+1)
	defeated = make([]bool, N+1)
	itemFor = make([][]Item, N+1)
	// 난이도 입력
	for i := 1; i <= N; i++ {
		difficulty[i] = scanInt()
	}

	p = scanInt()
	for i := 1; i <= p; i++ {
		a, b, t = scanInt(), scanInt(), scanInt()
		difficulty[b] += t                          // a번 아이템이 없으므로 b번 몬스터의 난이도 t만큼 증가
		itemFor[a] = append(itemFor[a], Item{b, t}) // a번 아이템을 얻으면 b번 아이템을 얻기 위한 난이도가 t만큼 줄어드는 관계 설명
	}

	// 새로운 가이드 생성, 우선 순위 큐 초기화
	guide = new(Guide)
	for i := 1; i <= N; i++ {
		heap.Push(guide, Monster{i, difficulty[i]})
	}
}

func Solve() {
	defeatedCount := 0 // 잡은 몬스터의 수
	maxDifficulty := 0 // 가이드를 따라 게임을 클리어했을 때의 최대 난이도

	// 가이드에 잡을 몬스터가 남아있으며 잡은 몬스터의 수가 M보다 작은 경우
	for len(*guide) > 0 && defeatedCount < M {
		monster := heap.Pop(guide).(Monster)

		if defeated[monster.number] {
			// 이미 해당 번호의 몬스터를 잡은 경우
			continue
		}

		// 몬스터를 잡고 난이도 및 잡은 몬스터의 수 갱신
		defeated[monster.number] = true
		maxDifficulty = max(maxDifficulty, monster.difficulty)
		defeatedCount++

		// 몬스터를 잡고 아이템을 획득했으므로 해당 아이템과 관련된 몬스터들의 난이도 조정
		for _, next := range itemFor[monster.number] {
			difficulty[next.number] -= next.difficultyInc
			heap.Push(guide, Monster{next.number, difficulty[next.number]}) // 우선 순위 큐에 새롭게 갱신된 난이도의 몬스터 정보를 추가
		}
	}

	fmt.Fprintln(writer, maxDifficulty) // 게임 클리어 난이도 출력
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
