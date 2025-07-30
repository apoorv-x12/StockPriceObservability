package main

import (
        "log"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/robfig/cron/v3"
)

func publish() {
fmt.Println("Running every 1 minute", time.Now())
}

func main() {
	log.Println("cron started")
	ctx,cancel:= context.WithCancel(context.Background())
	defer cancel()

	sign:=make(chan os.Signal,1)
	signal.Notify(sign,syscall.SIGTERM,syscall.SIGINT)

	c := cron.New(cron.WithSeconds())

	_,err:= c.AddFunc("@every 1m",func(){
		go publish()
	})
	if err!=nil{
		log.Fatal(err,"this issue")

	}

	c.Start()

	// Block the program
	select {}
}
