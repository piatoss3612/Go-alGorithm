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
	N, K    int
)

type Queue []int

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() int {
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// 1158번: 요세푸스 문제
// https://www.acmicpc.net/problem/1158
// 난이도: 실버 4
// 메모리: 89744 KB
// 시간: 324 ms
// 분류: 구현, 자료 구조, 큐
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
}

func Solve() {
	q := NewQueue()
	for i := 1; i <= N; i++ {
		q.Push(i)
	}

	fmt.Fprintf(writer, "<")
	for !q.IsEmpty() {
		for i := 0; i < K-1; i++ {
			q.Push(q.Pop())
		}
		fmt.Fprintf(writer, "%d", q.Pop())
		if !q.IsEmpty() {
			fmt.Fprintf(writer, ", ")
		}
	}
	fmt.Fprintf(writer, ">\n")
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
