package main

import "fmt"

func main() {
	mulThreeDigits()
}

func add() { //5단계 A+B (0 < A, B < 10)
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("a + b = %d", a+b)
}

func sub() { //6단계 A-B (0 < A, B < 10)
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("a - b = %d", a-b)
}

func mul() { //7단계 A*B (0 < A, B < 10)
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("a * b = %d", a*b)
}

func div() { //8단계 A/B (0 < A, B < 10)
	var a, b float64
	fmt.Scanf("%f %f", &a, &b)
	fmt.Printf("a / b = %f", a/b)
}

func mulThreeDigits() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	var hunds, tens, nums int
	hunds = (b / 100) * a
	tens = ((b % 100) / 10) * a
	nums = (b % 10) * a
	fmt.Println(nums)
	fmt.Println(tens)
	fmt.Println(hunds)
	fmt.Println(nums + (tens * 10) + (hunds * 100))
}
