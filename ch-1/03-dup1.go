// dup1 输出标准输入中出现次数大于1的行, 前面是次数

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // map 提供键值对集合, 提供常量时间操作存储或获取某个元素. make 函数新建 map
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意: 忽略 input.Err() 中可能得错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d  %s\n", n, line)
		}
	}
}

// TODO: 没有看到输出, 这是为啥
