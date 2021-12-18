package client

import (
	"context"
	"fmt"
	"grpc-protobuf-server/grpcPb/src/pb"
	"grpc-protobuf-server/grpcServer/src/util"
	"math/rand"
	"strconv"
	"time"

	"github.com/monnand/goredis"
	"google.golang.org/grpc"
)

var (
	baseSecond    float64           = 1
	receiveRecord string            = "receiveRecord"
	redisHelper   *util.RedisHelper = &util.RedisHelper{Client: &goredis.Client{}}
)

type testClient struct {
}

func NewTestClient() pb.TestserviceClient {
	return &testClient{}
}

func (c *testClient) Get(ctx context.Context, r *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	s := sleep(baseSecond)
	fmt.Printf("%v Get-request-id: %v （%v s）\n", r.ClientId, r.Id, s)
	res := &pb.GetResponse{}
	res.Result = r.Id
	defer UpdateRecord()
	return res, nil
}

func (c *testClient) Post(ctx context.Context, r *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	s := sleep(baseSecond)
	fmt.Printf("%v Post-request-id: %v （%v s）\n", r.ClientId, r.Id, s)
	res := &pb.GetResponse{}
	res.Result = r.Id
	defer UpdateRecord()
	return res, nil
}

func (c *testClient) Put(ctx context.Context, r *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	s := sleep(baseSecond)
	fmt.Printf("%v Put-request-id: %v （%v s）\n", r.ClientId, r.Id, s)
	res := &pb.GetResponse{}
	res.Result = r.Id
	defer UpdateRecord()
	return res, nil
}

func (c *testClient) Delete(ctx context.Context, r *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	s := sleep(baseSecond)
	fmt.Printf("%v Delete-request-id: %v （%v s）\n", r.ClientId, r.Id, s)
	res := &pb.GetResponse{}
	res.Result = r.Id
	defer UpdateRecord()
	return res, nil
}

//模拟耗时操作（基础时间 + 随机时间）
func sleep(baseSecond float64) float64 {
	random := float64(float64(rand.Intn(100)) / 100)
	result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", baseSecond+random), 64)
	timeData := int64(result * 1000000000)
	time.Sleep(time.Duration(timeData))
	return result
}

//计数 总共多少请求 记到redis里
func UpdateRecord() {
	redisHelper.ConnRedis()
	redisHelper.Client.Incr(receiveRecord)
}
