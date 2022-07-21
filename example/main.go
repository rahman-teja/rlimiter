package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rahman-teja/rlimiter"
)

type senderDummy struct{}

func (s senderDummy) Send(msg interface{}) {
	time.Sleep(time.Second * 3)

	log.Println(msg)
}

func main() {
	t0 := time.Now()
	log.Println("Start")

	sender := senderDummy{}

	mg, _ := rlimiter.NewManager(
		rlimiter.NewConfig().
			SetSender(sender).
			SetMaxLimit(5).
			SetWorker(5),
	)

	mg.Start()

	go func() {
		for i := 0; i < 4; i++ {
			mg.Send(fmt.Sprintf("msg << %d", i+1))
		}
	}()

	csignal := make(chan os.Signal, 1)
	signal.Notify(csignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// WAIT FOR IT
	<-csignal

	log.Println("Finish", time.Since(t0))
}
