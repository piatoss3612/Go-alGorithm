package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	notPrime [2000001]bool // 소수 판별
	pp       []int         // 소수 & 팰린드롬인 수들
	N        int
)

const LIMIT = 2000000 // 1000000보다 크거나 같은 소수이면서 팰린드롬인 수를 찾기위한 경곗값 설정

func init() {
	// 0과 1은 소수가 아니다
	notPrime[0] = true
	notPrime[1] = true

	for i := 2; i <= LIMIT; i++ {
		// i가 소수인 경우
		if !notPrime[i] {
			// i의 배수들은 소수가 아니게 된다
			for j := i * i; j <= LIMIT; j += i {
				notPrime[j] = true
			}

			// i가 팰린드롬인지 판별
			origin := i
			reverse := 0

			for origin > 0 {
				reverse *= 10
				reverse += origin % 10
				origin /= 10
			}

			// i가 소수이면서 팰린드롬인 경우
			if reverse == i {
				pp = append(pp, i)
			}
		}
	}
}

// 난이도: Silver 1
// 메모리: 2904KB
// 시간: 20ms
// 분류: 브루트포스 알고리즘, 정수론, 소수 판정, 에라토스테네스의 체
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
}

func Solve() {
	// N보다 크거나 같은 소수이면서 팰린드롬인 수 찾기
	for _, v := range pp {
		if v >= N {
			fmt.Fprintln(writer, v)
			return
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
