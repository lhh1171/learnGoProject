package main

import (
	"com/lhh/micro/kit/library-user-service/dao"
	"com/lhh/micro/kit/library-user-service/endpoint"
	"com/lhh/micro/kit/library-user-service/service"
	"com/lhh/micro/kit/library-user-service/transport"
	"com/lhh/micro/kit/pkg/circuitbreakers"
	"com/lhh/micro/kit/pkg/configs"
	"com/lhh/micro/kit/pkg/databases"
	"com/lhh/micro/kit/pkg/ratelimits"
	"com/lhh/micro/kit/pkg/registers"
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"github.com/juju/ratelimit"
)

var confFile = flag.String("f", "./user.yaml", "user config file")

// 请求完成之后才会有页面显示
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

	findBooksEndpoint := endpoint.MakeFindBooksEndpoint
	grpcClient := registers.GRPCClient(configs.Conf, findBooksEndpoint, logger)

	//创建hystrix熔断器进行配置
	hystrix.ConfigureCommand("Find books", hystrix.CommandConfig{Timeout: 10})
	//应该是嵌套一下hystrix
	grpcClient = circuitbreakers.Hystrix("Find books", "book-rpc-service currently unavailable")(grpcClient)

	bucket := ratelimit.NewBucket(time.Second*time.Duration(configs.Conf.RatelimitConfig.FillInterval), int64(configs.Conf.RatelimitConfig.Capacity))
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

	//启动hystrix，NewStreamHandler 返回一个服务器，该服务器能够通过 HTTP 公开仪表板指标
	hystrixStreamStreamHandler := hystrix.NewStreamHandler()

	//Start 开始观察内存中的断路器以获取指标
	hystrixStreamStreamHandler.Start()
	//接收hystrix-dashboard的请求并且返回
	go func() {
		fmt.Println("streamPort-----------", configs.Conf.StreamPort)
		errChan <- http.ListenAndServe(net.JoinHostPort("", configs.Conf.StreamPort), hystrixStreamStreamHandler)
		fmt.Println(errChan)
	}()

	go func() {
		//进行服务注册
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
