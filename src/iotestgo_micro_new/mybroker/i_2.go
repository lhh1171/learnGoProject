package mybroker

import (
	"bufio"
	"fmt"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"os"
)


func sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		log.Infof("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
func I1_2() {
	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error: %v", err)
	}
	go sub()
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
