export PATH=$PATH:~/go/bin
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative  ./internal/pkg/protofiles/*.proto