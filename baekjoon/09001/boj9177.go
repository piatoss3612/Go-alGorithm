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
	a, b, c []byte
	dp      [][]int
)

// 메모리: 17900KB
// 시간: 44ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()

	for i := 1; i <= t; i++ {
		a, b, c = scanBytes(), scanBytes(), scanBytes()

		// dp 초기화
		// dp[len(a)+1][len(b)+1]
		// +1씩 해준 이유는 어느 한쪽에 있는 문자를 모두 사용한 경우에 대비
		dp = make([][]int, len(a)+1)
		for i := 0; i < len(dp); i++ {
			dp[i] = make([]int, len(b)+1)
		}

		if rec2(0, 0, 0) == len(c) {
			fmt.Fprintf(writer, "Data set %d: yes\n", i)
		} else {
			fmt.Fprintf(writer, "Data set %d: no\n", i)
		}
	}
}

// 단어를 직접 비교하지 않고 메모이제이션을 활용하여 길이만으로 비교하는 풀이
func rec2(x, y, z int) int {
	// 기저 사례: 모든 비교가 끝난 경우
	if x == len(a) && y == len(b) {
		return 0
	}

	ret := &dp[x][y]

	if *ret != 0 {
		return *ret
	}

	// a에 있는 문자를 모두 사용한 경우
	if x == len(a) {
		if b[y] == c[z] {
			*ret = max(*ret, rec2(x, y+1, z+1)+1)
		}
		return *ret
	}

	// b에 있는 문자를 모두 사용한 경우
	if y == len(b) {
		if a[x] == c[z] {
			*ret = max(*ret, rec2(x+1, y, z+1)+1)
		}
		return *ret
	}

	// c는 반드시 a+b만큼의 길이를 가지므로 여기서는 궂이 유효성을 검사할 필요가 없다

	// a의 x번째 문자와 c의 z번째 문자가 같은 경우
	if a[x] == c[z] {
		*ret = max(*ret, rec2(x+1, y, z+1)+1) // 최댓값 갱신
	}

	// b의 y번째 문자와 c의 z번째 문자가 같은 경우
	if b[y] == c[z] {
		*ret = max(*ret, rec2(x, y+1, z+1)+1) // 최댓값 갱신
	}
	return *ret // 최댓값 반환
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 시간초과: 실제 문자열을 더해가면서 매칭하기 때문
func rec1(x, y, z int) []byte {
	if x == len(a) {
		return b[y:len(b)]
	}

	if y == len(b) {
		return a[x:len(a)]
	}

	var res []byte

	if a[x] == c[z] {
		tmp1 := rec1(x+1, y, z+1)
		if len(tmp1)+1 > len(res) {
			res = append(c[z:z+1], tmp1...)
		}
	}

	if b[y] == c[z] {
		tmp2 := rec1(x, y+1, z+1)
		if len(tmp2)+1 > len(res) {
			res = append(c[z:z+1], tmp2...)
		}
	}
	return res
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
