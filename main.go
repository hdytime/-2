package main

import (
	"fmt"
)

func Calculator(num1 int, num2 int, CMD func(int, int) int) int {
	return CMD(num1, num2)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func main() {
	a := 0
	b := 0
	result := 0
	fmt.Println("请输入任意两个整数以进行运算")
	fmt.Scan(&a, &b)
	fmt.Println("请选择运算方式\nadd\nsubtract\nmultiply\ndivide")
	var cmd string = "0"
	fmt.Scan(&cmd)
	switch {
	case cmd == "add":
		result = Calculator(a, b, add)
		fmt.Println("结果为:", result)

	case cmd == "subtract":
		result = Calculator(a, b, subtract)
		fmt.Println("结果为:", result)

	case cmd == "multiply":
		result = Calculator(a, b, multiply)
		fmt.Println("结果为:", result)

	case cmd == "divide":
		result = Calculator(a, b, divide)
		fmt.Println("结果为:", result)
	default:
		fmt.Println("请检查运算方式是否输入正确")
	}
}
