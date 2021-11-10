package bj10809

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanBytes)

	pos := make([]int, 26)
	for i := range pos {
		pos[i] = -1
	}
	cnt := 0

	for scanner.Scan() {
		s := scanner.Bytes()
		if s[0] == '\n' {
			break
		}
		if pos[s[0]-'a'] > -1 {
			cnt++
			continue
		}
		pos[s[0]-'a'] = cnt
		cnt++
	}
	for _, v := range pos {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}
