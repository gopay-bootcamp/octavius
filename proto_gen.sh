export PATH=$PATH:~/go/bin
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative  ./internal/pkg/protofiles/client_cp/*.proto
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative  ./internal/pkg/protofiles/executor_cp/*.proto