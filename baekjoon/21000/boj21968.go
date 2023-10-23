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
	T, N    int
)

// 21968번: 선린의 터를
// https://www.acmicpc.net/problem/21968
// 난이도: 실버 3
// 메모리: 868 KB
// 시간: 4 ms
// 분류: 수학, 정수론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	T = scanInt()
	for i := 0; i < T; i++ {
		Solve()
	}
}

func Solve() {
	N = scanInt()

	ans := 0

	/*

		1(= 2^0)번째: 1 (= 3^0)
		2(= 2^1)번째: 3 (= 3^1)
		3(= 2^0 + 2^1)번째: 4 = 1 + 3 (= 3^0 + 3^1)
		4(= 2^2)번째: 9 (= 3^2)
		5(= 2^0 + 2^2)번째: 10 = 1 + 9 (= 3^0 + 3^2)
		6(= 2^1 + 2^2)번째: 12 = 3 + 9 (= 3^1 + 3^2)
		7(= 2^0 + 2^1 + 2^2)번째: 13 = 1 + 3 + 9 (= 3^0 + 3^1 + 3^2)

		2^0 + 2^1 + 2^2 + ... + 2^i = N 이라고 할 때,
		N번째 선린의 터를 구하기 위해서는
		3^0 + 3^1 + 3^2 + ... + 3^i 의 합을 구하면 된다.
	*/

	for i := 0; ; i++ {
		x := 1 << i
		// N보다 큰 2의 제곱수를 찾으면 종료한다.
		if x > N {
			break
		}

		// N의 i번째 비트가 1이면 3^i를 더한다.
		if N&x != 0 {
			ans += pow(3, i)
		}
	}

	fmt.Fprintln(writer, ans)
}

func pow(n, e int) int {
	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		return pow(n*n, e/2)
	}

	return pow(n*n, e/2) * n
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
