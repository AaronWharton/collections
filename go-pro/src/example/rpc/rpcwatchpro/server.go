package main

import (
	"collections/go-pro/src/example/rpc/rpcwatchpro/api"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type WatchService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func (p *WatchService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *WatchService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	k, v := kv[0], kv[1]

	if oldValue := p.m[k]; oldValue != v {
		for _, fn := range p.filter {
			fn(k)
		}
	}

	p.m[k] = v
	return nil
}

func (p *WatchService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	p.mu.Lock()
	p.filter[id] = func(key string) {
		ch <- key
	}
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("time out")
	case key := <-ch:
		*keyChanged = key
		fmt.Println("value has been changed!")
		return nil
	}

	return nil
}

func main() {
	// 不使用这种方式注册rpc，因为有些rpc要使用成员变量，
	// 使用第一种方式注册的会因为成员变量未初始化而crash。
	// api.RegisterWatchService(new(WatchService))
	api.RegisterWatchService(NewWatchService())

	listener, err := net.Listen("tcp", ":1112")
	if err != nil {
		log.Fatal("Listener TCP error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}

}

// 初始化成员变量
func NewWatchService() *WatchService {
	return &WatchService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
		mu:     sync.Mutex{},
	}
}
