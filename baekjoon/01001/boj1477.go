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
	N, M, L int
	inp     []int
)

// 메모리: 904KB
// 시간: 4ms
// 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M, L = scanInt(), scanInt(), scanInt()
	inp = make([]int, N+2)
	// inp[0]는 고속도로 시작 위치 0
	for i := 1; i <= N; i++ {
		inp[i] = scanInt()
	}
	// inp[N+1]은 고속도로가 끝나는 위치 L
	inp[N+1] = L

	sort.Ints(inp) // 현재 고속도로 상의 휴게소 위치를 오름차순으로 정렬

	// 이분 탐색
	// 현재 고속도로 상의 휴게소들 사이에 새로운 휴게소를 지은다고 했을 때
	// 새로운 휴게소들 간의 간격이 mid인 경우
	// 최소 M개의 휴게소를 지을 수 있는 lower bound를 찾는다

	// l을 0으로 잡으면 mid값이 0이 될 수 있으므로
	// divided by zero 런타임 에러가 발생하니 주의
	l, r := 1, L
	for l <= r {
		mid := (l + r) / 2
		count := 0

		for i := 1; i < len(inp); i++ {
			available := (inp[i] - inp[i-1]) // i번째 휴게소와 i-1번째 휴게소의 간격
			count += available / mid         // 간격/새로운 휴게소의 간격으로 지을 수 있는 휴게소의 개수를 구한다

			// 간격이 mid로 나누어 떨어지는 경우는 휴게소를 지을 수 없는 경우이므로
			// 더해준 값을 다시 빼준다
			if (available % mid) == 0 {
				count--
			}
		}

		// 새롭게 지을 수 있는 휴게소의 개수가 M보다 많은 경우
		if count > M {
			l = mid + 1
		} else {
			// 새롭게 지을 수 있는 휴게소의 개수가 M보다 작거나 같을 경우
			r = mid - 1
		}
	}

	fmt.Fprintln(writer, l)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
