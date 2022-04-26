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
	input   []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	input = make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = scanInt()
	}

	rec(0)
}

/*
모든 값이 0으로 채우져 있는 길이가 n인 배열 a를 배열 b로 변환하는 문제
이는 반대로 생각해보면 배열 b의 모든 값을 0으로 만들면 된다는 것이다

연산의 순서는 배열의 모든 값을 2로 나누어 주는 것을 제외하면 아무런 상관이 없는 문제
*/

func rec(x int) {
	same := true
	for i := 0; i < n; i++ {
		if input[i] != 0 {
			same = false
			break
		}
	}

	if same {
		fmt.Fprintln(writer, x)
		return
	}

	flag := true
	// 배열 b의 특정값이 2로 나누어 떨어지지 않는 경우: 해당 값을 짝수로 만들어 주고 재귀 호출
	for i := 0; i < n; i++ {
		if input[i]%2 != 0 {
			flag = false
			input[i] -= 1
			rec(x + 1)
			break
		}
	}

	// 배열 b의 모든 값이 2로 나누어 떨어지는 경우: 모든 값을 2로 나누어주고 재귀 호출
	if flag {
		for i := 0; i < n; i++ {
			input[i] /= 2
		}
		rec(x + 1)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
