package main

import (
	"fmt"

	cron "github.com/robfig/cron/v3"
)

//Run a cron worker to make http request every 1m and output it to the file

func main() {
	scheduler := cron.New()
	scheduler.AddFunc("@every 1m", func() {
		fmt.Println("Making http request")
	})
	scheduler.AddFunc("@every 5m", func() {
		fmt.Println("Sending email")
	})
	fmt.Println("Running scheduler")
	scheduler.Run()
}
