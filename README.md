# G1

## Description
This is a fun project to imporove my knowledge about GRPC. This project maintain room processing business API. There are three mode delivery communication services:
* CLI
* gRPC Gateway
* Rest API

Here component in this project:
* Golang 
* Gorm (still using postgre sql)
* gRPC (using libary evans for running using cli)
* gRPC Gateway (Define end point service)

## Prerequisites
See Makefile to:

### Generate [filename].pb.go and [filename].pb.gw.go From Proto File
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

###  Generate Swagger File For Documentation API.
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
## Running App
### Run gRPC server
```
start_room_grpc_server:
	go run services/room/*.go -port 50051 -mode cli
```
To start gRPC server:
```
make start_room_grpc_server
```

### Run gRPC gateway
```
start_room_grpc_gateway:
	go run services/room/*.go -port 8080 -mode grpc_gateway -grpcAddress 50051
```
To start gRPC gateway:
```
make start_room_grpc_gateway
```

## End Point
See file [filename].proto to see end point dan file swagger to see documentation API.

Room Services
```
post:"/v1/room"
get:"/v1/rooms"
get:"/v1/room/{room_id}"
put:"/v1/room"
delete:"/v1/room/{room_id}"
```

## Tools Used
* All library can be seen in file go.mod

## Reference
* https://github.com/bxcodec/go-clean-arch

## Next Step
This project is still not finished and still doing:
* Unit testing
* Solve some isue
* Build delivery rest api
* Docker for deploymet
* Develop User service

