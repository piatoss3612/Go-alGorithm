package bj1764

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	check := make(map[string]int)
	cnt := 0
	s := []string{}
	for i := 0; i < n+m; i++ {
		scanner.Scan()
		name := scanner.Text()
		if i < n {
			check[name] = 1
		} else {
			_, ok := check[name]
			if ok {
				cnt++
				s = append(s, name)
			}
		}
	}

	sort.Strings(s)

	fmt.Fprintln(writer, cnt)
	for _, v := range s {
		fmt.Fprintln(writer, v)
	}
}
