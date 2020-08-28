export PATH=$PATH:~/go/bin
protoc --go_out=. --go_opt=paths=source_relative  ./internal/pkg/protofiles/executor_cp/*.proto
protoc --go-grpc_out=. --go-grpc_opt=requireUnimplementedServers=false,paths=source_relative  ./internal/pkg/protofiles/executor_cp/*.proto