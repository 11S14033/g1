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
* gorilla mux
## Prerequisites
See Makefile for

Generate file pb
```
make gen_room_pb
```

To generate swagger file:
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

### Run gRPC gateway server
```
start_room_grpc_gateway_server:
	go run services/room/*.go -port 8080 -mode grpc_gateway -grpcAddress 50051
```
To start gRPC gateway server:
```
make start_room_grpc_gateway_server
```

### Run rest server
```
start_room_rest_server:
	go run services/room/*.go -port 8080 -mode grpc_gateway -grpcAddress 50051
```
To start rest server:
```
make start_room_rest_server
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

