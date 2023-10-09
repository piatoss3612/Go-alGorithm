package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	T        int
	K, M, P  int
	inDegree []int   // 진입 차수
	edge     [][]int // edge[a][b]: a -> b
)

// 난이도: Gold 3
// 메모리: 936 KB
// 시간: 4ms
// 분류: 위상 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	K, M, P = scanInt(), scanInt(), scanInt()
	inDegree = make([]int, M+1)
	edge = make([][]int, M+1)

	for i := 1; i <= P; i++ {
		a, b := scanInt(), scanInt()
		inDegree[b] += 1             // 노드 b의 진입 차수 증가
		edge[a] = append(edge[a], b) // 노드 a에서 b로 가는 간선 추가
	}
}

type Strahler struct {
	maxOrderIn int
	count      int
}

// 위상 정렬을 이용하여 strahler 순서를 구한다.
func Solve() {
	strahler := make([]Strahler, M+1)

	// 노드 번호와 strahler 순서를 저장하는 큐
	q := []struct {
		nodeNumber int
		order      int
	}{}

	// 진입 차수가 0인 노드를 큐에 삽입, strahler 순서는 1
	for i := 1; i <= M; i++ {
		if inDegree[i] == 0 {
			strahler[i].maxOrderIn = 1
			q = append(q, struct {
				nodeNumber int
				order      int
			}{i, 1})
		}
	}

	ans := 0 // strahler 순서의 최댓값

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		ans = max(ans, x.order) // strahler 순서의 최댓값 갱신

		// 현재 집입 차수가 0인 노드와 연결된 노드들의 진입 차수를 감소시키고 strahler 순서를 갱신
		for _, next := range edge[x.nodeNumber] {
			// 현재 노드의 strahler 순서가 다음 노드의 strahler 순서보다 크면 갱신
			if strahler[next].maxOrderIn < x.order {
				strahler[next].maxOrderIn = x.order
				strahler[next].count = 1
				// 현재 노드의 strahler 순서가 다음 노드의 strahler 순서와 같으면 count 증가
			} else if strahler[next].maxOrderIn == x.order {
				strahler[next].count++
			}

			inDegree[next]-- // 다음 노드의 진입 차수 감소

			// 다음 노드의 진입 차수가 0이면 큐에 삽입
			if inDegree[next] == 0 {
				order := strahler[next].maxOrderIn
				// 다음 노드로 들어가는 강의 순서가 가장 큰 강의 개수가 2개 이상이면 다음 노드의 순서는 order+1
				if strahler[next].count >= 2 {
					order += 1
				}

				q = append(q, struct {
					nodeNumber int
					order      int
				}{next, order})
			}
		}
	}

	fmt.Fprintln(writer, K, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
