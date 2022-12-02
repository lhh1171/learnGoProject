package myapi

import (
	"context"
	proto "github.com/micro/go-micro/v2/api/proto"
	log "github.com/micro/go-micro/v2/logger"
)

// Event
// 切记，事件订阅结构的所有公有方法都会被执行 方法名没有限制，但是方法一定要接收ctx，event
type Event struct{}

func (e *Event) Process(ctx context.Context, event *proto.Event) error {
	log.Info("公有方法Process 收到事件，", event.Name)
	log.Info("公有方法Process 数据", event.Data)
	return nil
}

func (e *Event) Process2(ctx context.Context, event *proto.Event) error {
	log.Info("公有方法Process2 收到事件，", event.Name)
	log.Info("公有方法Process2 数据，", event.Data)
	return nil
}

/** 打开本注释后，会导致侦听器无法工作
func (e *Event) Process3() error {
    log.Log("公有方法Process3 收到事件3，不解析参数")
    return nil
}**/

func (e *Event) process(ctx context.Context, event *proto.Event) error {
	log.Info("私有方法process，收到事件，", event.Name)
	log.Info("私有方法process，数据", event.Data)
	return nil
}

func II2() {
	service := micro.NewService(
		// 服务名可以随意
		micro.Name("user2"),
	)
	service.Init()

	// register subscriber
	micro.RegisterSubscriber("go.micro.api.user", service.Server(), new(Event))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}