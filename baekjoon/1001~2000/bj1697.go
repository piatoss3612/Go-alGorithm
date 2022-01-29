package bj1697

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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, k := scanInt(), scanInt()
	visited = make([]bool, 100001)
	result := SBKG(n, k)
	fmt.Fprintln(writer, result)
}

// 숨바꼭질: 깊이 우선 탐색
func SBKG(start, target int) int {
	visited[start] = true
	result := 0
	queue := [][]int{{start, 0}}
	for len(queue) > 0 {
		idx := queue[0][0] // 현재 위치
		cnt := queue[0][1] // 동생을 찾는데 걸린 시간 = 깊이
		queue = queue[1:]
		if idx == target {
			result = cnt
			break
		}
		// x - 1, x + 1, x * 2 탐색
		if idx-1 >= 0 && !visited[idx-1] {
			visited[idx-1] = true
			queue = append(queue, []int{idx - 1, cnt + 1})
		}
		if idx+1 <= 100000 && !visited[idx+1] {
			visited[idx+1] = true
			queue = append(queue, []int{idx + 1, cnt + 1})
		}
		if idx*2 <= 100000 && !visited[idx*2] {
			visited[idx*2] = true
			queue = append(queue, []int{idx * 2, cnt + 1})
		}
	}
	return result
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
