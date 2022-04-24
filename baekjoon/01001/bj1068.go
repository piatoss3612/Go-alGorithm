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
	tree := make([]int, n)
	for i := 0; i < n; i++ {
		tree[i] = scanInt()
	}

	// 제거한 노드 및 연결된 모든 노드를 -2로 표시하여 제거한다
	rmv := []int{scanInt()}
	tree[rmv[0]] = -2
	for len(rmv) > 0 {
		front := rmv[0]
		rmv = rmv[1:]
		for i := 0; i < n; i++ {
			if tree[i] == front {
				tree[i] = -2
				rmv = append(rmv, i)
			}
		}
	}

	check := []int{}
	for i := 0; i < n; i++ {
		if tree[i] == -1 {
			check = append(check, i)
			break
		}
	}

	if len(check) == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 너비 우선 탐색
	// isLeaf 변수를 사용해 자식 노드가 있는지 확인하고
	// 없을 경우 ans를 1늘려준다
	ans := 0
	for len(check) > 0 {
		front := check[0]
		check = check[1:]
		isLeaf := true
		for i := 0; i < n; i++ {
			if tree[i] == front {
				isLeaf = false
				check = append(check, i)
			}
		}
		if isLeaf {
			ans += 1
		}
	}
	fmt.Fprintln(writer, ans)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
