package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	str      [1000001]int
	redTeam  [1000001]int
	blueTeam [1000002]int
)

// 난이도: Gold 5
// 메모리: 28896KB
// 시간: 156ms
// 분류: 애드 혹, 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()

	// 기준선 i의 왼쪽 1~i번 참가자는 홍팀
	// 오른쪽 i+1~N번 참가자는 청팀
	for i := 1; i <= N; i++ {
		str[i] = scanInt()
		redTeam[i] = max(redTeam[i-1], str[i]) // 홍팀에서 가장 강한 사람 구하기
	}
	for i := N; i >= 1; i-- {
		blueTeam[i] = max(blueTeam[i+1], str[i]) // 청팀에서 가장 강한 사람 구하기
	}
}

func Solve() {
	red, blue := 0, 0
	// 모든 기준선에 대해 각 팀이 이기는 경우의 수 구하기
	for i := 1; i <= N-1; i++ {
		if redTeam[i] > blueTeam[i+1] {
			red++
		} else if redTeam[i] < blueTeam[i+1] {
			blue++
		}
	}

	if red > blue {
		// 홍팀이 이기는 경우가 많을 경우
		fmt.Fprintln(writer, "R")
	} else if red < blue {
		// 청팀이 이기는 경우가 많을 경우
		fmt.Fprintln(writer, "B")
	} else {
		// 동일한 경우
		fmt.Fprintln(writer, "X")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
