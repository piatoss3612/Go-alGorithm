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
	N       int
	points  [][2]int
)

// 난이도: platinum 2
// 메모리: 11148KB
// 시간: 188ms
// 분류: 분할 정복
// 5620번과 동일하지만 풀이가 약간 다름
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

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
}

func Solve() {
	fmt.Fprintln(writer, ClosestPair(0, N-1))
}

func ClosestPair(left, right int) int {
	if left+1 == right {
		return Distance(points[left], points[right])
	}

	if left+2 == right {
		return min(Distance(points[left], points[left+1]), Distance(points[left+1], points[right]))
	}

	mid := (left + right) / 2
	ret := min(ClosestPair(left, mid), ClosestPair(mid+1, right))

	band := make([][2]int, 0, right-left+1)

	for i := left; i <= right; i++ {
		diff := points[i][0] - points[mid][0]
		if diff*diff < ret {
			band = append(band, points[i])
		}
	}

	sort.Slice(band, func(i, j int) bool {
		return band[i][1] < band[j][1]
	})

	for i := 0; i < len(band)-1; i++ {
		for j := i + 1; j < len(band); j++ {
			/*
				5620번에서 사용한 방식
				dist := Distance(band[i], band[j])
				if dist >= ret {
					break
				}
				ret = dist
			*/


			// 5620번과 다른 방식
			// 5620번 방식으로는 입력 데이터의 수가 많아지면 잘못된 값이 나오는데
			// 이 방식으로는 잘못된 값이 나오지 않음 -> 왜 인지는 아직 모름...
			// 아무래도 좌표가 중복되는 점이 존재하는 경우가 있어서 그런 것 같음
			diff := band[i][1] - band[j][1]
			if diff*diff < ret {
				ret = min(ret, Distance(band[i], band[j]))
			} else {
				break
			}
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