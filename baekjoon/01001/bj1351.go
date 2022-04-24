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
	seq     map[int]int
	p, q    int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	seq = make(map[int]int) // 맵 초기화
	p, q = scanInt(), scanInt()
	seq[0] = 1
	fmt.Fprintln(writer, solve(n))
}

// n의 범위가 10의 12승까지이므로 슬라이스로 다이나믹 프로그래밍 연산 불가
// 맵을 사용한 재귀 함수를 통해 무한수열의 n번째 값을 찾을 수 있다
func solve(n int) int {
	if seq[n] != 0 {
		return seq[n]
	}
	seq[n] = solve(n/p) + solve(n/q)
	return seq[n]
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
