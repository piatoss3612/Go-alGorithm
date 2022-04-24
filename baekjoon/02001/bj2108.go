package bj2108

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "sort"
)

var (
  scanner = bufio.NewScanner(os.Stdin)
  writer = bufio.NewWriter(os.Stdout)
)

func main() {
  defer writer.Flush()
  scanner.Split(bufio.ScanWords)

  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())

  nums := make([]int, 0, n)

  for i := 0; i < n; i++ {
    scanner.Scan()
    num, _ := strconv.Atoi(scanner.Text())
    nums = append(nums, num)
  }
  sort.Ints(nums)
  fmt.Fprintln(writer, getAvg(nums))
  fmt.Fprintln(writer, getMedian(nums))
  fmt.Fprintln(writer, getSecondMode(nums))
  fmt.Fprintln(writer, getRange(nums))
 
}

func getAvg(nums []int) int {
  sum := 0
  for _, v := range nums {
    sum += v
  }
  if sum % len(nums) > len(nums) / 2 {
    return (sum / len(nums)) + 1
  } else if sum % len(nums) < -(len(nums) / 2) {
    return (sum / len(nums)) - 1
  }
  return sum / len(nums)
}

func getMedian(nums []int) int {
  return nums[(len(nums) - 1) / 2]
}

func getSecondMode(nums []int) int {
  counts := make(map[int]int)
  for _, v := range nums {
    _, ok := counts[v]
    if ok {
      counts[v] += 1
    } else {
      counts[v] = 1
    }
  }
  max := 0
  for _, v := range counts {
    if v > max {
      max = v
    }
  }
  same := []int{}
  for k, v := range counts {
    if v == max {
      same = append(same, k)
    }
  }
  sort.Ints(same)
  if len(same) == 1 {
    return same[0]
  }
  return same[1]
}

func getRange(nums []int) int {
  if len(nums) == 1 {
    return 0
  }
  return nums[len(nums) - 1] - nums[0]
}
