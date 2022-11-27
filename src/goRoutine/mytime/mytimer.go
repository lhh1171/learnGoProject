package mytime

import (
	"fmt"
	"time"
)

func MyTimer() {
	timer1 := time.NewTimer(time.Second * 1)
	//相当于sleep
	c := <-timer1.C
	//能够打印时间
	fmt.Println("时间是", c)
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		//相当于sleep
		<-timer2.C
		fmt.Println("Timer 2 expired....")
	}()

	//time.Sleep(4 * time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
