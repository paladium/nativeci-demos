package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	cron "github.com/robfig/cron/v3"
)

//Run a cron worker to make http request every 1m and output it to the file
//Show to use env variables

func getRocketData() (*string, error) {
	resp, err := http.Get("https://api.spacexdata.com/v4/launches/latest")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(body)
	return &bodyString, nil
}

func main() {
	scheduler := cron.New()
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scheduler.AddFunc("@every 1m", func() {
		fmt.Println("Writing to file")
		rockets, err := getRocketData()
		if err != nil {
			return
		}
		file.WriteString(*rockets + "\n")
	})
	fmt.Println("Running scheduler")
	scheduler.Run()
}
