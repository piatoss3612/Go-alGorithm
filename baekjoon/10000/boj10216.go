package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	T, N    int
	inp     []*camp
	parent  []int
)

// 적군 진영 정보
type camp struct {
	x, y, r int
}

// 메모리: 	6768KB
// 시간: 	3608ms
// 분리 집합
// 4연속 시간 초과를 맞고 float64 타입과 math 패키지 사용이 연산 속도에 영향을 줄 것이라고 생각하여 빼버렸다
// 입출력도 라인 단위로 개선해보고 슬라이스도 혹시 몰라서 테스트 케이스가 끝날 때 비워주었다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	T = scanInt()

	for i := 1; i <= T; i++ {
		TestCase()
	}
}

func TestCase() {
	N = scanInt()
	inp = make([]*camp, N)
	parent = make([]int, N)

	for i := 0; i < N; i++ {
		parent[i] = i
		inp[i] = scanCamp()
	}

	ans := N // 진영의 개수

	// 시간 복잡도 O(N^2)
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			dist_camp := getDist(i, j)                              // 캠프 사이의 거리의 제곱값
			dist_r := (inp[i].r + inp[j].r) * (inp[i].r + inp[j].r) // 통신 반경 거리의 합의 제곱값
			pi, pj := find(i), find(j)
			// 통신 반경 거리의 합의 제곱값이 캠프 사이의 거리의 제곱값보다 크거나 같다면
			// 캠프 i와 캠프 j는 통신이 가능한 같은 진영에 속한다
			// 또한 캠프 i와 캠프 j가 아직 같은 진영에 속해 있지 않은 경우에만
			// union 연산을 해주고 2개의 캠프가 1개의 진영으로 병합되었으므로
			// 진영의 개수를 줄인다
			if dist_camp <= dist_r && pi != pj {
				parent[pj] = pi
				ans--
			}
		}
	}

	fmt.Fprintln(writer, ans)

	// 테스트 케이스가 종료되었으므로 inp와 parent 슬라이스를 비워준다
	inp = nil
	parent = nil
}

// 좌표 사이의 거리는 루트 연산을 하지 않은 값을 반환
func getDist(i, j int) int {
	return (inp[i].x-inp[j].x)*(inp[i].x-inp[j].x) + (inp[i].y-inp[j].y)*(inp[i].y-inp[j].y)
}

func find(x int) int {
	if parent[x] == x {
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

func scanCamp() *camp {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	x, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])
	r, _ := strconv.Atoi(fields[2])
	return &camp{x, y, r}
}
