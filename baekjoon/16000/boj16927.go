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
	N, M, R int
	arr     [][]int
)

// 16927번: 배열 돌리기 2
// hhttps://www.acmicpc.net/problem/16927
// 난이도: 골드 5
// 메모리: 5704 KB
// 시간: 48 ms
// 분류: 구현, 배열
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, R = scanInt(), scanInt(), scanInt()
	arr = make([][]int, N)
	for i := 0; i < N; i++ {
		arr[i] = make([]int, M)
		for j := 0; j < M; j++ {
			arr[i][j] = scanInt()
		}
	}
}

func Solve() {
	rotate()

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Fprintf(writer, "%d ", arr[i][j])
		}
		fmt.Fprintln(writer)
	}
}

func rotate() {
	layers := min(N, M) / 2

	for i := 0; i < layers; i++ {
		rotateLayer(i)
	}
}

func rotateLayer(layer int) {
	q := []int{}

	// 위쪽 행
	for i := layer; i < M-layer; i++ {
		q = append(q, arr[layer][i])
	}

	// 오른쪽 열
	for i := layer + 1; i < N-layer; i++ {
		q = append(q, arr[i][M-layer-1])
	}

	// 아래쪽 행
	for i := M - layer - 2; i >= layer; i-- {
		q = append(q, arr[N-layer-1][i])
	}

	// 왼쪽 열
	for i := N - layer - 2; i > layer; i-- {
		q = append(q, arr[i][layer])
	}

	// 회전
	r := R % len(q)             // 실제 회전 횟수 (제자리로 돌아오는 경우를 루프에서 제외)
	q = append(q[r:], q[:r]...) // 반시계 방향으로 회전

	idx := 0

	// 위쪽 행
	for i := layer; i < M-layer; i++ {
		arr[layer][i] = q[idx]
		idx++
	}

	// 오른쪽 열
	for i := layer + 1; i < N-layer; i++ {
		arr[i][M-layer-1] = q[idx]
		idx++
	}

	// 아래쪽 행
	for i := M - layer - 2; i >= layer; i-- {
		arr[N-layer-1][i] = q[idx]
		idx++
	}

	// 왼쪽 열
	for i := N - layer - 2; i > layer; i-- {
		arr[i][layer] = q[idx]
		idx++
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
