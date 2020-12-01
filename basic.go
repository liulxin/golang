package main

import "fmt"

// 包内变量
var aa int = 2
var bb string = "bb"
var (
	cc int = 4
	dd string = "dd"
)

func variableZeroValue() {
	var num int
	var str string
	fmt.Printf("%d %q \n",num, str)
}

func variableInitialValue() {
	var num1, num2 int = 3, 4
	var str string = "sxy"
	fmt.Println(num1, num2, str)
}

func variableTypeDeduction() {
	var num1, num2, str = 3, 4, "sxy"
	fmt.Println(num1, num2, str)
}

func variableShorter() {
	num1, num2, str := 3, 4, "sxy"
	fmt.Println(num1, num2, str)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(cc, bb)
}