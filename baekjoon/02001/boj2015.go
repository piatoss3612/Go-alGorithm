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
	N, K    int
	psum    []int // A[1]~A[i]까지의 누적합
	memory  map[int]int
)

// 메모리: 12668KB
// 시간: 36ms
// 누적 합, 해시 맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
	psum = make([]int, N+1)
	memory = make(map[int]int)

	ans := 0

	for i := 1; i <= N; i++ {
		psum[i] = psum[i-1] + scanInt() // A[1]~A[i]까지의 누적합 구하기

		// 1~i까지의 누적합이 K와 일치하는 경우
		if psum[i] == K {
			ans++
		}

		// 누적합이 K가 되는지 판별하는 방법
		// psum[i] - psum[j] (j < i)가 K가 되는 경우
		// 즉, 누적합이 psum[j] (psum[i] - K)인 값이 해시 맵에 존재한다면
		// 그 값 등장한 경우의 수 만큼 ans에 누적해서 더해준다
		ans += memory[psum[i]-K]

		// psum[i]가 등장한 경우의 수 증가
		memory[psum[i]]++
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
