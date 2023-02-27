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
	N, Q    int
	trees   []int // 나무들의 위치
	subSum  []int // 나무들의 위치값의 누적합
)

// 난이도: Gold 4
// 메모리: 5440KB
// 시간: 104ms
// 분류: 이분 탐색, 누적합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, Q = scanInt(), scanInt()
	trees = make([]int, N+1)
	subSum = make([]int, N+1)
	// 나무들의 위치 입력
	for i := 1; i <= N; i++ {
		trees[i] = scanInt()
	}
	// #필수# 나무들의 위치 오름차순 정렬
	sort.Ints(trees)
	// 나무들의 위치값의 누적합
	for i := 1; i <= N; i++ {
		subSum[i] = subSum[i-1] + trees[i]
	}
}

func Solve() {
	for i := 1; i <= Q; i++ {
		fmt.Fprintln(writer, BinarySearch(scanInt()))
	}
}

func BinarySearch(x int) int {
	// 이분 탐색을 통해 사진을 찍는 위치 x보다 나무의 위치가 작거나 같은 인덱스의 최댓값(upper bound)을 찾는다
	l, r := 1, N
	for l <= r {
		m := (l + r) / 2
		if trees[m] >= x {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	// 누적합을 활용해 구간을 나눠 절댓값의 합 구하기
	score := 0
	score += r*x - subSum[r]
	score += subSum[N] - subSum[r] - (N-r)*x
	return score
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
