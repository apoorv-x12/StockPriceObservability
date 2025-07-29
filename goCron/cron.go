package main

import (
	"fmt"
	"time"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("@every 1m", func() {
		fmt.Println("Running every 1 minute", time.Now())
	})

	c.Start()

	// Block the program
	select {}
}
