package bj1284

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	var result []int
	for r.Scan() {
		address := r.Text()
		if address != "0" {
			getWidth(address, &result)
		} else {
			break
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

func getWidth(s string, n *[]int) {
	width := 1
	for _, v := range s {
		if v == '1' {
			width += 2
		} else if v == '0' {
			width += 4
		} else {
			width += 3
		}
		width++
	}
	*n = append(*n, width)
}
