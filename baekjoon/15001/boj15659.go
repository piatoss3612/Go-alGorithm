package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner        = bufio.NewScanner(os.Stdin)
	writer         = bufio.NewWriter(os.Stdout)
	N              int
	nums           []int                   // N개의 수
	oprs           [5]int                  // 사용할 수 있는 각 연산자의 개수
	selected       []int                   // 연산식을 구성하기 위해 선택한 연산자의 순서
	ansMax, ansMin = -987654321, 987654321 // 만들 수 있는 연산식의 결과가 최대인 것, 최소인 것
)

const limit = 1000000000 // 가능한 연산 결과의 최댓값

// 난이도: Gold 3
// 메모리: 7036KB
// 시간: 24ms
// 분류: 브루트포스 알고리즘, 백트래킹, 스택
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	nums = make([]int, N+1)
	selected = make([]int, N)
	for i := 1; i <= N; i++ {
		nums[i] = scanInt()
	}
	for i := 1; i <= 4; i++ {
		oprs[i] = scanInt()
	}
}

func Solve() {
	bruteForce(1)
	fmt.Fprintln(writer, ansMax)
	fmt.Fprintln(writer, ansMin)
}

// 브루트포스 알고리즘: 모든 경우의 수를 탐색
func bruteForce(cnt int) {
	// 기저 사례: N-1개의 연산자를 선택한 경우
	if cnt == N {
		calc()
		return
	}

	for i := 1; i <= 4; i++ {
		if oprs[i] > 0 {
			selected[cnt] = i   // i번 연산자를 선택
			oprs[i]--           // i번 연산자의 개수 감소
			bruteForce(cnt + 1) // 다음 연산자 선택
			oprs[i]++           // i번 연산자의 개수 되돌리기
		}
	}
}

func calc() {
	exp := []int{nums[1]}

	// N개의 수 중에 첫번째 수를 연산식에 먼저 추가하고
	// i-1번째 연산자와 i번째 수를 연산식에 추가하는 과정을 반복하여 연산식 구성
	for i := 2; i <= N; i++ {
		switch selected[i-1] {
		// 곱하기(3)과 나누기(4)는 연산자 우선순위를 고려하여 값을 미리 계산하여 연산식에 추가
		case 3:
			exp[len(exp)-1] *= nums[i]
		case 4:
			exp[len(exp)-1] /= nums[i]
		// 더하기(1)와 빼기(2)의 경우, 연산자에 limit를 더한 뒤에 연산식에 추가
		default:
			exp = append(exp, limit+selected[i-1], nums[i])
		}
	}

	// 더하기와 빼기의 연산자 우선순위를 고려하여 연산식의 앞에서부터 계산
	for len(exp) > 2 {
		num := exp[0]
		opr := exp[1] - limit
		switch opr {
		case 1:
			exp[2] = num + exp[2]
		case 2:
			exp[2] = num - exp[2]
		}
		exp = exp[2:]
	}

	ansMax = max(ansMax, exp[0]) // 최댓값 갱신
	ansMin = min(ansMin, exp[0]) // 최솟값 갱신
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
