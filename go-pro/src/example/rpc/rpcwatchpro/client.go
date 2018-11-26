package main

import (
	"collections/go-pro/src/example/rpc/rpcwatchpro/api"
	"fmt"
	"log"
	"time"
)

func main() {
	client, err := api.DialWatchService("tcp", "localhost:1112")
	if err != nil {
		log.Fatal("Dialing...", err)
	}

	DoJob(client)
}
func DoJob(client *api.WatchServiceClient) {
	go func() {
		var keyChanged string
		// 没写api接口封装的时候要这样写：
		// err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		err := client.Watch(30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch: ", keyChanged)
	}()

	// 没写api接口封装的时候要这样写：
	// err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-values"}, new(struct{}))
	err := client.Set([2]string{"abc", "abc-value"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)
}
