package bj1259

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "0" {
			break
		}
		digits := strings.Split(s, "")
		isOk := true
		for i := 0; i < len(digits)/2; i++ {
			if digits[i] != digits[len(digits)-i-1] {
				isOk = false
				break
			}
		}
		if isOk {
			fmt.Fprintln(writer, "yes")
		} else {
			fmt.Fprintln(writer, "no")
		}
	}
}
