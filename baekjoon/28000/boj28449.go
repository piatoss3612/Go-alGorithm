package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner             = bufio.NewScanner(os.Stdin)
	writer              = bufio.NewWriter(os.Stdout)
	N, M                int
	HI, ARC             []int
	HIWin, ARCWin, Draw int
)

// 난이도: Gold 5
// 메모리: 4132KB
// 시간: 76ms
// 분류: 이분 탐색, 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	HI = make([]int, N)
	ARC = make([]int, M)

	for i := 0; i < N; i++ {
		HI[i] = scanInt()
	}

	for i := 0; i < M; i++ {
		ARC[i] = scanInt()
	}

	// 정렬 필수
	sort.Ints(HI)
	sort.Ints(ARC)
}

func Solve() {
	for i := 0; i < N; i++ {
		binarySearch(i)
	}

	fmt.Fprintln(writer, HIWin, ARCWin, Draw)
}

func binarySearch(idx int) {
	// ARC팀에서 HI[idx]보다 작거나 같은 수의 최대 개수를 찾는다 (upper bound)
	l1, r1 := 0, M-1
	for l1 <= r1 {
		mid := (l1 + r1) / 2
		if ARC[mid] > HI[idx] {
			r1 = mid - 1
		} else {
			l1 = mid + 1
		}
	}

	ARCWin += M - (r1 + 1) // ARC팀의 승리 횟수를 갱신

	// ARC팀에서 HI[idx]보다 작은 수의 최대 개수를 찾는다 (upper bound)
	l2, r2 := 0, M-1
	for l2 <= r2 {
		mid := (l2 + r2) / 2
		if ARC[mid] >= HI[idx] {
			r2 = mid - 1
		} else {
			l2 = mid + 1
		}
	}

	Draw += r1 - r2 // 무승부는 ARC팀에서 HI[idx]보다 작거나 같은 수의 개수에서 HI[idx]보다 작은 수의 개수를 뺀다
	HIWin += r2 + 1 // HI팀의 승리 횟수를 갱신
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
