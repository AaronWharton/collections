// 客户端
package main

import (
	"collections/go-pro/src/example/rpcpro/api"
	"fmt"
	"log"
)

func main() {
	client, err := api.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing...: ", err)
	}

	var response string
	err = client.Hello("Aaron", &response)
	if err != nil {
		log.Fatal(err)
	}

	// 打印输出
	fmt.Println(response)
}
