gen-grpc:
	protoc  \
		--go_out=./internal/service/grpc/ \
		--go_opt=module=github.com/victor8titov/go-grpc-example  \
		--go-grpc_out=./internal/service/grpc \
		--go-grpc_opt=module=github.com/victor8titov/go-grpc-example  \
		./proto/base.proto