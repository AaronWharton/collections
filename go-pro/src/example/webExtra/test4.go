package main

import (
	"encoding/json"
	"fmt"
)

type JServer struct {
	// 为保证json正确解析struct成员必须首字母大写，
	// 使用tag自定义输出json的大小写。
	ServerName string `json:"server_name"`
	ServerIP   string `json:"server_ip"`
}

type JsonServer struct {
	JsonServers []JServer `json:"json_servers"`
}

func main() {
	var s JsonServer
	s.JsonServers = append(s.JsonServers, JServer{"beijing_vpn", "127.0.2.3"})
	s.JsonServers = append(s.JsonServers, JServer{"shanghai_vpn", "127.0.3.4"})
	b, err := json.Marshal(&s)
	if err != nil {
		fmt.Println("json error: ", err)
	}
	fmt.Println(string(b))
}
