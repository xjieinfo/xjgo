package main

import (
	"gitee.com/xjieinfo/xjgo/xjrpc"
	"github.com/rcrowley/go-metrics"
	"time"
)

func main() {
	s := xjrpc.NewServer()

	addr := "127.0.0.1:12002"
	etcdAddr := []string{"127.0.0.1:2379"}
	basePath := "/xjrpc/admin"
	s.EtcdRegister(xjrpc.EtcdRegister{
		ServiceAddress: "tcp@" + addr,
		EtcdServers:    etcdAddr,
		BasePath:       basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	})

	// 创建计算实例
	remoteMathUtil := new(RemoteMathUtil)
	s.RegisterName("RemoteMathUtil", remoteMathUtil, "")

	s.Serve("tcp", addr)
}
