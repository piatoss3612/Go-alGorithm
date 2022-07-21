package bj2675

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		test := strings.Split(scanner.Text(), " ")
		rep, _ := strconv.Atoi(test[0])
		str := strings.Split(test[1], "")
		getQR(rep, str)
	}
}

func getQR(rep int, str []string) {
	for _, v := range str {
		for i := 0; i < rep; i++ {
			writer.WriteString(v)
		}
	}
	writer.WriteByte('\n')
}
