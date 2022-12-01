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
	square  map[int]int
)

const MOD = 1000000007

func init() {
	square = make(map[int]int)
	square[0] = 1
	square[1] = 2
}

// 난이도: Gold 5
// 메모리: 6120KB
// 시간: 24ms
// 분류: 조합론, 분할 정복을 이용한 거듭제곱
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		Input()
		Solve()
	}

}

func Input() {
	N = scanInt()
}

func Solve() {
	// 점화식
	// S1 = 1, S2 = 1
	// S(n+1) = Sn*2 (n >= 2)
	if N-2 < 0 {
		fmt.Fprintln(writer, 1)
		return
	}
	fmt.Fprintln(writer, DAQ(N-2))
}

// 분할 정복 + 거듭제곱
func DAQ(x int) int {
	_, ok := square[x]
	if ok {
		return square[x]
	}
	l := x / 2
	r := x - l

	_, ok = square[l]
	if !ok {
		square[l] = DAQ(l)
	}

	_, ok = square[r]
	if !ok {
		square[r] = DAQ(r)
	}

	square[x] = square[l] * square[r]
	square[x] %= MOD
	return square[x]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
