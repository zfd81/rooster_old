package main

import (
	"flag"
	"fmt"
	"github.com/zfd81/rooster/cluster"
	"github.com/zfd81/rooster/conf"
	"github.com/zfd81/rooster/http"
	"github.com/zfd81/rooster/meta"
	"os"
	"time"
)

const (
	fileName = "config.toml"
)

var (
	Rre    map[string]interface{} = make(map[string]interface{}) // Rooster runtime environment
	config                        = conf.GetGlobalConfig()
	h      bool
	v      bool
	hport  int
)

func usage() {
	fmt.Fprintf(os.Stderr, `%s version: %s
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`, config.Name, config.Version)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&v, "v", false, "版本")
	flag.IntVar(&hport, "hport", -1, "http服务端口")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
	} else if v {
		fmt.Println(config.Version)
	} else {
		if hport != -1 {
			config.Http.Port = hport
		}
		meta.WatchMeta()                    //监测元数据
		meta.LoadMeta()                     //加载元数据
		cluster.Register(time.Now().Unix()) //集群注册
		http.Start()                        //启动服务
	}

}
