package main

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	binary  []int
)

// 메모리: 920KB
// 시간: 4ms
// 그리디 알고리즘, 비트마스크
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	A := scanInt()
	binary = make([]int, 0, 61) // 2^59 < 10^18 < 2^60
	// int 슬라이스에 A를 2진수 형태로 변환하여 저장
	// 0번 인덱스는 A의 최하위 비트
	// ex. A=43 binary =[]int{1,1,0,1,0,1}
	for A > 0 {
		binary = append(binary, A%2)
		A /= 2
	}

	// 그리디 알고리즘으로 최적해 찾기

	lb := LowerBound() // A보다 작고 A에 가까우며 이진수로 A와 같은 1의 개수를 가지는 수
	ub := UpperBound() // A보다 크고 A에 가까우며 이진수로 A와 같은 1의 개수를 가지는 수
	fmt.Fprintln(writer, lb, ub)
}

func UpperBound() int {
	ub := 0

	/*
		43(101011)보다 크고 43에 가까우며 이진수로 A와 같은 1의 개수를 가지는 수 B

		B를 구하는 최적의 방법은 최하위비트부터 탐색하면서
		맞붙어있는 1(하위비트)과 0(상위비트)을 발견하면 자리를 스왑하는 것이다

		즉 101011의 2번째 1과 3번째 0을 스왑하여 101101, 45라는 값을 구할 수 있게 된다

		그런데 이렇게 구해진 B는 A보다 큰 최솟값이라는 것이 보장되지 않는다

		예를 들어 A가 22(10110)인 경우
		3번째 1과 4번째 0을 스왑하게 되면 26(11010)이라는 값을 얻게 된다
		그런데 2번째 1을 오른쪽으로 1만큼 시프트하면 1의 개수가 같으면서 더 작은 값인 25(11001)를 얻을 수 있게 된다

		즉, i번째 1과 i+1번째 0을 스왑한 후에
		i번째 비트 이하의 모든 1들을 최하위 비트부터 채워넣음으로써
		A보다 큰 최솟값을 찾을 수 있다
	*/

	for i := 0; i < len(binary)-1; i++ {
		// 0과 1을 스왑할 수 있는 경우
		if binary[i] == 1 && binary[i+1] == 0 {
			// i번째 1과 i+1번째 0을 스왑하고
			// i+1번째 비트부터 탐색하여 ub를 갱신한다
			binary[i], binary[i+1] = binary[i+1], binary[i]
			for j := i + 1; j < len(binary); j++ {
				if binary[j] == 1 {
					ub |= 1 << j
				}
			}

			// 남은 1의 개수를 구하고
			// 최하위 비트부터 채워넣는다
			cnt := 0
			for j := 0; j <= i; j++ {
				if binary[j] == 1 {
					cnt++
				}
			}

			for j := 0; j < cnt; j++ {
				ub |= 1 << j
			}

			return ub
		}
	}

	// 0과 1을 스왑할 수 없는 경우
	// 7, 15, 31(111, 1111, 11111)처럼 1로만 채워진 경우
	// 1로만 채워진 2진수보다 크고 최솟값이면서 1의 개수가 동일한 수는
	// 최상위 비트를 왼쪽으로 1만큼 시프트하고 나머지를 1로 채운
	// 10111, 101111, 1011111 형태로 표현된다
	cnt := 0
	for i := 0; i < len(binary)-1; i++ {
		if binary[i] == 1 {
			cnt++
		}
	}

	for i := 0; i < cnt; i++ {
		ub |= 1 << i
	}

	ub |= 1 << len(binary)

	return ub
}

func LowerBound() int {
	lb := 0

	/*
		43(101011)보다 작고 43에 가까우며 이진수로 A와 같은 1의 개수를 가지는 수 B

		B를 구하는 최적의 방법은 최하위비트부터 탐색하면서
		맞붙어있는 0(하위비트)과 1(상위비트)을 발견하면 자리를 스왑하는 것이다

		이 경우에도 UpperBound와 마찬가지로 B가 A보다 작은 최댓값이라는 것이 보장되지 않는다

		예를 들어 A가 25(11001)인 경우
		3번째 0과 4번째 1을 스왑하게 되면 21(10101)이라는 값을 얻게 된다
		그런데 1번째 1을 왼쪽으로 1만큼 시프트하면 1의 개수가 같으면서 더 큰 값인 22(10110)를 얻을 수 있게 된다

		즉, i번째 0과 i+1번째 1을 스왑한 후에
		i번째 비트 이하의 모든 1들을 i번째 비트부터 하위 비트로 채워넣음으로써
		A보다 작은 최댓값을 찾을 수 있다
	*/

	for i := 0; i < len(binary)-1; i++ {
		// 0과 1을 스왑할 수 있는 경우
		if binary[i] == 0 && binary[i+1] == 1 {
			binary[i], binary[i+1] = binary[i+1], binary[i]
			// i+1부터 탐색하여 lb 갱신
			for j := i + 1; j < len(binary); j++ {
				if binary[j] == 1 {
					lb |= 1 << j
				}
			}

			// i번째 비트부터 최하위 비트까지
			// 모든 1의 개수를 구하고
			cnt := 0
			for j := 0; j <= i; j++ {
				if binary[j] == 1 {
					cnt++
				}
			}

			// i번째 비트부터 채워넣는다
			j := i
			for ; cnt > 0; cnt-- {
				lb |= 1 << j
				j--
			}

			// 슬라이스는 포인터 참조 방식을 사용하므로
			// 스왑했던 구간을 원래대로 되돌려 놓는다
			binary[i], binary[i+1] = binary[i+1], binary[i]
			return lb
		}
	}

	return lb
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
