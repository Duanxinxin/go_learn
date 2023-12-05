// @Time    : 2023/12/4 20:35
// @File    : 07-fetchAll.go
// @User    : Recci
// @Software: GoLand

// 并发获取URL, 并报告它们的时间和大小

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintln(err) // 发送到通道 ch
		return
	}
	nBytes, err := io.Copy(io.Discard, resp.Body) // 读取响应内容, 通过写入 Discard 输出流进行丢弃, Copy 返回字节数以及所有错误
	cErr := resp.Body.Close()
	if cErr != nil {
		return
	} // 关闭避免资源泄露
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // 写入通道
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nBytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string) // 创建一个字符串通道
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 启动一个goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道 ch 接收
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

/*
goroutine:  是一个并发执行函数.
chan:
	是用于在 goroutine 之间进行通信的一种数据类型。
	chan 是 channel 的缩写，它提供了一种在不同 goroutine 之间传递数据的方式，实现了同步和通信。

	通道是一种允许某一个协程向另一个协程传递指定类型的值的通信机制

当一个 goroutine 试图在一个通道上进行发送或接收操作时, 它会阻塞, 直到另一个goroutine试图进行接收或发送操作才传递值, 并开始处理两个 goroutine
*/
