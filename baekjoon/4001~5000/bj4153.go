package bj4153

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Text() == "0 0 0" {
			break
		}
		slice := strings.Split(scanner.Text(), " ")
		sides := make([]int, 0, 3)
		for _, v := range slice {
			side, _ := strconv.Atoi(v)
			sides = append(sides, side)
		}
		sort.Ints(sides)
		a := sides[0] * sides[0]
		b := sides[1] * sides[1]
		c := sides[2] * sides[2]
		if (a + b) == c {
			fmt.Fprintln(writer, "right")
		} else {
			fmt.Fprintln(writer, "wrong")
		}
	}
}
