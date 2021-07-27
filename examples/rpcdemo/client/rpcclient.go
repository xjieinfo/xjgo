package main

import (
	"fmt"
	"github.com/xjieinfo/xjgo/examples/rpcdemo"
	"github.com/xjieinfo/xjgo/xjrpc"
	"log"
)

func main() {
	//创建rpc客户端
	RpcClient := &xjrpc.Client{
		EtcdEndpoints: []string{"127.0.0.1:2379"},
	}
	var req float64 = 10 //请求参数
	var ret float64      //接收参数
	//rpc服务调用
	err := RpcClient.Call("/rpcdemo", "RemoteMathUtil.CalculateCircleArea", req, &ret)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("area is : %v \n", ret)
	}

	req1 := rpcdemo.Req1{
		I1: 10,
		I2: 20,
		I3: 30,
	}
	var ret1 int
	err = RpcClient.Call("/rpcdemo", "RemoteMathUtil.Sum", req1, &ret1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("sum is : %v \n", ret1)
	}

	var ret2 rpcdemo.Resp1

	err = RpcClient.Call("/rpcdemo", "RemoteMathUtil.Sum2", req1, &ret2)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("sum is : %v \n", ret2)
	}

	err = RpcClient.Call("/rpcdemo", "RemoteMathUtil.Add", 33, &ret2)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("sum is : %v \n", ret2)
	}
}
