package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	snowball []int // 눈덩이의 지름
)

// 난이도: Gold 3
// 메모리: 932KB
// 시간: 184ms
// 분류: 정렬, 두포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

// 왜 87%에서 자꾸 틀리나 했는데 두 눈사람의 키차이의 최댓값이 10^9-1-(1-1)이라고 보면
// INF값이 999999999 이상이어야 하는데 987654321로 설정해서 그랬다
const INF = 9876543210

func Input() {
	N = scanInt()
	snowball = make([]int, N)
	for i := 0; i < N; i++ {
		snowball[i] = scanInt()
	}
	sort.Ints(snowball) // 눈덩이를 지름의 크기순으로 오름차순 정렬
}

func Solve() {
	ans := INF
	for i := 0; i < N-3; i++ {
		for j := i + 3; j < N; j++ {
			// 1. i와 j를 경계로하는 범위에서 i번째 눈덩이와 j번째 눈덩이로 첫번째 눈사람을 만든다
			// 2. l(init:i+1)번째 눈덩이와 r(init:j-1)번째 눈덩이로 두번째 눈사람을 만든다
			// 3. 첫번째 눈사람과 두번째 눈사람의 높이 차이의 절댓값을 구하고 ans를 최솟값으로 갱신한다
			// 4. 첫번째 눈사람의 높이가 두번째 눈사람의 높이보다 크다면 두번째 눈사람의 높이를 늘린다(l++)
			// 5. 첫번째 눈사람의 높이가 두번째 눈사람의 높이보다 작다면 두번째 눈사람의 높이를 줄인다(r--)
			// 6. 두번째 눈사람을 만들 수 있는 범위(i < l,r < j && l < r) 내에서 2~5번을 반복한다
			snowman1 := snowball[i] + snowball[j]
			l, r := i+1, j-1
			for l < r {
				snowman2 := snowball[l] + snowball[r]
				ans = min(ans, abs(snowman1-snowman2))
				if snowman2 > snowman1 {
					r--
				} else {
					l++
				}
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
