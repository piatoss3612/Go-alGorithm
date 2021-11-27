package bj4949

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
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Text() == "." {
			break
		}
		slice := strings.Split(scanner.Text(), "")
		checkBalanced(slice)
	}
}

func checkBalanced(slice []string) {
	matches := []string{}
	for _, v := range slice {
		currentLen := len(matches)
		if v == "(" || v == "[" {
			matches = append(matches, v)
		} else if v == ")" {
			if currentLen == 0 || matches[currentLen-1] != "(" {
				fmt.Fprintln(writer, "no")
				return
			} else {
				matches = matches[:currentLen-1]
			}
		} else if v == "]" {
			if currentLen == 0 || matches[currentLen-1] != "[" {
				fmt.Fprintln(writer, "no")
				return
			} else {
				matches = matches[:currentLen-1]
			}
		}
	}

	if len(matches) == 0 {
		fmt.Fprintln(writer, "yes")
	} else {
		fmt.Fprintln(writer, "no")
	}
}
