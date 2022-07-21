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
	visited [1001][1001]bool // 이모티콘 개수 x 클립 보드에 있는 이모티콘 개수
)

// 메모리: 4752KB
// 시간: 8ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	q := [][]int{{1, 0, 0}}
	visited[1][0] = true

	// 깊이 우선 탐색
	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		emoji := front[0]
		sec := front[1]
		clip := front[2]

		if emoji == n {
			fmt.Fprintln(writer, sec)
			return
		}

		if emoji > 0 && emoji <= 1000 {
			if !visited[emoji][emoji] {
				visited[emoji][emoji] = true
				q = append(q, []int{emoji, sec + 1, emoji})
			}

			if !visited[emoji-1][clip] {
				visited[emoji-1][clip] = true
				q = append(q, []int{emoji - 1, sec + 1, clip})
			}
		}

		if clip > 0 && emoji+clip <= 1000 {
			if !visited[emoji+clip][clip] {
				visited[emoji+clip][clip] = true
				q = append(q, []int{emoji + clip, sec + 1, clip})
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
