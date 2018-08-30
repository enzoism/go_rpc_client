package main

import (
	"net/rpc"
	"fmt"
)

type Args struct {
	A, B int
}

func main() {
	fmt.Println("-------------请先开启服务端")
	//连接服务器端，创建一个client
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	args := &Args{7, 8}
	var reply int
	//通过Call方法调用Arith类型的Multiply方法，注意形参
	client.Call("Arith.Multiply", args, &reply)
	//得到调用结果，输出Arith: 7*8=56
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}