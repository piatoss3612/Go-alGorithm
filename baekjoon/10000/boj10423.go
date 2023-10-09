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
	N, M, K int
	powered []bool  // i번 도시가 발전소와 연결되어 있는지 여부
	parent  []int   // i번 도시와 연결된 부모 도시, 부모 도시는 발전소와 연결되어 있어야 한다
	cables  []Cable // 설치 가능한 케이블의 정보
	ans     int     // 모든 도시에 전기를 공급할 수 있도록 케이블을 설치하는 데 드는 최소비용
)

type Cable struct {
	u, v, w int
}

// 난이도: Gold 2
// 메모리: 4568KB
// 시간: 288ms
// 분류: 최소 스패닝 트리
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M, K = scanInt(), scanInt(), scanInt()
	powered = make([]bool, N+1)
	parent = make([]int, N+1)
	// 부모 도시를 자기자신으로 초기화
	for i := 1; i <= N; i++ {
		parent[i] = i
	}
	// 발전소가 설치된 도시 입력
	for i := 1; i <= K; i++ {
		powered[scanInt()] = true
	}
	cables = make([]Cable, M)
	// 설치 가능한 케이블 정보 입력
	for i := 0; i < M; i++ {
		cables[i] = Cable{scanInt(), scanInt(), scanInt()}
	}

	// 비용이 가장 적게 드는 케이블을 기준으로 오름차순 정렬
	sort.Slice(cables, func(i, j int) bool {
		return cables[i].w < cables[j].w
	})
}

func Solve() {
	remain := N - K // 아직 전기가 공급되지 않은 도시의 수
	edges := 0      // 최소 스패닝 트리를 구성하는 간선의 개수

	// 이 문제의 경우는 최소 스패닝 트리를 구하는 것은 맞지만
	// N-1개의 간선으로 N개의 도시를 연결하기 전에
	// N개의 도시에 전력 공급이 가능하도록 케이블을 설치한다면
	// 궂이 N-1개의 간선을 설치할 필요가 없다
	for len(cables) > 0 && remain > 0 && edges < N-1 {
		for i := 0; i < len(cables); i++ {
			pu, pv := find(cables[i].u), find(cables[i].v) // u번 도시와 v번 도시와 연결된 부모 도시 탐색

			// pu번 도시와 pv번 도시가 아직 연결되어 있지 않고
			// 두 도시가 모두 전력이 연결되어 있는 상태가 아닌 경우에만
			// (두 도시가 모두 전력이 연결되어 있다면 연결할 필요가 없으므로)
			if pu != pv && !(powered[pu] && powered[pv]) {
				// pu번 도시에만 전력이 연결되어 있는 경우
				if powered[pu] {
					parent[pv] = pu                              // pv번 도시와 케이블 연결
					powered[pv] = true                           // pv번 도시에 전력 공급
					ans += cables[i].w                           // 연결한 케이블 비용 추가
					remain--                                     // 전력이 공급되지 않은 도시의 수 감소
					edges++                                      // 간선의 개수 추가
					cables = append(cables[:i], cables[i+1:]...) // 설치한 케이블을 케이블 목록에서 제거
					break
				}

				// pv번 도시에만 전력이 연결되어 있는 경우
				if powered[pv] {
					parent[pu] = pv
					powered[pu] = true
					ans += cables[i].w
					remain--
					edges++
					cables = append(cables[:i], cables[i+1:]...)
					break
				}
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func find(x int) int {
	if x == parent[x] {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
