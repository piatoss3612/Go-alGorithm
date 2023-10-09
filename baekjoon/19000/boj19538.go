package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner     = bufio.NewScanner(os.Stdin)
	writer      = bufio.NewWriter(os.Stdout)
	N, M        int
	neighborOf  [][]int // i번 사람의 주변인들
	trustRumor  []int   // i번 사람이 루머를 믿기 시작하는 시간
	spreaders   []int   // 루머를 퍼트리는 사람들
	newTrusters []int   // 새로운 루머 신봉자들
	currentTime int     // 현재 시간
)

// 난이도: Gold 4
// 메모리: 88240KB
// 시간: 2804ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	neighborOf = make([][]int, N+1)
	trustRumor = make([]int, N+1)

	for i := 1; i <= N; i++ {
		for {
			neighbor := scanInt()
			if neighbor == 0 {
				break
			}

			neighborOf[i] = append(neighborOf[i], neighbor)
		}

		trustRumor[i] = -1
	}

	M = scanInt()

	for i := 0; i < M; i++ {
		spreader := scanInt()
		trustRumor[spreader] = 0
		spreaders = append(spreaders, spreader)
	}
}

func Solve() {
	// currentTime에 모든 루머 유포자가 동시에 루머를 유포하기 시작
	// spreaders의 길이가 0이 되어 모든 루머 유포자가 주변인을 설득하는 작업을 마친 경우에만
	// 새로운 루머 신봉자들을 교육시킬 수 있다
	for len(spreaders) > 0 {
		curr := spreaders[0]
		spreaders = spreaders[1:]

		findNewTrusters(curr) // 루머 유포자가 루머를 믿지 않는 주변인들 설득 시작

		// currentTime에 모든 루머 유포자가 주변인에게 루머를 유포하는 작업을 마친 경우
		if len(spreaders) == 0 {
			educateNewTrusters() // 새롭게 루머를 믿는 사람들을 교육하여 루머를 퍼트리도록 만든다
		}
	}

	for i := 1; i <= N; i++ {
		fmt.Fprintf(writer, "%d ", trustRumor[i])
	}
	fmt.Fprintln(writer)
}

// 주변인들 중에 루머를 믿는 사람들의 수 구하기
func trustersAmongNeighbors(idx int) int {
	cnt := 0

	for _, neighbor := range neighborOf[idx] {
		if trustRumor[neighbor] >= 0 {
			cnt++
		}
	}

	return cnt
}

// 새로운 루머 신봉자 찾기
func findNewTrusters(spreader int) {
	for _, neighbor := range neighborOf[spreader] {
		if trustRumor[neighbor] != -1 {
			continue
		}

		numOfTrusters := trustersAmongNeighbors(neighbor)
		if numOfTrusters >= (len(neighborOf[neighbor])+1)/2 {
			newTrusters = append(newTrusters, neighbor)
		}
	}
}

// 새로운 유머 신봉자를 교육하여 루머를 퍼트리도록 만들기
func educateNewTrusters() {
	currentTime++
	for len(newTrusters) > 0 {
		newbie := newTrusters[0]
		newTrusters = newTrusters[1:]
		trustRumor[newbie] = currentTime
		spreaders = append(spreaders, newbie)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
