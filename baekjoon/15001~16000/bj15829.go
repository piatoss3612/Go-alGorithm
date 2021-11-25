package bj15829

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanBytes)
	s := ""
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			break
		}
		s += scanner.Text()
	}
	l, _ := strconv.Atoi(s)

	var hash int64 = 0
	var r int64 = 1
	for i := 0; i < l; i++ {
		scanner.Scan()
		hash = (hash + int64(scanner.Bytes()[0]-96)*r) % 1234567891
		r = (r * 31) % 1234567891
	}
	fmt.Fprintln(writer, hash)
}
