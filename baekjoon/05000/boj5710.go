package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	A, B    int
)

// 메모리: 896KB
// 시간: 4ms
// 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for {
		A, B = scanInt(), scanInt()
		if A == 0 && B == 0 {
			return
		}

		totalWatt := getWatt(A) // 전체 전기 사용량

		l := 0         // 전기를 안 쓴 경우
		r := totalWatt // 상근이가 모든 전기를 다 쓴 경우

		// 이분 탐색
		for l <= r {
			mid := (l + r) / 2

			sg := getCost(mid)             // 샹근이의 전기 사용량이 mid일 때 지불해야 하는 비용
			nb := getCost(totalWatt - mid) // 이웃이 지불해야 하는 비용

			// #답은 항상 유일한 경우만 주어진다
			// 상근이의 비용은 항상 이웃보다 적어야 한다
			// 이웃과 상근이의 비용의 차가 B와 동일하다면
			// 상근이의 비용을 출력하고 다음 테스트 케이스로 이동
			if nb-sg == B {
				fmt.Fprintln(writer, sg)
				break
			}

			if nb-sg > B {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
}

func getCost(watt int) int {
	cost := 0

	if watt > 1000000 {
		cost += 4979900 + (watt-1000000)*7
	} else if watt > 10000 {
		cost += 29900 + (watt-10000)*5
	} else if watt > 100 {
		cost += 200 + (watt-100)*3
	} else {
		cost += watt * 2
	}

	return cost
}

func getWatt(cost int) int {
	watt := 0

	if cost > 4979900 {
		watt += 1000000 + (cost-4979900)/7
	} else if cost > 29900 {
		watt += 10000 + (cost-29900)/5
	} else if cost > 200 {
		watt += 100 + (cost-200)/3
	} else {
		watt += cost / 2
	}

	return watt
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
