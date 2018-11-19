package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 解析json

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	// 解析到结构体，解析前就知道数据类型
	var s ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	// 解析到interface，解析前不知道数据类型。
	// 可通过interface{}+type assertion进行类型推导
	b := []byte(`{"Name":"Aaron", "Age":21, "RoomMate":["Allen", "Jack"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Println("Parse error: ", err)
	}

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, " is string ", vv)
		case float64:
			fmt.Println(k, " is int ", vv)
		case bool:
			fmt.Println(k, " is bool ", vv)
		case []interface{}:
			fmt.Println(k, " is an array: ")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, " is an unknown type ", vv)
		}
	}
}
