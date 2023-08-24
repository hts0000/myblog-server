package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello/myio")
	r := strings.NewReader("hello world")
	w1 := os.Stdout
	fp, err := os.OpenFile("./hello/io/log/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	w2 := bufio.NewWriter(fp)
	ConcurrencyWrtie(r, [2]io.Writer{w1, w2})
}

func ConcurrencyWrtie(src io.Reader, dest [2]io.Writer) (err error) {
	errCh := make(chan error, 1)

	// 管道，主要是用来写、读流转化
	pr, pw := io.Pipe()
	// teeReader ，主要是用来 IO 流分叉
	wr := io.TeeReader(src, pw)

	// 并发写入
	go func() {
		var _err error
		defer func() {
			pr.CloseWithError(_err)
			errCh <- _err
		}()
		_, _err = io.Copy(dest[1], pr)
	}()

	defer func() {
		// TODO：异常处理
		pw.Close()
		_err := <-errCh
		_ = _err
	}()

	// 数据写入
	_, err = io.Copy(dest[0], wr)

	return err
}
