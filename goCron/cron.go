package main

import (
	"fmt"
	"time"
	"github.com/robfig/cron/v3"
)

func publish() {
fmt.Println("Running every 1 minute", time.Now())
}

func main() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("@every 1m",func(){
		go publish()
	})

	c.Start()

	// Block the program
	select {}
}
