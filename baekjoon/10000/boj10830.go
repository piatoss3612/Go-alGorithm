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
	n       int
	findMat map[int][][]int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	b := scanInt()
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		// *중요: 결과를 1000으로 나누어주기 때문에 입력값이 1000이면 결국 0이 되므로
		// 미리 %연산을 해준다
		for j := 0; j < n; j++ {
			mat[i][j] = scanInt() % 1000
		}
	}
	findMat = make(map[int][][]int)
	findMat[1] = mat

	ans := solve(b)

	for _, i := range ans {
		for _, v := range i {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
	}
}

func solve(k int) [][]int {
	if k == 1 {
		return findMat[1]
	}

	// 분할 정복
	a := int(k / 2)
	b := k - a

	// 행렬의 a제곱에 해당하는 값
	_, ok1 := findMat[a]
	if !ok1 {
		// 없다면 a제곱에 해당하는 값을 찾아준다
		findMat[a] = solve(a)
	}

	// 행렬의 b제곱에 해당하는 값
	_, ok2 := findMat[b]
	if !ok2 {
		// 없다면 b제곱에 해당하는 값을 찾아준다
		findMat[b] = solve(b)
	}

	// 행렬의 k제곱의 해당하는 값 = 행렬의 a제곱 결과 * 행렬의 b제곱 결과
	findMat[k] = multiplyMatrix(findMat[a], findMat[b])
	return findMat[k]
}

// 행렬의 곱셈
func multiplyMatrix(m1, m2 [][]int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
			res[i][j] %= 1000
		}
	}
	return res
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
