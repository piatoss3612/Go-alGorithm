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
	inp     []Wire
	N       int
)

type Wire struct {
	line     int
	sw, bulb int
}

// 메모리: 2092KB
// 시간: 12ms
// 이분탐색을 사용하여 최장 증가 수열(LIS)를 구하는 문제인데
// LIS를 구성하는 요소들이 무엇인지 트래킹을 해야 한다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	inp = make([]Wire, N+1)

	for i := 1; i <= N; i++ {
		inp[i].line = i
		x := scanInt()
		inp[x].sw = i // x번 스위치의 위치는 i
	}

	for i := 1; i <= N; i++ {
		y := scanInt()
		inp[y].bulb = i // y번 전구의 위치는 i
	}

	inp = inp[1:] // 정렬하기 전에 0번 인덱스 항목 제거

	// 스위치 번호로 오름차순 정렬
	sort.Slice(inp, func(i, j int) bool {
		return inp[i].sw < inp[j].sw
	})

	LIS := []int{}
	order := make([]int, N)

	maxIdx := 0 // LIS가 가질 수 있는 정렬된 inp의 특정값의 인덱스의 최댓값
	cnt := -1   // inp의 특정값이 LIS의 몇 번째 자리에 올 수 있는지, LIS의 인덱스가 0부터 시작하므로 -1로 초기화

	for i := 0; i < N; i++ {
		if len(LIS) == 0 || inp[i].bulb > LIS[len(LIS)-1] {
			LIS = append(LIS, inp[i].bulb)
			maxIdx = i // maxIdx 갱신
			cnt++      // cnt 갱신
			order[i] = cnt
			continue
		}

		// 이분 탐색: lower bound 찾기
		l, r := 0, len(LIS)-1
		for l < r {
			m := (l + r) / 2
			if inp[i].bulb <= LIS[m] {
				r = m
			} else {
				l = m + 1
			}
		}

		LIS[r] = inp[i].bulb
		order[i] = r // inp의 i번째 값은 LIS의 r번째 자리에 들어갈 수 있다
	}

	fmt.Fprintln(writer, len(LIS))

	res := []int{inp[maxIdx].line}

	cnt-- // maxIdx를 제외하고 구해야하는 LIS 원소의 개수

	for i := maxIdx - 1; i >= 0; i-- {
		if cnt == -1 {
			break
		}

		// inp의 i번째 값이 LIS의 cnt 자리에 오고
		// inp의 i번째 값의 전구 위치가 LIS의 이전 원소의 전구 위치보다 작은 경우
		if order[i] == cnt && inp[i].bulb < inp[maxIdx].bulb {
			res = append(res, inp[i].line)
			cnt--      // 찾아야하는 LIS의 원소수 갱신
			maxIdx = i // LIS의 이전 원소 갱신
		}
	}

	sort.Ints(res)

	for _, v := range res {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
