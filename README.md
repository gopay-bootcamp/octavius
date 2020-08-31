# OCTAVIUS
The new version of proctor, primarily used for automating tasks in kubernetes cluster. The architecture is inspired by Kubernetes in terms of Control Plane, Storage (single Etcd), API Server, and CLI Client. It is also inspired by Gitlab Runner in terms of Decentralized executor.


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

1. `git clone https://github.com/gopay-bootcamp/octavius.git`
2. `cd octavius`


## Running 

To create the binaries and protobuf files -:

`make build`

To run the server -:

`_output/bin/control_plane start`

To run the client -:

`_output/bin/cli <command> <args>`


## CONTRIBUTION

Contributions are welcomed! Please read the [contributing.md](./docs/contributing.md) before adding one.

## GUIDELINES

1. Refrain from using `github.com/gogo/protobuf` and instead use `github.com/golang/protobuf` as previous one is failing when marshalling proto messages from string.


