package bj18111

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m, b := scanInt(), scanInt(), scanInt()
	board := make([][]int, n)
	max, min := 0, 256
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		for j := 0; j < m; j++ {
			board[i][j] = scanInt()
			if board[i][j] > max { // 높이 최댓값 찾기
				max = board[i][j]
			}
			if board[i][j] < min { // 높이 최솟값 찾기
				min = board[i][j]
			}
		}
	}

	totalTime := 2147483647 // int타입이 가질 수 있는 최댓값
	height := 0

	for i := min; i <= max; i++ {
		cntRemove := 0
		cntStack := 0
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				if board[j][k] >= i {
					cntRemove += board[j][k] - i
				} else {
					cntStack += i - board[j][k]
				}
			}
		}
		// 제거하는 블록과 인벤토리에 있는 블록의 합이 쌓아야 하는 블록보다 크거나 같은 경우
		if cntRemove+b >= cntStack {
			time := cntRemove*2 + cntStack // 작업 시간 계산
			if time < totalTime {
				totalTime = time
				height = i
			}
			if time == totalTime { // 시간이 같은 경우, 높이가 높은 것으로 변경
				if height < i {
					height = i
				}
			}
		}
	}
	fmt.Fprintln(writer, totalTime, height)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
