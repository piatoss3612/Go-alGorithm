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

	N, T, G int
	checked [100000]bool
)

// 난이도: Gold 4
// 메모리: 5668KB
// 시간: 12ms
// 분류: 너비 우선 탐색
// 시간복잡도: O(N) -> 1<=N<=99999
// 공간복잡도: O(N)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, T, G = scanInt(), scanInt(), scanInt()
}

func Solve() {
	q := [][2]int{}
	q = append(q, [2]int{N, 0})
	checked[N] = true

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		num, try := front[0], front[1]

		// G를 만든 경우: 시도 횟수를 출력하고 종료
		if num == G {
			fmt.Fprintln(writer, try)
			return
		}

		// T번 시도한 경우: 더 이상 시도할 수 없으므로 넘어간다
		if try == T {
			continue
		}

		// 버튼 A를 누른 경우: 1을 더한 수를 큐에 넣는다
		A := num + 1
		if A <= 99999 && !checked[A] {
			checked[A] = true
			q = append(q, [2]int{A, try + 1})
		}

		// 버튼 B를 누른 경우
		// num에 2를 곱한 순간 수가 99,999를 넘어간다면, 높은 자릿수의 수를 1 낮췄을때 99,999를 넘지 않는다고 해도 탈출에 실패
		if num*2 > 99999 {
			continue
		}

		// num * 2에서 높은 자릿수를 1 낮춘 수를 큐에 넣는다
		B := calB(num)
		if B >= 0 && B <= 99999 && !checked[B] {
			checked[B] = true
			q = append(q, [2]int{B, try + 1})
		}
	}

	fmt.Fprintln(writer, "ANG")
}

func calB(n int) int {
	n = n * 2
	B := 0
	d := 1

	for n > 9 {

		r := n % 10
		n /= 10
		B += r * d
		d *= 10
	}

	B += (n - 1) * d
	return B
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
