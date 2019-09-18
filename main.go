package main

import (
	"github.com/zfd81/rooster/conf"
	"github.com/zfd81/rooster/meta"
	"os"
)

const (
	fileName = "config.toml"
)

// Rooster runtime environment
var (
	Rre    map[string]interface{} = make(map[string]interface{})
	config                        = conf.NewConfig()
)

func main() {
	_, err := os.Stat(fileName) //os.Stat获取文件信息
	if err == nil {             //判断配置文件是否存在
		config.Load(fileName) //加载配置文件信息
	}

	meta.ReadMeta(config.Path)
	ins := meta.Instance{}
	ins.Name = "hello"
}
