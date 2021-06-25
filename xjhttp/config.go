package xjhttp

import "gitee.com/xjieinfo/xjgo/xjcore/xjtypes"

type AppConfig struct {
	App           xjtypes.App           //服务的名称等
	Register      xjtypes.Register      //如果服务都注册到ETCD,这个可以不要
	Gateway       xjtypes.Gateway       //如果服务间的调用都通过rpc,这个可以不要
	Config        xjtypes.Config        //如果配置使用本地配置文件, 这个可以不要
	Etcd          xjtypes.Etcd          //http和rpc服务的注册管理
	Auth          xjtypes.Auth          //jwt验证的密钥和有效时间
	Datasource    xjtypes.Datasource    //数据源
	Redis         xjtypes.Redis         //缓存
	ElasticSearch xjtypes.ElasticSearch //es搜索
	Nsq           xjtypes.Nsq           //NSQ消息
}
