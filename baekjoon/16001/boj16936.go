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
	n       int
	input   []int
	dic     map[int]int
)

func main() {
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	input = make([]int, n)
	dic = make(map[int]int)
	// 입력값으로 받는 값을 맵에 저장
	// 3을 3개 입력받으면 dic[3] = 3이 된다
	for i := 0; i < n; i++ {
		x := scanInt()
		input[i] = x
		dic[x] += 1
	}

	// 답이 반드시 존재하므로
	// i번째 입력값으로 시작하는 슬라이스를 매개변수로 넘겨
	// 백트래킹 기법으로 문제 풀이
	for i := 0; i < n; i++ {
		solve([]int{input[i]})
	}
}

func solve(arr []int) {
	// 슬라이스의 길이가 n인 경우 출력하고 프로그램 종료
	if len(arr) == n {
		for _, v := range arr {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
		writer.Flush()
		os.Exit(0)
	}

	x := arr[len(arr)-1]

	// 나3: 맵에 x/3을 키로 갖는 값이 0보다 큰 경우, x는 3으로 나누어 떨어져야 한다
	if x%3 == 0 && dic[x/3] > 0 {
		dic[x/3] -= 1
		tmp := append(arr, x/3)
		solve(tmp)
		dic[x/3] += 1
	}

	// 곱2: 맵에 x*2를 키로 갖는 값이 0보다 큰 경우
	if dic[x*2] > 0 {
		dic[x*2] -= 1
		tmp := append(arr, x*2)
		solve(tmp)
		dic[x*2] += 1
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
