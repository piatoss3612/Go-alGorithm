package bj1100

import (
	"fmt"
	"strings"
)

func main() {
	var chess [8][8]string
	var lines [8]string
	for i, _ := range lines {
		fmt.Scan(&lines[i])
	}

	for i := 0; i < len(chess); i++ {
		slice := strings.Split(lines[i], "")
		copy(chess[i][:], slice)
	}

	sw := true
	cnt := 0

	for i := 0; i < len(chess); i++ {
		line := chess[i]
		for j := 0; j < len(line); j++ {
			if sw && line[j] == "F" {
				cnt++
			}
			sw = !sw
		}
		sw = !sw
	}

	fmt.Print(cnt)
}
