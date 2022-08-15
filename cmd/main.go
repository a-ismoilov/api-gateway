package main

import (
	"github.com/akbarshoh/microOLX/api"
	"github.com/akbarshoh/microOLX/api/handlers"
	"github.com/akbarshoh/microOLX/proto/orderproto"
	"github.com/akbarshoh/microOLX/proto/userproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	connOrder, err := grpc.Dial(":1111", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	connUser, err := grpc.Dial(":2222", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	api.Handle(handlers.New(orderproto.NewOrderServiceClient(connOrder), userproto.NewUserServiceClient(connUser)))
}
