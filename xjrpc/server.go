package xjrpc

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
	etcdRegister EtcdRegister
}

func NewServer() *Server {
	return new(Server)
}

func (s *Server) EtcdRegister(etcdRegister EtcdRegister) {
	s.etcdRegister = etcdRegister
}

func (s *Server) RegisterName(name string, rcvr interface{}, metadata string) error {
	// 将对象注册到rpc服务中
	err := rpc.Register(rcvr)
	return err
}

func (s *Server) Serve(network, address string) (err error) {
	//通过该函数把mathUtil中提供的服务注册到HTTP协议上，方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	//在特定的端口进行监听
	listen, err := net.Listen(network, address)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Listening and serving RPC on %s\n", address)
	http.Serve(listen, nil)
	return nil
}
