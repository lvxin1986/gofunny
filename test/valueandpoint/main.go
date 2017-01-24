package main

import (
	"fmt"
)

//简单的一个函数，实现了参数+1的操作
func add1(a int) int {
	a = a+1 // 我们改变了a的值
	return a //返回一个新值
}

func add2(a *int) int {
	*a = *a+1 // 我们改变了a的值
	return *a //返回一个新值
}



func ValueAndPointTest() {
	fmt.Println("begin test the value copy...")
	x1 := 3
	fmt.Println("x1 = ", x1)  // 应该输出 "x = 3"

	x11 := add1(x1)  //调用add1(x)

	fmt.Println("x1+1 = ", x11) // 应该输出"x+1 = 4"
	fmt.Println("x1 = ", x1)    // 应该输出"x = 3"

	fmt.Println("begin test the point copy...")
	x2 := 3
	fmt.Println("x2 = ", x2)  // 应该输出 "x = 3"

	x21 := add2(&x2)  //调用add1(x)

	fmt.Println("x2+1 = ", x21) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x2)    // 应该输出"x = 3"

}

func main() {
	ValueAndPointTest()
}