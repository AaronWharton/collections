package api

import "net/rpc"

// server

// 为rpc方法设一个唯一独特的名字
const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, response *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// client
type HelloServiceClient struct {
	*rpc.Client
}

func (p *HelloServiceClient) Hello(request string, response *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, response)
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{Client: c}, nil
}
