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
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}


## Development

1. `git clone https://github.com/gopay-bootcamp/octavius.git`
2. `cd octavius`


## Running 

To create the binaries and protobuf files -:

`make build`

To run the controller -:

`_output/bin/controller start`

To run the executor -:

`_output/bin/executor start`

To run the client -:

`_output/bin/cli <command> <args>`


## CONTRIBUTION

Contributions are welcomed! Please read the [contributing.md](./docs/contributing.md) before adding one.

## TROUBLESHOOT GUIDELINES

1. Refrain from using `github.com/gogo/protobuf` and instead use `github.com/golang/protobuf` as previous one is failing when marshalling proto messages from string.

2. gRPC version `>=1.30.x` has a name conflict with etcd. As a result it is better to stick to grpc `1.27.0` for the foreseeable future unless the upstream resolve their conflicts.


