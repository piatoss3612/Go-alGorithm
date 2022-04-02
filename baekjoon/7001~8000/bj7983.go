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

type Homework struct {
	consume, deadLine int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	homeworks := make([]Homework, n)
	for i := 0; i < n; i++ {
		homeworks[i] = Homework{scanInt(), scanInt()}
	}

	/*
		예제 입력:
		3
		2 8
		1 13
		3 10

		입력값을 마감일을 기준으로 내림차순으로 정렬
	*/

	sort.Slice(homeworks, func(i, j int) bool {
		return homeworks[i].deadLine > homeworks[j].deadLine
	})

	/*
		예제 정렬:
		1 13
		3 10
		2 8

		예제 풀이:
		마감일 13일의 마지막날에 1시간 걸리는 과제를 함으로써 처음 얻을 수 있는 최대 연속으로 놀 수 있는 날은 12일까지
		마감일이 10일인 과제는 놀 수 있는 11, 12일을 제외하고 최대 연속으로 놀 수 있는 날을 10 - 3 = 7로 갱신
		마감일이 8일인 과제는 앞의 과제를 끝내고 최대 놀 수 있는 날이 7일까지이므로
		2일이 걸리는 과제를 마치고 연속으로 놀 수 있는 최댓값은 7 - 2 = 5로 갱신

		예제 출력:
		5
	*/

	left := homeworks[0].deadLine - homeworks[0].consume
	for i := 1; i < n; i++ {
		if homeworks[i].deadLine <= left {
			left = homeworks[i].deadLine - homeworks[i].consume
		} else {
			left -= homeworks[i].consume
		}
	}
	fmt.Fprintln(writer, left)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
