package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"project/auth-service/infra"
	pb "project/auth-service/proto"
)

func main() {
	var err error

	var ctx *infra.ServiceContext
	if ctx, err = infra.NewServiceContext(); err != nil {
		log.Fatal("can't init service context %w", err)
	}

	var listener net.Listener
	if listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", ctx.Cfg.GrpcIp, ctx.Cfg.GrpcPort)); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, ctx.Svc.Auth)

	log.Printf("auth-service started %s:%s", ctx.Cfg.GrpcIp, ctx.Cfg.GrpcPort)
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
