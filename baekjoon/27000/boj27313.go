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
	anime   []int
)

// 27313. 효율적인 애니메이션 감상
// 난이도: Gold 4
// 메모리: 3388KB
// 시간: 44ms
// 분류: 이분 탐색, 매개 변수 탐색, 정렬, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M, K = scanInt(), scanInt(), scanInt()
	anime = make([]int, N)
	for i := 0; i < N; i++ {
		anime[i] = scanInt()
	}
	sort.Ints(anime) // 오름차순 정렬
}

func Solve() {
	l, r := 0, N-1 // anime 슬라이스의 시작과 끝 인덱스
 
	// 이분 탐색을 통해 최대 시청 가능한 애니메이션 수를 구함
	for l <= r {
		mid := (l + r) / 2
		// mid개의 애니메이션을 M시간 내에 시청할 수 있는지 확인
		if possible(mid) {
			l = mid + 1 // 가능하다면 mid를 늘려서 더 많은 애니메이션을 시청할 수 있는지 확인
		} else {
			r = mid - 1 // 불가능하다면 mid를 줄여서 더 적은 애니메이션을 시청할 수 있는지 확인
		}
	}

	fmt.Fprintln(writer, r+1) // 0~r까지의 애니메이션을 시청할 수 있으므로 r+1 출력
}

func possible(param int) bool {
	total := 0 // 애니메이션을 시청한 시간

	/* 
		0~param까지의 애니메이션을 K개씩 묶어서 시청하는 것이 가장 효율적이나
		앞에서부터 K개씩 묶을지 뒤에서부터 K개씩 묶을지가 문제

		예제 입력 2의 경우
		3 15 2
		10 5 10
		이므로 anime 슬라이스를 오름차순 정렬하면
		5 10 10
		이 된다. 이 경우 앞에서부터 K개씩 묶으면
		(5 10) (10)
		이 되고, 뒤에서부터 K개씩 묶으면
		(5) (10 10)
		이 된다. 앞에서부터 묶을 경우 20분, 뒤에서부터 묶을 경우 15분이 걸리므로
		뒤에서부터 묶는 것이 더 효율적이다.

		마찬가지로 애니메이션을 K개씩 묶을 때 총 시청 시간은
		각 그룹의 최댓값의 합이다. 이 문제의 경우는 뒤에서부터 묶음으로써 최적의 해를 구할 수 있다.
	*/

	// 뒤에서부터 K개씩 묶어서 시청
	for i := param; i >= 0; i -= K {
		total += anime[i] // 각 그룹의 최댓값을 더함
	}

	return total <= M // param+1개의 애니메이션을 M시간 내에 시청할 수 있는지 반환
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