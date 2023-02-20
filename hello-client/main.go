package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-self/hello-server/proto"
	"log"
)

func main() {
	//连接server 此处没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	//建立连接
	client := pb.NewSayHelloClient(conn)
	//执行rpc调用(这个服务器端来实现并返回结果)
	hello, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "hello wzl"})
	fmt.Println(hello)
}
