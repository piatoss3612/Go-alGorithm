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

	N     int
	cases map[int]int
)

const MOD = 1000000007

// 난이도: Gold 5
// 메모리: 972KB
// 시간: 4ms
// 분류: 분할 정복을 이용한 거듭제곱
// 시간 복잡도: O(logN)
// 공간 복잡도: O(logN)
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	cases = make(map[int]int)
}

func Solve() {
	// 점화식 :
	// f(n) = f(n-1) * 2
	// f(0) = 1
	// f(1) = 2
	fmt.Fprintln(writer, dac(N-1))
}

// 분할 정복을 이용한 거듭제곱
func dac(x int) int {
	if x <= 0 {
		return 1
	}

	if x == 1 {
		return 2
	}

	_, ok := cases[x]
	if ok {
		return cases[x]
	}

	if x%2 == 0 {
		half := dac(x / 2)
		cases[x] = (half * half) % MOD
	} else {
		cases[x] = (dac(x-1) * 2) % MOD
	}

	return cases[x]
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
