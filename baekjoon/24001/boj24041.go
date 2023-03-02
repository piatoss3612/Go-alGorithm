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
	N, G, K int
	ings    []Ingredient
)

type Ingredient struct {
	S, L, O int
}

// 난이도: Gold 4
// 메모리: 46968KB
// 시간: 536ms
// 분류: 이분 탐색, 매개 변수 탐색, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, G, K = scanInt(), scanInt(), scanInt()
	ings = make([]Ingredient, N)
	for i := 0; i < N; i++ {
		ings[i] = Ingredient{
			scanInt(), scanInt(), scanInt(),
		}
	}
}

func Solve() {
	l, r := 0, 2000000000
	for l <= r {
		m := (l + r) / 2
		// 밀키트를 구매하고 x일 후에도 밀키트를 먹을 수 있는 경우
		if CanEat(m) {
			l = m + 1
		} else {
			r = m - 1
		}

	}
	fmt.Fprintln(writer, r) // 최댓값(upper bound 출력)
}

func CanEat(expected int) bool {
	totalGerms := 0
	candidates := []int{}

	for _, ing := range ings {
		/*
			여기서 유통기한이 지나지 않은 재료의 세균 수를 계산하지 않아 5번 틀림
			if expected < ing.L {
				continue
			}
		*/
		germs := ing.S * max(1, expected-ing.L)
		totalGerms += germs

		// 중요하지 않은 재료는 제거될 후보로 등록
		if ing.O == 1 {
			candidates = append(candidates, germs)
		}
	}

	// 세균 수가 많은 순서로 오름차순 정렬
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})

	cnt := K
	// 최대 K개의 중요하지 않은 재료를 세균 수가 많은 것부터 제거
	for len(candidates) > 0 && cnt > 0 {
		x := candidates[0]
		candidates = candidates[1:]
		totalGerms -= x
		cnt--
	}

	// 세균 수가 G이하인 경우
	if totalGerms <= G {
		return true
	}
	return false
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
