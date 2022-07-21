package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}
	sort.Ints(input)

	s, e := 0, n-1
	var idx1, idx2 int
	min := math.MaxInt64

	/*
		예제 입력:
		5
		-2 4 -99 -1 98

		예제 입력을 정렬: -99 -2 -1 4 98
		시작 s = 0, 끝 e = n - 1부터 탐색을 시작
		합이 0보다 작은 경우는 s를 앞으로 이동
		합이 0보다 큰 경우는 e를 뒤로 이동
		하다 보면 합이 0에 가까운 인덱스 idx1과 idx2를 구할 수 있다

		예제 출력:
		-99 98
	*/

	for s < e {
		sum := input[s] + input[e]
		tmp := int(math.Abs(float64(sum)))

		if tmp < min {
			min = tmp
			idx1 = s
			idx2 = e
		}

		if sum < 0 {
			s += 1
		} else {
			e -= 1
		}
	}

	fmt.Fprintln(writer, input[idx1], input[idx2])
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
