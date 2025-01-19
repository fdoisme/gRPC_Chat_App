package main

import (
	"fmt"
	"log"
	"net"
	"project/chat-service/handler"
	"project/chat-service/infra"

	pb "project/chat-service/proto"

	"google.golang.org/grpc"
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
	pb.RegisterChatServiceServer(server, &handler.ChatHandler{
		Service: ctx.Ctl.ChatHandler.Service,
		Logger:  ctx.Log,
	})

	log.Printf("chat-service started %s:%s", ctx.Cfg.GrpcIp, ctx.Cfg.GrpcPort)
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
