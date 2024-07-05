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
	cows    []Cow
)

type Cow struct {
	Count  int
	Output int
}

// 14527번: Paired Up
// hhttps://www.acmicpc.net/problem/14527
// 난이도: 골드 5
// 메모리: 2464 KB
// 시간: 56 ms
// 분류: 정렬, 두 포인터, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	cows = make([]Cow, N)
	for i := 0; i < N; i++ {
		cows[i] = Cow{
			Count:  scanInt(),
			Output: scanInt(),
		}
	}
}

func Solve() {
	// 소를 두 마리씩 짝지어서 우유를 생산
	// 두 소의 우유 생산량의 합이 작업에 필요한 시간
	// 이 시간이 최소가 되도록 짝을 지어야 함
	// 따라서 가장 작은 우유 생산량을 가진 소와 가장 큰 우유 생산량을 가진 소를 짝지어주는 것이 최적 (그리디)
	// 소의 생산량을 기준으로 오름차순 정렬 후, 두 포인터를 사용하여 양 끝에서 짝을 지어주면서 최대 시간의 최솟값을 구함

	// 생산량을 기준으로 오름차순 정렬
	sort.Slice(cows, func(i, j int) bool {
		return cows[i].Output < cows[j].Output
	})

	maxTime := 0

	l, r := 0, N-1
	for l <= r {
		// l과 r이 같은 경우
		if l == r {
			maxTime = max(maxTime, cows[l].Output*2)
			break
		}

		pair := min(cows[l].Count, cows[r].Count) // 두 소 중 작은 수만큼 짝을 지어줌
		time := cows[l].Output + cows[r].Output   // 짝을 지어준 두 소의 우유 생산량의 합 = 작업에 필요한 시간

		maxTime = max(maxTime, time) // 최대 시간의 최솟값을 구함
		cows[l].Count -= pair        // 짝을 지어준 소의 수만큼 각 소의 마릿수를 감소시킴
		cows[r].Count -= pair        // 짝을 지어준 소의 수만큼 각 소의 마릿수를 감소시킴

		if cows[l].Count == 0 { // 소의 마릿수가 0이 되면 다음 소로 이동
			l++
		}

		if cows[r].Count == 0 { // 소의 마릿수가 0이 되면 다음 소로 이동
			r--
		}
	}

	fmt.Fprintln(writer, maxTime)
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
