package main

import "fmt"

func main() {
	add()
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

func calculator(op string) {
}
