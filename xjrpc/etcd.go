package xjrpc

import (
	"sync"
	"time"

	"github.com/docker/libkv/store"
	metrics "github.com/rcrowley/go-metrics"
	//etcd "github.com/smallnest/libkv-etcdv3-store"
)

var (
	EtcdServers []string
)

type EtcdRegister struct {
	// service address, for example, tcp@127.0.0.1:8972, quic@127.0.0.1:1234
	ServiceAddress string
	// etcd addresses
	EtcdServers []string
	// base path for rpcx server, for example com/example/rpcx
	BasePath string
	Metrics  metrics.Registry
	// Registered services
	Services       []string
	metasLock      sync.RWMutex
	metas          map[string]string
	UpdateInterval time.Duration

	Options *store.Config
	kv      store.Store

	dying chan struct{}
	done  chan struct{}
}
