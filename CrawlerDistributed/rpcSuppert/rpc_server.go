package rpcSuppert

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//ServerRpc
func ServerRpc(host string, service interface{}) error {
	// rpc 包装服务
	_ = rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}

	return jsonrpc.NewClient(conn), nil
}
