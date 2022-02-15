package bj12851

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner        = bufio.NewScanner(os.Stdin)
	writer         = bufio.NewWriter(os.Stdout)
	minTime, count int
	visited        [100001]bool
	n, k           int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, k = scanInt(), scanInt()

	BFS()
	fmt.Fprintf(writer, "%d\n%d\n", minTime, count)
}

func BFS() {
	visited[n] = true
	queue := [][]int{{n, 0}}

	for len(queue) > 0 {
		idx := queue[0][0]
		time := queue[0][1]

		visited[idx] = true
		queue = queue[1:]

		if (count != 0) && (minTime == time) && (idx == k) {
			count += 1
		}

		if (count == 0) && (idx == k) {
			minTime = time
			count = 1
		}

		// 숨바꼭질 1과 달리, 큐에 데이터를 추가하는 단계에서
		// n이 1인 경우: 1+1, 1-1, 1*2 세 개의 경우의 수가 있는데
		// 1+1을 검사하고 visited[1+1] = true를 해버리면
		// 1*2를 체크하지 않고 넘어가게 되므로 모든 경우의 수를 탐색할 수 없다
		// 따라서 큐에서 팝 연산을 하는 앞단계에서 visited를 체크해 준다

		if valid(idx-1) && !visited[idx-1] {
			queue = append(queue, []int{idx - 1, time + 1})
		}

		if valid(idx+1) && !visited[idx+1] {
			queue = append(queue, []int{idx + 1, time + 1})
		}

		if valid(idx*2) && !visited[idx*2] {
			queue = append(queue, []int{idx * 2, time + 1})
		}
	}
}

func valid(v int) bool {
	return v >= 0 && v <= 100000
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
