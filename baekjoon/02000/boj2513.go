package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner     = bufio.NewScanner(os.Stdin)
	writer      = bufio.NewWriter(os.Stdout)
	left, right []*Apt // 학교를 기준으로 왼쪽, 오른쪽에 있는 아파트들
	N, K, S     int
)

// 아파트 정보
type Apt struct {
	pos     int // 위치
	student int // 학생 수
}

// 메모리: 2688KB
// 시간: 16ms
// 그리디 알고리즘: 학교를 기준으로 왼쪽, 오른쪽 어느 한 방향으로
// 거리가 가장 먼 아파트부터 들러서 학생들을 등교시키면 총 이동거리의 최솟값을 구할 수 있다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K, S = scanInt(), scanInt(), scanInt()
	for i := 0; i < N; i++ {
		temp := Apt{scanInt(), scanInt()}
		if temp.pos >= S {
			// 아파트 위치가 학교 위치 S보다 크거나 같은 경우
			right = append(right, &temp)
		} else {
			// 아파트 위치가 학교 위치 S보다 작은 경우
			left = append(left, &temp)
		}
	}

	// 학교의 오른쪽에 있는 아파트들을 위치값의 내림차순으로 정렬
	sort.Slice(right, func(i, j int) bool {
		return right[i].pos > right[j].pos
	})

	// 학교의 왼쪽에 있는 아파트들을 위치값의 오름차순으로 정렬
	sort.Slice(left, func(i, j int) bool {
		return left[i].pos < left[j].pos
	})

	ans := 0

	// 오른쪽에 있는 아파트들 탐색
	for len(right) > 0 {
		ridable := K
		// 학교에서 가장 먼 아파트를 먼저 들릴 때 편도 이동 거리
		dist := right[0].pos - S

		// 학교에서 가장 먼 아파트부터 탐색하면서
		// 가능한 버스에 남는 자리가 없도록
		// 돌아오는 길에도 학생을 태워서 꽉꽉채워야 한다
		for ridable > 0 && len(right) > 0 {
			// 버스에 태워야 하는 학생의 수가
			// 버스에 남은 자리보다 작거나 같다면
			if right[0].student <= ridable {
				// 해당 아파트는 다시 들리지 않아도 되므로 슬라이스에서 제거한다
				ridable -= right[0].student
				right = right[1:]
			} else if right[0].student > ridable {
				// 태워야 하는 학생 수가 태울 수 있는 학생 수보다 많다면
				// 다시 들려야만 한다
				right[0].student -= ridable
				ridable = 0
			}
		}

		ans += dist * 2 // 왕복 이동 거리 누적
	}

	for len(left) > 0 {
		ridable := K
		dist := left[0].pos - S // 항상 음수값을 가진다 (ex. 아파트 위치 0번 - 학교 위치 4번)

		for ridable > 0 && len(left) > 0 {
			if left[0].student <= ridable {
				ridable -= left[0].student
				left = left[1:]
			} else if left[0].student > ridable {
				left[0].student -= ridable
				ridable = 0
			}
		}

		ans -= dist * 2 // dist가 음수이므로 dist*2도 음수가 된다
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
