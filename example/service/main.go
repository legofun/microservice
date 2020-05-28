package main

import (
	pb "github.com/limitedlee/microservice/example/proto"
	"github.com/limitedlee/microservice/micro"
)

func main() {
	micro := &micro.MicService{}
	//micro.Route["/ping"]=handles.WsHandler
	//micro.Route["/"]
	micro.NewServer()
	pb.RegisterUserServiceServer(micro.GrpcServer, &pb.UserService{})
	micro.Start()
}
