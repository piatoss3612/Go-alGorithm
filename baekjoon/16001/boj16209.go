package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

// 메모리: 17176KB
// 시간: 284ms
// 그리디 알고리즘: 수열의 인접한 원소의 곱들의 합을 최대화하는 방법
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N := scanInt()
	// 이 문제의 입력값은 절댓값이 100만 이하
	// 즉 음수를 포함하고 있다
	// 따라서 입력값을 0이상의 양수와 음수로 구분하여 저장한다
	positive := make([]int, 0, 500000)
	negative := make([]int, 0, 500000)
	for i := 0; i < N; i++ {
		x := scanInt()
		if x >= 0 {
			positive = append(positive, x)
		} else {
			negative = append(negative, x)
		}
	}

	// 양수 수열 오름차순 정렬
	sort.Ints(positive)

	n := len(positive)
	positiveSort := make([]int, n)

	// 양수 수열의 인접한 원소의 곱들의 합을 최대화하는 방법
	// 가장 작은 값을 왼쪽, 그 다음 작은 값은 오른쪽으로 정렬하는 과정을 반복
	// 예를 들어, 1 2 3 4 5 6을 입력받으면 정렬된 수열은 1 3 5 6 4 2가 된다
	// 반복 과정에서 왼쪽 오른쪽을 바꾸면 2 4 6 4 3 1도 가능한 결과다
	for i := 0; i < n/2; i++ {
		positiveSort[i] = positive[0]
		positive = positive[1:]
		positiveSort[n-i-1] = positive[0]
		positive = positive[1:]
	}

	// 양수의 개수가 홀수 개인 경우
	if n%2 != 0 {
		positiveSort[n/2] = positive[0]
	}

	// 음수 수열 내림차순 정렬
	sort.Slice(negative, func(i, j int) bool {
		return negative[i] > negative[j]
	})

	m := len(negative)
	negativeSort := make([]int, m)

	// 음수 수열의 인접한 원소의 곱들의 합을 최대화하는 방법
	// 가장 큰 값을 왼쪽, 그 다음 큰 값은 오른쪽으로 정렬하는 과정을 반복
	// 예를 들어, -1 -2 -3 -4를 입력받으면 정렬된 수열은 -1 -3 -4 -2가 된다
	// 그런데 정렬된 음수 수열과 양수 수열을 이어붙여야 하므로
	// 두 수열의 이어지는 부분의 인접한 원소의 곱이 최소가 되도록 하려면
	// 가장 큰 값이 오른쪽으로 가도록 정렬해야 한다
	// -1 -2 -3 1 2 4 7을 정렬한다고 하면
	// 양수 수열: 1 4 7 2
	// 음수 수열: -2 -3 -1
	// 전체 수열: -2 -3 -1 1 4 7 2
	for i := 0; i < m/2; i++ {
		negativeSort[m-i-1] = negative[0]
		negative = negative[1:]
		negativeSort[i] = negative[0]
		negative = negative[1:]
	}

	// 음수의 개수가 홀수 개인 경우
	if m%2 != 0 {
		negativeSort[m/2] = negative[0]
	}

	for i := 0; i < m; i++ {
		fmt.Fprintf(writer, "%d ", negativeSort[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", positiveSort[i])
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
