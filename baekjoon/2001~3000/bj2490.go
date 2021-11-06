package bj2490

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	kGame := []string{"E", "A", "B", "C", "D"}
	scores := []int{}
	scanner.Split(bufio.ScanLines)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		s := scanner.Text()
		cnt := strings.Count(s, "0")
		scores = append(scores, cnt)
	}

	for _, v := range scores {
		fmt.Println(kGame[v])
	}
}
