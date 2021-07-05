package main

import (
	"fmt"
	"github.com/xjieinfo/xjgo/examples/rpcdemo"
	"github.com/xjieinfo/xjgo/xjrpc"
	"log"
	"net/rpc"
)

type ClientMathUtil struct{}

// CaculateCircleArea 计算圆的面积
func (m *ClientMathUtil) CaculateCircleArea(req float64) (float64, error) {
	client, err := rpc.DialHTTP("tcp", "localhost:10102")
	if err != nil {
		panic(err.Error())
	}
	//var req float64 //请求值
	//req = 3
	//
	var resp *float64 //返回值
	err = client.Call("RemoteMathUtil.CalculateCircleArea", req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)
	return *resp, nil
}

func main2() {
	area, err := new(ClientMathUtil).CaculateCircleArea(10)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("area is : %f \n", area)
	}

}

func main() {
	var req float64 = 10
	var ret float64
	err := new(xjrpc.Client).Call("/xjrpc/admin", "RemoteMathUtil.CalculateCircleArea", req, &ret)
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
	err = new(xjrpc.Client).Call("/xjrpc/admin", "RemoteMathUtil.Sum", req1, &ret1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("sum is : %v \n", ret1)
	}

	var ret2 rpcdemo.Resp1

	err = new(xjrpc.Client).Call("/xjrpc/admin", "RemoteMathUtil.Sum2", req1, &ret2)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("sum is : %v \n", ret2)
	}

	err = new(xjrpc.Client).Call("/xjrpc/admin", "RemoteMathUtil.Add", 33, &ret2)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("sum is : %v \n", ret2)
	}
}
