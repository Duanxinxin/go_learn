// echo1 输出其命令行参数

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string // 声明两个 string 类型的变量 s 和 sep, 没有明确的进行初始化, 隐式的初始化为 空值, 此处为 ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*
for 循环是 go 中唯一的循环语句

for initialization; condition; post {
	//
}

三部分均可省略

传统的 while 循环

for condition {
}

三部分都不存在, 传统的 无限循环
for {
}

循环可以通过, break 或 return 终止
*/
