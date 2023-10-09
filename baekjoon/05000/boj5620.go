package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	points  [][2]int
)

// 난이도: Platinum 3
// 메모리: 46568KB
// 시간: 788ms
// 분류: 분할 정복, 라인 스위핑
// 분할 정복 풀이
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	points = make([][2]int, N)
	for i := 0; i < N; i++ {
		points[i] = [2]int{scanInt(), scanInt()}
	}
	// x좌표 기준 정렬
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
}

func Solve() {
	// 0번 좌표부터 N-1번 좌표까지의 최단거리를 분할 정복으로 구한다.
	fmt.Fprintln(writer, ClosestPair(0, N-1))
}

func ClosestPair(left, right int) int {
	// 기저 사례: 점이 1개일 경우, 아주 큰 값을 반환한다.
	if left == right {
		return math.MaxInt64
	}

	// 기저 사례: 점이 2개일 경우, 두 점 사이의 거리를 반환한다.
	if left+1 == right {
		return Distance(points[left], points[right])
	}

	mid := (left + right) / 2 // 중앙 인덱스
	ret := min(ClosestPair(left, mid), ClosestPair(mid+1, right)) // 왼쪽 절반과 오른쪽 절반에 각각 포함된 좌표들 사이의 최단거리를 구한다.

	// band에는 중앙(m번째 좌표)의 x축과의 거리가 ret보다 작은 좌표들을 저장한다. 
	// => ret보다 큰 경우는 거리를 구할 필요가 없다. 
	band := make([][2]int, 0, right-left+1) 
	for i := left; i <= right; i++ {
		dist := (points[i][0] - points[mid][0]) * (points[i][0] - points[mid][0])
		if dist < ret {
			band = append(band, points[i])
		}
	}

	// band의 좌표들을 y축 기준으로 정렬한다.
	sort.Slice(band, func(i, j int) bool {
		return band[i][1] < band[j][1]
	})

	// band의 좌표들 사이의 최단거리를 구한다.
	for i := 0; i < len(band)-1; i++ {
		for j := i + 1; j < len(band); j++ {
			dist := Distance(band[i], band[j])
			if dist >= ret {
				break
			}
			ret = dist
		}
	}

	return ret
}

func Distance(a, b [2]int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
