package xjregister

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/xjieinfo/xjgo/xjcore/xjtypes"
	"go.etcd.io/etcd/clientv3"
	"log"
	"os"
	"strconv"
	"time"
)

//创建租约注册服务
type ServiceReg struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	canclefunc    func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
}

type Val struct {
	StartTime  string
	UpdateTime string
	Status     int    //状态:1正常 2下线
	Weight     int    //权重:默认为1
	Metatata   string //元数据
	Health     int    //健康状态:1健康 2不健康 3更不健康 ,数值越大越不健康
}

func NewServiceRegister(etcd xjtypes.Etcd) *ServiceReg {
	return NewServiceReg(etcd.Addrs, int64(etcd.LeaseTime))
}
func NewServiceReg(addr []string, timeNum int64) *ServiceReg {
	log.Println("开始连接etcd...")
	var err error
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}

	var (
		client *clientv3.Client
	)

	if clientTem, err := clientv3.New(conf); err == nil {
		client = clientTem
	} else {
		log.Println("连接etcd出错了")
		log.Println(err)
		os.Exit(1)
	}

	ser := &ServiceReg{
		client: client,
	}

	if err := ser.setLease(timeNum); err != nil {
		log.Println("连接etcd出错了")
		log.Println(err)
		os.Exit(1)
	}
	go ser.ListenLeaseRespChan()
	if err != nil {
		log.Println("连接etcd出错了")
		log.Println(err)
		os.Exit(1)
	}
	log.Println("连接etcd...OK")
	return ser
}

//设置租约
func (this *ServiceReg) setLease(timeNum int64) error {
	lease := clientv3.NewLease(this.client)

	//设置租约时间
	leaseResp, err := lease.Grant(context.TODO(), timeNum)
	if err != nil {
		return err
	}

	//设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)

	if err != nil {
		return err
	}

	this.lease = lease
	this.leaseResp = leaseResp
	this.canclefunc = cancelFunc
	this.keepAliveChan = leaseRespChan
	return nil
}

//监听 续租情况
func (this *ServiceReg) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-this.keepAliveChan:
			if leaseKeepResp == nil {
				fmt.Printf("已经关闭续租功能\n")
				return
			} else {
				//fmt.Printf("续租成功\n")
			}
		}
	}
}

func (this *ServiceReg) RegisterService(app xjtypes.App) error {
	return this.PutService(app.Name+"/"+app.Url+":"+strconv.Itoa(int(app.Port)), "")
}

//通过租约 注册服务
func (this *ServiceReg) PutService(key, val string) error {
	kv := clientv3.NewKV(this.client)
	key = "services/" + key
	if val == "" {
		value := Val{
			StartTime:  time.Now().Format("2006-01-02 15:04:05"),
			UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
			Status:     1,
			Weight:     1,
			Metatata:   "",
			Health:     1,
		}
		data, _ := json.Marshal(&value)
		val = string(data)
	}
	_, err := kv.Put(context.TODO(), key, val, clientv3.WithLease(this.leaseResp.ID))
	return err
}

//撤销租约
func (this *ServiceReg) RevokeLease() error {
	this.canclefunc()
	time.Sleep(2 * time.Second)
	_, err := this.lease.Revoke(context.TODO(), this.leaseResp.ID)
	return err
}

//func main() {
//	//ser,_ := NewServiceReg([]string{"192.168.147.151:2379"},5)
//	ser,_ := NewServiceReg([]string{"127.0.0.1:2379"},5)
//	ser.PutService("config","192.168.1.101:8001")
//	ser.PutService("config","192.168.1.102:8001")
//	select{}
//}
