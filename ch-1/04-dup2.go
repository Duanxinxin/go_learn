// @Time    : 2023/12/1 14:35
// @File    : 04-dup2.go
// @User    : Recci
// @Software: GoLand

// dup2 打印输入中多次数显行的个数和文本
// 是 stdin 或执行的文本列表读取

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // 打开文件
			if err != nil {
				_, err = fmt.Fprintf(os.Stderr, "dup2: %v\n", err) // 格式化输出, 输出到指定位置, 这里输出为标准输出
				continue
			}
			countLines(f, counts)
			_ = f.Close()
		}
	}
}
