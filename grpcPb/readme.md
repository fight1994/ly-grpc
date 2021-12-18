环境：
1、安装protoc 并配置Path环境变量
https://github.com/protocolbuffers/protobuf/releases

2、添加protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go


生成命令：
cd .\grpcPb\src\proto\
protoc --go_out=plugins=grpc:. .\*.proto