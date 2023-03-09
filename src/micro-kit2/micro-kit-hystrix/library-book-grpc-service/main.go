package main

import (
	"com/lhh/micro/kit/library-book-grpc-service/dao"
	"com/lhh/micro/kit/library-book-grpc-service/endpoint"
	"com/lhh/micro/kit/library-book-grpc-service/service"
	"com/lhh/micro/kit/library-book-grpc-service/transport"
	"com/lhh/micro/kit/pkg/configs"
	"com/lhh/micro/kit/pkg/databases"
	"com/lhh/micro/kit/pkg/ratelimits"
	"com/lhh/micro/kit/pkg/registers"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	pbbook "com/lhh/micro/kit/protos/book"

	"github.com/go-kit/kit/log"
	"github.com/hashicorp/consul/api"
	"github.com/juju/ratelimit"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "library-book-grpc-service/book_rpc.yaml", "book rpc config file")
var quiteChan = make(chan error, 1)

func main() {
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	err := configs.Init(*configFile)
	if err != nil {
		panic(err)
	}
	err = databases.InitMySql(configs.Conf.MySQLConfig)
	if err != nil {
		fmt.Println("load mysql failed")
	}

	ctx := context.Background()

	bucket := ratelimit.NewBucket(time.Second*time.Duration(configs.Conf.RatelimitConfig.FillInterval), int64(configs.Conf.RatelimitConfig.Capacity))
	ratelimit := ratelimits.NewTokenBucketLimiter(bucket)

	bookDao := dao.NewBookDaoImpl()
	bookService := service.NewBookServiceImpl(bookDao)
	endpoints := endpoint.BookEndpoints{
		FindBooksByUserIDEndpoint: ratelimit(endpoint.NewFindBooksByUserIDEndpoint(bookService)),
	}

	// check := registers.GRPCCheck(configs.Conf.ServerConfig.Port, configs.Conf.ConsulConfig.Interval, configs.Conf.ConsulConfig.Timeout)
	var check api.AgentServiceCheck
	registrar := registers.InitRegister(configs.Conf, check, logger)

	go func() {
		registrar.Register()
		handler := transport.NewBookServer(ctx, endpoints)
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", strconv.Itoa(configs.Conf.ServerConfig.Port)))
		if err != nil {
			fmt.Println("listen tcp err", err)
			quiteChan <- err
			return
		}
		gRPCServer := grpc.NewServer()
		pbbook.RegisterBookServer(gRPCServer, handler)
		err = gRPCServer.Serve(listener)
		if err != nil {
			fmt.Println("gRPCServer Serve err", err)
			quiteChan <- err
			return
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		quiteChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-quiteChan)
	registrar.Deregister()
}
