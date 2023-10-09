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

	isPrime map[int]bool

	T       int
	N, A, B int
)

// 난이도: Gold 5
// 메모리: 14520KB
// 시간: 420ms
// 분류: 에라토스테네스의 체, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	T = scanInt()

	isPrime = make(map[int]bool)
	getPrimes()

	for i := 1; i <= T; i++ {
		Input()
		Solve()
	}
}

func getPrimes() {
	checked := make([]bool, 100001)

	for i := 2; i <= 100000; i++ {
		if !checked[i] {
			isPrime[i] = true
			for j := i * i; j <= 100000; j += i {
				checked[j] = true
			}
		}
	}
}

func Input() {
	N, A, B = scanInt(), scanInt(), scanInt()
}

func Solve() {
	q := [][2]int{}
	visited := [1000001]bool{}

	q = append(q, [2]int{N, 0})
	visited[N] = true

	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		creatures := front[0]
		snapped := front[1]

		if A <= creatures && creatures <= B && isPrime[creatures] {
			fmt.Fprintln(writer, snapped)
			return
		}

		var temp int

		temp = creatures / 2
		if !visited[temp] {
			visited[temp] = true
			q = append(q, [2]int{temp, snapped + 1})
		}

		temp = creatures / 3
		if !visited[temp] {
			visited[temp] = true
			q = append(q, [2]int{temp, snapped + 1})
		}

		temp = creatures + 1
		if temp <= 1000000 && !visited[temp] {
			visited[temp] = true
			q = append(q, [2]int{temp, snapped + 1})
		}

		temp = creatures - 1
		if temp > 0 && !visited[temp] {
			visited[temp] = true
			q = append(q, [2]int{temp, snapped + 1})
		}
	}

	fmt.Fprintln(writer, -1)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
