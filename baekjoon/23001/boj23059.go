package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	N         int
	A, B      string
	nextItems map[string][]string // nextItems[A]: A의 다음 단계 아이템들
	inDegree  map[string]int      // inDegree[A]: A를 구매하기 위해 조건에 맞게 먼저 구매해야 하는 아이템의 개수 (진입 차수)
	total     int                 // 구매해야 하는 아이템의 전체 개수
)

// 난이도: Gold 2
// 메모리: 148216KB
// 시간: 1020ms
// 분류: 위상 정렬, 해시를 사용한 집합과 맵
// 회고:
// 반복문 안에서 candidates 후보 목록을 사용하지 않고 inDegree 맵을 직접 range 키워드로 순회함으로 인해 시간 초과 발생
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	nextItems = make(map[string][]string)
	inDegree = make(map[string]int)

	for i := 1; i <= N; i++ {
		A, B = scanItem(), scanItem()  // A -> B: 아이템 A는 B를 구매하기 위한 조건
		inDegree[B]++                  // 아이템 B의 진입 차수 증가
		if _, ok := inDegree[A]; !ok { // 아이템 A의 진입 차수가 정해지지 않은 경우 0으로 초기화
			inDegree[A] = 0
		}
		nextItems[A] = append(nextItems[A], B) // A의 다음 단계 아이템으로 B를 추가
	}

	total = len(inDegree) // 구매해야 하는 아이템의 전체 개수는 inDegree의 길이와 같다

	TopologicalSort() // 위상 정렬
}

/*
# 문제 조건

1. 현재 구매할 수 있는 아이템 중 아직 구매하지 않은 아이템을 모두 찾는다.
2. 찾은 아이템을 사전 순으로 모두 구매한다.
*/

func TopologicalSort() {
	ans := []string{}        // 문제 조건에 맞는 구매 결과
	toPurchase := []string{} // 구매할 아이템들

	// 현재 구매할 수 있는 아이템들 중 진입 차수가 0인, 즉 선행 구매해야 하는 아이템이 없는 아이템들을 먼저 구매
	for k, v := range inDegree {
		if v == 0 {
			toPurchase = append(toPurchase, k)
		}
	}

	for {
		// 구매할 수 있는 아이템이 없는 경우
		if len(toPurchase) == 0 {
			break
		}

		sort.Strings(toPurchase) // 구매할 아이템 이름을 사전순으로 정렬

		candidates := []string{} // 현재 구매할 수 있는 아이템을 모두 구매하고 난 다음 차례로 구매할 후보 리스트

		for len(toPurchase) > 0 {
			x := toPurchase[0]
			toPurchase = toPurchase[1:]
			ans = append(ans, x)

			// 특정 아이템을 구매함으로써 다음 단계의 아이템을 구매할 수 있는 조건이 일부 충족
			for _, next := range nextItems[x] {
				inDegree[next]--
				// next 아이템의 진입 차수가 0이 된 경우, 즉 구매할 수 있는 조건이 완전히 충족된 경우
				if inDegree[next] == 0 {
					// 다음에 구매할 후보 리스트에 추가
					candidates = append(candidates, next)
				}
			}
		}

		toPurchase = append(toPurchase, candidates...) // 구매할 아이템 목록에 후보 리스트를 추가
	}

	if len(ans) != total {
		// 구매한 아이템의 개수가 전체 아이템의 개수와 다른 경우
		fmt.Fprintln(writer, -1)
	} else {
		// 구매한 아이템을 한 줄에 하나씩 출력
		for _, item := range ans {
			fmt.Fprintln(writer, item)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanItem() string {
	scanner.Scan()
	return scanner.Text()
}
