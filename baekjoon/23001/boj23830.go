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
	N          int
	scores     []int
	p, q, r, S int
)

// 난이도: Gold 4
// 메모리: 2316KB
// 시간: 32ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	scores = make([]int, N+1)
	for i := 1; i <= N; i++ {
		scores[i] = scanInt()
	}
	p, q, r, S = scanInt(), scanInt(), scanInt(), scanInt()
}

func Solve() {
	l, r := 1, 110000 // K가 양의 정수여야 하므로 l은 0이 아닌 1이어야 한다!!!!!
	for l <= r {
		m := (l + r) / 2
		// 태영이가 청소를 하지 않아도 되는 경우
		if IsOK(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	// 어떻게 해도 태영이가 청소를 피할 수 없는 경우
	if l > 110000 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, l) // lower bound 출력
	}
}

func IsOK(K int) bool {
	sum := 0
	for i := 1; i <= N; i++ {
		score := scores[i]
		if score > K+r {
			score -= p
		} else if score < K {
			score += q
		}
		sum += score
		if sum >= S {
			return true
		}
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
