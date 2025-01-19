package helper

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func MustConnect(clientTarget string) *grpc.ClientConn {
	conn, err := grpc.NewClient(clientTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("can't init grpc client %w", err)
	}
	return conn
}
