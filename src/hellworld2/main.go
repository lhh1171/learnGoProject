package main

import (
	"hellworld2/handler"
	pb "hellworld2/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("hellworld2"),
	)

	// Register handler
	pb.RegisterHellworld2Handler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
