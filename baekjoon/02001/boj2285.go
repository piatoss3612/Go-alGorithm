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
	N       int
	inp     []Village
)

type Village struct {
	pos    int
	people int
}

// 메모리: 5604KB
// 시간: 76ms
// 그리디 알고리즘: 우체국은 각 마을의 각 사람까지 거리의 합이 최소가 되는 위치에 세워져야 한다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	inp = make([]Village, N)

	total := 0

	for i := 0; i < N; i++ {
		inp[i] = Village{scanInt(), scanInt()}
		total += inp[i].people
	}

	// 마을 위치 오름차순으로 정렬
	sort.Slice(inp, func(i, j int) bool {
		return inp[i].pos < inp[j].pos
	})

	// 최적의 값을 구하기 위해서는
	// 우체국의 위치에서 왼쪽, 오른쪽으로 분포된 마을 사람들의 수의 차가 최소가 되어야 한다
	// 따라서 앞에서부터 마을 사람들의 수를 누적하여 더하고
	// 그 값이 전체 마을 사람의 수의 절반을 넘어가는 위치에 우체국을 세운다
	count := 0
	for i := 0; i < N; i++ {
		count += inp[i].people
		if count > total/2 {
			fmt.Fprintln(writer, inp[i].pos)
			return
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
