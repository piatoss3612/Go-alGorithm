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
	N, M, P int
	matrix  map[int]*[][]int
)

// 난이도: Gold 4
// 메모리: 5536KB
// 시간: 168ms
// 분류: 분할 정복을 이용한 거듭제곱
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
}

func Input() {
	for {
		N, M, P = scanInt(), scanInt(), scanInt()
		if N == 0 && M == 0 && P == 0 {
			return
		}

		matrix = make(map[int]*[][]int)

		inp := make([][]int, N)
		for i := 0; i < N; i++ {
			inp[i] = make([]int, N)
			for j := 0; j < N; j++ {
				inp[i][j] = scanInt()
			}
		}
		matrix[1] = &inp
		Solve()
	}
}

func Solve() {
	rec(P)
	printMatrix(P)
}

func rec(n int) *[][]int {
	_, ok := matrix[n]
	if ok {
		return matrix[n]
	}

	// n이 짝수라면 l과 r 모두 짝수
	// n이 홀수라면 l은 짝수 r은 홀수
	l := n / 2
	r := n - l

	_, ok = matrix[l]
	if !ok {
		matrix[l] = rec(l)
	}

	_, ok = matrix[r]
	if !ok {
		matrix[r] = rec(r)
	}

	matrix[n] = mulMatrix(l, r)
	return matrix[n]
}

// N by N 행렬의 곱셈
func mulMatrix(l, r int) *[][]int {
	a := matrix[l]
	b := matrix[r]
	c := createMatrix()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				(*c)[i][j] += (*a)[i][k] * (*b)[k][j]
			}
			(*c)[i][j] %= M
		}
	}
	return c
}

// N by N 행렬 생성
func createMatrix() *[][]int {
	m := make([][]int, N)
	for i := 0; i < N; i++ {
		m[i] = make([]int, N)
	}
	return &m
}

// N by N 행렬 출력
func printMatrix(n int) {
	matrixToPrint := *matrix[n]
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fprintf(writer, "%d ", matrixToPrint[i][j])
		}
		fmt.Fprintln(writer)
	}
	fmt.Fprintln(writer) // 출력 조건에 따라 행렬을 출력하고 마지막 비어있는 라인을 추가해줘야 한다
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
