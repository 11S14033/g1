#==================================================Room#==================================================#
gen_room_pb:
	protoc \
	-I /usr/local/bin \
	-I. \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	--grpc-gateway_out=logtostderr=true:. \
	--go_out=plugins=grpc:. services/room/commons/protocs/Room.proto

gen_room_swagger:
	protoc \
	-I /usr/local/bin \
	-I. \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	--grpc-gateway_out=logtostderr=true:pb \
	--swagger_out=allow_merge=true,merge_file_name=greetapi:./swagger \
	--go_out=plugins=grpc:pkg services/room/commons/protocs/Room.proto


start_room_grpc_server:
	go run services/room/*.go -port 50051 -mode cli

start_room_grpc_gateway:
	go run services/room/*.go -port 8080 -mode grpc_gateway -grpcAddress 50051

