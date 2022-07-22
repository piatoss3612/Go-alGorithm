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
	target  []int // i번째 학생이 팀을 맺고 싶은 학생
	visited []int // 팀 결성 여부
	count   int   // 팀에 속하지 못한 학생 수
	T, n    int
)

// 메모리: 	61256KB
// 시간: 648ms
// 깊이 우선 탐색
// 비트마스킹을 사용할 수 있지 않을까 했는데 학생 번호가 10만까지여서 포기
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	// T개의 테스트 케이스 실행
	for i := 1; i <= T; i++ {
		TestCase()
	}
}

func TestCase() {
	count = 0
	n = scanInt()
	target = make([]int, n+1)
	visited = make([]int, n+1)
	for i := 1; i <= n; i++ {
		target[i] = scanInt()
		if target[i] == i { // 혼자 팀을 구성하는 학생은 미리 방문 처리
			visited[i] = i
		}
	}

	// 덱(dequeue)라고 하기에는 민망하지만 비슷한 역할을 한다
	// 같은 팀이 되기를 희망하는 학생 관계의 체인을 저장
	dq := []int{}

	for i := 1; i <= n; i++ {
		if visited[i] == 0 { // 아직 방문 처리하지 않은 학생
			visited[i] = i     // i값으로 임시 팀 번호 부여
			dq = append(dq, i) // 덱에 i번 학생 추가
			DFS(i, i, &dq)     // 깊이 우선 탐색 시작

			// 팀을 구성하고 덱에 남은 친구들
			for len(dq) > 0 {
				visited[dq[0]] = -1
				dq = dq[1:]
				count++
			}
		}
	}

	fmt.Fprintln(writer, count)
}

// member: 현재 학생 번호
// team: 팀 번호
// dq: 학생 관계 체인을 나타내는 정수형 슬라이스의 포인터
// 슬라이스가 이미 참조값이라고 해도 슬라이스의 포인터를 사용하는 것이 더 정확하고 안전하다
func DFS(member, team int, dq *[]int) {
	// 현재 학생이 팀을 맺고 싶어하는 학생이 이미 같은 팀에 속해있다면
	// 해당 학생으로부터 시작되는 사이클을 덱에서 제거하고
	// 깊이 우선 탐색을 종료한다
	if visited[target[member]] == team {
		idx := 0
		for i := len(*dq) - 1; i >= 0; i-- {
			if (*dq)[i] == target[member] {
				idx = i
				break
			}
		}

		// 사이클 제거
		*dq = (*dq)[:idx]
		return
	}

	// 현재 학생이 팀을 맺고 싶어하는 학생이 아직 팀을 정하지 않았다면
	// 팀으로 임시 영입하고 깊이 우선 탐색으로 그 다음 학생을 탐색한다
	if visited[target[member]] == 0 {
		visited[target[member]] = team
		*dq = append(*dq, target[member])
		DFS(target[member], team, dq)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
