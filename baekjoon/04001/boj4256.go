package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	preorder []int
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 0; i < t; i++ {
		testCase()
	}
}

func testCase() {
	n := scanInt()
	preorder = make([]int, n)
	inorder := make([]int, n)
	for i := 0; i < n; i++ {
		preorder[i] = scanInt()
	}
	for i := 0; i < n; i++ {
		inorder[i] = scanInt()
	}

	rec(inorder)
	fmt.Fprintln(writer)
}

// 메모리: 	3928KB
// 시간: 48ms

/*
# 전위 순회와 중위 순회의 결과를 가지고 후위 순회 결과를 출력하기

전위 순회: 부모 노드 출력 -> 왼쪽 탐색 -> 오른쪽 탐색
중위 순회: 왼쪽 탐색 -> 부모 노드 출력 -> 오른쪽 탐색
후위 순회: 왼쪽 탐색 -> 오른쪽 탐색 -> 부모 노드 출력

후위 순회는 전위 순회에서 부모 노드를 출력하는 것을 마지막으로 미룬 형태.
즉, 전위 순회 결과를 앞에서부터 추적하면서 해당 노드를 마지막에 출력하면 될 것이다.
그런데 왼쪽과 오른쪽에 어떤 노드들이 있는지 어떻게 알 수 있을까?

중위 순회의 결과는 (왼쪽 노드들) - 부모 노드 - (오른쪽 노드들)의 형태를 띄고 있다.
즉, 부모 노드가 무엇인지 알 수 있다면 중위 순회 결과를 왼쪽과 오른쪽으로 분할하여 재귀 호출하면 된다.
여기서 우리는 전위 순회 결과를 가지고 있으므로 부모 노드가 무엇인지 알 수 있다.
*/

func rec(sub []int) {
	if len(preorder) == 0 {
		return
	}

	tmp := preorder[0] // 부모 노드
	preorder = preorder[1:]

	// 중위 순회 결과에서 부모 노드와 일치하는 값의 인덱스 찾기
	idx := 0
	for i := 0; i < len(sub); i++ {
		if sub[i] == tmp {
			idx = i
			break
		}
	}

	// 중위 순회 결과에서 부모 노드와 일치하는 값을 기준으로 왼쪽, 오른쪽으로 분할
	left := sub[:idx]
	right := sub[idx+1:]

	// 왼쪽 중위 순회 결과의 길이가 0이 아니라면 재귀 호출
	if len(left) != 0 {
		rec(left)
	}

	// 오른쪽 중위 순회 결과의 길이가 0이 아니라면 재귀 호출
	if len(right) != 0 {
		rec(right)
	}

	// 부모 노드 출력
	fmt.Fprint(writer, tmp, " ")
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
