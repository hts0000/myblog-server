package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"net"
	"os"
	"sort"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

func main() {
	fmt.Println("hello/demo")
	ans := 100
	fmt.Println(max(ans, 1001))
	fmt.Println(min(ans, 1001))
	nums := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 77, 7, 33, 3, 1, 4, 5}
	sort.Slice(nums, func(i, j int) bool { return cmp.Less[int](nums[i], nums[j]) })
	slices.MaxFunc[[]int, int](nums, func(a, b int) int {
		if a > b {
			return a
		}
		return b
	})
	// 泛型自动推断类型
	slices.Max(nums)
	// 泛型指定类型参数
	slices.Sort[[]int, int](nums)

	endpointReader, endpointWriter := io.Pipe()
	defer endpointWriter.Close()

	// 同时往日志文件和控制台中输出
	fp, err := os.OpenFile("./hello/demo/log/info.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	w1 := bufio.NewWriter(fp)
	w2 := os.Stdout

	r1, pw := io.Pipe()
	r2 := io.TeeReader(endpointReader, pw)
	go func() {
		io.Copy(w1, r1)
	}()
	go func() {
		io.Copy(w2, r2)
	}()

	slog.SetDefault(slog.New(slog.NewTextHandler(endpointWriter, nil)))

	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)

	node1 := NewLinkNode1()
	node2 := NewLinkNode2()
	slog.Info("linknode:", node1, node1.Val, node1.Next)
	slog.Info("linknode:", node2, node2.Val, node2.Next)
}

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewLinkNode1() *LinkNode {
	return new(LinkNode)
}

func NewLinkNode2() *LinkNode {
	return &LinkNode{}
}
