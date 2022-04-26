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

type work struct {
	t  int
	dl int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	input := make([]work, n)
	for i := 0; i < n; i++ {
		input[i] = work{scanInt(), scanInt()}
	}

	// 일을 가장 늦게 시작할 수 있는 방법은 가능한 모든 일을 마감 시간에 맞춰서 끝내는 것이다

	// 마감 시간이 가장 큰 작업부터 내림차 순으로 정렬
	sort.Slice(input, func(i, j int) bool {
		return input[i].dl > input[j].dl
	})

	// 남은 시간
	ans := input[0].dl - input[0].t

	for i := 1; i < n; i++ {
		// 마감 시간이 남은 시간보다 작거나 같은 경우
		if input[i].dl <= ans {
			ans = input[i].dl - input[i].t
			// 마감 시간이 남은 시간보다 큰 경우
		} else {
			ans -= input[i].t
		}
	}

	// 시간 안에 일을 끝내지 못한 경우
	if ans < 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
