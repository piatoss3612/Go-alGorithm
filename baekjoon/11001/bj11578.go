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
	student []int // 각 학생이 풀 수 있는 문제를 비트마스크 기법으로 저장
	goal    int   // 모든 문제를 풀 수 있는 경우의 비트마스킹 값
	ans     = 987654321
)

// 메모리: 916KB
// 시간: 4ms
// 비트마스크, 브루트포스
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()

	// 학생이 풀 수 있는 문제를 비트마스트 기법으로 표현
	// 예를 들어, 예제 입력처럼 5개의 문제가 있는데
	// 1번 학생이 2번 4번 문제를 풀 수 있다면 01010으로 표현
	student = make([]int, M+1)
	for i := 1; i <= M; i++ {
		O := scanInt()
		for j := 1; j <= O; j++ {
			student[i] |= 1 << (scanInt() - 1)
		}
	}

	goal = (1 << N) - 1

	// 브루트포스: 첫번째 학생부터 마지막 학생까지 전수 조사
	for i := 1; i <= M; i++ {
		rec(student[i], i, 1)
	}

	// 모든 문제를 풀 수 있는 팀을 만들지 못한 경우
	if ans == 987654321 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func rec(solvable, member, count int) {
	// 모든 문제를 풀 수 있는 경우
	if solvable == goal {
		ans = min(ans, count) // 팀원 수를 최솟값으로 갱신
		return
	}

	// member 학생의 번호보다 빠른 번호를 가진 학생들은 조사 불필요
	// 왜? 전수 조사하면서 앞에서 이미 확인을 했으므로
	for i := member + 1; i <= M; i++ {
		rec(solvable|student[i], i, count+1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
