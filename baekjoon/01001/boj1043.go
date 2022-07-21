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
	knowTruth      []bool  // 진실을 아는 사람
	cannotJoin     []bool  // 파티에 참여할 수 있는지 여부
	parties        [][]int // 파티 정보
	memberBelongTo [][]int // 누가 어느 파티에 속해 있는지
	n, m           int
)

// 메모리: 948KB
// 시간: 4ms
// 거짓말아~ 지민이말 잘 듣고 오래오래 행복해야 된다~
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	knowTruth = make([]bool, n+1)
	cannotJoin = make([]bool, m+1)
	parties = make([][]int, m+1)
	memberBelongTo = make([][]int, n+1)

	// 큐를 사용하여 진실을 아는 사람이 누군지 저장
	q := []int{}

	t := scanInt()
	for i := 1; i <= t; i++ {
		temp := scanInt()
		knowTruth[temp] = true
		q = append(q, temp)
	}

	for i := 1; i <= m; i++ {
		p := scanInt()
		party := make([]int, p)
		for j := 0; j < p; j++ {
			party[j] = scanInt()
			memberBelongTo[party[j]] = append(memberBelongTo[party[j]], i)
		}
		parties[i] = party
	}

	// 큐에 진실을 아는 사람을 넣고 빼는 작업 반복
	for len(q) > 0 {
		x := q[0] // x는 진실을 아는 사람
		q = q[1:]

		// x가 속한 파티 탐색
		for i := 0; i < len(memberBelongTo[x]); i++ {
			// x가 속한 i번째 파티에 지민이가 참여할 수 있다면
			if !cannotJoin[memberBelongTo[x][i]] {
				cannotJoin[memberBelongTo[x][i]] = true // 이제는 못간다 애송이
				// 못가게 된 파티의 멤버들을 탐색
				for j := 0; j < len(parties[memberBelongTo[x][i]]); j++ {
					// 만약 파티의 j번째 멤버가 아직 진실을 모른다면
					if !knowTruth[parties[memberBelongTo[x][i]][j]] {
						knowTruth[parties[memberBelongTo[x][i]][j]] = true // 이제는 알게될 거야
						q = append(q, parties[memberBelongTo[x][i]][j])    // q에 진실을 알게된 멤버를 추가
					}
				}
			}
		}
	}

	cnt := 0
	// 참여할 수 있는 파티의 수를 카운트
	for i := 1; i <= m; i++ {
		if !cannotJoin[i] {
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
