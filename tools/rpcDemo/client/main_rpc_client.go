package main

import (
	"NewsSpider/tools/rpcDemo"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	// rpc 包装 conn
	client := jsonrpc.NewClient(conn)

	var result float64

	err = client.Call("DemoService.Div", rpcDemo.Params{10, 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(result, err)

	err = client.Call("DemoService.Div", rpcDemo.Params{10, 0}, &result)
	fmt.Println(result, err)

	err = client.Call("DemoService.Div", rpcDemo.Params{4, 2}, &result)
	fmt.Println(result, err)
}
