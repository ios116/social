gen:
	protoc -I/usr/local/include -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. api/*.proto

gwgen:
	protoc -I/usr/local/include -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:.   api/api.proto

build: gen
	go build -o clncnd main.go

swg:
	protoc -I/usr/local/include -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. api/api.proto

mock_user:
	mockgen  -destination=./internal/domain/usecase/mock_repo.go  -package=usecase -source=./internal/domain/entities/users.go UserRepository

rpc_impl:
	touch internal/grpcserver/fee.go; echo package grpcserver > internal/grpcserver/fee.go; impl 's *RPCServer' social/internal/grpcserver.UsersServer >> internal/grpcserver/fee.go

