# golang如何实现rpc服务（基于服务发现）

golang对rpc服务提供了很好地支持，在原生rpc服务的基础上进行简单封装，再加上服务发现功能，就能实现基本的rpc服务功能，可以在此基础上进行优化和完善，以供http服务访问，具体实现如下：

# 编写服务
## 建立一个用于数学计算的struct：RemoteMathUtil
```go
// MathUtil 用于数学计算
type RemoteMathUtil struct{}

// CaculateCircleArea 计算圆的面积
func (m *RemoteMathUtil) CalculateCircleArea(req float64, resp *float64) error {
	*resp = math.Pi * req * req
	return nil
}
```
# 服务启动
创建rpc服务，并将etcd来管理服务发现，代码实现如下，有详细注释：
```go
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
		EtcdServers:    etcdAddr, //etcd地址
		BasePath:       basePath, //服务基础地址
		LeaseTime:      50,  //etcd保留时长（秒）
		RpcAddress:     addr, //rpc服务地址
	})

	// 创建计算工具服务
	remoteMathUtil := new(RemoteMathUtil)
	//将服务注册到rpc
	s.RegisterName("RemoteMathUtil", remoteMathUtil, "")
	//启动xjrpc服务器
	s.Serve("tcp", addr)
}
```
## 服务端运行，显示如下：
```
2021/07/27 17:24:27 开始连接etcd...
2021/07/27 17:24:27 连接etcd...OK
Listening and serving RPC on 127.0.0.1:12002
```
# 客户访问
创建rpc客户端，进行rpc服务调用（通过etcd进行服务发现)，具体代码如下：
```go
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
	var ret float64 //接收参数
	//rpc服务调用
	err := RpcClient.Call("/rpcdemo", "RemoteMathUtil.CalculateCircleArea", req, &ret)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("area is : %v \n", ret)
	}	
}
```

## 客户端运行，显示如下：
```
area is : 314.1592653589793 
```

完整的源代码整理如下：[https://github.com/xjieinfo/xjgo/tree/main/examples/rpcdemo](https://github.com/xjieinfo/xjgo/tree/main/examples/rpcdemo)

本程序用到的工具为：[https://github.com/xjieinfo/xjgo](https://github.com/xjieinfo/xjgo)

如果此项目对你有所帮助或启发，请给个star支持一下，谢谢！

