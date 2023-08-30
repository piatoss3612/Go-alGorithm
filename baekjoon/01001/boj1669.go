package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	X, Y    int
)

// 난이도: Gold 5
// 메모리: 904KB
// 시간: 4ms
// 분류: 수학
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	X, Y = scanInt(), scanInt()
}

func Solve() {
	// X == Y인 경우는 0일이 걸린다.
	if X == Y {
		fmt.Fprintln(writer, 0)
		return
	}

	diff := Y - X

	/*
		시작과 마지막은 무조건 1cm만큼 늘릴 수 있다.

		diff = 1 -> 1 (1일)
		diff = 2 -> 1 1 (2일)
		diff = 3 -> 1 1 1 (3일)
		diff = 4 -> 1 2 1 (3일)
		diff = 5 -> 1 2 1 1 (4일)
		diff = 6 -> 1 2 2 1 (4일)
		diff = 7 -> 1 2 2 1 1 (5일)
		diff = 8 -> 1 2 2 2 1 (5일)
		diff = 9 -> 1 2 3 2 1 (5일)
		diff = 10 -> 1 2 3 2 1 1 (6일)
		diff = 11 -> 1 2 3 2 2 1 (6일)
		diff = 12 -> 1 2 3 3 2 1 (6일)
		diff = 13 -> 1 2 3 3 2 1 1 (7일)
		diff = 14 -> 1 2 3 3 2 2 1 (7일)
		diff = 15 -> 1 2 3 3 3 2 1 (7일)
		diff = 16 -> 1 2 3 4 3 2 1 (7일)
		diff = 17 -> 1 2 3 4 3 2 1 1 (8일)
		diff = 18 -> 1 2 3 4 3 2 2 1 (8일)
		diff = 19 -> 1 2 3 4 3 3 2 1 (8일)
		diff = 20 -> 1 2 3 4 4 3 2 1 (8일)
		diff = 21 -> 1 2 3 4 4 3 2 1 1 (9일)

		1까지 늘렸다가 줄이는 경우: 1 (1)
		2까지 늘렸다가 줄이는 경우: 2~4 (2, 3, 3)
		3까지 늘렸다가 줄이는 경우: 5~9 (4, 4, 5, 5, 5)
		4까지 늘렸다가 줄이는 경우: 10~16 (6, 6, 6, 7, 7, 7, 7)
		5까지 늘렸다가 줄이는 경우: 17~25 (8, 8, 8, 8, 9, 9, 9, 9, 9)

		i^2 < diff <= (i+1)^2

		sqrt(diff)를 n이라고 했을 때, n을 반올림하여 제곱한 값이
		(i+1)^2보다 작으면 n * 2가 답이고,
		(i+1)^2보다 크거나 같으면 n * 2 - 1이 답이다.
	*/

	n := math.Round(math.Sqrt(float64(diff)))

	if n*n < float64(diff) {
		fmt.Fprintln(writer, int(n)*2)
	} else {
		fmt.Fprintln(writer, int(n)*2-1)
	}
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
