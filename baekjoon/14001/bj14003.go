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
	inp     [1000001]int // LIS를 역추적하기 위해 어쩔 수 없이 입력값 저장
	lis     [1000001]int // 입력값이 가질 수 있는 LIS 최댓값을 저장
)

// 메모리: 69648KB
// 시간: 416ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()

	check := []int{-1000000001}
	maxIdx := 0 // 역추적을 시작하기 위해 LIS 값이 최댓값이 되는 인덱스를 저장
	cnt := 0    // LIS 최댓값

	for i := 1; i <= n; i++ {
		inp[i] = scanInt()
		// 입력값이 check 슬라이스의 마지막 값보가 큰 경우
		if inp[i] > check[cnt] {
			check = append(check, inp[i])
			maxIdx = i
			cnt += 1
			lis[i] = cnt
			continue
		}

		//이분 탐색: lower bound를 찾아 최적의 위치에 inp[i]를 삽입
		l, r := 0, cnt
		for l < r {
			m := (l + r) / 2
			if check[m] >= inp[i] {
				r = m
			} else {
				l = m + 1
			}
		}
		check[r] = inp[i]

		lis[i] = r // i번째 입력값이 가질 수 있는 LIS 최댓값은 r이 된다
	}

	fmt.Fprintln(writer, cnt)

	// 역추적 시작
	seq := []int{inp[maxIdx]}
	cnt -= 1

	for i := maxIdx - 1; i >= 1; i-- {
		if cnt != 0 && lis[i] == cnt && inp[i] < seq[len(seq)-1] {
			seq = append(seq, inp[i])
			cnt -= 1
		}
	}

	for i := len(seq) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%d ", seq[i])
	}
	fmt.Fprintln(writer)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
