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
)

// 메모리: 16020KB
// 시간: 220ms

/*
길이가 최대 1백만이므로 다이나믹 프로그래밍 O(n^2)으로는 풀 수 없는 문제
이분 탐색을 사용해서 (nlogn) 시간에 풀어야 한다
*/
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	// append를 사용하면 슬라이스의 메모리가 2배가 되어 낭비되는 메모리가 생긴다
	// 슬라이스를 수열 A의 크기의 최댓값인 1백만 + 1로 초기화하고
	// 세그먼트의 현재 길이를 따로 관리함으로써 메모리 낭비를 줄일 수 있다
	seg := make([]int, 1000001)
	seg[0] = 0  // 첫번째 입력값과 비교를 위해 입력값의 최솟값보다 작은 값으로 초기화
	segLen := 1 // 현재 세그먼트의 길이

	ans := 0 // LIS

	var t int       // 입력받는 값
	var l, r, m int // 이분 탐색 왼쪽, 오른쪽, 가운데
	for i := 0; i < n; i++ {
		t = scanInt()

		// 1. 입력받은 값이 세그먼트의 마지막 값보다 큰 경우
		if t > seg[segLen-1] {
			seg[segLen] = t // 세그먼트의 마지막에 입력받은 값 추가
			segLen += 1     // 세그먼트 길이 증가
			ans += 1        // LIS 증가

			// 2. 입력받은 값의 세그먼트의 마지막 값보다 작거나 같은 경우
		} else {

			// 이분 탐색: lower bound, 즉 세그먼트의 m번 인덱스의 값보다 t가 작거나 같은 경우를 찾아 최적의 위치에 t를 삽입
			// 왜? t보다 작거나 같은 값의 위치를 찾아서 t를 삽입해야 하는가?

			// 10, 20, 40을 입력받아 seg는 []int{10, 20, 40}, ans는 3이라고 생각해보자

			// 새로운 입력값으로 25를 입력받는다
			// 25는 20보다 크고 40보다 작거나 같으므로 seg = []int{10, 20, 25}가 되지만 ans는 3으로
			// 입력된 값들의 LIS는 {10, 20, 40}, {10, 20, 25} 2가지인 것이다

			// 다음 입력으로 30을 입력받는다
			// 30은 25보다 크므로 seg = []int{10, 20, 25, 30}, ans는 1을 더해 4가 되며 여기서 LIS가 {10, 20 25, 30}으로 갱신된다
			// 즉, 입력되는 값들의 최적의 위치를 찾아줌으로써 이전의 LIS 형태와 길이를 기억하지 않고도 LIS의 최댓값을 찾을 수 있다

			// 그러나 세그먼트가 LIS가 되는 것은 아니다
			// 입력값이 10 20 40 50 30 인 경우,
			// LIS는 10 20 40 50 이지만, 세그먼트는 10 20 30 50이 되기 때문이다
			l, r = 0, segLen
			for l < r {
				m = (l + r) / 2
				if seg[m] >= t {
					r = m
				} else {
					l = m + 1
				}
			}
			seg[r] = t
		}

		/*
			예제 입력:

			6
			10 20 10 30 20 50

			seg:

			1. 10 입력: [0 10] 조건문 1번, ans = 1

			2. 20 입력: [0 10 20] 조건문 1번, ans = 2

			3. 10 입력: [0 10 20] 조건문 2번, r = 1, ans = 2

			4. 30 입력: [0 10 20 30] 조건문 1번, ans = 3

			5. 20 입력: [0 10 20 30] 조건문 2번, r = 2, ans = 3

			6. 50 입력: [0 10 20 30 50] 조건문 1번, ans = 4
		*/
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
