package server

import (
	"context"
	"grpc-protobuf-server/grpcPb/src/pb"
	"grpc-protobuf-server/grpcServer/src/client"
	"grpc-protobuf-server/grpcServer/src/util"

	"github.com/monnand/goredis"
)

var (
	CallCount   int64             = 0
	redisHelper *util.RedisHelper = &util.RedisHelper{Client: &goredis.Client{}}
)

func NewTestServer() pb.TestserviceServer {
	return &testServer{}
}

type testServer struct {
	testClient pb.TestserviceClient
}

func (s *testServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	s.testClient = client.NewTestClient()
	return s.testClient.Get(ctx, r)
}

func (s *testServer) Post(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	s.testClient = client.NewTestClient()
	return s.testClient.Post(ctx, r)
}

func (s *testServer) Put(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	s.testClient = client.NewTestClient()
	return s.testClient.Put(ctx, r)
}

func (s *testServer) Delete(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	s.testClient = client.NewTestClient()
	return s.testClient.Delete(ctx, r)
}
