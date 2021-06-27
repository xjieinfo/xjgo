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
		return "localhost:12002"
	case "/xjrpc/member":
		return "localhost:12003"
	default:
		return ""
	}
}
