package main

import (
	"github.com/juju/ratelimit"
	"micro/kit/library-user-service/dao"
	"micro/kit/library-user-service/endpoint"
	"micro/kit/library-user-service/service"
	"micro/kit/library-user-service/transport"
	"micro/kit/pkg/configs"
	"micro/kit/pkg/databases"
	"micro/kit/pkg/ratelimits"
	"time"

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
	//Parse 从 os 解析命令行标志。
	//参数[1:]。必须在定义所有标志之后和程序访问标志之前调用
	//cmd
	//go run main.go -f ./user.yaml
	flag.Parse()

	//读取yaml并转换成对象
	err := configs.Init(*confFile)
	if err != nil {
		panic(err)
	}
	//连接数据库了
	err = databases.InitMySql(configs.Conf.MySQLConfig)
	if err != nil {
		fmt.Println("load mysql failed")
	}

	//Background 返回一个非零的空Context 。
	//它永远不会被取消，没有价值，也没有截止日期。
	//它通常由 main 函数、初始化和测试使用，并作为传入请求的顶级上下文
	ctx := context.Background()

	//Dial 创建到给定目标的客户端连接
	conn, err := grpc.Dial("127.0.0.1:10088", grpc.WithInsecure())
	if err != nil {
		log.Println("连接user grpc 错误", err)
		panic(err)
	}
	defer conn.Close()

	//grpClient的初始化
	bookClient := pbbook.NewBookClient(conn)
	log.Println(configs.Conf.FillInterval, configs.Conf.RatelimitConfig.Capacity)
	// 创建令牌桶
	bucket := ratelimit.NewBucket(time.Second*time.Duration(configs.Conf.FillInterval), int64(configs.Conf.RatelimitConfig.Capacity))
	limiter := ratelimits.NewTokenBucketLimiter(bucket)

	//启用服务的对象
	userDao := dao.NewUserDaoImpl()
	userService := service.NewUserServiceImpl(userDao, bookClient)
	userEndpoints := &endpoint.UserEndpoints{
		RegisterEndpoint:          limiter(endpoint.MakeRegisterEndpoint(userService)),
		FindByIDEndpoint:          limiter(endpoint.MakeFindByIDEndpoint(userService)),
		FindByEmailEndpoint:       limiter(endpoint.MakeFindByEmailEndpoint(userService)),
		FindBooksByUserIDEndpoint: limiter(endpoint.MakeFindBooksByUserIDEndpoint(userService)),
	}

	//配置文件，kv装入context
	ctx = context.WithValue(ctx, "ginMod", configs.Conf.ServerConfig.Mode)
	//启用网络服务
	r := transport.NewHttpHandler(ctx, userEndpoints)

	//发生错误和接收到ctl+c这种管道请求可以打印并处理
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
