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
	N       int
	limit   [100001]int // i번째 대회 이전에 모은 상금에 대한 제한
	prize   [100001]int // i번째 대회의 상금
	dp      [100001]int // i번째 대회까지의 총 상금
)

// 난이도: Gold 4
// 메모리: 5780KB
// 시간: 48ms
// 분류: 다이나믹 프로그래밍, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Solve()
}

func Solve() {
	N = scanInt()
	absentCount := 0         // 대회에 불참한 횟수
	absentContestNumber := 0 // 불참한 대회 번호

	for i := 1; i <= N; i++ {
		limit[i], prize[i] = scanInt(), scanInt()

		// 2개 이상의 대회를 불참한 경우: 연두는 다시 냉동인간으로 돌아가야 한다...
		if absentCount == 2 {
			continue
		}

		// 대회 참여 조건을 만족한 경우
		if dp[i-1] <= limit[i] {
			dp[i] = dp[i-1] + prize[i]
			continue
		}

		// 처음으로 대회를 불참하는 경우: 일단 불참 처리하고 진행
		if absentCount == 0 {
			dp[i] = dp[i-1]
			absentContestNumber = i
			absentCount++
			continue
		}

		// 이미 불참 처리한 대회가 있는데 i번째 대회도 불참해야 될 위기에 놓인 경우:
		// 최대 1개의 대회에 불참하면서 i번째 대회에 참여하기 위해 limit[i]를 만족하도록 '총 상금을 하향 조정'할 필요가 있다

		// 총 상금을 조정하기 위해서
		// 1. 이전의 불참 선언한 absentContestNumber번째 대회를 참가할 수 있도록 만들어야 한다
		// 2. absentContestNumber번째 대회에 참가하기 위해 limit[absentContestNumber]를 만족시켜야 하므로
		// 3. absentContestNumber번째 대회 이전에 참여 처리한 대회 중 하나를 골라 불참 처리해야 한다
		// 4. 불참 처리하는 대회는 상금액이 가장 큰 대회로 선택하여 limit[absentContestNumber]를 만족하면서
		// 5. 이후에 열리는 대회의 참여 조건에도 대비할 수 있도록 한다 (그리디 알고리즘)
		// 5. 물론 이 과정을 거치고도 i번째 대회에 참여할 수 없는 경우는 냉동인간 엔딩

		// absentContestNumber 이전에 참여 처리한 대회 중 가장 상금이 큰 대회 선택
		tempContest, tempPrize := 0, 0
		for j := 1; j <= absentContestNumber-1; j++ {
			if prize[j] > tempPrize {
				tempPrize = prize[j]
				tempContest = j
			}
		}

		// 1. absentContestNumber 대회에 참여하기 위한 조건
		// 2. i번 대회에 참여하기 위한 조건
		// 조건 1과 2를 동시에 만족하는 경우
		if dp[absentContestNumber-1]-tempPrize <= limit[absentContestNumber] && dp[i-1]-tempPrize+prize[absentContestNumber] <= limit[i] {
			dp[i] = dp[i-1] - tempPrize + prize[absentContestNumber] + prize[i]
			absentContestNumber = tempContest // 불참하는 대회의 번호 갱신
		} else {
			// 어떠한 노력에도 i번째 대회에 참여할 수 없는 경우
			dp[i] = dp[i-1]
			absentCount++
		}
	}

	if absentCount == 2 {
		fmt.Fprintln(writer, "Zzz") // 냉동인간 엔딩
	} else {
		fmt.Fprintln(writer, "Kkeo-eok") // 상금 꺼억
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
