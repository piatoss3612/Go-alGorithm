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
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	inorder := make([]int, n)
	postorder := make([]int, n)
	for i := 0; i < n; i++ {
		inorder[i] = scanInt()
	}
	for i := 0; i < n; i++ {
		postorder[i] = scanInt()
	}

	rec(postorder, inorder)
	fmt.Fprintln(writer)
}

// 메모리: 35320KB
// 시간: 988ms

/*
4256번과 유사한 문제.
다만 이 문제에서는 전위 순회와 중위 순회 결과가 아닌, 중위 순회와 후위 순회 결과를 사용해 전위 순회 결과를 찾아야 한다.


중위 순회: (왼쪽 노드들) - 부모 노드 - (오른쪽 노드들)
후위 순회: (왼쪽 노드들) - (오른쪽 노드들) - 부모 노드

# 전위 순회 결과 출력

1. 후위 순회 결과를 기준으로 가장 오른쪽에 있는 부모 노드를 선택, 출력한다
2. 중위 순회 결과를 선택된 부모 노드를 기준으로 왼쪽과 오른쪽으로 나눈다
3. 부모 노드를 제거한 후위 순회 결과를 앞서 나뉘어진 왼쪽 노드들과 오른쪽 노드들의 길이 만큼 양분한다
4. 중위 순회 결과의 왼쪽 노드들과 후위 순회 결과의 왼쪽 노드들을 재귀 탐색한다.
5. 중위 순회 결과의 왼쪽 노드들과 후위 순회 결과의 오른쪽 노드들을 재귀 탐색한다.
*/

func rec(post, in []int) {
	if len(post) == 0 {
		return
	}

	tmp := post[len(post)-1]
	fmt.Fprint(writer, tmp, " ")

	post = post[:len(post)-1]

	idx := 0
	for i := 0; i < len(in); i++ {
		if in[i] == tmp {
			idx = i
			break
		}
	}

	inLeft := in[:idx]
	inRight := in[idx+1:]

	postLeft := post[:len(inLeft)]
	postRight := post[len(inLeft):]

	if len(inLeft) != 0 {
		rec(postLeft, inLeft)
	}

	if len(inRight) != 0 {
		rec(postRight, inRight)
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
