package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	N, P, E    int
	accept     [][2]int
	candidates []int
)

// 18233번: 러버덕을 사랑하는 모임
// hhttps://www.acmicpc.net/problem/18233
// 난이도: 골드 5
// 메모리: 868 KB
// 시간: 12 ms
// 분류: 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, P, E = scanInt(), scanInt(), scanInt()
	accept = make([][2]int, N)
	for i := 0; i < N; i++ {
		accept[i][0], accept[i][1] = scanInt(), scanInt()
	}
	candidates = make([]int, 0, N)
}

func Solve() {
	// P가 N보다 크면 P명에게 선물을 나눠주는 것은 불가능
	if P > N {
		fmt.Fprintln(writer, -1)
		return
	}

	// 완전 탐색
	for i := 0; i < N; i++ {
		if bruteForce(i, 0, 0, 0) {
			break
		}
	}

	// P명에게 E개의 선물을 줄 수 없는 경우
	if len(candidates) != P {
		fmt.Fprintln(writer, -1)
		return
	}

	// P명에게 E개의 선물 분배
	distribution := make([]int, N)

	// 최소한의 선물을 받도록 우선 분배
	for i := 0; i < len(candidates); i++ {
		candidate := candidates[i]
		amount := accept[candidate][0]
		distribution[candidate] += amount
		E -= amount
	}

	// 남은 선물은 최대치를 채우도록 임의로 분배
	if E > 0 {
		for i := 0; i < len(candidates); i++ {
			if E == 0 {
				break
			}

			candidate := candidates[i]
			amount := accept[candidate][1] - accept[candidate][0]
			if amount > E {
				amount = E
			}
			distribution[candidate] += amount
			E -= amount
		}
	}

	// 출력
	for i := 0; i < N; i++ {
		fmt.Fprintf(writer, "%d ", distribution[i])
	}
	fmt.Fprintln(writer)
}

func bruteForce(idx, cnt, minAmount, maxAmount int) bool {
	candidates = append(candidates, idx)
	cnt += 1
	minAmount += accept[idx][0]
	maxAmount += accept[idx][1]

	rollback := func() {
		minAmount -= accept[idx][0]
		maxAmount -= accept[idx][1]
		cnt -= 1
		candidates = candidates[:len(candidates)-1]
	}

	if cnt == P {
		if minAmount <= E && E <= maxAmount {
			return true
		}

		rollback()

		return false
	}

	for i := idx + 1; i < N; i++ {
		if bruteForce(i, cnt, minAmount, maxAmount) {
			return true
		}
	}

	rollback()

	return false
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
