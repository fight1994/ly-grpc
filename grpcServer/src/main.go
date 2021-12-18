package main

import (
	"fmt"
	"grpc-protobuf-server/grpcPb/src/pb"
	"grpc-protobuf-server/grpcServer/src/server"
	"net"

	"google.golang.org/grpc"
)

func RegisterGrpcServer(grpcServer *grpc.Server) {
	//声明接口服务 * n
	testServer := server.NewTestServer()
	//把接口注册到grpc上 * n
	pb.RegisterTestserviceServer(grpcServer, testServer)

	fmt.Printf("Register success!\n")
}

func main() {
	//声明服务端口
	target := ":6688"
	tcp, err := net.Listen("tcp", target)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Listening %v\n", target)

	//声明grpc服务
	grpcServer := grpc.NewServer()

	//注册grpc服务
	RegisterGrpcServer(grpcServer)

	//启动grpc服务
	err = grpcServer.Serve(tcp)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
