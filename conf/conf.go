package conf

import (
	"github.com/BurntSushi/toml"
	"os/user"
)

type Config struct {
	Name    string  `toml:"name"`
	Version string  `toml:"version"`
	Path    string  `toml:"path"`
	Http    Http    `toml:"http"`
	Etcd    Etcd    `toml:"etcd"`
	Meta    Meta    `toml:"meta"`
	Cluster Cluster `toml:"cluster"`
	Data    Data    `toml:"data"`
}

type Http struct {
	Port int `toml:"port"`
}

type Etcd struct {
	Endpoints      []string `toml:"endpoints"`
	DialTimeout    int      `toml:"dial-timeout"`
	RequestTimeout int      `toml:"request-timeout"`
}

type Meta struct {
	InstanceSuffix string `toml:"instance-suffix"`
	DatabaseSuffix string `toml:"database-suffix"`
	TableSuffix    string `toml:"table-suffix"`
	Root           string `toml:"root"`
}

type Cluster struct {
	Root                     string `toml:"root"`
	HeartbeatInterval        int64  `toml:"heartbeat-interval"`
	HeartbeatRecheckInterval int64  `toml:"heartbeat-recheck-interval"`
}

type Data struct {
	BlockSize   int `toml:"block-size"`  //数据块大小，以行为单位
	Replication int `toml:"replication"` //数据块副本数量
}

func (c *Config) Load(confFile string) error {
	_, err := toml.DecodeFile(confFile, c)
	return err
}

var defaultConf = Config{
	Name:    "Rooster",
	Version: "1.0.0",
	Path:    "/rooster/meta",
	Http: Http{
		Port: 8143,
	},
	Etcd: Etcd{
		Endpoints:      []string{"127.0.0.1:2379"},
		DialTimeout:    5,
		RequestTimeout: 5,
	},
	Meta: Meta{
		InstanceSuffix: ".ins",
		DatabaseSuffix: ".db",
		TableSuffix:    ".tbl",
		Root:           "/rooster/meta",
	},
	Cluster: Cluster{
		Root:              "/rooster/cluster",
		HeartbeatInterval: 9, //5秒
	},
	Data: Data{
		BlockSize:   100000,
		Replication: 3,
	},
}

var globalConf = defaultConf

func init() {
	user, _ := user.Current()
	defaultConf.Path = user.HomeDir + "/rooster"
}

func NewConfig() *Config {
	conf := defaultConf
	return &conf
}

func GetGlobalConfig() *Config {
	return &globalConf
}
