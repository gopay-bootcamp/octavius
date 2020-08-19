export PATH=$PATH:~/go/bin
protoc --go_out=. --go_opt=paths=source_relative  ./pkg/protobuf/process.proto
protoc --go-grpc_out=. --go-grpc_opt=requireUnimplementedServers=false,paths=source_relative  ./pkg/protobuf/process.proto