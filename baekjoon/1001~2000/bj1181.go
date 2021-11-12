package bj1181

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
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	words := make([]string, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		word := scanner.Text()
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			if strings.Compare(words[i], words[j]) == -1 {
				return true
			} else {
				return false
			}
		}
		return len(words[i]) < len(words[j])
	})

	fmt.Fprintln(writer, words[0])
	for i := 1; i < len(words); i++ {
		if words[i] == words[i-1] {
			continue
		}
		fmt.Fprintln(writer, words[i])
	}
}
