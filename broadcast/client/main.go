package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/johanna-hub/graduation/broadcast/client/pusher"
)

//type instructions struct {
//direction string
//mode      string
//}

func main() {
	options := pusher.Options{
		Cluster:       "eu",
		Client:        "go",
		ClientVersion: "0.1",
		APIVersion:    "7",
		Protocol:      "wss",
		AppId:         "c24dabd6884e70c4eafb",
	}
	pusherURL := pusher.NewURL(options)
	//fmt.Println(pusherURL)

	pusherConnect := pusher.Connection{}
	err := pusherConnect.Connect(pusherURL, "http://localhost/")
	if err != nil {
		fmt.Println(err)
	}

	events, err := pusherConnect.Subscribe("my-channel")
	fmt.Println(events)
	if err != nil {
		fmt.Println(err)
	}

	for m := range events {
		resp, err := http.Post("http://172.16.14.244:7001/input", "application/json", bytes.NewBufferString(m.Data))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		fmt.Println(resp)
		//ins := instructions{}
		//_ = json.Unmarshal(m.Event, ins)
		//fmt.Println(ins)
	}
}

//func main() {
/*
	origin := "http://localhost/"

	options := pusher.Options{
		Cluster:       "eu",
		Client:        "go",
		ClientVersion: "0.1",
		APIVersion:    "7",
		Protocol:      "wss",
		AppId:         "c24dabd6884e70c4eafb",
	}
*/
//}
