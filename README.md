# G1

## Description
This is a fun project to imporove my knowledge about GRPC.


## Prerequisites
* Golang
* Using GRPC
* Using GRPC Gateway
* Using Gorm


## Running App
* See Makefile to:
** Generate [filename].pb.go and [filename].pb.gw.go from proto file
```
gen_room_pb:
	protoc \
	-I /usr/local/bin \
	-I. \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	--grpc-gateway_out=logtostderr=true:. \
	--go_out=plugins=grpc:. services/room/commons/protocs/Room.proto
```
To generate file:
```
make gen_room_pb
```

** Generate swagger file for documentation API.
```
gen_room_swagger:
	protoc \
	-I /usr/local/bin \
	-I. \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	--grpc-gateway_out=logtostderr=true:pb \
	--swagger_out=allow_merge=true,merge_file_name=roomapi:./swagger \
	--go_out=plugins=grpc:pkg services/room/commons/protocs/Room.proto
```
To generate file:
```
make gen_room_swagger
```

** Running GRPC server
```
start_room_grpc_server:
	go run services/room/*.go -port 50051 -mode cli
```
To generate file:
```
make start_room_grpc_server
```

** Running GRPC gateway
```
start_room_grpc_gateway:
	go run services/room/*.go -port 8080 -mode grpc_gateway -grpcAddress 50051
```
To generate file:
```
make start_room_grpc_gateway
```

## Tools Used
* All library can be seen in file go.mod

## Reference
* https://github.com/bxcodec/go-clean-arch

