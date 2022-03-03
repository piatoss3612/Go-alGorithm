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
	graph   [][]int
	n       int
	visited []bool
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	graph = make([][]int, n+1)
	visited = make([]bool, n+1)

	for i := 1; i <= n; i++ {
		graph[byteToInt()] = []int{byteToInt(), byteToInt()}
	}

	preorder(1)
	fmt.Fprintln(writer)
	visited = make([]bool, n+1)

	inorder(1)
	fmt.Fprintln(writer)
	visited = make([]bool, n+1)

	postorder(1)
	fmt.Fprintln(writer)
}

// 전위 순위: 깊이 우선 탐색과 유사
func preorder(idx int) {
	visited[idx] = true
	fmt.Fprint(writer, string(idx+64))

	if graph[idx][0] != 0 && !visited[graph[idx][0]] {
		preorder(graph[idx][0])
	}

	if graph[idx][1] != 0 && !visited[graph[idx][1]] {
		preorder(graph[idx][1])
	}
}

// 중위 탐색: 왼쪽 자식부터 부모, 오른쪽 자식으로 이동하면 다시 왼쪽 자식부터 부모를 반복
func inorder(idx int) {
	if graph[idx][0] != 0 && !visited[graph[idx][0]] {
		inorder(graph[idx][0])
		fmt.Fprint(writer, string(idx+64))
	} else {
		fmt.Fprint(writer, string(idx+64))
	}

	if graph[idx][1] != 0 && !visited[graph[idx][1]] {
		inorder(graph[idx][1])
	}
}

// 후위 탐색: 왼쪽 깊이를 먼저 탐색하고 오른쪽을 탐색하는 것을 반복한 후에 자기 자신 출력
func postorder(idx int) {
	if graph[idx][0] != 0 && !visited[graph[idx][0]] {
		postorder(graph[idx][0])
	}

	if graph[idx][1] != 0 && !visited[graph[idx][1]] {
		postorder(graph[idx][1])
	}

	fmt.Fprint(writer, string(idx+64))
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanByte() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}

func byteToInt() int {
	b := scanByte()
	if b == '.' {
		return 0
	}
	return int(b - 64)
}
