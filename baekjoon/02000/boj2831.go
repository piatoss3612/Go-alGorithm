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
	higherM []int // 자신보다 키가 큰 여성을 선호하는 남자
	lowerM  []int // 자신보다 키가 작은 여성을 선호하는 남자
	higherW []int // 자신보다 키가 큰 남성을 선호하는 여자
	lowerW  []int // 자신보다 키가 작은 여성을 선호하는 여자
	N       int
)

// 메모리: 3912KB
// 시간: 56ms
// 두 포인터, 그리디 알고리즘
// 문제가 안풀리면 문제 쪼개기, 시간 초과나면 문제 쪼개기, 쪼개기 쪼개기 쪼개기 쪼개기 쪼개기 쪼개기 쪼개기 쪼개기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	N = scanInt()

	higherM = make([]int, 0, N)
	lowerM = make([]int, 0, N)
	higherW = make([]int, 0, N)
	lowerW = make([]int, 0, N)

	for i := 0; i < N; i++ {
		x := scanInt()
		if x >= 0 {
			higherM = append(higherM, x) // 키가 큰 여성 선호
		} else {
			lowerM = append(lowerM, -x) // 키가 작은 여성 선호
		}
	}

	for i := 0; i < N; i++ {
		y := scanInt()
		if y >= 0 {
			higherW = append(higherW, y) // 키가 큰 남성 선호
		} else {
			lowerW = append(lowerW, -y) // 키가 작은 남성 성호
		}
	}

	// 정렬
	sort.Ints(lowerM)
	sort.Ints(higherM)
	sort.Ints(lowerW)
	sort.Ints(higherW)

	ans := 0

	// 두 포인터 1
	m, w := 0, 0
	for m < len(lowerM) && w < len(higherW) {
		// 작은 키의 여성을 선호하는 남자의 키가 큰 키를 선호하는 여성보다 큰 경우
		if lowerM[m] > higherW[w] {
			m++
			w++
			ans++
		} else {
			m++
		}
	}

	// 두 포인터 2
	m, w = 0, 0
	for m < len(higherM) && w < len(lowerW) {
		// 작은 키의 남성을 선호하는 여자의 키가 큰 키를 선호하는 남성보다 큰 경우
		if lowerW[w] > higherM[m] {
			m++
			w++
			ans++
		} else {
			w++
		}
	}

	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
