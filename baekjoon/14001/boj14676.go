package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	N, M, K   int
	X, Y      int     // X는 Y를 건설하기 위해 필요한 건물
	nextLevel [][]int // nextLevel[i]: i번 건물을 건설한 다음에야 지을 수 있는 직접 연결된 상위 건물들
	inDegree  []int   // 각 건물을 건설하는데 필요한 직접 연결된 하위 건물의 개수 (진입 차수)
	preBuilt  []int   // preBuilt[i]: i번 건물을 건설하기 위해 필요한 하위 건물의 개수가 충족되었는지 여부
	count     []int   // 각 번호의 건물이 지어진 개수
)

// 난이도: Gold 3
// 메모리: 12944KB
// 시간: 80ms
// 분류: 구현, 그래프 이론, 위상 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M, K = scanInt(), scanInt(), scanInt()
	nextLevel = make([][]int, N+1)
	inDegree = make([]int, N+1)
	preBuilt = make([]int, N+1)
	count = make([]int, N+1)

	for i := 1; i <= M; i++ {
		X, Y = scanInt(), scanInt()
		nextLevel[X] = append(nextLevel[X], Y) // Y번 건물은 X번 건물의 다음 단계
		inDegree[Y]++                          // Y번 건물의 진입 차수 증가
	}

	ans := true
	for i := 1; i <= K; i++ {
		cmd, a := scanInt(), scanInt()

		// 입력 과정에서 치트를 사용했음이 밝혀진 경우
		if !ans {
			continue
		}

		// 명령어 cmd에 따라 분기 처리
		switch cmd {
		case 1:

			if preBuilt[a] == inDegree[a] {
				// a번 건물을 건설하기 위한 전제조건이 갖추어진 경우
				if count[a] == 0 {
					// a번 건물의 다음 단계 건물을 건설하기 위한 일부 조건이 충족된다
					for _, next := range nextLevel[a] {
						preBuilt[next]++
					}
				}
				count[a]++ // a번 건물의 개수 증가
			} else {
				// a번 건물을 건설하기 위한 전제조건이 갖추어지지 않은 경우
				// 건설할 수 없는 건물은 건설했으므로 치트를 사용했다고 판정
				ans = false
			}
		case 2:
			if count[a] > 0 {
				// a번 건물을 파괴할 수 있는 경우
				count[a]-- // a번 건물의 개수 감소
				if count[a] == 0 {
					// a번 건물의 개수가 0이된 경우
					// a번 건물의 다음 단게 건물을 건설하기 위한 일부 조건이 더 이상 충족될 수 없다
					for _, next := range nextLevel[a] {
						preBuilt[next]--
					}
				}
			} else {
				// a번 건물이 건설되어 있지 않은 경우
				// 파괴 할 수 없는 건물을 파괴했으므로 치트를 사용했다고 판정
				ans = false
			}
		}
	}

	if ans {
		fmt.Fprintln(writer, "King-God-Emperor")
	} else {
		fmt.Fprintln(writer, "Lier!")
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
