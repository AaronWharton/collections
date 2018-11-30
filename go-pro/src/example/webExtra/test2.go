package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// go生成xml文件

type server struct {
	ServerName string `xml:"server_name"`
	ServerIP   string `xml:"server_ip"`
}

type host struct {
	HostName string `xml:"host_name"`
	HostIP   string `xml:"host_ip"`
}

type Servers struct {
	XMLName xml.Name `xml:"server"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
	Hosts   []host   `xml:"hosts"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	v.Hosts = append(v.Hosts, host{"localhost", "127.0.0.1"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	if _, err := os.Stdout.Write([]byte(xml.Header)); err != nil {
		log.Println(err)
	}

	if _, err := os.Stdout.Write(output); err != nil {
		log.Println(err)
	}
}
