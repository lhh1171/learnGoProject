package main

import (
	"micro/kit/library-user-service/dao"
	"micro/kit/library-user-service/endpoint"
	"micro/kit/library-user-service/service"
	"micro/kit/library-user-service/transport"
	"micro/kit/pkg/configs"
	"micro/kit/pkg/databases"

	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	pbbook "micro/kit/protos/book"

	"google.golang.org/grpc"
)

var confFile = flag.String("f", "library-user-service/user.yaml", "user config file")

func main() {
	flag.Parse()

	err := configs.Init(*confFile)
	if err != nil {
		panic(err)
	}

	err = databases.InitMySql(configs.Conf.MySQLConfig)
	if err != nil {
		fmt.Println("load mysql failed")
	}

	ctx := context.Background()

	conn, err := grpc.Dial("127.0.0.1:10088", grpc.WithInsecure())
	if err != nil {
		log.Println("连接user rpc 错误", err)
		panic(err)
	}
	defer conn.Close()
	bookClient := pbbook.NewBookClient(conn)

	userDao := dao.NewUserDaoImpl()
	userService := service.NewUserServiceImpl(userDao, bookClient)
	userEndpoints := &endpoint.UserEndpoints{
		RegisterEndpoint:          endpoint.MakeRegisterEndpoint(userService),
		FindByIDEndpoint:          endpoint.MakeFindByIDEndpoint(userService),
		FindByEmailEndpoint:       endpoint.MakeFindByEmailEndpoint(userService),
		FindBooksByUserIDEndpoint: endpoint.MakeFindBooksByUserIDEndpoint(userService),
	}

	ctx = context.WithValue(ctx, "ginMod", configs.Conf.ServerConfig.Mode)
	r := transport.NewHttpHandler(ctx, userEndpoints)

	errChan := make(chan error)
	go func() {
		errChan <- r.Run(fmt.Sprintf(":%s", strconv.Itoa(configs.Conf.ServerConfig.Port)))
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
}
