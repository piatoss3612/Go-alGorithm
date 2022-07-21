package bj7785

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
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	logs := make(map[string]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		log := strings.Split(scanner.Text(), " ")
		logs[log[0]] = log[1]
		if log[1] == "leave" {
			delete(logs, log[0])
		}
	}

	names := make([]string, 0, len(logs))
	for name := range logs {
		names = append(names, name)
	}
	sort.Strings(names)

	for i := len(names) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%s\n", names[i])
	}
}
