package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// go解析xml文件
// 熟悉tag的使用
// 解析struct成员大写（导出保证正常解析）

type Recurlyservers struct {
	XMLName     xml.Name  `xml:"servers"`
	Version     string    `xml:"version,attr"`
	Svs         []servers `xml:"servers"`
	Description string    `xml:",innerxml"`
}

type servers struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
