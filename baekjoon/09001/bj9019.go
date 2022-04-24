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
	visited []bool
)

type DSLR struct {
	current  int
	register string
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	for i := 1; i <= n; i++ {
		a, b := scanInt(), scanInt()
		visited = make([]bool, 10000)
		solve(a, b)
	}
}

// 메모리: 35996KB
// 시간: 7432ms
// 방문 여부를 정확히 확인하지 않아서 시간 초과 오류 발생 2회
func solve(a, b int) {
	visited[a] = true
	queue := []DSLR{{a, ""}}

	for len(queue) > 0 {
		cur := queue[0].current
		reg := queue[0].register
		queue = queue[1:]

		if cur == b {
			fmt.Fprintln(writer, reg)
			return
		}

		d := (cur * 2) % 10000

		if !visited[d] {
			visited[d] = true
			queue = append(queue, DSLR{d, reg + "D"})
		}

		s := cur - 1
		if s < 0 {
			s = 9999
		}

		if !visited[s] {
			visited[s] = true
			queue = append(queue, DSLR{s, reg + "S"})
		}

		// 1, 10, 100을 왼쪽으로 돌리면 각각 10, 100, 1000이 되어야 한다
		l := (cur%1000)*10 + (cur / 1000)

		if !visited[l] {
			visited[l] = true
			queue = append(queue, DSLR{l, reg + "L"})
		}

		// 10, 100, 1000을 오른쪽으로 돌리면 각각 1, 10, 100이 되어야 한다
		r := (cur%10)*1000 + (cur / 10)

		if !visited[r] {
			visited[r] = true
			queue = append(queue, DSLR{r, reg + "R"})
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
