module gitee.com/xjieinfo/xjgo

go 1.16

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/kr/pretty v0.2.0 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/stretchr/testify v1.7.0
	go.etcd.io/etcd v0.0.0-20200402134248-51bdeb39e698
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.11
)

//replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
