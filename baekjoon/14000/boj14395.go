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
	visited map[int]bool
	s, t    int
	max     int = 1e9
)

type operation struct {
	cnt int
	ops string
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	s, t = scanInt(), scanInt()

	if s == t {
		fmt.Fprintln(writer, 0)
		return
	}

	visited = make(map[int]bool)

	ans := solve()
	fmt.Fprintln(writer, ans)
}

// 아스키 코드 순서 *, +, -, / 대로 연산
// 뺄셈은 어짜피 0이 되므로 연산이 의미가 없다
// 따라서 뺄셈은 전체 연산에서 제외한다
// 나눗셈은 반드시 1이 된다
func solve() string {
	queue := []operation{{s, ""}}

	for len(queue) > 0 {
		cnt := queue[0].cnt
		ops := queue[0].ops
		queue = queue[1:]

		if cnt == t {
			return ops
		}

		tmp := cnt * cnt
		if tmp <= t && !visited[tmp] {
			queue = append(queue, operation{tmp, ops + "*"})
			visited[tmp] = true
		}

		tmp = cnt + cnt
		if tmp <= t && !visited[tmp] {
			queue = append(queue, operation{tmp, ops + "+"})
			visited[tmp] = true
		}

		if !visited[1] {
			queue = append(queue, operation{cnt / cnt, ops + "/"})
			visited[1] = true
		}
	}
	return "-1"
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
