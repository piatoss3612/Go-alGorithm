package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	T, N, B int
	parts   map[string][]Part
)

type Part struct {
	price   int
	quality int
}

// 난이도: Gold 3
// 메모리: 7276KB
// 시간: 72ms
// 분류: 이분 탐색, 매개 변수 탐색, 해시맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	T = scanInt()
	for i := 1; i <= T; i++ {
		TC()
	}
}

func TC() {
	N, B = scanSetupVars()
	parts = make(map[string][]Part)
	for i := 1; i <= N; i++ {
		t, _, p, q := scanPart()
		parts[t] = append(parts[t], Part{p, q})
	}

	for _, v := range parts {
		// 각 부품 타입의 부품들을 가격순으로 오름차순 정렬
		sort.Slice(v, func(i, j int) bool {
			return v[i].price < v[j].price
		})
	}

	Solve()
}

func Solve() {
	l, r := 0, 1000000000
	for l <= r {
		m := (l + r) / 2
		// 성능이 m인 컴퓨터를 구할 수 있는 경우
		if MeetPerformance(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	fmt.Fprintln(writer, r) // 최댓값(upper bound) 출력
}

func MeetPerformance(expected int) bool {
	totalCost := 0
	for _, part := range parts {
		bought := false // 구매하고자 하는 타입의 부품을 구매했는지 여부 확인
		for _, p := range part {
			// 부품의 성능이 expected 이상인 경우에만 구매
			// 가격순으로 정렬해 놓았으므로 가장 싼 부품을 산다
			if p.quality >= expected {
				totalCost += p.price
				bought = true
				break
			}
		}

		// 성능이 expected 이상인 부품이 없어서 해당 타입의 부품을 구매할 수 없는 경우
		if !bought {
			return false
		}
	}

	// 성능이 expected 이상인 컴퓨터를 조립하는데 필요한 비용이 B이하인 경우 -> 구매
	return totalCost <= B
}

func scanLine() string {
	scanner.Scan()
	return scanner.Text()
}

func scanSetupVars() (int, int) {
	fields := strings.Split(scanLine(), " ")
	return parseInt(fields[0]), parseInt(fields[1])
}

func scanPart() (string, string, int, int) {
	fields := strings.Split(scanLine(), " ")
	return fields[0], fields[1], parseInt(fields[2]), parseInt(fields[3])
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
