package bj10610

import (
	"bufio"
	"fmt"
	"os"
)

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	defer writer.Flush()
	b, _ := reader.ReadBytes('\n')
	b = b[:len(b)-1]
	inputs := make([]int, 10)
	sum := 0
	hasZero := false
	for _, v := range b {
		n := int(v - '0')
		inputs[n] += 1
		sum += n
		if n == 0 {
			hasZero = true
		}
	}

	if hasZero && sum%3 == 0 {
		for i := 9; i >= 0; i-- {
			for j := 1; j <= inputs[i]; j++ {
				fmt.Fprint(writer, i)
			}
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, -1)
	}
}
