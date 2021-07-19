package xjrpc

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"math/rand"
	"net/rpc"
	"time"
)

type Client struct {
	EtcdEndpoints []string
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

func (this *Client) GetAddress2(basePath string) string {
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

func (this *Client) GetAddress(basePath string) string {
	path := "/xjrpc" + basePath + "/"
	list := this.EtcdList(path)
	item := list[rand.Int()%len(list)]
	l := len(path)
	if len(item) > l {
		return item[l:]
	} else {
		return ""
	}
}

func (this *Client) EtcdList(key string) []string {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   this.EtcdEndpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return nil
	}

	fmt.Println("connect succ")
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return nil
	}
	list := make([]string, 0)
	for _, ev := range resp.Kvs {
		key := fmt.Sprintf("%s", ev.Key)
		list = append(list, key)
	}
	return list
}
