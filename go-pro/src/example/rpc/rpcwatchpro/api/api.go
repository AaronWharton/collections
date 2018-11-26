package api

import (
	"net/rpc"
)

/*服务端的简单封装*/

// 独特的方法用唯一的路径名区分
const WatchServiceName string = "path/to/pkg.WatchServiceName"

// 实现的方法列表
type WatchServiceInterface = interface {
	Get(key string, value *string) error
	Set(kv [2]string, reply *struct{}) error
	Watch(timeoutSecond int, keyChanged *string) error
}

// 具体到一个方法的注册函数
func RegisterWatchService(service WatchServiceInterface) error {
	return rpc.RegisterName(WatchServiceName, service)
}

/*客户端的简单封装*/

// 具体的一个Client封装（这里专指Watch这个rpc方法所封装）
type WatchServiceClient struct {
	*rpc.Client
}

// ???
var _ WatchServiceInterface = (*WatchServiceClient)(nil)

// 具体一个rpc方法调用的封装
func DialWatchService(network, address string) (*WatchServiceClient, error) {
	conn, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &WatchServiceClient{Client: conn}, nil
}

// 对具体方法的封装
func (p *WatchServiceClient) Get(key string, value *string) error {
	return p.Client.Call(WatchServiceName+".Get", key, value)
}

func (p *WatchServiceClient) Set(kv [2]string, reply *struct{}) error {
	return p.Client.Call(WatchServiceName+".Set", kv, reply)
}

func (p *WatchServiceClient) Watch(timeoutSecond int, keyChanged *string) error {
	return p.Client.Call(WatchServiceName+".Watch", timeoutSecond, keyChanged)
}
