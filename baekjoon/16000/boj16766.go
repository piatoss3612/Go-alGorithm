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
	N, M, C int
	cows    []int
)

// 난이도: Gold 4
// 메모리: 2748KB
// 시간: 44ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M, C = scanInt(), scanInt(), scanInt()
	cows = make([]int, N)
	for i := 0; i < N; i++ {
		cows[i] = scanInt()
	}
	sort.Ints(cows) // 공항에 도착하는 시간 순으로 오름차순 정렬
}

func Solve() {
	l, r := 0, cows[N-1]-cows[0] // 가장 오래 기다려야 하는 경우: 가장 나중에 공항에 도착한 소의 시간 - 가장 먼저 도착한 소의 시간
	for l <= r {
		mid := (l + r) / 2 // 매개 변수: 각 소가 버스를 타고 이동하기 위해 기다려야 하는 최대 대기 시간의 최솟값

		if check(mid) {
			// mid 시간이 조건을 만족하는 경우
			// 시간을 더 줄여본다
			r = mid - 1
		} else {
			// 시간을 늘린다
			l = mid + 1
		}
	}

	fmt.Fprintln(writer, l)
}

func check(maxInterval int) bool {
	busCnt := 0  // 버스의 개수
	cowCnt := 0  // 소들의 수
	l, r := 0, 0 // 두 포인터

	for busCnt < M {
		if l < N && r < N {
			if r+1 < N {
				// r+1번째 소를 버스에 태웠을 때,
				// 기다리는 시간의 최댓값이 maxInterval보다 작거나 같고
				// 버스의 수용 인원 한계를 넘지 않았을 때
				if cows[r+1]-cows[l] <= maxInterval && (r-l+1) <= C-1 {
					r++
				} else {
					// 버스 출발
					busCnt += 1
					cowCnt += (r - l + 1)
					l = r + 1
					r = l
				}
			} else {
				// 버스 출발
				busCnt += 1
				cowCnt += (r - l + 1)
				break
			}
		} else {
			break
		}
	}

	if cowCnt == N {
		return true
	}

	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
