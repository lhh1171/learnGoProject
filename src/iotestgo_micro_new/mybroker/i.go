package mybroker

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/config/cmd"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	topic = "go.micro.learning.topic.log"
	b     broker.Broker
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		log.Infof(broker.String())
		if err := broker.Publish(topic, msg); err != nil {
			log.Infof("[pub] Message publication failed: %v", err)
		} else {
			fmt.Println("[pub] Message published: ", string(msg.Body))
		}
		i++
	}
}
func I1() {

	cmd.Init()
	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error: %v", err)
	}

	go pub()
	<-time.After(time.Second * 20)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
