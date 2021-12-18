package api

import (
	"context"
	"fmt"
	"grpc-protobuf-server/grpcClient/src/util"
	"grpc-protobuf-server/grpcPb/src/pb"
	"sync"

	"github.com/monnand/goredis"
)

var (
	sendRecord  string            = "sendRecord"
	redisHelper *util.RedisHelper = &util.RedisHelper{Client: &goredis.Client{}}
)

func Get(wg *sync.WaitGroup, ctx *context.Context, client pb.TestserviceClient, clientId string, requestId string) {
	resp, err := client.Get(*ctx, &pb.GetRequest{ClientId: clientId, Id: requestId})
	defer wg.Done()
	if err != nil || resp == nil {
		fmt.Printf("接口请求失败，Error: %v\n", err)
		return
	}
	UpdateRecord()
	fmt.Printf("resp: %v\n", resp)

}

func Post(wg *sync.WaitGroup, ctx *context.Context, client pb.TestserviceClient, clientId string, requestId string) {
	resp, err := client.Post(*ctx, &pb.GetRequest{ClientId: clientId, Id: requestId})
	defer wg.Done()
	if err != nil {
		fmt.Printf("接口请求失败，Error: %v\n", err)
		return
	}
	UpdateRecord()
	fmt.Printf("resp: %v\n", resp)
}

func Put(wg *sync.WaitGroup, ctx *context.Context, client pb.TestserviceClient, clientId string, requestId string) {
	resp, err := client.Put(*ctx, &pb.GetRequest{ClientId: clientId, Id: requestId})
	defer wg.Done()
	if err != nil {
		fmt.Printf("接口请求失败，Error: %v\n", err)
		return
	}
	UpdateRecord()
	fmt.Printf("resp: %v\n", resp)
}

func Delete(wg *sync.WaitGroup, ctx *context.Context, client pb.TestserviceClient, clientId string, requestId string) {
	resp, err := client.Delete(*ctx, &pb.GetRequest{ClientId: clientId, Id: requestId})
	defer wg.Done()
	if err != nil {
		fmt.Printf("接口请求失败，Error: %v\n", err)
		return
	}
	UpdateRecord()
	fmt.Printf("resp: %v\n", resp)
}

func UpdateRecord() {
	redisHelper.Client.Incr(sendRecord)
}
