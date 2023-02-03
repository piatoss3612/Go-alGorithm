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
	T       int
	f       [1000001]int // f[x]: 자연수 x의 모든 약수의 합
	g       [1000001]int // g[x]: x보다 작거나 같은 모든 자연수 y의 f[x]값을 더한 값
)

func init() {
	// i의 배수는 모두 i를 약수로 가지고 있다는 사실에 근거한 풀이
	for i := 1; i <= 1000000; i++ {
		for j := i; j <= 1000000; j += i {
			f[j] += i
		}
		g[i] = g[i-1] + f[i] // f[i]를 구한 후에 g[i] 갱신
	}
}

// 난이도: Gold 4
// 메모리: 18212KB
// 시간: 132ms
// 분류: 누적 합, 에라토스테네스의 체
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	T = scanInt()
}

func Solve() {
	for i := 1; i <= T; i++ {
		fmt.Fprintln(writer, g[scanInt()])
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
