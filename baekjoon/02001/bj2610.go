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
	n, m    int
	graph   [101][101]int
	checked [101]bool
	INF     = 10000000
)

// 전처리: 플로이드 와샬 알고리즘을 적용하기 위해 그래프 초기화
func init() {
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if i != j {
				graph[i][j] = INF
			}
		}
	}
}

// 메모리: 992KB
// 시간: 8ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	var a, b int
	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		// 회의 참석자 a,b의 관계를 양방향 그래프로 표현
		graph[a][b] = 1
		graph[b][a] = 1
	}

	// 플로이드 와샬 알고리즘:
	// i->j로 의견이 전달되는 최단거리를 갱신
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	leaders := []int{} // 각 위원회의 리더를 저장

	for i := 1; i <= n; i++ {
		// i번째 사람이 어느 위원회에 속해있는지 이미 확인한 경우
		if checked[i] {
			continue
		}

		// 새로운 위원회 발견
		checked[i] = true                // i확인 처리
		leader := i                      // i가 속한 위원회의 리더를 i로 임시 설정
		maxDistance := getMaxDistance(i) // i가 위원회의 다른 회원들에게 의견을 전달하는 최장거리

		// i가 속한 위원회의 다른 회원들 탐색
		for j := 1; j <= n; j++ {
			// j가 i인 경우 또는
			// j가 속한 위원회를 이미 확인한 경우 또는
			// i에서 j로 의견을 전달할 수 없는 경우는 건너뛴다
			if i == j || checked[j] || graph[i][j] == INF {
				continue
			}

			// i와 같은 위원회에 속해있는 j를 찾는다
			checked[j] = true                 // j확인 처리
			tempDistance := getMaxDistance(j) // j가 위원회의 다른 회원들에게 의견을 전달하는 최장거리

			// j의 최장거리가 임시 리더의 최장거리보다 짧은 경우
			if tempDistance < maxDistance {
				leader = j                 // 리더를 j로 갱신
				maxDistance = tempDistance // 최장거리 갱신
			}
		}

		// 같은 위원회에 속한 모든 회원들을 탐색한 뒤

		leaders = append(leaders, leader) // 리더 목록에 방금 탐색한 위원회의 리더 추가
	}

	fmt.Fprintln(writer, len(leaders)) // 위원회의 수 = 리더의 수 출력

	sort.Ints(leaders) // 반드시 정렬 필요
	for _, l := range leaders {
		fmt.Fprintln(writer, l) // 오름차순으로 리더 번호 출력
	}
}

func getMaxDistance(x int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if graph[x][i] != INF {
			ans = max(ans, graph[x][i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
