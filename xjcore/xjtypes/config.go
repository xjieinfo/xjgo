package xjtypes

//服务的名称等
type App struct {
	Name    string
	Url     string
	Port    int
	Rpcport int
}

//如果服务都注册到ETCD,这个可以不要
type Register struct {
	Url  string
	Port int
}

//如果服务间的调用都通过rpc,这个可以不要
type Gateway struct {
	Url  string
	Port int
}

//如果配置使用本地配置文件, 这个可以不要
type Config struct {
	Url  string
	Port int
}

//http和rpc服务的注册管理
type Etcd struct {
	Addrs     []string
	LeaseTime int
}

//jwt验证的密钥和有效时间
type Auth struct {
	AccessSecret string
	AccessExpire int
}

//数据源
type Datasource struct {
	Drivername string
	Dsn        string
}

//缓存
type Redis struct {
	Addr     string
	Password string
	Db       int
}

//es搜索
type ElasticSearch struct {
	Addr     string
	Username string
	Password string
}

//NSQ消息
type Nsq struct {
	Nsqds       []string
	Nsqlookupds []string
}
