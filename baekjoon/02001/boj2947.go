package bj2947

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

	list := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		list = append(list, n)
	}

	for !(list[0] == 1 && list[1] == 2 && list[2] == 3 && list[3] == 4 && list[4] == 5) {
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				printList(list)
			}
		}
	}
}

func printList(list []int) {
	for _, k := range list {
		fmt.Fprintf(writer, "%d ", k)
	}
	fmt.Fprintln(writer)
}
