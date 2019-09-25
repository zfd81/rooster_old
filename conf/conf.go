package conf

import (
	"github.com/BurntSushi/toml"
	"os/user"
)

type Config struct {
	Path string `toml:"path"`
	Http Http   `toml:"http"`
	Etcd Etcd   `toml:"etcd"`
	Meta Meta   `toml:"meta"`
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

func (c *Config) Load(confFile string) error {
	_, err := toml.DecodeFile(confFile, c)
	return err
}

var defaultConf = Config{
	Path: "/rooster/meta",
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
}

var globalConf = defaultConf

func init() {
	user, _ := user.Current()
	defaultConf.Path = user.HomeDir + "/rooster/meta"
}

func NewConfig() *Config {
	conf := defaultConf
	return &conf
}

func GetGlobalConfig() *Config {
	return &globalConf
}
