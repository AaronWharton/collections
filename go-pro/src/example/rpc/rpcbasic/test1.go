// 注册RPC服务
package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

// go语言rpc规则：方法只能有两个可序列化的参数，第二个必须为指针，并返回一个error
func (p *HelloService) Hello(request string, response *string) error {
	*response = "Hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	// 放在for循环里面循环获取请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}

}
