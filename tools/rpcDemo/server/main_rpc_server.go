package main

import (
	"NewsSpider/tools/rpcDemo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// rpc 包装服务
	_ = rpc.Register(rpcDemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
