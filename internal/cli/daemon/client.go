package main

import (
	"context"
	"google.golang.org/grpc"
	"octavius/pkg/protobuf"

)

type Client interface {
	createJob() ()
}

type client struct {

}

func NewClient() Client {
	return &client{
		}
}

func (c *client) createJob() {
	conn, _ := grpc.Dial("localhost:8000",grpc.WithInsecure())
	grpcClient:=protobuf.NewProcServiceClient(conn)
	var args []*protobuf.Arg
	var sec []*protobuf.Secret
	args=append(args,&protobuf.Arg{Name: "name",Description: "name of proct"})
	sec=append(sec,&protobuf.Secret{Name: "secret",Description: "name of secret"})
	envVar:=&protobuf.EnvVars{Args: args,Secrets: sec}

	request:=protobuf.Metadata{EnvVars:envVar}
	grpcClient.CreateJob(context.Background(),&request)

}

func main() {
	client:=NewClient()
	client.createJob()
}