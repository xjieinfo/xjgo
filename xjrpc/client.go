package xjrpc

import (
	"net/rpc"
)

type Client struct {
}

func (this *Client) Call(basePath, servicePath string, args interface{}, reply interface{}) (err error) {
	address := this.GetAddress(basePath)
	client, err := rpc.DialHTTP("tcp", address)
	defer client.Close()
	if err != nil {
		panic(err.Error())
	}
	err = client.Call(servicePath, args, reply)
	return
}

func (this *Client) GetAddress(basePath string) string {
	switch basePath {
	case "/xjrpc/admin":
		return "127.0.0.1:12002"
	case "/xjrpc/member":
		return "127.0.0.1:12003"
	case "/xjrpc/product":
		return "127.0.0.1:12004"
	case "/xjrpc/marketing":
		return "127.0.0.1:12008"
	default:
		return ""
	}
}
