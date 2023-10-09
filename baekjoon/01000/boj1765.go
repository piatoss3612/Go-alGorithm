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
	team    []int   // 각 학생이 속한 팀의 리더
	enemy   [][]int // 각 학생의 원수를 저장하는 슬라이스
	N, M    int
)

// 메모리: 1232KB
// 시간: 16ms
// 분리 집합 문제
// 왜 친구 관계는 저장 않하고 원수 관계만 저장하는가?

// 처음에는 친구의 친구를 찾기 위해 친구 관계를 저장했지만
// union-find 과정에서 친구와 친구의 친구 또한 같은 팀으로 구성된다
// 원수의 원수가 친구가 되었을 때는 친구(원수의 원수)의 친구들도
// 자동으로 팀으로 편입되므로 따로 따져볼 필요가 없다고 결론내렸다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	team = make([]int, N+1)
	enemy = make([][]int, N+1)

	for i := 1; i <= N; i++ {
		team[i] = i
	}

	var F rune
	var p, q int
	for i := 1; i <= M; i++ {
		F, p, q = scanRune(), scanInt(), scanInt()
		switch F {
		case 'F':
			union(p, q)
			p = team[p]
		case 'E':
			enemy[p] = append(enemy[p], q)
			enemy[q] = append(enemy[q], p)
		}
	}

	for i := 1; i <= N; i++ {

		// directEnemy: i번째 학생의 직접적인 원수
		for _, directEnemy := range enemy[i] {
			// enemyOfEnemy: 원수의 원수는 친구가 된다
			// 따라서 union 연산을 실행
			for _, enemyOfEnemy := range enemy[directEnemy] {
				union(i, enemyOfEnemy)
			}
		}
	}

	cnt := 0

	for i := 1; i <= N; i++ {
		find(i) // 삑사리날 수도 있으니까 확실하게 i가 속한 팀의 리더 정보 갱신

		// 팀의 리더가 자기자신인 경우
		if team[i] == i {
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
}

func find(x int) int {
	if team[x] == x {
		return x
	}
	team[x] = find(team[x])
	return team[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		team[y] = x
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanRune() rune {
	scanner.Scan()
	return rune(scanner.Bytes()[0])
}
