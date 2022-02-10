package bj16953

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	a, b    int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	a, b = scanInt(), scanInt()

	// 방문 여부를 확인하기 위해 visited 슬라이스를 만들면 메모리 초과 발생
	// 전체 연산이 a -> 2a || a*10 + 1 이므로
	// 중복되는 연산 결과가 없을 것이라고 예상하고 방문 여부를 확인하지 않았다
	BFS(a)
}

func BFS(n int) {
	queue := [][]int{{n, 1}}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front[0] == b {
			fmt.Fprintln(writer, front[1])
			return
		}

		if front[0]*2 <= b {
			queue = append(queue, []int{front[0] * 2, front[1] + 1})
		}

		if front[0]*10+1 <= b {
			queue = append(queue, []int{front[0]*10 + 1, front[1] + 1})
		}
	}
	fmt.Fprintln(writer, -1)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
