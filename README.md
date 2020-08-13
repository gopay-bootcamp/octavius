# OCTAVIUS
The new version of proctor, primarily used for automating tasks in kubernetes cluster.


## Prerequisites

1. golang
2. Docker
3. Setting up etcd container
### installation
    docker run -d --name etcd-server \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
4. protobuf
### installation
    brew install protobuf
5. protoc-gen-go
### installation
    go get -u google.golang.org/protobuf/cmd/protoc-gen-go

    go install google.golang.org/protobuf/cmd/protoc-gen-go

6. protoc-gen-go-grpc
### installation
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

## Development

1. `git clone <url>`
2. `cd octavius`
3. `go mod tidy`


## Running 

Build the proto file -:

`protoc --go_out=. --go_opt=paths=source_relative  ./pkg/protobuf/process.proto`

`protoc --go-grpc_out=. --go-grpc_opt=requireUnimplementedServers=false,paths=source_relative  ./pkg/protobuf/process.proto` 


To create the binaries -:

`make build`


To run the server -:

`_output/bin/server start`

To run the client -:

`_output/bin/cli <command> <args>`



