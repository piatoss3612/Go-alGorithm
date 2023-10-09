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
	T, N    int
)

const MOD = 1000000007

// 메모리: 7080KB
// 시간: 312ms
// 그리디 알고리즘, 최소 힙
// 아~ 전생슬 아시는구나~ 리무르 겁.나. 셉니다~
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		TestCase()
	}
}

func TestCase() {
	N = scanInt()
	slimes := make(Slimes, 0, N) // 메모리 사용량이 N보다 커지지 않을 것이므로 N으로 초기화
	heap.Init(&slimes)           // 최소 힙 초기화

	for i := 1; i <= N; i++ {
		heap.Push(&slimes, scanInt())
	}

	cost := 1

	// 그리디 알고리즘:
	// 슬라임을 합성하는 비용을 최소화하려면 에너지가 적게 필요한 슬라임들부터 합성해야 한다
	for len(slimes) > 1 {
		a, b := heap.Pop(&slimes).(int), heap.Pop(&slimes).(int)
		temp := a * b
		heap.Push(&slimes, temp)
		cost *= temp % MOD // temp를 MOD로 나눈 나머지를 곱해준다
		cost %= MOD        // 오버플로우가 발생하지 않도록 확실하게
	}

	fmt.Fprintln(writer, cost)
}

type Slimes []int

func (s Slimes) Len() int { return len(s) }
func (s Slimes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Slimes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s *Slimes) Push(x interface{}) {
	*s = append(*s, x.(int))
}
func (s *Slimes) Pop() interface{} {
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
