package mybroker

import (
	"bufio"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	//"github.com/micro/go-plugins/broker/kafka/v2"
	"log"
	"os"
)

func II2_2() {
	var BrokerURLs = []string{
		0: "192.168.137.8:9092",
		1: "192.168.137.9:9092",
		2: "192.168.137.10:9092",
	}
	micro.NewService(
		micro.Name("abc"),
		micro.Version("0.0.2"),
		micro.Metadata(map[string]string{
			"type": "XXXXX",
		}),
		micro.Broker(kafka.NewBroker(func(o *broker.Options) {
			o.Addrs = BrokerURLs
		})),
	)
	if err := broker.Connect(); err != nil {
		log.Fatal(err.Error())
	}
	broker.Subscribe("abcabc", func(p broker.Event) error {
		brokerHeader := p.Message().Header
		aaa := brokerHeader["AAA"]
		bbb := string(p.Message().Body)
		fmt.Println(aaa, bbb)
		return nil
	})
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
