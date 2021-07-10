package conf

import (
	//"github.com/xjieinfo/xjgo/xjcore/mapping"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetAppConfig(profile string, v interface{}) error {
	filename := "./conf/conf.yml"
	if profile != "" {
		filename = "./conf/conf-" + profile + ".yml"
	}
	content, _ := ioutil.ReadFile(filename)
	//return mapping.UnmarshalYamlBytes(content, v)
	return yaml.Unmarshal(content, v)
}

func GetAppConfigStr(profile string) string {
	filename := "./conf/conf.yml"
	if profile != "" {
		filename = "./conf/conf-" + profile + ".yml"
	}
	data, _ := ioutil.ReadFile(filename)
	return string(data)
}
