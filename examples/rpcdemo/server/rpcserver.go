package main

import (
	"github.com/xjieinfo/xjgo/xjrpc"
)

func main() {
	//创建xjrpc服务器
	s := xjrpc.NewServer()
	//服务发现配置
	addr := "127.0.0.1:12002"
	etcdAddr := []string{"127.0.0.1:2379"}
	basePath := "/rpcdemo"
	s.EtcdRegister(xjrpc.EtcdRegister{
		EtcdServers: etcdAddr, //etcd地址
		BasePath:    basePath, //服务基础地址
		LeaseTime:   50,       //etcd保留时长（秒）
		RpcAddress:  addr,     //rpc服务地址
	})

	// 创建计算工具服务
	remoteMathUtil := new(RemoteMathUtil)
	//将服务注册到rpc
	s.RegisterName("RemoteMathUtil", remoteMathUtil, "")
	//启动xjrpc服务器
	s.Serve("tcp", addr)
}
