package xjcrypto

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	md5data := md5.Sum([]byte(str)) //校验和
	md5str := fmt.Sprintf("%x", md5data)
	return md5str
}
