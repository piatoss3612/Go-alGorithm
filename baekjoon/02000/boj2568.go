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
	inp     []wire
	lis     []int
)

// 전봇대 A, B사이의 전선
type wire struct {
	l, r int
}

// 메모리: 7716KB
// 시간: 88ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	inp = make([]wire, n)
	for i := 0; i < n; i++ {
		inp[i] = wire{scanInt(), scanInt()}
	}

	// A 전봇대를 기준으로 전선의 위치를 오름차순으로 정렬
	// sort.Slice를 사용하기 위해 배열이 아닌 슬라이스를 사용했다
	sort.Slice(inp, func(i, j int) bool {
		return inp[i].l < inp[j].l
	})

	lis = make([]int, n)
	check := []int{0}
	maxIdx := 0
	cnt := 0

	for i := 0; i < n; i++ {
		if inp[i].r > check[cnt] {
			check = append(check, inp[i].r)
			maxIdx = i
			cnt += 1
			lis[i] = cnt
			continue
		}

		// 이분 탐색: lower bound
		l, r := 0, cnt
		for l < r {
			m := (l + r) / 2
			if inp[i].r <= check[m] {
				r = m
			} else {
				l = m + 1
			}
		}

		check[r] = inp[i].r
		lis[i] = r // i번째 입력값의 LIS는 r
	}

	fmt.Fprintln(writer, n-cnt)

	rm := []int{} // 제거해야될 전선

	// 위치가 maxIdx보다 큰 전선들을 모두 필요없는 상황이므로 제거
	for i := maxIdx + 1; i < n; i++ {
		rm = append(rm, inp[i].l)
	}

	cnt -= 1 // maxIdx 값을 제외

	// 위치가 maxIdx보다 작은 전선들을 체크
	for i := maxIdx - 1; i >= 0; i-- {
		// LIS에 포함되는 경우
		if cnt != 0 && lis[i] == cnt && inp[i].r < inp[maxIdx].r {
			cnt -= 1
			maxIdx = i
		} else {
			// LIS에 포함되지 않는 경우는 제거
			rm = append(rm, inp[i].l)
		}
	}

	sort.Ints(rm) // 위치 번호 오름차순 정렬
	for _, v := range rm {
		fmt.Fprintln(writer, v)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
