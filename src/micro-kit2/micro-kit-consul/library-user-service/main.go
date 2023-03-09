package main

import (
	"com/lhh/micro/kit/library-user-service/dao"
	"com/lhh/micro/kit/library-user-service/endpoint"
	"com/lhh/micro/kit/library-user-service/service"
	"com/lhh/micro/kit/library-user-service/transport"
	"com/lhh/micro/kit/pkg/configs"
	"com/lhh/micro/kit/pkg/databases"
	"com/lhh/micro/kit/pkg/ratelimits"
	"com/lhh/micro/kit/pkg/registers"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/juju/ratelimit"
)

var confFile = flag.String("f", "library-user-service/user.yaml", "user config file")

func main() {
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	err := configs.Init(*confFile)
	if err != nil {
		panic(err)
	}

	err = databases.InitMySql(configs.Conf.MySQLConfig)
	if err != nil {
		fmt.Println("load mysql failed")
	}

	ctx := context.Background()

	//生成grpc客户端和构建服务映射,基于
	findBooksEndpoint := endpoint.MakeFindBooksEndpoint
	//注册consul的调用
	grpcClient := registers.GRPCClient(configs.Conf, findBooksEndpoint, logger)

	//限流的令牌筒
	bucket := ratelimit.NewBucket(time.Second*time.Duration(configs.Conf.RatelimitConfig.FillInterval), int64(configs.Conf.RatelimitConfig.Capacity))
	//返回值是func
	limiter := ratelimits.NewTokenBucketLimiter(bucket)

	userDao := dao.NewUserDaoImpl()
	userService := service.NewUserServiceImpl(userDao, grpcClient)
	userEndpoints := &endpoint.UserEndpoints{
		RegisterEndpoint:          limiter(endpoint.MakeRegisterEndpoint(userService)),
		FindByIDEndpoint:          limiter(endpoint.MakeFindByIDEndpoint(userService)),
		FindByEmailEndpoint:       limiter(endpoint.MakeFindByEmailEndpoint(userService)),
		FindBooksByUserIDEndpoint: limiter(endpoint.MakeFindBooksByUserIDEndpoint(userService)),
		HealthEndpoint:            limiter(endpoint.MakeHealthEndpoint(userService)),
	}

	ctx = context.WithValue(ctx, "ginMod", configs.Conf.ServerConfig.Mode)
	r := transport.NewHttpHandler(ctx, userEndpoints)

	check := registers.HttpCheck(configs.Conf.ServerConfig.Port, configs.Conf.ConsulConfig.Interval, configs.Conf.ConsulConfig.Timeout)
	registrar := registers.InitRegister(configs.Conf, check, logger)

	errChan := make(chan error)
	go func() {
		registrar.Register()
		errChan <- r.Run(fmt.Sprintf(":%s", strconv.Itoa(configs.Conf.ServerConfig.Port)))
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
	registrar.Deregister()
}
