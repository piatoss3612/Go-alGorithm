package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner       = bufio.NewScanner(os.Stdin)
	writer        = bufio.NewWriter(os.Stdout)
	D, P, C, F, S int
	paths         []Path // 도로 및 제트 비행 경로
	earn          []int  // 최대 수입
)

type Path struct {
	from, to, cost int
}

// 난이도: Gold 3
// 메모리: 980KB
// 시간: 4ms
// 분류: 벨만-포드 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

const INF = -987654321 // 최댓값을 구하는 문제이므로 최솟값을 무한대로 설정

func Setup() {
	D, P, C, F, S = scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
	paths = make([]Path, P+F)

	// 도로
	for i := 1; i <= P; i++ {
		paths = append(paths, Path{scanInt(), scanInt(), D})
	}

	// 비행 경로
	for i := 1; i <= F; i++ {
		paths = append(paths, Path{scanInt(), scanInt(), -scanInt() + D}) // 수입 = 기본 수입(D) - 비행 비용
	}

	// 최대 수입 초기화
	earn = make([]int, C+1)
	for i := 1; i <= C; i++ {
		earn[i] = INF
	}
	earn[S] = D // 출발지의 최대 수입은 기본 수입(D)
}

func Solve() {
	// 수입이 무한으로 증가하는 사이클이 존재하는 경우
	if BellmanFord() {
		fmt.Fprintln(writer, -1)
	}

	// 수입이 무한으로 증가하는 사이클이 존재하지 않는 경우 최대 수입 출력
	// 벨만 포드 알고리즘을 사용하여 최대 C번(2 <= C <= 220) 완화를 실행하므로
	// S와 연결되어 있지 않은 도시의 수입은 항상 D(1 <= D <= 1000,INF + C*D를 한들 0 보다 커질 수 없음)보다 작다
	// 따라서 S가 다른 도시와 연결되어 있지 않은 경우의 최댓값은 D가 되므로
	// 다른 도시와 연결되어 있는지 여부는 확인할 필요가 없다
	ans := 0
	for i := 1; i <= C; i++ {
		ans = max(ans, earn[i])
	}
	fmt.Fprintln(writer, ans)
}

func BellmanFord() bool {
	// C-1번 반복
	for i := 1; i <= C-1; i++ {
		// 수입이 갱신되지 않으면 종료
		if !relax() {
			return false
		}
	}

	// C번째 반복
	return relax()
}

func relax() bool {
	relaxed := false

	// 모든 경로에 대해 최대 수입 갱신
	for _, path := range paths {
		if earn[path.to] < earn[path.from]+path.cost {
			earn[path.to] = earn[path.from] + path.cost
			relaxed = true
		}
	}

	return relaxed
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
