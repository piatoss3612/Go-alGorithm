package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N, M, K int
	beer    []Beer
)

type Beer struct {
	Preference int
	Alcohol    int
}

// 난이도: Silver 1
// 메모리: 7196KB
// 시간: 136ms
// 분류: 이분 탐색, 정렬, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	beer = make([]Beer, K)
	for i := 0; i < K; i++ {
		beer[i] = Beer{Preference: scanInt(), Alcohol: scanInt()}
	}
	// 정렬 - 맥주의 선호도가 높은 순으로, 선호도가 같다면 도수가 낮은 순으로
	sort.Slice(beer, func(i, j int) bool {
		if beer[i].Preference == beer[j].Preference {
			return beer[i].Alcohol < beer[j].Alcohol
		}
		return beer[i].Preference > beer[j].Preference
	})
}

func Solve() {
	l, r := 1, 1<<31

	// 이분 탐색
	for l <= r {
		mid := (l + r) / 2
		// 간 레벨이 mid일 때, N일 동안 매일 맥주를 마셔서 M 이상의 선호도를 얻을 수 있는지 확인
		if Check(mid) {
			r = mid - 1 // 가능하면 레벨을 낮춰서 다시 확인
		} else {
			l = mid + 1 // 불가능하면 레벨을 높여서 다시 확인
		}
	}

	if l == 1<<31+1 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, l) // 간 레벨의 최솟값 출력
	}
}

func Check(mid int) bool {
	cnt := 0 // 마신 맥주의 수
	sum := 0 // 마신 맥주의 선호도 합
	for i := 0; i < K; i++ {
		// 도수가 mid보다 크면 무시
		if beer[i].Alcohol > mid {
			continue
		}

		cnt++ // 맥주를 마심
		sum += beer[i].Preference // 선호도 합산
		// N개의 맥주를 마시고, 선호도 합이 M 이상이면 true
		if cnt == N && sum >= M {
			return true
		}
	}

	// N개의 맥주를 마시지 못하거나, 선호도 합이 M 미만이면 false
	return false
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
