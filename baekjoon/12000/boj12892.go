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
	N, D    int
	inp     []Present
)

type Present struct {
	P, V int
}

// 메모리: 5572KB -> 5516KB
// 시간: 88ms -> 72ms
// 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, D = scanInt(), scanInt()
	inp = make([]Present, N)
	for i := 0; i < N; i++ {
		inp[i] = Present{scanInt(), scanInt()}
	}

	// 가격 오름차순으로 정렬
	sort.Slice(inp, func(i, j int) bool {
		// 없어도 되는데 왜 시간하고 메모리가 늘어날까?
		// if inp[i].P == inp[j].P {
		// 	return inp[i].V < inp[j].V
		// }
		return inp[i].P < inp[j].P
	})

	l, r := 0, 0
	satisfaction := 0
	res := 0

	// 두 포인터 탐색
	for l <= r && r < len(inp) {
		// 가격이 가장 싼 선물과 가격이 가장 비싼 선물의 가격의 차가 D보다 작다면
		if inp[r].P-inp[l].P < D {
			satisfaction += inp[r].V
			res = max(res, satisfaction)
			r++
		} else {
			satisfaction -= inp[l].V
			// res = max(res, satisfaction) 불필요한 라인
			l++
		}
	}

	fmt.Fprintln(writer, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
