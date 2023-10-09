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
	N, M    int
	paper   [101]bool
)

// 난이도: Silver 3
// 메모리: 1012KB
// 시간: 4ms
// 분류: 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		paper[scanInt()] = true
	}
}

func Solve() {
	printer := []int{}
	ans := 0

	for i := 1; i <= N; i++ {
		// 논문의 i번 페이지가 비어있는 경우
		if !paper[i] {

			// 프린터의 대기열이 비어있는 경우
			if len(printer) == 0 {
				printer = append(printer, i)
				continue
			}

			n := len(printer)
			back := printer[n-1]

			// 프린터의 대기열 마지막에 i-1번 페이지가 들어있는 경우
			if back == i-1 {
				printer = append(printer, i)
				continue
			}

			split := 5 + 5 + 2 + 2   // 프린터의 대기열 마지막에 있는 페이지와 i번째 페이지를 따로 출력하는 경우 필요한 잉크의 양
			join := 5 + 2*(i-back+1) // 프린터의 대기열 마지막에 있는 페이지부터 i번째 페이지까지 연속으로 출력하는 경우 필요한 잉크의 양

			// 연속해서 출력하는 비용이 더 작거나 같은 경우
			if join <= split {
				// 프린터 대기열에 back+1~i번째 페이지 추가
				for j := back + 1; j <= i; j++ {
					printer = append(printer, j)
				}
			} else {
				ans += 5 + 2*n     // 프린터를 작동하여 대기열에 있는 페이지들을 출력
				printer = []int{i} // 프린터 대기열에 i번 페이지만 추가
			}
		}
	}

	// 프린터 대기열이 남아있는 경우
	if len(printer) > 0 {
		ans += 5 + 2*len(printer)
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
