package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hellow World")

	// 整数型の変数
	var num1 int = 123
	fmt.Println(num1)

	// 右辺で型が決まるなら型名は不要
	var num2 = 123
	fmt.Println(num2)

	if true {
		x := 2
		fmt.Println("X=", x)
	}

	// 単行コメント　
	/*
		複数行コメント
	*/

}
