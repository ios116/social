mock_user:
	mockgen  -destination=./internal/domain/usecase/mock_repo.go  -package=usecase -source=./internal/domain/entities/users.go UserRepository

rpc_impl:
	touch internal/grpcserver/fee.go; echo package grpcserver > internal/grpcserver/fee.go; impl 's *RPCServer' social/internal/grpcserver.UsersServer >> internal/grpcserver/fee.go

