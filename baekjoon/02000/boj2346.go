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

	N        int
	balloons []Balloon
)

type Balloon struct {
	idx int
	val int
}

// 난이도: Silver 3
// 메모리: 144884KB
// 시간: 452ms
// 분류: 자료 구조, 덱
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	balloons = make([]Balloon, N)

	for i := 1; i <= N; i++ {
		balloons[i-1] = Balloon{idx: i, val: scanInt()}
	}
}

func Solve() {
	ans := make([]int, 0, N)

	for len(balloons) > 0 {
		front := balloons[0]
		balloons = balloons[1:]
		ans = append(ans, front.idx)

		if len(balloons) == 0 {
			break
		}

		if front.val > 0 {
			for i := 1; i <= (front.val-1)%len(balloons); i++ {
				balloons = append(balloons, balloons[0])
				balloons = balloons[1:]
			}
		} else {
			for i := 1; i <= (-front.val)%len(balloons); i++ {
				balloons = append([]Balloon{balloons[len(balloons)-1]}, balloons[:len(balloons)-1]...)
			}
		}
	}

	for _, v := range ans {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
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
