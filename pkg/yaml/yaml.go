package yaml

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Yaml struct {
	Account struct {
		User string `yaml:"user"`
		Password string `yaml:"password"`
	}
}

func ReadYaml(File string) (user string, pass string, err error){
	conf := Yaml{}
	yamlFile, err := ioutil.ReadFile(File)
	if err != nil {
		fmt.Printf("读取配置文件错误：%v ", err)
		return "","",err
	}
	if err = yaml.Unmarshal(yamlFile, &conf); err != nil {
		fmt.Printf("解析配置文件错误: %v", err)
		return "", "", err
	}
	return conf.Account.User, conf.Account.Password, nil
}