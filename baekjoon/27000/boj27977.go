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
	L, N, K int
	station [100002]int // 충전소 위치: N의 최댓값이 100,000이지만 도착지점을 추가하기 위해 100,001까지 선언
)

// 난이도: Gold 4
// 메모리: 2532KB
// 시간: 24ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	L, N, K = scanInt(), scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		station[i] = scanInt()
	}
	station[N+1] = L // 도착지점 추가
}

func Solve() {
	l, r := 1, L // 이분 탐색을 위한 left, right

	for l <= r {
		mid := (l + r) / 2
		// mid 만큼의 배터리를 가진 킥보드로 충전소에 최대 K번 들리면서 도착지점까지 갈 수 있는지 확인
		if possible(mid) {
			r = mid - 1 // 가능하다면 배터리를 줄여서 다시 확인
		} else {
			l = mid + 1 // 불가능하다면 배터리를 늘려서 다시 확인
		}
	}

	fmt.Fprintln(writer, l) // 최소 배터리 출력
}

func possible(param int) bool {
	cnt := 0 			    // 충전 횟수
	here := station[1]      // 1번 충전소를 방문했다고 가정
	battery := param - here // param 만큼의 배터리를 가지고 출발하여 1번 충전소를 방문했을 때 남은 배터리

	// 1번 충전소를 방문했을 때 남은 배터리가 0보다 작다면 잘못된 값
	if battery < 0 {
		return false
	}

	// 2번 충전소부터 N+1번(도착지) 충전소까지 순회
	for i := 2; i <= N+1; i++ {
		next := station[i] // 다음 충전소

		// 다음 충전소까지의 거리가 param 보다 크다면 이동할 수 없음
		if next-here > param {
			return false
		}

		// 다음 충전소까지의 거리가 남은 배터리보다 크다면 충전해야함
		if battery == 0 || battery < next-here {
			battery = param - (next - here) // 충전하고 이동
			cnt++
		} else {
			battery -= next - here // 이동
		}

		// 충전 횟수가 K보다 크다면 잘못된 값
		if cnt > K {
			return false
		}

		here = next // 현재 충전소를 다음 충전소로 변경
	}

	return true
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
