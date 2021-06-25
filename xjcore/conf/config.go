package conf

import (
	"gitee.com/xjieinfo/xjgo/xjcore/mapping"
	"io/ioutil"
)

func GetAppConfig(profile string, v interface{}) error {
	filename := "./conf/conf.yml"
	if profile != "" {
		filename = "./conf/conf-" + profile + ".yml"
	}
	content, _ := ioutil.ReadFile(filename)
	return mapping.UnmarshalYamlBytes(content, v)
}

func GetAppConfigStr(profile string) string {
	filename := "./conf/conf.yml"
	if profile != "" {
		filename = "./conf/conf-" + profile + ".yml"
	}
	data, _ := ioutil.ReadFile(filename)
	return string(data)
}
