package bj1371

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	counts := make(map[string]int)
	r := bufio.NewReader(os.Stdin)
	s, _ := ioutil.ReadAll(r)
	file := bytes.NewBuffer(s).String()
	lines := strings.Split(file, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			letters := strings.Split(word, "")
			for _, letter := range letters {
				_, found := counts[letter]
				if found {
					counts[letter]++
					continue
				}
				counts[letter] = 1
			}
		}
	}
	k, m := getMax(&counts)
	result := getAllMax(&counts, k, m)
	sort.Strings(result)
	for _, s := range result {
		fmt.Print(s)
	}
	fmt.Println()
}

func getMax(counts *map[string]int) (string, int) {
	var key string
	var max int
	for k, v := range *counts {
		if v > max {
			max = v
			key = k
		}
	}
	return key, max
}

func getAllMax(counts *map[string]int, key string, max int) []string {
	var result []string
	for k, v := range *counts {
		if v == max {
			result = append(result, k)
		}
	}
	return result
}
