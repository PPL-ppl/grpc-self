package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hb "grpc-self/hello-server/proto"
	"net"
)

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", ":8080")
	//创建grpc服务
	newServer := grpc.NewServer()
	//注册服务
	hb.RegisterSayHelloServer(newServer, &server{})
	//启动服务
	err := newServer.Serve(listen)
	if err != nil {
		fmt.Println("启动失败")
		return
	}
}

type server struct {
	hb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *hb.HelloRequest) (*hb.HelloResponse, error) {
	return &hb.HelloResponse{RequestName: req.RequestName}, nil
}
