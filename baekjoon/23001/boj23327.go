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

	N, Q int
	psum []int // psum[i]: 1~i번째 팀의 인기의 총합
	pjoy []int // pjoy[i]: 디비전을 1~i번째 팀으로 구성했을 때의 즐거움의 총합
)

// 난이도: Gold 4
// 메모리: 6352KB
// 시간: 152ms
// 분류: 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, Q = scanInt(), scanInt()
	psum = make([]int, N+1)
	pjoy = make([]int, N+1)

	// 1~i번째 팀의 인기의 총합을 구한다.
	for i := 1; i <= N; i++ {
		psum[i] = psum[i-1] + scanInt()
	}

	// 1~i번째 팀으로 디비전을 구성했을 때의 즐거움의 총합을 구한다.
	for i := 2; i <= N; i++ {
		pjoy[i] = pjoy[i-1] + psum[i-1]*(psum[i]-psum[i-1])
	}
}

func Solve() {
	for i := 1; i <= Q; i++ {
		l, r := scanInt(), scanInt()
		// 디비전을 l~r번째 팀으로 구성했을 때의 즐거움의 총합을 구한다.
		ans := pjoy[r] - pjoy[l-1] - (psum[r]-psum[l-1])*psum[l-1]
		fmt.Fprintln(writer, ans)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
