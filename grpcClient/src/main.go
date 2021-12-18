package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-protobuf-server/grpcClient/src/api"
	"grpc-protobuf-server/grpcPb/src/pb"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var (
	clientCount  int   = 10
	requestCount int   = 10000
	timeout      int64 = 12 * 10
)

func NewClient(conn *grpc.ClientConn) *pb.TestserviceClient {
	testClient := pb.NewTestserviceClient(conn)
	return &testClient
}

func Working(conn *grpc.ClientConn, ctx *context.Context) {
	workCount := clientCount * requestCount * 4
	wg := &sync.WaitGroup{}
	wg.Add(workCount)
	for i := 1; i <= clientCount; i++ {
		client := NewClient(conn)
		clientId := fmt.Sprintf("Client-%v", i)
		for j := 1; j <= requestCount; j++ {
			requestId := fmt.Sprintf("id-%v", j)
			go api.Get(wg, ctx, *client, clientId, requestId)
			go api.Post(wg, ctx, *client, clientId, requestId)
			go api.Put(wg, ctx, *client, clientId, requestId)
			go api.Delete(wg, ctx, *client, clientId, requestId)
		}
		fmt.Printf("Working进度:%v/%v\n", (i * requestCount), workCount)
	}
	wg.Wait()
	fmt.Printf("All Work done :%v\n", workCount)
}

func Workflow(wg *sync.WaitGroup) {
	target := ":6688"
	//连接grpc服务端口
	conn, err := grpc.Dial(target, grpc.WithInsecure()) //grpc.WithInsecure()，禁用HTTP2.0的SSL安全传输
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer conn.Close()

	tm := int64(time.Second) * timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(tm))
	defer cancel()

	fmt.Printf("Connection Success %v\n", target)

	Working(conn, &ctx)
	defer wg.Done()
}

func main() {
	GetParameter()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Workflow(wg)
	// go Workflow(wg)
	wg.Wait()

	Readkey()
}

func GetParameter() {
	clientCount = *flag.Int("c", clientCount, "clientCount")
	requestCount = *flag.Int("r", requestCount, "requestCount")
	timeout = *flag.Int64("t", timeout, "timeout")
}

func Readkey() {
	fmt.Printf("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
