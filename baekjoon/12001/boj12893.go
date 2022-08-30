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
	N, M    int
	enemies [2001]int
	parent  [2001]int
)

// 메모리: 5748KB
// 시간: 228ms
// 적의 적은 친구가 된다는 이론이 성립하는지 분리 집합을 통해 검증하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	N, M = scanInt(), scanInt()

	// 자기 자신을 그룹의 부모로 초기화
	for i := 1; i <= N; i++ {
		parent[i] = i
	}

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		a, b = find(a), find(b) // a가 속한 그룹과 b가 속한 그룹 찾기

		// a와 b가 같다면 이미 친구 관계라는 의미이며
		// 이 둘이 적대관계일 수는 없으므로 모순이 발생한다
		// 따라서 이론은 성립하지 않는다
		if a == b {
			fmt.Fprintln(writer, 0)
			return
		}

		// a의 적이 존재한다면
		if enemies[a] != 0 {
			union(b, enemies[a]) // a와 b는 적대관계이므로 a의 적과 b는 친구관계
		} else {
			enemies[a] = b // a의 적은 b
		}

		// b의 적이 존재한다면
		if enemies[b] != 0 {
			union(a, enemies[b]) // a와 b는 적대관계이므로 b의 적과 a는 친구관계
		} else {
			enemies[b] = a // b의 적은 a
		}
	}

	// 성공적으로 적의 적은 친구 이론을 검증한 경우
	fmt.Fprintln(writer, 1)
}

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		parent[y] = x
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
