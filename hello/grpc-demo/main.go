package main

import (
	"context"
	"fmt"
	hellopb "hello/grpc-demo/api/gen/v1"
	"net"

	"google.golang.org/grpc"
)

type Service struct {
	hellopb.UnimplementedHelloServer
}

func (s *Service) SayHello(ctx context.Context, req *hellopb.GetHelloReq) (rsp *hellopb.GetHelloRsp, err error) {
	fmt.Println("hello world")
	return &hellopb.GetHelloRsp{}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8018")
	s := grpc.NewServer()
	hellopb.RegisterHelloServer(s, &Service{})
	_ = s.Serve(lis)
}
