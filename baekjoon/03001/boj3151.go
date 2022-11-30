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
	student []int // 학생들의 코딩 실력
)

// 난이도: Gold 4
// 메모리: 1044KB
// 시간: 148ms
// 분류: 브루트포스 알고리즘, 정렬, 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	student = make([]int, N)
	for i := 0; i < N; i++ {
		student[i] = scanInt()
	}
	sort.Ints(student) // 코딩 실력을 오름차순으로 정렬
}

func Solve() {
	ans := 0
	// 브루트포스: 모든 경우의 수 탐색
	for i := 0; i < N-2; i++ {
		// 기저 사례: 두 포인터 탐색의 기준이 되는 i번째 학생의 코딩 실력이 0보다 크면
		// 세 학생의 코딩 실력이 0이 되는 경우를 찾을 수 없다
		if student[i] > 0 {
			break
		}

		l, r := i+1, N-1 // 두 포인터 l과 r
		for l < r {
			// 세 학생의 실력의 합에 따른 분기 처리
			switch sum := student[i] + student[l] + student[r]; {
			// 1. 세 학생의 실력의 합이 0인 경우
			case sum == 0:
				// 왼쪽 포인터와 오른쪽 포인터의 값이 같은 경우
				if student[l] == student[r] {
					ans += r - l // 왼쪽 포인터를 기준으로 r-l개의 경우의 수를 만들 수 있다
					l++          // 왼쪽 포인터를 오른쪽으로 한 칸 이동
					continue     // switch문 탈출, for문 초기 단계로 재진입
				}

				dl, dr := 1, 1 // 왼쪽 포인터와 중복된 값의 개수, 오른쪽 포인터와 중복된 값의 개수

				// 왼쪽 포인터와 중복된 값 찾기
				for student[l] == student[l+1] {
					dl++
					l++
				}

				// 오른쪽 포인터와 중복된 값 찾기
				for student[r] == student[r-1] {
					dr++
					r--
				}

				ans += dl * dr // 경우의 수 추가

				// l, r 포인터가 위치한 구간은 탐색을 마쳤으므로
				l++ // 왼쪽 포인터를 오른쪽으로 한 칸 이동
				r-- // 오른쪽 포인터를 왼쪽으로 한 칸 이동

			// 2. 0보다 큰 경우
			case sum > 0:
				r-- // 오른쪽 포인터를 왼쪽으로 한 칸 이동

			// 3. 0보다 작은 경우
			case sum < 0:
				l++ // 왼쪽 포인터를 오른쪽으로 한 칸 이동
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
