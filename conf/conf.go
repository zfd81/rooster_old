package conf

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Name string
	Age  int `toml:"age"`
}

func (c *Config) Load(confFile string) error {
	_, err := toml.DecodeFile(confFile, c)
	return err
}

var defaultConf = Config{"zfd", 33}

var globalConf = defaultConf

func NewConfig() *Config {
	conf := defaultConf
	return &conf
}

func GetGlobalConfig() *Config {
	return &globalConf
}
