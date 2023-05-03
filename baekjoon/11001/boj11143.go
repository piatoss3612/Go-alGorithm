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
	B, P, Q int
	segment [300000]int
)

// 난이도: Gold 1
// 메모리: 5812KB
// 시간: 120ms
// 분류: 세그먼트 트리
// 회고: 구슬을 박스에 'put'한다는게 말 그대로 놓는다는 의미로 구슬을 누적해서 더해줘야 한다. 왜 'change'한다는 의미로 이해했는지 모르겠다.
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	T = scanInt()

	for i := 1; i <= T; i++ {
		Setup()
		Solve()
	}
}

func Setup() {
	B, P, Q = scanInt(), scanInt(), scanInt()
	segment = [300000]int{}
}

func Solve() {
	for i := 1; i <= P+Q; i++ {
		op, x, y := scanByte(), scanInt(), scanInt()

		switch op {
		case 'P':
			Update(x, y, 1, B, 1)
		case 'Q':
			fmt.Fprintln(writer, Query(x, y, 1, B, 1))
		}
	}
}

func Update(target, val, left, right, node int) {
	if right < target || target < left {
		return
	}

	if left == right {
		segment[node] += val
		return
	}

	segment[node] += val
	mid := (left + right) / 2
	Update(target, val, left, mid, node*2)
	Update(target, val, mid+1, right, node*2+1)
}

func Query(left, right, qLeft, qRight, node int) int {
	if right < qLeft || qRight < left {
		return 0
	}

	if left <= qLeft && qRight <= right {
		return segment[node]
	}

	mid := (qLeft + qRight) / 2
	return Query(left, right, qLeft, mid, node*2) + Query(left, right, mid+1, qRight, node*2+1)
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}

func scanByte() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}
