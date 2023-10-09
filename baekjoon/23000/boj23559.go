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
	N, X    int
	sb      *SnackBar
)

// 우선순위 큐 정의 및 구현
type Menu struct {
	A, B int
}

type SnackBar []*Menu

func (s SnackBar) Len() int { return len(s) }
func (s SnackBar) Less(i, j int) bool {
	return s[i].A-s[i].B > s[j].A-s[j].B
}
func (s SnackBar) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s *SnackBar) Push(x interface{}) {
	*s = append(*s, x.(*Menu))
}
func (s *SnackBar) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

// 난이도: Gold 5
// 메모리: 7956KB
// 시간: 108ms
// 분류: 우선순위 큐, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, X = scanInt(), scanInt()
	sb = new(SnackBar)
	heap.Init(sb)

	for i := 1; i <= N; i++ {
		heap.Push(sb, &Menu{scanInt(), scanInt()})
	}
}

func Solve() {
	ans := 0

	// 그리디 알고리즘:
	// 주어지는 금액으로 코너 A의 메뉴(5000원)를 사먹을 수 있는 횟수는 한정적이다
	// 대부분은 코너 B의 1000원 짜리 메뉴를 선택해야 한다
	// 맛의 합을 최대화하기 위해서는 코너 A와 B의 메뉴의 맛 차이가 가장 큰 날부터 우선 순위로 탐색

	for len(*sb) > 0 {
		menu := heap.Pop(sb).(*Menu)

		// B코너의 메뉴가 더 맛있는 경우 또는 맛의 수치가 동일한 경우
		if menu.B >= menu.A {
			ans += menu.B
			X -= 1000
			continue
		}

		// A 코너의 메뉴를 선택하고 남은 일수를 모두 밥을 사먹을 수 있는 경우
		if X-5000 >= len(*sb)*1000 {
			ans += menu.A
			X -= 5000
		} else {
			ans += menu.B
			X -= 1000
		}
	}

	fmt.Fprintln(writer, ans)
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
