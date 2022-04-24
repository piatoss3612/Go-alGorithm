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
	n, c    int
	houses  []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, c = scanInt(), scanInt()
	houses = make([]int, n)
	for i := 0; i < n; i++ {
		houses[i] = scanInt()
	}
	sort.Ints(houses)

	left := 1           // 거리 최솟값
	right := 1000000000 // 거리 최댓값

	// 최대 거리를 구하기 위한 이분 탐색
	for left < right {
		mid := (left + right) / 2
		// 공유기 사이의 거리(mid)를 충족하는 공유기의 최대 갯수가 c보다 작은 경우
		if getMaxRouter(mid) < c {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Fprintln(writer, left-1)
}

func getMaxRouter(dist int) int {
	cnt := 1
	prev := houses[0]

	for i := 1; i < n; i++ {
		if houses[i]-prev >= dist {
			cnt += 1
			prev = houses[i]
		}
	}
	return cnt
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
