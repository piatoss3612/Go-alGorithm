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

	N, H    int
	section [500001]int
)

// 난이도: Gold 5
// 메모리: 6356KB
// 시간: 48ms
// 분류: 누적합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, H = scanInt(), scanInt()

	for i := 1; i <= N; i++ {
		x := scanInt()
		if i%2 == 0 {
			section[H-x+1] += 1 // 종유석은 H-X+1부터 H까지
		} else {
			// 석순은 1부터 X까지
			section[1] += 1
			section[x+1] -= 1 // X+1부터 H까지는 석순이 없으므로 -1
		}
	}

	// 누적합
	for i := 1; i <= H; i++ {
		section[i] += section[i-1]
	}
}

func Solve() {
	obst := section[1]

	// 최솟값 구하기
	for i := 2; i <= H; i++ {
		obst = min(obst, section[i])
	}

	// 최솟값의 개수 구하기
	cnt := 0
	for i := 1; i <= H; i++ {
		if section[i] == obst {
			cnt++
		}
	}

	fmt.Fprintln(writer, obst, cnt)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
